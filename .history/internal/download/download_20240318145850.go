package download

import (
	"bytes"
	"demola/vino/internal/fileio"
	"io"
)

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

func (dm *DownloadManager) Download(url string) (*DownloadResult, error) {
	result := &DownloadResult{Url: url}
	if err := downloadFile(url); err != nil {
		result.Error = err
		return result, err
	}

	result.Size, result.Received = getFileSizeAndReceived(url, result.Error)

	return result, nil
}


func getFileSizeAndReceived(url string, err error) size
	// return size, received

}

func (d *DownloadResult) LogResult() {}

func downloadFile(url string) error {
	var src io.Reader
	var dst bytes.Buffer

	fileio.SaveToFile()



}