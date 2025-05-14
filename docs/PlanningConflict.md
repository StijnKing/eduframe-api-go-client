# PlanningConflict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Type of the resource | 
**ResourceId** | **int32** | Unique identifier of the resource | 
**ResourceName** | **string** | Name of the resource | 
**EventId** | **int32** | Unique identifier of the event | 
**EventType** | **string** | Type of the event | 
**EventStartDateTime** | **string** | Timestamp of the start of the event | 
**EventEndDateTime** | **string** | Timestamp of the end of the event | 
**ConflictingEventId** | **int32** | Unique identifier of the conflicting event | 
**ConflictingEventType** | **string** | Type of the conflicting event | 
**ConflictingEventStartDateTime** | **string** | Timestamp of the start of the conflicting event | 
**ConflictingEventEndDateTime** | **string** | Timestamp of the end of the conflicting event | 
**StartDateTime** | **string** | Timestamp of the start of the conflict | 
**EndDateTime** | **string** | Timestamp of the end of the conflict | 

## Methods

### NewPlanningConflict

`func NewPlanningConflict(resourceType string, resourceId int32, resourceName string, eventId int32, eventType string, eventStartDateTime string, eventEndDateTime string, conflictingEventId int32, conflictingEventType string, conflictingEventStartDateTime string, conflictingEventEndDateTime string, startDateTime string, endDateTime string, ) *PlanningConflict`

NewPlanningConflict instantiates a new PlanningConflict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanningConflictWithDefaults

`func NewPlanningConflictWithDefaults() *PlanningConflict`

NewPlanningConflictWithDefaults instantiates a new PlanningConflict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *PlanningConflict) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *PlanningConflict) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *PlanningConflict) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetResourceId

`func (o *PlanningConflict) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PlanningConflict) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PlanningConflict) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.


### GetResourceName

`func (o *PlanningConflict) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PlanningConflict) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PlanningConflict) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetEventId

`func (o *PlanningConflict) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PlanningConflict) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PlanningConflict) SetEventId(v int32)`

SetEventId sets EventId field to given value.


### GetEventType

`func (o *PlanningConflict) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PlanningConflict) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PlanningConflict) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventStartDateTime

`func (o *PlanningConflict) GetEventStartDateTime() string`

GetEventStartDateTime returns the EventStartDateTime field if non-nil, zero value otherwise.

### GetEventStartDateTimeOk

`func (o *PlanningConflict) GetEventStartDateTimeOk() (*string, bool)`

GetEventStartDateTimeOk returns a tuple with the EventStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStartDateTime

`func (o *PlanningConflict) SetEventStartDateTime(v string)`

SetEventStartDateTime sets EventStartDateTime field to given value.


### GetEventEndDateTime

`func (o *PlanningConflict) GetEventEndDateTime() string`

GetEventEndDateTime returns the EventEndDateTime field if non-nil, zero value otherwise.

### GetEventEndDateTimeOk

`func (o *PlanningConflict) GetEventEndDateTimeOk() (*string, bool)`

GetEventEndDateTimeOk returns a tuple with the EventEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventEndDateTime

`func (o *PlanningConflict) SetEventEndDateTime(v string)`

SetEventEndDateTime sets EventEndDateTime field to given value.


### GetConflictingEventId

`func (o *PlanningConflict) GetConflictingEventId() int32`

GetConflictingEventId returns the ConflictingEventId field if non-nil, zero value otherwise.

### GetConflictingEventIdOk

`func (o *PlanningConflict) GetConflictingEventIdOk() (*int32, bool)`

GetConflictingEventIdOk returns a tuple with the ConflictingEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingEventId

`func (o *PlanningConflict) SetConflictingEventId(v int32)`

SetConflictingEventId sets ConflictingEventId field to given value.


### GetConflictingEventType

`func (o *PlanningConflict) GetConflictingEventType() string`

GetConflictingEventType returns the ConflictingEventType field if non-nil, zero value otherwise.

### GetConflictingEventTypeOk

`func (o *PlanningConflict) GetConflictingEventTypeOk() (*string, bool)`

GetConflictingEventTypeOk returns a tuple with the ConflictingEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingEventType

`func (o *PlanningConflict) SetConflictingEventType(v string)`

SetConflictingEventType sets ConflictingEventType field to given value.


### GetConflictingEventStartDateTime

`func (o *PlanningConflict) GetConflictingEventStartDateTime() string`

GetConflictingEventStartDateTime returns the ConflictingEventStartDateTime field if non-nil, zero value otherwise.

### GetConflictingEventStartDateTimeOk

`func (o *PlanningConflict) GetConflictingEventStartDateTimeOk() (*string, bool)`

GetConflictingEventStartDateTimeOk returns a tuple with the ConflictingEventStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingEventStartDateTime

`func (o *PlanningConflict) SetConflictingEventStartDateTime(v string)`

SetConflictingEventStartDateTime sets ConflictingEventStartDateTime field to given value.


### GetConflictingEventEndDateTime

`func (o *PlanningConflict) GetConflictingEventEndDateTime() string`

GetConflictingEventEndDateTime returns the ConflictingEventEndDateTime field if non-nil, zero value otherwise.

### GetConflictingEventEndDateTimeOk

`func (o *PlanningConflict) GetConflictingEventEndDateTimeOk() (*string, bool)`

GetConflictingEventEndDateTimeOk returns a tuple with the ConflictingEventEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingEventEndDateTime

`func (o *PlanningConflict) SetConflictingEventEndDateTime(v string)`

SetConflictingEventEndDateTime sets ConflictingEventEndDateTime field to given value.


### GetStartDateTime

`func (o *PlanningConflict) GetStartDateTime() string`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *PlanningConflict) GetStartDateTimeOk() (*string, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *PlanningConflict) SetStartDateTime(v string)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *PlanningConflict) GetEndDateTime() string`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *PlanningConflict) GetEndDateTimeOk() (*string, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *PlanningConflict) SetEndDateTime(v string)`

SetEndDateTime sets EndDateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


