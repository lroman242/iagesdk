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

// CreateCampaignRequest describe API request to iAGE platform to create new campaign
type CreateCampaignRequest struct {
	AdvertiserID int       `json:"-"`         //required
	Name         string    `json:"name"`      //required, max 255
	Status       int       `json:"status"`    //required, 1 or 2
	StartDate    time.Time `json:"startDate"` //required, RFC-3339, > today
	EndDate      time.Time `json:"endDate"`   //required, RFC-3339, > today, > StartDate
	Budget       float64   `json:"budget"`
}

// NewCreateCampaignRequest initialize CreateCampaignRequest based on Campaign instance and Advertiser id
func NewCreateCampaignRequest(c *Campaign, advertiserID int) *CreateCampaignRequest {
	return &CreateCampaignRequest{
		AdvertiserID: advertiserID,
		Name:         c.Name,
		Status:       c.Status,
		StartDate:    c.StartDate,
		EndDate:      c.EndDate,
		Budget:       c.Budget,
	}
}

// URL return API request entrypoint (URI)
func (ccr *CreateCampaignRequest) URL() string {
	return fmt.Sprintf(`/v1/advertisers/%v/campaigns`, ccr.AdvertiserID)
}

// Method return API request http method
func (ccr *CreateCampaignRequest) Method() string {
	return http.MethodPost
}

// Body generate body content of API request
func (ccr *CreateCampaignRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(ccr)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
