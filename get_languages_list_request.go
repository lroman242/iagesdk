package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetLanguagesListRequest describe API request to iAGE platform to fetch list of languages
type GetLanguagesListRequest struct {
	Page int `json:"-"`
	Size int `json:"-"`
}

// NewGetLanguagesListRequest initialize GetLanguagesListRequest
// with available query params
func NewGetLanguagesListRequest(page int, size int) *GetLanguagesListRequest {
	return &GetLanguagesListRequest{
		Page: page,
		Size: size,
	}
}

// URL return API request entrypoint (URI)
func (gllr *GetLanguagesListRequest) URL() string {
	uri := `/v1/languages/`

	u := url.URL{}
	if gllr.Page > 0 {
		u.Query().Add("page", strconv.Itoa(gllr.Page))
	}
	if gllr.Size > 0 {
		u.Query().Add("size", strconv.Itoa(gllr.Size))
	}

	if gllr.Page > 0 || gllr.Size > 0 {
		return fmt.Sprintf("%s?%s", uri, u.Query().Encode())
	}

	return uri
}

// Method return API request http method
func (gllr *GetLanguagesListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gllr *GetLanguagesListRequest) Body() io.Reader {
	return nil
}
