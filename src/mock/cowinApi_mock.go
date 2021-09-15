package mock

import (
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	cowinrepo "github.com/mjmhtjain/vaccine-alert-service/src/repo/cowinRepo"
)

func NewMockCowinAPI(setResponse model.Appointments) cowinrepo.CowinAPI {
	return &CowinAPIMockImpl{
		response: setResponse,
	}
}

type CowinAPIMockImpl struct {
	response model.Appointments
}

func (mock *CowinAPIMockImpl) AppointmentSessionByDistrictAndCalendar(
	districtId string, date string,
) (*model.Appointments, error) {
	return &mock.response, nil

	// var appointmentData *model.Appointments = new(model.Appointments)

	// if mock.mockResponseFilePath == "" {
	// 	return nil, errors.New("")
	// }

	// path, err := filepath.Abs(mock.mockResponseFilePath) //filepath.Abs("../mock/appointmentSessionMock.json")
	// if err != nil {
	// 	logger.ERROR.Printf("Invalid filepath.. \n %v \n", err)
	// 	return nil, err
	// }

	// data, err := util.Readfile(path)
	// if err != nil {
	// 	return nil, err
	// }

	// err = json.Unmarshal(data, appointmentData)
	// if err != nil {
	// 	return nil, err
	// }

	// return appointmentData, nil
}
