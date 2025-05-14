# PlanningMaterial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaterialId** | **int32** | Unique identifier of the material | 
**Amount** | **NullableInt32** | Amount of the material | 
**Comment** | **NullableString** | Comment on the planning material | 

## Methods

### NewPlanningMaterial

`func NewPlanningMaterial(materialId int32, amount NullableInt32, comment NullableString, ) *PlanningMaterial`

NewPlanningMaterial instantiates a new PlanningMaterial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanningMaterialWithDefaults

`func NewPlanningMaterialWithDefaults() *PlanningMaterial`

NewPlanningMaterialWithDefaults instantiates a new PlanningMaterial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaterialId

`func (o *PlanningMaterial) GetMaterialId() int32`

GetMaterialId returns the MaterialId field if non-nil, zero value otherwise.

### GetMaterialIdOk

`func (o *PlanningMaterial) GetMaterialIdOk() (*int32, bool)`

GetMaterialIdOk returns a tuple with the MaterialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialId

`func (o *PlanningMaterial) SetMaterialId(v int32)`

SetMaterialId sets MaterialId field to given value.


### GetAmount

`func (o *PlanningMaterial) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PlanningMaterial) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PlanningMaterial) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### SetAmountNil

`func (o *PlanningMaterial) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *PlanningMaterial) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetComment

`func (o *PlanningMaterial) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PlanningMaterial) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PlanningMaterial) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *PlanningMaterial) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *PlanningMaterial) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


