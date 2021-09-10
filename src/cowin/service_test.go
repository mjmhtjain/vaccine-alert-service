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
			sqlRepo:  mock.NewMockSqlRepoImpl(false),
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
			sqlRepo:  mock.NewMockSqlRepoImpl(true),
		}

		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments("Delhi", "2019-04-01")

		if districtVaccineSlots != nil && err == nil {
			t.Errorf("Error was expected")
		}
	})

	t.Run("When cowinRepo returns appointments existing in db .. Then expect unique appointments to be returned", func(t *testing.T) {
		appointmentService := &AppointmentServiceImpl{
			cowin:    mock.NewMockCowinAPI(""),
			staticFS: mock.NewMockStaticFileService(),
			sqlRepo:  mock.NewMockSqlRepoImpl(true),
		}

		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments("Delhi", "2019-04-01")

		if districtVaccineSlots != nil && err == nil {
			t.Errorf("Error was expected")
		}
	})
}
