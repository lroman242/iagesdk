package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// UpdateNativeVideoAppInstallRequest describe API request to iAGE platform to update new native video app install
type UpdateNativeVideoAppInstallRequest struct {
	ID             int      `json:"-"`
	TemplateType   int      `json:"templateType"`
	Disabled       bool     `json:"disabled"`
	ClickURL       string   `json:"clickUrl"`
	ThirdPartyURLs []string `json:"thirdPartyUrls"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Button         string   `json:"button"`
	AppURL         string   `json:"appUrl"`
	AppStore       string   `json:"appStore"`
	Brand          string   `json:"brand"`
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

// NewUpdateNativeVideoAppInstallRequest initialize UpdateNativeVideoAppInstallRequest
func NewUpdateNativeVideoAppInstallRequest(id int, templateType int, title string, brand string, description string, appURL string, appStore string, rating float64, likes int, downloads int, button string, disabled bool, clickURL string, iconURL string, imageURL string, videoFirstURL string, videoSecondURL string, videoStartDelay []int, thirdPartyUrls []string, events map[string][]string) *UpdateNativeVideoAppInstallRequest {
	request := &UpdateNativeVideoAppInstallRequest{
		ID:               id,
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
func (req *UpdateNativeVideoAppInstallRequest) URL() string {
	return fmt.Sprintf(`/v1/creatives/%v`, req.ID)
}

// Method return API request http method
func (req *UpdateNativeVideoAppInstallRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (req *UpdateNativeVideoAppInstallRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
