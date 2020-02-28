package iagesdk

// GetBrowsersListResponse describe success response for
// GetBrowsersListRequest API request
type GetBrowsersListResponse struct {
	TotalPages int       `json:"totalPages"`
	TotalCount int       `json:"totalCount"`
	Data       []Browser `json:"data"`
}
