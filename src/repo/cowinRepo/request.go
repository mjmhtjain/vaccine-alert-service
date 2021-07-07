package cowinrepo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

type CowinAPI interface {
	AppointmentSessionByDistrictAndCalendar(districtId string, date string) (*model.Appointments, error)
}

type CowinAPIImpl struct {
	httpClient *http.Client
}

func NewCowinAPI() CowinAPI {
	return &CowinAPIImpl{
		httpClient: &http.Client{},
	}
}

func (cowin *CowinAPIImpl) AppointmentSessionByDistrictAndCalendar(districtId string, date string) (*model.Appointments, error) {
	logger.INFO.Printf("requestAppointmentsFromCentres: date: %v\n", date)
	var appointmentResp *model.Appointments = new(model.Appointments)

	req, err := http.NewRequest("GET", util.URL_AppointmentSessionForWeek, nil)
	if err != nil {
		logger.ERROR.Printf("Invalid Http Request \n %v \n", err)

	}

	q := req.URL.Query()
	q.Add("district_id", fmt.Sprint(districtId))
	q.Add("date", date)
	req.URL.RawQuery = q.Encode()

	resp, err := cowin.httpClient.Do(req)

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
