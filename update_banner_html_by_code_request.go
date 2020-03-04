package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// UpdateBannerHTMLByCodeRequest describe API request to iAGE platform to update banner HTML by code
type UpdateBannerHTMLByCodeRequest struct {
	ID             int      `json:"-"`
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

// NewUpdateBannerHTMLByCodeRequest initialize UpdateBannerHTMLByCodeRequest based on Advertiser instance and password
func NewUpdateBannerHTMLByCodeRequest(id int, templateType int, name string, disabled bool, clickURL string, content string, thirdPartyURLs []string, width int, height int) *UpdateBannerHTMLByCodeRequest {
	return &UpdateBannerHTMLByCodeRequest{
		ID:             id,
		TemplateType:   templateType,
		BannerType:     BannerHTMLByCode,
		Name:           name,
		Disabled:       disabled,
		ClickURL:       clickURL,
		Content:        content,
		API:            []int{MRAID2},
		ThirdPartyURLs: thirdPartyURLs,
		CustomWidth:    width,
		CustomHeight:   height,
	}
}

// URL return API request entrypoint (URI)
func (ubhcr *UpdateBannerHTMLByCodeRequest) URL() string {
	return fmt.Sprintf(`/v1/creatives/%v`, ubhcr.ID)
}

// Method return API request http method
func (ubhcr *UpdateBannerHTMLByCodeRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (ubhcr *UpdateBannerHTMLByCodeRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(ubhcr)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
