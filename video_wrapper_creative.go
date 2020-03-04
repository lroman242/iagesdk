package iagesdk

import "net/url"

type VideoWrapperCreative struct {
	Creative
	BannerType         int                 `json:"bannerType"`
	Name               string              `json:"name"`
	ThirdPartyUrls     string              `json:"thirdPartyUrls"`
	Protocols          []int               `json:"protocols"`
	VastTrackingEvents map[string][]string `json:"vastTrackingEvents"`
	BannerFile         map[string]string   `json:"bannerFile"`
}
