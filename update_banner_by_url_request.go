package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// UpdateBannerByURLRequest describe API request to iAGE platform to update banner by URL
type UpdateBannerByURLRequest struct {
	ID             int      `json:"-"`
	TemplateType   int      `json:"templateType"`
	BannerType     int      `json:"bannerType"`
	Name           string   `json:"name"`
	Disabled       bool     `json:"disabled"`
	ClickURL       string   `json:"clickUrl"`
	BannerURL      string   `json:"bannerUrl"`
	API            []int    `json:"api"`
	ThirdPartyURLs []string `json:"thirdPartyUrls"`
	CustomWidth    int      `json:"customWidth,omitempty"`
	CustomHeight   int      `json:"customHeight,omitempty"`
}

// NewUpdateBannerByURLRequest initialize UpdateBannerByURLRequest based on Advertiser instance and password
func NewUpdateBannerByURLRequest(id int, templateType int, name string, disabled bool, clickURL string, bannerURL string, thirdPartyURLs []string, width int, height int) *UpdateBannerByURLRequest {
	return &UpdateBannerByURLRequest{
		ID:             id,
		TemplateType:   templateType,
		BannerType:     BannerHTMLByCode,
		Name:           name,
		Disabled:       disabled,
		ClickURL:       clickURL,
		BannerURL:      bannerURL,
		API:            []int{MRAID2},
		ThirdPartyURLs: thirdPartyURLs,
		CustomWidth:    width,
		CustomHeight:   height,
	}
}

// URL return API request entrypoint (URI)
func (ubur *UpdateBannerByURLRequest) URL() string {
	return fmt.Sprintf(`/v1/creatives/%v`, ubur.ID)
}

// Method return API request http method
func (ubur *UpdateBannerByURLRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (ubur *UpdateBannerByURLRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(ubur)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
