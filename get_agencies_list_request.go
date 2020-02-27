package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetAgenciesListRequest describe API request to iAGE platform to fetch list of agencies
type GetAgenciesListRequest struct {
	Page int `json:"-"`
	Size int `json:"-"`
}

// NewGetAgenciesListRequest initialize GetAgenciesListRequest
// with available query params
func NewGetAgenciesListRequest(page int, size int) *GetAgenciesListRequest {
	return &GetAgenciesListRequest{
		Page: page,
		Size: size,
	}
}

// URL return API request entrypoint (URI)
func (r *GetAgenciesListRequest) URL() string {
	uri := `/v1/agencies/`

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
func (r *GetAgenciesListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (r *GetAgenciesListRequest) Body() io.Reader {
	return nil
}
