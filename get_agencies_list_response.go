package iagesdk

// GetAgenciesListResponse describe success response for
// GetAgenciesListRequest API request
type GetAgenciesListResponse struct {
	TotalPages int      `json:"totalPages"`
	TotalCount int      `json:"totalCount"`
	Data       []Agency `json:"data"`
}
