package coinbase

/*
func (c *Client) GetExchangeRates(ctx context.Context, currency string) (*ExchangeRates, error) {
	path := "exchange-rates"
	method := "GET"

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s?currency=%s", c.baseURL.String(), path, currency), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	now := time.Now().Unix() // Within 30 seconds of server time
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("CB-ACCESS-KEY", c.config.apiKey)
	req.Header.Set("CB-ACCESS-SIGN", c.signHeader(now, path, strings.ToUpper(method)))
	req.Header.Set("CB-ACCESS-TIMESTAMP", fmt.Sprint(now))

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

	var exchangeRates ExchangeRates

	if err = json.NewDecoder(res.Body).Decode(&exchangeRates); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return &exchangeRates, nil
}

func (er *ExchangeRates) String() string {
	fmt.Printf("currency: %s\n", er.Data.Currency)
	fmt.Println("rates:")
	for k, v := range er.Data.Rates {
		fmt.Printf("\t%s = %s\n", k, v)
	}

	return ""
}
*/
