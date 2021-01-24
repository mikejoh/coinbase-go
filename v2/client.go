package v2

import (
	"net/http"
	"time"
)

const (
	apiBaseURL       = "https://api.coinbase.com/v2"
	defaultUserAgent = "coinbase-go"
)

type Client struct {
	client *http.Client
	Config *Config
}

type Config struct {
	apiKey    string
	secret    string
	baseURL   string
	userAgent string
	timeout   time.Duration
	debug     bool
}

func NewConfig(options ...Option) *Config {
	config := &Config{
		baseURL:   apiBaseURL,
		userAgent: defaultUserAgent,
	}

	for _, addOpt := range options {
		addOpt(config)
	}

	return config
}

func NewClient(config *Config) *Client {
	client := &Client{
		client: &http.Client{
			Timeout: config.timeout,
		},
		Config: config,
	}

	return client
}

type Option func(*Config)

func ApiKey(k string) func(*Config) {
	return func(c *Config) { c.apiKey = k }
}

func Secret(s string) func(*Config) {
	return func(c *Config) { c.secret = s }
}

func Timeout(t time.Duration) func(*Config) {
	return func(c *Config) { c.timeout = t }
}

func UserAgent(ua string) func(*Config) {
	return func(c *Config) { c.userAgent = ua }
}

func Debug(b bool) func(*Config) {
	return func(c *Config) { c.debug = b }
}
