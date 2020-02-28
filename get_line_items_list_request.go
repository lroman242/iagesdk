package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetLineItemsListRequest describe API request to iAGE platform to fetch list of line items
type GetLineItemsListRequest struct {
	Page int `json:"-"`
	Size int `json:"-"`
}

// NewGetLineItemsListRequest initialize GetLineItemsListRequest
// with available query params
func NewGetLineItemsListRequest(page int, size int) *GetLineItemsListRequest {
	return &GetLineItemsListRequest{
		Page: page,
		Size: size,
	}
}

// URL return API request entrypoint (URI)
func (gclr *GetLineItemsListRequest) URL() string {
	uri := `/v1/lineitems/`

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
func (gclr *GetLineItemsListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gclr *GetLineItemsListRequest) Body() io.Reader {
	return nil
}
