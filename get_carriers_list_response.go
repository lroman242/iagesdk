package iagesdk

// GetCarriersListResponse describe success response for
// GetCarriersListRequest API request
type GetCarriersListResponse struct {
	TotalPages int        `json:"totalPages"`
	TotalCount int        `json:"totalCount"`
	Data       []Carrier  `json:"data"`
}
