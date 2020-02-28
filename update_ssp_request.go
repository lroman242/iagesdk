package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// UpdateSSPRequest describe API request to iAGE platform to update SSP
type UpdateSSPRequest struct {
	ID                   int    `json:"-"`                    //required
	Name                 string `json:"name"`                 //required, max 255
	Disabled             bool   `json:"disabled"`             //required
	ClickMacrosUnescaped string `json:"clickMacrosUnescaped"` //required, max 255
	ClickMacrosEscaped   string `json:"clickMacrosEscaped"`   //required, max 255
}

// NewUpdateSSPRequest initialize UpdateSSPRequest based on Agency instance
func NewUpdateSSPRequest(ssp *SSP) *UpdateSSPRequest {
	return &UpdateSSPRequest{
		ID:                   ssp.ID,
		Name:                 ssp.Name,
		Disabled:             ssp.Disabled,
		ClickMacrosUnescaped: ssp.ClickMacrosUnescaped,
		ClickMacrosEscaped:   ssp.ClickMacrosEscaped,
	}
}

// URL return API request entrypoint (URI)
func (usr *UpdateSSPRequest) URL() string {
	return fmt.Sprintf("/v1/ssps/%v", usr.ID)
}

// Method return API request http method
func (usr *UpdateSSPRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (usr *UpdateSSPRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(usr)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
