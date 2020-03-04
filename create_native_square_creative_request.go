package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateNativeSquareCreativeRequest describe API request to iAGE platform to create new native square
type CreateNativeSquareCreativeRequest struct {
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

// NewCreateNativeSquareCreativeRequest initialize CreateNativeSquareCreativeRequest
func NewCreateNativeSquareCreativeRequest(lineItemID int, templateType int, title string, description string, domain string, button string, disabled bool, clickURL string, imageURL string, thirdPartyUrls []string) *CreateNativeSquareCreativeRequest {
	request := &CreateNativeSquareCreativeRequest{
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
func (req *CreateNativeSquareCreativeRequest) URL() string {
	return fmt.Sprintf(`/v1/lineitems/%v/creatives`, req.LineItemID)
}

// Method return API request http method
func (req *CreateNativeSquareCreativeRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (req *CreateNativeSquareCreativeRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
