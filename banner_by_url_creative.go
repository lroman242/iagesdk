package iagesdk

// BannerByURLCreative describe iAGE`s banner created from URL
type BannerByURLCreative struct {
	Creative

	BannerType     int       `json:"bannerType"`
	Name           string    `json:"name"`
	ClickURL       string    `json:"clickUrl"`
	LastDomain     string    `json:"lastDomain"`
	ThirdPartyUrls []string  `json:"thirdPartyUrls"`
	BannerFile     ImageFile `json:"bannerFile"`
}
