# CreateTeacherEnrollmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlannedCourseId** | **int32** | Unique identifier of the planned course. | 
**TeacherId** | **int32** | Unique identifier of the teacher. | 
**TeacherRoleId** | Pointer to **int32** | Unique identifier of the teacher role.  ![Teacher Roles](https://img.shields.io/badge/Feature-Teacher_Roles-blue)  | [optional] 

## Methods

### NewCreateTeacherEnrollmentRequest

`func NewCreateTeacherEnrollmentRequest(plannedCourseId int32, teacherId int32, ) *CreateTeacherEnrollmentRequest`

NewCreateTeacherEnrollmentRequest instantiates a new CreateTeacherEnrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTeacherEnrollmentRequestWithDefaults

`func NewCreateTeacherEnrollmentRequestWithDefaults() *CreateTeacherEnrollmentRequest`

NewCreateTeacherEnrollmentRequestWithDefaults instantiates a new CreateTeacherEnrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlannedCourseId

`func (o *CreateTeacherEnrollmentRequest) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *CreateTeacherEnrollmentRequest) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *CreateTeacherEnrollmentRequest) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.


### GetTeacherId

`func (o *CreateTeacherEnrollmentRequest) GetTeacherId() int32`

GetTeacherId returns the TeacherId field if non-nil, zero value otherwise.

### GetTeacherIdOk

`func (o *CreateTeacherEnrollmentRequest) GetTeacherIdOk() (*int32, bool)`

GetTeacherIdOk returns a tuple with the TeacherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherId

`func (o *CreateTeacherEnrollmentRequest) SetTeacherId(v int32)`

SetTeacherId sets TeacherId field to given value.


### GetTeacherRoleId

`func (o *CreateTeacherEnrollmentRequest) GetTeacherRoleId() int32`

GetTeacherRoleId returns the TeacherRoleId field if non-nil, zero value otherwise.

### GetTeacherRoleIdOk

`func (o *CreateTeacherEnrollmentRequest) GetTeacherRoleIdOk() (*int32, bool)`

GetTeacherRoleIdOk returns a tuple with the TeacherRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherRoleId

`func (o *CreateTeacherEnrollmentRequest) SetTeacherRoleId(v int32)`

SetTeacherRoleId sets TeacherRoleId field to given value.

### HasTeacherRoleId

`func (o *CreateTeacherEnrollmentRequest) HasTeacherRoleId() bool`

HasTeacherRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


