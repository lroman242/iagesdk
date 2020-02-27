package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateAgencyAdvertiserRequest describe API request to iAGE platform to create new advertiser for existed agency
type CreateAgencyAdvertiserRequest struct {
	CreateAdvertiserRequest
	AgencyID int `json:"-"` //required
}

// NewCreateAgencyAdvertiserRequest initialize CreateAgencyAdvertiserRequest based on Advertiser instance and password
func NewCreateAgencyAdvertiserRequest(agencyID int, adv *Advertiser, password string) *CreateAgencyAdvertiserRequest {
	return &CreateAgencyAdvertiserRequest{
		AgencyID: agencyID,
		CreateAdvertiserRequest: CreateAdvertiserRequest{
			FirstName: adv.FirstName,
			LastName:  adv.LastName,
			Status:    adv.Status,
			Email:     adv.Email,
			Password:  password,
			Country:   adv.Country,
			City:      adv.City,
			Phone:     adv.Phone,
			Skype:     adv.Skype,
			Currency:  adv.Currency,
		},
	}
}

// URL return API request entrypoint (URI)
func (caar *CreateAgencyAdvertiserRequest) URL() string {
	return fmt.Sprintf(`/v1/agencies/%v/advertisers/`, caar.AgencyID)
}

// Method return API request http method
func (caar *CreateAgencyAdvertiserRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (caar *CreateAgencyAdvertiserRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(caar)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
