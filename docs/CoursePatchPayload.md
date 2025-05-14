# CoursePatchPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **int32** | Identifier of the category of the course. | [optional] 
**Name** | Pointer to **string** | The name of the course. | [optional] 
**Code** | Pointer to **string** | The code of the course. | [optional] 
**Duration** | Pointer to **NullableString** | The duration of the course. | [optional] 
**Level** | Pointer to **NullableString** | A string indicating the level of the course. | [optional] 
**MetaTitle** | Pointer to **NullableString** | Meta title of the course for SEO purposes. | [optional] 
**MetaDescription** | Pointer to **NullableString** | Meta description of the course for SEO purposes. | [optional] 
**Result** | Pointer to **NullableString** | The result of the course | [optional] 
**Cost** | Pointer to **NullableString** | The price to be paid for this course. | [optional] 
**CostScheme** | Pointer to **string** | How should the course be paid by default. | [optional] 
**IsPublished** | Pointer to **bool** | Boolean representing the publishable status of the course. | [optional] 
**Conditions** | Pointer to **string** | The conditions of the course. | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the program. | [optional] 
**CustomAssociations** | Pointer to [**[]CustomAssociationsRecordsPayload**](CustomAssociationsRecordsPayload.md) | The new set of custom records linked to a course.  ![Custom Objects](https://img.shields.io/badge/Feature-Custom_objects-1d8127)  | [optional] 
**CourseTabContentsAttributes** | Pointer to [**[]CoursePatchPayloadCourseTabContentsAttributesInner**](CoursePatchPayloadCourseTabContentsAttributesInner.md) |  | [optional] 

## Methods

### NewCoursePatchPayload

`func NewCoursePatchPayload() *CoursePatchPayload`

NewCoursePatchPayload instantiates a new CoursePatchPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoursePatchPayloadWithDefaults

`func NewCoursePatchPayloadWithDefaults() *CoursePatchPayload`

NewCoursePatchPayloadWithDefaults instantiates a new CoursePatchPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *CoursePatchPayload) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CoursePatchPayload) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CoursePatchPayload) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CoursePatchPayload) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetName

`func (o *CoursePatchPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoursePatchPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoursePatchPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoursePatchPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *CoursePatchPayload) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CoursePatchPayload) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CoursePatchPayload) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CoursePatchPayload) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDuration

`func (o *CoursePatchPayload) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CoursePatchPayload) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CoursePatchPayload) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CoursePatchPayload) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *CoursePatchPayload) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *CoursePatchPayload) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetLevel

`func (o *CoursePatchPayload) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CoursePatchPayload) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CoursePatchPayload) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CoursePatchPayload) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *CoursePatchPayload) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *CoursePatchPayload) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetMetaTitle

`func (o *CoursePatchPayload) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *CoursePatchPayload) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *CoursePatchPayload) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *CoursePatchPayload) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### SetMetaTitleNil

`func (o *CoursePatchPayload) SetMetaTitleNil(b bool)`

 SetMetaTitleNil sets the value for MetaTitle to be an explicit nil

### UnsetMetaTitle
`func (o *CoursePatchPayload) UnsetMetaTitle()`

UnsetMetaTitle ensures that no value is present for MetaTitle, not even an explicit nil
### GetMetaDescription

`func (o *CoursePatchPayload) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *CoursePatchPayload) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *CoursePatchPayload) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *CoursePatchPayload) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### SetMetaDescriptionNil

`func (o *CoursePatchPayload) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *CoursePatchPayload) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
### GetResult

`func (o *CoursePatchPayload) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CoursePatchPayload) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CoursePatchPayload) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CoursePatchPayload) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *CoursePatchPayload) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *CoursePatchPayload) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetCost

`func (o *CoursePatchPayload) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CoursePatchPayload) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CoursePatchPayload) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CoursePatchPayload) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *CoursePatchPayload) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *CoursePatchPayload) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostScheme

`func (o *CoursePatchPayload) GetCostScheme() string`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *CoursePatchPayload) GetCostSchemeOk() (*string, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *CoursePatchPayload) SetCostScheme(v string)`

SetCostScheme sets CostScheme field to given value.

### HasCostScheme

`func (o *CoursePatchPayload) HasCostScheme() bool`

HasCostScheme returns a boolean if a field has been set.

### GetIsPublished

`func (o *CoursePatchPayload) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *CoursePatchPayload) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *CoursePatchPayload) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *CoursePatchPayload) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetConditions

`func (o *CoursePatchPayload) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CoursePatchPayload) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CoursePatchPayload) SetConditions(v string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CoursePatchPayload) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCustom

`func (o *CoursePatchPayload) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CoursePatchPayload) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CoursePatchPayload) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *CoursePatchPayload) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomAssociations

`func (o *CoursePatchPayload) GetCustomAssociations() []CustomAssociationsRecordsPayload`

GetCustomAssociations returns the CustomAssociations field if non-nil, zero value otherwise.

### GetCustomAssociationsOk

`func (o *CoursePatchPayload) GetCustomAssociationsOk() (*[]CustomAssociationsRecordsPayload, bool)`

GetCustomAssociationsOk returns a tuple with the CustomAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssociations

`func (o *CoursePatchPayload) SetCustomAssociations(v []CustomAssociationsRecordsPayload)`

SetCustomAssociations sets CustomAssociations field to given value.

### HasCustomAssociations

`func (o *CoursePatchPayload) HasCustomAssociations() bool`

HasCustomAssociations returns a boolean if a field has been set.

### GetCourseTabContentsAttributes

`func (o *CoursePatchPayload) GetCourseTabContentsAttributes() []CoursePatchPayloadCourseTabContentsAttributesInner`

GetCourseTabContentsAttributes returns the CourseTabContentsAttributes field if non-nil, zero value otherwise.

### GetCourseTabContentsAttributesOk

`func (o *CoursePatchPayload) GetCourseTabContentsAttributesOk() (*[]CoursePatchPayloadCourseTabContentsAttributesInner, bool)`

GetCourseTabContentsAttributesOk returns a tuple with the CourseTabContentsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseTabContentsAttributes

`func (o *CoursePatchPayload) SetCourseTabContentsAttributes(v []CoursePatchPayloadCourseTabContentsAttributesInner)`

SetCourseTabContentsAttributes sets CourseTabContentsAttributes field to given value.

### HasCourseTabContentsAttributes

`func (o *CoursePatchPayload) HasCourseTabContentsAttributes() bool`

HasCourseTabContentsAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


