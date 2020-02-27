package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetAgencyRequest describe API request to iAGE platform to get agency by ID
type GetAgencyRequest struct {
	ID int
}

// NewGetAgencyRequest initialize GetAgencyRequest based on Agency ID
func NewGetAgencyRequest(ID int) *GetAgencyRequest {
	return &GetAgencyRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gar *GetAgencyRequest) URL() string {
	return fmt.Sprintf(`/v1/agencies/%v`, gar.ID)
}

// Method return API request http method
func (gar *GetAgencyRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gar *GetAgencyRequest) Body() io.Reader {
	return nil
}
