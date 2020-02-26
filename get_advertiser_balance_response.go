package iagesdk

// GetAdvertiserBalanceResponse describe success response for GetAdvertiserBalanceRequest
type GetAdvertiserBalanceResponse struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}
