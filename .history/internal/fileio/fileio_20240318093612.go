package fileio

type SaveCompleteMessage struct {
	FullPath, SuccessMessage, Err string
}

func saveToFile(saveDialogValue string, fileContent []string) (string, error) {
	var path, fileName string

	if saveDialogValue == "" {
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
