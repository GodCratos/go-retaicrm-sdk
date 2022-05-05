package retaicrm

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
)

const (
	baseUrl = "https://%s/api/v5"

	endpointOrdersGet = "/orders?%s"
)

type Client struct {
	DomainName string
	ApiKey     string
	Debug      bool
}

func NewClient(domainName, apiKey string, debug bool) *Client {
	return &Client{
		DomainName: domainName,
		ApiKey:     apiKey,
		Debug:      debug,
	}
}

//{WORKING WITH ORDERS

const ByID = "id"
const ByExternalID = "externalId"

func (c *Client) OrdersGet(parameters OrdersRequest) (*OrdersResponse, int, error) {
	var dataResp OrdersResponse
	param, err := query.Values(parameters)
	if err != nil {
		return &dataResp, 0, err
	}
	url := fmt.Sprintf("%s%s", fmt.Sprintf(baseUrl, c.DomainName), fmt.Sprintf(endpointOrdersGet, param.Encode()))
	resp, err := c.SendGetRequest(url)
	if err != nil {
		return &dataResp, 0, err
	}
	json.NewDecoder(resp.Body).Decode(&dataResp)
	return &dataResp, resp.StatusCode, nil
}

//WORKING WITH ORDERS}
