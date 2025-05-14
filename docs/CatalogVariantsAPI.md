# \CatalogVariantsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogVariantById**](CatalogVariantsAPI.md#GetCatalogVariantById) | **Get** /catalog/variants/{id} | Get a catalog variant record
[**GetCatalogVariants**](CatalogVariantsAPI.md#GetCatalogVariants) | **Get** /catalog/variants | Get all catalog variants
[**UpdateCatalogVariantById**](CatalogVariantsAPI.md#UpdateCatalogVariantById) | **Patch** /catalog/variants/{id} | Update a catalog variant



## GetCatalogVariantById

> VariantWithIncludes GetCatalogVariantById(ctx, id).Execute()

Get a catalog variant record

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
	resp, r, err := apiClient.CatalogVariantsAPI.GetCatalogVariantById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogVariantsAPI.GetCatalogVariantById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogVariantById`: VariantWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CatalogVariantsAPI.GetCatalogVariantById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogVariantByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VariantWithIncludes**](VariantWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogVariants

> []VariantWithIncludes GetCatalogVariants(ctx).PublishedPublic(publishedPublic).ProductId(productId).VariantableId(variantableId).VariantableType(variantableType).Cursor(cursor).PerPage(perPage).Execute()

Get all catalog variants

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
	publishedPublic := "publishedPublic_example" // string | Only show published variants and planned_courses that are either planned or in progress (optional)
	productId := int32(56) // int32 | Filter results on product_id (optional)
	variantableId := int32(56) // int32 | Filter results on variantable_id (optional)
	variantableType := "variantableType_example" // string | Filter results on variantable_type (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogVariantsAPI.GetCatalogVariants(context.Background()).PublishedPublic(publishedPublic).ProductId(productId).VariantableId(variantableId).VariantableType(variantableType).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogVariantsAPI.GetCatalogVariants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogVariants`: []VariantWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CatalogVariantsAPI.GetCatalogVariants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogVariantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publishedPublic** | **string** | Only show published variants and planned_courses that are either planned or in progress | 
 **productId** | **int32** | Filter results on product_id | 
 **variantableId** | **int32** | Filter results on variantable_id | 
 **variantableType** | **string** | Filter results on variantable_type | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]VariantWithIncludes**](VariantWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogVariantById

> VariantWithIncludes UpdateCatalogVariantById(ctx, id).UpdateCatalogVariantByIdRequest(updateCatalogVariantByIdRequest).Execute()

Update a catalog variant

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
	updateCatalogVariantByIdRequest := *openapiclient.NewUpdateCatalogVariantByIdRequest(false) // UpdateCatalogVariantByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogVariantsAPI.UpdateCatalogVariantById(context.Background(), id).UpdateCatalogVariantByIdRequest(updateCatalogVariantByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogVariantsAPI.UpdateCatalogVariantById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCatalogVariantById`: VariantWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CatalogVariantsAPI.UpdateCatalogVariantById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogVariantByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCatalogVariantByIdRequest** | [**UpdateCatalogVariantByIdRequest**](UpdateCatalogVariantByIdRequest.md) |  | 

### Return type

[**VariantWithIncludes**](VariantWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

