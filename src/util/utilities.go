package util

import (
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"path/filepath"

	"github.com/mjmhtjain/vaccine-alert-service/src/logger"
)

func ReadStaticFile(fileName string) ([]byte, error) {
	logger.INFO.Printf("readStaticFile: fileName: %v \n", fileName)

	basePath, err := BasePath()
	if err != nil {
		logger.ERROR.Printf("Could not fetch basePath..\n %v \n", err)
		return nil, err
	}

	filename := filepath.Join(basePath, "src", "staticData", fileName)

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.ERROR.Printf("Error on reading file.. \n %v \n", err)
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
