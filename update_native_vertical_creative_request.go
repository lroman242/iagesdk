package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// UpdateNativeVerticalCreativeRequest describe API request to iAGE platform to update new native vertical creative
type UpdateNativeVerticalCreativeRequest struct {
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

// NewUpdateNativeVerticalCreativeRequest initialize UpdateNativeVerticalCreativeRequest
func NewUpdateNativeVerticalCreativeRequest(id int, templateType int, title string, description string, button string, disabled bool, clickURL string, imageURL string, thirdPartyUrls []string) *UpdateNativeVerticalCreativeRequest {
	request := &UpdateNativeVerticalCreativeRequest{
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
func (req *UpdateNativeVerticalCreativeRequest) URL() string {
	return fmt.Sprintf(`/v1/creatives/%v`, req.ID)
}

// Method return API request http method
func (req *UpdateNativeVerticalCreativeRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (req *UpdateNativeVerticalCreativeRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
