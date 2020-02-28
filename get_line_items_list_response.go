package iagesdk

// GetLineItemsListResponse describe success response for
// GetLineItemsListRequest API request
type GetLineItemsListResponse struct {
	TotalPages int        `json:"totalPages"`
	TotalCount int        `json:"totalCount"`
	Data       []LineItem `json:"data"` // short description. only id, campaignId, name, startDate, endDate, status, cpm, dailyBudget, dailyBudget
}
