# CreateMaterialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the material. | 
**UseType** | Pointer to **string** | Type of material. | [optional] 
**MaterialGroupId** | **int32** | Unique identifier of the material group. | 

## Methods

### NewCreateMaterialRequest

`func NewCreateMaterialRequest(name string, materialGroupId int32, ) *CreateMaterialRequest`

NewCreateMaterialRequest instantiates a new CreateMaterialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMaterialRequestWithDefaults

`func NewCreateMaterialRequestWithDefaults() *CreateMaterialRequest`

NewCreateMaterialRequestWithDefaults instantiates a new CreateMaterialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMaterialRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMaterialRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMaterialRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUseType

`func (o *CreateMaterialRequest) GetUseType() string`

GetUseType returns the UseType field if non-nil, zero value otherwise.

### GetUseTypeOk

`func (o *CreateMaterialRequest) GetUseTypeOk() (*string, bool)`

GetUseTypeOk returns a tuple with the UseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseType

`func (o *CreateMaterialRequest) SetUseType(v string)`

SetUseType sets UseType field to given value.

### HasUseType

`func (o *CreateMaterialRequest) HasUseType() bool`

HasUseType returns a boolean if a field has been set.

### GetMaterialGroupId

`func (o *CreateMaterialRequest) GetMaterialGroupId() int32`

GetMaterialGroupId returns the MaterialGroupId field if non-nil, zero value otherwise.

### GetMaterialGroupIdOk

`func (o *CreateMaterialRequest) GetMaterialGroupIdOk() (*int32, bool)`

GetMaterialGroupIdOk returns a tuple with the MaterialGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialGroupId

`func (o *CreateMaterialRequest) SetMaterialGroupId(v int32)`

SetMaterialGroupId sets MaterialGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


