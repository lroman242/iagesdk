package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetCarrierRequest describe API request to iAGE platform to get Carrier by ID
type GetCarrierRequest struct {
	ID int
}

// NewGetCarrierRequest initialize GetCarrierRequest based on Carrier ID
func NewGetCarrierRequest(ID int) *GetCarrierRequest {
	return &GetCarrierRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gar *GetCarrierRequest) URL() string {
	return fmt.Sprintf(`/v1/carriers/%v`, gar.ID)
}

// Method return API request http method
func (gar *GetCarrierRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gar *GetCarrierRequest) Body() io.Reader {
	return nil
}
