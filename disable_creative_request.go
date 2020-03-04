package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// DisableCreativeRequest describe API request to iAGE platform to disable creative by ID
type DisableCreativeRequest struct {
	ID int //required
}

// NewDisableCreativeRequest initialize DisableCreativeRequest based on Creative ID
func NewDisableCreativeRequest(ID int) *DisableCreativeRequest {
	return &DisableCreativeRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (dsr *DisableCreativeRequest) URL() string {
	return fmt.Sprintf("/v1/creatives/%v/disable", dsr.ID)
}

// Method return API request http method
func (dsr *DisableCreativeRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (dsr *DisableCreativeRequest) Body() io.Reader {
	return nil
}
