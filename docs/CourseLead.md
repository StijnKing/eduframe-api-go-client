# CourseLead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the course lead | [readonly] 
**LeadId** | **int32** | Unique identifier of the lead | 
**CourseId** | **int32** | ID of the course | 
**PlannedCourseId** | **NullableInt32** | ID of the planned course | 

## Methods

### NewCourseLead

`func NewCourseLead(id int32, leadId int32, courseId int32, plannedCourseId NullableInt32, ) *CourseLead`

NewCourseLead instantiates a new CourseLead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseLeadWithDefaults

`func NewCourseLeadWithDefaults() *CourseLead`

NewCourseLeadWithDefaults instantiates a new CourseLead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseLead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseLead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseLead) SetId(v int32)`

SetId sets Id field to given value.


### GetLeadId

`func (o *CourseLead) GetLeadId() int32`

GetLeadId returns the LeadId field if non-nil, zero value otherwise.

### GetLeadIdOk

`func (o *CourseLead) GetLeadIdOk() (*int32, bool)`

GetLeadIdOk returns a tuple with the LeadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadId

`func (o *CourseLead) SetLeadId(v int32)`

SetLeadId sets LeadId field to given value.


### GetCourseId

`func (o *CourseLead) GetCourseId() int32`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseLead) GetCourseIdOk() (*int32, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseLead) SetCourseId(v int32)`

SetCourseId sets CourseId field to given value.


### GetPlannedCourseId

`func (o *CourseLead) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *CourseLead) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *CourseLead) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.


### SetPlannedCourseIdNil

`func (o *CourseLead) SetPlannedCourseIdNil(b bool)`

 SetPlannedCourseIdNil sets the value for PlannedCourseId to be an explicit nil

### UnsetPlannedCourseId
`func (o *CourseLead) UnsetPlannedCourseId()`

UnsetPlannedCourseId ensures that no value is present for PlannedCourseId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


