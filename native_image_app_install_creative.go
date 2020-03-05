package iagesdk

// NativeImageAppInstallCreative describe iAGE`s image app install creative data structure
type NativeImageAppInstallCreative struct {
	Creative

	ClickURL        string    `json:"clickUrl"`
	LastDomain      string    `json:"lastDomain"`
	ThirdPartyUrls  []string  `json:"thirdPartyUrls"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Button          string    `json:"button"`
	AppURL          string    `json:"appUrl"`
	AppStore        string    `json:"appStore"`
	Rating          float64   `json:"rating"`
	Likes           int       `json:"likes"`
	Downloads       int       `json:"downloads"`
	NativeImageFile ImageFile `json:"nativeImageFile"`
	NativeIconFile  ImageFile `json:"nativeIconFile"`
}
