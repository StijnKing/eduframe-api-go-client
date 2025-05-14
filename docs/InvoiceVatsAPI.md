# \InvoiceVatsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoiceVat**](InvoiceVatsAPI.md#CreateInvoiceVat) | **Post** /invoice_vats | Create an invoice vat.
[**GetInvoiceVats**](InvoiceVatsAPI.md#GetInvoiceVats) | **Get** /invoice_vats | Get all invoice vat records



## CreateInvoiceVat

> InvoiceVatWithIncludes CreateInvoiceVat(ctx).CreateInvoiceVatRequest(createInvoiceVatRequest).Execute()

Create an invoice vat.

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
	createInvoiceVatRequest := *openapiclient.NewCreateInvoiceVatRequest("Name_example", "Percentage_example") // CreateInvoiceVatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceVatsAPI.CreateInvoiceVat(context.Background()).CreateInvoiceVatRequest(createInvoiceVatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceVatsAPI.CreateInvoiceVat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoiceVat`: InvoiceVatWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `InvoiceVatsAPI.CreateInvoiceVat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceVatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInvoiceVatRequest** | [**CreateInvoiceVatRequest**](CreateInvoiceVatRequest.md) |  | 

### Return type

[**InvoiceVatWithIncludes**](InvoiceVatWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceVats

> []InvoiceVatWithIncludes GetInvoiceVats(ctx).Cursor(cursor).PerPage(perPage).Execute()

Get all invoice vat records

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
	cursor := "cursor_example" // string | **Note**: It's almost never necessary to use this parameter directly, the URL   should be retrieved from the `Link` header.  The cursor used to fetch the next result set.  (optional)
	perPage := int32(10) // int32 | The number of results to retrieve for this page. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceVatsAPI.GetInvoiceVats(context.Background()).Cursor(cursor).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceVatsAPI.GetInvoiceVats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceVats`: []InvoiceVatWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `InvoiceVatsAPI.GetInvoiceVats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceVatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | **Note**: It&#39;s almost never necessary to use this parameter directly, the URL   should be retrieved from the &#x60;Link&#x60; header.  The cursor used to fetch the next result set.  | 
 **perPage** | **int32** | The number of results to retrieve for this page. | [default to 25]

### Return type

[**[]InvoiceVatWithIncludes**](InvoiceVatWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

