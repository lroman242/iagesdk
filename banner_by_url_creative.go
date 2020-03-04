package iagesdk

import "net/url"

type BannerByUrlCreative struct {
	Creative
	BannerType     int       `json:"bannerType"`
	Name           string    `json:"name"`
	ClickUrl       string    `json:"clickUrl"`
	LastDomain     string    `json:"lastDomain"`
	ThirdPartyUrls []string  `json:"thirdPartyUrls"`
	BannerFile     ImageFile `json:"bannerFile"`
}
