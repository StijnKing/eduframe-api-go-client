# \TeacherRolesAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeacherRole**](TeacherRolesAPI.md#CreateTeacherRole) | **Post** /teacher_roles | Create a teacher role.
[**DeleteTeacherRoleById**](TeacherRolesAPI.md#DeleteTeacherRoleById) | **Delete** /teacher_roles/{id} | Delete a teacher role.
[**GetTeacherRoleById**](TeacherRolesAPI.md#GetTeacherRoleById) | **Get** /teacher_roles/{id} | Get a teacher role
[**GetTeacherRoles**](TeacherRolesAPI.md#GetTeacherRoles) | **Get** /teacher_roles | Get all teacher roles
[**UpdateTeacherRoleById**](TeacherRolesAPI.md#UpdateTeacherRoleById) | **Patch** /teacher_roles/{id} | Update a teacher role.



## CreateTeacherRole

> TeacherRole CreateTeacherRole(ctx).CreateTeacherRoleRequest(createTeacherRoleRequest).Execute()

Create a teacher role.

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
	createTeacherRoleRequest := *openapiclient.NewCreateTeacherRoleRequest("Name_example") // CreateTeacherRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeacherRolesAPI.CreateTeacherRole(context.Background()).CreateTeacherRoleRequest(createTeacherRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeacherRolesAPI.CreateTeacherRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeacherRole`: TeacherRole
	fmt.Fprintf(os.Stdout, "Response from `TeacherRolesAPI.CreateTeacherRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeacherRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeacherRoleRequest** | [**CreateTeacherRoleRequest**](CreateTeacherRoleRequest.md) |  | 

### Return type

[**TeacherRole**](TeacherRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeacherRoleById

> DeleteTeacherRoleById(ctx, id).Execute()

Delete a teacher role.

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
	r, err := apiClient.TeacherRolesAPI.DeleteTeacherRoleById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeacherRolesAPI.DeleteTeacherRoleById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTeacherRoleByIdRequest struct via the builder pattern


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


## GetTeacherRoleById

> TeacherRole GetTeacherRoleById(ctx, id).Execute()

Get a teacher role

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
	resp, r, err := apiClient.TeacherRolesAPI.GetTeacherRoleById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeacherRolesAPI.GetTeacherRoleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeacherRoleById`: TeacherRole
	fmt.Fprintf(os.Stdout, "Response from `TeacherRolesAPI.GetTeacherRoleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeacherRoleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeacherRole**](TeacherRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeacherRoles

> []TeacherRole GetTeacherRoles(ctx).Cursor(cursor).PerPage(perPage).Execute()

Get all teacher roles

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
	resp, r, err := apiClient.TeacherRolesAPI.GetTeacherRoles(context.Background()).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeacherRolesAPI.GetTeacherRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeacherRoles`: []TeacherRole
	fmt.Fprintf(os.Stdout, "Response from `TeacherRolesAPI.GetTeacherRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeacherRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]TeacherRole**](TeacherRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeacherRoleById

> TeacherRole UpdateTeacherRoleById(ctx, id).CreateTeacherRoleRequest(createTeacherRoleRequest).Execute()

Update a teacher role.

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
	createTeacherRoleRequest := *openapiclient.NewCreateTeacherRoleRequest("Name_example") // CreateTeacherRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeacherRolesAPI.UpdateTeacherRoleById(context.Background(), id).CreateTeacherRoleRequest(createTeacherRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeacherRolesAPI.UpdateTeacherRoleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeacherRoleById`: TeacherRole
	fmt.Fprintf(os.Stdout, "Response from `TeacherRolesAPI.UpdateTeacherRoleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeacherRoleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTeacherRoleRequest** | [**CreateTeacherRoleRequest**](CreateTeacherRoleRequest.md) |  | 

### Return type

[**TeacherRole**](TeacherRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

