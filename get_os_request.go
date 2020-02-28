package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetOSRequest describe API request to iAGE platform to get OS by ID
type GetOSRequest struct {
	ID int
}

// NewGetOSRequest initialize GetOSRequest based on OS ID
func NewGetOSRequest(ID int) *GetOSRequest {
	return &GetOSRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gor *GetOSRequest) URL() string {
	return fmt.Sprintf(`/v1/os/%v`, gor.ID)
}

// Method return API request http method
func (gor *GetOSRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gor *GetOSRequest) Body() io.Reader {
	return nil
}
