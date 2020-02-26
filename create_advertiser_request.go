package iagesdk

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// CreateAdvertiserRequest describe API request to iAGE platform to create new advertiser
type CreateAdvertiserRequest struct {
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

// NewCreateAdvertiserRequest initialize CreateAdvertiserRequest based on Advertiser instance and password
func NewCreateAdvertiserRequest(adv *Advertiser, password string) *CreateAdvertiserRequest {
	return &CreateAdvertiserRequest{
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
func (car *CreateAdvertiserRequest) URL() string {
	//TODO:
	return `/v1/advertisers/`
}

// Method return API request http method
func (car *CreateAdvertiserRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (car *CreateAdvertiserRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(car)
	if err != nil {
		//TODO: log
		log.Print(err)
		return nil
	}

	return body
}
