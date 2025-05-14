# CreateCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | A string representing the content of a comment. | 
**CommentableId** | **int32** | Identifier of the subject the comment is linked to. | 
**CommentableType** | [**CommentType**](CommentType.md) |  | 

## Methods

### NewCreateCommentRequest

`func NewCreateCommentRequest(content string, commentableId int32, commentableType CommentType, ) *CreateCommentRequest`

NewCreateCommentRequest instantiates a new CreateCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCommentRequestWithDefaults

`func NewCreateCommentRequestWithDefaults() *CreateCommentRequest`

NewCreateCommentRequestWithDefaults instantiates a new CreateCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateCommentRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateCommentRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateCommentRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetCommentableId

`func (o *CreateCommentRequest) GetCommentableId() int32`

GetCommentableId returns the CommentableId field if non-nil, zero value otherwise.

### GetCommentableIdOk

`func (o *CreateCommentRequest) GetCommentableIdOk() (*int32, bool)`

GetCommentableIdOk returns a tuple with the CommentableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentableId

`func (o *CreateCommentRequest) SetCommentableId(v int32)`

SetCommentableId sets CommentableId field to given value.


### GetCommentableType

`func (o *CreateCommentRequest) GetCommentableType() CommentType`

GetCommentableType returns the CommentableType field if non-nil, zero value otherwise.

### GetCommentableTypeOk

`func (o *CreateCommentRequest) GetCommentableTypeOk() (*CommentType, bool)`

GetCommentableTypeOk returns a tuple with the CommentableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentableType

`func (o *CreateCommentRequest) SetCommentableType(v CommentType)`

SetCommentableType sets CommentableType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


