package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// DisableSSPRequest describe API request to iAGE platform to enable SSP by ID
type DisableSSPRequest struct {
	ID int //required
}

// NewDisableSSPRequest initialize DisableSSPRequest based on SSP ID
func NewDisableSSPRequest(ID int) *DisableSSPRequest {
	return &DisableSSPRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (dsr *DisableSSPRequest) URL() string {
	return fmt.Sprintf("/v1/ssps/%v/disable", dsr.ID)
}

// Method return API request http method
func (dsr *DisableSSPRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (dsr *DisableSSPRequest) Body() io.Reader {
	return nil
}
