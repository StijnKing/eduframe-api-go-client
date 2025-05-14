# \CustomAssociationsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssociationsOfObject**](CustomAssociationsAPI.md#GetAssociationsOfObject) | **Get** /custom/{object_type}/associations | Get all associations of an system object



## GetAssociationsOfObject

> []CustomAssociation GetAssociationsOfObject(ctx, objectType).Execute()

Get all associations of an system object



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
	objectType := "objectType_example" // string | The type of the object the custom field is for. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAssociationsAPI.GetAssociationsOfObject(context.Background(), objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAssociationsAPI.GetAssociationsOfObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssociationsOfObject`: []CustomAssociation
	fmt.Fprintf(os.Stdout, "Response from `CustomAssociationsAPI.GetAssociationsOfObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of the object the custom field is for.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssociationsOfObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CustomAssociation**](CustomAssociation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

