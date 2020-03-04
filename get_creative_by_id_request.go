package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetCreativeByIDRequest describe API request to iAGE platform to get creative ID
type GetCreativeByIDRequest struct {
	ID int
}

// NewGetCreativeByIDRequest initialize GetCreativeByIDRequest based on Browser ID
func NewGetCreativeByIDRequest(ID int) *GetCreativeByIDRequest {
	return &GetCreativeByIDRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gcr *GetCreativeByIDRequest) URL() string {
	return fmt.Sprintf(`/v1/creatives/%v`, gcr.ID)
}

// Method return API request http method
func (gcr *GetCreativeByIDRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gcr *GetCreativeByIDRequest) Body() io.Reader {
	return nil
}
