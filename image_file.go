package iagesdk

// ImageFile describe data structure used in image creatives
type ImageFile struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	SourceURL string `json:"sourceUrl"`
	MimeType  string `json:"mimeType"`
}
