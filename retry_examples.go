/*
Eduframe - API

Example usage of the retry helper for handling rate limits.

API version: v1
*/

package eduframe

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// ExampleWithRetryHelper demonstrates how to use the retry helper with the Eduframe API client
func ExampleWithRetryHelper() {
	// Create the API client configuration
	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)

	// Set up context with API key for authentication
	ctx := context.WithValue(context.Background(), ContextAPIKeys, map[string]APIKey{
		"Bearer": {Key: os.Getenv("EDUFRAME_API_KEY")},
	})

	// Configure retry settings (optional - defaults are usually fine)
	retryConfig := &RetryConfig{
		MaxRetries:      3,                         // Maximum number of retry attempts
		BaseDelay:       1 * time.Second,           // Base delay for exponential backoff
		MaxDelay:        60 * time.Second,          // Maximum delay between retries
		RetryOnStatuses: []int{429, 500, 502, 503, 504}, // HTTP status codes to retry on
	}

	// Example 1: Using RetryableAPICall wrapper (recommended)
	accounts, _, err := RetryableAPICall(ctx, retryConfig, func() ([]AccountWithIncludes, *http.Response, error) {
		return apiClient.AccountsAPI.GetAccounts(ctx).Execute()
	})

	if err != nil {
		log.Printf("Error fetching accounts with retry: %v", err)
		return
	}

	fmt.Printf("Successfully fetched %d accounts\n", len(accounts))

	// Example 2: Using WithRetry for lower-level control
	retryFunc := func() (*http.Response, error) {
		// This example shows how to wrap any function that returns (*http.Response, error)
		_, resp, err := apiClient.AccountsAPI.GetAccounts(ctx).Execute()
		return resp, err
	}

	finalResp, finalErr := WithRetry(ctx, retryConfig, retryFunc)
	if finalErr != nil {
		log.Printf("Error with manual retry: %v", finalErr)
		return
	}

	fmt.Printf("Final response status: %d\n", finalResp.StatusCode)
}

// ExampleWithDefaultRetry demonstrates using the retry helper with default settings
func ExampleWithDefaultRetry() {
	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)

	ctx := context.WithValue(context.Background(), ContextAPIKeys, map[string]APIKey{
		"Bearer": {Key: os.Getenv("EDUFRAME_API_KEY")},
	})

	// Use default retry configuration (nil config will use defaults)
	accounts, _, err := RetryableAPICall(ctx, nil, func() ([]AccountWithIncludes, *http.Response, error) {
		return apiClient.AccountsAPI.GetAccounts(ctx).Execute()
	})

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Fetched %d accounts with default retry settings\n", len(accounts))
}

// ExampleCustomRetryStatuses shows how to configure which HTTP status codes trigger retries
func ExampleCustomRetryStatuses() {
	retryConfig := &RetryConfig{
		MaxRetries:      2,
		BaseDelay:       500 * time.Millisecond, // 500ms
		MaxDelay:        5 * time.Second,        // 5 seconds
		// Only retry on rate limits, not server errors
		RetryOnStatuses: []int{429},
	}

	// Use this config with any API call as shown in previous examples
	fmt.Printf("Custom retry config: %+v\n", retryConfig)
}