package formatter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPrettyJsonStringAsLines(t *testing.T) {
	jsonStr := `{"key1": "value1", "key2": "value2"}`

	lines := PrettyJsonStringAsLines(jsonStr)

	require.Equal(t, 3, len(lines))
	require.Contains(t, lines[0], "{")
	require.Contains(t, lines[1], "  \"key1\": \"value1\",")
	require.Contains(t, lines[2], "}")
}

func TestFormatTime(t *testing.T) {
	tm := time.Date(2020, 1, 15, 17, 25, 0, 0, time.UTC)

	formatted := FormatTime(tm)

	require.Equal(t, "2020-01-15T17:25:00", formatted)
}

func TestFormatTimeNs(t *testing.T) {
	ns := int64(1579092000000000000) // Jan 15, 2020 17:25:00 UTC

	formatted := FormatTimeNs(ns)

	require.Equal(t, "2020-01-15T17:25:00", formatted)
}

func TestFormatTimeNsSinceNow(t *testing.T) {
	now := time.Now()
	ns := now.Add(-2 * time.Hour).UnixNano()

	formatted := FormatTimeNsSinceNow(ns)

	require.Equal(t, "2 hours", formatted)
}

func TestPluralize(t *testing.T) {
	plural := pluralize("item", 2.0)
	require.Equal(t, "items", plural)

	singular := pluralize("item", 1.0)
	require.Equal(t, "item", singular)
}
