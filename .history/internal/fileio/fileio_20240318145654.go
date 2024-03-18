package fileio

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type SaveCompleteMessage struct {
	FullPath       string
	SuccessMessage string
	Err            error
}

// SaveToFile saves fileContent to a file at savePathValue with the specified fileType.
func SaveToFile(savePathValue string, fileContent []string, fileType string, verbose bool) (*SaveCompleteMessage, error) {
	fullPath, err := CreateFile(savePathValue, "file", fileType, fileContent)
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
	fullPath := filepath.Join(path, fileName+"."+fileType)

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


	// 
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
