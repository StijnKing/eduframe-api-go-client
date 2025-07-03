# Retry Helper for Rate Limiting

The Eduframe API Go client includes a built-in retry helper to handle rate limiting (HTTP 429 responses) and other transient errors gracefully.

## Features

- **Automatic Rate Limit Handling**: Detects HTTP 429 responses and respects `Retry-After` headers
- **Exponential Backoff**: Falls back to exponential backoff when no `Retry-After` header is present
- **Configurable**: Customize retry behavior including max attempts, delays, and which status codes to retry
- **Context-Aware**: Supports Go context for cancellation and timeouts
- **Non-Intrusive**: Works as a wrapper around existing API client calls

## Quick Start

### Basic Usage with Default Settings

```go
import (
    "context"
    eduframe "github.com/StijnKing/eduframe-api-go-client"
)

// Use default retry configuration
ctx := context.Background()
accounts, resp, err := eduframe.RetryableAPICall(ctx, nil, func() ([]eduframe.AccountWithIncludes, *http.Response, error) {
    return apiClient.AccountsAPI.GetAccounts(ctx).Execute()
})
```

### Custom Retry Configuration

```go
retryConfig := &eduframe.RetryConfig{
    MaxRetries:      3,                         // Maximum retry attempts
    BaseDelay:       1 * time.Second,           // Base delay for exponential backoff
    MaxDelay:        60 * time.Second,          // Maximum delay between retries
    RetryOnStatuses: []int{429, 500, 502, 503, 504}, // Status codes to retry
}

accounts, resp, err := eduframe.RetryableAPICall(ctx, retryConfig, func() ([]eduframe.AccountWithIncludes, *http.Response, error) {
    return apiClient.AccountsAPI.GetAccounts(ctx).Execute()
})
```

## Configuration Options

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `MaxRetries` | `int` | `3` | Maximum number of retry attempts |
| `BaseDelay` | `time.Duration` | `1s` | Base delay for exponential backoff |
| `MaxDelay` | `time.Duration` | `60s` | Maximum delay between retries |
| `RetryOnStatuses` | `[]int` | `[429, 500, 502, 503, 504]` | HTTP status codes that trigger retries |

## How It Works

1. **Rate Limit Detection**: The helper automatically detects HTTP 429 (Too Many Requests) responses
2. **Retry-After Header**: When present, the helper respects the `Retry-After` header for wait time
3. **Exponential Backoff**: If no `Retry-After` header is provided, uses exponential backoff starting from `BaseDelay`
4. **Max Delay Cap**: All delays are capped at `MaxDelay` to prevent excessively long waits
5. **Context Cancellation**: Respects context cancellation and timeouts throughout the retry process

## Advanced Usage

### Lower-Level Control with WithRetry

For more control over the retry process, you can use the `WithRetry` function directly:

```go
retryFunc := func() (*http.Response, error) {
    _, resp, err := apiClient.AccountsAPI.GetAccounts(ctx).Execute()
    return resp, err
}

finalResp, finalErr := eduframe.WithRetry(ctx, retryConfig, retryFunc)
```

### Rate Limit Only Retries

To only retry on rate limits and not server errors:

```go
retryConfig := &eduframe.RetryConfig{
    MaxRetries:      5,
    BaseDelay:       500 * time.Millisecond,
    MaxDelay:        30 * time.Second,
    RetryOnStatuses: []int{429}, // Only retry rate limits
}
```

## Error Handling

The retry helper preserves the original error and response from the final attempt. This means:

- If all retries are exhausted, you get the last error/response
- The helper doesn't mask or transform the original API errors
- HTTP response details (status codes, headers, body) are preserved

## Best Practices

1. **Use Appropriate Timeouts**: Set context timeouts to prevent indefinite waiting
2. **Monitor Retry Patterns**: Log retry attempts to understand API usage patterns
3. **Respect Rate Limits**: The helper automatically respects `Retry-After` headers
4. **Configure for Your Use Case**: Adjust retry settings based on your application's requirements

## Examples


See `retry_examples.go` for comprehensive usage examples.
