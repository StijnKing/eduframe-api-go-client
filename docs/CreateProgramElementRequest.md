# CreateProgramElementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CourseId** | **int32** | The identifier of the associated course. | 
**EditionId** | **int32** | The identifier of the associated course. | 
**PlannedCourseId** | Pointer to **NullableInt32** | The identifier of the associated course. | [optional] 

## Methods

### NewCreateProgramElementRequest

`func NewCreateProgramElementRequest(courseId int32, editionId int32, ) *CreateProgramElementRequest`

NewCreateProgramElementRequest instantiates a new CreateProgramElementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProgramElementRequestWithDefaults

`func NewCreateProgramElementRequestWithDefaults() *CreateProgramElementRequest`

NewCreateProgramElementRequestWithDefaults instantiates a new CreateProgramElementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCourseId

`func (o *CreateProgramElementRequest) GetCourseId() int32`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CreateProgramElementRequest) GetCourseIdOk() (*int32, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CreateProgramElementRequest) SetCourseId(v int32)`

SetCourseId sets CourseId field to given value.


### GetEditionId

`func (o *CreateProgramElementRequest) GetEditionId() int32`

GetEditionId returns the EditionId field if non-nil, zero value otherwise.

### GetEditionIdOk

`func (o *CreateProgramElementRequest) GetEditionIdOk() (*int32, bool)`

GetEditionIdOk returns a tuple with the EditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionId

`func (o *CreateProgramElementRequest) SetEditionId(v int32)`

SetEditionId sets EditionId field to given value.


### GetPlannedCourseId

`func (o *CreateProgramElementRequest) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *CreateProgramElementRequest) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *CreateProgramElementRequest) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.

### HasPlannedCourseId

`func (o *CreateProgramElementRequest) HasPlannedCourseId() bool`

HasPlannedCourseId returns a boolean if a field has been set.

### SetPlannedCourseIdNil

`func (o *CreateProgramElementRequest) SetPlannedCourseIdNil(b bool)`

 SetPlannedCourseIdNil sets the value for PlannedCourseId to be an explicit nil

### UnsetPlannedCourseId
`func (o *CreateProgramElementRequest) UnsetPlannedCourseId()`

UnsetPlannedCourseId ensures that no value is present for PlannedCourseId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


