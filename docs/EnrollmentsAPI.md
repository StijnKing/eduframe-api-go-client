# \EnrollmentsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelEnrollmentById**](EnrollmentsAPI.md#CancelEnrollmentById) | **Put** /enrollments/{id}/cancel | Cancel an enrollment
[**GetEnrollmentById**](EnrollmentsAPI.md#GetEnrollmentById) | **Get** /enrollments/{id} | Get an enrollment record
[**GetEnrollments**](EnrollmentsAPI.md#GetEnrollments) | **Get** /enrollments | Get all enrollment records
[**UpdateEnrollmentById**](EnrollmentsAPI.md#UpdateEnrollmentById) | **Patch** /enrollments/{id} | Update an enrollment



## CancelEnrollmentById

> EnrollmentWithIncludes CancelEnrollmentById(ctx, id).Execute()

Cancel an enrollment

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
	resp, r, err := apiClient.EnrollmentsAPI.CancelEnrollmentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentsAPI.CancelEnrollmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelEnrollmentById`: EnrollmentWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentsAPI.CancelEnrollmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelEnrollmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrollmentWithIncludes**](EnrollmentWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentById

> EnrollmentWithIncludes GetEnrollmentById(ctx, id).Execute()

Get an enrollment record

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
	resp, r, err := apiClient.EnrollmentsAPI.GetEnrollmentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentsAPI.GetEnrollmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnrollmentById`: EnrollmentWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentsAPI.GetEnrollmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrollmentWithIncludes**](EnrollmentWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollments

> []EnrollmentWithIncludes GetEnrollments(ctx).StudentId(studentId).PlannedCourseId(plannedCourseId).Status(status).WithCanceled(withCanceled).Cursor(cursor).PerPage(perPage).Execute()

Get all enrollment records

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
	studentId := int32(56) // int32 | Filter results on student_id (optional)
	plannedCourseId := int32(56) // int32 | Filter results on planned_course_id (optional)
	status := []string{"Status_example"} // []string | Filter results on status (optional)
	withCanceled := true // bool | Filter results based on whether they include a canceled status or not (optional) (default to true)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentsAPI.GetEnrollments(context.Background()).StudentId(studentId).PlannedCourseId(plannedCourseId).Status(status).WithCanceled(withCanceled).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentsAPI.GetEnrollments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnrollments`: []EnrollmentWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentsAPI.GetEnrollments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **studentId** | **int32** | Filter results on student_id | 
 **plannedCourseId** | **int32** | Filter results on planned_course_id | 
 **status** | **[]string** | Filter results on status | 
 **withCanceled** | **bool** | Filter results based on whether they include a canceled status or not | [default to true]
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]EnrollmentWithIncludes**](EnrollmentWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnrollmentById

> EnrollmentWithIncludes UpdateEnrollmentById(ctx, id).UpdateEnrollmentByIdRequest(updateEnrollmentByIdRequest).Execute()

Update an enrollment

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
	updateEnrollmentByIdRequest := *openapiclient.NewUpdateEnrollmentByIdRequest() // UpdateEnrollmentByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentsAPI.UpdateEnrollmentById(context.Background(), id).UpdateEnrollmentByIdRequest(updateEnrollmentByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentsAPI.UpdateEnrollmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEnrollmentById`: EnrollmentWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentsAPI.UpdateEnrollmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnrollmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEnrollmentByIdRequest** | [**UpdateEnrollmentByIdRequest**](UpdateEnrollmentByIdRequest.md) |  | 

### Return type

[**EnrollmentWithIncludes**](EnrollmentWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

