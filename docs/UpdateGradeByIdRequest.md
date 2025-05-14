# UpdateGradeByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grade** | Pointer to **NullableString** | The grade awarded (at least one of grade and score is required) | [optional] 
**Score** | Pointer to **NullableFloat32** | The score awarded (at least one of grade and score is required) | [optional] 
**GradeableId** | Pointer to **int32** | Unique model identifier of the gradeable (enrollment / ...) | [optional] 
**GradeableType** | Pointer to **string** | Model type of the gradeable (enrollment / ...) | [optional] 
**Comment** | Pointer to **NullableString** | Additional comment about the grade | [optional] 
**EnrollmentId** | Pointer to **int32** | Unique identifier of the enrollment | [optional] 

## Methods

### NewUpdateGradeByIdRequest

`func NewUpdateGradeByIdRequest() *UpdateGradeByIdRequest`

NewUpdateGradeByIdRequest instantiates a new UpdateGradeByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGradeByIdRequestWithDefaults

`func NewUpdateGradeByIdRequestWithDefaults() *UpdateGradeByIdRequest`

NewUpdateGradeByIdRequestWithDefaults instantiates a new UpdateGradeByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrade

`func (o *UpdateGradeByIdRequest) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *UpdateGradeByIdRequest) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *UpdateGradeByIdRequest) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *UpdateGradeByIdRequest) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### SetGradeNil

`func (o *UpdateGradeByIdRequest) SetGradeNil(b bool)`

 SetGradeNil sets the value for Grade to be an explicit nil

### UnsetGrade
`func (o *UpdateGradeByIdRequest) UnsetGrade()`

UnsetGrade ensures that no value is present for Grade, not even an explicit nil
### GetScore

`func (o *UpdateGradeByIdRequest) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *UpdateGradeByIdRequest) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *UpdateGradeByIdRequest) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *UpdateGradeByIdRequest) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *UpdateGradeByIdRequest) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *UpdateGradeByIdRequest) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetGradeableId

`func (o *UpdateGradeByIdRequest) GetGradeableId() int32`

GetGradeableId returns the GradeableId field if non-nil, zero value otherwise.

### GetGradeableIdOk

`func (o *UpdateGradeByIdRequest) GetGradeableIdOk() (*int32, bool)`

GetGradeableIdOk returns a tuple with the GradeableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeableId

`func (o *UpdateGradeByIdRequest) SetGradeableId(v int32)`

SetGradeableId sets GradeableId field to given value.

### HasGradeableId

`func (o *UpdateGradeByIdRequest) HasGradeableId() bool`

HasGradeableId returns a boolean if a field has been set.

### GetGradeableType

`func (o *UpdateGradeByIdRequest) GetGradeableType() string`

GetGradeableType returns the GradeableType field if non-nil, zero value otherwise.

### GetGradeableTypeOk

`func (o *UpdateGradeByIdRequest) GetGradeableTypeOk() (*string, bool)`

GetGradeableTypeOk returns a tuple with the GradeableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeableType

`func (o *UpdateGradeByIdRequest) SetGradeableType(v string)`

SetGradeableType sets GradeableType field to given value.

### HasGradeableType

`func (o *UpdateGradeByIdRequest) HasGradeableType() bool`

HasGradeableType returns a boolean if a field has been set.

### GetComment

`func (o *UpdateGradeByIdRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateGradeByIdRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateGradeByIdRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateGradeByIdRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *UpdateGradeByIdRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *UpdateGradeByIdRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetEnrollmentId

`func (o *UpdateGradeByIdRequest) GetEnrollmentId() int32`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *UpdateGradeByIdRequest) GetEnrollmentIdOk() (*int32, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *UpdateGradeByIdRequest) SetEnrollmentId(v int32)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *UpdateGradeByIdRequest) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


