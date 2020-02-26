package iagesdk

import (
	"fmt"
	"io"
	"net/http"
)

// GetAdvertiserBalanceRequest describe API request to iAGE platform to get advertiser`s balance
type GetAdvertiserBalanceRequest struct {
	ID int `json:"-"`
}

// NewGetAdvertiserBalanceRequest initialize GetAdvertiserBalanceRequest based on advertiser ID
func NewGetAdvertiserBalanceRequest(ID int) *GetAdvertiserBalanceRequest {
	return &GetAdvertiserBalanceRequest{
		ID: ID,
	}
}

// URL return API request entrypoint (URI)
func (gabr *GetAdvertiserBalanceRequest) URL() string {
	return fmt.Sprintf("/v1/advertisers/%v/balance", gabr.ID)
}

// Method return API request http method
func (gabr *GetAdvertiserBalanceRequest) Method() string {
	return http.MethodGet
}

// Body generate body content of API request
func (gabr *GetAdvertiserBalanceRequest) Body() io.Reader {
	return nil
}
