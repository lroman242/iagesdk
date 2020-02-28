package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetCampaignLineItemsListRequest describe API request to iAGE platform to fetch list of line items related to Campaign
type GetCampaignLineItemsListRequest struct {
	CampaignID int `json:"-"`
	Page       int `json:"-"`
	Size       int `json:"-"`
}

// NewGetCampaignLineItemsListRequest initialize GetCampaignLineItemsListRequest
// with available query params
func NewGetCampaignLineItemsListRequest(campaignID, page, size int) *GetCampaignLineItemsListRequest {
	return &GetCampaignLineItemsListRequest{
		CampaignID: campaignID,
		Page:       page,
		Size:       size,
	}
}

// URL return API request entrypoint (URI)
func (gclilr *GetCampaignLineItemsListRequest) URL() string {
	uri := `/v1/campaigns/%v/lineitems/`

	u := url.URL{}
	if gclilr.Page > 0 {
		u.Query().Add("page", strconv.Itoa(gclilr.Page))
	}
	if gclilr.Size > 0 {
		u.Query().Add("size", strconv.Itoa(gclilr.Size))
	}

	if gclilr.Page > 0 || gclilr.Size > 0 {
		return fmt.Sprintf("%s?%s", uri, u.Query().Encode())
	}

	return uri
}

// Method return API request http method
func (gclilr *GetCampaignLineItemsListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gclilr *GetCampaignLineItemsListRequest) Body() io.Reader {
	return nil
}
