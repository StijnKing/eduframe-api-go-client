# \WebhookNotificationsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhookNotificationsByWebhookId**](WebhookNotificationsAPI.md#GetWebhookNotificationsByWebhookId) | **Get** /webhooks/{webhook_id}/notifications | Get the notifications for a specific webhook
[**GetWebhookNotificationsFailed**](WebhookNotificationsAPI.md#GetWebhookNotificationsFailed) | **Get** /webhooks/{webhook_id}/notifications/failed | Get the failed webhook notifications



## GetWebhookNotificationsByWebhookId

> []WebhookNotificationWithIncludes GetWebhookNotificationsByWebhookId(ctx, webhookId).Execute()

Get the notifications for a specific webhook

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	webhookId := "webhookId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookNotificationsAPI.GetWebhookNotificationsByWebhookId(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookNotificationsAPI.GetWebhookNotificationsByWebhookId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookNotificationsByWebhookId`: []WebhookNotificationWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `WebhookNotificationsAPI.GetWebhookNotificationsByWebhookId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookNotificationsByWebhookIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]WebhookNotificationWithIncludes**](WebhookNotificationWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookNotificationsFailed

> []WebhookNotificationFailed GetWebhookNotificationsFailed(ctx, webhookId).Start(start).End(end).Cursor(cursor).PerPage(perPage).Execute()

Get the failed webhook notifications

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	webhookId := "webhookId_example" // string | 
	start := time.Now() // time.Time | Only show failed notifications created after this date and time (optional)
	end := time.Now() // time.Time | Only show failed notifications starting before this date and time (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookNotificationsAPI.GetWebhookNotificationsFailed(context.Background(), webhookId).Start(start).End(end).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookNotificationsAPI.GetWebhookNotificationsFailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookNotificationsFailed`: []WebhookNotificationFailed
	fmt.Fprintf(os.Stdout, "Response from `WebhookNotificationsAPI.GetWebhookNotificationsFailed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookNotificationsFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **time.Time** | Only show failed notifications created after this date and time | 
 **end** | **time.Time** | Only show failed notifications starting before this date and time | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]WebhookNotificationFailed**](WebhookNotificationFailed.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

