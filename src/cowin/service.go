package cowin

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
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
	fileData, err := readStaticFile("districts.json")
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

func requestAppointmentsFromCentres(districts *model.StateDistricts, date string) (*model.Appointments, error) {
	logger.INFO.Printf("requestAppointmentsFromCentres: date: %v\n", date)

	var appointmentResp *model.Appointments = new(model.Appointments)
	d := districts.Districts[0]
	districtId := d.DistrictID
	url := fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?district_id=%v&date=%v", districtId, date)

	resp, err := http.Get(url)
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

func readStaticFile(fileName string) ([]byte, error) {
	logger.INFO.Printf("readStaticFile: fileName: %v \n", fileName)

	out, err := exec.Command("pwd").Output()
	if err != nil {
		logger.ERROR.Printf("Could not execute command.. \n %v \n", err)
		return nil, err
	}

	basePath := string(out)
	basePath = basePath[:len(basePath)-1]
	filename := filepath.Join(basePath, "src", "staticData", fileName)

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.ERROR.Printf("Error on reading file.. \n %v \n", err)
		return nil, err
	}

	return fileData, nil
}
