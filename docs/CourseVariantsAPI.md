# \CourseVariantsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCourseVariant**](CourseVariantsAPI.md#CreateCourseVariant) | **Post** /course_variants | Create a course variant
[**GetCourseVariantById**](CourseVariantsAPI.md#GetCourseVariantById) | **Get** /course_variants/{id} | Get a course variant record
[**GetCourseVariants**](CourseVariantsAPI.md#GetCourseVariants) | **Get** /course_variants | Get all course variant records



## CreateCourseVariant

> CourseVariant CreateCourseVariant(ctx).CreateCourseVariantRequest(createCourseVariantRequest).Execute()

Create a course variant

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
	createCourseVariantRequest := *openapiclient.NewCreateCourseVariantRequest("Name_example") // CreateCourseVariantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseVariantsAPI.CreateCourseVariant(context.Background()).CreateCourseVariantRequest(createCourseVariantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseVariantsAPI.CreateCourseVariant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCourseVariant`: CourseVariant
	fmt.Fprintf(os.Stdout, "Response from `CourseVariantsAPI.CreateCourseVariant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCourseVariantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCourseVariantRequest** | [**CreateCourseVariantRequest**](CreateCourseVariantRequest.md) |  | 

### Return type

[**CourseVariant**](CourseVariant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseVariantById

> CourseVariant GetCourseVariantById(ctx, id).Execute()

Get a course variant record

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
	resp, r, err := apiClient.CourseVariantsAPI.GetCourseVariantById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseVariantsAPI.GetCourseVariantById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseVariantById`: CourseVariant
	fmt.Fprintf(os.Stdout, "Response from `CourseVariantsAPI.GetCourseVariantById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseVariantByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CourseVariant**](CourseVariant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseVariants

> []CourseVariant GetCourseVariants(ctx).Cursor(cursor).PerPage(perPage).Execute()

Get all course variant records

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
	resp, r, err := apiClient.CourseVariantsAPI.GetCourseVariants(context.Background()).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseVariantsAPI.GetCourseVariants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseVariants`: []CourseVariant
	fmt.Fprintf(os.Stdout, "Response from `CourseVariantsAPI.GetCourseVariants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]CourseVariant**](CourseVariant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

