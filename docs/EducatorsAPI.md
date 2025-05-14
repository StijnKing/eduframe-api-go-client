# \EducatorsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentEducator**](EducatorsAPI.md#GetCurrentEducator) | **Get** /educators/current | Get an educator record



## GetCurrentEducator

> EducatorWithIncludes GetCurrentEducator(ctx).Execute()

Get an educator record

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EducatorsAPI.GetCurrentEducator(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EducatorsAPI.GetCurrentEducator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentEducator`: EducatorWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `EducatorsAPI.GetCurrentEducator`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentEducatorRequest struct via the builder pattern


### Return type

[**EducatorWithIncludes**](EducatorWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

