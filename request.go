package retaicrm

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (c *Client) SendGetRequest(url string) (*http.Response, error) {
	if c.DomainName == "" || c.ApiKey == "" {
		return nil, fmt.Errorf("domain name or api key parameters are not filled")
	}
	httpClient := &http.Client{Timeout: time.Minute}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", c.ApiKey)

	if c.Debug {
		logMessage := fmt.Sprintf("API Request: %s %s", url, c.ApiKey)
		log.Println(logMessage)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode >= http.StatusInternalServerError {
		return resp, fmt.Errorf("HTTP request error. Status code: %d", resp.StatusCode)
	}

	return resp, nil
}
