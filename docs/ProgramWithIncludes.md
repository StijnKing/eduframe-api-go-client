# ProgramWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the program. | [readonly] 
**SignupUrl** | **string** | URL to the signup page for this program. | 
**Name** | **string** | Name of the program. | 
**Cost** | **NullableString** | The price to be paid for this program. | 
**CostScheme** | [**CostScheme**](CostScheme.md) |  | 
**IsPublished** | **bool** | Boolean representing the publishable status of the program. | 
**CategoryId** | **int32** | Identifier of the category of the program. | 
**Conditions** | **NullableString** | Conditions for this program. | 
**Slug** | **string** | Human readable identifier, unique per educator. | 

## Methods

### NewProgramWithIncludes

`func NewProgramWithIncludes(id int32, signupUrl string, name string, cost NullableString, costScheme CostScheme, isPublished bool, categoryId int32, conditions NullableString, slug string, ) *ProgramWithIncludes`

NewProgramWithIncludes instantiates a new ProgramWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramWithIncludesWithDefaults

`func NewProgramWithIncludesWithDefaults() *ProgramWithIncludes`

NewProgramWithIncludesWithDefaults instantiates a new ProgramWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProgramWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProgramWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProgramWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetSignupUrl

`func (o *ProgramWithIncludes) GetSignupUrl() string`

GetSignupUrl returns the SignupUrl field if non-nil, zero value otherwise.

### GetSignupUrlOk

`func (o *ProgramWithIncludes) GetSignupUrlOk() (*string, bool)`

GetSignupUrlOk returns a tuple with the SignupUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupUrl

`func (o *ProgramWithIncludes) SetSignupUrl(v string)`

SetSignupUrl sets SignupUrl field to given value.


### GetName

`func (o *ProgramWithIncludes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProgramWithIncludes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProgramWithIncludes) SetName(v string)`

SetName sets Name field to given value.


### GetCost

`func (o *ProgramWithIncludes) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProgramWithIncludes) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProgramWithIncludes) SetCost(v string)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *ProgramWithIncludes) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *ProgramWithIncludes) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostScheme

`func (o *ProgramWithIncludes) GetCostScheme() CostScheme`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *ProgramWithIncludes) GetCostSchemeOk() (*CostScheme, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *ProgramWithIncludes) SetCostScheme(v CostScheme)`

SetCostScheme sets CostScheme field to given value.


### GetIsPublished

`func (o *ProgramWithIncludes) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *ProgramWithIncludes) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *ProgramWithIncludes) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.


### GetCategoryId

`func (o *ProgramWithIncludes) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ProgramWithIncludes) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ProgramWithIncludes) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetConditions

`func (o *ProgramWithIncludes) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ProgramWithIncludes) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ProgramWithIncludes) SetConditions(v string)`

SetConditions sets Conditions field to given value.


### SetConditionsNil

`func (o *ProgramWithIncludes) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *ProgramWithIncludes) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetSlug

`func (o *ProgramWithIncludes) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProgramWithIncludes) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProgramWithIncludes) SetSlug(v string)`

SetSlug sets Slug field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


