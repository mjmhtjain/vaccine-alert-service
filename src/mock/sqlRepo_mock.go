package mock

import (
	"context"

	customerror "github.com/mjmhtjain/vaccine-alert-service/src/customError"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	"github.com/mjmhtjain/vaccine-alert-service/src/repo/sql"
)

func NewMockSqlRepoImpl() sql.SqlRepo {
	return &MockSqlRepoImpl{
		centerMap:      make(map[int]bool),
		vaccineNameMap: make(map[string]bool),
		sessionMap:     make(map[string]bool),
	}
}

type MockSqlRepoImpl struct {
	centerMap      map[int]bool
	vaccineNameMap map[string]bool
	sessionMap     map[string]bool
}

func (impl *MockSqlRepoImpl) FindCenterWithCenterId(ctx context.Context, center model.Center) (*model.CenterORM, error) {
	if _, ok := impl.centerMap[center.CenterID]; ok {
		return &model.CenterORM{
			Id: center.CenterID,
		}, nil
	}

	return nil, &customerror.NoRecordExists{}
}

func (impl *MockSqlRepoImpl) InsertCenterInfo(ctx context.Context, center model.Center) *model.CenterORM {
	impl.centerMap[center.CenterID] = true

	return &model.CenterORM{
		Id: center.CenterID,
	}
}

func (impl *MockSqlRepoImpl) FindVaccineByName(ctx context.Context, name string) (*model.VaccineORM, error) {
	if _, ok := impl.vaccineNameMap[name]; ok {
		return &model.VaccineORM{
			Id:   "123",
			Name: name,
		}, nil
	}

	return nil, &customerror.NoRecordExists{}
}

func (impl *MockSqlRepoImpl) InsertVaccine(ctx context.Context, vaccineName string) *model.VaccineORM {
	impl.vaccineNameMap[vaccineName] = true

	return &model.VaccineORM{
		Id:   "123",
		Name: vaccineName,
	}
}

func (impl *MockSqlRepoImpl) FindSessionWithSessionId(ctx context.Context, sess *model.Session) (*model.AppointmentSessionORM, error) {
	if _, ok := impl.sessionMap[sess.SessionID]; ok {
		return &model.AppointmentSessionORM{
			Id: sess.SessionID,
		}, nil
	}

	return nil, &customerror.NoRecordExists{}
}

func (impl *MockSqlRepoImpl) InsertAppointmentSession(ctx context.Context, appSess *model.Session, centerId int, vaccineId string) *model.AppointmentSessionORM {
	impl.sessionMap[appSess.SessionID] = true

	return &model.AppointmentSessionORM{
		Id: appSess.SessionID,
	}
}
