package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateNativeVerticalCreativeRequest describe API request to iAGE platform to create new native vertical
type CreateNativeVerticalCreativeRequest struct {
	LineItemID     int      `json:"-"`
	TemplateType   int      `json:"templateType"`
	Disabled       bool     `json:"disabled"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Domain         string   `json:"domain"`
	Button         string   `json:"button"`
	ClickURL       string   `json:"clickUrl"`
	ImageURL       string   `json:"imageUrl"`
	ThirdPartyURLs []string `json:"thirdPartyUrls"`
}

// NewCreateNativeVerticalCreativeRequest initialize CreateNativeVerticalCreativeRequest
func NewCreateNativeVerticalCreativeRequest(lineItemID int, templateType int, title string, description string, domain string, button string, disabled bool, clickURL string, imageURL string, thirdPartyUrls []string) *CreateNativeVerticalCreativeRequest {
	request := &CreateNativeVerticalCreativeRequest{
		LineItemID:     lineItemID,
		TemplateType:   templateType,
		Disabled:       disabled,
		ClickURL:       clickURL,
		Title:          title,
		Description:    description,
		Domain:         domain,
		Button:         button,
		ImageURL:       imageURL,
		ThirdPartyURLs: thirdPartyUrls,
	}

	return request
}

// URL return API request entrypoint (URI)
func (req *CreateNativeVerticalCreativeRequest) URL() string {
	return fmt.Sprintf(`/v1/lineitems/%v/creatives`, req.LineItemID)
}

// Method return API request http method
func (req *CreateNativeVerticalCreativeRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (req *CreateNativeVerticalCreativeRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
