package cowin

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	cowinrepo "github.com/mjmhtjain/vaccine-alert-service/src/repo/cowinRepo"
	staticfile "github.com/mjmhtjain/vaccine-alert-service/src/staticFile"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

func TestCowinService(t *testing.T) {
	appointmentService := NewAppointmentServiceWithMockedCowinAPICall()

	t.Run("given valid arguments, cowin should return appointments", func(t *testing.T) {
		appointments, err := appointmentService.FetchVaccineAppointments("9", "2019-04-01")
		if err != nil {
			t.Errorf("Error in fetching appointments: %s", err)
		}

		if len(appointments.Centers) < 1 {
			t.Error("expected appointments to be populated")
		}
	})
}

func NewAppointmentServiceWithMockedCowinAPICall() AppointmentService {
	return &AppointmentServiceImpl{
		cowin:    NewMockCowinAPI(),
		staticFS: NewMockStaticFileService(),
	}
}

func NewMockStaticFileService() staticfile.FileService {
	return &MockStaticFileServiceImpl{}
}

type MockStaticFileServiceImpl struct {
}

func (mock *MockStaticFileServiceImpl) Read(name string) ([]byte, error) {
	return []byte(`{
		"districts": [
			{
				"district_id": 141,
				"district_name": "Central Delhi"
			}
		],
		"ttl": 24
	}`), nil
}

func NewMockCowinAPI() cowinrepo.CowinAPI {
	return &MockCowinAPIImpl{}
}

type MockCowinAPIImpl struct {
}

func (mock *MockCowinAPIImpl) AppointmentSessionByDistrictAndCalendar(districtId string, date string) (*model.Appointments, error) {
	var appointmentData *model.Appointments = new(model.Appointments)
	path, err := filepath.Abs("./mock/appointmentSessionMock.json")
	if err != nil {
		logger.ERROR.Printf("Invalid filepath.. \n %v \n", err)
		return nil, err
	}

	data, err := util.Readfile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, appointmentData)
	if err != nil {
		return nil, err
	}

	return appointmentData, nil
}