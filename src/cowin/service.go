package cowin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	customerror "github.com/mjmhtjain/vaccine-alert-service/src/customError"
	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	cowinrepo "github.com/mjmhtjain/vaccine-alert-service/src/repo/cowinRepo"
	"github.com/mjmhtjain/vaccine-alert-service/src/repo/sql"
	staticfile "github.com/mjmhtjain/vaccine-alert-service/src/staticFile"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

type AppointmentService interface {
	FetchVaccineAppointments(ctx context.Context, stateName string, date string) ([]model.AppointmentSessionORM, error)
}

type AppointmentServiceImpl struct {
	cowin    cowinrepo.CowinAPI
	staticFS staticfile.FileService
	sqlRepo  sql.SqlRepo
}

func NewAppointmentService() AppointmentService {
	return &AppointmentServiceImpl{
		cowin:    cowinrepo.NewCowinAPI(),
		staticFS: staticfile.NewFileService(),
		sqlRepo:  sql.NewSqlRepo(),
	}
}

func (service *AppointmentServiceImpl) FetchVaccineAppointments(
	ctx context.Context,
	stateName string,
	date string,
) ([]model.AppointmentSessionORM, error) {
	logger.INFO.Printf("FetchVaccineAppointments stateId: %v date: %v \n", stateName, date)
	var (
		stateId              string
		err                  error
		districts            *model.StateDistricts
		districtAppointments []model.Appointments
	)

	// fetch state and districtIds for the state
	stateId, err = service.fetchStateId(stateName)
	util.ErrorPanic(err)

	districts, err = service.fetchDistricts(stateId)
	util.ErrorPanic(err)

	// fetch all appointments for a district
	districtAppointments, err = service.requestAppointmentsFromCentres(ctx, districts, date)
	util.ErrorPanic(err)

	// filter out stale appointments
	filteredAppointments := service.filterAppointments(ctx, districtAppointments)

	return filteredAppointments, nil
}

func (service *AppointmentServiceImpl) fetchStateId(stateName string) (string, error) {
	var data model.States
	fileData, err := service.staticFS.Read("states.json")
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return "", fmt.Errorf("unmarshalling error: %s", err)
	}

	stateMap := make(map[string]string)
	for _, e := range data.States {

		key := strings.ToLower(e.StateName)
		val := fmt.Sprint(e.StateID)

		stateMap[key] = val
	}

	return stateMap[strings.ToLower(stateName)], nil
}

func (service *AppointmentServiceImpl) fetchDistricts(stateId string) (*model.StateDistricts, error) {
	logger.DEBUG.Printf("fetchDistricts: stateId: %v \n", stateId)

	var data model.StateDistricts

	//fetching predetermined districts
	fileData, err := service.staticFS.Read("districts.json")
	if err != nil {
		logger.ERROR.Printf("Error occured while reading file.. \n %v \n", err)
		return nil, err
	}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		logger.ERROR.Printf("Error occured while unmarshalling.. \n %v \n", err)
		return nil, err
	}

	return &data, nil
}

func (appService *AppointmentServiceImpl) requestAppointmentsFromCentres(
	ctx context.Context,
	districts *model.StateDistricts,
	date string,
) ([]model.Appointments, error) {
	logger.DEBUG.Printf("requestAppointmentsFromCentres: date: %v\n", date)

	var appoitments []model.Appointments
	resChan := make(chan model.CowinAppointmentResponse)
	defer close(resChan)

	workerCount := 5
	districtChan := make(chan string)

	// start worker go routines
	for i := 0; i < workerCount; i++ {
		go appService.requestWorker(date, resChan, districtChan)
	}

	// send jobs to worker routines in async manner
	go func() {
		defer close(districtChan) // writer closes channel

		for _, d := range districts.Districts {

			select {
			// case districtChan <- fmt.Sprint(d.DistrictID):
			case <-ctx.Done():
				return // exit go routine if ctx cancels
			default:
				// ctx is not cancelled
			}

			districtChan <- fmt.Sprint(d.DistrictID)
		}
	}()

	// collect output from workers in sync manner
	// TODO: need to think about timeouts
	for i := 0; i < len(districts.Districts); i++ {
		select {
		case res := <-resChan:
			if res.Err != nil {
				logger.ERROR.Printf("%v\n", res.Err)
			} else {
				appoitments = append(appoitments, res.AppointmentData)
			}
		case <-ctx.Done():
			return nil, errors.New("context cancelled")
		}
	}

	if len(appoitments) == 0 {
		return nil, errors.New("no data found")
	}

	return appoitments, nil
}

func (appService *AppointmentServiceImpl) requestWorker(
	date string,
	resChan chan model.CowinAppointmentResponse,
	districtChan chan string,
) {

	for d := range districtChan {
		app, err := appService.cowin.AppointmentSessionByDistrictAndCalendar(d, date)
		cowinRes := model.CowinAppointmentResponse{}
		if err != nil {
			cowinRes.Err = err
		} else {
			cowinRes.AppointmentData = *app
		}

		resChan <- cowinRes
	}

}

func (service *AppointmentServiceImpl) filterAppointments(
	ctx context.Context,
	districtAppointments []model.Appointments,
) []model.AppointmentSessionORM {
	var (
		appSessArr       = []model.AppointmentSessionORM{}
		noRecordExistErr *customerror.NoRecordExists
	)

	for _, distApp := range districtAppointments {
		for _, center := range distApp.Centers {
			for _, sess := range center.Sessions {

				_, err := service.sqlRepo.FindSessionWithSessionId(ctx, &sess)
				if err != nil && errors.As(err, &noRecordExistErr) {

					// find/insert center info
					centerOrm, err := service.sqlRepo.FindCenterWithCenterId(ctx, center)
					if err != nil && errors.As(err, &noRecordExistErr) {
						centerOrm = service.sqlRepo.InsertCenterInfo(ctx, center)
					}

					// find/insert vaccine info
					vaccinOrm, err := service.sqlRepo.FindVaccineByName(ctx, sess.Vaccine)
					if err != nil && errors.As(err, &noRecordExistErr) {
						vaccinOrm = service.sqlRepo.InsertVaccine(ctx, sess.Vaccine)
					}

					// insert session info
					sessionOrm := service.sqlRepo.InsertAppointmentSession(ctx, &sess, centerOrm.Id, vaccinOrm.Id)

					appSessArr = append(appSessArr, *sessionOrm)
				}
			}
		}
	}

	return appSessArr
}
