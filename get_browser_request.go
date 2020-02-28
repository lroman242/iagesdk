package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetBrowserRequest describe API request to iAGE platform to get Browser by ID
type GetBrowserRequest struct {
	ID int
}

// NewGetBrowserRequest initialize GetBrowserRequest based on Browser ID
func NewGetBrowserRequest(ID int) *GetBrowserRequest {
	return &GetBrowserRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gbr *GetBrowserRequest) URL() string {
	return fmt.Sprintf(`/v1/browsers/%v`, gbr.ID)
}

// Method return API request http method
func (gbr *GetBrowserRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gbr *GetBrowserRequest) Body() io.Reader {
	return nil
}
