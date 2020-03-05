package iagesdk

import (
	"time"
)

//FrequencyPeriodDay - value used for LineItem.FrequencyPeriod to indicate daily frequency
const FrequencyPeriodDay = 1

//FrequencyPeriodWeek - value used for LineItem.FrequencyPeriod to indicate weekly frequency
const FrequencyPeriodWeek = 2

//FrequencyPeriodMonth - value used for LineItem.FrequencyPeriod to indicate monthly frequency
const FrequencyPeriodMonth = 3

// LineItem describe iAGE`s line item data structure
type LineItem struct { //
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`

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

	DomainsAndPages []string `json:"domainsAndPages"`
	AppBundles      []string `json:"appBundles"`
	// Available formats:
	// - 192.168.0.1 single IP
	// - 192.168.0.0 - 192.168.210.19 range of IPs
	// - 192.168 short form of range (equal to 192.168.0.0 - 192.168.255.255)
	IPAddresses []string `json:"ipAddresses"`

	Countries []Country `json:"countries"`
	Regions   []Region  `json:"regions"`
	Cities    []City    `json:"cities"`
	GeoAreas  []Area    `json:"geoAreas"`

	// Available values:
	// - 1 "AND" condition. All advertising segments in which a particular user is located must be targeted for this advertising campaign.
	// - 2 "OR" condition. At least one advertising segment in which a particular user is located must be targeted for this advertising campaign.
	// - 3 "NOT" condition. No advertising segment in which a particular user is located should be targeted for this advertising campaign.
	Segments []int `json:"segments"`

	Makers     []Maker    `json:"makers"`
	Browsers   []Browser  `json:"browsers"`
	Categories []Category `json:"categories"`
	Carriers   []Carrier  `json:"carriers"`
	OS         []OS       `json:"os"`
	Languages  []Language `json:"languages"`

	Name       string `json:"name"`
	ID         int    `json:"id"`
	CampaignID int    `json:"campaignId"`
	Status     int    `json:"status"`
	// Cost per thousand impressions
	CPM          float64 `json:"cpm"`
	FrequencyCap int     `json:"frequencyCap"`
	// FrequencyPeriodDay or FrequencyPeriodWeek or FrequencyPeriodMonth
	FrequencyPeriod        int     `json:"frequencyPeriod"`
	DailyBudget            float64 `json:"dailyBudget"`
	SegmentsMatchCondition int     `json:"segmentsMatchCondition"`

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

// ScheduleWeek - is an array of boolean arrays, consisting of time schedules.
// An array consists of 24 elements corresponding to 24 hours. The first element of the array sets the delivery schedule from 00:00 to 01:00.
// true - show ads at the specified hour, false - do not show.
// JSON representation:
//  	{
// 			"week": [
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, false],
// 				[false, false, false, false, false, false, true],
// 				[false, false, false, false, false, false, true],
// 				[false, false, false, false, false, false, true],
// 				[false, false, false, false, false, false, true],
// 				[false, false, false, false, false, false, true],
// 				[false, false, false, false, false, false, true],
// 				[false, false, false, false, false, false, true],
// 				[false, false, false, false, false, false, true]
// 			],
// 			"showInHolidays": false,
// 			"useLocalTime": false
// 		}
type ScheduleWeek struct {
	// An array of Boolean variables, consisting of flags for impressions by day of the week.
	// The array consists of 7 elements corresponding to 7 days of the week.
	// The first element of the array sets the ad to run on Monday.
	// true - show ads on the specified day of the week, false - do not show.
	Week           [24][7]bool `json:"week"`
	ShowInHolidays bool        `json:"showInHolidays"`
	UseLocalTime   bool        `json:"useLocalTime"`
}
