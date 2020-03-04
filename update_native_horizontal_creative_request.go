package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// UpdateNativeHorizontalCreativeRequest describe API request to iAGE platform to update new native horizontal creative
type UpdateNativeHorizontalCreativeRequest struct {
	ID             int      `json:"-"`
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

// NewUpdateNativeHorizontalCreativeRequest initialize UpdateNativeHorizontalCreativeRequest
func NewUpdateNativeHorizontalCreativeRequest(id int, templateType int, title string, description string, button string, disabled bool, clickURL string, imageURL string, thirdPartyUrls []string) *UpdateNativeHorizontalCreativeRequest {
	request := &UpdateNativeHorizontalCreativeRequest{
		ID:             id,
		TemplateType:   templateType,
		Disabled:       disabled,
		ClickURL:       clickURL,
		Title:          title,
		Description:    description,
		Button:         button,
		ImageURL:       imageURL,
		ThirdPartyURLs: thirdPartyUrls,
	}

	return request
}

// URL return API request entrypoint (URI)
func (req *UpdateNativeHorizontalCreativeRequest) URL() string {
	return fmt.Sprintf(`/v1/creatives/%v`, req.ID)
}

// Method return API request http method
func (req *UpdateNativeHorizontalCreativeRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (req *UpdateNativeHorizontalCreativeRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
