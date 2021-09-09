package mock

import (
	"fmt"

	staticfile "github.com/mjmhtjain/vaccine-alert-service/src/staticFile"
)

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
						},
						{
							"district_id": 143,
							"district_name": "test_district"
						}
					],
					"ttl": 24
				}`), nil

	default:
		return nil, fmt.Errorf("invalid input")
	}
}
