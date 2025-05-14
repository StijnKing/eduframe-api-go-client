# \AffiliationsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAffiliation**](AffiliationsAPI.md#CreateAffiliation) | **Post** /affiliations | Create an affiliation affiliations
[**DeleteAffiliationById**](AffiliationsAPI.md#DeleteAffiliationById) | **Delete** /affiliations/{id} | Delete an affiliation
[**GetAffiliations**](AffiliationsAPI.md#GetAffiliations) | **Get** /affiliations | Get all affiliations
[**UpdateAffiliationById**](AffiliationsAPI.md#UpdateAffiliationById) | **Patch** /affiliations/{id} | Update an affiliation.



## CreateAffiliation

> Affiliation CreateAffiliation(ctx).CreateAffiliationRequest(createAffiliationRequest).Execute()

Create an affiliation affiliations

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
	createAffiliationRequest := *openapiclient.NewCreateAffiliationRequest(int32(123), int32(123)) // CreateAffiliationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliationsAPI.CreateAffiliation(context.Background()).CreateAffiliationRequest(createAffiliationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliationsAPI.CreateAffiliation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAffiliation`: Affiliation
	fmt.Fprintf(os.Stdout, "Response from `AffiliationsAPI.CreateAffiliation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAffiliationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAffiliationRequest** | [**CreateAffiliationRequest**](CreateAffiliationRequest.md) |  | 

### Return type

[**Affiliation**](Affiliation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAffiliationById

> DeleteAffiliationById(ctx, id).Execute()

Delete an affiliation

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
	r, err := apiClient.AffiliationsAPI.DeleteAffiliationById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliationsAPI.DeleteAffiliationById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAffiliationByIdRequest struct via the builder pattern


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


## GetAffiliations

> []Affiliation GetAffiliations(ctx).UserId(userId).AccountId(accountId).Cursor(cursor).PerPage(perPage).Execute()

Get all affiliations

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
	userId := int32(56) // int32 | Filter results on user_id (optional)
	accountId := int32(56) // int32 | Filter results on account_id (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliationsAPI.GetAffiliations(context.Background()).UserId(userId).AccountId(accountId).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliationsAPI.GetAffiliations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliations`: []Affiliation
	fmt.Fprintf(os.Stdout, "Response from `AffiliationsAPI.GetAffiliations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32** | Filter results on user_id | 
 **accountId** | **int32** | Filter results on account_id | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]Affiliation**](Affiliation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAffiliationById

> Affiliation UpdateAffiliationById(ctx, id).UpdateAffiliationByIdRequest(updateAffiliationByIdRequest).Execute()

Update an affiliation.

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
	updateAffiliationByIdRequest := *openapiclient.NewUpdateAffiliationByIdRequest() // UpdateAffiliationByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliationsAPI.UpdateAffiliationById(context.Background(), id).UpdateAffiliationByIdRequest(updateAffiliationByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliationsAPI.UpdateAffiliationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAffiliationById`: Affiliation
	fmt.Fprintf(os.Stdout, "Response from `AffiliationsAPI.UpdateAffiliationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAffiliationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAffiliationByIdRequest** | [**UpdateAffiliationByIdRequest**](UpdateAffiliationByIdRequest.md) |  | 

### Return type

[**Affiliation**](Affiliation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

