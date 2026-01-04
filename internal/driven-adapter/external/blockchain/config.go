package blockchain

import "time"

// Config holds the configuration for the blockchain client
type Config struct {
	BaseURL        string
	Timeout        time.Duration
	MaxRetries     int
	RetryDelay     time.Duration
	EnableLogging  bool
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		BaseURL:        "https://x24.i247.com",
		Timeout:        30 * time.Second,
		MaxRetries:     3,
		RetryDelay:     1 * time.Second,
		EnableLogging:  true,
	}
}

// WithBaseURL sets a custom base URL
func (c *Config) WithBaseURL(baseURL string) *Config {
	c.BaseURL = baseURL
	return c
}

// WithTimeout sets a custom timeout
func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}

// WithRetry sets retry configuration
func (c *Config) WithRetry(maxRetries int, retryDelay time.Duration) *Config {
	c.MaxRetries = maxRetries
	c.RetryDelay = retryDelay
	return c
}

// WithLogging enables or disables logging
func (c *Config) WithLogging(enabled bool) *Config {
	c.EnableLogging = enabled
	return c
}
