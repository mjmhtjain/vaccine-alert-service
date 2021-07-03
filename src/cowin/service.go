package cowin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

type AppointmentService interface {
	FetchVaccineAppointments(stateId string, date string) (*model.Appointments, error)
}

type AppointmentServiceImpl struct {
}

func NewAppointmentService() AppointmentService {
	return &AppointmentServiceImpl{}
}

func (service *AppointmentServiceImpl) FetchVaccineAppointments(stateId string, date string) (*model.Appointments, error) {
	logger.INFO.Printf("FetchVaccineAppointments %v %v \n", stateId, date)

	districts, err := fetchDistricts(stateId)
	if err != nil {
		return nil, err
	}

	resp, err := requestAppointmentsFromCentres(districts, date)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func fetchDistricts(stateId string) (*model.StateDistricts, error) {
	logger.INFO.Printf("fetchDistricts: stateId: %v \n", stateId)

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
func requestAppointmentsFromCentres(districts *model.StateDistricts, date string) (*model.Appointments, error) {
	logger.INFO.Printf("requestAppointmentsFromCentres: date: %v\n", date)

	var appointmentResp *model.Appointments = new(model.Appointments)
	d := districts.Districts[0]
	districtId := d.DistrictID

	req, err := http.NewRequest("GET", util.URL_AppointmentSessionForWeek, nil)
	if err != nil {
		logger.ERROR.Printf("Invalid Http Request \n %v \n", err)

	}

	q := req.URL.Query()
	q.Add("district_id", fmt.Sprint(districtId))
	q.Add("date", date)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		logger.ERROR.Printf("request failed .. \n %v \n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ERROR.Printf("body unreadable .. \n %v \n", err)
		return nil, err
	}

	err = json.Unmarshal(body, appointmentResp)
	if err != nil {
		logger.ERROR.Printf("unmarshalling error .. \n %v \n", err)
		return nil, err
	}

	return appointmentResp, nil
}
