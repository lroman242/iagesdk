package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetAdvertiserRequest describe API request to iAGE platform to get advertiser by ID
type GetAdvertiserRequest struct {
	ID int
}

// NewGetAdvertiserRequest initialize GetAdvertiserRequest based on Advertiser ID
func NewGetAdvertiserRequest(ID int) *GetAdvertiserRequest {
	return &GetAdvertiserRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gar *GetAdvertiserRequest) URL() string {
	return fmt.Sprintf("/v1/advertisers/%v", gar.ID)
}

// Method return API request http method
func (gar *GetAdvertiserRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gar *GetAdvertiserRequest) Body() io.Reader {
	return nil
}
