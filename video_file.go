package iagesdk

import (
	"time"
)

// VideoFile describe data structure used in video creatives
type VideoFile struct {
	Width     int           `json:"width"`
	Height    int           `json:"height"`
	SourceURL string        `json:"sourceUrl"`
	MimeType  string        `json:"mimeType"`
	Bitrate   int           `json:"bitrate"`
	Duration  time.Duration `json:"duration"`
}
