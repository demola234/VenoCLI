package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
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

func FormatTimeNsSinceNow(t int64) string {
	tm := time.Unix(0, t).UTC()
	since := time.Since(tm)
	if secs := since.Seconds(); secs > 0 && secs < 60 {
		val := math.Floor(secs)
		out := fmt.Sprintf("%.0f second", val)
		return pluralize(out, val)
	}
	if mins := since.Minutes(); mins > 1 && mins < 60 {
		val := math.Floor(mins)
		out := fmt.Sprintf("%.0f minute", val)
		return pluralize(out, val)
	}
	if hrs := since.Hours(); hrs > 1 && hrs < 24 {
		val := math.Floor(hrs)
		out := fmt.Sprintf("%.0f hour", val)
		return pluralize(out, val)
	}
	if days := since.Hours() / 24; days > 1 && days < 365.25 {
		val := math.Floor(days)
		out := fmt.Sprintf("%.0f day", val)
		return pluralize(out, val)
	}
	if years := since.Hours() / 24 / 365.25; years > 1 {
		val := math.Floor(years)
		out := fmt.Sprintf("%.0f year", val)
		return pluralize(out, val)
	}
	return ""
}
