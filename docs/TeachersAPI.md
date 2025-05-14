# \TeachersAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateTeacherById**](TeachersAPI.md#ActivateTeacherById) | **Post** /teachers/{id}/activate | Mark teacher as active
[**CreateTeacher**](TeachersAPI.md#CreateTeacher) | **Post** /teachers | Create a new teacher
[**DeactivateTeacherById**](TeachersAPI.md#DeactivateTeacherById) | **Post** /teachers/{id}/deactivate | Mark teacher as inactive
[**GetTeacherById**](TeachersAPI.md#GetTeacherById) | **Get** /teachers/{id} | Get a teacher record
[**GetTeachers**](TeachersAPI.md#GetTeachers) | **Get** /teachers | Get all teacher records



## ActivateTeacherById

> ActivateTeacherById(ctx, id).Execute()

Mark teacher as active

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
	r, err := apiClient.TeachersAPI.ActivateTeacherById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeachersAPI.ActivateTeacherById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActivateTeacherByIdRequest struct via the builder pattern


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


## CreateTeacher

> CreateTeacher201Response CreateTeacher(ctx).CreateTeacherRequest(createTeacherRequest).Execute()

Create a new teacher

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
	createTeacherRequest := *openapiclient.NewCreateTeacherRequest() // CreateTeacherRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeachersAPI.CreateTeacher(context.Background()).CreateTeacherRequest(createTeacherRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeachersAPI.CreateTeacher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeacher`: CreateTeacher201Response
	fmt.Fprintf(os.Stdout, "Response from `TeachersAPI.CreateTeacher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeacherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeacherRequest** | [**CreateTeacherRequest**](CreateTeacherRequest.md) |  | 

### Return type

[**CreateTeacher201Response**](CreateTeacher201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateTeacherById

> DeactivateTeacherById(ctx, id).Execute()

Mark teacher as inactive

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
	r, err := apiClient.TeachersAPI.DeactivateTeacherById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeachersAPI.DeactivateTeacherById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeactivateTeacherByIdRequest struct via the builder pattern


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


## GetTeacherById

> TeacherWithIncludes GetTeacherById(ctx, id).Execute()

Get a teacher record

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
	resp, r, err := apiClient.TeachersAPI.GetTeacherById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeachersAPI.GetTeacherById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeacherById`: TeacherWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `TeachersAPI.GetTeacherById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeacherByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeacherWithIncludes**](TeacherWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeachers

> []TeacherWithIncludes GetTeachers(ctx).Search(search).LabelId(labelId).Id(id).Cursor(cursor).PerPage(perPage).Execute()

Get all teacher records

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
	search := "search_example" // string | Filter results on search (optional)
	labelId := []int32{int32(123)} // []int32 | Filter results on label_id (optional)
	id := []int32{int32(123)} // []int32 | Filter results on id (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeachersAPI.GetTeachers(context.Background()).Search(search).LabelId(labelId).Id(id).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeachersAPI.GetTeachers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeachers`: []TeacherWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `TeachersAPI.GetTeachers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeachersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Filter results on search | 
 **labelId** | **[]int32** | Filter results on label_id | 
 **id** | **[]int32** | Filter results on id | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]TeacherWithIncludes**](TeacherWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

