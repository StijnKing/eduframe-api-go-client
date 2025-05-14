# Payment

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

### NewPayment

`func NewPayment(id int32, invoiceId int32, amount string, currency Currency, paymentMethodId int32, status string, gatewayId NullableInt32, date string, ) *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Payment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v int32)`

SetId sets Id field to given value.


### GetInvoiceId

`func (o *Payment) GetInvoiceId() int32`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Payment) GetInvoiceIdOk() (*int32, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Payment) SetInvoiceId(v int32)`

SetInvoiceId sets InvoiceId field to given value.


### GetAmount

`func (o *Payment) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payment) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payment) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *Payment) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Payment) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Payment) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetPaymentMethodId

`func (o *Payment) GetPaymentMethodId() int32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *Payment) GetPaymentMethodIdOk() (*int32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *Payment) SetPaymentMethodId(v int32)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### GetStatus

`func (o *Payment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetGatewayId

`func (o *Payment) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *Payment) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *Payment) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.


### SetGatewayIdNil

`func (o *Payment) SetGatewayIdNil(b bool)`

 SetGatewayIdNil sets the value for GatewayId to be an explicit nil

### UnsetGatewayId
`func (o *Payment) UnsetGatewayId()`

UnsetGatewayId ensures that no value is present for GatewayId, not even an explicit nil
### GetDate

`func (o *Payment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Payment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Payment) SetDate(v string)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


