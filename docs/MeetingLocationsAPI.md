# \MeetingLocationsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMeetingLocation**](MeetingLocationsAPI.md#CreateMeetingLocation) | **Post** /meeting_locations | Create a meeting location.
[**DeleteMeetingLocationById**](MeetingLocationsAPI.md#DeleteMeetingLocationById) | **Delete** /meeting_locations/{id} | Delete a course location.
[**GetMeetingLocationById**](MeetingLocationsAPI.md#GetMeetingLocationById) | **Get** /meeting_locations/{id} | Get an meeting location
[**GetMeetingLocations**](MeetingLocationsAPI.md#GetMeetingLocations) | **Get** /meeting_locations | Get all meeting location records
[**UpdateMeetingLocationById**](MeetingLocationsAPI.md#UpdateMeetingLocationById) | **Patch** /meeting_locations/{id} | Update a meeting location.



## CreateMeetingLocation

> MeetingLocationWithIncludes CreateMeetingLocation(ctx).CreateMeetingLocationRequest(createMeetingLocationRequest).Execute()

Create a meeting location.

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
	createMeetingLocationRequest := *openapiclient.NewCreateMeetingLocationRequest("Name_example", int32(123)) // CreateMeetingLocationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetingLocationsAPI.CreateMeetingLocation(context.Background()).CreateMeetingLocationRequest(createMeetingLocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingLocationsAPI.CreateMeetingLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMeetingLocation`: MeetingLocationWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `MeetingLocationsAPI.CreateMeetingLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMeetingLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMeetingLocationRequest** | [**CreateMeetingLocationRequest**](CreateMeetingLocationRequest.md) |  | 

### Return type

[**MeetingLocationWithIncludes**](MeetingLocationWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMeetingLocationById

> DeleteMeetingLocationById(ctx, id).Execute()

Delete a course location.

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
	r, err := apiClient.MeetingLocationsAPI.DeleteMeetingLocationById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingLocationsAPI.DeleteMeetingLocationById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMeetingLocationByIdRequest struct via the builder pattern


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


## GetMeetingLocationById

> MeetingLocationWithIncludes GetMeetingLocationById(ctx, id).Execute()

Get an meeting location

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
	resp, r, err := apiClient.MeetingLocationsAPI.GetMeetingLocationById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingLocationsAPI.GetMeetingLocationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeetingLocationById`: MeetingLocationWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `MeetingLocationsAPI.GetMeetingLocationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeetingLocationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MeetingLocationWithIncludes**](MeetingLocationWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeetingLocations

> []MeetingLocationWithIncludes GetMeetingLocations(ctx).Cursor(cursor).PerPage(perPage).Execute()

Get all meeting location records

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
	resp, r, err := apiClient.MeetingLocationsAPI.GetMeetingLocations(context.Background()).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingLocationsAPI.GetMeetingLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeetingLocations`: []MeetingLocationWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `MeetingLocationsAPI.GetMeetingLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMeetingLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]MeetingLocationWithIncludes**](MeetingLocationWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMeetingLocationById

> MeetingLocationWithIncludes UpdateMeetingLocationById(ctx, id).UpdateMeetingLocationByIdRequest(updateMeetingLocationByIdRequest).Execute()

Update a meeting location.

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
	updateMeetingLocationByIdRequest := *openapiclient.NewUpdateMeetingLocationByIdRequest() // UpdateMeetingLocationByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetingLocationsAPI.UpdateMeetingLocationById(context.Background(), id).UpdateMeetingLocationByIdRequest(updateMeetingLocationByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingLocationsAPI.UpdateMeetingLocationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMeetingLocationById`: MeetingLocationWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `MeetingLocationsAPI.UpdateMeetingLocationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMeetingLocationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMeetingLocationByIdRequest** | [**UpdateMeetingLocationByIdRequest**](UpdateMeetingLocationByIdRequest.md) |  | 

### Return type

[**MeetingLocationWithIncludes**](MeetingLocationWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

