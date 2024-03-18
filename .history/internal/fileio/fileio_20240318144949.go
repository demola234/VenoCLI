package fileio

import (
	"bytes"
	"fmt"
	"io"
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

func SaveToFile(savePathValue string, fileContent []string, fileType string) (*SaveCompleteMessage, error) {
	fullPath, err := CreateFile(savePathValue, "file", fileType, fileContent)

	if err != nil {
		return nil, err
	}

	path := filepath.Join(savePathValue, fullPath)

	info(path, true)

	src = bytes.NewBufferString(strings.Repeat("Some random input data", 1000))

	// Wrap it with our custom io.Reader.
	_ = &PassThru{Reader: src}

	count, err := io.Copy(&dst, src)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &SaveCompleteMessage{
		FullPath:       fullPath,
		SuccessMessage: "File saved successfully",
		Err:            nil,
	}, nil
}

func info(text string, verbose bool) {
	log.Printf("INFO: " + text)
	if verbose {
		fmt.Println("VinoCLI: " + text)
	}
}

func CreateFile(path, fileName string, fileType string, fileContent []string) (string, error) {
	fullPath := filepath.Join(path, fileName+"."+fileType)

	exists, err := fileOrDirectoryExists(fullPath)

	if err != nil {
		return "", err
	}

	if exists {
		return "", fmt.Errorf("file already exists at %s", fullPath)

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
