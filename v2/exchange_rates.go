package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type exchangeRates struct {
	Data struct {
		Currency string            `json:"currency"`
		Rates    map[string]string `json:"rates"`
	} `json:"data"`
}

// ExchangeRates
func (c *Client) ExchangeRates(ctx context.Context, currency string) (*exchangeRates, error) {
	endpoint := "exchange-rates"
	method := "GET"

	path := fmt.Sprintf("%s/%s", c.Config.baseURL, endpoint)

	if currency != "" {
		path = fmt.Sprintf("%s/%s?currency=%s", c.Config.baseURL, endpoint, currency)
	}

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse

		errRes.Code = res.StatusCode
		errRes.URL = req.URL.String()

		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return nil, errRes
		}

		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	var exchangeRates exchangeRates

	if err = json.NewDecoder(res.Body).Decode(&exchangeRates); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return &exchangeRates, nil
}

func (er *exchangeRates) String() string {
	fmt.Printf("currency: %s\n", er.Data.Currency)
	fmt.Println("rates:")
	for k, v := range er.Data.Rates {
		fmt.Printf("\t%s = %s\n", k, v)
	}

	return ""
}
