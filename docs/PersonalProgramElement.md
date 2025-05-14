# PersonalProgramElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the personal program element. | [readonly] 
**EnrollmentId** | **int32** | The identifier of the associated enrollment. | 
**CourseId** | **int32** | The identifier of the associated course. | 
**CourseEnrollmentId** | **NullableInt32** | The identifier of the associated course enrollment. | 
**ElementId** | **NullableInt32** | The identifier of the associated element. | 

## Methods

### NewPersonalProgramElement

`func NewPersonalProgramElement(id int32, enrollmentId int32, courseId int32, courseEnrollmentId NullableInt32, elementId NullableInt32, ) *PersonalProgramElement`

NewPersonalProgramElement instantiates a new PersonalProgramElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalProgramElementWithDefaults

`func NewPersonalProgramElementWithDefaults() *PersonalProgramElement`

NewPersonalProgramElementWithDefaults instantiates a new PersonalProgramElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonalProgramElement) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonalProgramElement) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonalProgramElement) SetId(v int32)`

SetId sets Id field to given value.


### GetEnrollmentId

`func (o *PersonalProgramElement) GetEnrollmentId() int32`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *PersonalProgramElement) GetEnrollmentIdOk() (*int32, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *PersonalProgramElement) SetEnrollmentId(v int32)`

SetEnrollmentId sets EnrollmentId field to given value.


### GetCourseId

`func (o *PersonalProgramElement) GetCourseId() int32`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *PersonalProgramElement) GetCourseIdOk() (*int32, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *PersonalProgramElement) SetCourseId(v int32)`

SetCourseId sets CourseId field to given value.


### GetCourseEnrollmentId

`func (o *PersonalProgramElement) GetCourseEnrollmentId() int32`

GetCourseEnrollmentId returns the CourseEnrollmentId field if non-nil, zero value otherwise.

### GetCourseEnrollmentIdOk

`func (o *PersonalProgramElement) GetCourseEnrollmentIdOk() (*int32, bool)`

GetCourseEnrollmentIdOk returns a tuple with the CourseEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseEnrollmentId

`func (o *PersonalProgramElement) SetCourseEnrollmentId(v int32)`

SetCourseEnrollmentId sets CourseEnrollmentId field to given value.


### SetCourseEnrollmentIdNil

`func (o *PersonalProgramElement) SetCourseEnrollmentIdNil(b bool)`

 SetCourseEnrollmentIdNil sets the value for CourseEnrollmentId to be an explicit nil

### UnsetCourseEnrollmentId
`func (o *PersonalProgramElement) UnsetCourseEnrollmentId()`

UnsetCourseEnrollmentId ensures that no value is present for CourseEnrollmentId, not even an explicit nil
### GetElementId

`func (o *PersonalProgramElement) GetElementId() int32`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *PersonalProgramElement) GetElementIdOk() (*int32, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *PersonalProgramElement) SetElementId(v int32)`

SetElementId sets ElementId field to given value.


### SetElementIdNil

`func (o *PersonalProgramElement) SetElementIdNil(b bool)`

 SetElementIdNil sets the value for ElementId to be an explicit nil

### UnsetElementId
`func (o *PersonalProgramElement) UnsetElementId()`

UnsetElementId ensures that no value is present for ElementId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


