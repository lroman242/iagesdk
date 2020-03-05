package iagesdk

// NativeVideoAppInstallCreative describe iAGE`s video app install creative data structure
type NativeVideoAppInstallCreative struct {
	Creative

	ClickURL           string              `json:"clickUrl"`
	LastDomain         string              `json:"lastDomain"`
	ThirdPartyUrls     []string            `json:"thirdPartyUrls"`
	SkipOffset         int                 `json:"skipoffset	"`
	Protocols          []int               `json:"protocols"`
	VastTrackingEvents map[string][]string `json:"vastTrackingEvents"`
	Title              string              `json:"title"`
	Description        string              `json:"description"`
	Button             string              `json:"button"`
	AppURL             string              `json:"appUrl"`
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
