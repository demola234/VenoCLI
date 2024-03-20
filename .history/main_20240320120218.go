package main

import (
	"bytes"
	"context"
	"demola/vino/internal/extractor"
	"demola/vino/internal/fileutils"
	"demola/vino/utils/errors"
	"encoding/json"
	errse "errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

func downloadVideo(url string) (*fileutils.Video, error) {
	return GetVideo(url, context.Background())
}

func GetVideo(videoUrl string, ctx context.Context) (*fileutils.Video, error) {
	id, err := extractor.ExtractVideoID(videoUrl)
	if err != nil {
		return nil, fmt.Errorf("extractVideoID failed: %w", err)
	}

	body, err := videoDataByInnertube(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("videoDataByInnertube failed: %w", err)
	}

	v := fileutils.Video{
		ID: id,
	}

	if err = v.ParseVideoInfo(body); err == nil {
		return &v, nil
	}

	if errse.Is(err, errors.ErrNotPlayableInEmbed) {
		html, err := httpGetBodyBytes(ctx, "https://www.youtube.com/watch?v="+id+"&bpctr=9999999999&has_verified=1")
		if err != nil {
			return nil, err
		}

		if err := v.ParseVideoPage(html); err != nil {
			return nil, fmt.Errorf("failed to parse video page: %w", err)
		}

		return &v, nil
	}

	return &v, err
}

func httpGetBodyBytes(ctx context.Context, url string) ([]byte, error) {
	resp, err := httpGet(ctx, url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func httpGet(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpDo(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("non-OK status code: %d", resp.StatusCode)
	}

	return resp, nil
}

func videoDataByInnertube(ctx context.Context, id string) ([]byte, error) {
	data := map[string]interface{}{
		// your data
	}

	return httpPostBodyBytes(ctx, "https://www.youtube.com/youtubei/v1/player?key="+"your_api_key", data)
}

func httpPostBodyBytes(ctx context.Context, url string, body interface{}) ([]byte, error) {
	resp, err := httpPost(ctx, url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func httpPost(ctx context.Context, url string, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	resp, err := httpDo(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("non-OK status code: %d", resp.StatusCode)
	}

	return resp, nil
}

func httpDo(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	if client == nil {
		client = http.DefaultClient
	}

	req.Header.Set("Origin", "https://youtube.com")
	req.Header.Set("Sec-Fetch-Mode", "navigate")

	consentID := strconv.Itoa(rand.Intn(899) + 100)
	req.AddCookie(&http.Cookie{
		Name:   "CONSENT",
		Value:  "YES+cb.20210328-17-p0.en+FX+" + consentID,
		Path:   "/",
		Domain: ".youtube.com",
	})

	return client.Do(req)
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
				wee, err := downloadVideo(currentURL)
				fmt.Println(wee)
				fmt.Println(err)
				return err
			}
		})
	}

	return eg.Wait()
}

func main() {
	ytUrl := fmt.Sprintln("https://www.youtube.com/watch?v=7qMAmI_Lzv4&t=33s")
	err := download([]string{ytUrl})

	zerolog.Logger
}
