/*
Eduframe - API

Retry helper for handling rate limits and network errors.

API version: v1
*/

package eduframe

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// RetryConfig configures retry behavior for API calls
type RetryConfig struct {
	// MaxRetries is the maximum number of retry attempts (default: 3)
	MaxRetries int
	// BaseDelay is the base delay for exponential backoff (default: 1 second)
	BaseDelay time.Duration
	// MaxDelay is the maximum delay between retries (default: 60 seconds)
	MaxDelay time.Duration
	// RetryOnStatuses are HTTP status codes that should trigger a retry (default: [429, 500, 502, 503, 504])
	RetryOnStatuses []int
}

// DefaultRetryConfig returns a default retry configuration
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries:      3,
		BaseDelay:       1 * time.Second,
		MaxDelay:        60 * time.Second,
		RetryOnStatuses: []int{429, 500, 502, 503, 504},
	}
}

// RetryableFunc represents a function that can be retried
type RetryableFunc func() (*http.Response, error)

// WithRetry wraps a function with retry logic for handling rate limits and transient errors
func WithRetry(ctx context.Context, config *RetryConfig, fn RetryableFunc) (*http.Response, error) {
	if config == nil {
		config = DefaultRetryConfig()
	}

	var lastResponse *http.Response
	var lastError error

	for attempt := 0; attempt <= config.MaxRetries; attempt++ {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return lastResponse, ctx.Err()
		default:
		}

		lastResponse, lastError = fn()

		// If no error and successful response, return immediately
		if lastError == nil && lastResponse != nil && !shouldRetry(lastResponse.StatusCode, config.RetryOnStatuses) {
			return lastResponse, nil
		}

		// If this was the last attempt, return the error
		if attempt == config.MaxRetries {
			break
		}

		// Calculate delay for next attempt
		delay := calculateDelay(lastResponse, attempt, config)

		// Wait before retrying
		select {
		case <-ctx.Done():
			return lastResponse, ctx.Err()
		case <-time.After(delay):
			// Continue to next attempt
		}
	}

	return lastResponse, lastError
}

// shouldRetry determines if a status code should trigger a retry
func shouldRetry(statusCode int, retryOnStatuses []int) bool {
	for _, code := range retryOnStatuses {
		if statusCode == code {
			return true
		}
	}
	return false
}

// calculateDelay calculates the delay before the next retry attempt
func calculateDelay(response *http.Response, attempt int, config *RetryConfig) time.Duration {
	// First, try to parse Retry-After header if we have a response
	if response != nil {
		if retryAfter := parseRetryAfter(response.Header.Get("Retry-After")); retryAfter > 0 {
			// Respect the server's retry-after header, but cap it at MaxDelay
			if retryAfter > config.MaxDelay {
				return config.MaxDelay
			}
			return retryAfter
		}
	}

	// Fall back to exponential backoff
	delay := time.Duration(float64(config.BaseDelay) * math.Pow(2, float64(attempt)))
	if delay > config.MaxDelay {
		delay = config.MaxDelay
	}
	return delay
}

// parseRetryAfter parses the Retry-After header value
// It supports both delay-seconds (integer) and HTTP-date formats
func parseRetryAfter(retryAfter string) time.Duration {
	if retryAfter == "" {
		return 0
	}

	// Try parsing as delay-seconds (integer)
	if seconds, err := strconv.Atoi(strings.TrimSpace(retryAfter)); err == nil && seconds > 0 {
		return time.Duration(seconds) * time.Second
	}

	// Try parsing as HTTP-date
	if httpDate, err := time.Parse(time.RFC1123, retryAfter); err == nil {
		delay := time.Until(httpDate)
		if delay > 0 {
			return delay
		}
	}

	// Alternative HTTP-date format
	if httpDate, err := time.Parse(time.RFC1123Z, retryAfter); err == nil {
		delay := time.Until(httpDate)
		if delay > 0 {
			return delay
		}
	}

	return 0
}

// RetryableAPICall wraps an API call with retry logic
// This is a convenience function for wrapping generated API client calls
func RetryableAPICall[T any](ctx context.Context, config *RetryConfig, apiCall func() (T, *http.Response, error)) (T, *http.Response, error) {
	var result T
	var response *http.Response
	var err error

	retryFunc := func() (*http.Response, error) {
		result, response, err = apiCall()
		return response, err
	}

	finalResponse, finalErr := WithRetry(ctx, config, retryFunc)
	return result, finalResponse, finalErr
}
