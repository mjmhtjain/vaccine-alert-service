package cowin

import (
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
)

type AppointmentService interface {
	FetchVaccineAppointments(stateId string, date string) ([]model.AppointmentSessionORM, error)
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

func (service *AppointmentServiceImpl) FetchVaccineAppointments(stateName string, date string) ([]model.AppointmentSessionORM, error) {
	logger.INFO.Printf("FetchVaccineAppointments stateId: %v date: %v \n", stateName, date)
	var (
		stateId              string
		err                  error
		districts            *model.StateDistricts
		districtAppointments []model.Appointments
	)

	stateId, err = service.fetchStateId(stateName)
	if err != nil {
		return nil, err
	}
	districts, err = service.fetchDistricts(stateId)
	if err != nil {
		return nil, err
	}

	districtAppointments, err = service.requestAppointmentsFromCentres(districts, date)
	if err != nil {
		panic(err)
	}

	filteredAppointments := service.filterAppointments(districtAppointments)

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
	districts *model.StateDistricts,
	date string,
) ([]model.Appointments, error) {
	logger.DEBUG.Printf("requestAppointmentsFromCentres: date: %v\n", date)

	var appoitments []model.Appointments
	resChan := make(chan model.CowinAppointmentResponse)
	defer close(resChan)

	districtCount := len(districts.Districts)
	workerCount := 5
	districtChan := make(chan string)
	defer close(districtChan)

	// start workers
	for i := 0; i < workerCount; i++ {
		go appService.requestWorker(date, resChan, districtChan)
	}

	// send all the data
	go func() {
		for _, d := range districts.Districts {
			districtChan <- fmt.Sprint(d.DistrictID)
		}
	}()

	// collect data
	// TODO: need to think about timeouts
	for i := 0; i < districtCount; i++ {
		res := <-resChan

		if res.Err != nil {
			logger.ERROR.Printf("%v\n", res.Err)
		} else {
			appoitments = append(appoitments, res.AppointmentData)
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

func (service *AppointmentServiceImpl) filterAppointments(districtAppointments []model.Appointments) []model.AppointmentSessionORM {
	var (
		appSessArr       = []model.AppointmentSessionORM{}
		noRecordExistErr *customerror.NoRecordExists
	)

	for _, distApp := range districtAppointments {
		for _, center := range distApp.Centers {
			for _, sess := range center.Sessions {

				_, err := service.sqlRepo.FindSessionWithSessionId(&sess)
				if err != nil && errors.As(err, &noRecordExistErr) {

					// find/insert center info
					centerOrm, err := service.sqlRepo.FindCenterWithCenterId(center)
					if err != nil && errors.As(err, &noRecordExistErr) {
						centerOrm = service.sqlRepo.InsertCenterInfo(center)
					}

					// find/insert vaccine info
					vaccinOrm, err := service.sqlRepo.FindVaccineByName(sess.Vaccine)
					if err != nil && errors.As(err, &noRecordExistErr) {
						vaccinOrm = service.sqlRepo.InsertVaccine(sess.Vaccine)
					}

					// insert session info
					sessionOrm := service.sqlRepo.InsertAppointmentSession(&sess, centerOrm.Id, vaccinOrm.Id)

					appSessArr = append(appSessArr, *sessionOrm)
				}
			}
		}
	}

	return appSessArr
}

// func generateAppointmentSession(sess model.Session) model.AppointmentSession {
// 	appSess := model.AppointmentSession{
// 		CenterIDFK:             -1,
// 		SessionID:              sess.SessionID,
// 		Date:                   sess.Date,
// 		AvailableCapacity:      sess.AvailableCapacity,
// 		MinAgeLimit:            sess.MinAgeLimit,
// 		VaccineIDKF:            "-1",
// 		AvailableCapacityDose1: sess.AvailableCapacityDose1,
// 		AvailableCapacityDose2: sess.AvailableCapacityDose2,
// 	}
// 	return appSess
// }
