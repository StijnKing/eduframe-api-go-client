# Material

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the material | [readonly] 
**Name** | **string** | Name of the material. | 
**UseType** | **string** | Type of material. | 
**MaterialGroupId** | **int32** | Unique identifier of the material group. | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewMaterial

`func NewMaterial(id int32, name string, useType string, materialGroupId int32, updatedAt time.Time, createdAt time.Time, ) *Material`

NewMaterial instantiates a new Material object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaterialWithDefaults

`func NewMaterialWithDefaults() *Material`

NewMaterialWithDefaults instantiates a new Material object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Material) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Material) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Material) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Material) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Material) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Material) SetName(v string)`

SetName sets Name field to given value.


### GetUseType

`func (o *Material) GetUseType() string`

GetUseType returns the UseType field if non-nil, zero value otherwise.

### GetUseTypeOk

`func (o *Material) GetUseTypeOk() (*string, bool)`

GetUseTypeOk returns a tuple with the UseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseType

`func (o *Material) SetUseType(v string)`

SetUseType sets UseType field to given value.


### GetMaterialGroupId

`func (o *Material) GetMaterialGroupId() int32`

GetMaterialGroupId returns the MaterialGroupId field if non-nil, zero value otherwise.

### GetMaterialGroupIdOk

`func (o *Material) GetMaterialGroupIdOk() (*int32, bool)`

GetMaterialGroupIdOk returns a tuple with the MaterialGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialGroupId

`func (o *Material) SetMaterialGroupId(v int32)`

SetMaterialGroupId sets MaterialGroupId field to given value.


### GetUpdatedAt

`func (o *Material) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Material) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Material) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Material) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Material) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Material) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


