package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type prices struct {
	Data struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency`
	} `json:"data"`
	PriceType    string `json:"price_type"`
	CurrencyPair string `json:"currency_pair"`
}

// Prices
func (c *Client) Prices(ctx context.Context, currencyPair string, priceType string) (*prices, error) {
	endpoint := "prices"
	method := "GET"

	switch priceType {
	case "sell":
	case "buy":
	case "spot":
	default:
		return nil, errors.New("unsupported price type: " + priceType)
	}

	path := fmt.Sprintf("%s/%s/%s/%s", c.Config.baseURL, endpoint, currencyPair, priceType)

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

	var prices prices

	prices.PriceType = priceType
	prices.CurrencyPair = currencyPair

	if err = json.NewDecoder(res.Body).Decode(&prices); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return &prices, nil
}

func (p *prices) String() string {
	return fmt.Sprintf("%s %s @ %s %s", p.CurrencyPair, p.PriceType, p.Data.Amount, p.Data.Currency)
}
