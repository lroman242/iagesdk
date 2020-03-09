package iagesdk

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
	"time"
)

func TestClient_Sandbox(t *testing.T) {
	c := Client{
		APIToken: "someToken",
	}

	c.Sandbox()

	if c.baseURL != sandboxAPI {
		t.Errorf("invalid client endpoint. sandbox endpoint expected but got %v", c.baseURL)
	}
}

func TestClient_Production(t *testing.T) {
	c := Client{
		APIToken: "someToken",
	}

	c.Production()

	if c.baseURL != productionAPI {
		t.Errorf("invalid client endpoint. production endpoint expected but got %v", c.baseURL)
	}
}

func TestNewClient_Sandbox(t *testing.T) {
	token := "someToken1"
	c := NewClient(token, true)

	if c.APIToken != token {
		t.Errorf("invalid token received. expected %v but got %v", token, c.APIToken)
	}

	if c.baseURL != sandboxAPI {
		t.Error("invalid client endpoint. sandbox endpoint expected")
	}
}

func TestNewClient_Production(t *testing.T) {
	token := "someToken2"
	c := NewClient(token, false)

	if c.APIToken != token {
		t.Errorf("invalid token received. expected %v but got %v", token, c.APIToken)
	}

	if c.baseURL != productionAPI {
		t.Error("invalid client endpoint. production endpoint expected")
	}
}

func TestClient_getURL_Sandbox(t *testing.T) {
	c := Client{
		APIToken: "someToken",
	}

	c.Sandbox()

	if c.getURL() != sandboxAPI {
		t.Errorf("invalid client endpoint. sandbox endpoint expected but got %v", c.baseURL)
	}
}

func TestClient_getURL_Production(t *testing.T) {
	c := Client{
		APIToken: "someToken",
	}

	c.Production()

	if c.getURL() != productionAPI {
		t.Errorf("invalid client endpoint. production endpoint expected but got %v", c.baseURL)
	}
}

func TestClient_getURL_Empty(t *testing.T) {
	c := Client{
		APIToken: "someToken",
	}

	if c.getURL() != productionAPI {
		t.Errorf("invalid client endpoint. production endpoint expected but got %v", c.baseURL)
	}
}

func TestClient_SendRequest(t *testing.T) {
	godotenv.Load(".env")
	c := NewClient(os.Getenv("IAGE_TOKEN"), true)

	t.Run("sendRequest_success", func(t *testing.T) {
		timestamp := time.Now().Unix()
		car := NewCreateAgencyRequest(&Agency{
			Name:       fmt.Sprintf("TestAgency-%v", timestamp),
			Country:    "Malaysia",
			Status:     1,
			Email:      fmt.Sprintf("test-%v@agency.com", timestamp),
			WebSiteURL: "https://test-agency.com",
		})

		var resp interface{}

		err := c.SendRequest(car, &resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("sendRequest_error", func(t *testing.T) {
		timestamp := time.Now().Unix()
		car := NewCreateAgencyRequest(&Agency{
			Name:    fmt.Sprintf("TestAgency-%v", timestamp),
			Country: "Malaysia",
			Status:  1,
			// cause validation error
			//Email: fmt.Sprintf("test-%v@agency.com", timestamp),
			WebSiteURL: "https://test-agency.com",
		})

		var resp interface{}

		err := c.SendRequest(car, &resp)
		if err == nil {
			t.Error("validation error expected")
		}
	})

	t.Run("sendRequest_errorUnmarshal", func(t *testing.T) {
		timestamp := time.Now().Unix()
		car := NewCreateAgencyRequest(&Agency{
			Name:       fmt.Sprintf("TestAgency-%v", timestamp),
			Country:    "Malaysia",
			Status:     1,
			Email:      fmt.Sprintf("test-%v@agency.com", timestamp),
			WebSiteURL: "https://test-agency.com",
		})

		// invalid response struct
		// cause unmarshal error
		var resp int

		err := c.SendRequest(car, &resp)
		if err == nil {
			t.Error("unmarshal error expected")
		}
	})

}
