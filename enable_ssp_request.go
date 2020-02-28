package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// EnableSSPRequest describe API request to iAGE platform to enable SSP by ID
type EnableSSPRequest struct {
	ID int //required
}

// NewEnableSSPRequest initialize EnableSSPRequest based on SSP ID
func NewEnableSSPRequest(ID int) *EnableSSPRequest {
	return &EnableSSPRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (esr *EnableSSPRequest) URL() string {
	return fmt.Sprintf("/v1/ssps/%v/enable", esr.ID)
}

// Method return API request http method
func (esr *EnableSSPRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (esr *EnableSSPRequest) Body() io.Reader {
	return nil
}
