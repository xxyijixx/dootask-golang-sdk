package sdk

import (
	"time"
)

// Config holds client configuration
type Config struct {
	BaseURL    string
	Timeout    time.Duration
	UserAgent  string
	Insecure   bool
	Debug      bool
	RetryCount int
}

// DefaultConfig returns default client configuration
func DefaultConfig() *Config {
	return &Config{
		Timeout:    30 * time.Second,
		UserAgent:  "dootask-golang-sdk/1.0",
		Insecure:   false,
		Debug:      false,
		RetryCount: 3,
	}
}

// WithTimeout sets the timeout for HTTP requests
func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}

// WithUserAgent sets the user agent for HTTP requests
func (c *Config) WithUserAgent(userAgent string) *Config {
	c.UserAgent = userAgent
	return c
}

// WithInsecure sets whether to skip TLS verification
func (c *Config) WithInsecure(insecure bool) *Config {
	c.Insecure = insecure
	return c
}

// WithDebug enables debug logging
func (c *Config) WithDebug(debug bool) *Config {
	c.Debug = debug
	return c
}

// WithRetryCount sets the number of retries for failed requests
func (c *Config) WithRetryCount(count int) *Config {
	c.RetryCount = count
	return c
}
