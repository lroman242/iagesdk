package iagesdk

// Advertiser describe iAGE`s advertiser data structure
type Advertiser struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    int    `json:"status"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Phone     string `json:"phone"`
	Skype     string `json:"skype"`
	Currency  string `json:"currency"`
}
