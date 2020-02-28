package iagesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

// UpdateLineItemRequest describe API request to iAGE platform to update advertiser
type UpdateLineItemRequest struct {
	ID                        int                 `json:"-"`
	CampaignID                int                 `json:"campaignId"`
	Name                      string              `json:"name"`
	StartDate                 time.Time           `json:"startDate"`
	EndDate                   time.Time           `json:"endDate"`
	Status                    int                 `json:"status"`
	CPM                       float64             `json:"cpm"`
	FrequencyCap              int                 `json:"frequencyCap"`
	FrequencyPeriod           int                 `json:"frequencyPeriod"`
	DailyBudget               float64             `json:"dailyBudget"`
	IsDailyBudgetEven         bool                `json:"isDailyBudgetEven"`
	IsAdPositionsExcluded     bool                `json:"isAdPositionsExcluded"`
	AdPositions               []int               `json:"adPositions"`
	Devices                   []int               `json:"devices"`
	IsDomainsAndPagesExcluded bool                `json:"isDomainsAndPagesExcluded"`
	DomainsAndPages           []string            `json:"domainsAndPages"`
	IsAppBundlesExcluded      bool                `json:"isAppBundlesExcluded"`
	AppBundles                []string            `json:"appBundles"`
	IsIpAddressesExcluded     bool                `json:"isIpAddressesExcluded"`
	IpAddresses               []net.IP            `json:"ipAddresses"`
	IsGeoExcluded             bool                `json:"isGeoExcluded"`
	Countries                 []int               `json:"countries"`
	Regions                   []int               `json:"regions"`
	Cities                    []int               `json:"cities"`
	IsGeoAreasExcluded        bool                `json:"isGeoAreasExcluded"`
	GeoAreas                  []Area              `json:"geoAreas"`
	Segments                  []int               `json:"segments"`
	SegmentsMatchCondition    int                 `json:"segmentsMatchCondition"`
	IsMakersExcluded          bool                `json:"isMakersExcluded"`
	Makers                    map[string][]string `json:"makers"`
	IsBrowsersExcluded        bool                `json:"isBrowsersExcluded"`
	Browsers                  []int               `json:"browsers"`
	IsCategoriesExcluded      bool                `json:"isCategoriesExcluded"`
	Categories                []int               `json:"categories"`
	IsCarriersExcluded        bool                `json:"isCarriersExcluded"`
	Carriers                  []int               `json:"carriers"`
	IsOsExcluded              bool                `json:"isOsExcluded"`
	OS                        map[string][]string `json:"os"`
	IsLanguagesExcluded       bool                `json:"isLanguagesExcluded"`
	Languages                 []int               `json:"languages"`
	SSPs                      []int               `json:"ssps"`
	Schedule                  ScheduleWeek        `json:"schedule"`
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
