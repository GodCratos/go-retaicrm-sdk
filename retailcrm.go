package retaicrm

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	baseUrl = "https://%s/api/v5"

	endpointOrdersGet   = "/orders?%s"
	endpointOrderCreate = "/orders/create"
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

// Example:

//	client := retailcrm.NewClient("example.retailcrm.ru", "<your_api_key>", true)
//	filter := retailcrm.OrdersRequest{Filter: retailcrm.OrdersFilter{Ids: []int{1}}, Page: 1}
//	data, _, err := client.OrdersGet(filter)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	log.Println(data)

func (c *Client) OrdersGet(parameters OrdersRequest) (*OrdersResponse, int, error) {
	var bodyResp OrdersResponse
	param, err := query.Values(parameters)
	if err != nil {
		return &bodyResp, 0, err
	}
	url := fmt.Sprintf("%s%s", fmt.Sprintf(baseUrl, c.DomainName), fmt.Sprintf(endpointOrdersGet, param.Encode()))
	resp, err := c.sendGetRequest(url)
	if err != nil {
		return &bodyResp, 0, err
	}
	json.NewDecoder(resp.Body).Decode(&bodyResp)
	return &bodyResp, resp.StatusCode, nil
}

/* Example:

 	client := retailcrm.NewClient("example.retailcrm.ru", "<your_api_key>", true)
	order := retailcrm.Order{
		FirstName:  "Ivan",
		LastName:   "Ivanov",
		Patronymic: "Ivanovich",
		Email:      "ivanov@example.com",
	}
	data, status, err := client.OrderCreate(&order, "<your_site>")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(data.ErrorMsg)

	if data.Success {
		log.Printf("%v\n %d", data.ID, status)
	}
	}*/
func (c *Client) OrderCreate(order *Order, site string) (*OrderCreateResponse, int, error) {
	var bodyResp OrderCreateResponse

	if order == nil || site == "" {
		return &bodyResp, 0, fmt.Errorf("order or site parameters are not filled")
	}
	orderJSON, err := json.Marshal(&order)
	if err != nil {
		return &bodyResp, 0, err
	}

	p := url.Values{
		"order": {string(orderJSON)}}
	p.Add("site", site)

	url := fmt.Sprintf("%s%s", fmt.Sprintf(baseUrl, c.DomainName), endpointOrderCreate)
	resp, err := c.sendPostRequest(url, strings.NewReader(p.Encode()))
	if err != nil {
		return nil, 0, err
	}
	json.NewDecoder(resp.Body).Decode(&bodyResp)
	return &bodyResp, 0, nil
}

func checkBy(by string) string {
	var context = ByID
	if context != by {
		context = ByExternalID
	}
	return context
}

//WORKING WITH ORDERS}
