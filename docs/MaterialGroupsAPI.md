# \MaterialGroupsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMaterialGroup**](MaterialGroupsAPI.md#CreateMaterialGroup) | **Post** /material_groups | Create a material group.
[**DeleteMaterialGroupById**](MaterialGroupsAPI.md#DeleteMaterialGroupById) | **Delete** /material_groups/{id} | Delete a material group.
[**GetMaterialGroupById**](MaterialGroupsAPI.md#GetMaterialGroupById) | **Get** /material_groups/{id} | Get a material group record
[**GetMaterialGroups**](MaterialGroupsAPI.md#GetMaterialGroups) | **Get** /material_groups | Get all material group records
[**UpdateMaterialGroupById**](MaterialGroupsAPI.md#UpdateMaterialGroupById) | **Patch** /material_groups/{id} | Update a material group.



## CreateMaterialGroup

> MaterialGroup CreateMaterialGroup(ctx).CreateMaterialGroupRequest(createMaterialGroupRequest).Execute()

Create a material group.

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
	createMaterialGroupRequest := *openapiclient.NewCreateMaterialGroupRequest("Name_example") // CreateMaterialGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaterialGroupsAPI.CreateMaterialGroup(context.Background()).CreateMaterialGroupRequest(createMaterialGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaterialGroupsAPI.CreateMaterialGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMaterialGroup`: MaterialGroup
	fmt.Fprintf(os.Stdout, "Response from `MaterialGroupsAPI.CreateMaterialGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMaterialGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMaterialGroupRequest** | [**CreateMaterialGroupRequest**](CreateMaterialGroupRequest.md) |  | 

### Return type

[**MaterialGroup**](MaterialGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMaterialGroupById

> DeleteMaterialGroupById(ctx, id).Execute()

Delete a material group.

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
	r, err := apiClient.MaterialGroupsAPI.DeleteMaterialGroupById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaterialGroupsAPI.DeleteMaterialGroupById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMaterialGroupByIdRequest struct via the builder pattern


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


## GetMaterialGroupById

> MaterialGroup GetMaterialGroupById(ctx, id).Execute()

Get a material group record

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
	resp, r, err := apiClient.MaterialGroupsAPI.GetMaterialGroupById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaterialGroupsAPI.GetMaterialGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaterialGroupById`: MaterialGroup
	fmt.Fprintf(os.Stdout, "Response from `MaterialGroupsAPI.GetMaterialGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaterialGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MaterialGroup**](MaterialGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaterialGroups

> []MaterialGroup GetMaterialGroups(ctx).Cursor(cursor).PerPage(perPage).Execute()

Get all material group records

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
	resp, r, err := apiClient.MaterialGroupsAPI.GetMaterialGroups(context.Background()).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaterialGroupsAPI.GetMaterialGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaterialGroups`: []MaterialGroup
	fmt.Fprintf(os.Stdout, "Response from `MaterialGroupsAPI.GetMaterialGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMaterialGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]MaterialGroup**](MaterialGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMaterialGroupById

> MaterialGroup UpdateMaterialGroupById(ctx, id).UpdateMaterialGroupByIdRequest(updateMaterialGroupByIdRequest).Execute()

Update a material group.

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
	updateMaterialGroupByIdRequest := *openapiclient.NewUpdateMaterialGroupByIdRequest() // UpdateMaterialGroupByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaterialGroupsAPI.UpdateMaterialGroupById(context.Background(), id).UpdateMaterialGroupByIdRequest(updateMaterialGroupByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaterialGroupsAPI.UpdateMaterialGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMaterialGroupById`: MaterialGroup
	fmt.Fprintf(os.Stdout, "Response from `MaterialGroupsAPI.UpdateMaterialGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMaterialGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMaterialGroupByIdRequest** | [**UpdateMaterialGroupByIdRequest**](UpdateMaterialGroupByIdRequest.md) |  | 

### Return type

[**MaterialGroup**](MaterialGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

