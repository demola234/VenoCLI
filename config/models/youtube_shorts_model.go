package main

import "encoding/json"

func UnmarshalYoutubeShortPayload(data []byte) (YoutubeShortPayload, error) {
	var r YoutubeShortPayload
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *YoutubeShortPayload) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type YoutubeShortPayload struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type Data struct {
	AudioFormats     []AudioFormat     `json:"audio_formats"`
	DefaultSelected  int64             `json:"default_selected"`
	Duration         int64             `json:"duration"`
	DurationLabel    string            `json:"durationLabel"`
	FromCache        bool              `json:"fromCache"`
	ID               string            `json:"id"`
	Key              string            `json:"key"`
	Thumbnail        string            `json:"thumbnail"`
	ThumbnailFormats []ThumbnailFormat `json:"thumbnail_formats"`
	Title            string            `json:"title"`
	TitleSlug        string            `json:"titleSlug"`
	URL              string            `json:"url"`
	VideoFormats     []VideoFormat     `json:"video_formats"`
}

type AudioFormat struct {
	Label   string      `json:"label"`
	Quality int64       `json:"quality"`
	URL     interface{} `json:"url"`
}

type ThumbnailFormat struct {
	Label   string `json:"label"`
	Quality string `json:"quality"`
	URL     string `json:"url"`
	Value   string `json:"value"`
}

type VideoFormat struct {
	DefaultSelected int64   `json:"default_selected"`
	Height          int64   `json:"height"`
	Label           string  `json:"label"`
	Quality         int64   `json:"quality"`
	URL             *string `json:"url"`
	Width           int64   `json:"width"`
}
