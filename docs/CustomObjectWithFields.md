# CustomObjectWithFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique identifier of the custom object | 
**NameSingular** | **string** | The singular name of the custom object | 
**NamePlural** | **string** | The plural name of the custom object | 
**Fields** | [**[]CustomObjectField**](CustomObjectField.md) |  | 

## Methods

### NewCustomObjectWithFields

`func NewCustomObjectWithFields(id int32, nameSingular string, namePlural string, fields []CustomObjectField, ) *CustomObjectWithFields`

NewCustomObjectWithFields instantiates a new CustomObjectWithFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomObjectWithFieldsWithDefaults

`func NewCustomObjectWithFieldsWithDefaults() *CustomObjectWithFields`

NewCustomObjectWithFieldsWithDefaults instantiates a new CustomObjectWithFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomObjectWithFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomObjectWithFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomObjectWithFields) SetId(v int32)`

SetId sets Id field to given value.


### GetNameSingular

`func (o *CustomObjectWithFields) GetNameSingular() string`

GetNameSingular returns the NameSingular field if non-nil, zero value otherwise.

### GetNameSingularOk

`func (o *CustomObjectWithFields) GetNameSingularOk() (*string, bool)`

GetNameSingularOk returns a tuple with the NameSingular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSingular

`func (o *CustomObjectWithFields) SetNameSingular(v string)`

SetNameSingular sets NameSingular field to given value.


### GetNamePlural

`func (o *CustomObjectWithFields) GetNamePlural() string`

GetNamePlural returns the NamePlural field if non-nil, zero value otherwise.

### GetNamePluralOk

`func (o *CustomObjectWithFields) GetNamePluralOk() (*string, bool)`

GetNamePluralOk returns a tuple with the NamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePlural

`func (o *CustomObjectWithFields) SetNamePlural(v string)`

SetNamePlural sets NamePlural field to given value.


### GetFields

`func (o *CustomObjectWithFields) GetFields() []CustomObjectField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomObjectWithFields) GetFieldsOk() (*[]CustomObjectField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomObjectWithFields) SetFields(v []CustomObjectField)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


