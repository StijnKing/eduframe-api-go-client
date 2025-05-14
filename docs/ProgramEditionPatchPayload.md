# ProgramEditionPatchPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **int32** | Unique identifier of associated program. | [optional] 
**Name** | Pointer to **string** | Name of the program edition. | [optional] 
**Cost** | Pointer to **NullableString** | The price to be paid for this edition. | [optional] 
**CostScheme** | Pointer to **string** | How should the edition be paid by default. | [optional] 
**MinParticipants** | Pointer to **NullableInt32** | A number representing the minimum number of participants. | [optional] 
**MaxParticipants** | Pointer to **NullableInt32** | A number representing the maximum number of participants. | [optional] 
**IsPublished** | Pointer to **bool** | Boolean representing the publishable status of the edition. | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the edition. | [optional] 
**CustomAssociations** | Pointer to [**[]CustomAssociationsRecordsPayload**](CustomAssociationsRecordsPayload.md) | The new set of custom records to be associated with the program.  ![Custom Objects](https://img.shields.io/badge/Feature-Custom_objects-1d8127)  | [optional] 

## Methods

### NewProgramEditionPatchPayload

`func NewProgramEditionPatchPayload() *ProgramEditionPatchPayload`

NewProgramEditionPatchPayload instantiates a new ProgramEditionPatchPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramEditionPatchPayloadWithDefaults

`func NewProgramEditionPatchPayloadWithDefaults() *ProgramEditionPatchPayload`

NewProgramEditionPatchPayloadWithDefaults instantiates a new ProgramEditionPatchPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *ProgramEditionPatchPayload) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ProgramEditionPatchPayload) GetProgramIdOk() (*int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *ProgramEditionPatchPayload) SetProgramId(v int32)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *ProgramEditionPatchPayload) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetName

`func (o *ProgramEditionPatchPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProgramEditionPatchPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProgramEditionPatchPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProgramEditionPatchPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCost

`func (o *ProgramEditionPatchPayload) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProgramEditionPatchPayload) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProgramEditionPatchPayload) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProgramEditionPatchPayload) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *ProgramEditionPatchPayload) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *ProgramEditionPatchPayload) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostScheme

`func (o *ProgramEditionPatchPayload) GetCostScheme() string`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *ProgramEditionPatchPayload) GetCostSchemeOk() (*string, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *ProgramEditionPatchPayload) SetCostScheme(v string)`

SetCostScheme sets CostScheme field to given value.

### HasCostScheme

`func (o *ProgramEditionPatchPayload) HasCostScheme() bool`

HasCostScheme returns a boolean if a field has been set.

### GetMinParticipants

`func (o *ProgramEditionPatchPayload) GetMinParticipants() int32`

GetMinParticipants returns the MinParticipants field if non-nil, zero value otherwise.

### GetMinParticipantsOk

`func (o *ProgramEditionPatchPayload) GetMinParticipantsOk() (*int32, bool)`

GetMinParticipantsOk returns a tuple with the MinParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinParticipants

`func (o *ProgramEditionPatchPayload) SetMinParticipants(v int32)`

SetMinParticipants sets MinParticipants field to given value.

### HasMinParticipants

`func (o *ProgramEditionPatchPayload) HasMinParticipants() bool`

HasMinParticipants returns a boolean if a field has been set.

### SetMinParticipantsNil

`func (o *ProgramEditionPatchPayload) SetMinParticipantsNil(b bool)`

 SetMinParticipantsNil sets the value for MinParticipants to be an explicit nil

### UnsetMinParticipants
`func (o *ProgramEditionPatchPayload) UnsetMinParticipants()`

UnsetMinParticipants ensures that no value is present for MinParticipants, not even an explicit nil
### GetMaxParticipants

`func (o *ProgramEditionPatchPayload) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *ProgramEditionPatchPayload) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *ProgramEditionPatchPayload) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *ProgramEditionPatchPayload) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### SetMaxParticipantsNil

`func (o *ProgramEditionPatchPayload) SetMaxParticipantsNil(b bool)`

 SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil

### UnsetMaxParticipants
`func (o *ProgramEditionPatchPayload) UnsetMaxParticipants()`

UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
### GetIsPublished

`func (o *ProgramEditionPatchPayload) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *ProgramEditionPatchPayload) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *ProgramEditionPatchPayload) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *ProgramEditionPatchPayload) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetCustom

`func (o *ProgramEditionPatchPayload) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ProgramEditionPatchPayload) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ProgramEditionPatchPayload) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *ProgramEditionPatchPayload) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomAssociations

`func (o *ProgramEditionPatchPayload) GetCustomAssociations() []CustomAssociationsRecordsPayload`

GetCustomAssociations returns the CustomAssociations field if non-nil, zero value otherwise.

### GetCustomAssociationsOk

`func (o *ProgramEditionPatchPayload) GetCustomAssociationsOk() (*[]CustomAssociationsRecordsPayload, bool)`

GetCustomAssociationsOk returns a tuple with the CustomAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssociations

`func (o *ProgramEditionPatchPayload) SetCustomAssociations(v []CustomAssociationsRecordsPayload)`

SetCustomAssociations sets CustomAssociations field to given value.

### HasCustomAssociations

`func (o *ProgramEditionPatchPayload) HasCustomAssociations() bool`

HasCustomAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


