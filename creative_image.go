package iagesdk

type CreativeImage struct {
	Creative

	LineItemId int `json:"lineItemId"`

	// Creative title
	Title string `json:"title"` //max 40

	// Additional text description
	Description string `json:"description"` // max 90

	// Final landing url where ClickURL follows
	Domain     string `json:"domain"` //max 40
	LastDomain string `json:"lastDomain"`

	// Text on button
	Button string `json:"button"` //max 15

	// Tracking/Target link
	ClickURL string `json:"clickUrl"`

	// An array of strings representing impression tracking links.
	// Each element of the array must be a valid URL.
	ThirdPartyURLs []string `json:"thirdPartyUrls"`

	NativeImageFile ImageFile `json:"nativeImageFile"`
}
