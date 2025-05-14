# CoursePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int32** | Identifier of the category of the course. | 
**Name** | **string** | The name of the course. | 
**Code** | **string** | The code of the course. | 
**CostScheme** | Pointer to **string** | How should the course be paid by default. | [optional] 
**Cost** | Pointer to **NullableString** | The price to be paid for this course. Required if cost_scheme is student (default value) or order. | [optional] 
**MetaTitle** | Pointer to **NullableString** | Meta title of the course for SEO purposes. | [optional] 
**MetaDescription** | Pointer to **NullableString** | Meta description of the course for SEO purposes. | [optional] 
**IsPublished** | Pointer to **bool** | Boolean representing the publishable status of the course. | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the course. | [optional] 
**CustomAssociations** | Pointer to [**[]CustomAssociationsRecordsPayload**](CustomAssociationsRecordsPayload.md) | Custom associations are a way to link custom records to a course.  ![Custom Objects](https://img.shields.io/badge/Feature-Custom_objects-1d8127)  | [optional] 
**CourseTabContentsAttributes** | Pointer to [**[]CoursePayloadCourseTabContentsAttributesInner**](CoursePayloadCourseTabContentsAttributesInner.md) |  | [optional] 

## Methods

### NewCoursePayload

`func NewCoursePayload(categoryId int32, name string, code string, ) *CoursePayload`

NewCoursePayload instantiates a new CoursePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoursePayloadWithDefaults

`func NewCoursePayloadWithDefaults() *CoursePayload`

NewCoursePayloadWithDefaults instantiates a new CoursePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *CoursePayload) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CoursePayload) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CoursePayload) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetName

`func (o *CoursePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoursePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoursePayload) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *CoursePayload) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CoursePayload) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CoursePayload) SetCode(v string)`

SetCode sets Code field to given value.


### GetCostScheme

`func (o *CoursePayload) GetCostScheme() string`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *CoursePayload) GetCostSchemeOk() (*string, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *CoursePayload) SetCostScheme(v string)`

SetCostScheme sets CostScheme field to given value.

### HasCostScheme

`func (o *CoursePayload) HasCostScheme() bool`

HasCostScheme returns a boolean if a field has been set.

### GetCost

`func (o *CoursePayload) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CoursePayload) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CoursePayload) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CoursePayload) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *CoursePayload) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *CoursePayload) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetMetaTitle

`func (o *CoursePayload) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *CoursePayload) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *CoursePayload) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *CoursePayload) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### SetMetaTitleNil

`func (o *CoursePayload) SetMetaTitleNil(b bool)`

 SetMetaTitleNil sets the value for MetaTitle to be an explicit nil

### UnsetMetaTitle
`func (o *CoursePayload) UnsetMetaTitle()`

UnsetMetaTitle ensures that no value is present for MetaTitle, not even an explicit nil
### GetMetaDescription

`func (o *CoursePayload) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *CoursePayload) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *CoursePayload) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *CoursePayload) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### SetMetaDescriptionNil

`func (o *CoursePayload) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *CoursePayload) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
### GetIsPublished

`func (o *CoursePayload) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *CoursePayload) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *CoursePayload) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *CoursePayload) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetCustom

`func (o *CoursePayload) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CoursePayload) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CoursePayload) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *CoursePayload) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomAssociations

`func (o *CoursePayload) GetCustomAssociations() []CustomAssociationsRecordsPayload`

GetCustomAssociations returns the CustomAssociations field if non-nil, zero value otherwise.

### GetCustomAssociationsOk

`func (o *CoursePayload) GetCustomAssociationsOk() (*[]CustomAssociationsRecordsPayload, bool)`

GetCustomAssociationsOk returns a tuple with the CustomAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssociations

`func (o *CoursePayload) SetCustomAssociations(v []CustomAssociationsRecordsPayload)`

SetCustomAssociations sets CustomAssociations field to given value.

### HasCustomAssociations

`func (o *CoursePayload) HasCustomAssociations() bool`

HasCustomAssociations returns a boolean if a field has been set.

### GetCourseTabContentsAttributes

`func (o *CoursePayload) GetCourseTabContentsAttributes() []CoursePayloadCourseTabContentsAttributesInner`

GetCourseTabContentsAttributes returns the CourseTabContentsAttributes field if non-nil, zero value otherwise.

### GetCourseTabContentsAttributesOk

`func (o *CoursePayload) GetCourseTabContentsAttributesOk() (*[]CoursePayloadCourseTabContentsAttributesInner, bool)`

GetCourseTabContentsAttributesOk returns a tuple with the CourseTabContentsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseTabContentsAttributes

`func (o *CoursePayload) SetCourseTabContentsAttributes(v []CoursePayloadCourseTabContentsAttributesInner)`

SetCourseTabContentsAttributes sets CourseTabContentsAttributes field to given value.

### HasCourseTabContentsAttributes

`func (o *CoursePayload) HasCourseTabContentsAttributes() bool`

HasCourseTabContentsAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


