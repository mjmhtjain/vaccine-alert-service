package mock

import (
	"encoding/json"
	"path/filepath"
	"sync"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
	"github.com/mjmhtjain/vaccine-alert-service/src/model"
	cowinrepo "github.com/mjmhtjain/vaccine-alert-service/src/repo/cowinRepo"
	"github.com/mjmhtjain/vaccine-alert-service/src/util"
)

func NewMockCowinAPI() cowinrepo.CowinAPI {
	mockCowinApiImpl := &MockCowinAPIImpl{
		callCount: 0,
		mutex:     &sync.Mutex{},
	}

	return mockCowinApiImpl
}

type MockCowinAPIImpl struct {
	callCount int
	mutex     *sync.Mutex
}

func (mock *MockCowinAPIImpl) AppointmentSessionByDistrictAndCalendar(districtId string, date string) (*model.Appointments, error) {
	mock.AddCallCount()

	var appointmentData *model.Appointments = new(model.Appointments)
	path, err := filepath.Abs("../mock/appointmentSessionMock.json")
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

func (spy *MockCowinAPIImpl) AddCallCount() {
	spy.mutex.Lock()
	spy.callCount++
	spy.mutex.Unlock()
}
