package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateBannerByURLRequest describe API request to iAGE platform to create new banner by URL
type CreateBannerByURLRequest struct {
	LineItemID     int      `json:"-"`
	TemplateType   int      `json:"templateType"`
	BannerType     int      `json:"bannerType"`
	Name           string   `json:"name"`
	Disabled       bool     `json:"disabled"`
	ClickURL       string   `json:"clickUrl"`
	BannerURL      string   `json:"bannerUrl"`
	API            []int    `json:"api"`
	ThirdPartyURLs []string `json:"thirdPartyUrls"`
	CustomWidth    int      `json:"customWidth,omitempty"`  // required only for TemplateType = 221!
	CustomHeight   int      `json:"customHeight,omitempty"` // required only for TemplateType = 221!
}

// NewCreateBannerByURLRequest initialize CreateBannerByURLRequest
func NewCreateBannerByURLRequest(lineItemID int, templateType int, name string, disabled bool, clickURL string, bannerURL string, thirdPartyUrls []string, width int, height int) *CreateBannerByURLRequest {
	return &CreateBannerByURLRequest{
		LineItemID:     lineItemID,
		TemplateType:   templateType,
		BannerType:     BannerHTMLByURL,
		Name:           name,
		Disabled:       disabled,
		ClickURL:       clickURL,
		BannerURL:      bannerURL,
		API:            []int{MRAID2},
		ThirdPartyURLs: thirdPartyUrls,
		CustomWidth:    width,
		CustomHeight:   height,
	}
}

// URL return API request entrypoint (URI)
func (cbhcr *CreateBannerByURLRequest) URL() string {
	return fmt.Sprintf(`/v1/lineitems/%v/creatives`, cbhcr.LineItemID)
}

// Method return API request http method
func (cbhcr *CreateBannerByURLRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (cbhcr *CreateBannerByURLRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(cbhcr)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
