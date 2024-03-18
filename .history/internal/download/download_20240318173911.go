package download

import (
	"bytes"
	"demola/vino/internal/https"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func (dm *downloadManager) downloadFile(url string, reqHandler https.RequestHandler, unmarshaler https.Unmarshaler) error {
	var src io.Reader // Source file/url/etc
	var dst bytes.Buffer
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

	// Create some random input data.
	src = bytes.NewBufferString(strings.Repeat("Some random input data", 1000))

	// Wrap it with our custom io.Reader.
	_ = &PassThru{Reader: src}

	count, err := io.Copy(&dst, src)

	//* Check status code
	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status code: " + resp.Status)
	}

	//*create body
	var body io.Reader

	if !dm.verbose {
		body = resp.Body
	} else {
		go displayStatus(string(count))
		body = io.TeeReader(resp.Body, &writeCounter{0, videoSize}) // Pipe stream
	}

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

func (dm *downloadManager) displayStatus(count string) string {
	fmt.Println("Transferred", count, "bytes")
	for dm.percent < 100 {
		fmt.Printf("\rGoTube: Download progress: %%%d complete", dm.percent)
	}

	return count
}
