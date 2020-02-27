package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetCampaignRequest describe API request to iAGE platform to get Campaign by ID
type GetCampaignRequest struct {
	ID int
}

// NewGetCampaignRequest initialize GetCampaignRequest based on Campaign ID
func NewGetCampaignRequest(ID int) *GetCampaignRequest {
	return &GetCampaignRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gar *GetCampaignRequest) URL() string {
	return fmt.Sprintf(`/v1/campaigns/%v`, gar.ID)
}

// Method return API request http method
func (gar *GetCampaignRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gar *GetCampaignRequest) Body() io.Reader {
	return nil
}
