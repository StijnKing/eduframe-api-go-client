# ProgramProgramPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the program. | 
**CategoryId** | **int32** | Identifier of the category of the program. | 
**CostScheme** | **string** | How should the program be paid by default. | 
**Cost** | Pointer to **NullableInt32** | The price to be paid for this program. Required if cost_scheme is student (default value) or order. | [optional] 
**IsPublished** | Pointer to **bool** | Boolean representing the publishable status of the program. | [optional] 
**LabelIds** | Pointer to **[]int32** | IDs of the labels | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the program. | [optional] 
**CustomAssociations** | Pointer to [**[]CustomAssociationsRecordsPayload**](CustomAssociationsRecordsPayload.md) | Custom associations are a way to link custom records to a program.  ![Custom Objects](https://img.shields.io/badge/Feature-Custom_objects-1d8127)  | [optional] 
**CourseTabContentsAttributes** | Pointer to [**[]CoursePayloadCourseTabContentsAttributesInner**](CoursePayloadCourseTabContentsAttributesInner.md) |  | [optional] 

## Methods

### NewProgramProgramPayload

`func NewProgramProgramPayload(name string, categoryId int32, costScheme string, ) *ProgramProgramPayload`

NewProgramProgramPayload instantiates a new ProgramProgramPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramProgramPayloadWithDefaults

`func NewProgramProgramPayloadWithDefaults() *ProgramProgramPayload`

NewProgramProgramPayloadWithDefaults instantiates a new ProgramProgramPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProgramProgramPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProgramProgramPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProgramProgramPayload) SetName(v string)`

SetName sets Name field to given value.


### GetCategoryId

`func (o *ProgramProgramPayload) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ProgramProgramPayload) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ProgramProgramPayload) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetCostScheme

`func (o *ProgramProgramPayload) GetCostScheme() string`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *ProgramProgramPayload) GetCostSchemeOk() (*string, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *ProgramProgramPayload) SetCostScheme(v string)`

SetCostScheme sets CostScheme field to given value.


### GetCost

`func (o *ProgramProgramPayload) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProgramProgramPayload) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProgramProgramPayload) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProgramProgramPayload) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *ProgramProgramPayload) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *ProgramProgramPayload) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetIsPublished

`func (o *ProgramProgramPayload) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *ProgramProgramPayload) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *ProgramProgramPayload) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *ProgramProgramPayload) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetLabelIds

`func (o *ProgramProgramPayload) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *ProgramProgramPayload) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *ProgramProgramPayload) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *ProgramProgramPayload) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetCustom

`func (o *ProgramProgramPayload) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ProgramProgramPayload) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ProgramProgramPayload) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *ProgramProgramPayload) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomAssociations

`func (o *ProgramProgramPayload) GetCustomAssociations() []CustomAssociationsRecordsPayload`

GetCustomAssociations returns the CustomAssociations field if non-nil, zero value otherwise.

### GetCustomAssociationsOk

`func (o *ProgramProgramPayload) GetCustomAssociationsOk() (*[]CustomAssociationsRecordsPayload, bool)`

GetCustomAssociationsOk returns a tuple with the CustomAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssociations

`func (o *ProgramProgramPayload) SetCustomAssociations(v []CustomAssociationsRecordsPayload)`

SetCustomAssociations sets CustomAssociations field to given value.

### HasCustomAssociations

`func (o *ProgramProgramPayload) HasCustomAssociations() bool`

HasCustomAssociations returns a boolean if a field has been set.

### GetCourseTabContentsAttributes

`func (o *ProgramProgramPayload) GetCourseTabContentsAttributes() []CoursePayloadCourseTabContentsAttributesInner`

GetCourseTabContentsAttributes returns the CourseTabContentsAttributes field if non-nil, zero value otherwise.

### GetCourseTabContentsAttributesOk

`func (o *ProgramProgramPayload) GetCourseTabContentsAttributesOk() (*[]CoursePayloadCourseTabContentsAttributesInner, bool)`

GetCourseTabContentsAttributesOk returns a tuple with the CourseTabContentsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseTabContentsAttributes

`func (o *ProgramProgramPayload) SetCourseTabContentsAttributes(v []CoursePayloadCourseTabContentsAttributesInner)`

SetCourseTabContentsAttributes sets CourseTabContentsAttributes field to given value.

### HasCourseTabContentsAttributes

`func (o *ProgramProgramPayload) HasCourseTabContentsAttributes() bool`

HasCourseTabContentsAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


