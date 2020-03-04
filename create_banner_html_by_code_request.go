package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateBannerHTMLByCodeRequest describe API request to iAGE platform to create new banner by HTML
type CreateBannerHTMLByCodeRequest struct {
	LineItemID     int      `json:"-"`
	TemplateType   int      `json:"templateType"`
	BannerType     int      `json:"bannerType"`
	Name           string   `json:"name"`
	Disabled       bool     `json:"disabled"`
	ClickURL       string   `json:"clickUrl"`
	Content        string   `json:"content"`
	API            []int    `json:"api"`
	ThirdPartyURLs []string `json:"thirdPartyUrls"`
	CustomWidth    int      `json:"customWidth,omitempty"`
	CustomHeight   int      `json:"customHeight,omitempty"`
}

// NewCreateBannerHTMLByCodeRequest initialize CreateBannerHTMLByCodeRequest
func NewCreateBannerHTMLByCodeRequest(lineItemID int, templateType int, name string, disabled bool, clickURL string, content string, thirdPartyUrls []string, width int, height int) *CreateBannerHTMLByCodeRequest {
	return &CreateBannerHTMLByCodeRequest{
		LineItemID:     lineItemID,
		TemplateType:   templateType,
		BannerType:     BannerHTMLByCode,
		Name:           name,
		Disabled:       disabled,
		ClickURL:       clickURL,
		Content:        content,
		API:            []int{MRAID2},
		ThirdPartyURLs: thirdPartyUrls,
		CustomWidth:    width,
		CustomHeight:   height,
	}
}

// URL return API request entrypoint (URI)
func (cbhcr *CreateBannerHTMLByCodeRequest) URL() string {
	return fmt.Sprintf(`/v1/lineitems/%v/creatives`, cbhcr.LineItemID)
}

// Method return API request http method
func (cbhcr *CreateBannerHTMLByCodeRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (cbhcr *CreateBannerHTMLByCodeRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(cbhcr)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
