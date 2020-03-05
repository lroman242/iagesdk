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

// CreateLineItemRequest describe API request to iAGE platform to create new line item
type CreateLineItemRequest struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`

	AdPositions     []int    `json:"adPositions"`
	Devices         []int    `json:"devices"`
	DomainsAndPages []string `json:"domainsAndPages"`
	IPAddresses     []string `json:"ipAddresses"`
	AppBundles      []string `json:"appBundles"`
	Countries       []int    `json:"countries"`
	Regions         []int    `json:"regions"`
	Cities          []int    `json:"cities"`
	GeoAreas        []Area   `json:"geoAreas"`
	Segments        []int    `json:"segments"`
	Browsers        []int    `json:"browsers"`
	Categories      []int    `json:"categories"`
	Carriers        []int    `json:"carriers"`
	Languages       []int    `json:"languages"`
	SSPs            []int    `json:"ssps"`

	Name                   string  `json:"name"`
	CampaignID             int     `json:"campaignId"`
	Status                 int     `json:"status"`
	CPM                    float64 `json:"cpm"`
	FrequencyCap           int     `json:"frequencyCap"`
	FrequencyPeriod        int     `json:"frequencyPeriod"`
	DailyBudget            float64 `json:"dailyBudget"`
	SegmentsMatchCondition int     `json:"segmentsMatchCondition"`

	OS     map[string][]string `json:"os"`
	Makers map[string][]string `json:"makers"`

	Schedule ScheduleWeek `json:"schedule"`

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

// NewCreateLineItemRequest initialize CreateLineItemRequest based on ...
//TODO:
//func NewCreateLineItemRequest() *CreateLineItemRequest {
//	return &CreateLineItemRequest{
//
//	}
//}

// URL return API request entrypoint (URI)
func (clir *CreateLineItemRequest) URL() string {
	return fmt.Sprintf(`/v1/campaigns/%v/lineitems`, clir.CampaignID)
}

// Method return API request http method
func (clir *CreateLineItemRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (clir *CreateLineItemRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(clir)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
