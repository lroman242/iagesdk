package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// UpdateNativeImageAppInstallRequest describe API request to iAGE platform to update new native video app install
type UpdateNativeImageAppInstallRequest struct {
	ID             int      `json:"-"`
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

// NewUpdateNativeImageAppInstallRequest initialize UpdateNativeImageAppInstallRequest
func NewUpdateNativeImageAppInstallRequest(id int, templateType int, title string, description string, appURL string, appStore string, rating float64, likes int, downloads int, button string, disabled bool, clickURL string, iconURL string, imageURL string, thirdPartyUrls []string) *UpdateNativeImageAppInstallRequest {
	request := &UpdateNativeImageAppInstallRequest{
		ID:             id,
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
func (req *UpdateNativeImageAppInstallRequest) URL() string {
	return fmt.Sprintf(`/v1/creatives/%v`, req.ID)
}

// Method return API request http method
func (req *UpdateNativeImageAppInstallRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (req *UpdateNativeImageAppInstallRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
