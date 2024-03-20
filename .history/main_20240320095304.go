package main

import (
	"bytes"
	"context"
	"demola/vino/internal/extractor"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func downloadVideo(videoUrl string) error {
	// check video and extract id
	id, err := extractor.ExtractVideoID(videoUrl)
	if err != nil {
		return err
	}

}

type innertubeRequest struct {
	VideoID         string            `json:"videoId,omitempty"`
	BrowseID        string            `json:"browseId,omitempty"`
	Continuation    string            `json:"continuation,omitempty"`
	Context         inntertubeContext `json:"context"`
	PlaybackContext *playbackContext  `json:"playbackContext,omitempty"`
	ContentCheckOK  bool              `json:"contentCheckOk,omitempty"`
	RacyCheckOk     bool              `json:"racyCheckOk,omitempty"`
	Params          string            `json:"params"`
}

type playbackContext struct {
	ContentPlaybackContext contentPlaybackContext `json:"contentPlaybackContext"`
}

type contentPlaybackContext struct {
	// SignatureTimestamp string `json:"signatureTimestamp"`
	HTML5Preference string `json:"html5Preference"`
}

type inntertubeContext struct {
	Client innertubeClient `json:"client"`
}

type innertubeClient struct {
	HL                string `json:"hl"`
	GL                string `json:"gl"`
	ClientName        string `json:"clientName"`
	ClientVersion     string `json:"clientVersion"`
	AndroidSDKVersion int    `json:"androidSDKVersion,omitempty"`
	UserAgent         string `json:"userAgent,omitempty"`
	TimeZone          string `json:"timeZone"`
	UTCOffset         int    `json:"utcOffsetMinutes"`
}

// client info for the innertube API
type clientInfo struct {
	name           string
	key            string
	version        string
	userAgent      string
	androidVersion int
}


func prepareInnertubeContext(clientInfo clientInfo) inntertubeContext {
	return inntertubeContext{
		Client: innertubeClient{
			HL:                "en",
			GL:                "US",
			TimeZone:          "UTC",
			ClientName:        clientInfo.name,
			ClientVersion:     clientInfo.version,
			AndroidSDKVersion: clientInfo.androidVersion,
			UserAgent:         clientInfo.userAgent,
		},
	}
}
func (videoDataByInnertube(ctx context., id string) ([]byte, error) {
	data := innertubeRequest{
		VideoID:        id,
		Context:        prepareInnertubeContext(*c.client),
		ContentCheckOK: true,
		RacyCheckOk:    true,
		Params:         playerParams,
		PlaybackContext: &playbackContext{
			ContentPlaybackContext: contentPlaybackContext{
				// SignatureTimestamp: sts,
				HTML5Preference: "HTML5_PREF_WANTS",
			},
		},
	}

	return c.httpPostBodyBytes(ctx, "https://www.youtube.com/youtubei/v1/player?key="+c.client.key, data)
}

func getMetaDataFromVideo(id string) (string, string, error) {
	url := "https://www.youtube.com/youtubei/v1/player?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"

	// Prepare request body
	requestBody := map[string]interface{}{
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

	// Convert request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Make POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch video metadata: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to fetch video metadata: %s", resp.Status)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal response body into YoutubePayload struct
	youtubePayload, err := UnmarshalYoutubePayload(body)
	if err != nil {
		return "", "", fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	// // Parse JSON response
	// var responseData map[string]interface{}
	// if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
	// 	return "", "", fmt.Errorf("failed to decode response body: %v", err)
	// }

	res := fmt.Sprintln(youtubePayload.StreamingData.Formats[0].URL)

	fmt.Println(res)

	// Extract video metadata
	downloadUrl := res
	fileName := youtubePayload.VideoDetails.Title
	// description := videoDetails["shortDescription"].(string)

	return fileName, downloadUrl, nil
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
