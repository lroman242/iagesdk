package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// DeleteAgencyRequest describe API request to iAGE platform to delete agency by ID
type DeleteAgencyRequest struct {
	ID int //required
}

// NewDeleteAgencyRequest initialize DeleteAgencyRequest based on Agency ID
func NewDeleteAgencyRequest(ID int) *DeleteAgencyRequest {
	return &DeleteAgencyRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (dar *DeleteAgencyRequest) URL() string {
	return fmt.Sprintf("%s%v", v1Agencies, dar.ID)
}

// Method return API request http method
func (dar *DeleteAgencyRequest) Method() string {
	return http.MethodDelete
}

// Body generate body content of API request
func (dar *DeleteAgencyRequest) Body() io.Reader {
	return nil
}
