package fileio

import ("time")

type SaveCompleteMessage struct {
	FullPath, SuccessMessage, Err string
}

func saveToFile(savePathValue string, fileContent []string) (string, error) {
	var path, fileName string

	if savePathValue == "" {
		path = "."
		fileName = formatter.FormatTime(time.Now())
	}

}

func createFile() error {
	return nil
}

func fileOrDirectoryExists(path string) bool {
	return false
}
