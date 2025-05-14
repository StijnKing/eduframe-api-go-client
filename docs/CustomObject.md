# CustomObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique identifier of the custom object | 
**NameSingular** | **string** | The singular name of the custom object | 
**NamePlural** | **string** | The plural name of the custom object | 

## Methods

### NewCustomObject

`func NewCustomObject(id int32, nameSingular string, namePlural string, ) *CustomObject`

NewCustomObject instantiates a new CustomObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomObjectWithDefaults

`func NewCustomObjectWithDefaults() *CustomObject`

NewCustomObjectWithDefaults instantiates a new CustomObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomObject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomObject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomObject) SetId(v int32)`

SetId sets Id field to given value.


### GetNameSingular

`func (o *CustomObject) GetNameSingular() string`

GetNameSingular returns the NameSingular field if non-nil, zero value otherwise.

### GetNameSingularOk

`func (o *CustomObject) GetNameSingularOk() (*string, bool)`

GetNameSingularOk returns a tuple with the NameSingular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSingular

`func (o *CustomObject) SetNameSingular(v string)`

SetNameSingular sets NameSingular field to given value.


### GetNamePlural

`func (o *CustomObject) GetNamePlural() string`

GetNamePlural returns the NamePlural field if non-nil, zero value otherwise.

### GetNamePluralOk

`func (o *CustomObject) GetNamePluralOk() (*string, bool)`

GetNamePluralOk returns a tuple with the NamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePlural

`func (o *CustomObject) SetNamePlural(v string)`

SetNamePlural sets NamePlural field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


