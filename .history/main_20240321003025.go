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
	"os"
	"strconv"
	"strings"

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
			return nil, fmt.Errorf("failed to parse video page: %w", err)
		}

		if err := v.ParseVideoPage(html); err != nil {
			return nil, fmt.Errorf("failed to parse video page: %w", err)
		}

		return &v, nil
	}

	err =downloading(v.Formats[0].URL, v.Title)

	return &v, err
}

func httpGetBodyBytes(ctx context.Context, url string) ([]byte, error) {
	resp, err := httpGet(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse video page: %w", err)
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func httpGet(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse video page: %w", err)
	}

	resp, err := httpDo(req)
	if err != nil {
		return nil, fmt.Errorf("failed to parse video page: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("non-OK status code: %d", resp.StatusCode)
	}

	return resp, nil
}

func videoDataByInnertube(ctx context.Context, id string) ([]byte, error) {
	data := map[string]interface{}{
		"context": map[string]interface{}{
			"client": map[string]interface{}{
				"hl":               "en",
				"clientName":       "WEB",
				"clientVersion":    "2.20210721.00.00",
				"clientFormFactor": "UNKNOWN_FORM_FACTOR",
				"clientScreen":     "WATCH",
				"mainAppWebInfo": map[string]interface{}{
					"graftUrl": "/watch?v=" + id,
				},
			},
			"user": map[string]interface{}{
				"lockedSafetyMode": false,
			},
			"request": map[string]interface{}{
				"useSsl":                  true,
				"internalExperimentFlags": []interface{}{},
				"consistencyTokenJars":    []interface{}{},
			},
		},
		"videoId": id,
		"playbackContext": map[string]interface{}{
			"contentPlaybackContext": map[string]interface{}{
				"vis":                   0,
				"splay":                 false,
				"autoCaptionsDefaultOn": false,
				"autonavState":          "STATE_NONE",
				"html5Preference":       "HTML5_PREF_WANTS",
				"lactMilliseconds":      "-1",
			},
		},
		"racyCheckOk":    false,
		"contentCheckOk": false,
	}

	return httpPostBodyBytes(ctx, "https://www.youtube.com/youtubei/v1/player?key="+"AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w", data)
}

func httpPostBodyBytes(ctx context.Context, url string, body interface{}) ([]byte, error) {
	resp, err := httpPost(ctx, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse video page: %w", err)
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func httpPost(ctx context.Context, url string, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse video page: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse video page: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	resp, err := httpDo(req)
	if err != nil {
		return nil, fmt.Errorf("failed to parse video page: %w", err)
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

type PassThru struct {
	io.Reader
	total int64 // Total # of bytes transferred
}

func downloading(url string, title string) error {
	videoPath := fmt.Sprintf("%s.mp4", title)

	var body io.Reader

	output, err := os.Create(videoPath)
	if err != nil {
		return fmt.Errorf("GoTube: Failed to create video file: %v", err)
	}
	defer output.Close()

	// Create some random input data.
	src := bytes.NewBufferString(strings.Repeat("Some random input data", 1000))
	_ = &PassThru{Reader: src}

	_, err = io.Copy(output, body)

	return nil
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
				return fmt.Errorf("failed to parse video page: %w", err)
			}
		})
	}

	return eg.Wait()
}

func main() {
	ytUrl := fmt.Sprintln("https://www.youtube.com/watch?v=7qMAmI_Lzv4&t=33s")
	err := download([]string{ytUrl})

	fmt.Errorf(err.Error())
}
