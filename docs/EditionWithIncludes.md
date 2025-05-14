# EditionWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the program edition. | [readonly] 
**Name** | **string** | Name of the program edition. | 
**ProgramId** | **int32** | Unique identifier of associated program. | 
**Cost** | **NullableString** | The price to be paid for this edition. | 
**CostScheme** | [**CostScheme**](CostScheme.md) |  | 
**MinParticipants** | **NullableInt32** | A number representing the minimum number of participants. | 
**MaxParticipants** | **NullableInt32** | A number representing the maximum number of participants. | 
**CurrentParticipants** | **int32** | The current amount of participants. | [readonly] 
**IsPublished** | **bool** | Boolean representing the publishable status of the edition. | [default to false]
**StartDate** | **NullableString** | Nominal start date of the edition. | 
**EndDate** | **NullableString** | Nominal end date of the edition (inclusive). | 

## Methods

### NewEditionWithIncludes

`func NewEditionWithIncludes(id int32, name string, programId int32, cost NullableString, costScheme CostScheme, minParticipants NullableInt32, maxParticipants NullableInt32, currentParticipants int32, isPublished bool, startDate NullableString, endDate NullableString, ) *EditionWithIncludes`

NewEditionWithIncludes instantiates a new EditionWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditionWithIncludesWithDefaults

`func NewEditionWithIncludesWithDefaults() *EditionWithIncludes`

NewEditionWithIncludesWithDefaults instantiates a new EditionWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditionWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditionWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditionWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *EditionWithIncludes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditionWithIncludes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditionWithIncludes) SetName(v string)`

SetName sets Name field to given value.


### GetProgramId

`func (o *EditionWithIncludes) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *EditionWithIncludes) GetProgramIdOk() (*int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *EditionWithIncludes) SetProgramId(v int32)`

SetProgramId sets ProgramId field to given value.


### GetCost

`func (o *EditionWithIncludes) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *EditionWithIncludes) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *EditionWithIncludes) SetCost(v string)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *EditionWithIncludes) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *EditionWithIncludes) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostScheme

`func (o *EditionWithIncludes) GetCostScheme() CostScheme`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *EditionWithIncludes) GetCostSchemeOk() (*CostScheme, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *EditionWithIncludes) SetCostScheme(v CostScheme)`

SetCostScheme sets CostScheme field to given value.


### GetMinParticipants

`func (o *EditionWithIncludes) GetMinParticipants() int32`

GetMinParticipants returns the MinParticipants field if non-nil, zero value otherwise.

### GetMinParticipantsOk

`func (o *EditionWithIncludes) GetMinParticipantsOk() (*int32, bool)`

GetMinParticipantsOk returns a tuple with the MinParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinParticipants

`func (o *EditionWithIncludes) SetMinParticipants(v int32)`

SetMinParticipants sets MinParticipants field to given value.


### SetMinParticipantsNil

`func (o *EditionWithIncludes) SetMinParticipantsNil(b bool)`

 SetMinParticipantsNil sets the value for MinParticipants to be an explicit nil

### UnsetMinParticipants
`func (o *EditionWithIncludes) UnsetMinParticipants()`

UnsetMinParticipants ensures that no value is present for MinParticipants, not even an explicit nil
### GetMaxParticipants

`func (o *EditionWithIncludes) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *EditionWithIncludes) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *EditionWithIncludes) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.


### SetMaxParticipantsNil

`func (o *EditionWithIncludes) SetMaxParticipantsNil(b bool)`

 SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil

### UnsetMaxParticipants
`func (o *EditionWithIncludes) UnsetMaxParticipants()`

UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
### GetCurrentParticipants

`func (o *EditionWithIncludes) GetCurrentParticipants() int32`

GetCurrentParticipants returns the CurrentParticipants field if non-nil, zero value otherwise.

### GetCurrentParticipantsOk

`func (o *EditionWithIncludes) GetCurrentParticipantsOk() (*int32, bool)`

GetCurrentParticipantsOk returns a tuple with the CurrentParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentParticipants

`func (o *EditionWithIncludes) SetCurrentParticipants(v int32)`

SetCurrentParticipants sets CurrentParticipants field to given value.


### GetIsPublished

`func (o *EditionWithIncludes) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *EditionWithIncludes) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *EditionWithIncludes) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.


### GetStartDate

`func (o *EditionWithIncludes) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EditionWithIncludes) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EditionWithIncludes) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### SetStartDateNil

`func (o *EditionWithIncludes) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *EditionWithIncludes) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *EditionWithIncludes) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EditionWithIncludes) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EditionWithIncludes) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *EditionWithIncludes) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *EditionWithIncludes) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


