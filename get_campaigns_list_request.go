package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetCampaignsListRequest describe API request to iAGE platform to fetch list of campaigns
type GetCampaignsListRequest struct {
	Page int `json:"-"`
	Size int `json:"-"`
}

// NewGetCampaignsListRequest initialize GetCampaignsListRequest
// with available query params
func NewGetCampaignsListRequest(page int, size int) *GetCampaignsListRequest {
	return &GetCampaignsListRequest{
		Page: page,
		Size: size,
	}
}

// URL return API request entrypoint (URI)
func (gclr *GetCampaignsListRequest) URL() string {
	uri := `/v1/campaigns/`

	u := url.URL{}
	if gclr.Page > 0 {
		u.Query().Add("page", strconv.Itoa(gclr.Page))
	}
	if gclr.Size > 0 {
		u.Query().Add("size", strconv.Itoa(gclr.Size))
	}

	if gclr.Page > 0 || gclr.Size > 0 {
		return fmt.Sprintf("%s?%s", uri, u.Query().Encode())
	}

	return uri
}

// Method return API request http method
func (gclr *GetCampaignsListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gclr *GetCampaignsListRequest) Body() io.Reader {
	return nil
}
