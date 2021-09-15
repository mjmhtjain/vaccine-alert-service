package sql

import (
	"errors"

	customerror "github.com/mjmhtjain/vaccine-alert-service/src/customError"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

var sqlRepoInstance SqlRepo

const (
	findSessionBySessionId string = `select session_id from appointment_session where session_id = ?`
	findVaccineByName      string = `select name from vaccine where name = ?`
	findCenterById                = `select center_id from center_info where center_id = ?`
	insertCenterInfo              = `insert into center_info 
									(center_id, name, address,state_name,district_name,pincode) 
									values (?, ?, ?,?, ?, ?)`
	insertVaccineInfo        string = `insert into vaccine (id, name) values (?, ?)`
	insertAppointmentSession string = `insert into appointment_session 
								(session_id, center_idfk, date, available_capacity, min_age_limit, vaccine_idfk, available_capacity_dose1,available_capacity_dose2) 
								values (?, ?,?, ?,?, ?,?, ?)`
)

type SqlRepo interface {
	FindCenterWithCenterId(center model.Center) (*model.CenterORM, error)
	FindSessionWithSessionId(sess *model.Session) (*model.AppointmentSessionORM, error)
	FindVaccineByName(name string) (*model.VaccineORM, error)

	InsertCenterInfo(center model.Center) *model.CenterORM
	InsertVaccine(vaccineName string) *model.VaccineORM
	InsertAppointmentSession(appSess *model.Session, centerId int, vaccineId string) *model.AppointmentSessionORM
}

type SqlRepoImpl struct {
	dbConn *gorm.DB
}

func NewSqlRepo() SqlRepo {
	if sqlRepoInstance == nil {
		sqlRepoInstance = &SqlRepoImpl{
			dbConn: GetConnection(),
		}
	}

	return sqlRepoInstance
}

// TODO:: handle context related complexities
func (impl *SqlRepoImpl) FindCenterWithCenterId(center model.Center) (*model.CenterORM, error) {
	centerOrm := model.CenterORM{}

	result := impl.dbConn.Find(&centerOrm, "id = ?", center.CenterID)
	if result.RowsAffected <= 0 || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &customerror.NoRecordExists{Msg: "No record found"}
	}

	if result.Error != nil {
		panic(result.Error)
	}

	return &centerOrm, nil
}

func (impl *SqlRepoImpl) FindSessionWithSessionId(sess *model.Session) (*model.AppointmentSessionORM, error) {
	app := model.AppointmentSessionORM{}

	result := impl.dbConn.Find(&app, "id = ?", sess.SessionID)
	if result.RowsAffected <= 0 || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &customerror.NoRecordExists{Msg: "No record found"}
	}

	if result.Error != nil {
		panic(result.Error)
	}
	return &app, nil
}

func (impl *SqlRepoImpl) FindVaccineByName(vaccineName string) (*model.VaccineORM, error) {
	vaccine := model.VaccineORM{}

	result := impl.dbConn.Find(&vaccine, "name = ?", vaccineName)
	if result.RowsAffected <= 0 || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &customerror.NoRecordExists{Msg: "No record found"}
	}

	if result.Error != nil {
		panic(result.Error)
	}
	return &vaccine, nil
}

func (impl *SqlRepoImpl) InsertCenterInfo(center model.Center) *model.CenterORM {
	centerEntry := model.CenterORM{
		Id:           center.CenterID,
		Name:         center.Name,
		Address:      center.Address,
		StateName:    center.StateName,
		DistrictName: center.DistrictName,
		Pincode:      center.Pincode,
	}

	result := impl.dbConn.Create(&centerEntry)
	if result.Error != nil {
		panic(result.Error)
	}

	return &centerEntry
}

func (impl *SqlRepoImpl) InsertVaccine(vaccineName string) *model.VaccineORM {
	vaccine := model.VaccineORM{
		Id:   generateUUID_8(),
		Name: vaccineName,
	}
	result := impl.dbConn.Create(&vaccine)
	if result.Error != nil {
		panic(result.Error)
	}

	return &vaccine
}

func (impl *SqlRepoImpl) InsertAppointmentSession(
	appSess *model.Session,
	centerId int,
	vaccineId string,
) *model.AppointmentSessionORM {

	appointment := model.AppointmentSessionORM{
		Id:                     appSess.SessionID,
		CenterIDFK:             centerId,
		Date:                   appSess.Date,
		AvailableCapacity:      appSess.AvailableCapacity,
		MinAgeLimit:            appSess.MinAgeLimit,
		VaccineIDKF:            vaccineId,
		AvailableCapacityDose1: appSess.AvailableCapacityDose1,
		AvailableCapacityDose2: appSess.AvailableCapacityDose2,
	}

	result := impl.dbConn.Create(&appointment)
	if result.Error != nil {
		panic(result.Error)
	}

	return &appointment
}

func generateUUID_8() string {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	uuid := u.String()[:8]
	return uuid
}
