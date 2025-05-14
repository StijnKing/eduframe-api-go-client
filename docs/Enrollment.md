# Enrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the enrollment. | [readonly] 
**StudentId** | **int32** | Unique identifier of the student. | 
**PlannedCourseId** | **int32** | Unique identifier of the planned course. | 
**OrderId** | **NullableInt32** | Unique identifier of the order. | 
**StartDate** | **string** | If it is an enrollment of a fixed course, it equals the start date. For a flexible course, it returns the enrollment specific start date. | 
**EndDate** | **NullableString** | If it is an enrollment of a fixed course, it equals the end date. For a flexible course, it returns the enrollment specific end date. | 
**Status** | [**EnrollmentStatus**](EnrollmentStatus.md) |  | 
**GraduationState** | [**GraduationState**](GraduationState.md) |  | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewEnrollment

`func NewEnrollment(id int32, studentId int32, plannedCourseId int32, orderId NullableInt32, startDate string, endDate NullableString, status EnrollmentStatus, graduationState GraduationState, updatedAt time.Time, createdAt time.Time, ) *Enrollment`

NewEnrollment instantiates a new Enrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentWithDefaults

`func NewEnrollmentWithDefaults() *Enrollment`

NewEnrollmentWithDefaults instantiates a new Enrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Enrollment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Enrollment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Enrollment) SetId(v int32)`

SetId sets Id field to given value.


### GetStudentId

`func (o *Enrollment) GetStudentId() int32`

GetStudentId returns the StudentId field if non-nil, zero value otherwise.

### GetStudentIdOk

`func (o *Enrollment) GetStudentIdOk() (*int32, bool)`

GetStudentIdOk returns a tuple with the StudentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentId

`func (o *Enrollment) SetStudentId(v int32)`

SetStudentId sets StudentId field to given value.


### GetPlannedCourseId

`func (o *Enrollment) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *Enrollment) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *Enrollment) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.


### GetOrderId

`func (o *Enrollment) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Enrollment) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Enrollment) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### SetOrderIdNil

`func (o *Enrollment) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *Enrollment) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetStartDate

`func (o *Enrollment) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Enrollment) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Enrollment) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *Enrollment) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Enrollment) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Enrollment) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *Enrollment) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Enrollment) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetStatus

`func (o *Enrollment) GetStatus() EnrollmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Enrollment) GetStatusOk() (*EnrollmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Enrollment) SetStatus(v EnrollmentStatus)`

SetStatus sets Status field to given value.


### GetGraduationState

`func (o *Enrollment) GetGraduationState() GraduationState`

GetGraduationState returns the GraduationState field if non-nil, zero value otherwise.

### GetGraduationStateOk

`func (o *Enrollment) GetGraduationStateOk() (*GraduationState, bool)`

GetGraduationStateOk returns a tuple with the GraduationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduationState

`func (o *Enrollment) SetGraduationState(v GraduationState)`

SetGraduationState sets GraduationState field to given value.


### GetUpdatedAt

`func (o *Enrollment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Enrollment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Enrollment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Enrollment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Enrollment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Enrollment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


