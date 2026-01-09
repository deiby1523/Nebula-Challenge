package ssllabs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Base URL for SSL Labs API v2
const baseURL = "https://api.ssllabs.com/api/v2"


type Client struct {
	httpClient *http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

/*
check() verifies basic connectivity with the
SSL Labs API
*/
func (c *Client) check() error {
	resp, err := c.httpClient.Get(baseURL + "/info")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return validateHTTPStatus(resp)
}

/*
analyze() triggers or retrieves a TLS analysis
for a given domain. If startNew is true, a new
analysis is explicitly started.
*/
func (c *Client) analyze(domain string, startNew bool) (*Response, error) {
	url := fmt.Sprintf("%s/analyze?host=%s&all=done", baseURL, domain)
	if startNew {
		url += "&startNew=on"
	}

	// use of http.Client instead of http.Get()
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := validateHTTPStatus(resp); err != nil {
		return nil, err
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}