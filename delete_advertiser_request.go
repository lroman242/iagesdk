package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// DeleteAdvertiserRequest describe API request to iAGE platform to delete advertiser by ID
type DeleteAdvertiserRequest struct {
	ID int //required
}

// NewDeleteAdvertiserRequest initialize DeleteAdvertiserRequest based on advertiser ID
func NewDeleteAdvertiserRequest(ID int) *DeleteAdvertiserRequest {
	return &DeleteAdvertiserRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (dar *DeleteAdvertiserRequest) URL() string {
	return fmt.Sprintf("/v1/advertisers/%v", dar.ID)
}

// Method return API request http method
func (dar *DeleteAdvertiserRequest) Method() string {
	return http.MethodDelete
}

// Body generate body content of API request
func (dar *DeleteAdvertiserRequest) Body() io.Reader {
	return nil
}
