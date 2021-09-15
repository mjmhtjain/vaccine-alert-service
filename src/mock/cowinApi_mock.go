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
}
