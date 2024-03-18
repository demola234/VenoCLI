package download

type DownloadManager interface {
	Download(url string) (*DownloadResult, error)
	DisplayStatuses(count int) (string, error)
}

type DownloadResult struct {
	Url      string `json:"url"`
	Size     int64
	Received int64
	Error    error
}

func DownloadPerDisplayStatusescentage(count int) (count string) {
	return count
}

func (d *DownloadResult) LogResult() {}

func downloadFile(url string) error {
	return nil
}