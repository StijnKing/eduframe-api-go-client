/*
Eduframe - API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eduframe

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the MeetingWithIncludes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeetingWithIncludes{}

// MeetingWithIncludes struct for MeetingWithIncludes
type MeetingWithIncludes struct {
	// Unique identifier of the meeting.
	Id int32 `json:"id"`
	// The name of the meeting.
	Name NullableString `json:"name"`
	// Unique identifier of the planned course.
	PlannedCourseId int32 `json:"planned_course_id"`
	// Date and time when the meeting is starting.
	StartDateTime time.Time `json:"start_date_time"`
	// The date and time when the meeting is ending.
	EndDateTime time.Time `json:"end_date_time"`
	// Unique identifier of the meeting location.
	MeetingLocationId NullableInt32 `json:"meeting_location_id"`
	PlanningMeetingLocationIds []int32 `json:"planning_meeting_location_ids"`
	// The description of the meeting.
	Description NullableString `json:"description"`
	// The description that will be shown in the dashboard.
	DescriptionDashboard NullableString `json:"description_dashboard"`
	// Timestamp of last update.
	UpdatedAt time.Time `json:"updated_at"`
	// Timestamp of creation.
	CreatedAt time.Time `json:"created_at"`
	PlanningAttendees []PlanningAttendee `json:"planning_attendees"`
	PlanningMaterials []PlanningMaterial `json:"planning_materials,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MeetingWithIncludes MeetingWithIncludes

// NewMeetingWithIncludes instantiates a new MeetingWithIncludes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeetingWithIncludes(id int32, name NullableString, plannedCourseId int32, startDateTime time.Time, endDateTime time.Time, meetingLocationId NullableInt32, planningMeetingLocationIds []int32, description NullableString, descriptionDashboard NullableString, updatedAt time.Time, createdAt time.Time, planningAttendees []PlanningAttendee) *MeetingWithIncludes {
	this := MeetingWithIncludes{}
	this.Id = id
	this.Name = name
	this.PlannedCourseId = plannedCourseId
	this.StartDateTime = startDateTime
	this.EndDateTime = endDateTime
	this.MeetingLocationId = meetingLocationId
	this.PlanningMeetingLocationIds = planningMeetingLocationIds
	this.Description = description
	this.DescriptionDashboard = descriptionDashboard
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	this.PlanningAttendees = planningAttendees
	return &this
}

// NewMeetingWithIncludesWithDefaults instantiates a new MeetingWithIncludes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeetingWithIncludesWithDefaults() *MeetingWithIncludes {
	this := MeetingWithIncludes{}
	return &this
}

// GetId returns the Id field value
func (o *MeetingWithIncludes) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MeetingWithIncludes) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MeetingWithIncludes) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MeetingWithIncludes) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeetingWithIncludes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *MeetingWithIncludes) SetName(v string) {
	o.Name.Set(&v)
}

// GetPlannedCourseId returns the PlannedCourseId field value
func (o *MeetingWithIncludes) GetPlannedCourseId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlannedCourseId
}

// GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field value
// and a boolean to check if the value has been set.
func (o *MeetingWithIncludes) GetPlannedCourseIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlannedCourseId, true
}

// SetPlannedCourseId sets field value
func (o *MeetingWithIncludes) SetPlannedCourseId(v int32) {
	o.PlannedCourseId = v
}

// GetStartDateTime returns the StartDateTime field value
func (o *MeetingWithIncludes) GetStartDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value
// and a boolean to check if the value has been set.
func (o *MeetingWithIncludes) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDateTime, true
}

// SetStartDateTime sets field value
func (o *MeetingWithIncludes) SetStartDateTime(v time.Time) {
	o.StartDateTime = v
}

// GetEndDateTime returns the EndDateTime field value
func (o *MeetingWithIncludes) GetEndDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value
// and a boolean to check if the value has been set.
func (o *MeetingWithIncludes) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDateTime, true
}

// SetEndDateTime sets field value
func (o *MeetingWithIncludes) SetEndDateTime(v time.Time) {
	o.EndDateTime = v
}

// GetMeetingLocationId returns the MeetingLocationId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *MeetingWithIncludes) GetMeetingLocationId() int32 {
	if o == nil || o.MeetingLocationId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MeetingLocationId.Get()
}

// GetMeetingLocationIdOk returns a tuple with the MeetingLocationId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeetingWithIncludes) GetMeetingLocationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeetingLocationId.Get(), o.MeetingLocationId.IsSet()
}

// SetMeetingLocationId sets field value
func (o *MeetingWithIncludes) SetMeetingLocationId(v int32) {
	o.MeetingLocationId.Set(&v)
}

// GetPlanningMeetingLocationIds returns the PlanningMeetingLocationIds field value
func (o *MeetingWithIncludes) GetPlanningMeetingLocationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.PlanningMeetingLocationIds
}

// GetPlanningMeetingLocationIdsOk returns a tuple with the PlanningMeetingLocationIds field value
// and a boolean to check if the value has been set.
func (o *MeetingWithIncludes) GetPlanningMeetingLocationIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanningMeetingLocationIds, true
}

// SetPlanningMeetingLocationIds sets field value
func (o *MeetingWithIncludes) SetPlanningMeetingLocationIds(v []int32) {
	o.PlanningMeetingLocationIds = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MeetingWithIncludes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeetingWithIncludes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *MeetingWithIncludes) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetDescriptionDashboard returns the DescriptionDashboard field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MeetingWithIncludes) GetDescriptionDashboard() string {
	if o == nil || o.DescriptionDashboard.Get() == nil {
		var ret string
		return ret
	}

	return *o.DescriptionDashboard.Get()
}

// GetDescriptionDashboardOk returns a tuple with the DescriptionDashboard field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeetingWithIncludes) GetDescriptionDashboardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DescriptionDashboard.Get(), o.DescriptionDashboard.IsSet()
}

// SetDescriptionDashboard sets field value
func (o *MeetingWithIncludes) SetDescriptionDashboard(v string) {
	o.DescriptionDashboard.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *MeetingWithIncludes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *MeetingWithIncludes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *MeetingWithIncludes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MeetingWithIncludes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MeetingWithIncludes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MeetingWithIncludes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetPlanningAttendees returns the PlanningAttendees field value
func (o *MeetingWithIncludes) GetPlanningAttendees() []PlanningAttendee {
	if o == nil {
		var ret []PlanningAttendee
		return ret
	}

	return o.PlanningAttendees
}

// GetPlanningAttendeesOk returns a tuple with the PlanningAttendees field value
// and a boolean to check if the value has been set.
func (o *MeetingWithIncludes) GetPlanningAttendeesOk() ([]PlanningAttendee, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanningAttendees, true
}

// SetPlanningAttendees sets field value
func (o *MeetingWithIncludes) SetPlanningAttendees(v []PlanningAttendee) {
	o.PlanningAttendees = v
}

// GetPlanningMaterials returns the PlanningMaterials field value if set, zero value otherwise.
func (o *MeetingWithIncludes) GetPlanningMaterials() []PlanningMaterial {
	if o == nil || IsNil(o.PlanningMaterials) {
		var ret []PlanningMaterial
		return ret
	}
	return o.PlanningMaterials
}

// GetPlanningMaterialsOk returns a tuple with the PlanningMaterials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingWithIncludes) GetPlanningMaterialsOk() ([]PlanningMaterial, bool) {
	if o == nil || IsNil(o.PlanningMaterials) {
		return nil, false
	}
	return o.PlanningMaterials, true
}

// HasPlanningMaterials returns a boolean if a field has been set.
func (o *MeetingWithIncludes) HasPlanningMaterials() bool {
	if o != nil && !IsNil(o.PlanningMaterials) {
		return true
	}

	return false
}

// SetPlanningMaterials gets a reference to the given []PlanningMaterial and assigns it to the PlanningMaterials field.
func (o *MeetingWithIncludes) SetPlanningMaterials(v []PlanningMaterial) {
	o.PlanningMaterials = v
}

func (o MeetingWithIncludes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeetingWithIncludes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["planned_course_id"] = o.PlannedCourseId
	toSerialize["start_date_time"] = o.StartDateTime
	toSerialize["end_date_time"] = o.EndDateTime
	toSerialize["meeting_location_id"] = o.MeetingLocationId.Get()
	toSerialize["planning_meeting_location_ids"] = o.PlanningMeetingLocationIds
	toSerialize["description"] = o.Description.Get()
	toSerialize["description_dashboard"] = o.DescriptionDashboard.Get()
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["planning_attendees"] = o.PlanningAttendees
	if !IsNil(o.PlanningMaterials) {
		toSerialize["planning_materials"] = o.PlanningMaterials
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MeetingWithIncludes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"planned_course_id",
		"start_date_time",
		"end_date_time",
		"meeting_location_id",
		"planning_meeting_location_ids",
		"description",
		"description_dashboard",
		"updated_at",
		"created_at",
		"planning_attendees",
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

	varMeetingWithIncludes := _MeetingWithIncludes{}

	err = json.Unmarshal(data, &varMeetingWithIncludes)

	if err != nil {
		return err
	}

	*o = MeetingWithIncludes(varMeetingWithIncludes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "planned_course_id")
		delete(additionalProperties, "start_date_time")
		delete(additionalProperties, "end_date_time")
		delete(additionalProperties, "meeting_location_id")
		delete(additionalProperties, "planning_meeting_location_ids")
		delete(additionalProperties, "description")
		delete(additionalProperties, "description_dashboard")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "planning_attendees")
		delete(additionalProperties, "planning_materials")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMeetingWithIncludes struct {
	value *MeetingWithIncludes
	isSet bool
}

func (v NullableMeetingWithIncludes) Get() *MeetingWithIncludes {
	return v.value
}

func (v *NullableMeetingWithIncludes) Set(val *MeetingWithIncludes) {
	v.value = val
	v.isSet = true
}

func (v NullableMeetingWithIncludes) IsSet() bool {
	return v.isSet
}

func (v *NullableMeetingWithIncludes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeetingWithIncludes(val *MeetingWithIncludes) *NullableMeetingWithIncludes {
	return &NullableMeetingWithIncludes{value: val, isSet: true}
}

func (v NullableMeetingWithIncludes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeetingWithIncludes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


