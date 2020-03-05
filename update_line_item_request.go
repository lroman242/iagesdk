package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// UpdateLineItemRequest describe API request to iAGE platform to update advertiser
type UpdateLineItemRequest struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`

	AdPositions []int `json:"adPositions"`
	Devices     []int `json:"devices"`

	DomainsAndPages []string `json:"domainsAndPages"`
	AppBundles      []string `json:"appBundles"`
	IPAddresses     []string `json:"ipAddresses"`

	Countries  []int  `json:"countries"`
	Regions    []int  `json:"regions"`
	Cities     []int  `json:"cities"`
	GeoAreas   []Area `json:"geoAreas"`
	Segments   []int  `json:"segments"`
	Browsers   []int  `json:"browsers"`
	Categories []int  `json:"categories"`
	Carriers   []int  `json:"carriers"`
	Languages  []int  `json:"languages"`
	SSPs       []int  `json:"ssps"`

	Name string `json:"name"`

	ID                     int                 `json:"-"`
	CampaignID             int                 `json:"campaignId"`
	Status                 int                 `json:"status"`
	CPM                    float64             `json:"cpm"`
	FrequencyCap           int                 `json:"frequencyCap"`
	FrequencyPeriod        int                 `json:"frequencyPeriod"`
	DailyBudget            float64             `json:"dailyBudget"`
	SegmentsMatchCondition int                 `json:"segmentsMatchCondition"`
	Makers                 map[string][]string `json:"makers"`
	OS                     map[string][]string `json:"os"`
	Schedule               ScheduleWeek        `json:"schedule"`

	IsDailyBudgetEven         bool `json:"isDailyBudgetEven"`
	IsAdPositionsExcluded     bool `json:"isAdPositionsExcluded"`
	IsDomainsAndPagesExcluded bool `json:"isDomainsAndPagesExcluded"`
	IsAppBundlesExcluded      bool `json:"isAppBundlesExcluded"`
	IsIPAddressesExcluded     bool `json:"isIpAddressesExcluded"`
	IsGeoExcluded             bool `json:"isGeoExcluded"`
	IsGeoAreasExcluded        bool `json:"isGeoAreasExcluded"`
	IsMakersExcluded          bool `json:"isMakersExcluded"`
	IsBrowsersExcluded        bool `json:"isBrowsersExcluded"`
	IsCategoriesExcluded      bool `json:"isCategoriesExcluded"`
	IsCarriersExcluded        bool `json:"isCarriersExcluded"`
	IsOsExcluded              bool `json:"isOsExcluded"`
	IsLanguagesExcluded       bool `json:"isLanguagesExcluded"`
}

// NewUpdateLineItemRequest initialize UpdateLineItemRequest based on Campaign instance
//TODO:
//func NewUpdateLineItemRequest() *UpdateLineItemRequest {
//	return &UpdateLineItemRequest{
//
//	}
//}

// URL return API request entrypoint (URI)
func (ulir *UpdateLineItemRequest) URL() string {
	return fmt.Sprintf(`/v1/lineitems/%v`, ulir.ID)
}

// Method return API request http method
func (ulir *UpdateLineItemRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (ulir *UpdateLineItemRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(ulir)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
