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

	"golang.org/x/sync/errgroup"
)

func GetVideo(videoUrl string, ctx context.Context) (*fileutils.Video, error) {
	// check video and extract id
	id, err := extractor.ExtractVideoID(videoUrl)
	if err != nil {
		return nil, fmt.Errorf("extractVideoID failed: %w", err)
	}

	body, err := videoDataByInnertube(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("extractVideoID failed: %w", err)
	}

	v := fileutils.Video{
		ID: id,
	}

	if err = v.ParseVideoInfo(body); err == nil {
		return &v, nil
	}

	// If the uploader has disabled embedding the video on other sites, parse video page
	if errse.Is(err, errors.ErrNotPlayableInEmbed) {
		// additional parameters are required to access clips with sensitiv content
		html, err := httpGetBodyBytes(ctx, "https://www.youtube.com/watch?v="+id+"&bpctr=9999999999&has_verified=1")
		if err != nil {
			return nil, err
		}

		return &v, v.ParseVideoPage(html)
	}

	// undefined error
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

// httpGet does a HTTP GET request, checks the response to be a 200 OK and returns it
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
		return nil, f
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

	return httpPostBodyBytes(ctx, "https://www.youtube.com/youtubei/v1/player?key="+"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8", data)
}

func httpPostBodyBytes(ctx context.Context, url string, body interface{}) ([]byte, error) {
	resp, err := httpPost(ctx, url, body)
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

func httpDo(req *http.Request) (*http.Response, error) {

	client := &http.Client{}

	if client == nil {
		client = http.DefaultClient
	}

	req.Header.Set("Origin", "https://youtube.com")
	req.Header.Set("Sec-Fetch-Mode", "navigate")

	consentID := strconv.Itoa(rand.Intn(899) + 100) //nolint:gosec

	req.AddCookie(&http.Cookie{
		Name:   "CONSENT",
		Value:  "YES+cb.20210328-17-p0.en+FX+" + consentID,
		Path:   "/",
		Domain: ".youtube.com",
	})

	res, err := client.Do(req)

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
