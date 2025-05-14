# PlannedCoursePatchPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPublished** | Pointer to **bool** | Boolean if is published on the website. | [optional] 
**CourseId** | Pointer to **int32** | Unique identifier of the course. | [optional] 
**StartDate** | Pointer to **string** | Date at which the planned course starts. Only needed for fixed planned courses. | [optional] 
**EndDate** | Pointer to **NullableString** | Date at which the planned course ends. Only needed for fixed planned courses. | [optional] 
**MinParticipants** | Pointer to **NullableInt32** | A number representing the minimum number of participants that can enroll for the planned course. | [optional] 
**MaxParticipants** | Pointer to **NullableInt32** | A number representing the maximum number of participants that can enroll for the planned course. | [optional] 
**CostScheme** | Pointer to **string** | The cost schema that the payment will follow for the specified course. | [optional] 
**Cost** | Pointer to **NullableFloat32** | A positive float representing the price of the planned course. | [optional] 
**CourseVariantId** | Pointer to **NullableInt32** | Unique identifier of the course variant. | [optional] 
**CourseLocationId** | Pointer to **NullableInt32** | Unique identifier of the course location. | [optional] 
**Duration** | Pointer to **float32** | The period of time of the planned course. Only needed for flexible planned courses. | [optional] 
**TeacherIds** | Pointer to **[]string** | The ids of the teachers in the course | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the planned course. | [optional] 
**CustomAssociations** | Pointer to [**[]CustomAssociationsRecordsPayload**](CustomAssociationsRecordsPayload.md) | Custom associations are a way to link custom records to a program. ![Custom Objects](https://img.shields.io/badge/Feature-Custom_objects-1d8127)  | [optional] 

## Methods

### NewPlannedCoursePatchPayload

`func NewPlannedCoursePatchPayload() *PlannedCoursePatchPayload`

NewPlannedCoursePatchPayload instantiates a new PlannedCoursePatchPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannedCoursePatchPayloadWithDefaults

`func NewPlannedCoursePatchPayloadWithDefaults() *PlannedCoursePatchPayload`

NewPlannedCoursePatchPayloadWithDefaults instantiates a new PlannedCoursePatchPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPublished

`func (o *PlannedCoursePatchPayload) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *PlannedCoursePatchPayload) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *PlannedCoursePatchPayload) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *PlannedCoursePatchPayload) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetCourseId

`func (o *PlannedCoursePatchPayload) GetCourseId() int32`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *PlannedCoursePatchPayload) GetCourseIdOk() (*int32, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *PlannedCoursePatchPayload) SetCourseId(v int32)`

SetCourseId sets CourseId field to given value.

### HasCourseId

`func (o *PlannedCoursePatchPayload) HasCourseId() bool`

HasCourseId returns a boolean if a field has been set.

### GetStartDate

`func (o *PlannedCoursePatchPayload) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PlannedCoursePatchPayload) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PlannedCoursePatchPayload) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PlannedCoursePatchPayload) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PlannedCoursePatchPayload) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PlannedCoursePatchPayload) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PlannedCoursePatchPayload) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PlannedCoursePatchPayload) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *PlannedCoursePatchPayload) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *PlannedCoursePatchPayload) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetMinParticipants

`func (o *PlannedCoursePatchPayload) GetMinParticipants() int32`

GetMinParticipants returns the MinParticipants field if non-nil, zero value otherwise.

### GetMinParticipantsOk

`func (o *PlannedCoursePatchPayload) GetMinParticipantsOk() (*int32, bool)`

GetMinParticipantsOk returns a tuple with the MinParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinParticipants

`func (o *PlannedCoursePatchPayload) SetMinParticipants(v int32)`

SetMinParticipants sets MinParticipants field to given value.

### HasMinParticipants

`func (o *PlannedCoursePatchPayload) HasMinParticipants() bool`

HasMinParticipants returns a boolean if a field has been set.

### SetMinParticipantsNil

`func (o *PlannedCoursePatchPayload) SetMinParticipantsNil(b bool)`

 SetMinParticipantsNil sets the value for MinParticipants to be an explicit nil

### UnsetMinParticipants
`func (o *PlannedCoursePatchPayload) UnsetMinParticipants()`

UnsetMinParticipants ensures that no value is present for MinParticipants, not even an explicit nil
### GetMaxParticipants

`func (o *PlannedCoursePatchPayload) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *PlannedCoursePatchPayload) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *PlannedCoursePatchPayload) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *PlannedCoursePatchPayload) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### SetMaxParticipantsNil

`func (o *PlannedCoursePatchPayload) SetMaxParticipantsNil(b bool)`

 SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil

### UnsetMaxParticipants
`func (o *PlannedCoursePatchPayload) UnsetMaxParticipants()`

UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
### GetCostScheme

`func (o *PlannedCoursePatchPayload) GetCostScheme() string`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *PlannedCoursePatchPayload) GetCostSchemeOk() (*string, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *PlannedCoursePatchPayload) SetCostScheme(v string)`

SetCostScheme sets CostScheme field to given value.

### HasCostScheme

`func (o *PlannedCoursePatchPayload) HasCostScheme() bool`

HasCostScheme returns a boolean if a field has been set.

### GetCost

`func (o *PlannedCoursePatchPayload) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *PlannedCoursePatchPayload) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *PlannedCoursePatchPayload) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *PlannedCoursePatchPayload) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *PlannedCoursePatchPayload) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *PlannedCoursePatchPayload) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCourseVariantId

`func (o *PlannedCoursePatchPayload) GetCourseVariantId() int32`

GetCourseVariantId returns the CourseVariantId field if non-nil, zero value otherwise.

### GetCourseVariantIdOk

`func (o *PlannedCoursePatchPayload) GetCourseVariantIdOk() (*int32, bool)`

GetCourseVariantIdOk returns a tuple with the CourseVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseVariantId

`func (o *PlannedCoursePatchPayload) SetCourseVariantId(v int32)`

SetCourseVariantId sets CourseVariantId field to given value.

### HasCourseVariantId

`func (o *PlannedCoursePatchPayload) HasCourseVariantId() bool`

HasCourseVariantId returns a boolean if a field has been set.

### SetCourseVariantIdNil

`func (o *PlannedCoursePatchPayload) SetCourseVariantIdNil(b bool)`

 SetCourseVariantIdNil sets the value for CourseVariantId to be an explicit nil

### UnsetCourseVariantId
`func (o *PlannedCoursePatchPayload) UnsetCourseVariantId()`

UnsetCourseVariantId ensures that no value is present for CourseVariantId, not even an explicit nil
### GetCourseLocationId

`func (o *PlannedCoursePatchPayload) GetCourseLocationId() int32`

GetCourseLocationId returns the CourseLocationId field if non-nil, zero value otherwise.

### GetCourseLocationIdOk

`func (o *PlannedCoursePatchPayload) GetCourseLocationIdOk() (*int32, bool)`

GetCourseLocationIdOk returns a tuple with the CourseLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseLocationId

`func (o *PlannedCoursePatchPayload) SetCourseLocationId(v int32)`

SetCourseLocationId sets CourseLocationId field to given value.

### HasCourseLocationId

`func (o *PlannedCoursePatchPayload) HasCourseLocationId() bool`

HasCourseLocationId returns a boolean if a field has been set.

### SetCourseLocationIdNil

`func (o *PlannedCoursePatchPayload) SetCourseLocationIdNil(b bool)`

 SetCourseLocationIdNil sets the value for CourseLocationId to be an explicit nil

### UnsetCourseLocationId
`func (o *PlannedCoursePatchPayload) UnsetCourseLocationId()`

UnsetCourseLocationId ensures that no value is present for CourseLocationId, not even an explicit nil
### GetDuration

`func (o *PlannedCoursePatchPayload) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PlannedCoursePatchPayload) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PlannedCoursePatchPayload) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PlannedCoursePatchPayload) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTeacherIds

`func (o *PlannedCoursePatchPayload) GetTeacherIds() []string`

GetTeacherIds returns the TeacherIds field if non-nil, zero value otherwise.

### GetTeacherIdsOk

`func (o *PlannedCoursePatchPayload) GetTeacherIdsOk() (*[]string, bool)`

GetTeacherIdsOk returns a tuple with the TeacherIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherIds

`func (o *PlannedCoursePatchPayload) SetTeacherIds(v []string)`

SetTeacherIds sets TeacherIds field to given value.

### HasTeacherIds

`func (o *PlannedCoursePatchPayload) HasTeacherIds() bool`

HasTeacherIds returns a boolean if a field has been set.

### GetCustom

`func (o *PlannedCoursePatchPayload) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *PlannedCoursePatchPayload) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *PlannedCoursePatchPayload) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *PlannedCoursePatchPayload) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomAssociations

`func (o *PlannedCoursePatchPayload) GetCustomAssociations() []CustomAssociationsRecordsPayload`

GetCustomAssociations returns the CustomAssociations field if non-nil, zero value otherwise.

### GetCustomAssociationsOk

`func (o *PlannedCoursePatchPayload) GetCustomAssociationsOk() (*[]CustomAssociationsRecordsPayload, bool)`

GetCustomAssociationsOk returns a tuple with the CustomAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssociations

`func (o *PlannedCoursePatchPayload) SetCustomAssociations(v []CustomAssociationsRecordsPayload)`

SetCustomAssociations sets CustomAssociations field to given value.

### HasCustomAssociations

`func (o *PlannedCoursePatchPayload) HasCustomAssociations() bool`

HasCustomAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


