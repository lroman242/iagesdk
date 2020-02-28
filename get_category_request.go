package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetCategoryRequest describe API request to iAGE platform to get Category by ID
type GetCategoryRequest struct {
	ID int
}

// NewGetCategoryRequest initialize GetCategoryRequest based on Category ID
func NewGetCategoryRequest(ID int) *GetCategoryRequest {
	return &GetCategoryRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gar *GetCategoryRequest) URL() string {
	return fmt.Sprintf(`/v1/category/%v`, gar.ID)
}

// Method return API request http method
func (gar *GetCategoryRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gar *GetCategoryRequest) Body() io.Reader {
	return nil
}
