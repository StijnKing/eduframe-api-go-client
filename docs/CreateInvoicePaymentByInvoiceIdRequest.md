# CreateInvoicePaymentByInvoiceIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | A number representing the total amount of the invoice. | 
**Currency** | Pointer to **string** | The currency used for the payment. | [optional] 
**Date** | Pointer to **string** | Date on which the payment was created. | [optional] 
**PaymentMethodId** | Pointer to **int32** | Identifier of the payment method. | [optional] 

## Methods

### NewCreateInvoicePaymentByInvoiceIdRequest

`func NewCreateInvoicePaymentByInvoiceIdRequest(amount string, ) *CreateInvoicePaymentByInvoiceIdRequest`

NewCreateInvoicePaymentByInvoiceIdRequest instantiates a new CreateInvoicePaymentByInvoiceIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvoicePaymentByInvoiceIdRequestWithDefaults

`func NewCreateInvoicePaymentByInvoiceIdRequestWithDefaults() *CreateInvoicePaymentByInvoiceIdRequest`

NewCreateInvoicePaymentByInvoiceIdRequestWithDefaults instantiates a new CreateInvoicePaymentByInvoiceIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreateInvoicePaymentByInvoiceIdRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateInvoicePaymentByInvoiceIdRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateInvoicePaymentByInvoiceIdRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *CreateInvoicePaymentByInvoiceIdRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateInvoicePaymentByInvoiceIdRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateInvoicePaymentByInvoiceIdRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateInvoicePaymentByInvoiceIdRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDate

`func (o *CreateInvoicePaymentByInvoiceIdRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CreateInvoicePaymentByInvoiceIdRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CreateInvoicePaymentByInvoiceIdRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CreateInvoicePaymentByInvoiceIdRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *CreateInvoicePaymentByInvoiceIdRequest) GetPaymentMethodId() int32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *CreateInvoicePaymentByInvoiceIdRequest) GetPaymentMethodIdOk() (*int32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *CreateInvoicePaymentByInvoiceIdRequest) SetPaymentMethodId(v int32)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *CreateInvoicePaymentByInvoiceIdRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


