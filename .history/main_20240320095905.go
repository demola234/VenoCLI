package main

import (
	"context"
	"demola/vino/internal/extractor"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func GetVideo(videoUrl string) error {
	// check video and extract id
	id, err := extractor.ExtractVideoID(videoUrl)
	if err != nil {
		return err
	}


	body, err := c.videoDataByInnertube(ctx, id)
	if err != nil {
		return nil, err
	}

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
