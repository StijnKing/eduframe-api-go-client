# \CatalogProductsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogProductById**](CatalogProductsAPI.md#GetCatalogProductById) | **Get** /catalog/products/{id} | Get a catalog product record
[**GetCatalogProducts**](CatalogProductsAPI.md#GetCatalogProducts) | **Get** /catalog/products | Get all catalog products
[**UpdateCatalogProductById**](CatalogProductsAPI.md#UpdateCatalogProductById) | **Patch** /catalog/products/{id} | Update a catalog product



## GetCatalogProductById

> ProductWithIncludes GetCatalogProductById(ctx, id).Execute()

Get a catalog product record

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
	resp, r, err := apiClient.CatalogProductsAPI.GetCatalogProductById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogProductsAPI.GetCatalogProductById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogProductById`: ProductWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CatalogProductsAPI.GetCatalogProductById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogProductByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductWithIncludes**](ProductWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogProducts

> []ProductWithIncludes GetCatalogProducts(ctx).Published(published).CategoryId(categoryId).ProductableType(productableType).Search(search).Sort(sort).Cursor(cursor).PerPage(perPage).Execute()

Get all catalog products

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
	published := "published_example" // string | Show only published products (optional)
	categoryId := int32(56) // int32 | Filter results on category_id (optional)
	productableType := "productableType_example" // string | Filter results on productable_type (optional)
	search := "search_example" // string | Filter results on search (optional)
	sort := []string{"Sort_example"} // []string | Sort the results. Can change order by using `<sort_by>:<direction>` where `<direction>` is either `asc` or `desc` (optional)
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogProductsAPI.GetCatalogProducts(context.Background()).Published(published).CategoryId(categoryId).ProductableType(productableType).Search(search).Sort(sort).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogProductsAPI.GetCatalogProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogProducts`: []ProductWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CatalogProductsAPI.GetCatalogProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **published** | **string** | Show only published products | 
 **categoryId** | **int32** | Filter results on category_id | 
 **productableType** | **string** | Filter results on productable_type | 
 **search** | **string** | Filter results on search | 
 **sort** | **[]string** | Sort the results. Can change order by using &#x60;&lt;sort_by&gt;:&lt;direction&gt;&#x60; where &#x60;&lt;direction&gt;&#x60; is either &#x60;asc&#x60; or &#x60;desc&#x60; | 
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]ProductWithIncludes**](ProductWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogProductById

> ProductWithIncludes UpdateCatalogProductById(ctx, id).UpdateCatalogProductByIdRequest(updateCatalogProductByIdRequest).Execute()

Update a catalog product

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
	updateCatalogProductByIdRequest := *openapiclient.NewUpdateCatalogProductByIdRequest() // UpdateCatalogProductByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogProductsAPI.UpdateCatalogProductById(context.Background(), id).UpdateCatalogProductByIdRequest(updateCatalogProductByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogProductsAPI.UpdateCatalogProductById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCatalogProductById`: ProductWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `CatalogProductsAPI.UpdateCatalogProductById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogProductByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCatalogProductByIdRequest** | [**UpdateCatalogProductByIdRequest**](UpdateCatalogProductByIdRequest.md) |  | 

### Return type

[**ProductWithIncludes**](ProductWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

