package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetSSPRequest describe API request to iAGE platform to get SSP by ID
type GetSSPRequest struct {
	ID int
}

// NewGetSSPRequest initialize GetSSPRequest based on SSP ID
func NewGetSSPRequest(ID int) *GetSSPRequest {
	return &GetSSPRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gsr *GetSSPRequest) URL() string {
	return fmt.Sprintf(`/v1/ssps/%v`, gsr.ID)
}

// Method return API request http method
func (gsr *GetSSPRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gsr *GetSSPRequest) Body() io.Reader {
	return nil
}
