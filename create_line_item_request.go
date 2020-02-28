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

// CreateLineItemRequest describe API request to iAGE platform to create new line item
type CreateLineItemRequest struct {
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
