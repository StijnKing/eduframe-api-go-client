# UpdateTeacherEnrollmentByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlannedCourseId** | Pointer to **int32** | Unique identifier of the planned course. | [optional] 
**TeacherId** | Pointer to **int32** | Unique identifier of the teacher. | [optional] 
**TeacherRoleId** | Pointer to **int32** | Unique identifier of the teacher role.  ![Teacher Roles](https://img.shields.io/badge/Feature-Teacher_Roles-blue)  | [optional] 

## Methods

### NewUpdateTeacherEnrollmentByIdRequest

`func NewUpdateTeacherEnrollmentByIdRequest() *UpdateTeacherEnrollmentByIdRequest`

NewUpdateTeacherEnrollmentByIdRequest instantiates a new UpdateTeacherEnrollmentByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeacherEnrollmentByIdRequestWithDefaults

`func NewUpdateTeacherEnrollmentByIdRequestWithDefaults() *UpdateTeacherEnrollmentByIdRequest`

NewUpdateTeacherEnrollmentByIdRequestWithDefaults instantiates a new UpdateTeacherEnrollmentByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlannedCourseId

`func (o *UpdateTeacherEnrollmentByIdRequest) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *UpdateTeacherEnrollmentByIdRequest) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *UpdateTeacherEnrollmentByIdRequest) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.

### HasPlannedCourseId

`func (o *UpdateTeacherEnrollmentByIdRequest) HasPlannedCourseId() bool`

HasPlannedCourseId returns a boolean if a field has been set.

### GetTeacherId

`func (o *UpdateTeacherEnrollmentByIdRequest) GetTeacherId() int32`

GetTeacherId returns the TeacherId field if non-nil, zero value otherwise.

### GetTeacherIdOk

`func (o *UpdateTeacherEnrollmentByIdRequest) GetTeacherIdOk() (*int32, bool)`

GetTeacherIdOk returns a tuple with the TeacherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherId

`func (o *UpdateTeacherEnrollmentByIdRequest) SetTeacherId(v int32)`

SetTeacherId sets TeacherId field to given value.

### HasTeacherId

`func (o *UpdateTeacherEnrollmentByIdRequest) HasTeacherId() bool`

HasTeacherId returns a boolean if a field has been set.

### GetTeacherRoleId

`func (o *UpdateTeacherEnrollmentByIdRequest) GetTeacherRoleId() int32`

GetTeacherRoleId returns the TeacherRoleId field if non-nil, zero value otherwise.

### GetTeacherRoleIdOk

`func (o *UpdateTeacherEnrollmentByIdRequest) GetTeacherRoleIdOk() (*int32, bool)`

GetTeacherRoleIdOk returns a tuple with the TeacherRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherRoleId

`func (o *UpdateTeacherEnrollmentByIdRequest) SetTeacherRoleId(v int32)`

SetTeacherRoleId sets TeacherRoleId field to given value.

### HasTeacherRoleId

`func (o *UpdateTeacherEnrollmentByIdRequest) HasTeacherRoleId() bool`

HasTeacherRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


