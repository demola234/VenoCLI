package fileio

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveToFile(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "fileio_test")
	if err != nil {
		t.Fatal("Failed to create temporary directory:", err)
	}
	defer os.RemoveAll(tempDir)

	testFilePath := filepath.Join(tempDir, "test.txt")
	fileContent := []string{"line 1", "line 2", "line 3"}
	fileType := ""
	verbose := true

	saveMessage, err := SaveToFile(testFilePath, fileContent, fileType, verbose)
	if err != nil {
		t.Errorf("SaveToFile returned an unexpected error: %v", err)
	}

	if saveMessage == nil {
		t.Error("SaveToFile did not return a save message")
	}

	if saveMessage.Err != nil {
		t.Errorf("SaveToFile returned an unexpected error in save message: %v", saveMessage.Err)
	}

	if saveMessage.FullPath != testFilePath {
		t.Errorf("Expected full path %s, got %s", testFilePath, saveMessage.FullPath)
	}
}
package fileio_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/your/package/path/fileio"
)

func TestCreateFile(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "fileio_test")
	if err != nil {
		t.Fatal("Failed to create temporary directory:", err)
	}
	defer os.RemoveAll(tempDir)

	testFileName := "test.txt"
	fileContent := []string{"line 1", "line 2", "line 3"}
	fileType := "txt"

	createdFilePath, err := fileio.CreateFile(tempDir, testFileName, fileType, fileContent)
	if err != nil {
		t.Errorf("CreateFile returned an unexpected error: %v", err)
	}

	expectedFilePath := filepath.Join(tempDir, testFileName)
	if createdFilePath != expectedFilePath {
		t.Errorf("Expected created file path %s, got %s", expectedFilePath, createdFilePath)
	}

	// Check if the file was created
	if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
		t.Errorf("Expected file %s to be created, but it doesn't exist", expectedFilePath)
	}

	// Read the file and check its content
	actualContent, err := ioutil.ReadFile(expectedFilePath)
	if err != nil {
		t.Errorf("Failed to read file: %v", err)
	}

	expectedContent := "line 1\nline 2\nline 3\n"
	if string(actualContent) != expectedContent {
		t.Errorf("File content mismatch. Expected:\n%s\nGot:\n%s", expectedContent, string(actualContent))
	}
}
