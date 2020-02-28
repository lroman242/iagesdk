package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetMakersListRequest describe API request to iAGE platform to fetch list of Makers
type GetMakersListRequest struct {
	Page int `json:"-"`
	Size int `json:"-"`
}

// NewGetMakersListRequest initialize GetMakersListRequest
// with available query params
func NewGetMakersListRequest(page int, size int) *GetMakersListRequest {
	return &GetMakersListRequest{
		Page: page,
		Size: size,
	}
}

// URL return API request entrypoint (URI)
func (golr *GetMakersListRequest) URL() string {
	uri := `/v1/makers/`

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
func (golr *GetMakersListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (golr *GetMakersListRequest) Body() io.Reader {
	return nil
}
