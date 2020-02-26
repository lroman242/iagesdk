package iagesdk

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// CreateAgencyRequest describe API request to iAGE platform to create new agency
type CreateAgencyRequest struct {
	Name       string `json:"name"`       //required, max 255
	Status     int    `json:"status"`     //required, 1 or 2
	Email      string `json:"email"`      //required, max 255
	WebSiteURL string `json:"webSiteUrl"` //required, max 1024
	Country    string `json:"country"`    //required, max 255
}

// NewCreateAgencyRequest initialize CreateAgencyRequest based on Agency instance
func NewCreateAgencyRequest(agency *Agency) *CreateAgencyRequest {
	return &CreateAgencyRequest{
		Name:       agency.Name,
		Status:     agency.Status,
		Email:      agency.Email,
		WebSiteURL: agency.WebSiteURL,
		Country:    agency.Country,
	}
}

// URL return API request entrypoint (URI)
func (car *CreateAgencyRequest) URL() string {
	return v1Agencies
}

// Method return API request http method
func (car *CreateAgencyRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (car *CreateAgencyRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(car)
	if err != nil {
		//TODO: log
		log.Print(err)
		return nil
	}

	return body
}
