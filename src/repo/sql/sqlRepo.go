package sql

import "github.com/mjmhtjain/vaccine-alert-service/src/model"

var sqlRepoInstance SqlRepo

type SqlRepo interface {
	IsSessionExist(sess model.Session) bool
	InsertCenterInfo(resp []model.Appointments, centerIndex int)
	InsertVaccine(sess model.Session) (string, error)
	InsertAppointmentSession(appSess model.AppointmentSession)
}

type SqlRepoImpl struct {
}

func NewSqlRepo() SqlRepo {
	if sqlRepoInstance == nil {
		sqlRepoInstance = &SqlRepoImpl{}
	}

	return sqlRepoInstance
}

func (impl *SqlRepoImpl) IsSessionExist(sess model.Session) bool {
	return false
}

func (impl *SqlRepoImpl) InsertCenterInfo(resp []model.Appointments, centerIndex int) {

}

func (impl *SqlRepoImpl) InsertVaccine(sess model.Session) (string, error) {
	return "", nil
}

func (impl *SqlRepoImpl) InsertAppointmentSession(appSess model.AppointmentSession) {

}