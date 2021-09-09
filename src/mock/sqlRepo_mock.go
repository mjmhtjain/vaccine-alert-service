package mock

import (
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	"github.com/mjmhtjain/vaccine-alert-service/src/repo/sql"
)

type MockSqlRepoImpl struct {
	isSessionExistResp bool
}

func NewMockSqlRepoImpl(isSessionExist bool) sql.SqlRepo {
	return &MockSqlRepoImpl{
		isSessionExistResp: isSessionExist,
	}
}

func (impl *MockSqlRepoImpl) IsSessionExist(sess model.Session) bool {
	return impl.isSessionExistResp
}

func (impl *MockSqlRepoImpl) InsertCenterInfo(resp []model.Appointments, centerIndex int) {

}

func (impl *MockSqlRepoImpl) InsertVaccine(sess model.Session) (string, error) {
	return "", nil
}

func (impl *MockSqlRepoImpl) InsertAppointmentSession(appSess model.AppointmentSession) {

}
