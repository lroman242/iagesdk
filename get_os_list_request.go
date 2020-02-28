package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetOSListRequest describe API request to iAGE platform to fetch list of OS
type GetOSListRequest struct {
	Page int `json:"-"`
	Size int `json:"-"`
}

// NewGetOSListRequest initialize GetOSListRequest
// with available query params
func NewGetOSListRequest(page int, size int) *GetOSListRequest {
	return &GetOSListRequest{
		Page: page,
		Size: size,
	}
}

// URL return API request entrypoint (URI)
func (golr *GetOSListRequest) URL() string {
	uri := `/v1/os/`

	u := url.URL{}
	if golr.Page > 0 {
		u.Query().Add("page", strconv.Itoa(golr.Page))
	}
	if golr.Size > 0 {
		u.Query().Add("size", strconv.Itoa(golr.Size))
	}

	if golr.Page > 0 || golr.Size > 0 {
		return fmt.Sprintf("%s?%s", uri, u.Query().Encode())
	}

	return uri
}

// Method return API request http method
func (golr *GetOSListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (golr *GetOSListRequest) Body() io.Reader {
	return nil
}
