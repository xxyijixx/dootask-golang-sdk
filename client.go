package sdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/xxyijixx/dootask-golang-sdk/api/dialog"
	"github.com/xxyijixx/dootask-golang-sdk/api/file"
	ihttp "github.com/xxyijixx/dootask-golang-sdk/internal/http"
)

// Client represents the main API client
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
	Config     *Config

	// Service modules
	File   *file.Service
	Dialog *dialog.Service
}

// NewClient creates a new API client
func NewClient(baseURL string) *Client {
	return NewClientWithConfig(baseURL, DefaultConfig())
}

// NewClientWithConfig creates a new API client with custom configuration
func NewClientWithConfig(baseURL string, config *Config) *Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.Insecure,
		},
	}

	client := &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout:   config.Timeout,
			Transport: transport,
		},
		Config: config,
	}

	// Initialize service modules
	client.File = file.New(client)
	client.Dialog = dialog.New(client)

	return client
}

// SetToken sets the authentication token
func (c *Client) SetToken(token string) {
	c.Token = token
}

// DoRequest performs an HTTP request
func (c *Client) DoRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	return c.DoRequestWithHeaders(method, endpoint, body, nil)
}

// DoRequestWithHeaders performs an HTTP request with additional headers
func (c *Client) DoRequestWithHeaders(method, endpoint string, body interface{}, headers map[string]string) (*http.Response, error) {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	url := c.BaseURL + endpoint
	if c.Config.Debug {
		ihttp.LogRequest(method, url, body)
	}

	req, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return nil, err
	}

	// Set default headers
	req.Header.Set("Content-Type", "application/json")
	if c.Config.UserAgent != "" {
		req.Header.Set("User-Agent", c.Config.UserAgent)
	}

	// Set authentication token if available
	if c.Token != "" {
		req.Header.Set("Token", c.Token)
	}

	// Set additional headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if c.Config.Debug {
		ihttp.LogResponse(resp)
	}

	return resp, nil
}
