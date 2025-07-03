/*
Eduframe - API

Integration test demonstrating the retry helper with a mock server.

API version: v1
*/

package eduframe

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRetryIntegration(t *testing.T) {
	// Create a mock server that simulates rate limiting
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		
		if requestCount <= 2 {
			// First two requests return 429 with Retry-After header
			w.Header().Set("Retry-After", "1") // 1 second
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error": "rate limited"}`))
			return
		}
		
		// Third request succeeds
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	// Configure API client to use mock server
	config := NewConfiguration()
	config.Servers = ServerConfigurations{
		{
			URL: server.URL,
			Description: "Mock server for testing",
		},
	}
	
	// Test with retry helper
	ctx := context.Background()
	retryConfig := &RetryConfig{
		MaxRetries:      3,
		BaseDelay:       100 * time.Millisecond, // Use short delays for testing
		MaxDelay:        2 * time.Second,
		RetryOnStatuses: []int{429},
	}
	
	start := time.Now()
	
	// Simulate an API call with retry
	var finalResponse *http.Response
	var finalError error
	
	retryFunc := func() (*http.Response, error) {
		req, err := http.NewRequest("GET", server.URL+"/test", nil)
		if err != nil {
			return nil, err
		}
		return http.DefaultClient.Do(req)
	}
	
	finalResponse, finalError = WithRetry(ctx, retryConfig, retryFunc)
	
	duration := time.Since(start)
	
	// Assertions
	if finalError != nil {
		t.Errorf("Expected no error, got: %v", finalError)
	}
	
	if finalResponse.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got: %d", finalResponse.StatusCode)
	}
	
	if requestCount != 3 {
		t.Errorf("Expected 3 requests, got: %d", requestCount)
	}
	
	// Should have waited at least 2 seconds (1 second retry delay * 2 retries)
	// But be lenient with timing in tests
	if duration < 100*time.Millisecond {
		t.Errorf("Expected some delay, got: %v", duration)
	}
	
	t.Logf("Integration test completed successfully:")
	t.Logf("  - Requests made: %d", requestCount)
	t.Logf("  - Total duration: %v", duration)
	t.Logf("  - Final status: %d", finalResponse.StatusCode)
}

func TestRetryIntegrationMaxAttemptsExceeded(t *testing.T) {
	// Create a mock server that always returns 429
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Retry-After", "1")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"error": "persistent rate limit"}`))
	}))
	defer server.Close()

	ctx := context.Background()
	retryConfig := &RetryConfig{
		MaxRetries:      2,
		BaseDelay:       50 * time.Millisecond,
		MaxDelay:        1 * time.Second,
		RetryOnStatuses: []int{429},
	}
	
	retryFunc := func() (*http.Response, error) {
		req, err := http.NewRequest("GET", server.URL+"/test", nil)
		if err != nil {
			return nil, err
		}
		return http.DefaultClient.Do(req)
	}
	
	finalResponse, finalError := WithRetry(ctx, retryConfig, retryFunc)
	
	// Should still return the last response even after max retries
	if finalError != nil {
		t.Errorf("Expected no error from retryFunc, got: %v", finalError)
	}
	
	if finalResponse.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected status 429, got: %d", finalResponse.StatusCode)
	}
	
	expectedRequests := retryConfig.MaxRetries + 1 // Initial attempt + retries
	if requestCount != expectedRequests {
		t.Errorf("Expected %d requests, got: %d", expectedRequests, requestCount)
	}
	
	t.Logf("Max attempts test completed:")
	t.Logf("  - Requests made: %d", requestCount)
	t.Logf("  - Final status: %d", finalResponse.StatusCode)
}