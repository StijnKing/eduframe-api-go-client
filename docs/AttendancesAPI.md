# \AttendancesAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttendances**](AttendancesAPI.md#GetAttendances) | **Get** /attendances | Get all attendance records
[**SetAttendance**](AttendancesAPI.md#SetAttendance) | **Post** /attendances/upsert | Set an attendance.



## GetAttendances

> []Attendance GetAttendances(ctx).MeetingId(meetingId).Cursor(cursor).PerPage(perPage).Execute()

Get all attendance records

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
	meetingId := int32(56) // int32 | Filter attendances on meeting_id (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttendancesAPI.GetAttendances(context.Background()).MeetingId(meetingId).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttendancesAPI.GetAttendances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttendances`: []Attendance
	fmt.Fprintf(os.Stdout, "Response from `AttendancesAPI.GetAttendances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttendancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **meetingId** | **int32** | Filter attendances on meeting_id | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]Attendance**](Attendance.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAttendance

> Attendance SetAttendance(ctx).SetAttendanceRequest(setAttendanceRequest).Execute()

Set an attendance.

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
	setAttendanceRequest := *openapiclient.NewSetAttendanceRequest(int32(123), int32(123)) // SetAttendanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttendancesAPI.SetAttendance(context.Background()).SetAttendanceRequest(setAttendanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttendancesAPI.SetAttendance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAttendance`: Attendance
	fmt.Fprintf(os.Stdout, "Response from `AttendancesAPI.SetAttendance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAttendanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setAttendanceRequest** | [**SetAttendanceRequest**](SetAttendanceRequest.md) |  | 

### Return type

[**Attendance**](Attendance.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

