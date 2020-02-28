package iagesdk

// GetMakersListResponse describe success response for
// GetMakersListRequest API request
type GetMakersListResponse struct {
	TotalPages int     `json:"totalPages"`
	TotalCount int     `json:"totalCount"`
	Data       []Maker `json:"data"`
}
