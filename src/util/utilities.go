package util

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
)

var EmbededFiles embed.FS

func ReadStaticFile(fileName string) ([]byte, error) {
	logger.DEBUG.Printf("readStaticFile: fileName: %v \n", fileName)

	fsys, _ := fs.Sub(EmbededFiles, "staticData")
	fileData, err := fs.ReadFile(fsys, fileName)
	if err != nil {
		return nil, fmt.Errorf("file %s not found", fileName)
	}

	return fileData, nil
}

func Readfile(path string) ([]byte, error) {
	logger.DEBUG.Printf("Readfile: path: %v \n", path)

	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		logger.ERROR.Printf("Error in reading file.. \n %v \n", err)
		return nil, err
	}

	return fileData, nil
}

func BasePath() (string, error) {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		logger.ERROR.Printf("Could not execute command.. \n %v \n", err)
		return "", err
	}

	basePath := string(out)
	basePath = basePath[:len(basePath)-1]

	return basePath, nil
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// TODO: extract date as per region specific UTC time
func TodaysDate() string {
	t := time.Now()
	return fmt.Sprintf("%02d-%02d-%d",
		t.Day(), t.Month(), t.Year())
}
