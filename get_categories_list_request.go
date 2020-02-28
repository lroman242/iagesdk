package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetCategoriesListRequest describe API request to iAGE platform to fetch list of categories
type GetCategoriesListRequest struct {
	Page int `json:"-"`
	Size int `json:"-"`
}

// NewGetCategoriesListRequest initialize GetCategoriesListRequest
// with available query params
func NewGetCategoriesListRequest(page int, size int) *GetCategoriesListRequest {
	return &GetCategoriesListRequest{
		Page: page,
		Size: size,
	}
}

// URL return API request entrypoint (URI)
func (gclr *GetCategoriesListRequest) URL() string {
	uri := `/v1/categories/`

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
func (gclr *GetCategoriesListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gclr *GetCategoriesListRequest) Body() io.Reader {
	return nil
}
