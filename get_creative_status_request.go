package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetCreativeStatusRequest describe API request to iAGE platform to get creative status by ID
type GetCreativeStatusRequest struct {
	ID int
}

// NewGetCreativeStatusRequest initialize GetCreativeStatusRequest based on Creative ID
func NewGetCreativeStatusRequest(ID int) *GetCreativeStatusRequest {
	return &GetCreativeStatusRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gcsr *GetCreativeStatusRequest) URL() string {
	return fmt.Sprintf(`/v1/creatives/%v/status`, gcsr.ID)
}

// Method return API request http method
func (gcsr *GetCreativeStatusRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gcsr *GetCreativeStatusRequest) Body() io.Reader {
	return nil
}
