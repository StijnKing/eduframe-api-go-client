# \TeacherEnrollmentsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeacherEnrollment**](TeacherEnrollmentsAPI.md#CreateTeacherEnrollment) | **Post** /teacher_enrollments | Enroll a teacher to a planned_course.
[**CreateTeacherEnrollmentByPlannedCourseId**](TeacherEnrollmentsAPI.md#CreateTeacherEnrollmentByPlannedCourseId) | **Post** /planned_courses/{planned_course_id}/teacher_enrollments | Enroll a teacher to the given planned course.
[**DeleteTeacherEnrollmentById**](TeacherEnrollmentsAPI.md#DeleteTeacherEnrollmentById) | **Delete** /teacher_enrollments/{id} | Delete a teacher enrollment.
[**GetTeacherEnrollments**](TeacherEnrollmentsAPI.md#GetTeacherEnrollments) | **Get** /teacher_enrollments | Get all teacher enrollments.
[**GetTeacherEnrollmentsByPlannedCourseId**](TeacherEnrollmentsAPI.md#GetTeacherEnrollmentsByPlannedCourseId) | **Get** /planned_courses/{planned_course_id}/teacher_enrollments | Get all teacher enrollments for given planned course.
[**UpdateTeacherEnrollmentById**](TeacherEnrollmentsAPI.md#UpdateTeacherEnrollmentById) | **Patch** /teacher_enrollments/{id} | Update a teacher enrollment.



## CreateTeacherEnrollment

> TeacherEnrollment CreateTeacherEnrollment(ctx).CreateTeacherEnrollmentRequest(createTeacherEnrollmentRequest).Execute()

Enroll a teacher to a planned_course.

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
	createTeacherEnrollmentRequest := *openapiclient.NewCreateTeacherEnrollmentRequest(int32(123), int32(123)) // CreateTeacherEnrollmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeacherEnrollmentsAPI.CreateTeacherEnrollment(context.Background()).CreateTeacherEnrollmentRequest(createTeacherEnrollmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeacherEnrollmentsAPI.CreateTeacherEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeacherEnrollment`: TeacherEnrollment
	fmt.Fprintf(os.Stdout, "Response from `TeacherEnrollmentsAPI.CreateTeacherEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeacherEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeacherEnrollmentRequest** | [**CreateTeacherEnrollmentRequest**](CreateTeacherEnrollmentRequest.md) |  | 

### Return type

[**TeacherEnrollment**](TeacherEnrollment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeacherEnrollmentByPlannedCourseId

> TeacherEnrollment CreateTeacherEnrollmentByPlannedCourseId(ctx, plannedCourseId).CreateTeacherEnrollmentByPlannedCourseIdRequest(createTeacherEnrollmentByPlannedCourseIdRequest).Execute()

Enroll a teacher to the given planned course.

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
	plannedCourseId := int32(56) // int32 | Filter results on planned_course_id
	createTeacherEnrollmentByPlannedCourseIdRequest := *openapiclient.NewCreateTeacherEnrollmentByPlannedCourseIdRequest(int32(123)) // CreateTeacherEnrollmentByPlannedCourseIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeacherEnrollmentsAPI.CreateTeacherEnrollmentByPlannedCourseId(context.Background(), plannedCourseId).CreateTeacherEnrollmentByPlannedCourseIdRequest(createTeacherEnrollmentByPlannedCourseIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeacherEnrollmentsAPI.CreateTeacherEnrollmentByPlannedCourseId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeacherEnrollmentByPlannedCourseId`: TeacherEnrollment
	fmt.Fprintf(os.Stdout, "Response from `TeacherEnrollmentsAPI.CreateTeacherEnrollmentByPlannedCourseId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannedCourseId** | **int32** | Filter results on planned_course_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeacherEnrollmentByPlannedCourseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTeacherEnrollmentByPlannedCourseIdRequest** | [**CreateTeacherEnrollmentByPlannedCourseIdRequest**](CreateTeacherEnrollmentByPlannedCourseIdRequest.md) |  | 

### Return type

[**TeacherEnrollment**](TeacherEnrollment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeacherEnrollmentById

> DeleteTeacherEnrollmentById(ctx, id).Execute()

Delete a teacher enrollment.

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
	r, err := apiClient.TeacherEnrollmentsAPI.DeleteTeacherEnrollmentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeacherEnrollmentsAPI.DeleteTeacherEnrollmentById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTeacherEnrollmentByIdRequest struct via the builder pattern


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


## GetTeacherEnrollments

> []TeacherEnrollment GetTeacherEnrollments(ctx).PlannedCourseId(plannedCourseId).Cursor(cursor).PerPage(perPage).Execute()

Get all teacher enrollments.

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
	plannedCourseId := int32(56) // int32 | Filter results on planned_course_id (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeacherEnrollmentsAPI.GetTeacherEnrollments(context.Background()).PlannedCourseId(plannedCourseId).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeacherEnrollmentsAPI.GetTeacherEnrollments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeacherEnrollments`: []TeacherEnrollment
	fmt.Fprintf(os.Stdout, "Response from `TeacherEnrollmentsAPI.GetTeacherEnrollments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeacherEnrollmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plannedCourseId** | **int32** | Filter results on planned_course_id | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]TeacherEnrollment**](TeacherEnrollment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeacherEnrollmentsByPlannedCourseId

> []TeacherEnrollment GetTeacherEnrollmentsByPlannedCourseId(ctx, plannedCourseId).Execute()

Get all teacher enrollments for given planned course.

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
	plannedCourseId := int32(56) // int32 | Filter results on planned_course_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeacherEnrollmentsAPI.GetTeacherEnrollmentsByPlannedCourseId(context.Background(), plannedCourseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeacherEnrollmentsAPI.GetTeacherEnrollmentsByPlannedCourseId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeacherEnrollmentsByPlannedCourseId`: []TeacherEnrollment
	fmt.Fprintf(os.Stdout, "Response from `TeacherEnrollmentsAPI.GetTeacherEnrollmentsByPlannedCourseId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannedCourseId** | **int32** | Filter results on planned_course_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeacherEnrollmentsByPlannedCourseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TeacherEnrollment**](TeacherEnrollment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeacherEnrollmentById

> TeacherEnrollment UpdateTeacherEnrollmentById(ctx, id).UpdateTeacherEnrollmentByIdRequest(updateTeacherEnrollmentByIdRequest).Execute()

Update a teacher enrollment.

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
	updateTeacherEnrollmentByIdRequest := *openapiclient.NewUpdateTeacherEnrollmentByIdRequest() // UpdateTeacherEnrollmentByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeacherEnrollmentsAPI.UpdateTeacherEnrollmentById(context.Background(), id).UpdateTeacherEnrollmentByIdRequest(updateTeacherEnrollmentByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeacherEnrollmentsAPI.UpdateTeacherEnrollmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeacherEnrollmentById`: TeacherEnrollment
	fmt.Fprintf(os.Stdout, "Response from `TeacherEnrollmentsAPI.UpdateTeacherEnrollmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeacherEnrollmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTeacherEnrollmentByIdRequest** | [**UpdateTeacherEnrollmentByIdRequest**](UpdateTeacherEnrollmentByIdRequest.md) |  | 

### Return type

[**TeacherEnrollment**](TeacherEnrollment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

