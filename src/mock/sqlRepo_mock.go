package mock

import (
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	"github.com/mjmhtjain/vaccine-alert-service/src/repo/sql"
)

// SetResponse Mock
func NewMockSqlRepoImpl_SetResponse(isSessionExist bool) sql.SqlRepo {
	return &MockSqlRepoImpl_SetResponse{
		isSessionExistResp: isSessionExist,
	}
}

type MockSqlRepoImpl_SetResponse struct {
	isSessionExistResp bool
}

func (impl *MockSqlRepoImpl_SetResponse) IsSessionExist(sess model.Session) bool {
	return impl.isSessionExistResp
}

func (impl *MockSqlRepoImpl_SetResponse) InsertCenterInfo(resp []model.Appointments, centerIndex int) (string, error) {
	return "", nil
}

func (impl *MockSqlRepoImpl_SetResponse) InsertVaccine(sess model.Session) (string, error) {
	return "", nil
}

func (impl *MockSqlRepoImpl_SetResponse) InsertAppointmentSession(appSess model.AppointmentSession) (string, error) {
	return "", nil
}

// RecordSessions Mock
func NewMockSqlRepoImpl_RecordSessions() sql.SqlRepo {
	return &MockSqlRepoImpl_RecordSessions{
		sessionMap: make(map[string]bool),
	}
}

type MockSqlRepoImpl_RecordSessions struct {
	sessionMap map[string]bool
}

func (impl *MockSqlRepoImpl_RecordSessions) IsSessionExist(sess model.Session) bool {
	if _, ok := impl.sessionMap[sess.SessionID]; ok {
		return true
	} else {
		impl.sessionMap[sess.SessionID] = true
		return false
	}
}

func (impl *MockSqlRepoImpl_RecordSessions) InsertCenterInfo(resp []model.Appointments, centerIndex int) (string, error) {
	return "", nil
}

func (impl *MockSqlRepoImpl_RecordSessions) InsertVaccine(sess model.Session) (string, error) {
	return "", nil
}

func (impl *MockSqlRepoImpl_RecordSessions) InsertAppointmentSession(appSess model.AppointmentSession) (string, error) {
	return "", nil
}
