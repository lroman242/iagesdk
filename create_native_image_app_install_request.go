package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateNativeImageAppInstallRequest describe API request to iAGE platform to create new native image app install
type CreateNativeImageAppInstallRequest struct {
	LineItemID     int      `json:"-"`
	TemplateType   int      `json:"templateType"`
	Disabled       bool     `json:"disabled"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Rating         float64  `json:"rating"`
	Likes          int      `json:"likes"`
	Downloads      int      `json:"downloads"`
	Button         string   `json:"button"`
	ClickURL       string   `json:"clickUrl"`
	AppURL         string   `json:"appUrl"`
	AppStore       string   `json:"appStore"`
	IconURL        string   `json:"iconUrl"`
	ImageURL       string   `json:"imageUrl"`
	ThirdPartyURLs []string `json:"thirdPartyUrls"`
}

// NewCreateNativeImageAppInstallRequest initialize CreateNativeImageAppInstallRequest
func NewCreateNativeImageAppInstallRequest(lineItemID int, templateType int, title string, description string, appURL string, appStore string, rating float64, likes int, downloads int, button string, disabled bool, clickURL string, iconURL string, imageURL string, thirdPartyUrls []string) *CreateNativeImageAppInstallRequest {
	request := &CreateNativeImageAppInstallRequest{
		LineItemID:     lineItemID,
		TemplateType:   templateType,
		Disabled:       disabled,
		ClickURL:       clickURL,
		Title:          title,
		Description:    description,
		Button:         button,
		AppURL:         appURL,
		AppStore:       appStore,
		Rating:         rating,
		Likes:          likes,
		Downloads:      downloads,
		IconURL:        iconURL,
		ImageURL:       imageURL,
		ThirdPartyURLs: thirdPartyUrls,
	}

	return request
}

// URL return API request entrypoint (URI)
func (req *CreateNativeImageAppInstallRequest) URL() string {
	return fmt.Sprintf(`/v1/lineitems/%v/creatives`, req.LineItemID)
}

// Method return API request http method
func (req *CreateNativeImageAppInstallRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (req *CreateNativeImageAppInstallRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
