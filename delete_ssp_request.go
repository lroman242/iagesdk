package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// DeleteSSPRequest describe API request to iAGE platform to delete SSP by ID
type DeleteSSPRequest struct {
	ID int //required
}

// NewDeleteSSPRequest initialize DeleteSSPRequest based on SSP ID
func NewDeleteSSPRequest(ID int) *DeleteSSPRequest {
	return &DeleteSSPRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (dsr *DeleteSSPRequest) URL() string {
	return fmt.Sprintf("/v1/ssps/%v", dsr.ID)
}

// Method return API request http method
func (dsr *DeleteSSPRequest) Method() string {
	return http.MethodDelete
}

// Body generate body content of API request
func (dsr *DeleteSSPRequest) Body() io.Reader {
	return nil
}
