# CreateInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int32** | Identifier of the account. | 
**Feature** | Pointer to **NullableString** | Some description of the invoice which is displayed on the invoice. | [optional] 
**Footnote** | Pointer to **NullableString** | The note displayed at the bottom of the invoice. | [optional] 
**InvoiceItemsAttributes** | Pointer to [**[]CreateInvoiceRequestInvoiceItemsAttributesInner**](CreateInvoiceRequestInvoiceItemsAttributesInner.md) |  | [optional] 

## Methods

### NewCreateInvoiceRequest

`func NewCreateInvoiceRequest(accountId int32, ) *CreateInvoiceRequest`

NewCreateInvoiceRequest instantiates a new CreateInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvoiceRequestWithDefaults

`func NewCreateInvoiceRequestWithDefaults() *CreateInvoiceRequest`

NewCreateInvoiceRequestWithDefaults instantiates a new CreateInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateInvoiceRequest) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateInvoiceRequest) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateInvoiceRequest) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetFeature

`func (o *CreateInvoiceRequest) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *CreateInvoiceRequest) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *CreateInvoiceRequest) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *CreateInvoiceRequest) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### SetFeatureNil

`func (o *CreateInvoiceRequest) SetFeatureNil(b bool)`

 SetFeatureNil sets the value for Feature to be an explicit nil

### UnsetFeature
`func (o *CreateInvoiceRequest) UnsetFeature()`

UnsetFeature ensures that no value is present for Feature, not even an explicit nil
### GetFootnote

`func (o *CreateInvoiceRequest) GetFootnote() string`

GetFootnote returns the Footnote field if non-nil, zero value otherwise.

### GetFootnoteOk

`func (o *CreateInvoiceRequest) GetFootnoteOk() (*string, bool)`

GetFootnoteOk returns a tuple with the Footnote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFootnote

`func (o *CreateInvoiceRequest) SetFootnote(v string)`

SetFootnote sets Footnote field to given value.

### HasFootnote

`func (o *CreateInvoiceRequest) HasFootnote() bool`

HasFootnote returns a boolean if a field has been set.

### SetFootnoteNil

`func (o *CreateInvoiceRequest) SetFootnoteNil(b bool)`

 SetFootnoteNil sets the value for Footnote to be an explicit nil

### UnsetFootnote
`func (o *CreateInvoiceRequest) UnsetFootnote()`

UnsetFootnote ensures that no value is present for Footnote, not even an explicit nil
### GetInvoiceItemsAttributes

`func (o *CreateInvoiceRequest) GetInvoiceItemsAttributes() []CreateInvoiceRequestInvoiceItemsAttributesInner`

GetInvoiceItemsAttributes returns the InvoiceItemsAttributes field if non-nil, zero value otherwise.

### GetInvoiceItemsAttributesOk

`func (o *CreateInvoiceRequest) GetInvoiceItemsAttributesOk() (*[]CreateInvoiceRequestInvoiceItemsAttributesInner, bool)`

GetInvoiceItemsAttributesOk returns a tuple with the InvoiceItemsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemsAttributes

`func (o *CreateInvoiceRequest) SetInvoiceItemsAttributes(v []CreateInvoiceRequestInvoiceItemsAttributesInner)`

SetInvoiceItemsAttributes sets InvoiceItemsAttributes field to given value.

### HasInvoiceItemsAttributes

`func (o *CreateInvoiceRequest) HasInvoiceItemsAttributes() bool`

HasInvoiceItemsAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


