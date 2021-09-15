package mock

import (
	customerror "github.com/mjmhtjain/vaccine-alert-service/src/customError"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	"github.com/mjmhtjain/vaccine-alert-service/src/repo/sql"
)

// SetResponse Mock
// func NewMockSqlRepoImpl_SetResponse(isSessionExist bool) sql.SqlRepo {
// 	return &MockSqlRepoImpl_SetResponse{
// 		isSessionExistResp: isSessionExist,
// 	}
// }

// type MockSqlRepoImpl_SetResponse struct {
// 	isSessionExistResp bool
// }

// func (impl *MockSqlRepoImpl) FindCenterWithCenterId(center model.Center) (*model.CenterORM, error) {
// 	return nil, nil
// }

// func (impl *MockSqlRepoImpl) FindSessionWithSessionId(sess model.Session) (*model.AppointmentSessionORM, error) {
// 	return nil, nil
// }

// func (impl *MockSqlRepoImpl) FindVaccineByName(name string) (*model.VaccineORM, error) {
// 	return nil, nil
// }

// func (impl *MockSqlRepoImpl) InsertCenterInfo(center model.Center) (*model.CenterORM, error) {
// 	return nil, nil
// }

// func (impl *MockSqlRepoImpl) InsertVaccine(sess model.Session) (*model.VaccineORM, error) {
// 	return nil, nil
// }

// func (impl *MockSqlRepoImpl) InsertAppointmentSession(appSess *model.Session, centerId int, vaccineId string) (*model.AppointmentSessionORM, error) {
// 	return nil, nil
// }

// RecordSessions Mock
func NewMockSqlRepoImpl(
	FindCenterWithCenterIdResponse *model.CenterORM,
	FindSessionWithSessionIdResponse *model.AppointmentSessionORM,
	FindVaccineByNameResponse *model.VaccineORM,
	InsertCenterInfoResponse *model.CenterORM,
	InsertVaccineResponse *model.VaccineORM,
	InsertAppointmentSessionResponse *model.AppointmentSessionORM,
) sql.SqlRepo {
	return &MockSqlRepoImpl{
		// FindCenterWithCenterIdResponse:   FindCenterWithCenterIdResponse,
		// FindSessionWithSessionIdResponse: FindSessionWithSessionIdResponse,
		// FindVaccineByNameResponse:        FindVaccineByNameResponse,
		// InsertCenterInfoResponse:         InsertCenterInfoResponse,
		// InsertVaccineResponse:            InsertVaccineResponse,
		// InsertAppointmentSessionResponse: InsertAppointmentSessionResponse,
		centerMap:      make(map[int]bool),
		vaccineNameMap: make(map[string]bool),
		sessionMap:     make(map[string]bool),
	}
}

type MockSqlRepoImpl struct {
	FindCenterWithCenterIdResponse   *model.CenterORM
	FindSessionWithSessionIdResponse *model.AppointmentSessionORM
	FindVaccineByNameResponse        *model.VaccineORM
	InsertCenterInfoResponse         *model.CenterORM
	InsertVaccineResponse            *model.VaccineORM
	InsertAppointmentSessionResponse *model.AppointmentSessionORM
	centerMap                        map[int]bool
	vaccineNameMap                   map[string]bool
	sessionMap                       map[string]bool
}

func (impl *MockSqlRepoImpl) FindCenterWithCenterId(center model.Center) (*model.CenterORM, error) {
	if _, ok := impl.centerMap[center.CenterID]; ok {
		return &model.CenterORM{
			Id: center.CenterID,
		}, nil
	}

	return nil, &customerror.NoRecordExists{}
}

func (impl *MockSqlRepoImpl) InsertCenterInfo(center model.Center) *model.CenterORM {
	impl.centerMap[center.CenterID] = true

	return &model.CenterORM{
		Id: center.CenterID,
	}
}

func (impl *MockSqlRepoImpl) FindVaccineByName(name string) (*model.VaccineORM, error) {
	if _, ok := impl.vaccineNameMap[name]; ok {
		return &model.VaccineORM{
			Id:   "123",
			Name: name,
		}, nil
	}

	return nil, &customerror.NoRecordExists{}
}

func (impl *MockSqlRepoImpl) InsertVaccine(vaccineName string) *model.VaccineORM {
	impl.vaccineNameMap[vaccineName] = true

	return &model.VaccineORM{
		Id:   "123",
		Name: vaccineName,
	}
}

func (impl *MockSqlRepoImpl) FindSessionWithSessionId(sess *model.Session) (*model.AppointmentSessionORM, error) {
	if _, ok := impl.sessionMap[sess.SessionID]; ok {
		return &model.AppointmentSessionORM{
			Id: sess.SessionID,
		}, nil
	}

	return nil, &customerror.NoRecordExists{}
}

func (impl *MockSqlRepoImpl) InsertAppointmentSession(appSess *model.Session, centerId int, vaccineId string) *model.AppointmentSessionORM {
	impl.sessionMap[appSess.SessionID] = true

	return &model.AppointmentSessionORM{
		Id: appSess.SessionID,
	}
}
