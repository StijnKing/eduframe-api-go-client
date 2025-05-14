# Grade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the grade | [readonly] 
**GraderId** | **int32** | Unique identifier of the user that awarded this grade | [readonly] 
**Result** | **string** | The actual grade which is either the grade or the score | 
**Grade** | **NullableString** | The grade awarded (at least one of grade and score is required) | 
**Score** | **NullableString** | The score awarded (at least one of grade and score is required) | 
**GradeableId** | **int32** | Unique model identifier of the gradeable (enrollment / ...) | 
**GradeableType** | **string** | Model type of the gradeable (enrollment / ...) | 
**Comment** | **NullableString** | Additional comment about the grade | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewGrade

`func NewGrade(id int32, graderId int32, result string, grade NullableString, score NullableString, gradeableId int32, gradeableType string, comment NullableString, updatedAt time.Time, createdAt time.Time, ) *Grade`

NewGrade instantiates a new Grade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGradeWithDefaults

`func NewGradeWithDefaults() *Grade`

NewGradeWithDefaults instantiates a new Grade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Grade) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Grade) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Grade) SetId(v int32)`

SetId sets Id field to given value.


### GetGraderId

`func (o *Grade) GetGraderId() int32`

GetGraderId returns the GraderId field if non-nil, zero value otherwise.

### GetGraderIdOk

`func (o *Grade) GetGraderIdOk() (*int32, bool)`

GetGraderIdOk returns a tuple with the GraderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraderId

`func (o *Grade) SetGraderId(v int32)`

SetGraderId sets GraderId field to given value.


### GetResult

`func (o *Grade) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Grade) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Grade) SetResult(v string)`

SetResult sets Result field to given value.


### GetGrade

`func (o *Grade) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *Grade) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *Grade) SetGrade(v string)`

SetGrade sets Grade field to given value.


### SetGradeNil

`func (o *Grade) SetGradeNil(b bool)`

 SetGradeNil sets the value for Grade to be an explicit nil

### UnsetGrade
`func (o *Grade) UnsetGrade()`

UnsetGrade ensures that no value is present for Grade, not even an explicit nil
### GetScore

`func (o *Grade) GetScore() string`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Grade) GetScoreOk() (*string, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Grade) SetScore(v string)`

SetScore sets Score field to given value.


### SetScoreNil

`func (o *Grade) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *Grade) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetGradeableId

`func (o *Grade) GetGradeableId() int32`

GetGradeableId returns the GradeableId field if non-nil, zero value otherwise.

### GetGradeableIdOk

`func (o *Grade) GetGradeableIdOk() (*int32, bool)`

GetGradeableIdOk returns a tuple with the GradeableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeableId

`func (o *Grade) SetGradeableId(v int32)`

SetGradeableId sets GradeableId field to given value.


### GetGradeableType

`func (o *Grade) GetGradeableType() string`

GetGradeableType returns the GradeableType field if non-nil, zero value otherwise.

### GetGradeableTypeOk

`func (o *Grade) GetGradeableTypeOk() (*string, bool)`

GetGradeableTypeOk returns a tuple with the GradeableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeableType

`func (o *Grade) SetGradeableType(v string)`

SetGradeableType sets GradeableType field to given value.


### GetComment

`func (o *Grade) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Grade) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Grade) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *Grade) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Grade) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetUpdatedAt

`func (o *Grade) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Grade) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Grade) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Grade) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Grade) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Grade) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


