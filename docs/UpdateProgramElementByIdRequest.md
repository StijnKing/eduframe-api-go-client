# UpdateProgramElementByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EditionId** | Pointer to **int32** | The identifier of the associated course. | [optional] 
**PlannedCourseId** | Pointer to **NullableInt32** | The identifier of the associated course. | [optional] 

## Methods

### NewUpdateProgramElementByIdRequest

`func NewUpdateProgramElementByIdRequest() *UpdateProgramElementByIdRequest`

NewUpdateProgramElementByIdRequest instantiates a new UpdateProgramElementByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProgramElementByIdRequestWithDefaults

`func NewUpdateProgramElementByIdRequestWithDefaults() *UpdateProgramElementByIdRequest`

NewUpdateProgramElementByIdRequestWithDefaults instantiates a new UpdateProgramElementByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditionId

`func (o *UpdateProgramElementByIdRequest) GetEditionId() int32`

GetEditionId returns the EditionId field if non-nil, zero value otherwise.

### GetEditionIdOk

`func (o *UpdateProgramElementByIdRequest) GetEditionIdOk() (*int32, bool)`

GetEditionIdOk returns a tuple with the EditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionId

`func (o *UpdateProgramElementByIdRequest) SetEditionId(v int32)`

SetEditionId sets EditionId field to given value.

### HasEditionId

`func (o *UpdateProgramElementByIdRequest) HasEditionId() bool`

HasEditionId returns a boolean if a field has been set.

### GetPlannedCourseId

`func (o *UpdateProgramElementByIdRequest) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *UpdateProgramElementByIdRequest) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *UpdateProgramElementByIdRequest) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.

### HasPlannedCourseId

`func (o *UpdateProgramElementByIdRequest) HasPlannedCourseId() bool`

HasPlannedCourseId returns a boolean if a field has been set.

### SetPlannedCourseIdNil

`func (o *UpdateProgramElementByIdRequest) SetPlannedCourseIdNil(b bool)`

 SetPlannedCourseIdNil sets the value for PlannedCourseId to be an explicit nil

### UnsetPlannedCourseId
`func (o *UpdateProgramElementByIdRequest) UnsetPlannedCourseId()`

UnsetPlannedCourseId ensures that no value is present for PlannedCourseId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


