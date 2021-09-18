package sql

import (
	"context"
	"errors"

	customerror "github.com/mjmhtjain/vaccine-alert-service/src/customError"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

var sqlRepoInstance SqlRepo

type SqlRepo interface {
	FindCenterWithCenterId(ctx context.Context, center model.Center) (*model.CenterORM, error)
	InsertCenterInfo(ctx context.Context, center model.Center) *model.CenterORM

	FindVaccineByName(ctx context.Context, name string) (*model.VaccineORM, error)
	InsertVaccine(ctx context.Context, vaccineName string) *model.VaccineORM

	FindSessionWithSessionId(ctx context.Context, sess *model.Session) (*model.AppointmentSessionORM, error)
	InsertAppointmentSession(ctx context.Context, appSess *model.Session, centerId int) *model.AppointmentSessionORM
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

func (impl *SqlRepoImpl) FindCenterWithCenterId(ctx context.Context, center model.Center) (*model.CenterORM, error) {
	centerOrm := model.CenterORM{}
	db := impl.dbConn.WithContext(ctx)

	result := db.Find(&centerOrm, "id = ?", center.CenterID)
	if result.RowsAffected <= 0 || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &customerror.NoRecordExists{Msg: "No record found"}
	}

	if result.Error != nil {
		panic(result.Error)
	}

	return &centerOrm, nil
}

func (impl *SqlRepoImpl) FindSessionWithSessionId(ctx context.Context, sess *model.Session) (*model.AppointmentSessionORM, error) {
	app := model.AppointmentSessionORM{}
	db := impl.dbConn.WithContext(ctx)

	result := db.Find(&app, "id = ?", sess.SessionID)
	if result.RowsAffected <= 0 || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &customerror.NoRecordExists{Msg: "No record found"}
	}

	if result.Error != nil {
		panic(result.Error)
	}
	return &app, nil
}

func (impl *SqlRepoImpl) FindVaccineByName(ctx context.Context, vaccineName string) (*model.VaccineORM, error) {
	vaccine := model.VaccineORM{}
	db := impl.dbConn.WithContext(ctx)

	result := db.Find(&vaccine, "name = ?", vaccineName)
	if result.RowsAffected <= 0 || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &customerror.NoRecordExists{Msg: "No record found"}
	}

	if result.Error != nil {
		panic(result.Error)
	}
	return &vaccine, nil
}

func (impl *SqlRepoImpl) InsertCenterInfo(ctx context.Context, center model.Center) *model.CenterORM {
	centerEntry := model.CenterORM{
		Id:           center.CenterID,
		Name:         center.Name,
		Address:      center.Address,
		StateName:    center.StateName,
		DistrictName: center.DistrictName,
		Pincode:      center.Pincode,
	}
	db := impl.dbConn.WithContext(ctx)

	result := db.Create(&centerEntry)
	if result.Error != nil {
		panic(result.Error)
	}

	return &centerEntry
}

func (impl *SqlRepoImpl) InsertVaccine(ctx context.Context, vaccineName string) *model.VaccineORM {
	vaccine := model.VaccineORM{
		Id:   generateUUID_8(),
		Name: vaccineName,
	}
	db := impl.dbConn.WithContext(ctx)

	result := db.Create(&vaccine)
	if result.Error != nil {
		panic(result.Error)
	}

	return &vaccine
}

func (impl *SqlRepoImpl) InsertAppointmentSession(
	ctx context.Context,
	appSess *model.Session,
	centerId int,
) *model.AppointmentSessionORM {
	db := impl.dbConn.WithContext(ctx)
	appointment := model.AppointmentSessionORM{
		Id:                     appSess.SessionID,
		CenterIDFK:             centerId,
		Date:                   appSess.Date,
		AvailableCapacity:      appSess.AvailableCapacity,
		MinAgeLimit:            appSess.MinAgeLimit,
		Vaccine:                appSess.Vaccine,
		AvailableCapacityDose1: appSess.AvailableCapacityDose1,
		AvailableCapacityDose2: appSess.AvailableCapacityDose2,
	}

	result := db.Create(&appointment)
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
