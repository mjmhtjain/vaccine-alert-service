package cowin

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	cowinrepo "github.com/mjmhtjain/vaccine-alert-service/src/repo/cowinRepo"
	staticfile "github.com/mjmhtjain/vaccine-alert-service/src/staticFile"
)

type AppointmentService interface {
	FetchVaccineAppointments(stateId string, date string) ([]model.Appointments, error)
}

type AppointmentServiceImpl struct {
	cowin    cowinrepo.CowinAPI
	staticFS staticfile.FileService
}

func NewAppointmentService() AppointmentService {
	return &AppointmentServiceImpl{
		cowin:    cowinrepo.NewCowinAPI(),
		staticFS: staticfile.NewFileService(),
	}
}

func (service *AppointmentServiceImpl) FetchVaccineAppointments(stateName string, date string) ([]model.Appointments, error) {
	logger.INFO.Printf("FetchVaccineAppointments stateId: %v date: %v \n", stateName, date)

	stateId, err := service.fetchStateId(stateName)
	if err != nil {
		return nil, err
	}
	districts, err := service.fetchDistricts(stateId)
	if err != nil {
		return nil, err
	}

	resp, err := service.requestAppointmentsFromCentres(districts, date)
	if err != nil {
		return nil, err
	}

	return resp, nil
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

// TODO: make parallel calls for each district
func (service *AppointmentServiceImpl) requestAppointmentsFromCentres(districts *model.StateDistricts, date string) ([]model.Appointments, error) {
	logger.DEBUG.Printf("requestAppointmentsFromCentres: date: %v\n", date)

	var appoitments []model.Appointments

	for _, d := range districts.Districts {
		districtId := fmt.Sprint(d.DistrictID)
		app, err := service.cowin.AppointmentSessionByDistrictAndCalendar(districtId, date)
		if err != nil {
			return nil, err
		}

		appoitments = append(appoitments, *app)
	}

	return appoitments, nil
}
