package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type user struct {
	Data struct {
		ID              string      `json:"id"`
		Name            string      `json:"name"`
		Username        string      `json:"username"`
		ProfileLocation interface{} `json:"profile_location"`
		ProfileBio      interface{} `json:"profile_bio"`
		ProfileURL      string      `json:"profile_url"`
		AvatarURL       string      `json:"avatar_url"`
		Resource        string      `json:"resource"`
		ResourcePath    string      `json:"resource_path"`
	} `json:"data"`
}

func (c *Client) GetUser(ctx context.Context) (*user, error) {
	endpoint := "user"
	method := "GET"
	rawURL := fmt.Sprintf("%s/%s", c.Config.baseURL, endpoint)

	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	fmt.Println(u.Path)

	req, err := c.NewRequest(method, u, true, nil)
	if err != nil {
		return nil, err
	}

	req.WithContext(ctx)

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

	var userData user

	if err = json.NewDecoder(res.Body).Decode(&userData); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return &userData, nil
}
