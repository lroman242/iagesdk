package iagesdk

// GetAgencyAdvertisersListResponse describe success response for
// GetAgencyAdvertisersListRequest API request
type GetAgencyAdvertisersListResponse struct {
	TotalPages int          `json:"totalPages"`
	TotalCount int          `json:"totalCount"`
	Data       []Advertiser `json:"data"`
}
