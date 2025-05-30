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
	"bytes"
	"fmt"
)

// checks if the PlannedCourse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlannedCourse{}

// PlannedCourse struct for PlannedCourse
type PlannedCourse struct {
	// Unique identifier of the user.
	Id int32 `json:"id"`
	Status PlannedCourseStatus `json:"status"`
	// The period of time of the planned course. For flexible planned courses this equals the provided +duration+. For fixed planned courses this equals the number of meetings if any, if the fixed planned course has no meetings, it's the number of days between the start and end date. 
	DurationInDays NullableInt32 `json:"duration_in_days"`
	AvailabilityState AvailabilityState `json:"availability_state"`
	// Boolean wether there are payments involved for this course. Basically its true if the cost_scheme is +student+ or +order+. 
	Payable bool `json:"payable"`
	// The current amount of participants.
	CurrentParticipants int32 `json:"current_participants"`
	// The amount of confirmed active and completed enrollments.
	ConfirmedActiveAndCompletedEnrollmentsCount int32 `json:"confirmed_active_and_completed_enrollments_count"`
	// The amount of requested enrollments. ie. the number of reserved seats
	RequestedEnrollmentsCount int32 `json:"requested_enrollments_count"`
	// Boolean if there are any places available.
	AvailablePlaces bool `json:"available_places"`
	// URL to the course in canvas. Is only returned if available.
	CanvasLink NullableString `json:"canvas_link,omitempty"`
	Currency Currency `json:"currency"`
	// A positive float representing the multiplier for the VAT. Say you have 21% VAT, this will return +1.21+. 
	CostMultiplier NullableString `json:"cost_multiplier"`
	// Boolean if is published on the website.
	IsPublished bool `json:"is_published"`
	// Unique identifier of the course.
	CourseId int32 `json:"course_id"`
	// The type of the course.
	Type string `json:"type"`
	// Date at which the planned course starts. Only needed for fixed planned courses.
	StartDate string `json:"start_date"`
	// Date at which the planned course ends. Only needed for fixed planned courses.
	EndDate NullableString `json:"end_date"`
	// A number representing the minimum number of participants that can enroll for the planned course.
	MinParticipants NullableInt32 `json:"min_participants"`
	// A number representing the maximum number of participants that can enroll for the planned course.
	MaxParticipants NullableInt32 `json:"max_participants"`
	CostScheme CostScheme `json:"cost_scheme"`
	// A positive float representing the price of the planned course.
	Cost NullableString `json:"cost"`
	// Unique identifier of the course variant.
	CourseVariantId NullableInt32 `json:"course_variant_id"`
	// Unique identifier of the course location.
	CourseLocationId NullableInt32 `json:"course_location_id"`
	// Timestamp of last update.
	UpdatedAt time.Time `json:"updated_at"`
	// Timestamp of creation.
	CreatedAt time.Time `json:"created_at"`
}

type _PlannedCourse PlannedCourse

// NewPlannedCourse instantiates a new PlannedCourse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlannedCourse(id int32, status PlannedCourseStatus, durationInDays NullableInt32, availabilityState AvailabilityState, payable bool, currentParticipants int32, confirmedActiveAndCompletedEnrollmentsCount int32, requestedEnrollmentsCount int32, availablePlaces bool, currency Currency, costMultiplier NullableString, isPublished bool, courseId int32, type_ string, startDate string, endDate NullableString, minParticipants NullableInt32, maxParticipants NullableInt32, costScheme CostScheme, cost NullableString, courseVariantId NullableInt32, courseLocationId NullableInt32, updatedAt time.Time, createdAt time.Time) *PlannedCourse {
	this := PlannedCourse{}
	this.Id = id
	this.Status = status
	this.DurationInDays = durationInDays
	this.AvailabilityState = availabilityState
	this.Payable = payable
	this.CurrentParticipants = currentParticipants
	this.ConfirmedActiveAndCompletedEnrollmentsCount = confirmedActiveAndCompletedEnrollmentsCount
	this.RequestedEnrollmentsCount = requestedEnrollmentsCount
	this.AvailablePlaces = availablePlaces
	this.Currency = currency
	this.CostMultiplier = costMultiplier
	this.IsPublished = isPublished
	this.CourseId = courseId
	this.Type = type_
	this.StartDate = startDate
	this.EndDate = endDate
	this.MinParticipants = minParticipants
	this.MaxParticipants = maxParticipants
	this.CostScheme = costScheme
	this.Cost = cost
	this.CourseVariantId = courseVariantId
	this.CourseLocationId = courseLocationId
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	return &this
}

// NewPlannedCourseWithDefaults instantiates a new PlannedCourse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlannedCourseWithDefaults() *PlannedCourse {
	this := PlannedCourse{}
	return &this
}

// GetId returns the Id field value
func (o *PlannedCourse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PlannedCourse) SetId(v int32) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *PlannedCourse) GetStatus() PlannedCourseStatus {
	if o == nil {
		var ret PlannedCourseStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetStatusOk() (*PlannedCourseStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PlannedCourse) SetStatus(v PlannedCourseStatus) {
	o.Status = v
}

// GetDurationInDays returns the DurationInDays field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PlannedCourse) GetDurationInDays() int32 {
	if o == nil || o.DurationInDays.Get() == nil {
		var ret int32
		return ret
	}

	return *o.DurationInDays.Get()
}

// GetDurationInDaysOk returns a tuple with the DurationInDays field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCourse) GetDurationInDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationInDays.Get(), o.DurationInDays.IsSet()
}

// SetDurationInDays sets field value
func (o *PlannedCourse) SetDurationInDays(v int32) {
	o.DurationInDays.Set(&v)
}

// GetAvailabilityState returns the AvailabilityState field value
func (o *PlannedCourse) GetAvailabilityState() AvailabilityState {
	if o == nil {
		var ret AvailabilityState
		return ret
	}

	return o.AvailabilityState
}

// GetAvailabilityStateOk returns a tuple with the AvailabilityState field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetAvailabilityStateOk() (*AvailabilityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityState, true
}

// SetAvailabilityState sets field value
func (o *PlannedCourse) SetAvailabilityState(v AvailabilityState) {
	o.AvailabilityState = v
}

// GetPayable returns the Payable field value
func (o *PlannedCourse) GetPayable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Payable
}

// GetPayableOk returns a tuple with the Payable field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetPayableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payable, true
}

// SetPayable sets field value
func (o *PlannedCourse) SetPayable(v bool) {
	o.Payable = v
}

// GetCurrentParticipants returns the CurrentParticipants field value
func (o *PlannedCourse) GetCurrentParticipants() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentParticipants
}

// GetCurrentParticipantsOk returns a tuple with the CurrentParticipants field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetCurrentParticipantsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentParticipants, true
}

// SetCurrentParticipants sets field value
func (o *PlannedCourse) SetCurrentParticipants(v int32) {
	o.CurrentParticipants = v
}

// GetConfirmedActiveAndCompletedEnrollmentsCount returns the ConfirmedActiveAndCompletedEnrollmentsCount field value
func (o *PlannedCourse) GetConfirmedActiveAndCompletedEnrollmentsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConfirmedActiveAndCompletedEnrollmentsCount
}

// GetConfirmedActiveAndCompletedEnrollmentsCountOk returns a tuple with the ConfirmedActiveAndCompletedEnrollmentsCount field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetConfirmedActiveAndCompletedEnrollmentsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmedActiveAndCompletedEnrollmentsCount, true
}

// SetConfirmedActiveAndCompletedEnrollmentsCount sets field value
func (o *PlannedCourse) SetConfirmedActiveAndCompletedEnrollmentsCount(v int32) {
	o.ConfirmedActiveAndCompletedEnrollmentsCount = v
}

// GetRequestedEnrollmentsCount returns the RequestedEnrollmentsCount field value
func (o *PlannedCourse) GetRequestedEnrollmentsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequestedEnrollmentsCount
}

// GetRequestedEnrollmentsCountOk returns a tuple with the RequestedEnrollmentsCount field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetRequestedEnrollmentsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedEnrollmentsCount, true
}

// SetRequestedEnrollmentsCount sets field value
func (o *PlannedCourse) SetRequestedEnrollmentsCount(v int32) {
	o.RequestedEnrollmentsCount = v
}

// GetAvailablePlaces returns the AvailablePlaces field value
func (o *PlannedCourse) GetAvailablePlaces() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AvailablePlaces
}

// GetAvailablePlacesOk returns a tuple with the AvailablePlaces field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetAvailablePlacesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailablePlaces, true
}

// SetAvailablePlaces sets field value
func (o *PlannedCourse) SetAvailablePlaces(v bool) {
	o.AvailablePlaces = v
}

// GetCanvasLink returns the CanvasLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannedCourse) GetCanvasLink() string {
	if o == nil || IsNil(o.CanvasLink.Get()) {
		var ret string
		return ret
	}
	return *o.CanvasLink.Get()
}

// GetCanvasLinkOk returns a tuple with the CanvasLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCourse) GetCanvasLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanvasLink.Get(), o.CanvasLink.IsSet()
}

// HasCanvasLink returns a boolean if a field has been set.
func (o *PlannedCourse) HasCanvasLink() bool {
	if o != nil && o.CanvasLink.IsSet() {
		return true
	}

	return false
}

// SetCanvasLink gets a reference to the given NullableString and assigns it to the CanvasLink field.
func (o *PlannedCourse) SetCanvasLink(v string) {
	o.CanvasLink.Set(&v)
}
// SetCanvasLinkNil sets the value for CanvasLink to be an explicit nil
func (o *PlannedCourse) SetCanvasLinkNil() {
	o.CanvasLink.Set(nil)
}

// UnsetCanvasLink ensures that no value is present for CanvasLink, not even an explicit nil
func (o *PlannedCourse) UnsetCanvasLink() {
	o.CanvasLink.Unset()
}

// GetCurrency returns the Currency field value
func (o *PlannedCourse) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PlannedCourse) SetCurrency(v Currency) {
	o.Currency = v
}

// GetCostMultiplier returns the CostMultiplier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PlannedCourse) GetCostMultiplier() string {
	if o == nil || o.CostMultiplier.Get() == nil {
		var ret string
		return ret
	}

	return *o.CostMultiplier.Get()
}

// GetCostMultiplierOk returns a tuple with the CostMultiplier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCourse) GetCostMultiplierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CostMultiplier.Get(), o.CostMultiplier.IsSet()
}

// SetCostMultiplier sets field value
func (o *PlannedCourse) SetCostMultiplier(v string) {
	o.CostMultiplier.Set(&v)
}

// GetIsPublished returns the IsPublished field value
func (o *PlannedCourse) GetIsPublished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetIsPublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublished, true
}

// SetIsPublished sets field value
func (o *PlannedCourse) SetIsPublished(v bool) {
	o.IsPublished = v
}

// GetCourseId returns the CourseId field value
func (o *PlannedCourse) GetCourseId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CourseId
}

// GetCourseIdOk returns a tuple with the CourseId field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetCourseIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CourseId, true
}

// SetCourseId sets field value
func (o *PlannedCourse) SetCourseId(v int32) {
	o.CourseId = v
}

// GetType returns the Type field value
func (o *PlannedCourse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PlannedCourse) SetType(v string) {
	o.Type = v
}

// GetStartDate returns the StartDate field value
func (o *PlannedCourse) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *PlannedCourse) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PlannedCourse) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCourse) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// SetEndDate sets field value
func (o *PlannedCourse) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// GetMinParticipants returns the MinParticipants field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PlannedCourse) GetMinParticipants() int32 {
	if o == nil || o.MinParticipants.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MinParticipants.Get()
}

// GetMinParticipantsOk returns a tuple with the MinParticipants field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCourse) GetMinParticipantsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinParticipants.Get(), o.MinParticipants.IsSet()
}

// SetMinParticipants sets field value
func (o *PlannedCourse) SetMinParticipants(v int32) {
	o.MinParticipants.Set(&v)
}

// GetMaxParticipants returns the MaxParticipants field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PlannedCourse) GetMaxParticipants() int32 {
	if o == nil || o.MaxParticipants.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MaxParticipants.Get()
}

// GetMaxParticipantsOk returns a tuple with the MaxParticipants field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCourse) GetMaxParticipantsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxParticipants.Get(), o.MaxParticipants.IsSet()
}

// SetMaxParticipants sets field value
func (o *PlannedCourse) SetMaxParticipants(v int32) {
	o.MaxParticipants.Set(&v)
}

// GetCostScheme returns the CostScheme field value
func (o *PlannedCourse) GetCostScheme() CostScheme {
	if o == nil {
		var ret CostScheme
		return ret
	}

	return o.CostScheme
}

// GetCostSchemeOk returns a tuple with the CostScheme field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetCostSchemeOk() (*CostScheme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostScheme, true
}

// SetCostScheme sets field value
func (o *PlannedCourse) SetCostScheme(v CostScheme) {
	o.CostScheme = v
}

// GetCost returns the Cost field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PlannedCourse) GetCost() string {
	if o == nil || o.Cost.Get() == nil {
		var ret string
		return ret
	}

	return *o.Cost.Get()
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCourse) GetCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cost.Get(), o.Cost.IsSet()
}

// SetCost sets field value
func (o *PlannedCourse) SetCost(v string) {
	o.Cost.Set(&v)
}

// GetCourseVariantId returns the CourseVariantId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PlannedCourse) GetCourseVariantId() int32 {
	if o == nil || o.CourseVariantId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CourseVariantId.Get()
}

// GetCourseVariantIdOk returns a tuple with the CourseVariantId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCourse) GetCourseVariantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CourseVariantId.Get(), o.CourseVariantId.IsSet()
}

// SetCourseVariantId sets field value
func (o *PlannedCourse) SetCourseVariantId(v int32) {
	o.CourseVariantId.Set(&v)
}

// GetCourseLocationId returns the CourseLocationId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PlannedCourse) GetCourseLocationId() int32 {
	if o == nil || o.CourseLocationId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CourseLocationId.Get()
}

// GetCourseLocationIdOk returns a tuple with the CourseLocationId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedCourse) GetCourseLocationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CourseLocationId.Get(), o.CourseLocationId.IsSet()
}

// SetCourseLocationId sets field value
func (o *PlannedCourse) SetCourseLocationId(v int32) {
	o.CourseLocationId.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PlannedCourse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PlannedCourse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PlannedCourse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PlannedCourse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PlannedCourse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o PlannedCourse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlannedCourse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["duration_in_days"] = o.DurationInDays.Get()
	toSerialize["availability_state"] = o.AvailabilityState
	toSerialize["payable"] = o.Payable
	toSerialize["current_participants"] = o.CurrentParticipants
	toSerialize["confirmed_active_and_completed_enrollments_count"] = o.ConfirmedActiveAndCompletedEnrollmentsCount
	toSerialize["requested_enrollments_count"] = o.RequestedEnrollmentsCount
	toSerialize["available_places"] = o.AvailablePlaces
	if o.CanvasLink.IsSet() {
		toSerialize["canvas_link"] = o.CanvasLink.Get()
	}
	toSerialize["currency"] = o.Currency
	toSerialize["cost_multiplier"] = o.CostMultiplier.Get()
	toSerialize["is_published"] = o.IsPublished
	toSerialize["course_id"] = o.CourseId
	toSerialize["type"] = o.Type
	toSerialize["start_date"] = o.StartDate
	toSerialize["end_date"] = o.EndDate.Get()
	toSerialize["min_participants"] = o.MinParticipants.Get()
	toSerialize["max_participants"] = o.MaxParticipants.Get()
	toSerialize["cost_scheme"] = o.CostScheme
	toSerialize["cost"] = o.Cost.Get()
	toSerialize["course_variant_id"] = o.CourseVariantId.Get()
	toSerialize["course_location_id"] = o.CourseLocationId.Get()
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *PlannedCourse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"duration_in_days",
		"availability_state",
		"payable",
		"current_participants",
		"confirmed_active_and_completed_enrollments_count",
		"requested_enrollments_count",
		"available_places",
		"currency",
		"cost_multiplier",
		"is_published",
		"course_id",
		"type",
		"start_date",
		"end_date",
		"min_participants",
		"max_participants",
		"cost_scheme",
		"cost",
		"course_variant_id",
		"course_location_id",
		"updated_at",
		"created_at",
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

	varPlannedCourse := _PlannedCourse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPlannedCourse)

	if err != nil {
		return err
	}

	*o = PlannedCourse(varPlannedCourse)

	return err
}

type NullablePlannedCourse struct {
	value *PlannedCourse
	isSet bool
}

func (v NullablePlannedCourse) Get() *PlannedCourse {
	return v.value
}

func (v *NullablePlannedCourse) Set(val *PlannedCourse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlannedCourse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlannedCourse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlannedCourse(val *PlannedCourse) *NullablePlannedCourse {
	return &NullablePlannedCourse{value: val, isSet: true}
}

func (v NullablePlannedCourse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlannedCourse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


