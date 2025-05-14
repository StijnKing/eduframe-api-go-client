# \CoursesAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCourse**](CoursesAPI.md#CreateCourse) | **Post** /courses | Create a course.
[**GetCourseById**](CoursesAPI.md#GetCourseById) | **Get** /courses/{id} | Get a course record
[**GetCourses**](CoursesAPI.md#GetCourses) | **Get** /courses | Get all course records
[**UpdateCourseById**](CoursesAPI.md#UpdateCourseById) | **Patch** /courses/{id} | Update a course.



## CreateCourse

> CourseWithIncludes CreateCourse(ctx).CoursePayload(coursePayload).Execute()

Create a course.

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
	coursePayload := *openapiclient.NewCoursePayload(int32(123), "Name_example", "Code_example") // CoursePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.CreateCourse(context.Background()).CoursePayload(coursePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.CreateCourse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCourse`: CourseWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.CreateCourse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCourseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coursePayload** | [**CoursePayload**](CoursePayload.md) |  | 

### Return type

[**CourseWithIncludes**](CourseWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseById

> CourseWithIncludes GetCourseById(ctx, id).Execute()

Get a course record

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
	resp, r, err := apiClient.CoursesAPI.GetCourseById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseById`: CourseWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CourseWithIncludes**](CourseWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourses

> []CourseWithIncludes GetCourses(ctx).Published(published).Cursor(cursor).PerPage(perPage).Execute()

Get all course records

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
	published := "published_example" // string | Show only published courses (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourses(context.Background()).Published(published).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourses`: []CourseWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCoursesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **published** | **string** | Show only published courses | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]CourseWithIncludes**](CourseWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCourseById

> CourseWithIncludes UpdateCourseById(ctx, id).CoursePatchPayload(coursePatchPayload).Execute()

Update a course.

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
	coursePatchPayload := *openapiclient.NewCoursePatchPayload() // CoursePatchPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.UpdateCourseById(context.Background(), id).CoursePatchPayload(coursePatchPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.UpdateCourseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCourseById`: CourseWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.UpdateCourseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCourseByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **coursePatchPayload** | [**CoursePatchPayload**](CoursePatchPayload.md) |  | 

### Return type

[**CourseWithIncludes**](CourseWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

