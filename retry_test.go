/*
Eduframe - API

Tests for retry helper functionality.

API version: v1
*/

package eduframe

import (
	"context"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestDefaultRetryConfig(t *testing.T) {
	config := DefaultRetryConfig()
	
	if config.MaxRetries != 3 {
		t.Errorf("Expected MaxRetries to be 3, got %d", config.MaxRetries)
	}
	
	if config.BaseDelay != time.Second {
		t.Errorf("Expected BaseDelay to be 1 second, got %v", config.BaseDelay)
	}
	
	if config.MaxDelay != 60*time.Second {
		t.Errorf("Expected MaxDelay to be 60 seconds, got %v", config.MaxDelay)
	}
	
	expectedStatuses := []int{429, 500, 502, 503, 504}
	if len(config.RetryOnStatuses) != len(expectedStatuses) {
		t.Errorf("Expected %d retry statuses, got %d", len(expectedStatuses), len(config.RetryOnStatuses))
	}
}

func TestShouldRetry(t *testing.T) {
	retryStatuses := []int{429, 500, 502, 503, 504}
	
	// Test status codes that should trigger retry
	for _, status := range retryStatuses {
		if !shouldRetry(status, retryStatuses) {
			t.Errorf("Expected status %d to trigger retry", status)
		}
	}
	
	// Test status codes that should not trigger retry
	nonRetryStatuses := []int{200, 201, 400, 401, 403, 404, 422}
	for _, status := range nonRetryStatuses {
		if shouldRetry(status, retryStatuses) {
			t.Errorf("Expected status %d to not trigger retry", status)
		}
	}
}

func TestParseRetryAfter(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected time.Duration
	}{
		{
			name:     "Empty value",
			value:    "",
			expected: 0,
		},
		{
			name:     "Delay seconds",
			value:    "30",
			expected: 30 * time.Second,
		},
		{
			name:     "Delay seconds with whitespace",
			value:    "  60  ",
			expected: 60 * time.Second,
		},
		{
			name:     "Zero delay",
			value:    "0",
			expected: 0,
		},
		{
			name:     "Invalid format",
			value:    "invalid",
			expected: 0,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseRetryAfter(tt.value)
			if result != tt.expected {
				t.Errorf("parseRetryAfter(%q) = %v, expected %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestCalculateDelay(t *testing.T) {
	config := &RetryConfig{
		BaseDelay: time.Second,
		MaxDelay:  10 * time.Second,
	}
	
	// Test exponential backoff without Retry-After header
	tests := []struct {
		attempt  int
		expected time.Duration
	}{
		{0, time.Second},
		{1, 2 * time.Second},
		{2, 4 * time.Second},
		{3, 8 * time.Second},
		{4, 10 * time.Second}, // Should be capped at MaxDelay
	}
	
	for _, tt := range tests {
		t.Run("attempt_"+strconv.Itoa(tt.attempt), func(t *testing.T) {
			delay := calculateDelay(nil, tt.attempt, config)
			if delay != tt.expected {
				t.Errorf("calculateDelay(attempt %d) = %v, expected %v", tt.attempt, delay, tt.expected)
			}
		})
	}
	
	// Test with Retry-After header
	t.Run("with_retry_after", func(t *testing.T) {
		resp := &http.Response{
			Header: http.Header{
				"Retry-After": []string{"5"},
			},
		}
		delay := calculateDelay(resp, 0, config)
		expected := 5 * time.Second
		if delay != expected {
			t.Errorf("calculateDelay with Retry-After = %v, expected %v", delay, expected)
		}
	})
	
	// Test Retry-After header capped at MaxDelay
	t.Run("retry_after_capped", func(t *testing.T) {
		resp := &http.Response{
			Header: http.Header{
				"Retry-After": []string{"30"},
			},
		}
		delay := calculateDelay(resp, 0, config)
		expected := config.MaxDelay
		if delay != expected {
			t.Errorf("calculateDelay with large Retry-After = %v, expected %v", delay, expected)
		}
	})
}

func TestWithRetrySuccess(t *testing.T) {
	ctx := context.Background()
	config := &RetryConfig{
		MaxRetries:      2,
		BaseDelay:       10 * time.Millisecond,
		MaxDelay:        100 * time.Millisecond,
		RetryOnStatuses: []int{429},
	}
	
	callCount := 0
	retryFunc := func() (*http.Response, error) {
		callCount++
		resp := &http.Response{StatusCode: http.StatusOK}
		return resp, nil
	}
	
	resp, err := WithRetry(ctx, config, retryFunc)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	
	if callCount != 1 {
		t.Errorf("Expected 1 call, got %d", callCount)
	}
}

func TestWithRetryRateLimit(t *testing.T) {
	ctx := context.Background()
	config := &RetryConfig{
		MaxRetries:      2,
		BaseDelay:       10 * time.Millisecond,
		MaxDelay:        100 * time.Millisecond,
		RetryOnStatuses: []int{429},
	}
	
	callCount := 0
	retryFunc := func() (*http.Response, error) {
		callCount++
		if callCount <= 2 {
			// Return 429 for first two calls with minimal retry delay for testing
			resp := &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{"Retry-After": []string{"0"}}, // Minimal delay
			}
			return resp, nil
		}
		// Return success on third call
		resp := &http.Response{StatusCode: http.StatusOK}
		return resp, nil
	}
	
	start := time.Now()
	resp, err := WithRetry(ctx, config, retryFunc)
	duration := time.Since(start)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	
	if callCount != 3 {
		t.Errorf("Expected 3 calls, got %d", callCount)
	}
	
	// Should have some delay but not too much for testing
	// Since we're using exponential backoff as fallback when Retry-After is 0
	expectedMaxDelay := 1 * time.Second // Should be reasonable for test
	if duration > expectedMaxDelay {
		t.Errorf("Expected delay less than %v, got %v", expectedMaxDelay, duration)
	}
}

func TestWithRetryMaxAttemptsExceeded(t *testing.T) {
	ctx := context.Background()
	config := &RetryConfig{
		MaxRetries:      1,
		BaseDelay:       10 * time.Millisecond,
		MaxDelay:        100 * time.Millisecond,
		RetryOnStatuses: []int{429},
	}
	
	callCount := 0
	retryFunc := func() (*http.Response, error) {
		callCount++
		resp := &http.Response{StatusCode: http.StatusTooManyRequests}
		return resp, nil
	}
	
	resp, err := WithRetry(ctx, config, retryFunc)
	
	if err != nil {
		t.Errorf("Expected no error from retryFunc, got %v", err)
	}
	
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected status 429, got %d", resp.StatusCode)
	}
	
	if callCount != 2 { // 1 initial + 1 retry
		t.Errorf("Expected 2 calls, got %d", callCount)
	}
}

func TestWithRetryContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	config := DefaultRetryConfig()
	
	callCount := 0
	retryFunc := func() (*http.Response, error) {
		callCount++
		if callCount == 1 {
			// Cancel context after first call
			cancel()
		}
		resp := &http.Response{StatusCode: http.StatusTooManyRequests}
		return resp, nil
	}
	
	_, err := WithRetry(ctx, config, retryFunc)
	
	if err != context.Canceled {
		t.Errorf("Expected context.Canceled error, got %v", err)
	}
	
	if callCount > 2 {
		t.Errorf("Expected at most 2 calls, got %d", callCount)
	}
}

func TestRetryableAPICall(t *testing.T) {
	ctx := context.Background()
	config := DefaultRetryConfig()
	
	// Mock API call that returns a string, response, and error
	apiCall := func() (string, *http.Response, error) {
		resp := &http.Response{StatusCode: http.StatusOK}
		return "success", resp, nil
	}
	
	result, resp, err := RetryableAPICall(ctx, config, apiCall)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if result != "success" {
		t.Errorf("Expected result 'success', got %q", result)
	}
	
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}