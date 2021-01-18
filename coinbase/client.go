package coinbase

import (
	"net/http"
	"time"
)

// TODO: Rate limit
// TODO: Pagination
// TODO: GET, POST, PUT, DELETE ops

const (
	apiBaseURL = "https://api.coinbase.com/v2/"
	userAgent  = "coinbase-go"
)

type Client struct {
	client    *http.Client
	apiKey    string
	secret    string
	baseURL   string
	userAgent string
	debug     bool
}

func New(options ...Option) *Client {
	c := &Client{
		client:    &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL:   apiBaseURL,
		userAgent: userAgent,
		debug:     false,
	}

	for _, addOpt := range options {
		addOpt(c)
	}

	return c
}

type Option func(*Client)

func ApiKey(apiKey string) func(*Client) {
	return func(c *Client) { c.apiKey = apiKey }
}

func Secret(secret string) func(*Client) {
	return func(c *Client) { c.secret = secret }
}

func Timeout(t time.Duration) func(*Client) {
	return func(c *Client) { c.client.Timeout = t }
}

func UserAgent(ua string) func(*Client) {
	return func(c *Client) { c.userAgent = ua }
}

func Debug(b bool) func(*Client) {
	return func(c *Client) { c.debug = b }
}
