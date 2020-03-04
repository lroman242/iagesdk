package iagesdk

type NativeImageAppInstallCreative struct {
	Creative

	ClickUrl        string    `json:"clickUrl"`
	LastDomain      string    `json:"lastDomain"`
	ThirdPartyUrls  []string  `json:"thirdPartyUrls"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Button          string    `json:"button"`
	AppUrl          string    `json:"appUrl"`
	AppStore        string    `json:"appStore"`
	Rating          float64   `json:"rating"`
	Likes           int       `json:"likes"`
	Downloads       int       `json:"downloads"`
	NativeImageFile ImageFile `json:"nativeImageFile"`
	NativeIconFile  ImageFile `json:"nativeIconFile"`
}
