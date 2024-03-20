package errors

import (
	"fmt"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrCipherNotFound             = Error("cipher not found")
	ErrSignatureTimestampNotFound = Error("signature timestamp not found")
	ErrInvalidCharactersInVideoID = Error("invalid characters in video id")
	ErrVideoIDMinLength           = Error("the video id must be at least 10 characters long")
	ErrReadOnClosedResBody        = Error("http: read on closed response body")
	ErrNotPlayableInEmbed         = Error("embedding of this video has been disabled")
	ErrLoginRequired              = Error("login required to confirm your age")
	ErrVideoPrivate               = Error("user restricted access to this video")
	ErrInvalidPlaylist            = Error("no playlist detected or invalid playlist ID")
)

type PlayabiltyError struct {
	Status string
	Reason string
}

func (err PlayabiltyError) Error() string {
	return fmt.Sprintf("cannot playback and download, status: %s, reason: %s", err.Status, err.Reason)
}

// UnexpectedStatusCodeError is returned on unexpected HTTP status codes
type UnexpectedStatusCodeError struct {
	Code int
}

func (err UnexpectedStatusCodeError) Error() string {
	return fmt.Sprintf("unexpected status code: %d", err.Code)
}

type PlaylistError struct {
	Reason string
}

func (err PlaylistError) Error() string {
	return fmt.Sprintf("could not load playlist: %s", err.Reason)
}
