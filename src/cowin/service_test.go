package cowin

import (
	"testing"

	"github.com/mjmhtjain/vaccine-alert-service/src/mock"
)

// var mockCowinApiImpl *MockCowinAPIImpl

func TestCowinService(t *testing.T) {

	t.Run("When cowinRepo returns appointments Then expect appointments", func(t *testing.T) {
		appointmentService := &AppointmentServiceImpl{
			cowin:    mock.NewMockCowinAPI("../mock/appointmentSessionMock.json"),
			staticFS: mock.NewMockStaticFileService(),
			sqlRepo:  mock.NewMockSqlRepoImpl_SetResponse(false),
		}

		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments("Delhi", "2019-04-01")
		if err != nil {
			t.Errorf("Error in fetching appointments: %s", err)
		}

		if len(districtVaccineSlots) < 3 {
			t.Error("expected 3 appointment sessions")
		}
	})

	t.Run("When cowinRepo returns error then expect error", func(t *testing.T) {
		appointmentService := &AppointmentServiceImpl{
			cowin:    mock.NewMockCowinAPI(""),
			staticFS: mock.NewMockStaticFileService(),
			sqlRepo:  mock.NewMockSqlRepoImpl_SetResponse(true),
		}

		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments("Delhi", "2019-04-01")

		if districtVaccineSlots != nil && err == nil {
			t.Errorf("Error was expected")
		}
	})

	t.Run("When cowinRepo returns all stale appointments .. Then expect 0 appointments returned", func(t *testing.T) {
		appointmentService := &AppointmentServiceImpl{
			cowin:    mock.NewMockCowinAPI("../mock/appointmentSessionMock.json"),
			staticFS: mock.NewMockStaticFileService(),
			sqlRepo:  mock.NewMockSqlRepoImpl_SetResponse(true),
		}

		districtVaccineSlots, _ := appointmentService.FetchVaccineAppointments("Delhi", "2019-04-01")

		if len(districtVaccineSlots) > 0 {
			t.Errorf("All stale appointments are expected to be filtered")
		}
	})

	t.Run("When cowinRepo returns some stale appointments .. Then expect filtered appointments returned", func(t *testing.T) {
		// using repeated sessionId to mock stale sessions
		appointmentService := &AppointmentServiceImpl{
			cowin:    mock.NewMockCowinAPI("../mock/appointmentSessionMock.json"),
			staticFS: mock.NewMockStaticFileService(),
			sqlRepo:  mock.NewMockSqlRepoImpl_RecordSessions(),
		}
		expectedAppointments := 1

		districtVaccineSlots, _ := appointmentService.FetchVaccineAppointments("Delhi", "2019-04-01")

		if len(districtVaccineSlots) != expectedAppointments {
			t.Errorf("expected number of records %v, actual number of records %v",
				expectedAppointments,
				len(districtVaccineSlots))
		}
	})
}
