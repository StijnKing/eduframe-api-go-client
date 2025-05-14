# MeetingWithIncludes

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
**PlanningAttendees** | [**[]PlanningAttendee**](PlanningAttendee.md) |  | 
**PlanningMaterials** | Pointer to [**[]PlanningMaterial**](PlanningMaterial.md) |  | [optional] 

## Methods

### NewMeetingWithIncludes

`func NewMeetingWithIncludes(id int32, name NullableString, plannedCourseId int32, startDateTime time.Time, endDateTime time.Time, meetingLocationId NullableInt32, planningMeetingLocationIds []int32, description NullableString, descriptionDashboard NullableString, updatedAt time.Time, createdAt time.Time, planningAttendees []PlanningAttendee, ) *MeetingWithIncludes`

NewMeetingWithIncludes instantiates a new MeetingWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingWithIncludesWithDefaults

`func NewMeetingWithIncludesWithDefaults() *MeetingWithIncludes`

NewMeetingWithIncludesWithDefaults instantiates a new MeetingWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MeetingWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeetingWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeetingWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *MeetingWithIncludes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeetingWithIncludes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeetingWithIncludes) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *MeetingWithIncludes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MeetingWithIncludes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPlannedCourseId

`func (o *MeetingWithIncludes) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *MeetingWithIncludes) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *MeetingWithIncludes) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.


### GetStartDateTime

`func (o *MeetingWithIncludes) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MeetingWithIncludes) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MeetingWithIncludes) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *MeetingWithIncludes) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MeetingWithIncludes) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MeetingWithIncludes) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.


### GetMeetingLocationId

`func (o *MeetingWithIncludes) GetMeetingLocationId() int32`

GetMeetingLocationId returns the MeetingLocationId field if non-nil, zero value otherwise.

### GetMeetingLocationIdOk

`func (o *MeetingWithIncludes) GetMeetingLocationIdOk() (*int32, bool)`

GetMeetingLocationIdOk returns a tuple with the MeetingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingLocationId

`func (o *MeetingWithIncludes) SetMeetingLocationId(v int32)`

SetMeetingLocationId sets MeetingLocationId field to given value.


### SetMeetingLocationIdNil

`func (o *MeetingWithIncludes) SetMeetingLocationIdNil(b bool)`

 SetMeetingLocationIdNil sets the value for MeetingLocationId to be an explicit nil

### UnsetMeetingLocationId
`func (o *MeetingWithIncludes) UnsetMeetingLocationId()`

UnsetMeetingLocationId ensures that no value is present for MeetingLocationId, not even an explicit nil
### GetPlanningMeetingLocationIds

`func (o *MeetingWithIncludes) GetPlanningMeetingLocationIds() []int32`

GetPlanningMeetingLocationIds returns the PlanningMeetingLocationIds field if non-nil, zero value otherwise.

### GetPlanningMeetingLocationIdsOk

`func (o *MeetingWithIncludes) GetPlanningMeetingLocationIdsOk() (*[]int32, bool)`

GetPlanningMeetingLocationIdsOk returns a tuple with the PlanningMeetingLocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningMeetingLocationIds

`func (o *MeetingWithIncludes) SetPlanningMeetingLocationIds(v []int32)`

SetPlanningMeetingLocationIds sets PlanningMeetingLocationIds field to given value.


### GetDescription

`func (o *MeetingWithIncludes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MeetingWithIncludes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MeetingWithIncludes) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *MeetingWithIncludes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MeetingWithIncludes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDescriptionDashboard

`func (o *MeetingWithIncludes) GetDescriptionDashboard() string`

GetDescriptionDashboard returns the DescriptionDashboard field if non-nil, zero value otherwise.

### GetDescriptionDashboardOk

`func (o *MeetingWithIncludes) GetDescriptionDashboardOk() (*string, bool)`

GetDescriptionDashboardOk returns a tuple with the DescriptionDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionDashboard

`func (o *MeetingWithIncludes) SetDescriptionDashboard(v string)`

SetDescriptionDashboard sets DescriptionDashboard field to given value.


### SetDescriptionDashboardNil

`func (o *MeetingWithIncludes) SetDescriptionDashboardNil(b bool)`

 SetDescriptionDashboardNil sets the value for DescriptionDashboard to be an explicit nil

### UnsetDescriptionDashboard
`func (o *MeetingWithIncludes) UnsetDescriptionDashboard()`

UnsetDescriptionDashboard ensures that no value is present for DescriptionDashboard, not even an explicit nil
### GetUpdatedAt

`func (o *MeetingWithIncludes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MeetingWithIncludes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MeetingWithIncludes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *MeetingWithIncludes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MeetingWithIncludes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MeetingWithIncludes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPlanningAttendees

`func (o *MeetingWithIncludes) GetPlanningAttendees() []PlanningAttendee`

GetPlanningAttendees returns the PlanningAttendees field if non-nil, zero value otherwise.

### GetPlanningAttendeesOk

`func (o *MeetingWithIncludes) GetPlanningAttendeesOk() (*[]PlanningAttendee, bool)`

GetPlanningAttendeesOk returns a tuple with the PlanningAttendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningAttendees

`func (o *MeetingWithIncludes) SetPlanningAttendees(v []PlanningAttendee)`

SetPlanningAttendees sets PlanningAttendees field to given value.


### GetPlanningMaterials

`func (o *MeetingWithIncludes) GetPlanningMaterials() []PlanningMaterial`

GetPlanningMaterials returns the PlanningMaterials field if non-nil, zero value otherwise.

### GetPlanningMaterialsOk

`func (o *MeetingWithIncludes) GetPlanningMaterialsOk() (*[]PlanningMaterial, bool)`

GetPlanningMaterialsOk returns a tuple with the PlanningMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningMaterials

`func (o *MeetingWithIncludes) SetPlanningMaterials(v []PlanningMaterial)`

SetPlanningMaterials sets PlanningMaterials field to given value.

### HasPlanningMaterials

`func (o *MeetingWithIncludes) HasPlanningMaterials() bool`

HasPlanningMaterials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


