# \CourseTabsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCourseTabs**](CourseTabsAPI.md#GetCourseTabs) | **Get** /course_tabs | Get all course tab records



## GetCourseTabs

> []CourseTab GetCourseTabs(ctx).Execute()

Get all course tab records

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
	resp, r, err := apiClient.CourseTabsAPI.GetCourseTabs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseTabsAPI.GetCourseTabs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseTabs`: []CourseTab
	fmt.Fprintf(os.Stdout, "Response from `CourseTabsAPI.GetCourseTabs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseTabsRequest struct via the builder pattern


### Return type

[**[]CourseTab**](CourseTab.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

