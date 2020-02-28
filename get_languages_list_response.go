package iagesdk

// GetLanguagesListResponse describe success response for
// GetLanguagesListRequest API request
type GetLanguagesListResponse struct {
	TotalPages int        `json:"totalPages"`
	TotalCount int        `json:"totalCount"`
	Data       []Language `json:"data"`
}
