package staticfile

import (
	"embed"
	"fmt"
	"io/fs"

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
		return nil, fmt.Errorf("file %s not found", name)
	}

	return fileData, nil
}
