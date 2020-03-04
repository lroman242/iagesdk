package iagesdk

type ImageFile struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	SourceUrl string `json:"sourceUrl"`
	MimeType  string `json:"mimeType"`
}
