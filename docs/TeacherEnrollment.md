# TeacherEnrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the teacher enrollment. | [readonly] 
**PlannedCourseId** | **int32** | Unique identifier of the planned course. | 
**TeacherId** | **int32** | Unique identifier of the teacher. | 
**TeacherRoleId** | Pointer to **int32** | Unique identifier of the teacher role.  ![Teacher Roles](https://img.shields.io/badge/Feature-Teacher_Roles-blue)  | [optional] 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewTeacherEnrollment

`func NewTeacherEnrollment(id int32, plannedCourseId int32, teacherId int32, updatedAt time.Time, createdAt time.Time, ) *TeacherEnrollment`

NewTeacherEnrollment instantiates a new TeacherEnrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeacherEnrollmentWithDefaults

`func NewTeacherEnrollmentWithDefaults() *TeacherEnrollment`

NewTeacherEnrollmentWithDefaults instantiates a new TeacherEnrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeacherEnrollment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeacherEnrollment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeacherEnrollment) SetId(v int32)`

SetId sets Id field to given value.


### GetPlannedCourseId

`func (o *TeacherEnrollment) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *TeacherEnrollment) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *TeacherEnrollment) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.


### GetTeacherId

`func (o *TeacherEnrollment) GetTeacherId() int32`

GetTeacherId returns the TeacherId field if non-nil, zero value otherwise.

### GetTeacherIdOk

`func (o *TeacherEnrollment) GetTeacherIdOk() (*int32, bool)`

GetTeacherIdOk returns a tuple with the TeacherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherId

`func (o *TeacherEnrollment) SetTeacherId(v int32)`

SetTeacherId sets TeacherId field to given value.


### GetTeacherRoleId

`func (o *TeacherEnrollment) GetTeacherRoleId() int32`

GetTeacherRoleId returns the TeacherRoleId field if non-nil, zero value otherwise.

### GetTeacherRoleIdOk

`func (o *TeacherEnrollment) GetTeacherRoleIdOk() (*int32, bool)`

GetTeacherRoleIdOk returns a tuple with the TeacherRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherRoleId

`func (o *TeacherEnrollment) SetTeacherRoleId(v int32)`

SetTeacherRoleId sets TeacherRoleId field to given value.

### HasTeacherRoleId

`func (o *TeacherEnrollment) HasTeacherRoleId() bool`

HasTeacherRoleId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TeacherEnrollment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeacherEnrollment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeacherEnrollment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *TeacherEnrollment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeacherEnrollment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeacherEnrollment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


