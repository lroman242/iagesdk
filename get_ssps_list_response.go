package iagesdk

// GetSSPsListResponse describe success response for
// GetSSPsListRequest API request
type GetSSPsListResponse struct {
	TotalPages int   `json:"totalPages"`
	TotalCount int   `json:"totalCount"`
	Data       []SSP `json:"data"`
}
