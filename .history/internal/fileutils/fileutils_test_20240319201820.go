package fileutils

import "testing"

func TestGetPreferredVideoExtension(t *testing.T) {
	tests := map[string]struct {
		mediaType string
		expected  string
	}{
		"Test MP4 MediaType": {
			mediaType: "video/mp4",
			expected:  ".mp4",
		},
		"Test QuickTime MediaType": {
			mediaType: "video/quicktime",
			expected:  ".mov",
		},
		"Test Unknown MediaType": {
			mediaType: "video/unknown",
			expected:  ".mp4",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := GetPreferredVideoExtension(test.mediaType)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}

func TestSanitizeFileName(t *testing.T) {
	tests := map[string]struct {
		fileName string
		expected string
	}{
		"Test Remove Invalid Characters": {
			fileName: "file:/filename?.mov",
			expected: "file filename.mov",
		},
		"Test Replace Whitespace": {
			fileName: "file  filename.mp4",
			expected: "file name.mp4",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := SanitizeFileName(test.fileName)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
