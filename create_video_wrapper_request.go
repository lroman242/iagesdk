package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateVideoWrapperRequest describe API request to iAGE platform to create new video wrapper
type CreateVideoWrapperRequest struct {
	LineItemID              int      `json:"-"`
	TemplateType            int      `json:"templateType"`
	BannerType              int      `json:"bannerType"`
	Name                    string   `json:"name"`
	Disabled                bool     `json:"disabled"`
	BannerURL               string   `json:"bannerUrl"`
	ThirdPartyURLs          []string `json:"thirdPartyUrls"`
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

// NewCreateVideoWrapperRequest initialize CreateVideoWrapperRequest
func NewCreateVideoWrapperRequest(lineItemID int, templateType int, name string, disabled bool, bannerURL string, thirdPartyUrls []string, events map[string][]string) *CreateVideoWrapperRequest {
	request := &CreateVideoWrapperRequest{
		LineItemID:     lineItemID,
		TemplateType:   templateType,
		BannerType:     BannerHTMLByURL,
		Name:           name,
		Disabled:       disabled,
		BannerURL:      bannerURL,
		ThirdPartyURLs: thirdPartyUrls,
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
func (cvwr *CreateVideoWrapperRequest) URL() string {
	return fmt.Sprintf(`/v1/lineitems/%v/creatives`, cvwr.LineItemID)
}

// Method return API request http method
func (cvwr *CreateVideoWrapperRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (cvwr *CreateVideoWrapperRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(cvwr)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
