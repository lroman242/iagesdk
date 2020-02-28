package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetCarriersListRequest describe API request to iAGE platform to fetch list of carriers
type GetCarriersListRequest struct {
	Page int `json:"-"`
	Size int `json:"-"`
}

// NewGetCarriersListRequest initialize GetCarriersListRequest
// with available query params
func NewGetCarriersListRequest(page int, size int) *GetCarriersListRequest {
	return &GetCarriersListRequest{
		Page: page,
		Size: size,
	}
}

// URL return API request entrypoint (URI)
func (gclr *GetCarriersListRequest) URL() string {
	uri := `/v1/carriers/`

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
func (gclr *GetCarriersListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gclr *GetCarriersListRequest) Body() io.Reader {
	return nil
}
