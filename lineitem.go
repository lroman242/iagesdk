package iagesdk

import (
	"time"
)

const FrequencyPeriodDay = 1
const FrequencyPeriodWeek = 2
const FrequencyPeriodMonth = 3

// LineItem describe iAGE`s line item data structure
type LineItem struct { //
	ID         int       `json:"id"`
	CampaignID int       `json:"campaignId"`
	Name       string    `json:"name"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	Status     int       `json:"status"`
	// Cost per thousand impressions
	CPM               float64 `json:"cpm"`
	FrequencyCap      int     `json:"frequencyCap"`
	FrequencyPeriod   int     `json:"frequencyPeriod"`
	DailyBudget       float64 `json:"dailyBudget"`
	IsDailyBudgetEven bool    `json:"isDailyBudgetEven"`

	IsAdPositionsExcluded bool `json:"isAdPositionsExcluded"`
	// Available values:
	// - 1 Top half of screen
	// - 3 Bottom half of screen
	// - 4 Head
	// - 5 Footer
	// - 6 Sidebar
	// - 7 All screen
	AdPositions []int `json:"adPositions"`

	// Available values:
	// - 1 Mobile
	// - 2 Tablet
	// - 3 Desktop
	Devices []int `json:"devices"`

	IsDomainsAndPagesExcluded bool     `json:"isDomainsAndPagesExcluded"`
	DomainsAndPages           []string `json:"domainsAndPages"`

	IsAppBundlesExcluded bool     `json:"isAppBundlesExcluded"`
	AppBundles           []string `json:"appBundles"`

	IsIpAddressesExcluded bool `json:"isIpAddressesExcluded"`
	// Available formats:
	// - 192.168.0.1 single IP
	// - 192.168.0.0 - 192.168.210.19 range of IPs
	// - 192.168 short form of range (equal to 192.168.0.0 - 192.168.255.255)
	IpAddresses []string `json:"ipAddresses"`

	IsGeoExcluded bool      `json:"isGeoExcluded"`
	Countries     []Country `json:"countries"`
	Regions       []Region  `json:"regions"`
	Cities        []City    `json:"cities"`

	IsGeoAreasExcluded bool   `json:"isGeoAreasExcluded"`
	GeoAreas           []Area `json:"geoAreas"`

	// Available values:
	// - 1 "AND" condition. All advertising segments in which a particular user is located must be targeted for this advertising campaign.
	// - 2 "OR" condition. At least one advertising segment in which a particular user is located must be targeted for this advertising campaign.
	// - 3 "NOT" condition. No advertising segment in which a particular user is located should be targeted for this advertising campaign.
	Segments               []int `json:"segments"`
	SegmentsMatchCondition int   `json:"segmentsMatchCondition"`

	IsMakersExcluded bool    `json:"isMakersExcluded"`
	Makers           []Maker `json:"makers"`

	IsBrowsersExcluded bool      `json:"isBrowsersExcluded"`
	Browsers           []Browser `json:"browsers"`

	IsCategoriesExcluded bool       `json:"isCategoriesExcluded"`
	Categories           []Category `json:"categories"`

	IsCarriersExcluded bool      `json:"isCarriersExcluded"`
	Carriers           []Carrier `json:"carriers"`

	IsOsExcluded bool `json:"isOsExcluded"`
	OS           []OS `json:"os"`

	IsLanguagesExcluded bool       `json:"isLanguagesExcluded"`
	Languages           []Language `json:"languages"`

	//SSPs []int

	Schedule ScheduleWeek `json:"schedule"`
}

type ScheduleWeek struct {
	Week           [24][7]bool `json:"week"`
	ShowInHolidays bool        `json:"showInHolidays"`
	UseLocalTime   bool        `json:"useLocalTime"`
}
