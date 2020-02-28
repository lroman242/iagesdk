package iagesdk

// GetCategoriesListResponse describe success response for
// GetCategoriesListRequest API request
type GetCategoriesListResponse struct {
	TotalPages int        `json:"totalPages"`
	TotalCount int        `json:"totalCount"`
	Data       []Category `json:"data"`
}
