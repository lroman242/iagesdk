package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// EnableCreativeRequest describe API request to iAGE platform to enable creative by ID
type EnableCreativeRequest struct {
	ID int
}

// NewEnableCreativeRequest initialize EnableCreativeRequest based on Creative ID
func NewEnableCreativeRequest(ID int) *EnableCreativeRequest {
	return &EnableCreativeRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gcsr *EnableCreativeRequest) URL() string {
	return fmt.Sprintf(`/v1/creatives/%v/enable`, gcsr.ID)
}

// Method return API request http method
func (gcsr *EnableCreativeRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gcsr *EnableCreativeRequest) Body() io.Reader {
	return nil
}
