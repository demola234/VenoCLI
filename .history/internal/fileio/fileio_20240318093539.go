package fileio

type SaveCompleteMessage struct {
	FullPath, SuccessMessage, Err string
}

func saveToFile(saveDialogValue string, fileContent []string) (string, error) {
	return nil
}

func createFile() error {
	return nil
}

func fileOrDirectoryExists(path string) bool {
	return false
}
