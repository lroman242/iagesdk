package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetSSPsListRequest describe API request to iAGE platform to fetch list of SSPs
type GetSSPsListRequest struct {
	Page int `json:"-"`
	Size int `json:"-"`
}

// NewGetSSPsListRequest initialize GetSSPsListRequest
// with available query params
func NewGetSSPsListRequest(page int, size int) *GetSSPsListRequest {
	return &GetSSPsListRequest{
		Page: page,
		Size: size,
	}
}

// URL return API request entrypoint (URI)
func (glilr *GetSSPsListRequest) URL() string {
	uri := `/v1/ssps/`

	u := url.URL{}
	if glilr.Page > 0 {
		u.Query().Add("page", strconv.Itoa(glilr.Page))
	}
	if glilr.Size > 0 {
		u.Query().Add("size", strconv.Itoa(glilr.Size))
	}

	if glilr.Page > 0 || glilr.Size > 0 {
		return fmt.Sprintf("%s?%s", uri, u.Query().Encode())
	}

	return uri
}

// Method return API request http method
func (glilr *GetSSPsListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (glilr *GetSSPsListRequest) Body() io.Reader {
	return nil
}
