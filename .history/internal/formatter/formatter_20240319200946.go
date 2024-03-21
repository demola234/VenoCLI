package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func prettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func PrettyJsonStringAsLines(logline string) []string {
	pretty, err := prettyPrint([]byte(logline))
	if err != nil {
		return []string{logline}
	}

	var splitLines []string
	for _, row := range bytes.Split(pretty, []byte("\n")) {
		splitLines = append(splitLines, string(row))
	}

	return splitLines
}

func FormatTime(t time.Time) string {
	if t.IsZero() {
		return "-"
	}
	local, err := time.LoadLocation("Local")
	if err != nil {
		return "-"
	}
	tLocal := t.In(local)
	return tLocal.Format("2006-01-02T15:04:05")
}

func FormatTimeNs(t int64) string {
	tm := time.Unix(0, t).UTC()
	return FormatTime(tm)
}

func pluralize(s string, q float64) string {
	if q > 1 {
		return s + "s"
	}
	return s
}

func FormatTimeNsSinceNow(t int64) string {
	tm := time.Unix(0, t).UTC()
	since := time.Since(tm)

	seconds := int(since.Seconds())
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24
	years := days / 365

	if years > 0 {
		return fmt.Sprintf("%d year%s", years, pluralize("s", float64(years)))
	}
	if days > 0 {
		return fmt.Sprintf("%d day%s", days, pluralize("s", float64(days)))
	}
	if hours > 0 {
		return fmt.Sprintf("%d hour%s", hours, pluralize("s", float64(hours)))
	}
	if minutes > 0 {
		return fmt.Sprintf("%d minute%s", minutes, pluralize("s", float64(minutes)))
	}
	return fmt.Sprintf("%d second%s", seconds, pluralize("s", float64(seconds)))
}
