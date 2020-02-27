package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// DeleteCampaignRequest describe API request to iAGE platform to delete campaign by ID
type DeleteCampaignRequest struct {
	ID int //required
}

// NewDeleteCampaignRequest initialize DeleteCampaignRequest based on Agency ID
func NewDeleteCampaignRequest(ID int) *DeleteCampaignRequest {
	return &DeleteCampaignRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (dcr *DeleteCampaignRequest) URL() string {
	return fmt.Sprintf("/v1/campaigns/%v", dcr.ID)
}

// Method return API request http method
func (dcr *DeleteCampaignRequest) Method() string {
	return http.MethodDelete
}

// Body generate body content of API request
func (dcr *DeleteCampaignRequest) Body() io.Reader {
	return nil
}
