package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateNativeImageFeedRequest describe API request to iAGE platform to create new native image feed
type CreateNativeImageFeedRequest struct {
	LineItemID     int      `json:"-"`
	TemplateType   int      `json:"templateType"`
	Disabled       bool     `json:"disabled"`
	ClickURL       string   `json:"clickUrl"`
	ThirdPartyURLs []string `json:"thirdPartyUrls"`
	Title          string   `json:"title"`
	Brand          string   `json:"brand"`
	Description    string   `json:"description"`
	Domain         string   `json:"domain"`
	Button         string   `json:"button"`
	IconURL        string   `json:"iconUrl"`
	ImageURL       string   `json:"imageUrl"`
}

// NewCreateNativeImageFeedRequest initialize CreateNativeImageFeedRequest
func NewCreateNativeImageFeedRequest(lineItemID int, templateType int, title string, brand string, description string, domain string, button string, disabled bool, clickURL string, iconURL string, imageURL string, thirdPartyUrls []string) *CreateNativeImageFeedRequest {
	request := &CreateNativeImageFeedRequest{
		LineItemID:     lineItemID,
		TemplateType:   templateType,
		Disabled:       disabled,
		ClickURL:       clickURL,
		Title:          title,
		Brand:          brand,
		Description:    description,
		Domain:         domain,
		Button:         button,
		IconURL:        iconURL,
		ImageURL:       imageURL,
		ThirdPartyURLs: thirdPartyUrls,
	}

	return request
}

// URL return API request entrypoint (URI)
func (req *CreateNativeImageFeedRequest) URL() string {
	return fmt.Sprintf(`/v1/lineitems/%v/creatives`, req.LineItemID)
}

// Method return API request http method
func (req *CreateNativeImageFeedRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (req *CreateNativeImageFeedRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
