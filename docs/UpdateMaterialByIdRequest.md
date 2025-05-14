# UpdateMaterialByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the material. | [optional] 
**MaterialGroupId** | Pointer to **int32** | Unique identifier of the material group. | [optional] 

## Methods

### NewUpdateMaterialByIdRequest

`func NewUpdateMaterialByIdRequest() *UpdateMaterialByIdRequest`

NewUpdateMaterialByIdRequest instantiates a new UpdateMaterialByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMaterialByIdRequestWithDefaults

`func NewUpdateMaterialByIdRequestWithDefaults() *UpdateMaterialByIdRequest`

NewUpdateMaterialByIdRequestWithDefaults instantiates a new UpdateMaterialByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMaterialByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMaterialByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMaterialByIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMaterialByIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMaterialGroupId

`func (o *UpdateMaterialByIdRequest) GetMaterialGroupId() int32`

GetMaterialGroupId returns the MaterialGroupId field if non-nil, zero value otherwise.

### GetMaterialGroupIdOk

`func (o *UpdateMaterialByIdRequest) GetMaterialGroupIdOk() (*int32, bool)`

GetMaterialGroupIdOk returns a tuple with the MaterialGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialGroupId

`func (o *UpdateMaterialByIdRequest) SetMaterialGroupId(v int32)`

SetMaterialGroupId sets MaterialGroupId field to given value.

### HasMaterialGroupId

`func (o *UpdateMaterialByIdRequest) HasMaterialGroupId() bool`

HasMaterialGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


