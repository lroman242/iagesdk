package iagesdk

import (
	"testing"
	"github.com/joho/godotenv"
	"time"
	"os"
	"fmt"
)

func TestNewCreateAdvertiserRequest(t *testing.T) {
	adv := Advertiser{
		Country: "Malaysia",
		Email: "advertiser@test.com",
		Status: 1,
		City: "Kuala Lumpur",
		FirstName: "Test",
		LastName: "Advertiser",
		Currency: "MYR",
		Phone: "+60399999999",
		Skype: "test_adv",
	}

	password := "secret123"

	r := NewCreateAdvertiserRequest(&adv, password)

	if r.FirstName != adv.FirstName {
		t.Errorf("invalid value of field FirstName populated to CreateAdvertiserRequest. expected %v but got %v\n", adv.FirstName, r.FirstName)
	}
	if r.LastName != adv.LastName {
		t.Errorf("invalid value of field LastName populated to CreateAdvertiserRequest. expected %v but got %v\n", adv.LastName, r.LastName)
	}
	if r.Status != adv.Status {
		t.Errorf("invalid value of field Status populated to CreateAdvertiserRequest. expected %v but got %v\n", adv.Status, r.Status)
	}
	if r.Country != adv.Country {
		t.Errorf("invalid value of field Country populated to CreateAdvertiserRequest. expected %v but got %v\n", adv.Country, r.Country)
	}
	if r.City != adv.City {
		t.Errorf("invalid value of field City populated to CreateAdvertiserRequest. expected %v but got %v\n", adv.City, r.City)
	}
	if r.Email != adv.Email {
		t.Errorf("invalid value of field Email populated to CreateAdvertiserRequest. expected %v but got %v\n", adv.Email, r.Email)
	}
	if r.Phone != adv.Phone {
		t.Errorf("invalid value of field Phone populated to CreateAdvertiserRequest. expected %v but got %v\n", adv.Phone, r.Phone)
	}
	if r.Skype != adv.Skype {
		t.Errorf("invalid value of field Skype populated to CreateAdvertiserRequest. expected %v but got %v\n", adv.Skype, r.Skype)
	}
	if r.Currency != adv.Currency {
		t.Errorf("invalid value of field Currency populated to CreateAdvertiserRequest. expected %v but got %v\n", adv.Currency, r.Currency)
	}

}

func TestCreateAdvertiserRequest(t *testing.T) {
	godotenv.Load(".env")
	c := NewClient(os.Getenv("IAGE_TOKEN"), true)
	timestamp := time.Now().Unix()

	adv := Advertiser{
		Country: "Malaysia",
		Email: fmt.Sprintf("advertiser-%v@test.com", timestamp),
		Status: 1,
		City: "Kuala Lumpur",
		FirstName: "Test",
		LastName: "Advertiser",
		Currency: "USD",
		Phone: "+60399999999",
		Skype: "test_adv",
	}

	password := "secret123"

	r := NewCreateAdvertiserRequest(&adv, password)

	resp := new(CreateAdvertiserResponse)

	err := c.SendRequest(r, resp)
	if err != nil {
		t.Errorf("unexpected error during CreateAdvertiserResponse API call: %v", err)
	} else {
		if r.FirstName != resp.FirstName {
			t.Errorf("invalid value of field FirstName received in CreateAdvertiserResponse. expected %v but got %v\n", r.FirstName, resp.FirstName)
		}
		if r.LastName != resp.LastName {
			t.Errorf("invalid value of field LastName received in CreateAdvertiserResponse. expected %v but got %v\n", r.LastName, resp.LastName)
		}
		if r.Status != resp.Status {
			t.Errorf("invalid value of field Status received in CreateAdvertiserResponse. expected %v but got %v\n", r.Status, resp.Status)
		}
		if r.Country != resp.Country {
			t.Errorf("invalid value of field Country received in CreateAdvertiserResponse. expected %v but got %v\n", r.Country, resp.Country)
		}
		if r.City != resp.City {
			t.Errorf("invalid value of field City received in CreateAdvertiserResponse. expected %v but got %v\n", r.City, resp.City)
		}
		if r.Email != resp.Email {
			t.Errorf("invalid value of field Email received in CreateAdvertiserResponse. expected %v but got %v\n", r.Email, resp.Email)
		}
		if r.Phone != resp.Phone {
			t.Errorf("invalid value of field Phone received in CreateAdvertiserResponse. expected %v but got %v\n", r.Phone, resp.Phone)
		}
		if r.Skype != resp.Skype {
			t.Errorf("invalid value of field Skype received in CreateAdvertiserResponse. expected %v but got %v\n", r.Skype, resp.Skype)
		}
		if r.Currency != resp.Currency {
			t.Errorf("invalid value of field Currency received in CreateAdvertiserResponse. expected %v but got %v\n", r.Currency, resp.Currency)
		}
	}
}