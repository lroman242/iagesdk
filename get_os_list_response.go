package iagesdk

// GetOSListResponse describe success response for
// GetOSListRequest API request
type GetOSListResponse struct {
	TotalPages int  `json:"totalPages"`
	TotalCount int  `json:"totalCount"`
	Data       []OS `json:"data"`
}
