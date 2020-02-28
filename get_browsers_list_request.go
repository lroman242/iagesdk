package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetBrowsersListRequest describe API request to iAGE platform to fetch list of browsers
type GetBrowsersListRequest struct {
	Page int `json:"-"`
	Size int `json:"-"`
}

// NewGetBrowsersListRequest initialize GetBrowsersListRequest
// with available query params
func NewGetBrowsersListRequest(page int, size int) *GetBrowsersListRequest {
	return &GetBrowsersListRequest{
		Page: page,
		Size: size,
	}
}

// URL return API request entrypoint (URI)
func (gblr *GetBrowsersListRequest) URL() string {
	uri := `/v1/browsers/`

	u := url.URL{}
	if gblr.Page > 0 {
		u.Query().Add("page", strconv.Itoa(gblr.Page))
	}
	if gblr.Size > 0 {
		u.Query().Add("size", strconv.Itoa(gblr.Size))
	}

	if gblr.Page > 0 || gblr.Size > 0 {
		return fmt.Sprintf("%s?%s", uri, u.Query().Encode())
	}

	return uri
}

// Method return API request http method
func (gblr *GetBrowsersListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gblr *GetBrowsersListRequest) Body() io.Reader {
	return nil
}
