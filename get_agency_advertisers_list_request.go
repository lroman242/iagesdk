package iagesdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetAgencyAdvertisersListRequest describe API request to iAGE platform to fetch list of advertisers related to the Agency
type GetAgencyAdvertisersListRequest struct {
	AgencyID int `json:"-"`
	Page     int `json:"-"`
	Size     int `json:"-"`
}

// NewGetAgencyAdvertisersListRequest initialize GetAgencyAdvertisersListRequest
// with available query params
func NewGetAgencyAdvertisersListRequest(page int, size int, agencyID int) *GetAgencyAdvertisersListRequest {
	return &GetAgencyAdvertisersListRequest{
		AgencyID: agencyID,
		Page:     page,
		Size:     size,
	}
}

// URL return API request entrypoint (URI)
func (r *GetAgencyAdvertisersListRequest) URL() string {
	uri := fmt.Sprintf(`/v1/agencies/%v/advertisers/`, r.AgencyID)

	u := url.URL{}
	if r.Page > 0 {
		u.Query().Add("page", strconv.Itoa(r.Page))
	}
	if r.Size > 0 {
		u.Query().Add("size", strconv.Itoa(r.Size))
	}

	if r.Page > 0 || r.Size > 0 {
		return fmt.Sprintf("%s?%s", uri, u.Query().Encode())
	}

	return uri
}

// Method return API request http method
func (r *GetAgencyAdvertisersListRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (r *GetAgencyAdvertisersListRequest) Body() io.Reader {
	return nil
}
