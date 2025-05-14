# CreateGradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grade** | **NullableString** | The grade awarded (at least one of grade and score is required) | 
**Score** | **NullableFloat32** | The score awarded (at least one of grade and score is required) | 
**GradeableId** | **int32** | Unique model identifier of the gradeable (enrollment / ...) | 
**GradeableType** | **string** | Model type of the gradeable (enrollment / ...) | 
**Comment** | Pointer to **NullableString** | Additional comment about the grade | [optional] 

## Methods

### NewCreateGradeRequest

`func NewCreateGradeRequest(grade NullableString, score NullableFloat32, gradeableId int32, gradeableType string, ) *CreateGradeRequest`

NewCreateGradeRequest instantiates a new CreateGradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGradeRequestWithDefaults

`func NewCreateGradeRequestWithDefaults() *CreateGradeRequest`

NewCreateGradeRequestWithDefaults instantiates a new CreateGradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrade

`func (o *CreateGradeRequest) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *CreateGradeRequest) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *CreateGradeRequest) SetGrade(v string)`

SetGrade sets Grade field to given value.


### SetGradeNil

`func (o *CreateGradeRequest) SetGradeNil(b bool)`

 SetGradeNil sets the value for Grade to be an explicit nil

### UnsetGrade
`func (o *CreateGradeRequest) UnsetGrade()`

UnsetGrade ensures that no value is present for Grade, not even an explicit nil
### GetScore

`func (o *CreateGradeRequest) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreateGradeRequest) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreateGradeRequest) SetScore(v float32)`

SetScore sets Score field to given value.


### SetScoreNil

`func (o *CreateGradeRequest) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *CreateGradeRequest) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetGradeableId

`func (o *CreateGradeRequest) GetGradeableId() int32`

GetGradeableId returns the GradeableId field if non-nil, zero value otherwise.

### GetGradeableIdOk

`func (o *CreateGradeRequest) GetGradeableIdOk() (*int32, bool)`

GetGradeableIdOk returns a tuple with the GradeableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeableId

`func (o *CreateGradeRequest) SetGradeableId(v int32)`

SetGradeableId sets GradeableId field to given value.


### GetGradeableType

`func (o *CreateGradeRequest) GetGradeableType() string`

GetGradeableType returns the GradeableType field if non-nil, zero value otherwise.

### GetGradeableTypeOk

`func (o *CreateGradeRequest) GetGradeableTypeOk() (*string, bool)`

GetGradeableTypeOk returns a tuple with the GradeableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeableType

`func (o *CreateGradeRequest) SetGradeableType(v string)`

SetGradeableType sets GradeableType field to given value.


### GetComment

`func (o *CreateGradeRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateGradeRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateGradeRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateGradeRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *CreateGradeRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *CreateGradeRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


