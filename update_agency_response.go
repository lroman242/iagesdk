package iagesdk

// UpdateAgencyResponse describe success response for
// UpdateAgencyRequest API request
type UpdateAgencyResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`       //max 255
	Status     int    `json:"status"`     // 1 or 2
	Email      string `json:"email"`      //max 255
	WebSiteURL string `json:"webSiteUrl"` //max 1024
	Country    string `json:"country"`    //max 255
}
