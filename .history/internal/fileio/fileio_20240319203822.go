package fileio

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type SaveCompleteMessage struct {
	FullPath       string
	SuccessMessage string
	Err            error
}

func SaveToFile(savePathValue string, fileContent []string, fileType string, verbose bool) (*SaveCompleteMessage, error) {
	fullPath, err := CreateFile(filepath.Dir(savePathValue), filepath.Base(savePathValue), fileType, fileContent)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %v", err)
	}

	if verbose {
		info("File saved successfully", true)
	}

	return &SaveCompleteMessage{
		FullPath:       fullPath,
		SuccessMessage: "File saved successfully",
		Err:            nil,
	}, nil
}

// info logs informational messages.
func info(text string, verbose bool) {
	log.Printf("INFO: " + text)
	if verbose {
		fmt.Println("VinoCLI: " + text)
	}
}

// CreateFile creates a file with the specified fileContent at the specified path.
func CreateFile(path, fileName string, fileType string, fileContent []string) (string, error) {
	if strings.Split(fileName, ".")[1 {
		fileName = fileName + "." + fileType
	}
	fullPath := filepath.Join(path, fileName)

	exists, err := fileOrDirectoryExists(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to check file existence: %v", err)
	}

	if exists {
		return "", fmt.Errorf("file already exists at %s", fullPath)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// File Content Writer
	for _, line := range fileContent {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			return "", fmt.Errorf("failed to write to file: %v", err)
		}
	}

	// Return the file path if created successfully
	return fullPath, nil
}

// fileOrDirectoryExists checks if a file or directory exists at the specified path.
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
