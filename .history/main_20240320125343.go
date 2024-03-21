package main

import (
	"bytes"
	"context"
	"demola/vino/internal/extractor"
	"demola/vino/internal/fileutils"
	"demola/vino/utils/errors"
	"demola/vino/internal/download"
	"encoding/json"
	errse "errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
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
				return fmt.Errorf("failed to parse video page: %w", err)
			}
		})
	}

	return eg.Wait()
}

var (
	OutputDir string // optional directory to store the file
)

func getOutputFile(v *fileutils.Video, format *fileutils.Format, outputFile string) (string, error) {
	if outputFile == "" {
		outputFile = fileutils.SanitizeFileName(v.Title)
		outputFile += fileutils.PreferredVideoExtension(format.MimeType)
	}

	if OutputDir != "" {
		if err := os.MkdirAll(OutputDir, 0o755); err != nil {
			return "", err
		}
		outputFile = filepath.Join(OutputDir, outputFile)
	}

	return outputFile, nil
}

func getVideoAudioFormats(v *fileutils.Video, quality string, mimetype, language string) (*fileutils.Format, *fileutils.Format, error) {
	var videoFormats, audioFormats fileutils.FormatList

	formats := v.Formats
	if mimetype != "" {
		formats = formats.Type(mimetype)
	}

	videoFormats = formats.Type("video").AudioChannels(0)
	audioFormats = formats.Type("audio")

	if quality != "" {
		videoFormats = videoFormats.Quality(quality)
	}

	if language != "" {
		audioFormats = audioFormats.Language(language)
	}

	if len(videoFormats) == 0 {
		return nil, nil, errse.New("no video format found after filtering")
	}

	if len(audioFormats) == 0 {
		return nil, nil, errse.New("no audio format found after filtering")
	}

	videoFormats.Sort()
	audioFormats.Sort()

	return &videoFormats[0], &audioFormats[0], nil
}

func videoDLWorker(ctx context.Context, out *os.File, video *fileutils.Video, format *fileutils.Format) error {
	stream, size, err := dl.GetStreamContext(ctx, video, format)
	if err != nil {
		return err
	}

	prog := &downloader.Progress{
	}

	// create progress bar
	// progress := mpb.New(mpb.WithWidth(64))
	// bar := progress.AddBar(
	// 	// int64(prog.contentLength),

	// 	mpb.PrependDecorators(
	// 		decor.CountersKibiByte("% .2f / % .2f"),
	// 		decor.Percentage(decor.WCSyncSpace),
	// 	),
	// 	mpb.AppendDecorators(
	// 		decor.EwmaETA(decor.ET_STYLE_GO, 90),
	// 		decor.Name(" ] "),
	// 		decor.EwmaSpeed(decor.UnitKiB, "% .2f", 60),
	// 	),
	// )

	reader := bar.ProxyReader(stream)
	mw := io.MultiWriter(out, prog)
	_, err = io.Copy(mw, reader)
	if err != nil {
		return err
	}

	progress.Wait()
	return nil
}

func DownloadFile(ctx context.Context, v *fileutils.Video, format *fileutils.Format, outputFile string) error {

	destFile, err := getOutputFile(v, format, outputFile)
	if err != nil {
		return err
	}

	// Create output file
	out, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer out.Close()

	return dl.videoDLWorker(ctx, out, v, format)
}

func DownloadComposite(ctx context.Context, outputFile string, v *fileutils.Video, quality string, mimetype, language string) error {
	videoFormat, audioFormat, err1 := getVideoAudioFormats(v, quality, mimetype, language)
	if err1 != nil {
		return err1
	}

	destFile, err := getOutputFile(v, videoFormat, outputFile)
	if err != nil {
		return err
	}
	outputDir := filepath.Dir(destFile)

	// Create temporary video file
	videoFile, err := os.CreateTemp(outputDir, "youtube_*.m4v")
	if err != nil {
		return err
	}
	defer os.Remove(videoFile.Name())

	// Create temporary audio file
	audioFile, err := os.CreateTemp(outputDir, "youtube_*.m4a")
	if err != nil {
		return err
	}
	defer os.Remove(audioFile.Name())

	err = dl.videoDLWorker(ctx, videoFile, v, videoFormat)
	if err != nil {
		return err
	}

	err = dl.videoDLWorker(ctx, audioFile, v, audioFormat)
	if err != nil {
		return err
	}

	//nolint:gosec
	ffmpegVersionCmd := exec.Command("ffmpeg", "-y",
		"-i", videoFile.Name(),
		"-i", audioFile.Name(),
		"-c", "copy", // Just copy without re-encoding
		"-shortest", // Finish encoding when the shortest input stream ends
		destFile,
		"-loglevel", "warning",
	)
	ffmpegVersionCmd.Stderr = os.Stderr
	ffmpegVersionCmd.Stdout = os.Stdout

	return ffmpegVersionCmd.Run()
}

func downloadVideo(id string) error {
	video, format, err := getVideoWithFormat(id)
	if err != nil {
		return err
	}

	log.Println("download to directory", outputDir)

	if strings.HasPrefix(outputQuality, "hd") {
		if err := checkFFMPEG(); err != nil {
			return err
		}
		return DownloadComposite(context.Background(), outputFile, video, outputQuality, mimetype, language)
	}

	return DownloadFile(context.Background(), video, format, outputFile)
}

func main() {
	ytUrl := fmt.Sprintln("https://www.youtube.com/watch?v=7qMAmI_Lzv4&t=33s")
	err := download([]string{ytUrl})

	fmt.Errorf(err.Error())
}
