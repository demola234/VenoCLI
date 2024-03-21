package main

import (
	"context"
	"demola/vino/internal/extractor"
	"fmt"
	"log"

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
