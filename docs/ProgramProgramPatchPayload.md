# ProgramProgramPatchPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the program. | [optional] 
**Cost** | Pointer to **NullableString** | The price to be paid for this program. | [optional] 
**CostScheme** | Pointer to **string** | How should the program be paid by default. | [optional] 
**IsPublished** | Pointer to **bool** | Boolean representing the publishable status of the program. | [optional] 
**CategoryId** | Pointer to **int32** | Identifier of the category of the program. | [optional] 
**Slug** | Pointer to **string** | Human readable identifier, unique per educator. | [optional] 
**LabelIds** | Pointer to **[]int32** | IDs of the labels | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the program. | [optional] 
**CustomAssociations** | Pointer to [**[]CustomAssociationsRecordsPayload**](CustomAssociationsRecordsPayload.md) | The new set of custom records to be associated with the program.  ![Custom Objects](https://img.shields.io/badge/Feature-Custom_objects-1d8127)  | [optional] 
**CourseTabContentsAttributes** | Pointer to [**[]ProgramProgramPatchPayloadCourseTabContentsAttributesInner**](ProgramProgramPatchPayloadCourseTabContentsAttributesInner.md) |  | [optional] 

## Methods

### NewProgramProgramPatchPayload

`func NewProgramProgramPatchPayload() *ProgramProgramPatchPayload`

NewProgramProgramPatchPayload instantiates a new ProgramProgramPatchPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramProgramPatchPayloadWithDefaults

`func NewProgramProgramPatchPayloadWithDefaults() *ProgramProgramPatchPayload`

NewProgramProgramPatchPayloadWithDefaults instantiates a new ProgramProgramPatchPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProgramProgramPatchPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProgramProgramPatchPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProgramProgramPatchPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProgramProgramPatchPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCost

`func (o *ProgramProgramPatchPayload) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProgramProgramPatchPayload) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProgramProgramPatchPayload) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProgramProgramPatchPayload) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *ProgramProgramPatchPayload) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *ProgramProgramPatchPayload) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostScheme

`func (o *ProgramProgramPatchPayload) GetCostScheme() string`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *ProgramProgramPatchPayload) GetCostSchemeOk() (*string, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *ProgramProgramPatchPayload) SetCostScheme(v string)`

SetCostScheme sets CostScheme field to given value.

### HasCostScheme

`func (o *ProgramProgramPatchPayload) HasCostScheme() bool`

HasCostScheme returns a boolean if a field has been set.

### GetIsPublished

`func (o *ProgramProgramPatchPayload) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *ProgramProgramPatchPayload) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *ProgramProgramPatchPayload) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *ProgramProgramPatchPayload) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetCategoryId

`func (o *ProgramProgramPatchPayload) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ProgramProgramPatchPayload) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ProgramProgramPatchPayload) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ProgramProgramPatchPayload) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetSlug

`func (o *ProgramProgramPatchPayload) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProgramProgramPatchPayload) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProgramProgramPatchPayload) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProgramProgramPatchPayload) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetLabelIds

`func (o *ProgramProgramPatchPayload) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *ProgramProgramPatchPayload) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *ProgramProgramPatchPayload) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *ProgramProgramPatchPayload) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetCustom

`func (o *ProgramProgramPatchPayload) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ProgramProgramPatchPayload) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ProgramProgramPatchPayload) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *ProgramProgramPatchPayload) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomAssociations

`func (o *ProgramProgramPatchPayload) GetCustomAssociations() []CustomAssociationsRecordsPayload`

GetCustomAssociations returns the CustomAssociations field if non-nil, zero value otherwise.

### GetCustomAssociationsOk

`func (o *ProgramProgramPatchPayload) GetCustomAssociationsOk() (*[]CustomAssociationsRecordsPayload, bool)`

GetCustomAssociationsOk returns a tuple with the CustomAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssociations

`func (o *ProgramProgramPatchPayload) SetCustomAssociations(v []CustomAssociationsRecordsPayload)`

SetCustomAssociations sets CustomAssociations field to given value.

### HasCustomAssociations

`func (o *ProgramProgramPatchPayload) HasCustomAssociations() bool`

HasCustomAssociations returns a boolean if a field has been set.

### GetCourseTabContentsAttributes

`func (o *ProgramProgramPatchPayload) GetCourseTabContentsAttributes() []ProgramProgramPatchPayloadCourseTabContentsAttributesInner`

GetCourseTabContentsAttributes returns the CourseTabContentsAttributes field if non-nil, zero value otherwise.

### GetCourseTabContentsAttributesOk

`func (o *ProgramProgramPatchPayload) GetCourseTabContentsAttributesOk() (*[]ProgramProgramPatchPayloadCourseTabContentsAttributesInner, bool)`

GetCourseTabContentsAttributesOk returns a tuple with the CourseTabContentsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseTabContentsAttributes

`func (o *ProgramProgramPatchPayload) SetCourseTabContentsAttributes(v []ProgramProgramPatchPayloadCourseTabContentsAttributesInner)`

SetCourseTabContentsAttributes sets CourseTabContentsAttributes field to given value.

### HasCourseTabContentsAttributes

`func (o *ProgramProgramPatchPayload) HasCourseTabContentsAttributes() bool`

HasCourseTabContentsAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


