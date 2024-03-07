package zebedee

import (
	"net/http"
)

// BaseURL is https://api.zebedee.io/v0 by default.
type Client struct {
	environment Environment
	HttpClient  *http.Client
	BaseURL     string
	APIKey      string

	// pluggable http requester to make
	// life easy to mock and modify http
	// requests for tests and different contexts
	Requester Requester
}

func New(apikey string) *Client {
	return &Client{
		BaseURL: endpoints[Production],
		APIKey:  apikey,

		environment: Production,
		HttpClient:  &http.Client{},
		Requester:   &EndpointRequester{},
	}
}

func NewSandboxClient(apikey string) *Client {
	return &Client{
		environment: Sandbox,
		BaseURL:     endpoints[Sandbox],
		APIKey:      apikey,
		HttpClient:  &http.Client{},
		Requester:   &EndpointRequester{},
	}
}

func NewPublicClient(apikey string) *Client {
	return &Client{
		environment: Public,
		BaseURL:     endpoints[Public],
		APIKey:      apikey,
		HttpClient:  &http.Client{},
		Requester:   &EndpointRequester{},
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
