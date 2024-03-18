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
	var path, fileName string

	if savePathValue == "" {
		path = "."
		fileName = formatter.FormatTime(time.Now())
	} else {
		if strings.Contains(saveDialogValue, "~") {
			currUser, userErr := user.Current()
			if userErr != nil {
				return "", userErr
			}
			saveDialogValue = strings.ReplaceAll(saveDialogValue, "~", currUser.HomeDir)
		}

		if strings.Contains(saveDialogValue, string(os.PathSeparator)) {
			path = filepath.Dir(saveDialogValue)
			fileName = filepath.Base(saveDialogValue)
		} else {
			path = "."
			fileName = saveDialogValue
		}
	}

}

func createFile(path, fileName string, fileContent []string) (string, error) {
	fullPath := filepath.Join(path, fileName)

	exists, err := fileOrDirectoryExists(fullPath)

	if err != nil {
		return err
	}

	if exists {
		return "", fmt.Errorf("file already exists at %s",
			fullPath)
	}

	file, err := os.Create(fullPath)

	
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
