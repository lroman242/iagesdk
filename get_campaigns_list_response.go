package iagesdk

// GetCampaignsListResponse describe success response for
// GetCampaignsListRequest API request
type GetCampaignsListResponse struct {
	TotalPages int        `json:"totalPages"`
	TotalCount int        `json:"totalCount"`
	Data       []Campaign `json:"data"`
}
