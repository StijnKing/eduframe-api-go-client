# \MaterialsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMaterial**](MaterialsAPI.md#CreateMaterial) | **Post** /materials | Create a material.
[**DeleteMaterialById**](MaterialsAPI.md#DeleteMaterialById) | **Delete** /materials/{id} | Delete a material.
[**GetMaterials**](MaterialsAPI.md#GetMaterials) | **Get** /materials | Get all material records
[**UpdateMaterialById**](MaterialsAPI.md#UpdateMaterialById) | **Patch** /materials/{id} | Update a material.



## CreateMaterial

> MaterialWithIncludes CreateMaterial(ctx).CreateMaterialRequest(createMaterialRequest).Execute()

Create a material.

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
	createMaterialRequest := *openapiclient.NewCreateMaterialRequest("Name_example", int32(123)) // CreateMaterialRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaterialsAPI.CreateMaterial(context.Background()).CreateMaterialRequest(createMaterialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaterialsAPI.CreateMaterial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMaterial`: MaterialWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `MaterialsAPI.CreateMaterial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMaterialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMaterialRequest** | [**CreateMaterialRequest**](CreateMaterialRequest.md) |  | 

### Return type

[**MaterialWithIncludes**](MaterialWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMaterialById

> DeleteMaterialById(ctx, id).Execute()

Delete a material.

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
	r, err := apiClient.MaterialsAPI.DeleteMaterialById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaterialsAPI.DeleteMaterialById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMaterialByIdRequest struct via the builder pattern


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


## GetMaterials

> []MaterialWithIncludes GetMaterials(ctx).Cursor(cursor).PerPage(perPage).Execute()

Get all material records

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
	resp, r, err := apiClient.MaterialsAPI.GetMaterials(context.Background()).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaterialsAPI.GetMaterials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaterials`: []MaterialWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `MaterialsAPI.GetMaterials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMaterialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]MaterialWithIncludes**](MaterialWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMaterialById

> MaterialWithIncludes UpdateMaterialById(ctx, id).UpdateMaterialByIdRequest(updateMaterialByIdRequest).Execute()

Update a material.

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
	updateMaterialByIdRequest := *openapiclient.NewUpdateMaterialByIdRequest() // UpdateMaterialByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaterialsAPI.UpdateMaterialById(context.Background(), id).UpdateMaterialByIdRequest(updateMaterialByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaterialsAPI.UpdateMaterialById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMaterialById`: MaterialWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `MaterialsAPI.UpdateMaterialById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMaterialByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMaterialByIdRequest** | [**UpdateMaterialByIdRequest**](UpdateMaterialByIdRequest.md) |  | 

### Return type

[**MaterialWithIncludes**](MaterialWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

