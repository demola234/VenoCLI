package fileio

import (
	"fmt"
	"os"
	"path/filepath"
)

type SaveCompleteMessage struct {
	FullPath, SuccessMessage, Err string
}

func saveToFile(savePathValue string, fileContent []string, fileType string) (string, error) {
	fullPath, err := createFile(savePathValue, "file", fileType, fileContent)

	if err != nil {
		return "", err
	}



	

}

func createFile(path, fileName string, fileType string, fileContent []string) (string, error) {
	fullPath := filepath.Join(path, fileName+"."+fileType)

	exists, err := fileOrDirectoryExists(fullPath)

	if err != nil {
		return "", err
	}

	if exists {
		// R
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	for _, line := range fileContent {
		_, err = file.WriteString(line + "\n")
	}

	if err != nil {
		return "", err
	}

	// Return the file path if created successfully
	return fullPath, nil

}

func fileOrDirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
