package iagesdk

import "net/url"

type VideoInlineCreative struct {
	Creative

	BannerType         int                 `json:"bannerType"`
	Name               string              `json:"name"`
	ClickUrl           string              `json:"clickUrl"`
	LastDomain         string              `json:"lastDomain"`
	ThirdPartyUrls     []string            `json:"thirdPartyUrls"`
	VideoStartDelays   []int               `json:"videoStartDelays"`
	SkipOffset         int                 `json:"skipoffset"`
	Protocols          []int               `json:"protocols"`
	VastTrackingEvents map[string][]string `json:"vastTrackingEvents"`
	VideoFirstFile     VideoFile           `json:"videoFirstFile"`
	VideoSecondFile    VideoFile           `json:"videoSecondFile"`
}
