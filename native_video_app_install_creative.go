package iagesdk

type NativeVideoAppInstallCreative struct {
	Creative

	ClickUrl           string              `json:"clickUrl"`
	LastDomain         string              `json:"lastDomain"`
	ThirdPartyUrls     []string            `json:"thirdPartyUrls"`
	SkipOffset         int                 `json:"skipoffset	"`
	Protocols          []int               `json:"protocols"`
	VastTrackingEvents map[string][]string `json:"vastTrackingEvents"`
	Title              string              `json:"title"`
	Description        string              `json:"description"`
	Button             string              `json:"button"`
	AppUrl             string              `json:"appUrl"`
	AppStore           string              `json:"appStore"`
	Brand              string              `json:"brand"`
	Rating             float64             `json:"rating"`
	Likes              int                 `json:"likes"`
	Downloads          int                 `json:"downloads"`
	VideoFirstFile     VideoFile           `json:"videoFirstFile"`
	VideoSecondFile    VideoFile           `json:"videoSecondFile"`
	NativeIconFile     ImageFile           `json:"nativeIconFile"`
	NativeImageFile    ImageFile           `json:"nativeImageFile"`
}
