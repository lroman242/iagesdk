package iagesdk

// Creative describe iAGE`s base creative data structure
type Creative struct {
	ID           int  `json:"id"`
	CampaignID   int  `json:"campaignId"`
	Disabled     bool `json:"disabled"`
	TemplateType int  `json:"templateType"`
}
