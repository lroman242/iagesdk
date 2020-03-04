package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// DeleteCreativeRequest describe API request to iAGE platform to delete Creative by ID
type DeleteCreativeRequest struct {
	ID int //required
}

// NewDeleteCreativeRequest initialize DeleteCreativeRequest based on Creative ID
func NewDeleteCreativeRequest(ID int) *DeleteCreativeRequest {
	return &DeleteCreativeRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (dsr *DeleteCreativeRequest) URL() string {
	return fmt.Sprintf("/v1/creatives/%v", dsr.ID)
}

// Method return API request http method
func (dsr *DeleteCreativeRequest) Method() string {
	return http.MethodDelete
}

// Body generate body content of API request
func (dsr *DeleteCreativeRequest) Body() io.Reader {
	return nil
}
