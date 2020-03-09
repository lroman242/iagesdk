package iagesdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const sandboxAPI = "https://demo.iae.one"
const productionAPI = "https://api.dsp.iage.com"

// Client manages communication with the iAGE API.
type Client struct {
	// Base URL for API requests.
	baseURL string

	// API token to authorize requests
	APIToken string
}

// Sandbox method set sandbox url to client
func (c *Client) Sandbox() *Client {
	c.baseURL = sandboxAPI

	return c
}

// Production method set production url to client
func (c *Client) Production() *Client {
	c.baseURL = productionAPI

	return c
}

// getUrl is a method to get base api url
func (c *Client) getURL() string {
	if c.baseURL == "" {
		c.Production()
	}

	return c.baseURL
}

// SendRequest sends an API request and populates the given interface with the parsed
// response. It does not make much sense to call SendRequest without a prepared
// interface instance.
func (c Client) SendRequest(r Request, resp interface{}) error {
	req, _ := http.NewRequest(
		r.Method(),
		fmt.Sprintf("%s%s", c.getURL(), r.URL()),
		r.Body(),
	)

	req.Header.Add("X-Auth-Token", c.APIToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	bodyContents, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		errorMessage, err := c.parseErrorResponse(bodyContents)
		if err != nil {
			return err
		}

		return errors.New(errorMessage)
	}

	if err := json.Unmarshal(bodyContents, resp); err != nil {
		return err
	}

	return nil
}

// parseErrorResponse pares received response and retrieve error message from it
func (c Client) parseErrorResponse(content []byte) (string, error) {
	errorResponse := ErrorResponse{}

	if err := json.Unmarshal(content, &errorResponse); err != nil {
		return "", err
	}

	return errorResponse.Error, nil
}

// NewClient - create new instance of iAGE`s HTTP API client
func NewClient(APIToken string, sandbox bool) *Client {
	c := &Client{
		APIToken: APIToken,
	}

	if sandbox {
		c.Sandbox()
	} else {
		c.Production()
	}

	return c
}
