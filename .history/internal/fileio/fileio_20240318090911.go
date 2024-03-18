package fileio

func saveToFile(data []byte, path string) error {
	f, err := os.Create(path)
  if err != nil
}

func fileOrDirectoryExists(path string) bool {}