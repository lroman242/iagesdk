package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// TopUpAdvertiserAccountRequest describe API request to iAGE platform to top up advertiser`s balance
type TopUpAdvertiserAccountRequest struct {
	ID     int     `json:"-"`
	Amount float64 `json:"amount"`
}

// NewTopUpAdvertiserAccountRequest initialize TopUpAdvertiserAccountRequest based on advertiser ID
func NewTopUpAdvertiserAccountRequest(ID int, amount float64) *TopUpAdvertiserAccountRequest {
	return &TopUpAdvertiserAccountRequest{
		ID:     ID,
		Amount: amount,
	}
}

// URL return API request entrypoint (URI)
func (tuaar *TopUpAdvertiserAccountRequest) URL() string {
	return fmt.Sprintf("/v1/advertisers/%v/topupaccount", tuaar.ID)
}

// Method return API request http method
func (tuaar *TopUpAdvertiserAccountRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (tuaar *TopUpAdvertiserAccountRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(tuaar)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
