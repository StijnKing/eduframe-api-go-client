# PlanningEventWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the meeting template | [readonly] 
**MeetingId** | **int32** | Identifier of the meeting | 
**MeetingLocationId** | **NullableInt32** | Unique identifier of the meeting location. | 
**Name** | **NullableString** | Name of the event | 
**StartDateTime** | **string** | Date and time when the event is starting | 
**EndDateTime** | **string** | The date and time when the event is ending | 
**PlanningMeetingLocationIds** | **[]int32** |  | 
**PlanningAttendees** | [**[]PlanningAttendee**](PlanningAttendee.md) |  | 
**PlanningMaterials** | Pointer to [**[]PlanningMaterial**](PlanningMaterial.md) |  | [optional] 

## Methods

### NewPlanningEventWithIncludes

`func NewPlanningEventWithIncludes(id int32, meetingId int32, meetingLocationId NullableInt32, name NullableString, startDateTime string, endDateTime string, planningMeetingLocationIds []int32, planningAttendees []PlanningAttendee, ) *PlanningEventWithIncludes`

NewPlanningEventWithIncludes instantiates a new PlanningEventWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanningEventWithIncludesWithDefaults

`func NewPlanningEventWithIncludesWithDefaults() *PlanningEventWithIncludes`

NewPlanningEventWithIncludesWithDefaults instantiates a new PlanningEventWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanningEventWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanningEventWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanningEventWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetMeetingId

`func (o *PlanningEventWithIncludes) GetMeetingId() int32`

GetMeetingId returns the MeetingId field if non-nil, zero value otherwise.

### GetMeetingIdOk

`func (o *PlanningEventWithIncludes) GetMeetingIdOk() (*int32, bool)`

GetMeetingIdOk returns a tuple with the MeetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingId

`func (o *PlanningEventWithIncludes) SetMeetingId(v int32)`

SetMeetingId sets MeetingId field to given value.


### GetMeetingLocationId

`func (o *PlanningEventWithIncludes) GetMeetingLocationId() int32`

GetMeetingLocationId returns the MeetingLocationId field if non-nil, zero value otherwise.

### GetMeetingLocationIdOk

`func (o *PlanningEventWithIncludes) GetMeetingLocationIdOk() (*int32, bool)`

GetMeetingLocationIdOk returns a tuple with the MeetingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingLocationId

`func (o *PlanningEventWithIncludes) SetMeetingLocationId(v int32)`

SetMeetingLocationId sets MeetingLocationId field to given value.


### SetMeetingLocationIdNil

`func (o *PlanningEventWithIncludes) SetMeetingLocationIdNil(b bool)`

 SetMeetingLocationIdNil sets the value for MeetingLocationId to be an explicit nil

### UnsetMeetingLocationId
`func (o *PlanningEventWithIncludes) UnsetMeetingLocationId()`

UnsetMeetingLocationId ensures that no value is present for MeetingLocationId, not even an explicit nil
### GetName

`func (o *PlanningEventWithIncludes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanningEventWithIncludes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanningEventWithIncludes) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PlanningEventWithIncludes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PlanningEventWithIncludes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartDateTime

`func (o *PlanningEventWithIncludes) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *PlanningEventWithIncludes) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *PlanningEventWithIncludes) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *PlanningEventWithIncludes) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *PlanningEventWithIncludes) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *PlanningEventWithIncludes) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.


### GetPlanningMeetingLocationIds

`func (o *PlanningEventWithIncludes) GetPlanningMeetingLocationIds() []int32`

GetPlanningMeetingLocationIds returns the PlanningMeetingLocationIds field if non-nil, zero value otherwise.

### GetPlanningMeetingLocationIdsOk

`func (o *PlanningEventWithIncludes) GetPlanningMeetingLocationIdsOk() (*[]int32, bool)`

GetPlanningMeetingLocationIdsOk returns a tuple with the PlanningMeetingLocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningMeetingLocationIds

`func (o *PlanningEventWithIncludes) SetPlanningMeetingLocationIds(v []int32)`

SetPlanningMeetingLocationIds sets PlanningMeetingLocationIds field to given value.


### GetPlanningAttendees

`func (o *PlanningEventWithIncludes) GetPlanningAttendees() []PlanningAttendee`

GetPlanningAttendees returns the PlanningAttendees field if non-nil, zero value otherwise.

### GetPlanningAttendeesOk

`func (o *PlanningEventWithIncludes) GetPlanningAttendeesOk() (*[]PlanningAttendee, bool)`

GetPlanningAttendeesOk returns a tuple with the PlanningAttendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningAttendees

`func (o *PlanningEventWithIncludes) SetPlanningAttendees(v []PlanningAttendee)`

SetPlanningAttendees sets PlanningAttendees field to given value.


### GetPlanningMaterials

`func (o *PlanningEventWithIncludes) GetPlanningMaterials() []PlanningMaterial`

GetPlanningMaterials returns the PlanningMaterials field if non-nil, zero value otherwise.

### GetPlanningMaterialsOk

`func (o *PlanningEventWithIncludes) GetPlanningMaterialsOk() (*[]PlanningMaterial, bool)`

GetPlanningMaterialsOk returns a tuple with the PlanningMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningMaterials

`func (o *PlanningEventWithIncludes) SetPlanningMaterials(v []PlanningMaterial)`

SetPlanningMaterials sets PlanningMaterials field to given value.

### HasPlanningMaterials

`func (o *PlanningEventWithIncludes) HasPlanningMaterials() bool`

HasPlanningMaterials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


