package main

import (
	"bytes"
	"context"
	"demola/vino/internal/extractor"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"strconv"

	"golang.org/x/sync/errgroup"
)

func GetVideo(videoUrl string, ctx context.Context) error {
	// check video and extract id
	id, err := extractor.ExtractVideoID(videoUrl)
	if err != nil {
		return err
	}

	body, err := videoDataByInnertube(ctx, id)
	if err != nil {
		return err
	}

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

	return httpPostBodyBytes(ctx, "https://www.youtube.com/youtubei/v1/player?key="+c.client.key, data)
}

func httpPostBodyBytes(ctx context.Context, url string, body interface{}) ([]byte, error) {
	resp, err := c.httpPost(ctx, url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// httpPost does a HTTP POST request with a body, checks the response to be a 200 OK and returns it
func httpPost(ctx context.Context, url string, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Youtube-Client-Name", "3")
	req.Header.Set("X-Youtube-Client-Version", "2.20210721.00.00")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	resp, err := httpDo(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf(err.Error())
	}

	return resp, nil
}

var (
	consentID string
)

func httpDo(req *http.Request) (*http.Response, error) {

	client := &http.Client{}


	if client == nil {
		client = http.DefaultClient
	}

	req.Header.Set("Origin", "https://youtube.com")
	req.Header.Set("Sec-Fetch-Mode", "navigate")

	if len(consentID) == 0 {
		cconsentID = strconv.Itoa(rand.Intn(899) + 100) //nolint:gosec
	}

	req.AddCookie(&http.Cookie{
		Name:   "CONSENT",
		Value:  "YES+cb.20210328-17-p0.en+FX+" + c.consentID,
		Path:   "/",
		Domain: ".youtube.com",
	})

	res, err := client.Do(req)

	log := slog.With("method", req.Method, "url", req.URL)

	if err != nil {
		log.Debug("HTTP request failed", "error", err)
	} else {
		log.Debug("HTTP request succeeded", "status", res.Status)
	}

	return res, err
}

func getMetaDataFromVideo()

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
				err := downloadVideo(currentURL)
				fmt.Println(err)
				return err
			}
		})
	}

	return eg.Wait()
}

func main() {

}
