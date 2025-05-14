# \SignupQuestionsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSignupQuestions**](SignupQuestionsAPI.md#GetSignupQuestions) | **Get** /signup_questions | Get all signup_question records



## GetSignupQuestions

> []SignupQuestion GetSignupQuestions(ctx).ForUser(forUser).ForAccount(forAccount).ForType(forType).Visibility(visibility).UseAsDuplicateIndicator(useAsDuplicateIndicator).Cursor(cursor).PerPage(perPage).Execute()

Get all signup_question records

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
	forUser := "forUser_example" // string | Filter results on for_user (optional)
	forAccount := "forAccount_example" // string | Filter results on for_account (optional)
	forType := "forType_example" // string | Filter results on for_type (optional)
	visibility := "visibility_example" // string | Filter results on visibility (optional)
	useAsDuplicateIndicator := true // bool | Filter results on use_as_duplicate_indicator (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignupQuestionsAPI.GetSignupQuestions(context.Background()).ForUser(forUser).ForAccount(forAccount).ForType(forType).Visibility(visibility).UseAsDuplicateIndicator(useAsDuplicateIndicator).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignupQuestionsAPI.GetSignupQuestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignupQuestions`: []SignupQuestion
	fmt.Fprintf(os.Stdout, "Response from `SignupQuestionsAPI.GetSignupQuestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignupQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forUser** | **string** | Filter results on for_user | 
 **forAccount** | **string** | Filter results on for_account | 
 **forType** | **string** | Filter results on for_type | 
 **visibility** | **string** | Filter results on visibility | 
 **useAsDuplicateIndicator** | **bool** | Filter results on use_as_duplicate_indicator | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]SignupQuestion**](SignupQuestion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

