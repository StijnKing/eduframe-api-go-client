# \CustomFieldOptionsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOptionToCustomField**](CustomFieldOptionsAPI.md#AddOptionToCustomField) | **Post** /custom/{object_type}/fields/{field_slug}/options | Add an option to a custom field
[**DeleteOptionOfCustomField**](CustomFieldOptionsAPI.md#DeleteOptionOfCustomField) | **Delete** /custom/{object_type}/fields/{field_slug}/options/{option_id} | Delete an option from custom field
[**GetOptionOfCustomField**](CustomFieldOptionsAPI.md#GetOptionOfCustomField) | **Get** /custom/{object_type}/fields/{field_slug}/options/{option_id} | Get an option of a custom field
[**GetOptionsOfCustomField**](CustomFieldOptionsAPI.md#GetOptionsOfCustomField) | **Get** /custom/{object_type}/fields/{field_slug}/options | Get all options of a custom field
[**UpdateOptionOfCustomField**](CustomFieldOptionsAPI.md#UpdateOptionOfCustomField) | **Patch** /custom/{object_type}/fields/{field_slug}/options/{option_id} | Update an option of a custom field



## AddOptionToCustomField

> CustomFieldOption AddOptionToCustomField(ctx, objectType, fieldSlug).CustomFieldOptionPayload(customFieldOptionPayload).Execute()

Add an option to a custom field

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
	fieldSlug := "fieldSlug_example" // string | The unique identifier of the custom field. 
	customFieldOptionPayload := *openapiclient.NewCustomFieldOptionPayload("Value_example", false) // CustomFieldOptionPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldOptionsAPI.AddOptionToCustomField(context.Background(), objectType, fieldSlug).CustomFieldOptionPayload(customFieldOptionPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldOptionsAPI.AddOptionToCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOptionToCustomField`: CustomFieldOption
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldOptionsAPI.AddOptionToCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of the object the custom field is for.  | 
**fieldSlug** | **string** | The unique identifier of the custom field.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOptionToCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customFieldOptionPayload** | [**CustomFieldOptionPayload**](CustomFieldOptionPayload.md) |  | 

### Return type

[**CustomFieldOption**](CustomFieldOption.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOptionOfCustomField

> DeleteOptionOfCustomField(ctx, objectType, fieldSlug, optionId).Execute()

Delete an option from custom field

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
	fieldSlug := "fieldSlug_example" // string | The unique identifier of the custom field. 
	optionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomFieldOptionsAPI.DeleteOptionOfCustomField(context.Background(), objectType, fieldSlug, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldOptionsAPI.DeleteOptionOfCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of the object the custom field is for.  | 
**fieldSlug** | **string** | The unique identifier of the custom field.  | 
**optionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOptionOfCustomFieldRequest struct via the builder pattern


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


## GetOptionOfCustomField

> CustomFieldOption GetOptionOfCustomField(ctx, objectType, fieldSlug, optionId).Execute()

Get an option of a custom field

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
	fieldSlug := "fieldSlug_example" // string | The unique identifier of the custom field. 
	optionId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldOptionsAPI.GetOptionOfCustomField(context.Background(), objectType, fieldSlug, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldOptionsAPI.GetOptionOfCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptionOfCustomField`: CustomFieldOption
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldOptionsAPI.GetOptionOfCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of the object the custom field is for.  | 
**fieldSlug** | **string** | The unique identifier of the custom field.  | 
**optionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionOfCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CustomFieldOption**](CustomFieldOption.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOptionsOfCustomField

> []CustomFieldOption GetOptionsOfCustomField(ctx, objectType, fieldSlug).Cursor(cursor).PerPage(perPage).Execute()

Get all options of a custom field

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
	fieldSlug := "fieldSlug_example" // string | The unique identifier of the custom field. 
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldOptionsAPI.GetOptionsOfCustomField(context.Background(), objectType, fieldSlug).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldOptionsAPI.GetOptionsOfCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptionsOfCustomField`: []CustomFieldOption
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldOptionsAPI.GetOptionsOfCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of the object the custom field is for.  | 
**fieldSlug** | **string** | The unique identifier of the custom field.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionsOfCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]CustomFieldOption**](CustomFieldOption.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOptionOfCustomField

> UpdateOptionOfCustomField(ctx, objectType, fieldSlug, optionId).CustomFieldOptionPatchPayload(customFieldOptionPatchPayload).Execute()

Update an option of a custom field

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
	fieldSlug := "fieldSlug_example" // string | The unique identifier of the custom field. 
	optionId := int32(56) // int32 | 
	customFieldOptionPatchPayload := *openapiclient.NewCustomFieldOptionPatchPayload() // CustomFieldOptionPatchPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomFieldOptionsAPI.UpdateOptionOfCustomField(context.Background(), objectType, fieldSlug, optionId).CustomFieldOptionPatchPayload(customFieldOptionPatchPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldOptionsAPI.UpdateOptionOfCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | The type of the object the custom field is for.  | 
**fieldSlug** | **string** | The unique identifier of the custom field.  | 
**optionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOptionOfCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **customFieldOptionPatchPayload** | [**CustomFieldOptionPatchPayload**](CustomFieldOptionPatchPayload.md) |  | 

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

