package cowin

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	cowinrepo "github.com/mjmhtjain/vaccine-alert-service/src/repo/cowinRepo"
	staticfile "github.com/mjmhtjain/vaccine-alert-service/src/staticFile"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

func TestCowinService(t *testing.T) {

	t.Run("cowinRepo returns appointments", func(t *testing.T) {
		appointmentService := NewAppointmentServiceWithMockedCowinAPICall()
		districtVaccineSlots, err := appointmentService.FetchVaccineAppointments("Delhi", "2019-04-01")
		if err != nil {
			t.Errorf("Error in fetching appointments: %s", err)
		}

		if len(districtVaccineSlots) < 1 {
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
	switch name {
	case "states.json":
		return []byte(`{
			"states": [
				{
					"state_id": 9,
					"state_name": "Delhi"
				}
			],
			"ttl": 24
		}`), nil

	case "districts.json":
		return []byte(`{
					"districts": [
						{
							"district_id": 141,
							"district_name": "Central Delhi"
						},
						{
							"district_id": 142,
							"district_name": "test_district"
						}
					],
					"ttl": 24
				}`), nil

	default:
		return nil, fmt.Errorf("invalid input")
	}
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
