package fileio

type SaveCompleteMessage struct {
	FullPath, SuccessMessage, Err string
}

func saveToFile(data []byte, path string) error {
	return nil
}

func createFile() error {
	return nil
}

func fileOrDirectoryExists(path string) bool {
	return false
}
