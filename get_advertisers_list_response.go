package iagesdk

// GetAdvertisersListResponse describe success response for
// GetAdvertisersListRequest API request
type GetAdvertisersListResponse struct {
	TotalPages int          `json:"totalPages"`
	TotalCount int          `json:"totalCount"`
	Data       []Advertiser `json:"data"`
}
