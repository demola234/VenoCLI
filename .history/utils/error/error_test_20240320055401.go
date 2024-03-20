package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	assert.Equal(t, "cipher not found", ErrCipherNotFound.Error())
	assert.Equal(t, "signature timestamp not found", ErrSignatureTimestampNotFound.Error())
	assert.Equal(t, "invalid characters in video id", ErrInvalidCharactersInVideoID.Error())
	assert.Equal(t, "the video id must be at least 10 characters long", ErrVideoIDMinLength.Error())
	assert.Equal(t, "http: read on closed response body", ErrReadOnClosedResBody.Error())
	assert.Equal(t, "embedding of this video has been disabled", ErrNotPlayableInEmbed.Error())
	assert.Equal(t, "login required to confirm your age", ErrLoginRequired.Error())
	assert.Equal(t, "user restricted access to this video", ErrVideoPrivate.Error())
	assert.Equal(t, "no playlist detected or invalid playlist ID", ErrInvalidPlaylist.Error())

	playErr := PlayabiltyError{Status: "error", Reason: "blocked in your country"}
	assert.Equal(t, "cannot playback and download, status: error, reason: blocked in your country", playErr.Error())

	unexpectedStatusCodeErr := UnexpectedStatusCodeError{Code: 404}
	assert.Equal(t, "unexpected status code: 404", unexpectedStatusCodeErr.Error())

	playlistErr := PlaylistError{Reason: "playlist not found"}
	assert.Equal(t, "could not load playlist: playlist not found", playlistErr.Error())
}
