# CreateTeacherEnrollmentByPlannedCourseIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeacherId** | **int32** | Unique identifier of the teacher. | 
**TeacherRoleId** | Pointer to **int32** | Unique identifier of the teacher role. | [optional] 

## Methods

### NewCreateTeacherEnrollmentByPlannedCourseIdRequest

`func NewCreateTeacherEnrollmentByPlannedCourseIdRequest(teacherId int32, ) *CreateTeacherEnrollmentByPlannedCourseIdRequest`

NewCreateTeacherEnrollmentByPlannedCourseIdRequest instantiates a new CreateTeacherEnrollmentByPlannedCourseIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTeacherEnrollmentByPlannedCourseIdRequestWithDefaults

`func NewCreateTeacherEnrollmentByPlannedCourseIdRequestWithDefaults() *CreateTeacherEnrollmentByPlannedCourseIdRequest`

NewCreateTeacherEnrollmentByPlannedCourseIdRequestWithDefaults instantiates a new CreateTeacherEnrollmentByPlannedCourseIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeacherId

`func (o *CreateTeacherEnrollmentByPlannedCourseIdRequest) GetTeacherId() int32`

GetTeacherId returns the TeacherId field if non-nil, zero value otherwise.

### GetTeacherIdOk

`func (o *CreateTeacherEnrollmentByPlannedCourseIdRequest) GetTeacherIdOk() (*int32, bool)`

GetTeacherIdOk returns a tuple with the TeacherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherId

`func (o *CreateTeacherEnrollmentByPlannedCourseIdRequest) SetTeacherId(v int32)`

SetTeacherId sets TeacherId field to given value.


### GetTeacherRoleId

`func (o *CreateTeacherEnrollmentByPlannedCourseIdRequest) GetTeacherRoleId() int32`

GetTeacherRoleId returns the TeacherRoleId field if non-nil, zero value otherwise.

### GetTeacherRoleIdOk

`func (o *CreateTeacherEnrollmentByPlannedCourseIdRequest) GetTeacherRoleIdOk() (*int32, bool)`

GetTeacherRoleIdOk returns a tuple with the TeacherRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherRoleId

`func (o *CreateTeacherEnrollmentByPlannedCourseIdRequest) SetTeacherRoleId(v int32)`

SetTeacherRoleId sets TeacherRoleId field to given value.

### HasTeacherRoleId

`func (o *CreateTeacherEnrollmentByPlannedCourseIdRequest) HasTeacherRoleId() bool`

HasTeacherRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


