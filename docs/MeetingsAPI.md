# \MeetingsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMeetingById**](MeetingsAPI.md#DeleteMeetingById) | **Delete** /meetings/{id} | Delete a meeting.
[**GetMeetingById**](MeetingsAPI.md#GetMeetingById) | **Get** /meetings/{id} | Get an meeting record
[**GetMeetings**](MeetingsAPI.md#GetMeetings) | **Get** /meetings | Get all meeting records
[**GetMeetingsByPlannedCourseId**](MeetingsAPI.md#GetMeetingsByPlannedCourseId) | **Get** /planned_courses/{planned_course_id}/meetings | Get all meeting records of a planned course



## DeleteMeetingById

> DeleteMeetingById(ctx, id).Execute()

Delete a meeting.

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
	r, err := apiClient.MeetingsAPI.DeleteMeetingById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.DeleteMeetingById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMeetingByIdRequest struct via the builder pattern


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


## GetMeetingById

> MeetingWithIncludes GetMeetingById(ctx, id).Execute()

Get an meeting record

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
	resp, r, err := apiClient.MeetingsAPI.GetMeetingById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.GetMeetingById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeetingById`: MeetingWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `MeetingsAPI.GetMeetingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeetingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MeetingWithIncludes**](MeetingWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeetings

> []MeetingWithIncludes GetMeetings(ctx).PlannedCourseId(plannedCourseId).Start(start).End(end).Cursor(cursor).PerPage(perPage).Execute()

Get all meeting records

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
	plannedCourseId := int32(56) // int32 | Filter results on planned_course_id (optional)
	start := time.Now() // time.Time | Only show meetings ending after this date and time (optional)
	end := time.Now() // time.Time | Only show meetings starting before this date and time (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetingsAPI.GetMeetings(context.Background()).PlannedCourseId(plannedCourseId).Start(start).End(end).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.GetMeetings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeetings`: []MeetingWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `MeetingsAPI.GetMeetings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMeetingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plannedCourseId** | **int32** | Filter results on planned_course_id | 
 **start** | **time.Time** | Only show meetings ending after this date and time | 
 **end** | **time.Time** | Only show meetings starting before this date and time | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]MeetingWithIncludes**](MeetingWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeetingsByPlannedCourseId

> []MeetingWithIncludes GetMeetingsByPlannedCourseId(ctx, plannedCourseId).Sort(sort).Execute()

Get all meeting records of a planned course

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
	plannedCourseId := "plannedCourseId_example" // string | Filter results on planned_course_id
	sort := []string{"Sort_example"} // []string | Sort the results. Can change order by using `<sort_by>:<direction>` where `<direction>` is either `asc` or `desc` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetingsAPI.GetMeetingsByPlannedCourseId(context.Background(), plannedCourseId).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.GetMeetingsByPlannedCourseId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeetingsByPlannedCourseId`: []MeetingWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `MeetingsAPI.GetMeetingsByPlannedCourseId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannedCourseId** | **string** | Filter results on planned_course_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeetingsByPlannedCourseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **[]string** | Sort the results. Can change order by using &#x60;&lt;sort_by&gt;:&lt;direction&gt;&#x60; where &#x60;&lt;direction&gt;&#x60; is either &#x60;asc&#x60; or &#x60;desc&#x60; | 

### Return type

[**[]MeetingWithIncludes**](MeetingWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

