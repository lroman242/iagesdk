package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// UpdateAdvertiserRequest describe API request to iAGE platform to update advertiser
type UpdateAdvertiserRequest struct {
	ID        int    `json:"-"`
	FirstName string `json:"firstName"`          //required, max 255
	LastName  string `json:"lastName"`           //required, max 255
	Status    int    `json:"status"`             //required, 1 or 2
	Email     string `json:"email"`              //required, max 255
	Password  string `json:"password"`           //required, min 6, max 255
	Country   string `json:"country"`            //required, max 255
	City      string `json:"city,omitempty"`     //max 255
	Phone     string `json:"phone"`              //required, max 255
	Skype     string `json:"skype,omitempty"`    //max 255
	Currency  string `json:"currency,omitempty"` //available values: USD, GBP, EUR, RUB (default USD)
}

// NewUpdateAdvertiserRequest initialize UpdateAdvertiserRequest based on Advertiser instance and password
func NewUpdateAdvertiserRequest(adv *Advertiser, password string) *UpdateAdvertiserRequest {
	return &UpdateAdvertiserRequest{
		ID:        adv.ID,
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
	}
}

// URL return API request entrypoint (URI)
func (uar *UpdateAdvertiserRequest) URL() string {
	return fmt.Sprintf(`/v1/advertisers/%v`, uar.ID)
}

// Method return API request http method
func (uar *UpdateAdvertiserRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (uar *UpdateAdvertiserRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(uar)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
