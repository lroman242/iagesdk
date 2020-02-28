package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetLanguageRequest describe API request to iAGE platform to get Language by ID
type GetLanguageRequest struct {
	ID int
}

// NewGetLanguageRequest initialize GetLanguageRequest based on Language ID
func NewGetLanguageRequest(ID int) *GetLanguageRequest {
	return &GetLanguageRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gor *GetLanguageRequest) URL() string {
	return fmt.Sprintf(`/v1/languages/%v`, gor.ID)
}

// Method return API request http method
func (gor *GetLanguageRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gor *GetLanguageRequest) Body() io.Reader {
	return nil
}
