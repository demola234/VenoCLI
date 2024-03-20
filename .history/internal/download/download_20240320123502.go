package downloader

// import (
// 	"context"
// 	"demola/vino/internal/fileutils"
// 	"errors"
// 	"io"
// 	"os"
// 	"os/exec"
// 	"path/filepath"

// 	logger "demola/vino/config/log"

// 	"github.com/vbauerster/mpb/v5"
// 	"github.com/vbauerster/mpb/v5/decor"
// )

// // Downloader offers high level functions to download videos into files
// type Downloader struct {
// 	OutputDir string // optional directory to store the files
// }

// func (dl *Downloader) getOutputFile(v *fileutils.Video, format *fileutils.Format, outputFile string) (string, error) {
// 	if outputFile == "" {
// 		outputFile = fileutils.SanitizeFilename(v.Title)
// 		outputFile += fileutils.PickIdealFileExtension(format.MimeType)
// 	}

// 	if dl.OutputDir != "" {
// 		if err := os.MkdirAll(dl.OutputDir, 0o755); err != nil {
// 			return "", err
// 		}
// 		outputFile = filepath.Join(dl.OutputDir, outputFile)
// 	}

// 	return outputFile, nil
// }

// // Download : Starting download video by arguments.
// func (dl *Downloader) Download(ctx context.Context, v *fileutils.Video, format *fileutils.Format, outputFile string) error {
// 	// logger.Logger.Info(
// 	// 	"Downloading video",
// 	// 	"id", v.ID,
// 	// 	"quality", format.Quality,
// 	// 	"mimeType", format.MimeType,
// 	// )
// 	destFile, err := dl.getOutputFile(v, format, outputFile)
// 	if err != nil {
// 		return err
// 	}

// 	// Create output file
// 	out, err := os.Create(destFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer out.Close()

// 	return dl.videoDLWorker(ctx, out, v, format)
// }

// // DownloadComposite : Downloads audio and video streams separately and merges them via ffmpeg.
// func (dl *Downloader) DownloadComposite(ctx context.Context, outputFile string, v *fileutils.Video, quality string, mimetype, language string) error {
// 	videoFormat, audioFormat, err1 := getVideoAudioFormats(v, quality, mimetype, language)
// 	if err1 != nil {
// 		return err1
// 	}

// 	log := logger.Logger.With("id", v.ID)

// 	// log.Info(
// 	// 	"Downloading composite video",
// 	// 	"videoQuality", videoFormat.QualityLabel,
// 	// 	"videoMimeType", videoFormat.MimeType,
// 	// 	"audioMimeType", audioFormat.MimeType,
// 	// )

// 	destFile, err := dl.getOutputFile(v, videoFormat, outputFile)
// 	if err != nil {
// 		return err
// 	}
// 	outputDir := filepath.Dir(destFile)

// 	// Create temporary video file
// 	videoFile, err := os.CreateTemp(outputDir, "youtube_*.m4v")
// 	if err != nil {
// 		return err
// 	}
// 	defer os.Remove(videoFile.Name())

// 	// Create temporary audio file
// 	audioFile, err := os.CreateTemp(outputDir, "youtube_*.m4a")
// 	if err != nil {
// 		return err
// 	}
// 	defer os.Remove(audioFile.Name())

// 	// log.Debug("Downloading video file...")
// 	err = dl.videoDLWorker(ctx, videoFile, v, videoFormat)
// 	if err != nil {
// 		return err
// 	}

// 	log.Debug("Downloading audio file...")
// 	err = dl.videoDLWorker(ctx, audioFile, v, audioFormat)
// 	if err != nil {
// 		return err
// 	}

// 	//nolint:gosec
// 	ffmpegVersionCmd := exec.Command("ffmpeg", "-y",
// 		"-i", videoFile.Name(),
// 		"-i", audioFile.Name(),
// 		"-c", "copy", // Just copy without re-encoding
// 		"-shortest", // Finish encoding when the shortest input stream ends
// 		destFile,
// 		"-loglevel", "warning",
// 	)
// 	ffmpegVersionCmd.Stderr = os.Stderr
// 	ffmpegVersionCmd.Stdout = os.Stdout
// 	log.Info("merging video and audio", "output", destFile)

// 	return ffmpegVersionCmd.Run()
// }

// func getVideoAudioFormats(v *fileutils.Video, quality string, mimetype, language string) (*fileutils.Format, *fileutils.Format, error) {
// 	var videoFormats, audioFormats fileutils.FormatList

// 	formats := v.Formats
// 	if mimetype != "" {
// 		formats = formats.Type(mimetype)
// 	}

// 	videoFormats = formats.Type("video").AudioChannels(0)
// 	audioFormats = formats.Type("audio")

// 	if quality != "" {
// 		videoFormats = videoFormats.Quality(quality)
// 	}

// 	if language != "" {
// 		audioFormats = audioFormats.Language(language)
// 	}

// 	if len(videoFormats) == 0 {
// 		return nil, nil, errors.New("no video format found after filtering")
// 	}

// 	if len(audioFormats) == 0 {
// 		return nil, nil, errors.New("no audio format found after filtering")
// 	}

// 	videoFormats.Sort()
// 	audioFormats.Sort()

// 	return &videoFormats[0], &audioFormats[0], nil
// }

// func (dl *Downloader) videoDLWorker(ctx context.Context, out *os.File, video *fileutils.Video, format *fileutils.Format) error {
// 	stream, size, err := dl.GetStreamContext(ctx, video, format)
// 	if err != nil {
// 		return err
// 	}

// 	prog := &progress{
// 		contentLength: float64(size),
// 	}

// 	// create progress bar
// 	progress := mpb.New(mpb.WithWidth(64))
// 	bar := progress.AddBar(
// 		int64(prog.contentLength),

// 		mpb.PrependDecorators(
// 			decor.CountersKibiByte("% .2f / % .2f"),
// 			decor.Percentage(decor.WCSyncSpace),
// 		),
// 		mpb.AppendDecorators(
// 			decor.EwmaETA(decor.ET_STYLE_GO, 90),
// 			decor.Name(" ] "),
// 			decor.EwmaSpeed(decor.UnitKiB, "% .2f", 60),
// 		),
// 	)

// 	reader := bar.ProxyReader(stream)
// 	mw := io.MultiWriter(out, prog)
// 	_, err = io.Copy(mw, reader)
// 	if err != nil {
// 		return err
// 	}

// 	progress.Wait()
// 	return nil
// }
