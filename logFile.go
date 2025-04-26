package logger

import (
	"os"
	"path/filepath"
)

func CreateNewLogFile(pathAndName ...string) error {
	var defaultPath string = "./Logger.log"
	var fullPath string

	if len(pathAndName) == 0 {
		fullPath = defaultPath
	} else {
		fullPath = pathAndName[0]
	}

	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
