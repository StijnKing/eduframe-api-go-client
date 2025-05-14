# \ProgramProgramsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProgram**](ProgramProgramsAPI.md#CreateProgram) | **Post** /program/programs | Create a program
[**DeleteProgramById**](ProgramProgramsAPI.md#DeleteProgramById) | **Delete** /program/programs/{id} | Delete a program
[**GetProgramById**](ProgramProgramsAPI.md#GetProgramById) | **Get** /program/programs/{id} | Get a program
[**GetPrograms**](ProgramProgramsAPI.md#GetPrograms) | **Get** /program/programs | Get all programs
[**UpdateProgramById**](ProgramProgramsAPI.md#UpdateProgramById) | **Patch** /program/programs/{id} | Update a program



## CreateProgram

> ProgramWithIncludes CreateProgram(ctx).ProgramProgramPayload(programProgramPayload).Execute()

Create a program



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
	programProgramPayload := *openapiclient.NewProgramProgramPayload("Name_example", int32(123), "CostScheme_example") // ProgramProgramPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramProgramsAPI.CreateProgram(context.Background()).ProgramProgramPayload(programProgramPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramProgramsAPI.CreateProgram``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProgram`: ProgramWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramProgramsAPI.CreateProgram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programProgramPayload** | [**ProgramProgramPayload**](ProgramProgramPayload.md) |  | 

### Return type

[**ProgramWithIncludes**](ProgramWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProgramById

> DeleteProgramById(ctx, id).Execute()

Delete a program



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
	r, err := apiClient.ProgramProgramsAPI.DeleteProgramById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramProgramsAPI.DeleteProgramById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteProgramByIdRequest struct via the builder pattern


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


## GetProgramById

> ProgramWithIncludes GetProgramById(ctx, id).Execute()

Get a program



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
	resp, r, err := apiClient.ProgramProgramsAPI.GetProgramById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramProgramsAPI.GetProgramById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProgramById`: ProgramWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramProgramsAPI.GetProgramById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProgramWithIncludes**](ProgramWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrograms

> []ProgramWithIncludes GetPrograms(ctx).Cursor(cursor).PerPage(perPage).Execute()

Get all programs



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
	resp, r, err := apiClient.ProgramProgramsAPI.GetPrograms(context.Background()).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramProgramsAPI.GetPrograms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrograms`: []ProgramWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramProgramsAPI.GetPrograms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]ProgramWithIncludes**](ProgramWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProgramById

> ProgramWithIncludes UpdateProgramById(ctx, id).ProgramProgramPatchPayload(programProgramPatchPayload).Execute()

Update a program



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
	programProgramPatchPayload := *openapiclient.NewProgramProgramPatchPayload() // ProgramProgramPatchPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramProgramsAPI.UpdateProgramById(context.Background(), id).ProgramProgramPatchPayload(programProgramPatchPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramProgramsAPI.UpdateProgramById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProgramById`: ProgramWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramProgramsAPI.UpdateProgramById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProgramByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **programProgramPatchPayload** | [**ProgramProgramPatchPayload**](ProgramProgramPatchPayload.md) |  | 

### Return type

[**ProgramWithIncludes**](ProgramWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

