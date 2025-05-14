# Credit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the credit. | [readonly] 
**StudentId** | **int32** | Unique identifier of the orders Student (User). | 
**Credits** | **float32** | The amount of awarded credits. | 
**Description** | **NullableString** | Description of the awarded credits. | 
**CourseId** | **NullableInt32** | Unique identifier of the Course. | 
**EnrollmentId** | **NullableInt32** | Unique identifier of the Enrollment. | 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewCredit

`func NewCredit(id int32, studentId int32, credits float32, description NullableString, courseId NullableInt32, enrollmentId NullableInt32, createdAt time.Time, ) *Credit`

NewCredit instantiates a new Credit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditWithDefaults

`func NewCreditWithDefaults() *Credit`

NewCreditWithDefaults instantiates a new Credit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Credit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Credit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Credit) SetId(v int32)`

SetId sets Id field to given value.


### GetStudentId

`func (o *Credit) GetStudentId() int32`

GetStudentId returns the StudentId field if non-nil, zero value otherwise.

### GetStudentIdOk

`func (o *Credit) GetStudentIdOk() (*int32, bool)`

GetStudentIdOk returns a tuple with the StudentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentId

`func (o *Credit) SetStudentId(v int32)`

SetStudentId sets StudentId field to given value.


### GetCredits

`func (o *Credit) GetCredits() float32`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *Credit) GetCreditsOk() (*float32, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *Credit) SetCredits(v float32)`

SetCredits sets Credits field to given value.


### GetDescription

`func (o *Credit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Credit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Credit) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Credit) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Credit) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCourseId

`func (o *Credit) GetCourseId() int32`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *Credit) GetCourseIdOk() (*int32, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *Credit) SetCourseId(v int32)`

SetCourseId sets CourseId field to given value.


### SetCourseIdNil

`func (o *Credit) SetCourseIdNil(b bool)`

 SetCourseIdNil sets the value for CourseId to be an explicit nil

### UnsetCourseId
`func (o *Credit) UnsetCourseId()`

UnsetCourseId ensures that no value is present for CourseId, not even an explicit nil
### GetEnrollmentId

`func (o *Credit) GetEnrollmentId() int32`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *Credit) GetEnrollmentIdOk() (*int32, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *Credit) SetEnrollmentId(v int32)`

SetEnrollmentId sets EnrollmentId field to given value.


### SetEnrollmentIdNil

`func (o *Credit) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *Credit) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetCreatedAt

`func (o *Credit) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Credit) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Credit) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


