# Meeting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the meeting. | [readonly] 
**Name** | **NullableString** | The name of the meeting. | 
**PlannedCourseId** | **int32** | Unique identifier of the planned course. | 
**StartDateTime** | **time.Time** | Date and time when the meeting is starting. | 
**EndDateTime** | **time.Time** | The date and time when the meeting is ending. | 
**MeetingLocationId** | **NullableInt32** | Unique identifier of the meeting location. | 
**PlanningMeetingLocationIds** | **[]int32** |  | 
**Description** | **NullableString** | The description of the meeting. | 
**DescriptionDashboard** | **NullableString** | The description that will be shown in the dashboard. | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewMeeting

`func NewMeeting(id int32, name NullableString, plannedCourseId int32, startDateTime time.Time, endDateTime time.Time, meetingLocationId NullableInt32, planningMeetingLocationIds []int32, description NullableString, descriptionDashboard NullableString, updatedAt time.Time, createdAt time.Time, ) *Meeting`

NewMeeting instantiates a new Meeting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingWithDefaults

`func NewMeetingWithDefaults() *Meeting`

NewMeetingWithDefaults instantiates a new Meeting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Meeting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Meeting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Meeting) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Meeting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Meeting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Meeting) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Meeting) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Meeting) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPlannedCourseId

`func (o *Meeting) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *Meeting) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *Meeting) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.


### GetStartDateTime

`func (o *Meeting) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *Meeting) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *Meeting) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *Meeting) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *Meeting) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *Meeting) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.


### GetMeetingLocationId

`func (o *Meeting) GetMeetingLocationId() int32`

GetMeetingLocationId returns the MeetingLocationId field if non-nil, zero value otherwise.

### GetMeetingLocationIdOk

`func (o *Meeting) GetMeetingLocationIdOk() (*int32, bool)`

GetMeetingLocationIdOk returns a tuple with the MeetingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingLocationId

`func (o *Meeting) SetMeetingLocationId(v int32)`

SetMeetingLocationId sets MeetingLocationId field to given value.


### SetMeetingLocationIdNil

`func (o *Meeting) SetMeetingLocationIdNil(b bool)`

 SetMeetingLocationIdNil sets the value for MeetingLocationId to be an explicit nil

### UnsetMeetingLocationId
`func (o *Meeting) UnsetMeetingLocationId()`

UnsetMeetingLocationId ensures that no value is present for MeetingLocationId, not even an explicit nil
### GetPlanningMeetingLocationIds

`func (o *Meeting) GetPlanningMeetingLocationIds() []int32`

GetPlanningMeetingLocationIds returns the PlanningMeetingLocationIds field if non-nil, zero value otherwise.

### GetPlanningMeetingLocationIdsOk

`func (o *Meeting) GetPlanningMeetingLocationIdsOk() (*[]int32, bool)`

GetPlanningMeetingLocationIdsOk returns a tuple with the PlanningMeetingLocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningMeetingLocationIds

`func (o *Meeting) SetPlanningMeetingLocationIds(v []int32)`

SetPlanningMeetingLocationIds sets PlanningMeetingLocationIds field to given value.


### GetDescription

`func (o *Meeting) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Meeting) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Meeting) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Meeting) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Meeting) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDescriptionDashboard

`func (o *Meeting) GetDescriptionDashboard() string`

GetDescriptionDashboard returns the DescriptionDashboard field if non-nil, zero value otherwise.

### GetDescriptionDashboardOk

`func (o *Meeting) GetDescriptionDashboardOk() (*string, bool)`

GetDescriptionDashboardOk returns a tuple with the DescriptionDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionDashboard

`func (o *Meeting) SetDescriptionDashboard(v string)`

SetDescriptionDashboard sets DescriptionDashboard field to given value.


### SetDescriptionDashboardNil

`func (o *Meeting) SetDescriptionDashboardNil(b bool)`

 SetDescriptionDashboardNil sets the value for DescriptionDashboard to be an explicit nil

### UnsetDescriptionDashboard
`func (o *Meeting) UnsetDescriptionDashboard()`

UnsetDescriptionDashboard ensures that no value is present for DescriptionDashboard, not even an explicit nil
### GetUpdatedAt

`func (o *Meeting) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Meeting) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Meeting) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Meeting) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Meeting) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Meeting) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


