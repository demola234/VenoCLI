package fileutils

import (
	"mime"
	"regexp"
)

const defaultVideoExtension = ".mp4"

var videoExtensions = map[string]string{
	"video/quicktime":  ".mov",
	"video/x-msvideo":  ".avi",
	"video/x-matroska": ".mkv",
	"video/mpeg":       ".mpeg",
	"video/webm":       ".webm",
	"video/3gpp2":      ".3g2",
	"video/x-flv":      ".flv",
	"video/3gpp":       ".3gp",
	"video/mp4":        ".mp4",
	"video/ogg":        ".ogv",
	"video/mp2t":       ".ts",
}

// GetPreferredVideoExtension returns the preferred file extension for a given video media type.
func GetPreferredVideoExtension(mediaType string) string {
	mediaType, _, err := mime.ParseMediaType(mediaType)
	if err != nil {
		return defaultVideoExtension
	}

	if extension, ok := videoExtensions[mediaType]; ok {
		return extension
	}

	extensions, err := mime.ExtensionsByType(mediaType)
	if err != nil || extensions == nil {
		return defaultVideoExtension
	}

	return extensions[0]
}

// SanitizeFileName removes invalid characters from a file name.
func SanitizeFileName(fileName string) string {
	invalidCharsRegex := regexp.MustCompile(`[:/<>\:"\\|?*]`)
	whitespaceRegex := regexp.MustCompile(`\s+`)

	fileName = invalidCharsRegex.ReplaceAllString(fileName, "")
	fileName = whitespaceRegex.ReplaceAllString(fileName, " ")

	return fileName
}
