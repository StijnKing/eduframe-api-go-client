# \PaymentsAPI

All URIs are relative to *https://api.eduframe.nl/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoicePaymentByInvoiceId**](PaymentsAPI.md#CreateInvoicePaymentByInvoiceId) | **Post** /invoices/{invoice_id}/payments | Create a payment.
[**DeleteInvoicePaymentByIdAndInvoiceId**](PaymentsAPI.md#DeleteInvoicePaymentByIdAndInvoiceId) | **Delete** /invoices/{invoice_id}/payments/{id} | Delete a payment.
[**GetInvoicePaymentsByInvoiceId**](PaymentsAPI.md#GetInvoicePaymentsByInvoiceId) | **Get** /invoices/{invoice_id}/payments | Get all payment records of an invoice
[**GetPaymentById**](PaymentsAPI.md#GetPaymentById) | **Get** /payments/{id} | Get one payment record



## CreateInvoicePaymentByInvoiceId

> PaymentWithFixedIncludes CreateInvoicePaymentByInvoiceId(ctx, invoiceId).CreateInvoicePaymentByInvoiceIdRequest(createInvoicePaymentByInvoiceIdRequest).Execute()

Create a payment.

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
	invoiceId := "invoiceId_example" // string | Filter results on invoice_id
	createInvoicePaymentByInvoiceIdRequest := *openapiclient.NewCreateInvoicePaymentByInvoiceIdRequest("Amount_example") // CreateInvoicePaymentByInvoiceIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.CreateInvoicePaymentByInvoiceId(context.Background(), invoiceId).CreateInvoicePaymentByInvoiceIdRequest(createInvoicePaymentByInvoiceIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.CreateInvoicePaymentByInvoiceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoicePaymentByInvoiceId`: PaymentWithFixedIncludes
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.CreateInvoicePaymentByInvoiceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | Filter results on invoice_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoicePaymentByInvoiceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInvoicePaymentByInvoiceIdRequest** | [**CreateInvoicePaymentByInvoiceIdRequest**](CreateInvoicePaymentByInvoiceIdRequest.md) |  | 

### Return type

[**PaymentWithFixedIncludes**](PaymentWithFixedIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvoicePaymentByIdAndInvoiceId

> DeleteInvoicePaymentByIdAndInvoiceId(ctx, id, invoiceId).Execute()

Delete a payment.

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
	invoiceId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PaymentsAPI.DeleteInvoicePaymentByIdAndInvoiceId(context.Background(), id, invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.DeleteInvoicePaymentByIdAndInvoiceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**invoiceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvoicePaymentByIdAndInvoiceIdRequest struct via the builder pattern


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


## GetInvoicePaymentsByInvoiceId

> []PaymentWithIncludes GetInvoicePaymentsByInvoiceId(ctx, invoiceId).Execute()

Get all payment records of an invoice

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
	invoiceId := "invoiceId_example" // string | Filter results on invoice_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.GetInvoicePaymentsByInvoiceId(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetInvoicePaymentsByInvoiceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoicePaymentsByInvoiceId`: []PaymentWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetInvoicePaymentsByInvoiceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | Filter results on invoice_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicePaymentsByInvoiceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PaymentWithIncludes**](PaymentWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentById

> PaymentWithIncludes GetPaymentById(ctx, id).Execute()

Get one payment record

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
	resp, r, err := apiClient.PaymentsAPI.GetPaymentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetPaymentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentById`: PaymentWithIncludes
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetPaymentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentWithIncludes**](PaymentWithIncludes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

