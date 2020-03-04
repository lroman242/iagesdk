package iagesdk

import "net/url"

type BannerHTMLByCodeCreative struct {
	Creative
	BannerType     int       `json:"bannerType"`
	Name           string    `json:"name"`
	ClickUrl       string    `json:"clickUrl"`
	LastDomain     string    `json:"lastDomain"`
	Content        string    `json:"content"`
	API            []int     `json:"api"`
	ThirdPartyUrls []string  `json:"thirdPartyUrls"`
	BannerFile     ImageFile `json:"bannerFile"`
}
