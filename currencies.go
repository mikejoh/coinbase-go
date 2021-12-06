package coinbase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type currency struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	MinSize string `json:"min_size"`
}

type currencies struct {
	Data []currency `json:"data"`
}

// Currencies lists known currencies. 
func (c *Client) Currencies(ctx context.Context) (*currencies, error) {
	endpoint := "currencies"
	method := "GET"

	path := fmt.Sprintf("%s/%s", c.Config.baseURL, endpoint)

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

	var currencies currencies

	if err = json.NewDecoder(res.Body).Decode(&currencies); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return &currencies, nil
}

func (c *currencies) String() string {
	fmt.Println("currencies:")
	for _, v := range c.Data {
		fmt.Printf("%s (%s) %s\n", v.ID, v.Name, v.MinSize)
	}

	return ""
}
