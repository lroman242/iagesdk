package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// UpdateAgencyRequest describe API request to iAGE platform to update agency
type UpdateAgencyRequest struct {
	ID         int    `json:"-"`          //required
	Name       string `json:"name"`       //max 255
	Status     int    `json:"status"`     // 1 or 2
	Email      string `json:"email"`      //max 255
	WebSiteURL string `json:"webSiteUrl"` //max 1024
	Country    string `json:"country"`    //max 255
}

// NewUpdateAgencyRequest initialize UpdateAgencyRequest based on Agency instance
func NewUpdateAgencyRequest(agency *Agency) *UpdateAgencyRequest {
	return &UpdateAgencyRequest{
		ID:         agency.ID,
		Name:       agency.Name,
		Status:     agency.Status,
		Email:      agency.Email,
		WebSiteURL: agency.WebSiteURL,
		Country:    agency.Country,
	}
}

// URL return API request entrypoint (URI)
func (uar *UpdateAgencyRequest) URL() string {
	return fmt.Sprintf("/v1/agencies/%v", uar.ID)
}

// Method return API request http method
func (uar *UpdateAgencyRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (uar *UpdateAgencyRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(uar)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
