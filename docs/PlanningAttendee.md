# PlanningAttendee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the planning attendee | [readonly] 
**AttendableId** | **int32** | Unique identifier of the object that is attended. | 
**AttendableType** | **string** | Type of the object that is attended. | 
**TeacherId** | **int32** | Unique identifier of the teacher. | 
**TeacherRoleId** | Pointer to **int32** | Unique identifier of the teacher role. | [optional] 
**Comment** | **NullableString** | Comment on the teacher assignment | 

## Methods

### NewPlanningAttendee

`func NewPlanningAttendee(id int32, attendableId int32, attendableType string, teacherId int32, comment NullableString, ) *PlanningAttendee`

NewPlanningAttendee instantiates a new PlanningAttendee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanningAttendeeWithDefaults

`func NewPlanningAttendeeWithDefaults() *PlanningAttendee`

NewPlanningAttendeeWithDefaults instantiates a new PlanningAttendee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanningAttendee) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanningAttendee) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanningAttendee) SetId(v int32)`

SetId sets Id field to given value.


### GetAttendableId

`func (o *PlanningAttendee) GetAttendableId() int32`

GetAttendableId returns the AttendableId field if non-nil, zero value otherwise.

### GetAttendableIdOk

`func (o *PlanningAttendee) GetAttendableIdOk() (*int32, bool)`

GetAttendableIdOk returns a tuple with the AttendableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendableId

`func (o *PlanningAttendee) SetAttendableId(v int32)`

SetAttendableId sets AttendableId field to given value.


### GetAttendableType

`func (o *PlanningAttendee) GetAttendableType() string`

GetAttendableType returns the AttendableType field if non-nil, zero value otherwise.

### GetAttendableTypeOk

`func (o *PlanningAttendee) GetAttendableTypeOk() (*string, bool)`

GetAttendableTypeOk returns a tuple with the AttendableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendableType

`func (o *PlanningAttendee) SetAttendableType(v string)`

SetAttendableType sets AttendableType field to given value.


### GetTeacherId

`func (o *PlanningAttendee) GetTeacherId() int32`

GetTeacherId returns the TeacherId field if non-nil, zero value otherwise.

### GetTeacherIdOk

`func (o *PlanningAttendee) GetTeacherIdOk() (*int32, bool)`

GetTeacherIdOk returns a tuple with the TeacherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherId

`func (o *PlanningAttendee) SetTeacherId(v int32)`

SetTeacherId sets TeacherId field to given value.


### GetTeacherRoleId

`func (o *PlanningAttendee) GetTeacherRoleId() int32`

GetTeacherRoleId returns the TeacherRoleId field if non-nil, zero value otherwise.

### GetTeacherRoleIdOk

`func (o *PlanningAttendee) GetTeacherRoleIdOk() (*int32, bool)`

GetTeacherRoleIdOk returns a tuple with the TeacherRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherRoleId

`func (o *PlanningAttendee) SetTeacherRoleId(v int32)`

SetTeacherRoleId sets TeacherRoleId field to given value.

### HasTeacherRoleId

`func (o *PlanningAttendee) HasTeacherRoleId() bool`

HasTeacherRoleId returns a boolean if a field has been set.

### GetComment

`func (o *PlanningAttendee) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PlanningAttendee) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PlanningAttendee) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *PlanningAttendee) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *PlanningAttendee) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


