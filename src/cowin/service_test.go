package cowin

import (
	"testing"

	"github.com/mjmhtjain/vaccine-alert-service/src/mock"
)

// var mockCowinApiImpl *MockCowinAPIImpl

func TestCowinService(t *testing.T) {

	t.Run("cowinRepo returns appointments", func(t *testing.T) {
		appointmentService := NewMockAppointmentService()
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

		// if mockCowinApiImpl.callCount != 3 {
		// 	t.Errorf("expected call count: %v, actual call count: %v", 2, mockCowinApiImpl.callCount)
		// }
	})
}

func NewMockAppointmentService() AppointmentService {
	return &AppointmentServiceImpl{
		cowin:    mock.NewMockCowinAPI(),
		staticFS: mock.NewMockStaticFileService(),
		sqlRepo:  mock.NewMockSqlRepoImpl(true),
	}
}

// func NewMockStaticFileService() staticfile.FileService {
// 	return &MockStaticFileServiceImpl{}
// }

// type MockStaticFileServiceImpl struct {
// }

// func (mock *MockStaticFileServiceImpl) Read(name string) ([]byte, error) {
// 	switch name {
// 	case "states.json":
// 		return []byte(`{
// 			"states": [
// 				{
// 					"state_id": 9,
// 					"state_name": "Delhi"
// 				}
// 			],
// 			"ttl": 24
// 		}`), nil

// 	case "districts.json":
// 		return []byte(`{
// 					"districts": [
// 						{
// 							"district_id": 141,
// 							"district_name": "Central Delhi"
// 						},
// 						{
// 							"district_id": 142,
// 							"district_name": "test_district"
// 						},
// 						{
// 							"district_id": 143,
// 							"district_name": "test_district"
// 						}
// 					],
// 					"ttl": 24
// 				}`), nil

// 	default:
// 		return nil, fmt.Errorf("invalid input")
// 	}
// }

// func NewMockCowinAPI() cowinrepo.CowinAPI {
// 	mockCowinApiImpl = &MockCowinAPIImpl{
// 		callCount: 0,
// 		mutex:     &sync.Mutex{},
// 	}

// 	return mockCowinApiImpl
// }

// type MockCowinAPIImpl struct {
// 	callCount int
// 	mutex     *sync.Mutex
// }

// func (mock *MockCowinAPIImpl) AppointmentSessionByDistrictAndCalendar(districtId string, date string) (*model.Appointments, error) {
// 	mock.AddCallCount()

// 	var appointmentData *model.Appointments = new(model.Appointments)
// 	path, err := filepath.Abs("../mock/appointmentSessionMock.json")
// 	if err != nil {
// 		logger.ERROR.Printf("Invalid filepath.. \n %v \n", err)
// 		return nil, err
// 	}

// 	data, err := util.Readfile(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = json.Unmarshal(data, appointmentData)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return appointmentData, nil
// }

// func (spy *MockCowinAPIImpl) AddCallCount() {
// 	spy.mutex.Lock()
// 	spy.callCount++
// 	spy.mutex.Unlock()
// }
