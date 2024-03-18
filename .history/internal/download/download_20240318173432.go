package download

import (
	"demola/vino/internal/https"
	"io"
	"net/http"
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

type writeCounter struct {
	BytesDownloaded int64
	TotalBytes      int64
}

type passThru struct {
	io.Reader
	total int64
}

type downloadManager struct {
	percent int
	verbose bool
}

// NewDownloadManager creates a new instance of DownloadManager.
func NewDownloadManager() DownloadManager {
	return &downloadManager{}
}

func (dm *downloadManager) Download(url string) (*DownloadResult, error) {
	result := &DownloadResult{Url: url}
	if err := downloadFile(url); err != nil {
		result.Error = err
		return result, err
	}

	result.Size, result.Received = getFileSizeAndReceived(url, result.Error)

	return result, nil
}

func (dm *downloadManager) DisplayStatuses(count int) (string, error) {
	// Implementation for displaying download statuses
	return "", nil
}

func downloadFile(url string, reqHandler https.RequestHandler, unmarshaler https.Unmarshaler) error {
	httpsClient := https.NewHTTPS(reqHandler, unmarshaler)

	// Placeholder implementation for downloading file

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

		return err
	}

	defer req.Body.Close()

	resp, err := httpsClient.ReqHandler.HandleRequest(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	//* Check status code 
	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status code: " + resp.Status)
		
	// Implement file download logic here

	return nil
}

func getFileSizeAndReceived(url string, err error) (int64, int64) {
	if err != nil {
		// Handle error case
		return 0, 0
	}
	// Placeholder implementation for obtaining file size and received size
	return 0, 0
}

func (d *DownloadResult) LogResult() {
	// Placeholder implementation for logging download result
}
