package fileio

import (
	"os"
	"path/filepath"
)

type SaveCompleteMessage struct {
	FullPath       string
	SuccessMessage string
	Err            error
}

func SaveToFile(savePathValue string, fileContent []string, fileType string) (*SaveCompleteMessage, error) {
	fullPath, err := CreateFile(savePathValue, "file", fileType, fileContent)

	if err != nil {
		return nil, err
	}

	path := filepath.Join(outputDirectory, fileName)

	return &SaveCompleteMessage{
		FullPath:       fullPath,
		SuccessMessage: "File saved successfully",
		Err:            nil,
	}, nil
}

func CreateFile(path, fileName string, fileType string, fileContent []string) (string, error) {
	fullPath := filepath.Join(path, fileName+"."+fileType)

	exists, err := fileOrDirectoryExists(fullPath)

	if err != nil {
		return "", err
	}

	if exists {
		// Return the file path if file already exists
		return fullPath, nil
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	for _, line := range fileContent {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			return "", err
		}
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
