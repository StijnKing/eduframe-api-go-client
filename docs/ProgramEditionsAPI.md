# \ProgramEditionsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProgramEdition**](ProgramEditionsAPI.md#CreateProgramEdition) | **Post** /program/editions | Create a program edition
[**DeleteProgramEditionById**](ProgramEditionsAPI.md#DeleteProgramEditionById) | **Delete** /program/editions/{id} | Delete a program edition
[**GetElementsOfProgramEdition**](ProgramEditionsAPI.md#GetElementsOfProgramEdition) | **Get** /program/editions/{id}/elements | Get the elements of a program edition
[**GetProgramEditionById**](ProgramEditionsAPI.md#GetProgramEditionById) | **Get** /program/editions/{id} | Get a program edition
[**GetProgramEditions**](ProgramEditionsAPI.md#GetProgramEditions) | **Get** /program/editions | Get all program editions
[**UpdateProgramEditionById**](ProgramEditionsAPI.md#UpdateProgramEditionById) | **Patch** /program/editions/{id} | Update a program edition



## CreateProgramEdition

> EditionWithIncludes CreateProgramEdition(ctx).ProgramEditionPayload(programEditionPayload).Execute()

Create a program edition



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
	programEditionPayload := *openapiclient.NewProgramEditionPayload(int32(123), "Name_example") // ProgramEditionPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramEditionsAPI.CreateProgramEdition(context.Background()).ProgramEditionPayload(programEditionPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramEditionsAPI.CreateProgramEdition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProgramEdition`: EditionWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramEditionsAPI.CreateProgramEdition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProgramEditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programEditionPayload** | [**ProgramEditionPayload**](ProgramEditionPayload.md) |  | 

### Return type

[**EditionWithIncludes**](EditionWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProgramEditionById

> DeleteProgramEditionById(ctx, id).Execute()

Delete a program edition



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
	r, err := apiClient.ProgramEditionsAPI.DeleteProgramEditionById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramEditionsAPI.DeleteProgramEditionById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteProgramEditionByIdRequest struct via the builder pattern


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


## GetElementsOfProgramEdition

> []ProgramEditionElement GetElementsOfProgramEdition(ctx, id).Execute()

Get the elements of a program edition



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
	resp, r, err := apiClient.ProgramEditionsAPI.GetElementsOfProgramEdition(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramEditionsAPI.GetElementsOfProgramEdition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetElementsOfProgramEdition`: []ProgramEditionElement
	fmt.Fprintf(os.Stdout, "Response from `ProgramEditionsAPI.GetElementsOfProgramEdition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementsOfProgramEditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProgramEditionElement**](ProgramEditionElement.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProgramEditionById

> EditionWithIncludes GetProgramEditionById(ctx, id).Execute()

Get a program edition



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
	resp, r, err := apiClient.ProgramEditionsAPI.GetProgramEditionById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramEditionsAPI.GetProgramEditionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProgramEditionById`: EditionWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramEditionsAPI.GetProgramEditionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramEditionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EditionWithIncludes**](EditionWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProgramEditions

> []EditionWithIncludes GetProgramEditions(ctx).ProgramId(programId).Cursor(cursor).PerPage(perPage).Execute()

Get all program editions



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
	programId := []string{"Inner_example"} // []string | Filter results on program_id (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramEditionsAPI.GetProgramEditions(context.Background()).ProgramId(programId).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramEditionsAPI.GetProgramEditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProgramEditions`: []EditionWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramEditionsAPI.GetProgramEditions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramEditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programId** | **[]string** | Filter results on program_id | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]EditionWithIncludes**](EditionWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProgramEditionById

> EditionWithIncludes UpdateProgramEditionById(ctx, id).ProgramEditionPatchPayload(programEditionPatchPayload).Execute()

Update a program edition



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
	programEditionPatchPayload := *openapiclient.NewProgramEditionPatchPayload() // ProgramEditionPatchPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramEditionsAPI.UpdateProgramEditionById(context.Background(), id).ProgramEditionPatchPayload(programEditionPatchPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramEditionsAPI.UpdateProgramEditionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProgramEditionById`: EditionWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramEditionsAPI.UpdateProgramEditionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProgramEditionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **programEditionPatchPayload** | [**ProgramEditionPatchPayload**](ProgramEditionPatchPayload.md) |  | 

### Return type

[**EditionWithIncludes**](EditionWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

