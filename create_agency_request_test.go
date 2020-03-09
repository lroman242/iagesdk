package iagesdk

import (
	"testing"
	"github.com/joho/godotenv"
	"os"
	"time"
	"fmt"
)

func TestNewCreateAgencyRequest(t *testing.T) {
	a := Agency{
		WebSiteURL: "website.com",
		Email:      "email@test.com",
		Status:     1,
		Country:    "Malaysia",
		Name:       "TestAgency",
		ID:         32,
	}

	r := NewCreateAgencyRequest(&a)

	if r.Name != a.Name {
		t.Errorf("invalid value of field Name populated to CreateAgencyRequest. expected %v but got %v\n", a.Name, r.Name)
	}
	if r.Status != a.Status {
		t.Errorf("invalid value of field Status populated to CreateAgencyRequest. expected %v but got %v\n", a.Status, r.Status)
	}
	if r.Country != a.Country {
		t.Errorf("invalid value of field Country populated to CreateAgencyRequest. expected %v but got %v\n", a.Country, r.Country)
	}
	if r.Email != a.Email {
		t.Errorf("invalid value of field Email populated to CreateAgencyRequest. expected %v but got %v\n", a.Email, r.Email)
	}
	if r.WebSiteURL != a.WebSiteURL {
		t.Errorf("invalid value of field WebSiteURL populated to CreateAgencyRequest. expected %v but got %v\n", a.WebSiteURL, r.WebSiteURL)
	}
}

func TestCreateAgencyRequest(t *testing.T) {
	godotenv.Load(".env")
	c := NewClient(os.Getenv("IAGE_TOKEN"), true)
	timestamp := time.Now().Unix()
	car := NewCreateAgencyRequest(&Agency{
		Name:       fmt.Sprintf("TestAgency-%v", timestamp),
		Country:    "Malaysia",
		Status:     1,
		Email:      fmt.Sprintf("test-%v@agency.com", timestamp),
		WebSiteURL: "https://test-agency.com",
	})

	resp := new(CreateAgencyResponse)

	err := c.SendRequest(car, &resp)
	if err != nil {
		t.Errorf("unexpected error during CreateAgencyRequest API call: %v", err)
	} else {
		if car.Name != resp.Name {
			t.Errorf("invalid value of field Name received in CreateAgencyResponse. expected %v but got %v\n", car.Name, resp.Name)
		}
		if car.Status != resp.Status {
			t.Errorf("invalid value of field Status received in CreateAgencyResponse. expected %v but got %v\n", car.Status, resp.Status)
		}
		if car.Country != resp.Country {
			t.Errorf("invalid value of field Country received in CreateAgencyResponse. expected %v but got %v\n", car.Country, resp.Country)
		}
		if car.Email != resp.Email {
			t.Errorf("invalid value of field Email received in CreateAgencyResponse. expected %v but got %v\n", car.Email, resp.Email)
		}
		if car.WebSiteURL != resp.WebSiteURL {
			t.Errorf("invalid value of field WebSiteURL received in CreateAgencyResponse. expected %v but got %v\n", car.WebSiteURL, resp.WebSiteURL)
		}
	}
}