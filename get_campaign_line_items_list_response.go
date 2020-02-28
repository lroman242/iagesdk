package iagesdk

// GetCampaignLineItemsListResponse describe success response for
// GetCampaignLineItemsListRequest API request
type GetCampaignLineItemsListResponse struct {
	TotalPages int `json:"totalPages"`
	TotalCount int `json:"totalCount"`
	Data       struct {
		ID        int        `json:"id"`
		Name      string     `json:"name"`
		Campaigns []LineItem `json:"campaigns"`
	}
	//Data       []LineItem `json:"data"` // short description. only id, campaignId, name, startDate, endDate, status, cpm, dailyBudget, dailyBudget
}
