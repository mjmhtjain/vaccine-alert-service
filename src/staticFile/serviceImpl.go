package staticfile

import (
	"embed"
	"io/fs"
	"log"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
)

type fileServiceImpl struct {
	embededFiles embed.FS
}

func (f *fileServiceImpl) Read(name string) ([]byte, error) {
	logger.DEBUG.Printf("Read: fileName: %v \n", name)

	fsys, _ := fs.Sub(f.embededFiles, "staticData")
	fileData, err := fs.ReadFile(fsys, name)
	if err != nil {
		log.Printf("file %s not found", name)
		return nil, err
	}

	return fileData, nil
}
