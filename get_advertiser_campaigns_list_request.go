package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetAdvertiserCampaignsListRequest describe API request to iAGE platform to fetch list of Campaigns related to Advertiser
type GetAdvertiserCampaignsListRequest struct {
	AdvertiserID int `json:"-"`
	Page         int `json:"-"`
	Size         int `json:"-"`
}

// NewGetAdvertiserCampaignsListRequest initialize GetAdvertiserCampaignsListRequest
// with available query params
func NewGetAdvertiserCampaignsListRequest(advertiserID, page, size int) *GetAdvertiserCampaignsListRequest {
	return &GetAdvertiserCampaignsListRequest{
		AdvertiserID: advertiserID,
		Page:         page,
		Size:         size,
	}
}

// URL return API request entrypoint (URI)
func (gaclr *GetAdvertiserCampaignsListRequest) URL() string {
	uri := fmt.Sprintf(`/v1/advertisers/%v/campaigns/`, gaclr.AdvertiserID)

	u := url.URL{}
	if gaclr.Page > 0 {
		u.Query().Add("page", strconv.Itoa(gaclr.Page))
	}
	if gaclr.Size > 0 {
		u.Query().Add("size", strconv.Itoa(gaclr.Size))
	}

	if gaclr.Page > 0 || gaclr.Size > 0 {
		return fmt.Sprintf("%s?%s", uri, u.Query().Encode())
	}

	return uri
}

// Method return API request http method
func (gaclr *GetAdvertiserCampaignsListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gaclr *GetAdvertiserCampaignsListRequest) Body() io.Reader {
	return nil
}
