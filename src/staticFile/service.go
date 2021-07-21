package staticfile

import "github.com/mjmhtjain/vaccine-alert-service/src/util"

var fileServiceInstance FileService

type FileService interface {
	Read(name string) ([]byte, error)
}

func NewFileService() FileService {
	if fileServiceInstance == nil {
		fileServiceInstance = &fileServiceImpl{
			embededFiles: util.EmbededFiles,
		}
	}

	return fileServiceInstance
}
