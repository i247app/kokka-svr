package http_client

import (
	"maps"
	"net/http"
	"time"
)

// Option is a function that configures the Client
type Option func(*Client)

// RequestOption is a function that configures a specific request
type RequestOption func(*RequestConfig)

// RequestConfig holds configuration for a specific request
type RequestConfig struct {
	headers     map[string]string
	queryParams map[string]string
}

// ===============================
// Client-level Options
// ===============================

// WithBaseURL sets the base URL for all requests
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithTimeout sets the HTTP client timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithHeader sets a custom header for all requests
func WithHeader(key, value string) Option {
	return func(c *Client) {
		c.headers[key] = value
	}
}

// WithHeaders sets multiple custom headers
func WithHeaders(headers map[string]string) Option {
	return func(c *Client) {
		maps.Copy(c.headers, headers)
	}
}

// WithQueryParam sets a query parameter for all requests
func WithQueryParam(key, value string) Option {
	return func(c *Client) {
		c.queryParams[key] = value
	}
}

// WithQueryParams sets multiple query parameters
func WithQueryParams(params map[string]string) Option {
	return func(c *Client) {
		maps.Copy(c.queryParams, params)
	}
}

// WithInterceptor adds an interceptor to the client
func WithInterceptor(interceptor Interceptor) Option {
	return func(c *Client) {
		c.interceptors = append(c.interceptors, interceptor)
	}
}

// WithRetry configures retry behavior
func WithRetry(maxRetries int, retryDelay time.Duration, retryableHTTPCodes ...int) Option {
	return func(c *Client) {
		c.retryConfig = &RetryConfig{
			MaxRetries:         maxRetries,
			RetryDelay:         retryDelay,
			RetryableHTTPCodes: retryableHTTPCodes,
		}
	}
}

// ===============================
// Request-level Options
// ===============================

// WithRequestHeader sets a header for a specific request
func WithRequestHeader(key, value string) RequestOption {
	return func(rc *RequestConfig) {
		rc.headers[key] = value
	}
}

// WithRequestHeaders sets multiple headers for a specific request
func WithRequestHeaders(headers map[string]string) RequestOption {
	return func(rc *RequestConfig) {
		maps.Copy(rc.headers, headers)
	}
}

// WithRequestQueryParam sets a query parameter for a specific request
func WithRequestQueryParam(key, value string) RequestOption {
	return func(rc *RequestConfig) {
		rc.queryParams[key] = value
	}
}

// WithRequestQueryParams sets multiple query parameters for a specific request
func WithRequestQueryParams(params map[string]string) RequestOption {
	return func(rc *RequestConfig) {
		maps.Copy(rc.queryParams, params)
	}
}
