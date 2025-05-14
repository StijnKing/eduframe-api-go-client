# CreditCategoryWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the credit category. | [readonly] 
**Name** | **string** | Name of the credit category. | 
**Types** | [**[]CreditType**](CreditType.md) |  | 

## Methods

### NewCreditCategoryWithIncludes

`func NewCreditCategoryWithIncludes(id int32, name string, types []CreditType, ) *CreditCategoryWithIncludes`

NewCreditCategoryWithIncludes instantiates a new CreditCategoryWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCategoryWithIncludesWithDefaults

`func NewCreditCategoryWithIncludesWithDefaults() *CreditCategoryWithIncludes`

NewCreditCategoryWithIncludesWithDefaults instantiates a new CreditCategoryWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditCategoryWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditCategoryWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditCategoryWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CreditCategoryWithIncludes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreditCategoryWithIncludes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreditCategoryWithIncludes) SetName(v string)`

SetName sets Name field to given value.


### GetTypes

`func (o *CreditCategoryWithIncludes) GetTypes() []CreditType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *CreditCategoryWithIncludes) GetTypesOk() (*[]CreditType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *CreditCategoryWithIncludes) SetTypes(v []CreditType)`

SetTypes sets Types field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


