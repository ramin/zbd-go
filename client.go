package zebedee

import (
	"net/http"
)

// BaseURL is https://api.zebedee.io/v0 by default.
type Client struct {
	BaseURL    string
	APIKey     string
	HttpClient *http.Client
}

func New(apikey string) *Client {
	return &Client{
		BaseURL:    endpoints[Production],
		APIKey:     apikey,
		HttpClient: &http.Client{},
	}
}

func NewSandboxClient(apikey string) *Client {
	return &Client{
		BaseURL:    endpoints[Sandbox],
		APIKey:     apikey,
		HttpClient: &http.Client{},
	}
}

func (c *Client) Sandbox() *Client {
	c.BaseURL = endpoints[Sandbox]
	return c
}

func (c *Client) WithEnvironment(env Environment) *Client {
	c.BaseURL = endpoints[env]
	return c
}
