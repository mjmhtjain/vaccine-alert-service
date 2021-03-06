package cowin

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/mjmhtjain/vaccine-alert-service/src/mock"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

func TestAppointmentService_SqlRepo(t *testing.T) {
	appointments := model.Appointments{}
	ctx := context.Background()
	ReadJsonFile("../mock/session_1.json", &appointments)

	mockCowin := mock.NewMockCowinAPI(&appointments)
	mockStaticFS := mock.NewMockStaticFileService()

	t.Run("When sqlRepo has no appointments data stored.. Then expect all sessions returned", func(t *testing.T) {
		mockSqlRepo := mock.NewMockSqlRepoImpl()

		appointmentService := &AppointmentServiceImpl{
			cowin:    mockCowin,
			staticFS: mockStaticFS,
			sqlRepo:  mockSqlRepo,
		}

		expectedAppointments := 1
		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments(ctx, "Delhi", "2019-04-01")
		if err != nil {
			t.Errorf("unexpected error in fetching appointments: %s", err)
		}

		if len(districtVaccineSlots) != expectedAppointments {
			t.Errorf(
				"expected number of records %v, actual number of records %v",
				expectedAppointments,
				len(districtVaccineSlots),
			)
		}
	})

	t.Run("When sqlRepo contains data for centers and vaccine but no session data.. Then expect all sessions returned", func(t *testing.T) {
		mockSqlRepo := mock.NewMockSqlRepoImpl()

		appointmentService := &AppointmentServiceImpl{
			cowin:    mockCowin,
			staticFS: mockStaticFS,
			sqlRepo:  mockSqlRepo,
		}

		// inserting center and vaccine data
		s := appointments.Centers[0].Sessions[0]
		s.SessionID = "123"
		mockSqlRepo.InsertAppointmentSession(ctx, &s, 123)

		// asserting
		expectedAppointments := 1
		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments(ctx, "Delhi", "2019-04-01")
		if err != nil {
			t.Errorf("unexpected error in fetching appointments: %s", err)
		}

		if len(districtVaccineSlots) != expectedAppointments {
			t.Errorf(
				"expected number of records %v, actual number of records %v",
				expectedAppointments,
				len(districtVaccineSlots),
			)
		}
	})

	t.Run("When all stale sessions data is given .. Then expect 0 sessions returned", func(t *testing.T) {
		mockSqlRepo := mock.NewMockSqlRepoImpl()

		appointmentService := &AppointmentServiceImpl{
			cowin:    mockCowin,
			staticFS: mockStaticFS,
			sqlRepo:  mockSqlRepo,
		}

		// inserting the appointments in advance
		for _, c := range appointments.Centers {
			for _, s := range c.Sessions {
				mockSqlRepo.InsertAppointmentSession(ctx, &s, 123)
			}
		}

		expectedAppointments := 0
		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments(ctx, "Delhi", "2019-04-01")
		if err != nil {
			t.Errorf("unexpected error in fetching appointments: %s", err)
		}

		if len(districtVaccineSlots) != expectedAppointments {
			t.Errorf(
				"expected number of records %v, actual number of records %v",
				expectedAppointments,
				len(districtVaccineSlots),
			)
		}
	})

	t.Run("When mix of fresh and stale session data is returned .. Then expect fresh sessions data returned", func(t *testing.T) {
		// adding fresh sessions mockCowin
		for _, id := range []string{"session_1", "session_2"} {
			tempSession := appointments.Centers[0].Sessions[0]
			tempSession.SessionID = id

			appointments.Centers[0].Sessions = append(appointments.Centers[0].Sessions, tempSession)
		}

		mockCowin := mock.NewMockCowinAPI(&appointments)
		mockSqlRepo := mock.NewMockSqlRepoImpl()

		appointmentService := &AppointmentServiceImpl{
			cowin:    mockCowin,
			staticFS: mockStaticFS,
			sqlRepo:  mockSqlRepo,
		}

		// inserting stale appointments in sqlRepo
		tempSession := appointments.Centers[0].Sessions[0]
		mockSqlRepo.InsertAppointmentSession(ctx, &tempSession, 123)

		// asserting
		expectedAppointments := 2
		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments(ctx, "Delhi", "2019-04-01")
		if err != nil {
			t.Errorf("unexpected error in fetching appointments: %s", err)
		}

		if len(districtVaccineSlots) != expectedAppointments {
			t.Errorf(
				"expected number of records %v, actual number of records %v",
				expectedAppointments,
				len(districtVaccineSlots),
			)
		}
	})
}

func TestAppointmentService_CowinService(t *testing.T) {
	ctx := context.Background()
	mockStaticFS := mock.NewMockStaticFileService()
	mockSqlRepo := mock.NewMockSqlRepoImpl()

	t.Run("When CowinService responds with fresh sessions.. Then expect sessions in result", func(t *testing.T) {
		appointments := model.Appointments{}
		ReadJsonFile("../mock/session_1.json", &appointments)
		mockCowin := mock.NewMockCowinAPI(&appointments)

		appointmentService := &AppointmentServiceImpl{
			cowin:    mockCowin,
			staticFS: mockStaticFS,
			sqlRepo:  mockSqlRepo,
		}

		// assert
		expectedAppointments := 1
		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments(ctx, "Delhi", "2019-04-01")
		if err != nil {
			t.Errorf("unexpected error in fetching appointments: %s", err)
		}

		if len(districtVaccineSlots) != expectedAppointments {
			t.Errorf(
				"expected number of records %v, actual number of records %v",
				expectedAppointments,
				len(districtVaccineSlots),
			)
		}
	})

	t.Run("When CowinService gives no sessions.. Then expect empty response", func(t *testing.T) {
		appointments := model.Appointments{}
		ReadJsonFile("../mock/session_1.json", &appointments)
		appointments.Centers = appointments.Centers[:0]

		mockCowin := mock.NewMockCowinAPI(&appointments)

		appointmentService := &AppointmentServiceImpl{
			cowin:    mockCowin,
			staticFS: mockStaticFS,
			sqlRepo:  mockSqlRepo,
		}

		// assert
		expectedAppointments := 0
		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments(ctx, "Delhi", "2019-04-01")
		if err != nil {
			t.Errorf("unexpected error in fetching appointments: %s", err)
		}

		if len(districtVaccineSlots) != expectedAppointments {
			t.Errorf(
				"expected number of records %v, actual number of records %v",
				expectedAppointments,
				len(districtVaccineSlots),
			)
		}
	})

	t.Run("When CowinService throws error.. Then expect panic", func(t *testing.T) {
		mockCowin := mock.NewMockCowinAPI(nil)
		defer PanicCheck(t)

		appointmentService := &AppointmentServiceImpl{
			cowin:    mockCowin,
			staticFS: mockStaticFS,
			sqlRepo:  mockSqlRepo,
		}

		// call appointment service
		appointmentService.FetchVaccineAppointments(ctx, "Delhi", "2019-04-01")
	})
}

func TestAppointmentService_ContextCancel(t *testing.T) {
	ctx := context.Background()
	mockStaticFS := mock.NewMockStaticFileService()
	mockSqlRepo := mock.NewMockSqlRepoImpl()

	t.Run("When Root context gets cancelled.. Then expect panic response", func(t *testing.T) {
		appointments := model.Appointments{}
		ctx, cancel := context.WithCancel(ctx)
		ReadJsonFile("../mock/session_1.json", &appointments)
		defer PanicCheck(t)

		mockCowin := mock.NewMockCowinAPI(&appointments)

		appointmentService := &AppointmentServiceImpl{
			cowin:    mockCowin,
			staticFS: mockStaticFS,
			sqlRepo:  mockSqlRepo,
		}

		// cancel context almost immediately
		go func() {
			cancel()
		}()
		appointmentService.FetchVaccineAppointments(ctx, "Delhi", "2019-04-01")
	})

}

func ReadJsonFile(relativeFilePath string, model interface{}) {
	path, err := filepath.Abs(relativeFilePath)
	util.ErrorPanic(err)

	data, err := util.Readfile(path)
	util.ErrorPanic(err)

	err = json.Unmarshal(data, model)
	util.ErrorPanic(err)
}

func PanicCheck(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("The code did not panic")
	}
}
