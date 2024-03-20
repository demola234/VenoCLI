package formatter

import (
	"testing"
	"time"
)

func TestPrettyPrint(t *testing.T) {
	input := []byte(`{"key": "value"}`)
	expected := "{\n  \"key\": \"value\"\n}"

	result, err := prettyPrint(input)
	if err != nil {
		t.Errorf("prettyPrint returned an error: %v", err)
	}

	if string(result) != expected {
		t.Errorf("prettyPrint returned unexpected result, got: %s, want: %s", result, expected)
	}
}

func TestPrettyJsonStringAsLines(t *testing.T) {
	input := `{"key": "value"}`
	expected := []string{"{", "  \"key\": \"value\"", "}"}

	result := PrettyJsonStringAsLines(input)

	if len(result) != len(expected) {
		t.Errorf("PrettyJsonStringAsLines returned unexpected number of lines, got: %d, want: %d", len(result), len(expected))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("PrettyJsonStringAsLines returned unexpected line, got: %s, want: %s", result[i], expected[i])
		}
	}
}

func TestFormatTime(t *testing.T) {
	now := time.Now()
	expected := now.Format("2006-01-02T15:04:05")

	result := FormatTime(now)

	if result != expected {
		t.Errorf("FormatTime returned unexpected result, got: %s, want: %s", result, expected)
	}
}

func TestFormatTimeNs(t *testing.T) {
	now := time.Now().UnixNano()
	expected := FormatTime(time.Unix(0, now))

	result := FormatTimeNs(now)

	if result != expected {
		t.Errorf("FormatTimeNs returned unexpected result, got: %s, want: %s", result, expected)
	}
}

func TestFormatTimeNsSinceNow(t *testing.T) {
	pastTime := time.Now().Add(-time.Hour)
	expected := "1 hour"

	result := FormatTimeNsSinceNow(pastTime.UnixNano())

	if result != expected {
		t.Errorf("FormatTimeNsSinceNow returned unexpected result, got: %s, want: %s", result, expected)
	}
	
}
