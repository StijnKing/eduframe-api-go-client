# \AuthenticationsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthentication**](AuthenticationsAPI.md#CreateAuthentication) | **Post** /authentications | Create an authentication.
[**DeleteAuthenticationFromUser**](AuthenticationsAPI.md#DeleteAuthenticationFromUser) | **Delete** /users/{user_id}/authentications/{id} | Remove an authentication from a user. NOTE: It is impossible to remove the last authentication for a user.
[**GetAuthenticationsByUserId**](AuthenticationsAPI.md#GetAuthenticationsByUserId) | **Get** /users/{user_id}/authentications | Get the authentications of an user



## CreateAuthentication

> Authentication CreateAuthentication(ctx).CreateAuthenticationRequest(createAuthenticationRequest).Execute()

Create an authentication.

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
	createAuthenticationRequest := *openapiclient.NewCreateAuthenticationRequest("Uid_example", int32(123), openapiclient.AuthenticationProviderType("azure_active_directory")) // CreateAuthenticationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationsAPI.CreateAuthentication(context.Background()).CreateAuthenticationRequest(createAuthenticationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationsAPI.CreateAuthentication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAuthentication`: Authentication
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationsAPI.CreateAuthentication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAuthenticationRequest** | [**CreateAuthenticationRequest**](CreateAuthenticationRequest.md) |  | 

### Return type

[**Authentication**](Authentication.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthenticationFromUser

> DeleteAuthenticationFromUser(ctx, userId, id).Execute()

Remove an authentication from a user. NOTE: It is impossible to remove the last authentication for a user.

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
	userId := int32(56) // int32 | 
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationsAPI.DeleteAuthenticationFromUser(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationsAPI.DeleteAuthenticationFromUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthenticationFromUserRequest struct via the builder pattern


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


## GetAuthenticationsByUserId

> []Authentication GetAuthenticationsByUserId(ctx, userId).Provider(provider).Execute()

Get the authentications of an user

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
	userId := int32(56) // int32 | 
	provider := openapiclient.AuthenticationProviderType("azure_active_directory") // AuthenticationProviderType | Filter results on provider (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationsAPI.GetAuthenticationsByUserId(context.Background(), userId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationsAPI.GetAuthenticationsByUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthenticationsByUserId`: []Authentication
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationsAPI.GetAuthenticationsByUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationsByUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | [**AuthenticationProviderType**](AuthenticationProviderType.md) | Filter results on provider | 

### Return type

[**[]Authentication**](Authentication.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

