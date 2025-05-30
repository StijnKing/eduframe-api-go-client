/*
Eduframe - API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eduframe

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PlanningEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanningEvent{}

// PlanningEvent struct for PlanningEvent
type PlanningEvent struct {
	// Unique identifier of the meeting template
	Id int32 `json:"id"`
	// Identifier of the meeting
	MeetingId int32 `json:"meeting_id"`
	// Unique identifier of the meeting location.
	MeetingLocationId NullableInt32 `json:"meeting_location_id"`
	// Name of the event
	Name NullableString `json:"name"`
	// Date and time when the event is starting
	StartDateTime string `json:"start_date_time"`
	// The date and time when the event is ending
	EndDateTime string `json:"end_date_time"`
	PlanningMeetingLocationIds []int32 `json:"planning_meeting_location_ids"`
}

type _PlanningEvent PlanningEvent

// NewPlanningEvent instantiates a new PlanningEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanningEvent(id int32, meetingId int32, meetingLocationId NullableInt32, name NullableString, startDateTime string, endDateTime string, planningMeetingLocationIds []int32) *PlanningEvent {
	this := PlanningEvent{}
	this.Id = id
	this.MeetingId = meetingId
	this.MeetingLocationId = meetingLocationId
	this.Name = name
	this.StartDateTime = startDateTime
	this.EndDateTime = endDateTime
	this.PlanningMeetingLocationIds = planningMeetingLocationIds
	return &this
}

// NewPlanningEventWithDefaults instantiates a new PlanningEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanningEventWithDefaults() *PlanningEvent {
	this := PlanningEvent{}
	return &this
}

// GetId returns the Id field value
func (o *PlanningEvent) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PlanningEvent) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PlanningEvent) SetId(v int32) {
	o.Id = v
}

// GetMeetingId returns the MeetingId field value
func (o *PlanningEvent) GetMeetingId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MeetingId
}

// GetMeetingIdOk returns a tuple with the MeetingId field value
// and a boolean to check if the value has been set.
func (o *PlanningEvent) GetMeetingIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeetingId, true
}

// SetMeetingId sets field value
func (o *PlanningEvent) SetMeetingId(v int32) {
	o.MeetingId = v
}

// GetMeetingLocationId returns the MeetingLocationId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PlanningEvent) GetMeetingLocationId() int32 {
	if o == nil || o.MeetingLocationId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MeetingLocationId.Get()
}

// GetMeetingLocationIdOk returns a tuple with the MeetingLocationId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlanningEvent) GetMeetingLocationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeetingLocationId.Get(), o.MeetingLocationId.IsSet()
}

// SetMeetingLocationId sets field value
func (o *PlanningEvent) SetMeetingLocationId(v int32) {
	o.MeetingLocationId.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PlanningEvent) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlanningEvent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *PlanningEvent) SetName(v string) {
	o.Name.Set(&v)
}

// GetStartDateTime returns the StartDateTime field value
func (o *PlanningEvent) GetStartDateTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value
// and a boolean to check if the value has been set.
func (o *PlanningEvent) GetStartDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDateTime, true
}

// SetStartDateTime sets field value
func (o *PlanningEvent) SetStartDateTime(v string) {
	o.StartDateTime = v
}

// GetEndDateTime returns the EndDateTime field value
func (o *PlanningEvent) GetEndDateTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value
// and a boolean to check if the value has been set.
func (o *PlanningEvent) GetEndDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDateTime, true
}

// SetEndDateTime sets field value
func (o *PlanningEvent) SetEndDateTime(v string) {
	o.EndDateTime = v
}

// GetPlanningMeetingLocationIds returns the PlanningMeetingLocationIds field value
func (o *PlanningEvent) GetPlanningMeetingLocationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.PlanningMeetingLocationIds
}

// GetPlanningMeetingLocationIdsOk returns a tuple with the PlanningMeetingLocationIds field value
// and a boolean to check if the value has been set.
func (o *PlanningEvent) GetPlanningMeetingLocationIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanningMeetingLocationIds, true
}

// SetPlanningMeetingLocationIds sets field value
func (o *PlanningEvent) SetPlanningMeetingLocationIds(v []int32) {
	o.PlanningMeetingLocationIds = v
}

func (o PlanningEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanningEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["meeting_id"] = o.MeetingId
	toSerialize["meeting_location_id"] = o.MeetingLocationId.Get()
	toSerialize["name"] = o.Name.Get()
	toSerialize["start_date_time"] = o.StartDateTime
	toSerialize["end_date_time"] = o.EndDateTime
	toSerialize["planning_meeting_location_ids"] = o.PlanningMeetingLocationIds
	return toSerialize, nil
}

func (o *PlanningEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"meeting_id",
		"meeting_location_id",
		"name",
		"start_date_time",
		"end_date_time",
		"planning_meeting_location_ids",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPlanningEvent := _PlanningEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPlanningEvent)

	if err != nil {
		return err
	}

	*o = PlanningEvent(varPlanningEvent)

	return err
}

type NullablePlanningEvent struct {
	value *PlanningEvent
	isSet bool
}

func (v NullablePlanningEvent) Get() *PlanningEvent {
	return v.value
}

func (v *NullablePlanningEvent) Set(val *PlanningEvent) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanningEvent) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanningEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanningEvent(val *PlanningEvent) *NullablePlanningEvent {
	return &NullablePlanningEvent{value: val, isSet: true}
}

func (v NullablePlanningEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanningEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


