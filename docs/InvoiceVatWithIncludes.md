# InvoiceVatWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique idenitfier of the invoice vat. | [readonly] 
**Name** | **string** | Name of the invoice vat. | 
**Percentage** | **string** | Number representing the VAT percentage. | 
**HasVat** | **bool** | Boolean indicating whether the invoice vat has VAT. | 

## Methods

### NewInvoiceVatWithIncludes

`func NewInvoiceVatWithIncludes(id int32, name string, percentage string, hasVat bool, ) *InvoiceVatWithIncludes`

NewInvoiceVatWithIncludes instantiates a new InvoiceVatWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceVatWithIncludesWithDefaults

`func NewInvoiceVatWithIncludesWithDefaults() *InvoiceVatWithIncludes`

NewInvoiceVatWithIncludesWithDefaults instantiates a new InvoiceVatWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceVatWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceVatWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceVatWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *InvoiceVatWithIncludes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceVatWithIncludes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceVatWithIncludes) SetName(v string)`

SetName sets Name field to given value.


### GetPercentage

`func (o *InvoiceVatWithIncludes) GetPercentage() string`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *InvoiceVatWithIncludes) GetPercentageOk() (*string, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *InvoiceVatWithIncludes) SetPercentage(v string)`

SetPercentage sets Percentage field to given value.


### GetHasVat

`func (o *InvoiceVatWithIncludes) GetHasVat() bool`

GetHasVat returns the HasVat field if non-nil, zero value otherwise.

### GetHasVatOk

`func (o *InvoiceVatWithIncludes) GetHasVatOk() (*bool, bool)`

GetHasVatOk returns a tuple with the HasVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVat

`func (o *InvoiceVatWithIncludes) SetHasVat(v bool)`

SetHasVat sets HasVat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


