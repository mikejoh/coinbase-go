package v2

import "testing"

func TestNewConfig(t *testing.T) {
	tests := map[string]struct {
		config *Config
		want   string
	}{
		"Set Base URL":     {NewConfig(BaseURL("https://mock.local")), "https://mock.local"},
		"Default Base URL": {NewConfig(), "https://api.coinbase.com/v2"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.config.baseURL
			if got != tc.want {
				t.Fatalf("expected %s got %s", tc.want, got)
			}
		})
	}
}

/*
func TestAuthenticatedNewRequest(t *testing.T) {
	fakeApiKey := "key"
	fakeApiSecret := "secret"
	fakeEndpoint := "https://api.fake"
	fakeJSON := []byte(`{"fake":"data"}`)

	config := NewConfig(
		ApiKey(fakeApiKey),
		Secret(fakeApiSecret),
	)

	client := NewClient(config)

	tests := map[string]struct {
		method         string
		endpoint       string
		json           []byte
		authenticate   bool
		expectedHeader http.Header
	}{
		"Authenticated GET":  {method: "GET", endpoint: fakeEndpoint, json: nil, authenticate: true, expectedHeader: http.Header{"CB-ACCESS-KEY": []string{"key"}}},
		"Authenticated POST": {method: "POST", endpoint: fakeEndpoint, json: fakeJSON, authenticate: true, expectedHeader: http.Header{"CB-ACCESS-KEY": []string{"key"}}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := client.NewRequest(tc.method, tc.endpoint, tc.authenticate, tc.json)
			if err != nil {
				t.Error(err)
			}

			got := req.Header.Get("CB-ACCESS-KEY")
			want := tc.expectedHeader.Get("CB-ACCESS-KEY")

			t.Log(tc.expectedHeader)

			if want != got {
				t.Fatalf("expected %s got %s", want, got)
			}
		})
	}
}
*/
