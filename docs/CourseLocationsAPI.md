# \CourseLocationsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCourseLocation**](CourseLocationsAPI.md#CreateCourseLocation) | **Post** /course_locations | Create a course location.
[**DeleteCourseLocationById**](CourseLocationsAPI.md#DeleteCourseLocationById) | **Delete** /course_locations/{id} | Delete a course location.
[**GetCourseLocationById**](CourseLocationsAPI.md#GetCourseLocationById) | **Get** /course_locations/{id} | Get a course location record
[**GetCourseLocations**](CourseLocationsAPI.md#GetCourseLocations) | **Get** /course_locations | Get all course location records
[**UpdateCourseLocationById**](CourseLocationsAPI.md#UpdateCourseLocationById) | **Patch** /course_locations/{id} | Update a course location.



## CreateCourseLocation

> CourseLocationWithIncludes CreateCourseLocation(ctx).CreateCourseLocationRequest(createCourseLocationRequest).Execute()

Create a course location.

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
	createCourseLocationRequest := *openapiclient.NewCreateCourseLocationRequest("Name_example") // CreateCourseLocationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseLocationsAPI.CreateCourseLocation(context.Background()).CreateCourseLocationRequest(createCourseLocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseLocationsAPI.CreateCourseLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCourseLocation`: CourseLocationWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CourseLocationsAPI.CreateCourseLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCourseLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCourseLocationRequest** | [**CreateCourseLocationRequest**](CreateCourseLocationRequest.md) |  | 

### Return type

[**CourseLocationWithIncludes**](CourseLocationWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCourseLocationById

> DeleteCourseLocationById(ctx, id).Execute()

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
	r, err := apiClient.CourseLocationsAPI.DeleteCourseLocationById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseLocationsAPI.DeleteCourseLocationById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCourseLocationByIdRequest struct via the builder pattern


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


## GetCourseLocationById

> CourseLocationWithIncludes GetCourseLocationById(ctx, id).Execute()

Get a course location record

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
	resp, r, err := apiClient.CourseLocationsAPI.GetCourseLocationById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseLocationsAPI.GetCourseLocationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseLocationById`: CourseLocationWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CourseLocationsAPI.GetCourseLocationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseLocationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CourseLocationWithIncludes**](CourseLocationWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseLocations

> []CourseLocationWithIncludes GetCourseLocations(ctx).Cursor(cursor).PerPage(perPage).Execute()

Get all course location records

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
	resp, r, err := apiClient.CourseLocationsAPI.GetCourseLocations(context.Background()).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseLocationsAPI.GetCourseLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseLocations`: []CourseLocationWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CourseLocationsAPI.GetCourseLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]CourseLocationWithIncludes**](CourseLocationWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCourseLocationById

> CourseLocationWithIncludes UpdateCourseLocationById(ctx, id).UpdateCourseLocationByIdRequest(updateCourseLocationByIdRequest).Execute()

Update a course location.

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
	updateCourseLocationByIdRequest := *openapiclient.NewUpdateCourseLocationByIdRequest() // UpdateCourseLocationByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseLocationsAPI.UpdateCourseLocationById(context.Background(), id).UpdateCourseLocationByIdRequest(updateCourseLocationByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseLocationsAPI.UpdateCourseLocationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCourseLocationById`: CourseLocationWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CourseLocationsAPI.UpdateCourseLocationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCourseLocationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCourseLocationByIdRequest** | [**UpdateCourseLocationByIdRequest**](UpdateCourseLocationByIdRequest.md) |  | 

### Return type

[**CourseLocationWithIncludes**](CourseLocationWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

