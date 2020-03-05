package iagesdk

// NativeVideoFeedCreative describe iAGE`s data structure for video feed creative
type NativeVideoFeedCreative struct {
	Creative

	ClickURL   string `json:"clickUrl"`
	LastDomain string `json:"lastDomain"`

	ThirdPartyUrls     []string            `json:"thirdPartyUrls"`
	SkipOffset         int                 `json:"skipoffset"`
	Protocols          []int               `json:"protocols"`
	VastTrackingEvents map[string][]string `json:"vastTrackingEvents"`
	Title              string              `json:"title"`
	Brand              string              `json:"brand"`
	Description        string              `json:"description"`
	Button             string              `json:"button"`
	Domain             string              `json:"domain"`
	VideoFirstFile     VideoFile           `json:"videoFirstFile"`
	VideoSecondFile    VideoFile           `json:"videoSecondFile"`
	NativeIconFile     ImageFile           `json:"nativeIconFile"`
	NativeImageFile    ImageFile           `json:"nativeImageFile"`
}
