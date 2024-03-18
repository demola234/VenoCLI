package download

import (
	"bytes"
	"context"
	"demola/vino/internal/fileio"
	"demola/vino/internal/https"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type DownloadManager interface {
	Download(url string) (*DownloadResult, error)
}

type DownloadResult struct {
	Url      string `json:"url"`
	Size     int64
	Received int64
	Error    error
}

type downloadManager struct{}

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

	result.Size, result.Received = getFileSizeAndReceived(url, nil)

	return result, nil
}

func downloadFile(url string) error {
	// Placeholder implementation for downloading file
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

func download(URLs []string) error {
	eg, ctx := errgroup.WithContext(context.Background())
	for _, currentURL := range URLs {
		log.Printf("URL: %s", currentURL)
		currentURL := currentURL
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				fmt.Println("Canceled:", currentURL)
				return nil
			default:
				err := downloadYTVideo(currentURL)
				fmt.Println(err)
				return err
			}
		})
	}

	return eg.Wait()
}

func downloadYTVideo(url string) error {
	// Placeholder implementation for downloading YouTube video
	return nil
}

func (dm *downloadManager) saveAudio(outputDirectory, fileName, path string) error {
	audioFile := filepath.Join(outputDirectory, strings.TrimRight(fileName, filepath.Ext(fileName))+".mp3")

	ffmpeg, err := exec.LookPath("ffmpeg")
	if err != nil {
		return fmt.Errorf("ffmpeg not found")
	}

	cmd := exec.Command(ffmpeg, "-i", path, "-vn", "-ar", "44100", "-ac", "1", "-b:a", "128k", "-f", "mp3", audioFile)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		return err
	}
	return nil
}
