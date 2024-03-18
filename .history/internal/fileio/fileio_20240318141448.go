package fileio

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

type SaveCompleteMessage struct {
	FullPath, SuccessMessage, Err string
}

func saveToFile(savePathValue string, fileContent []string) (string, error) {
	fullPath, err := createFile(savePathValue, "file.", fileContent)

}

func createFile(path, fileName string,  fileContent []string) (string, error) {
	fullPath := filepath.Join(path, fileName)

	exists, err := fileOrDirectoryExists(fullPath)

	if err != nil {
		return "", err
	}

	if exists {
		return "", fmt.Errorf("file already exists at %s",
			fullPath)
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
