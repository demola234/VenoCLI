package download

type DownloadManager interface {
	Download(url string) (*DownloadResult, error)
	DisplayStatuses(int count) 
}




func DownloadPercentage(count string) (count string) {
	return count
}



func (d *DownloadResult) LogResult() {}

func downloadFile(url string) error {
	return nil
}

