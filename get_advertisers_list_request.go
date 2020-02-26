package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetAdvertisersListRequest describe API request to iAGE platform to fetch list of advertisers
type GetAdvertisersListRequest struct {
	Page int `json:"-"`
	Size int `json:"-"`
}

// NewGetAdvertisersListRequest initialize GetAdvertisersListRequest
// with available query params
func NewGetAdvertisersListRequest(page int, size int) *GetAdvertisersListRequest {
	return &GetAdvertisersListRequest{
		Page: page,
		Size: size,
	}
}

// URL return API request entrypoint (URI)
func (r *GetAdvertisersListRequest) URL() string {
	uri := `/v1/advertisers/`

	u := url.URL{}
	if r.Page > 0 {
		u.Query().Add("page", strconv.Itoa(r.Page))
	}
	if r.Size > 0 {
		u.Query().Add("size", strconv.Itoa(r.Size))
	}

	if r.Page > 0 || r.Size > 0 {
		return fmt.Sprintf("%s?%s", uri, u.Query().Encode())
	}

	return uri
}

// Method return API request http method
func (r *GetAdvertisersListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (r *GetAdvertisersListRequest) Body() io.Reader {
	return nil
}
