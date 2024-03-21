package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractVideoID(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"https://www.youtube.com/watch?v=12345678901", "12345678901", nil},
		{"https://www.youtube.com/embed/12345678901", "12345678901", nil},
		{"https://youtu.be/12345678901", "12345678901", nil},
		{"https://www.youtube.com/shorts/12345678901", "12345678901", nil},
		{"12345678901", "", nil},
		{"http://www.example.com/?v=12345678901", "", ErrInvalidCharactersInVideoID},
		{"1234567", "", ErrVideoIDMinLength},
	}

	for _, test := range tests {
		result, err := ExtractVideoID(test.input)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.err, err)
	}
}
