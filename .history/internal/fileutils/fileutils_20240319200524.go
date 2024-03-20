package fileutils

import (
	"mime"
	"regexp"
)

const defaultExtension = ".mp4"


var videoExtensions = []string{
	