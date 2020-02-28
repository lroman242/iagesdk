package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetLineItemRequest describe API request to iAGE platform to get line item by ID
type GetLineItemRequest struct {
	ID int
}

// NewGetLineItemRequest initialize GetLineItemRequest based on Agency ID
func NewGetLineItemRequest(ID int) *GetLineItemRequest {
	return &GetLineItemRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gar *GetLineItemRequest) URL() string {
	return fmt.Sprintf(`/v1/lineitems/%v`, gar.ID)
}

// Method return API request http method
func (gar *GetLineItemRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gar *GetLineItemRequest) Body() io.Reader {
	return nil
}
