package v2

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func (c *Client) NewRequest(method string, u *url.URL, authenticate bool, json []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if authenticate {
		epoch := strconv.FormatInt(time.Now().Unix(), 10)

		unsigned := epoch + strings.ToUpper(method) + u.Path

		if len(json) > 0 {
			unsigned = unsigned + string(json)
		}

		req.Header.Set("CB-ACCESS-KEY", c.Config.apiKey)
		req.Header.Set("CB-ACCESS-TIMESTAMP", epoch)

		h := hmac.New(sha256.New, []byte(c.Config.secret))
		h.Write([]byte(unsigned))

		signed := hex.EncodeToString(h.Sum(nil))

		req.Header.Set("CB-ACCESS-SIGN", signed)
	}

	if req.Method != "GET" {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
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
