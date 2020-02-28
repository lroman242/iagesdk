package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetMakerRequest describe API request to iAGE platform to get Maker by ID
type GetMakerRequest struct {
	ID int
}

// NewGetMakerRequest initialize GetMakerRequest based on Maker ID
func NewGetMakerRequest(ID int) *GetMakerRequest {
	return &GetMakerRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gor *GetMakerRequest) URL() string {
	return fmt.Sprintf(`/v1/makers/%v`, gor.ID)
}

// Method return API request http method
func (gor *GetMakerRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gor *GetMakerRequest) Body() io.Reader {
	return nil
}
