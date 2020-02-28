package iagesdk

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// CreateSSPRequest describe API request to iAGE platform to create new SSP
type CreateSSPRequest struct {
	Name                 string `json:"name"`                 //required, max 255
	Disabled             bool   `json:"disabled"`             //required, true - stopped, false - active, default: active
	ClickMacrosUnescaped string `json:"clickMacrosUnescaped"` //required, max 255
	ClickMacrosEscaped   string `json:"clickMacrosEscaped"`   //required, max 255
}

// NewCreateSSPRequest initialize CreateSSPRequest based on SSP instance
func NewCreateSSPRequest(ssp SSP) *CreateSSPRequest {
	return &CreateSSPRequest{
		Name:                 ssp.Name,
		Disabled:             ssp.Disabled,
		ClickMacrosEscaped:   ssp.ClickMacrosEscaped,
		ClickMacrosUnescaped: ssp.ClickMacrosUnescaped,
	}
}

// URL return API request entrypoint (URI)
func (car *CreateSSPRequest) URL() string {
	return `/v1/ssps/`
}

// Method return API request http method
func (car *CreateSSPRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (car *CreateSSPRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(car)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
