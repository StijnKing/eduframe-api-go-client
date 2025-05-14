# UpdateCatalogProductByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **int32** | Identifier of the category of the course. | [optional] 
**IsPublished** | Pointer to **bool** | Boolean showing if the product is published or not. | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the product. | [optional] 
**CourseTabContentsAttributes** | Pointer to [**[]UpdateCatalogProductByIdRequestCourseTabContentsAttributesInner**](UpdateCatalogProductByIdRequestCourseTabContentsAttributesInner.md) |  | [optional] 

## Methods

### NewUpdateCatalogProductByIdRequest

`func NewUpdateCatalogProductByIdRequest() *UpdateCatalogProductByIdRequest`

NewUpdateCatalogProductByIdRequest instantiates a new UpdateCatalogProductByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCatalogProductByIdRequestWithDefaults

`func NewUpdateCatalogProductByIdRequestWithDefaults() *UpdateCatalogProductByIdRequest`

NewUpdateCatalogProductByIdRequestWithDefaults instantiates a new UpdateCatalogProductByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *UpdateCatalogProductByIdRequest) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *UpdateCatalogProductByIdRequest) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *UpdateCatalogProductByIdRequest) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *UpdateCatalogProductByIdRequest) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetIsPublished

`func (o *UpdateCatalogProductByIdRequest) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *UpdateCatalogProductByIdRequest) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *UpdateCatalogProductByIdRequest) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *UpdateCatalogProductByIdRequest) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetCustom

`func (o *UpdateCatalogProductByIdRequest) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UpdateCatalogProductByIdRequest) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UpdateCatalogProductByIdRequest) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *UpdateCatalogProductByIdRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCourseTabContentsAttributes

`func (o *UpdateCatalogProductByIdRequest) GetCourseTabContentsAttributes() []UpdateCatalogProductByIdRequestCourseTabContentsAttributesInner`

GetCourseTabContentsAttributes returns the CourseTabContentsAttributes field if non-nil, zero value otherwise.

### GetCourseTabContentsAttributesOk

`func (o *UpdateCatalogProductByIdRequest) GetCourseTabContentsAttributesOk() (*[]UpdateCatalogProductByIdRequestCourseTabContentsAttributesInner, bool)`

GetCourseTabContentsAttributesOk returns a tuple with the CourseTabContentsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseTabContentsAttributes

`func (o *UpdateCatalogProductByIdRequest) SetCourseTabContentsAttributes(v []UpdateCatalogProductByIdRequestCourseTabContentsAttributesInner)`

SetCourseTabContentsAttributes sets CourseTabContentsAttributes field to given value.

### HasCourseTabContentsAttributes

`func (o *UpdateCatalogProductByIdRequest) HasCourseTabContentsAttributes() bool`

HasCourseTabContentsAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


