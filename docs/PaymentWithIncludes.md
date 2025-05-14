# PaymentWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the payment. | [readonly] 
**InvoiceId** | **int32** | Unique identifier of the linked invoice. | 
**Amount** | **string** | A number representing the total amount of the invoice. | 
**Currency** | [**Currency**](Currency.md) |  | 
**PaymentMethodId** | **int32** | Identifier of the payment method. | 
**Status** | **string** | The status of the payment. | 
**GatewayId** | **NullableInt32** | The payment provider payment id. | 
**Date** | **string** | Date on which the payment was created. | 

## Methods

### NewPaymentWithIncludes

`func NewPaymentWithIncludes(id int32, invoiceId int32, amount string, currency Currency, paymentMethodId int32, status string, gatewayId NullableInt32, date string, ) *PaymentWithIncludes`

NewPaymentWithIncludes instantiates a new PaymentWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithIncludesWithDefaults

`func NewPaymentWithIncludesWithDefaults() *PaymentWithIncludes`

NewPaymentWithIncludesWithDefaults instantiates a new PaymentWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetInvoiceId

`func (o *PaymentWithIncludes) GetInvoiceId() int32`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaymentWithIncludes) GetInvoiceIdOk() (*int32, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaymentWithIncludes) SetInvoiceId(v int32)`

SetInvoiceId sets InvoiceId field to given value.


### GetAmount

`func (o *PaymentWithIncludes) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentWithIncludes) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentWithIncludes) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *PaymentWithIncludes) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentWithIncludes) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentWithIncludes) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetPaymentMethodId

`func (o *PaymentWithIncludes) GetPaymentMethodId() int32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *PaymentWithIncludes) GetPaymentMethodIdOk() (*int32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *PaymentWithIncludes) SetPaymentMethodId(v int32)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### GetStatus

`func (o *PaymentWithIncludes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentWithIncludes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentWithIncludes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetGatewayId

`func (o *PaymentWithIncludes) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *PaymentWithIncludes) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *PaymentWithIncludes) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.


### SetGatewayIdNil

`func (o *PaymentWithIncludes) SetGatewayIdNil(b bool)`

 SetGatewayIdNil sets the value for GatewayId to be an explicit nil

### UnsetGatewayId
`func (o *PaymentWithIncludes) UnsetGatewayId()`

UnsetGatewayId ensures that no value is present for GatewayId, not even an explicit nil
### GetDate

`func (o *PaymentWithIncludes) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PaymentWithIncludes) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PaymentWithIncludes) SetDate(v string)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


