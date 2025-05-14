# ProgramEnrollmentWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the enrollment | [readonly] 
**StudentId** | **int32** | The identifier of the associated student | 
**EditionId** | **int32** | The identifier of the associated edition | 
**PersonalProgramId** | **int32** | The identifier of the associated personal program | 
**LabelIds** | **[]int32** | An array of integers representing unique identifier values associated with labels.  | 
**OrderId** | **NullableInt32** | The identifier of the associated order | 
**State** | **string** | The state of the enrollment | 
**GraduationState** | [**GraduationState**](GraduationState.md) |  | 
**GraduationDate** | **NullableString** | The graduation date of the enrollment | 
**Grade** | Pointer to [**NullableGrade**](Grade.md) |  | [optional] 

## Methods

### NewProgramEnrollmentWithIncludes

`func NewProgramEnrollmentWithIncludes(id int32, studentId int32, editionId int32, personalProgramId int32, labelIds []int32, orderId NullableInt32, state string, graduationState GraduationState, graduationDate NullableString, ) *ProgramEnrollmentWithIncludes`

NewProgramEnrollmentWithIncludes instantiates a new ProgramEnrollmentWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramEnrollmentWithIncludesWithDefaults

`func NewProgramEnrollmentWithIncludesWithDefaults() *ProgramEnrollmentWithIncludes`

NewProgramEnrollmentWithIncludesWithDefaults instantiates a new ProgramEnrollmentWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProgramEnrollmentWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProgramEnrollmentWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProgramEnrollmentWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetStudentId

`func (o *ProgramEnrollmentWithIncludes) GetStudentId() int32`

GetStudentId returns the StudentId field if non-nil, zero value otherwise.

### GetStudentIdOk

`func (o *ProgramEnrollmentWithIncludes) GetStudentIdOk() (*int32, bool)`

GetStudentIdOk returns a tuple with the StudentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentId

`func (o *ProgramEnrollmentWithIncludes) SetStudentId(v int32)`

SetStudentId sets StudentId field to given value.


### GetEditionId

`func (o *ProgramEnrollmentWithIncludes) GetEditionId() int32`

GetEditionId returns the EditionId field if non-nil, zero value otherwise.

### GetEditionIdOk

`func (o *ProgramEnrollmentWithIncludes) GetEditionIdOk() (*int32, bool)`

GetEditionIdOk returns a tuple with the EditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionId

`func (o *ProgramEnrollmentWithIncludes) SetEditionId(v int32)`

SetEditionId sets EditionId field to given value.


### GetPersonalProgramId

`func (o *ProgramEnrollmentWithIncludes) GetPersonalProgramId() int32`

GetPersonalProgramId returns the PersonalProgramId field if non-nil, zero value otherwise.

### GetPersonalProgramIdOk

`func (o *ProgramEnrollmentWithIncludes) GetPersonalProgramIdOk() (*int32, bool)`

GetPersonalProgramIdOk returns a tuple with the PersonalProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalProgramId

`func (o *ProgramEnrollmentWithIncludes) SetPersonalProgramId(v int32)`

SetPersonalProgramId sets PersonalProgramId field to given value.


### GetLabelIds

`func (o *ProgramEnrollmentWithIncludes) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *ProgramEnrollmentWithIncludes) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *ProgramEnrollmentWithIncludes) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.


### GetOrderId

`func (o *ProgramEnrollmentWithIncludes) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ProgramEnrollmentWithIncludes) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ProgramEnrollmentWithIncludes) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### SetOrderIdNil

`func (o *ProgramEnrollmentWithIncludes) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *ProgramEnrollmentWithIncludes) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetState

`func (o *ProgramEnrollmentWithIncludes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProgramEnrollmentWithIncludes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProgramEnrollmentWithIncludes) SetState(v string)`

SetState sets State field to given value.


### GetGraduationState

`func (o *ProgramEnrollmentWithIncludes) GetGraduationState() GraduationState`

GetGraduationState returns the GraduationState field if non-nil, zero value otherwise.

### GetGraduationStateOk

`func (o *ProgramEnrollmentWithIncludes) GetGraduationStateOk() (*GraduationState, bool)`

GetGraduationStateOk returns a tuple with the GraduationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduationState

`func (o *ProgramEnrollmentWithIncludes) SetGraduationState(v GraduationState)`

SetGraduationState sets GraduationState field to given value.


### GetGraduationDate

`func (o *ProgramEnrollmentWithIncludes) GetGraduationDate() string`

GetGraduationDate returns the GraduationDate field if non-nil, zero value otherwise.

### GetGraduationDateOk

`func (o *ProgramEnrollmentWithIncludes) GetGraduationDateOk() (*string, bool)`

GetGraduationDateOk returns a tuple with the GraduationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduationDate

`func (o *ProgramEnrollmentWithIncludes) SetGraduationDate(v string)`

SetGraduationDate sets GraduationDate field to given value.


### SetGraduationDateNil

`func (o *ProgramEnrollmentWithIncludes) SetGraduationDateNil(b bool)`

 SetGraduationDateNil sets the value for GraduationDate to be an explicit nil

### UnsetGraduationDate
`func (o *ProgramEnrollmentWithIncludes) UnsetGraduationDate()`

UnsetGraduationDate ensures that no value is present for GraduationDate, not even an explicit nil
### GetGrade

`func (o *ProgramEnrollmentWithIncludes) GetGrade() Grade`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *ProgramEnrollmentWithIncludes) GetGradeOk() (*Grade, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *ProgramEnrollmentWithIncludes) SetGrade(v Grade)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *ProgramEnrollmentWithIncludes) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### SetGradeNil

`func (o *ProgramEnrollmentWithIncludes) SetGradeNil(b bool)`

 SetGradeNil sets the value for Grade to be an explicit nil

### UnsetGrade
`func (o *ProgramEnrollmentWithIncludes) UnsetGrade()`

UnsetGrade ensures that no value is present for Grade, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


