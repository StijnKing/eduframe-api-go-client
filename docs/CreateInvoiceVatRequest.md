# CreateInvoiceVatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the invoice vat. | 
**Percentage** | **string** | Number representing the VAT percentage. | 

## Methods

### NewCreateInvoiceVatRequest

`func NewCreateInvoiceVatRequest(name string, percentage string, ) *CreateInvoiceVatRequest`

NewCreateInvoiceVatRequest instantiates a new CreateInvoiceVatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvoiceVatRequestWithDefaults

`func NewCreateInvoiceVatRequestWithDefaults() *CreateInvoiceVatRequest`

NewCreateInvoiceVatRequestWithDefaults instantiates a new CreateInvoiceVatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateInvoiceVatRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInvoiceVatRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInvoiceVatRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPercentage

`func (o *CreateInvoiceVatRequest) GetPercentage() string`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *CreateInvoiceVatRequest) GetPercentageOk() (*string, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *CreateInvoiceVatRequest) SetPercentage(v string)`

SetPercentage sets Percentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


