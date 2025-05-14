# \CreditsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCredits**](CreditsAPI.md#GetCredits) | **Get** /credits | Get all credit records



## GetCredits

> []Credit GetCredits(ctx).StudentId(studentId).Cursor(cursor).PerPage(perPage).Execute()

Get all credit records

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
	studentId := int32(56) // int32 | Filter results on student_id (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditsAPI.GetCredits(context.Background()).StudentId(studentId).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditsAPI.GetCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredits`: []Credit
	fmt.Fprintf(os.Stdout, "Response from `CreditsAPI.GetCredits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **studentId** | **int32** | Filter results on student_id | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]Credit**](Credit.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

