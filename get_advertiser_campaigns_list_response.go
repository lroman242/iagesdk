package iagesdk

// GetAdvertiserCampaignsListResponse describe success response for
// GetAdvertiserCampaignsListRequest API request
type GetAdvertiserCampaignsListResponse struct {
	TotalPages int        `json:"totalPages"`
	TotalCount int        `json:"totalCount"`
	Data       Advertiser `json:"data"`
}
