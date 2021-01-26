package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type serverTime struct {
	Data struct {
		ISO   string `json:"iso"`
		Epoch int64  `json:"epoch"`
	} `json:"data"`
}

// Time
func (c *Client) Time(ctx context.Context) (*serverTime, error) {
	endpoint := "time"
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

	var serverTime serverTime

	if err = json.NewDecoder(res.Body).Decode(&serverTime); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return &serverTime, nil
}

func (t *serverTime) String() string {
	return fmt.Sprintf("%s (%d)", t.Data.ISO, t.Data.Epoch)
}
