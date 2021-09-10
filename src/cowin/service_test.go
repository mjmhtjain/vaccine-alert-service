package cowin

import (
	"testing"

	"github.com/mjmhtjain/vaccine-alert-service/src/mock"
)

// var mockCowinApiImpl *MockCowinAPIImpl

func TestCowinService(t *testing.T) {

	t.Run("cowinRepo returns appointments", func(t *testing.T) {
		appointmentService := &AppointmentServiceImpl{
			cowin:    mock.NewMockCowinAPI("../mock/appointmentSessionMock.json"),
			staticFS: mock.NewMockStaticFileService(),
			sqlRepo:  mock.NewMockSqlRepoImpl(true),
		}

		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments("Delhi", "2019-04-01")
		if err != nil {
			t.Errorf("Error in fetching appointments: %s", err)
		}

		if len(districtVaccineSlots) < 1 {
			t.Error("expected appointments to be populated")
		}

		for _, vs := range districtVaccineSlots {
			if len(vs.Centers) <= 0 {
				t.Error("expected atleast one appointment")
			}
		}
	})
}
