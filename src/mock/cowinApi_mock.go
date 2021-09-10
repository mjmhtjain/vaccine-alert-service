package mock

import (
	"encoding/json"
	"errors"
	"path/filepath"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	cowinrepo "github.com/mjmhtjain/vaccine-alert-service/src/repo/cowinRepo"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

func NewMockCowinAPI(responseFilePath string) cowinrepo.CowinAPI {
	return &CowinAPIMockImpl{
		mockResponseFilePath: responseFilePath,
	}
}

type CowinAPIMockImpl struct {
	mockResponseFilePath string
}

func (mock *CowinAPIMockImpl) AppointmentSessionByDistrictAndCalendar(
	districtId string, date string,
) (*model.Appointments, error) {

	var appointmentData *model.Appointments = new(model.Appointments)

	if mock.mockResponseFilePath == "" {
		return nil, errors.New("")
	}

	path, err := filepath.Abs(mock.mockResponseFilePath) //filepath.Abs("../mock/appointmentSessionMock.json")
	if err != nil {
		logger.ERROR.Printf("Invalid filepath.. \n %v \n", err)
		return nil, err
	}

	data, err := util.Readfile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, appointmentData)
	if err != nil {
		return nil, err
	}

	return appointmentData, nil
}
