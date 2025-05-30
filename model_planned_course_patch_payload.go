/*
Eduframe - API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eduframe

import (
	"encoding/json"
)

// checks if the PlannedCoursePatchPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlannedCoursePatchPayload{}

// PlannedCoursePatchPayload struct for PlannedCoursePatchPayload
type PlannedCoursePatchPayload struct {
	// Boolean if is published on the website.
	IsPublished *bool `json:"is_published,omitempty"`
	// Unique identifier of the course.
	CourseId *int32 `json:"course_id,omitempty"`
	// Date at which the planned course starts. Only needed for fixed planned courses.
	StartDate *string `json:"start_date,omitempty"`
	// Date at which the planned course ends. Only needed for fixed planned courses.
	EndDate NullableString `json:"end_date,omitempty"`
	// A number representing the minimum number of participants that can enroll for the planned course.
	MinParticipants NullableInt32 `json:"min_participants,omitempty"`
	// A number representing the maximum number of participants that can enroll for the planned course.
	MaxParticipants NullableInt32 `json:"max_participants,omitempty"`
	// The cost schema that the payment will follow for the specified course.
	CostScheme *string `json:"cost_scheme,omitempty"`
	// A positive float representing the price of the planned course.
	Cost NullableFloat32 `json:"cost,omitempty"`
	// Unique identifier of the course variant.
	CourseVariantId NullableInt32 `json:"course_variant_id,omitempty"`
	// Unique identifier of the course location.
	CourseLocationId NullableInt32 `json:"course_location_id,omitempty"`
	// The period of time of the planned course. Only needed for flexible planned courses.
	Duration *float32 `json:"duration,omitempty"`
	// The ids of the teachers in the course
	TeacherIds []string `json:"teacher_ids,omitempty"`
	// The custom properties of the planned course.
	Custom map[string]interface{} `json:"custom,omitempty"`
	// Custom associations are a way to link custom records to a program. ![Custom Objects](https://img.shields.io/badge/Feature-Custom_objects-1d8127) 
	CustomAssociations []CustomAssociationsRecordsPayload `json:"custom_associations,omitempty"`
}

// NewPlannedCoursePatchPayload instantiates a new PlannedCoursePatchPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlannedCoursePatchPayload() *PlannedCoursePatchPayload {
	this := PlannedCoursePatchPayload{}
	return &this
}

// NewPlannedCoursePatchPayloadWithDefaults instantiates a new PlannedCoursePatchPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlannedCoursePatchPayloadWithDefaults() *PlannedCoursePatchPayload {
	this := PlannedCoursePatchPayload{}
	return &this
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *PlannedCoursePatchPayload) GetIsPublished() bool {
	if o == nil || IsNil(o.IsPublished) {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannedCoursePatchPayload) GetIsPublishedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublished) {
		return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasIsPublished() bool {
	if o != nil && !IsNil(o.IsPublished) {
		return true
	}

	return false
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *PlannedCoursePatchPayload) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetCourseId returns the CourseId field value if set, zero value otherwise.
func (o *PlannedCoursePatchPayload) GetCourseId() int32 {
	if o == nil || IsNil(o.CourseId) {
		var ret int32
		return ret
	}
	return *o.CourseId
}

// GetCourseIdOk returns a tuple with the CourseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannedCoursePatchPayload) GetCourseIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CourseId) {
		return nil, false
	}
	return o.CourseId, true
}

// HasCourseId returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasCourseId() bool {
	if o != nil && !IsNil(o.CourseId) {
		return true
	}

	return false
}

// SetCourseId gets a reference to the given int32 and assigns it to the CourseId field.
func (o *PlannedCoursePatchPayload) SetCourseId(v int32) {
	o.CourseId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *PlannedCoursePatchPayload) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannedCoursePatchPayload) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *PlannedCoursePatchPayload) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannedCoursePatchPayload) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCoursePatchPayload) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *PlannedCoursePatchPayload) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *PlannedCoursePatchPayload) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *PlannedCoursePatchPayload) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetMinParticipants returns the MinParticipants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannedCoursePatchPayload) GetMinParticipants() int32 {
	if o == nil || IsNil(o.MinParticipants.Get()) {
		var ret int32
		return ret
	}
	return *o.MinParticipants.Get()
}

// GetMinParticipantsOk returns a tuple with the MinParticipants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCoursePatchPayload) GetMinParticipantsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinParticipants.Get(), o.MinParticipants.IsSet()
}

// HasMinParticipants returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasMinParticipants() bool {
	if o != nil && o.MinParticipants.IsSet() {
		return true
	}

	return false
}

// SetMinParticipants gets a reference to the given NullableInt32 and assigns it to the MinParticipants field.
func (o *PlannedCoursePatchPayload) SetMinParticipants(v int32) {
	o.MinParticipants.Set(&v)
}
// SetMinParticipantsNil sets the value for MinParticipants to be an explicit nil
func (o *PlannedCoursePatchPayload) SetMinParticipantsNil() {
	o.MinParticipants.Set(nil)
}

// UnsetMinParticipants ensures that no value is present for MinParticipants, not even an explicit nil
func (o *PlannedCoursePatchPayload) UnsetMinParticipants() {
	o.MinParticipants.Unset()
}

// GetMaxParticipants returns the MaxParticipants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannedCoursePatchPayload) GetMaxParticipants() int32 {
	if o == nil || IsNil(o.MaxParticipants.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxParticipants.Get()
}

// GetMaxParticipantsOk returns a tuple with the MaxParticipants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCoursePatchPayload) GetMaxParticipantsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxParticipants.Get(), o.MaxParticipants.IsSet()
}

// HasMaxParticipants returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasMaxParticipants() bool {
	if o != nil && o.MaxParticipants.IsSet() {
		return true
	}

	return false
}

// SetMaxParticipants gets a reference to the given NullableInt32 and assigns it to the MaxParticipants field.
func (o *PlannedCoursePatchPayload) SetMaxParticipants(v int32) {
	o.MaxParticipants.Set(&v)
}
// SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil
func (o *PlannedCoursePatchPayload) SetMaxParticipantsNil() {
	o.MaxParticipants.Set(nil)
}

// UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
func (o *PlannedCoursePatchPayload) UnsetMaxParticipants() {
	o.MaxParticipants.Unset()
}

// GetCostScheme returns the CostScheme field value if set, zero value otherwise.
func (o *PlannedCoursePatchPayload) GetCostScheme() string {
	if o == nil || IsNil(o.CostScheme) {
		var ret string
		return ret
	}
	return *o.CostScheme
}

// GetCostSchemeOk returns a tuple with the CostScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannedCoursePatchPayload) GetCostSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.CostScheme) {
		return nil, false
	}
	return o.CostScheme, true
}

// HasCostScheme returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasCostScheme() bool {
	if o != nil && !IsNil(o.CostScheme) {
		return true
	}

	return false
}

// SetCostScheme gets a reference to the given string and assigns it to the CostScheme field.
func (o *PlannedCoursePatchPayload) SetCostScheme(v string) {
	o.CostScheme = &v
}

// GetCost returns the Cost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannedCoursePatchPayload) GetCost() float32 {
	if o == nil || IsNil(o.Cost.Get()) {
		var ret float32
		return ret
	}
	return *o.Cost.Get()
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCoursePatchPayload) GetCostOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cost.Get(), o.Cost.IsSet()
}

// HasCost returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasCost() bool {
	if o != nil && o.Cost.IsSet() {
		return true
	}

	return false
}

// SetCost gets a reference to the given NullableFloat32 and assigns it to the Cost field.
func (o *PlannedCoursePatchPayload) SetCost(v float32) {
	o.Cost.Set(&v)
}
// SetCostNil sets the value for Cost to be an explicit nil
func (o *PlannedCoursePatchPayload) SetCostNil() {
	o.Cost.Set(nil)
}

// UnsetCost ensures that no value is present for Cost, not even an explicit nil
func (o *PlannedCoursePatchPayload) UnsetCost() {
	o.Cost.Unset()
}

// GetCourseVariantId returns the CourseVariantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannedCoursePatchPayload) GetCourseVariantId() int32 {
	if o == nil || IsNil(o.CourseVariantId.Get()) {
		var ret int32
		return ret
	}
	return *o.CourseVariantId.Get()
}

// GetCourseVariantIdOk returns a tuple with the CourseVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCoursePatchPayload) GetCourseVariantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CourseVariantId.Get(), o.CourseVariantId.IsSet()
}

// HasCourseVariantId returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasCourseVariantId() bool {
	if o != nil && o.CourseVariantId.IsSet() {
		return true
	}

	return false
}

// SetCourseVariantId gets a reference to the given NullableInt32 and assigns it to the CourseVariantId field.
func (o *PlannedCoursePatchPayload) SetCourseVariantId(v int32) {
	o.CourseVariantId.Set(&v)
}
// SetCourseVariantIdNil sets the value for CourseVariantId to be an explicit nil
func (o *PlannedCoursePatchPayload) SetCourseVariantIdNil() {
	o.CourseVariantId.Set(nil)
}

// UnsetCourseVariantId ensures that no value is present for CourseVariantId, not even an explicit nil
func (o *PlannedCoursePatchPayload) UnsetCourseVariantId() {
	o.CourseVariantId.Unset()
}

// GetCourseLocationId returns the CourseLocationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannedCoursePatchPayload) GetCourseLocationId() int32 {
	if o == nil || IsNil(o.CourseLocationId.Get()) {
		var ret int32
		return ret
	}
	return *o.CourseLocationId.Get()
}

// GetCourseLocationIdOk returns a tuple with the CourseLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCoursePatchPayload) GetCourseLocationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CourseLocationId.Get(), o.CourseLocationId.IsSet()
}

// HasCourseLocationId returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasCourseLocationId() bool {
	if o != nil && o.CourseLocationId.IsSet() {
		return true
	}

	return false
}

// SetCourseLocationId gets a reference to the given NullableInt32 and assigns it to the CourseLocationId field.
func (o *PlannedCoursePatchPayload) SetCourseLocationId(v int32) {
	o.CourseLocationId.Set(&v)
}
// SetCourseLocationIdNil sets the value for CourseLocationId to be an explicit nil
func (o *PlannedCoursePatchPayload) SetCourseLocationIdNil() {
	o.CourseLocationId.Set(nil)
}

// UnsetCourseLocationId ensures that no value is present for CourseLocationId, not even an explicit nil
func (o *PlannedCoursePatchPayload) UnsetCourseLocationId() {
	o.CourseLocationId.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *PlannedCoursePatchPayload) GetDuration() float32 {
	if o == nil || IsNil(o.Duration) {
		var ret float32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannedCoursePatchPayload) GetDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float32 and assigns it to the Duration field.
func (o *PlannedCoursePatchPayload) SetDuration(v float32) {
	o.Duration = &v
}

// GetTeacherIds returns the TeacherIds field value if set, zero value otherwise.
func (o *PlannedCoursePatchPayload) GetTeacherIds() []string {
	if o == nil || IsNil(o.TeacherIds) {
		var ret []string
		return ret
	}
	return o.TeacherIds
}

// GetTeacherIdsOk returns a tuple with the TeacherIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannedCoursePatchPayload) GetTeacherIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TeacherIds) {
		return nil, false
	}
	return o.TeacherIds, true
}

// HasTeacherIds returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasTeacherIds() bool {
	if o != nil && !IsNil(o.TeacherIds) {
		return true
	}

	return false
}

// SetTeacherIds gets a reference to the given []string and assigns it to the TeacherIds field.
func (o *PlannedCoursePatchPayload) SetTeacherIds(v []string) {
	o.TeacherIds = v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *PlannedCoursePatchPayload) GetCustom() map[string]interface{} {
	if o == nil || IsNil(o.Custom) {
		var ret map[string]interface{}
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannedCoursePatchPayload) GetCustomOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Custom) {
		return map[string]interface{}{}, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given map[string]interface{} and assigns it to the Custom field.
func (o *PlannedCoursePatchPayload) SetCustom(v map[string]interface{}) {
	o.Custom = v
}

// GetCustomAssociations returns the CustomAssociations field value if set, zero value otherwise.
func (o *PlannedCoursePatchPayload) GetCustomAssociations() []CustomAssociationsRecordsPayload {
	if o == nil || IsNil(o.CustomAssociations) {
		var ret []CustomAssociationsRecordsPayload
		return ret
	}
	return o.CustomAssociations
}

// GetCustomAssociationsOk returns a tuple with the CustomAssociations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannedCoursePatchPayload) GetCustomAssociationsOk() ([]CustomAssociationsRecordsPayload, bool) {
	if o == nil || IsNil(o.CustomAssociations) {
		return nil, false
	}
	return o.CustomAssociations, true
}

// HasCustomAssociations returns a boolean if a field has been set.
func (o *PlannedCoursePatchPayload) HasCustomAssociations() bool {
	if o != nil && !IsNil(o.CustomAssociations) {
		return true
	}

	return false
}

// SetCustomAssociations gets a reference to the given []CustomAssociationsRecordsPayload and assigns it to the CustomAssociations field.
func (o *PlannedCoursePatchPayload) SetCustomAssociations(v []CustomAssociationsRecordsPayload) {
	o.CustomAssociations = v
}

func (o PlannedCoursePatchPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlannedCoursePatchPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsPublished) {
		toSerialize["is_published"] = o.IsPublished
	}
	if !IsNil(o.CourseId) {
		toSerialize["course_id"] = o.CourseId
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.MinParticipants.IsSet() {
		toSerialize["min_participants"] = o.MinParticipants.Get()
	}
	if o.MaxParticipants.IsSet() {
		toSerialize["max_participants"] = o.MaxParticipants.Get()
	}
	if !IsNil(o.CostScheme) {
		toSerialize["cost_scheme"] = o.CostScheme
	}
	if o.Cost.IsSet() {
		toSerialize["cost"] = o.Cost.Get()
	}
	if o.CourseVariantId.IsSet() {
		toSerialize["course_variant_id"] = o.CourseVariantId.Get()
	}
	if o.CourseLocationId.IsSet() {
		toSerialize["course_location_id"] = o.CourseLocationId.Get()
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.TeacherIds) {
		toSerialize["teacher_ids"] = o.TeacherIds
	}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.CustomAssociations) {
		toSerialize["custom_associations"] = o.CustomAssociations
	}
	return toSerialize, nil
}

type NullablePlannedCoursePatchPayload struct {
	value *PlannedCoursePatchPayload
	isSet bool
}

func (v NullablePlannedCoursePatchPayload) Get() *PlannedCoursePatchPayload {
	return v.value
}

func (v *NullablePlannedCoursePatchPayload) Set(val *PlannedCoursePatchPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePlannedCoursePatchPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePlannedCoursePatchPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlannedCoursePatchPayload(val *PlannedCoursePatchPayload) *NullablePlannedCoursePatchPayload {
	return &NullablePlannedCoursePatchPayload{value: val, isSet: true}
}

func (v NullablePlannedCoursePatchPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlannedCoursePatchPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


