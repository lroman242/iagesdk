package iagesdk

// VideoInlineCreative describe iAGE`s video inline video creative
type VideoInlineCreative struct {
	Creative

	BannerType         int                 `json:"bannerType"`
	Name               string              `json:"name"`
	ClickURL           string              `json:"clickUrl"`
	LastDomain         string              `json:"lastDomain"`
	ThirdPartyUrls     []string            `json:"thirdPartyUrls"`
	VideoStartDelays   []int               `json:"videoStartDelays"`
	SkipOffset         int                 `json:"skipoffset"`
	Protocols          []int               `json:"protocols"`
	VastTrackingEvents map[string][]string `json:"vastTrackingEvents"`
	VideoFirstFile     VideoFile           `json:"videoFirstFile"`
	VideoSecondFile    VideoFile           `json:"videoSecondFile"`
}
