package iagesdk

import (
	"time"
)

type VideoFile struct {
	Width     int           `json:"width"`
	Height    int           `json:"height"`
	SourceUrl string        `json:"sourceUrl"`
	MimeType  string        `json:"mimeType"`
	Bitrate   int           `json:"bitrate"`
	Duration  time.Duration `json:"duration"`
}
