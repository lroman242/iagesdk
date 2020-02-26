package iagesdk

// TopUpAdvertiserAccountResponse describe success response for TopUpAdvertiserAccountRequest
type TopUpAdvertiserAccountResponse struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}
