# ProgramEditionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | **int32** | Unique identifier of associated program. | 
**Name** | **string** | Name of the program edition. | 
**Cost** | Pointer to **NullableFloat32** | The price to be paid for this edition. Required if cost_scheme is student (default value) or order. | [optional] 
**CostScheme** | Pointer to **string** | How should the edition be paid by default. | [optional] 
**MinParticipants** | Pointer to **NullableInt32** | A number representing the minimum number of participants. | [optional] 
**MaxParticipants** | Pointer to **NullableInt32** | A number representing the maximum number of participants. | [optional] 
**IsPublished** | Pointer to **bool** | Boolean representing the publishable status of the edition. | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the edition. | [optional] 
**CustomAssociations** | Pointer to [**[]CustomAssociationsRecordsPayload**](CustomAssociationsRecordsPayload.md) | Custom associations are a way to link custom records to a program edition.  ![Custom Objects](https://img.shields.io/badge/Feature-Custom_objects-1d8127)  | [optional] 

## Methods

### NewProgramEditionPayload

`func NewProgramEditionPayload(programId int32, name string, ) *ProgramEditionPayload`

NewProgramEditionPayload instantiates a new ProgramEditionPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramEditionPayloadWithDefaults

`func NewProgramEditionPayloadWithDefaults() *ProgramEditionPayload`

NewProgramEditionPayloadWithDefaults instantiates a new ProgramEditionPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *ProgramEditionPayload) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ProgramEditionPayload) GetProgramIdOk() (*int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *ProgramEditionPayload) SetProgramId(v int32)`

SetProgramId sets ProgramId field to given value.


### GetName

`func (o *ProgramEditionPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProgramEditionPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProgramEditionPayload) SetName(v string)`

SetName sets Name field to given value.


### GetCost

`func (o *ProgramEditionPayload) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProgramEditionPayload) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProgramEditionPayload) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProgramEditionPayload) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *ProgramEditionPayload) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *ProgramEditionPayload) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostScheme

`func (o *ProgramEditionPayload) GetCostScheme() string`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *ProgramEditionPayload) GetCostSchemeOk() (*string, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *ProgramEditionPayload) SetCostScheme(v string)`

SetCostScheme sets CostScheme field to given value.

### HasCostScheme

`func (o *ProgramEditionPayload) HasCostScheme() bool`

HasCostScheme returns a boolean if a field has been set.

### GetMinParticipants

`func (o *ProgramEditionPayload) GetMinParticipants() int32`

GetMinParticipants returns the MinParticipants field if non-nil, zero value otherwise.

### GetMinParticipantsOk

`func (o *ProgramEditionPayload) GetMinParticipantsOk() (*int32, bool)`

GetMinParticipantsOk returns a tuple with the MinParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinParticipants

`func (o *ProgramEditionPayload) SetMinParticipants(v int32)`

SetMinParticipants sets MinParticipants field to given value.

### HasMinParticipants

`func (o *ProgramEditionPayload) HasMinParticipants() bool`

HasMinParticipants returns a boolean if a field has been set.

### SetMinParticipantsNil

`func (o *ProgramEditionPayload) SetMinParticipantsNil(b bool)`

 SetMinParticipantsNil sets the value for MinParticipants to be an explicit nil

### UnsetMinParticipants
`func (o *ProgramEditionPayload) UnsetMinParticipants()`

UnsetMinParticipants ensures that no value is present for MinParticipants, not even an explicit nil
### GetMaxParticipants

`func (o *ProgramEditionPayload) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *ProgramEditionPayload) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *ProgramEditionPayload) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *ProgramEditionPayload) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### SetMaxParticipantsNil

`func (o *ProgramEditionPayload) SetMaxParticipantsNil(b bool)`

 SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil

### UnsetMaxParticipants
`func (o *ProgramEditionPayload) UnsetMaxParticipants()`

UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
### GetIsPublished

`func (o *ProgramEditionPayload) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *ProgramEditionPayload) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *ProgramEditionPayload) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *ProgramEditionPayload) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetCustom

`func (o *ProgramEditionPayload) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ProgramEditionPayload) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ProgramEditionPayload) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *ProgramEditionPayload) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomAssociations

`func (o *ProgramEditionPayload) GetCustomAssociations() []CustomAssociationsRecordsPayload`

GetCustomAssociations returns the CustomAssociations field if non-nil, zero value otherwise.

### GetCustomAssociationsOk

`func (o *ProgramEditionPayload) GetCustomAssociationsOk() (*[]CustomAssociationsRecordsPayload, bool)`

GetCustomAssociationsOk returns a tuple with the CustomAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssociations

`func (o *ProgramEditionPayload) SetCustomAssociations(v []CustomAssociationsRecordsPayload)`

SetCustomAssociations sets CustomAssociations field to given value.

### HasCustomAssociations

`func (o *ProgramEditionPayload) HasCustomAssociations() bool`

HasCustomAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


