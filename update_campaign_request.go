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

// UpdateCampaignRequest describe API request to iAGE platform to update advertiser
type UpdateCampaignRequest struct {
	ID        int       `json:"-"`         //required
	Name      string    `json:"name"`      //required, max 255
	Status    int       `json:"status"`    //required, 1 or 2
	StartDate time.Time `json:"startDate"` //required, RFC-3339, > today
	EndDate   time.Time `json:"endDate"`   //required, RFC-3339, > today, > StartDate
	Budget    float64   `json:"budget"`
}

// NewUpdateCampaignRequest initialize UpdateCampaignRequest based on Campaign instance
func NewUpdateCampaignRequest(c *Campaign) *UpdateCampaignRequest {
	return &UpdateCampaignRequest{
		ID:        c.ID,
		Name:      c.Name,
		Status:    c.Status,
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
		Budget:    c.Budget,
	}
}

// URL return API request entrypoint (URI)
func (ucr *UpdateCampaignRequest) URL() string {
	return fmt.Sprintf(`/v1/campaigns/%v`, ucr.ID)
}

// Method return API request http method
func (ucr *UpdateCampaignRequest) Method() string {
	return http.MethodPut
}

// Body generate body content of API request
func (ucr *UpdateCampaignRequest) Body() io.Reader {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(ucr)
	if err != nil {
		log.Print(err)
		return nil
	}

	return body
}
