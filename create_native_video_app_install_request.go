package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateNativeVideoAppInstallRequest describe API request to iAGE platform to create new native video app install
type CreateNativeVideoAppInstallRequest struct {
	LineItemID     int      `json:"-"`
	TemplateType   int      `json:"templateType"`
	Disabled       bool     `json:"disabled"`
	ClickURL       string   `json:"clickUrl"`
	ThirdPartyURLs []string `json:"thirdPartyUrls"`
	Title          string   `json:"title"`
	Brand          string   `json:"brand"`
	Description    string   `json:"description"`
	Button         string   `json:"button"`
	AppURL         string   `json:"appUrl"`
	AppStore       string   `json:"appStore"`
	Rating         float64  `json:"rating"`
	Likes          int      `json:"likes"`
	Downloads      int      `json:"downloads"`

	VideoFirstURL  string `json:"videoFirstUrl"`
	VideoSecondURL string `json:"videoSecondUrl"`
	IconURL        string `json:"iconUrl"`
	ImageURL       string `json:"imageUrl"`

	VideoStartDelays        []int    `json:"videoStartDelays"`
	VastEventStart          []string `json:"vastEventStart"`
	VastEventSkip           []string `json:"vastEventSkip"`
	VastEventMidpoint       []string `json:"vastEventMidpoint"`
	VastEventFirstQuartile  []string `json:"vastEventFirstQuartile"`
	VastEventThirdQuartile  []string `json:"vastEventThirdQuartile"`
	VastEventComplete       []string `json:"vastEventComplete"`
	VastEventMute           []string `json:"vastEventMute"`
	VastEventUnmute         []string `json:"vastEventUnmute"`
	VastEventPause          []string `json:"vastEventPause"`
	VastEventResume         []string `json:"vastEventResume"`
	VastEventFullscreen     []string `json:"vastEventFullscreen"`
	VastEventExitFullscreen []string `json:"vastEventExitFullscreen"`
	VastEventExpand         []string `json:"vastEventExpand"`
	VastEventCollapse       []string `json:"vastEventCollapse"`
	VastEventError          []string `json:"vastEventError"`
	VastEventClose          []string `json:"vastEventClose"`
	VastEventCreativeView   []string `json:"vastEventCreativeView"`
}

// NewCreateNativeVideoAppInstallRequest initialize CreateNativeVideoAppInstallRequest
func NewCreateNativeVideoAppInstallRequest(lineItemID int, templateType int, title string, brand string, description string, appURL string, appStore string, rating float64, likes int, downloads int, button string, disabled bool, clickURL string, iconURL string, imageURL string, videoFirstURL string, videoSecondURL string, videoStartDelay []int, thirdPartyUrls []string, events map[string][]string) *CreateNativeVideoAppInstallRequest {
	request := &CreateNativeVideoAppInstallRequest{
		LineItemID:       lineItemID,
		TemplateType:     templateType,
		Disabled:         disabled,
		ClickURL:         clickURL,
		Title:            title,
		Brand:            brand,
		Description:      description,
		Button:           button,
		AppURL:           appURL,
		AppStore:         appStore,
		Rating:           rating,
		Likes:            likes,
		Downloads:        downloads,
		IconURL:          iconURL,
		ImageURL:         imageURL,
		VideoFirstURL:    videoFirstURL,
		VideoSecondURL:   videoSecondURL,
		VideoStartDelays: videoStartDelay,
		ThirdPartyURLs:   thirdPartyUrls,
	}

	if value, ok := events["start"]; ok {
		request.VastEventStart = value
	}
	if value, ok := events["skip"]; ok {
		request.VastEventSkip = value
	}
	if value, ok := events["midpoint"]; ok {
		request.VastEventMidpoint = value
	}
	if value, ok := events["firstQuartile"]; ok {
		request.VastEventFirstQuartile = value
	}
	if value, ok := events["thirdQuartile"]; ok {
		request.VastEventThirdQuartile = value
	}
	if value, ok := events["complete"]; ok {
		request.VastEventComplete = value
	}
	if value, ok := events["mute"]; ok {
		request.VastEventMute = value
	}
	if value, ok := events["unmute"]; ok {
		request.VastEventUnmute = value
	}
	if value, ok := events["pause"]; ok {
		request.VastEventPause = value
	}
	if value, ok := events["resume"]; ok {
		request.VastEventResume = value
	}
	if value, ok := events["fullscreen"]; ok {
		request.VastEventFullscreen = value
	}
	if value, ok := events["exitFullscreen"]; ok {
		request.VastEventExitFullscreen = value
	}
	if value, ok := events["expand"]; ok {
		request.VastEventExpand = value
	}
	if value, ok := events["collapse"]; ok {
		request.VastEventCollapse = value
	}
	if value, ok := events["error"]; ok {
		request.VastEventError = value
	}
	if value, ok := events["close"]; ok {
		request.VastEventClose = value
	}
	if value, ok := events["creativeView"]; ok {
		request.VastEventCreativeView = value
	}

	return request
}

// URL return API request entrypoint (URI)
func (req *CreateNativeVideoAppInstallRequest) URL() string {
	return fmt.Sprintf(`/v1/lineitems/%v/creatives`, req.LineItemID)
}

// Method return API request http method
func (req *CreateNativeVideoAppInstallRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (req *CreateNativeVideoAppInstallRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
