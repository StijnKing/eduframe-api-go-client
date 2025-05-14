# \PlanningEventsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePlanningEventById**](PlanningEventsAPI.md#DeletePlanningEventById) | **Delete** /planning/events/{id} | Delete a planning event.
[**GetPlanningEventById**](PlanningEventsAPI.md#GetPlanningEventById) | **Get** /planning/events/{id} | Get an planning event record
[**GetPlanningEvents**](PlanningEventsAPI.md#GetPlanningEvents) | **Get** /planning/events | Get all planning event records



## DeletePlanningEventById

> DeletePlanningEventById(ctx, id).Execute()

Delete a planning event.

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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlanningEventsAPI.DeletePlanningEventById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanningEventsAPI.DeletePlanningEventById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanningEventByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanningEventById

> PlanningEventWithIncludes GetPlanningEventById(ctx, id).Execute()

Get an planning event record

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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanningEventsAPI.GetPlanningEventById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanningEventsAPI.GetPlanningEventById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanningEventById`: PlanningEventWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `PlanningEventsAPI.GetPlanningEventById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanningEventByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlanningEventWithIncludes**](PlanningEventWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanningEvents

> []PlanningEventWithIncludes GetPlanningEvents(ctx).Cursor(cursor).PerPage(perPage).Execute()

Get all planning event records

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
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanningEventsAPI.GetPlanningEvents(context.Background()).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanningEventsAPI.GetPlanningEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanningEvents`: []PlanningEventWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `PlanningEventsAPI.GetPlanningEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanningEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]PlanningEventWithIncludes**](PlanningEventWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

