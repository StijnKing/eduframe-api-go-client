# PlanningEvent

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

## Methods

### NewPlanningEvent

`func NewPlanningEvent(id int32, meetingId int32, meetingLocationId NullableInt32, name NullableString, startDateTime string, endDateTime string, planningMeetingLocationIds []int32, ) *PlanningEvent`

NewPlanningEvent instantiates a new PlanningEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanningEventWithDefaults

`func NewPlanningEventWithDefaults() *PlanningEvent`

NewPlanningEventWithDefaults instantiates a new PlanningEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanningEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanningEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanningEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetMeetingId

`func (o *PlanningEvent) GetMeetingId() int32`

GetMeetingId returns the MeetingId field if non-nil, zero value otherwise.

### GetMeetingIdOk

`func (o *PlanningEvent) GetMeetingIdOk() (*int32, bool)`

GetMeetingIdOk returns a tuple with the MeetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingId

`func (o *PlanningEvent) SetMeetingId(v int32)`

SetMeetingId sets MeetingId field to given value.


### GetMeetingLocationId

`func (o *PlanningEvent) GetMeetingLocationId() int32`

GetMeetingLocationId returns the MeetingLocationId field if non-nil, zero value otherwise.

### GetMeetingLocationIdOk

`func (o *PlanningEvent) GetMeetingLocationIdOk() (*int32, bool)`

GetMeetingLocationIdOk returns a tuple with the MeetingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingLocationId

`func (o *PlanningEvent) SetMeetingLocationId(v int32)`

SetMeetingLocationId sets MeetingLocationId field to given value.


### SetMeetingLocationIdNil

`func (o *PlanningEvent) SetMeetingLocationIdNil(b bool)`

 SetMeetingLocationIdNil sets the value for MeetingLocationId to be an explicit nil

### UnsetMeetingLocationId
`func (o *PlanningEvent) UnsetMeetingLocationId()`

UnsetMeetingLocationId ensures that no value is present for MeetingLocationId, not even an explicit nil
### GetName

`func (o *PlanningEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanningEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanningEvent) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PlanningEvent) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PlanningEvent) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartDateTime

`func (o *PlanningEvent) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *PlanningEvent) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *PlanningEvent) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *PlanningEvent) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *PlanningEvent) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *PlanningEvent) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.


### GetPlanningMeetingLocationIds

`func (o *PlanningEvent) GetPlanningMeetingLocationIds() []int32`

GetPlanningMeetingLocationIds returns the PlanningMeetingLocationIds field if non-nil, zero value otherwise.

### GetPlanningMeetingLocationIdsOk

`func (o *PlanningEvent) GetPlanningMeetingLocationIdsOk() (*[]int32, bool)`

GetPlanningMeetingLocationIdsOk returns a tuple with the PlanningMeetingLocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningMeetingLocationIds

`func (o *PlanningEvent) SetPlanningMeetingLocationIds(v []int32)`

SetPlanningMeetingLocationIds sets PlanningMeetingLocationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


