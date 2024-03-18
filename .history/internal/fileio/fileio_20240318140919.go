package fileio

import (
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

func createFile(path, fileName string, fileContent []string) error {
	fullPath := path + "/" + fileName

	err := ioutil
}

func fileOrDirectoryExists(path string) bool {
	return false
}
