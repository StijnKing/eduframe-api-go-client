# Edition

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

### NewEdition

`func NewEdition(id int32, name string, programId int32, cost NullableString, costScheme CostScheme, minParticipants NullableInt32, maxParticipants NullableInt32, currentParticipants int32, isPublished bool, startDate NullableString, endDate NullableString, ) *Edition`

NewEdition instantiates a new Edition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditionWithDefaults

`func NewEditionWithDefaults() *Edition`

NewEditionWithDefaults instantiates a new Edition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Edition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Edition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Edition) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Edition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Edition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Edition) SetName(v string)`

SetName sets Name field to given value.


### GetProgramId

`func (o *Edition) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *Edition) GetProgramIdOk() (*int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *Edition) SetProgramId(v int32)`

SetProgramId sets ProgramId field to given value.


### GetCost

`func (o *Edition) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Edition) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Edition) SetCost(v string)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *Edition) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *Edition) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostScheme

`func (o *Edition) GetCostScheme() CostScheme`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *Edition) GetCostSchemeOk() (*CostScheme, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *Edition) SetCostScheme(v CostScheme)`

SetCostScheme sets CostScheme field to given value.


### GetMinParticipants

`func (o *Edition) GetMinParticipants() int32`

GetMinParticipants returns the MinParticipants field if non-nil, zero value otherwise.

### GetMinParticipantsOk

`func (o *Edition) GetMinParticipantsOk() (*int32, bool)`

GetMinParticipantsOk returns a tuple with the MinParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinParticipants

`func (o *Edition) SetMinParticipants(v int32)`

SetMinParticipants sets MinParticipants field to given value.


### SetMinParticipantsNil

`func (o *Edition) SetMinParticipantsNil(b bool)`

 SetMinParticipantsNil sets the value for MinParticipants to be an explicit nil

### UnsetMinParticipants
`func (o *Edition) UnsetMinParticipants()`

UnsetMinParticipants ensures that no value is present for MinParticipants, not even an explicit nil
### GetMaxParticipants

`func (o *Edition) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *Edition) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *Edition) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.


### SetMaxParticipantsNil

`func (o *Edition) SetMaxParticipantsNil(b bool)`

 SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil

### UnsetMaxParticipants
`func (o *Edition) UnsetMaxParticipants()`

UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
### GetCurrentParticipants

`func (o *Edition) GetCurrentParticipants() int32`

GetCurrentParticipants returns the CurrentParticipants field if non-nil, zero value otherwise.

### GetCurrentParticipantsOk

`func (o *Edition) GetCurrentParticipantsOk() (*int32, bool)`

GetCurrentParticipantsOk returns a tuple with the CurrentParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentParticipants

`func (o *Edition) SetCurrentParticipants(v int32)`

SetCurrentParticipants sets CurrentParticipants field to given value.


### GetIsPublished

`func (o *Edition) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *Edition) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *Edition) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.


### GetStartDate

`func (o *Edition) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Edition) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Edition) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### SetStartDateNil

`func (o *Edition) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Edition) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *Edition) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Edition) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Edition) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *Edition) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Edition) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


