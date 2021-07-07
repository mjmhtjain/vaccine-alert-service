package cowin

import (
	"encoding/json"
	"fmt"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	cowinrepo "github.com/mjmhtjain/vaccine-alert-service/src/repo/cowinRepo"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

type AppointmentService interface {
	FetchVaccineAppointments(stateId string, date string) (*model.Appointments, error)
}

type AppointmentServiceImpl struct {
	cowin cowinrepo.CowinAPI
}

func NewAppointmentService() AppointmentService {
	return &AppointmentServiceImpl{
		cowin: cowinrepo.NewCowinAPI(),
	}
}

func (service *AppointmentServiceImpl) FetchVaccineAppointments(stateId string, date string) (*model.Appointments, error) {
	logger.INFO.Printf("FetchVaccineAppointments %v %v \n", stateId, date)

	districts, err := fetchDistricts(stateId)
	if err != nil {
		return nil, err
	}

	resp, err := service.requestAppointmentsFromCentres(districts, date)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func fetchDistricts(stateId string) (*model.StateDistricts, error) {
	logger.DEBUG.Printf("fetchDistricts: stateId: %v \n", stateId)

	var data model.StateDistricts

	//fetching predetermined districts
	fileData, err := util.ReadStaticFile("districts.json")

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
func (service *AppointmentServiceImpl) requestAppointmentsFromCentres(districts *model.StateDistricts, date string) (*model.Appointments, error) {
	logger.DEBUG.Printf("requestAppointmentsFromCentres: date: %v\n", date)
	districtId := fmt.Sprint(districts.Districts[0].DistrictID)

	appointments, err := service.cowin.AppointmentSessionByDistrictAndCalendar(districtId, date)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}
