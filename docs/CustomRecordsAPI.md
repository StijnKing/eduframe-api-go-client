# \CustomRecordsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomRecord**](CustomRecordsAPI.md#CreateCustomRecord) | **Post** /custom/objects/{object_id}/records | Create a custom record
[**DeleteCustomRecord**](CustomRecordsAPI.md#DeleteCustomRecord) | **Delete** /custom/objects/{object_id}/records/{record_id} | Delete a custom record
[**GetCustomRecord**](CustomRecordsAPI.md#GetCustomRecord) | **Get** /custom/objects/{object_id}/records/{record_id} | Get a custom record
[**GetCustomRecords**](CustomRecordsAPI.md#GetCustomRecords) | **Get** /custom/objects/{object_id}/records | Get all custom records
[**UpdateCustomRecord**](CustomRecordsAPI.md#UpdateCustomRecord) | **Patch** /custom/objects/{object_id}/records/{record_id} | Update a custom record



## CreateCustomRecord

> CustomRecord CreateCustomRecord(ctx, objectId).CustomRecordPayload(customRecordPayload).Execute()

Create a custom record



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
	objectId := int32(56) // int32 | The unique identifier of the custom object. 
	customRecordPayload := *openapiclient.NewCustomRecordPayload("DisplayName_example", map[string]CustomFieldValue{"key": openapiclient.CustomFieldValue{ArrayOfString: new([]string)}}) // CustomRecordPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomRecordsAPI.CreateCustomRecord(context.Background(), objectId).CustomRecordPayload(customRecordPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRecordsAPI.CreateCustomRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomRecord`: CustomRecord
	fmt.Fprintf(os.Stdout, "Response from `CustomRecordsAPI.CreateCustomRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **int32** | The unique identifier of the custom object.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customRecordPayload** | [**CustomRecordPayload**](CustomRecordPayload.md) |  | 

### Return type

[**CustomRecord**](CustomRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomRecord

> DeleteCustomRecord(ctx, objectId, recordId).Execute()

Delete a custom record



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
	objectId := int32(56) // int32 | The unique identifier of the custom object. 
	recordId := int32(56) // int32 | The unique identifier of the custom record. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomRecordsAPI.DeleteCustomRecord(context.Background(), objectId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRecordsAPI.DeleteCustomRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **int32** | The unique identifier of the custom object.  | 
**recordId** | **int32** | The unique identifier of the custom record.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomRecordRequest struct via the builder pattern


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


## GetCustomRecord

> CustomRecord GetCustomRecord(ctx, objectId, recordId).Execute()

Get a custom record



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
	objectId := int32(56) // int32 | The unique identifier of the custom object. 
	recordId := int32(56) // int32 | The unique identifier of the custom record. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomRecordsAPI.GetCustomRecord(context.Background(), objectId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRecordsAPI.GetCustomRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomRecord`: CustomRecord
	fmt.Fprintf(os.Stdout, "Response from `CustomRecordsAPI.GetCustomRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **int32** | The unique identifier of the custom object.  | 
**recordId** | **int32** | The unique identifier of the custom record.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomRecord**](CustomRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomRecords

> []CustomRecord GetCustomRecords(ctx, objectId).Cursor(cursor).PerPage(perPage).Execute()

Get all custom records



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
	objectId := int32(56) // int32 | The unique identifier of the custom object. 
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomRecordsAPI.GetCustomRecords(context.Background(), objectId).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRecordsAPI.GetCustomRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomRecords`: []CustomRecord
	fmt.Fprintf(os.Stdout, "Response from `CustomRecordsAPI.GetCustomRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **int32** | The unique identifier of the custom object.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]CustomRecord**](CustomRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomRecord

> UpdateCustomRecord(ctx, objectId, recordId).CustomRecordPatchPayload(customRecordPatchPayload).Execute()

Update a custom record



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
	objectId := int32(56) // int32 | The unique identifier of the custom object. 
	recordId := int32(56) // int32 | The unique identifier of the custom record. 
	customRecordPatchPayload := *openapiclient.NewCustomRecordPatchPayload() // CustomRecordPatchPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomRecordsAPI.UpdateCustomRecord(context.Background(), objectId, recordId).CustomRecordPatchPayload(customRecordPatchPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRecordsAPI.UpdateCustomRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **int32** | The unique identifier of the custom object.  | 
**recordId** | **int32** | The unique identifier of the custom record.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customRecordPatchPayload** | [**CustomRecordPatchPayload**](CustomRecordPatchPayload.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

