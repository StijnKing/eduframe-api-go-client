# \ProgramEnrollmentsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AwardCertificateToProgramEnrollmentById**](ProgramEnrollmentsAPI.md#AwardCertificateToProgramEnrollmentById) | **Put** /program/enrollments/{id}/award_certificate | Awards a certificate to a program enrollment
[**CancelProgramEnrollmentById**](ProgramEnrollmentsAPI.md#CancelProgramEnrollmentById) | **Put** /program/enrollments/{id}/cancel | Cancel a program enrollment
[**DeleteCertificateFromProgramEnrollmentById**](ProgramEnrollmentsAPI.md#DeleteCertificateFromProgramEnrollmentById) | **Post** /program/enrollments/{id}/delete_certificate | Deletes a certificate from a program enrollment
[**GetProgramEnrollmentById**](ProgramEnrollmentsAPI.md#GetProgramEnrollmentById) | **Get** /program/enrollments/{id} | Get a program enrollment record
[**GetProgramEnrollments**](ProgramEnrollmentsAPI.md#GetProgramEnrollments) | **Get** /program/enrollments | Get all program enrollments



## AwardCertificateToProgramEnrollmentById

> ProgramEnrollmentWithIncludes AwardCertificateToProgramEnrollmentById(ctx, id).AwardCertificateToProgramEnrollmentByIdRequest(awardCertificateToProgramEnrollmentByIdRequest).Execute()

Awards a certificate to a program enrollment



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
	awardCertificateToProgramEnrollmentByIdRequest := *openapiclient.NewAwardCertificateToProgramEnrollmentByIdRequest(int32(123)) // AwardCertificateToProgramEnrollmentByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramEnrollmentsAPI.AwardCertificateToProgramEnrollmentById(context.Background(), id).AwardCertificateToProgramEnrollmentByIdRequest(awardCertificateToProgramEnrollmentByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramEnrollmentsAPI.AwardCertificateToProgramEnrollmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwardCertificateToProgramEnrollmentById`: ProgramEnrollmentWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramEnrollmentsAPI.AwardCertificateToProgramEnrollmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAwardCertificateToProgramEnrollmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awardCertificateToProgramEnrollmentByIdRequest** | [**AwardCertificateToProgramEnrollmentByIdRequest**](AwardCertificateToProgramEnrollmentByIdRequest.md) |  | 

### Return type

[**ProgramEnrollmentWithIncludes**](ProgramEnrollmentWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelProgramEnrollmentById

> ProgramEnrollmentWithIncludes CancelProgramEnrollmentById(ctx, id).Execute()

Cancel a program enrollment



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
	resp, r, err := apiClient.ProgramEnrollmentsAPI.CancelProgramEnrollmentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramEnrollmentsAPI.CancelProgramEnrollmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelProgramEnrollmentById`: ProgramEnrollmentWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramEnrollmentsAPI.CancelProgramEnrollmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelProgramEnrollmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProgramEnrollmentWithIncludes**](ProgramEnrollmentWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificateFromProgramEnrollmentById

> ProgramEnrollmentWithIncludes DeleteCertificateFromProgramEnrollmentById(ctx, id).Execute()

Deletes a certificate from a program enrollment



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
	resp, r, err := apiClient.ProgramEnrollmentsAPI.DeleteCertificateFromProgramEnrollmentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramEnrollmentsAPI.DeleteCertificateFromProgramEnrollmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertificateFromProgramEnrollmentById`: ProgramEnrollmentWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramEnrollmentsAPI.DeleteCertificateFromProgramEnrollmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateFromProgramEnrollmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProgramEnrollmentWithIncludes**](ProgramEnrollmentWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProgramEnrollmentById

> ProgramEnrollmentWithIncludes GetProgramEnrollmentById(ctx, id).Execute()

Get a program enrollment record



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
	resp, r, err := apiClient.ProgramEnrollmentsAPI.GetProgramEnrollmentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramEnrollmentsAPI.GetProgramEnrollmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProgramEnrollmentById`: ProgramEnrollmentWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramEnrollmentsAPI.GetProgramEnrollmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramEnrollmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProgramEnrollmentWithIncludes**](ProgramEnrollmentWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProgramEnrollments

> []ProgramEnrollmentWithIncludes GetProgramEnrollments(ctx).StudentId(studentId).EditionId(editionId).Cursor(cursor).PerPage(perPage).Execute()

Get all program enrollments



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
	editionId := int32(56) // int32 | Filter results on edition_id (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramEnrollmentsAPI.GetProgramEnrollments(context.Background()).StudentId(studentId).EditionId(editionId).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramEnrollmentsAPI.GetProgramEnrollments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProgramEnrollments`: []ProgramEnrollmentWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `ProgramEnrollmentsAPI.GetProgramEnrollments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramEnrollmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **studentId** | **int32** | Filter results on student_id | 
 **editionId** | **int32** | Filter results on edition_id | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]ProgramEnrollmentWithIncludes**](ProgramEnrollmentWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

