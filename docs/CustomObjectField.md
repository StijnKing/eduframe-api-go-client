# CustomObjectField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the question. | [readonly] 
**Slug** | **string** | The slug of the custom field | [readonly] 
**Title** | **string** | The title and label of the question | 
**FieldType** | [**SignupQuestionFieldType**](SignupQuestionFieldType.md) |  | 

## Methods

### NewCustomObjectField

`func NewCustomObjectField(id int32, slug string, title string, fieldType SignupQuestionFieldType, ) *CustomObjectField`

NewCustomObjectField instantiates a new CustomObjectField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomObjectFieldWithDefaults

`func NewCustomObjectFieldWithDefaults() *CustomObjectField`

NewCustomObjectFieldWithDefaults instantiates a new CustomObjectField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomObjectField) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomObjectField) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomObjectField) SetId(v int32)`

SetId sets Id field to given value.


### GetSlug

`func (o *CustomObjectField) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CustomObjectField) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CustomObjectField) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *CustomObjectField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomObjectField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomObjectField) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetFieldType

`func (o *CustomObjectField) GetFieldType() SignupQuestionFieldType`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *CustomObjectField) GetFieldTypeOk() (*SignupQuestionFieldType, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *CustomObjectField) SetFieldType(v SignupQuestionFieldType)`

SetFieldType sets FieldType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


