package iagesdk

// Carrier describe iAGE`s mobile operator data structure
type Carrier struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
}
