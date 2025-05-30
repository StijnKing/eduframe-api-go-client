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

// checks if the TeacherEnrollment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeacherEnrollment{}

// TeacherEnrollment struct for TeacherEnrollment
type TeacherEnrollment struct {
	// Unique identifier of the teacher enrollment.
	Id int32 `json:"id"`
	// Unique identifier of the planned course.
	PlannedCourseId int32 `json:"planned_course_id"`
	// Unique identifier of the teacher.
	TeacherId int32 `json:"teacher_id"`
	// Unique identifier of the teacher role.  ![Teacher Roles](https://img.shields.io/badge/Feature-Teacher_Roles-blue) 
	TeacherRoleId *int32 `json:"teacher_role_id,omitempty"`
	// Timestamp of last update.
	UpdatedAt time.Time `json:"updated_at"`
	// Timestamp of creation.
	CreatedAt time.Time `json:"created_at"`
}

type _TeacherEnrollment TeacherEnrollment

// NewTeacherEnrollment instantiates a new TeacherEnrollment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeacherEnrollment(id int32, plannedCourseId int32, teacherId int32, updatedAt time.Time, createdAt time.Time) *TeacherEnrollment {
	this := TeacherEnrollment{}
	this.Id = id
	this.PlannedCourseId = plannedCourseId
	this.TeacherId = teacherId
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	return &this
}

// NewTeacherEnrollmentWithDefaults instantiates a new TeacherEnrollment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeacherEnrollmentWithDefaults() *TeacherEnrollment {
	this := TeacherEnrollment{}
	return &this
}

// GetId returns the Id field value
func (o *TeacherEnrollment) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeacherEnrollment) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeacherEnrollment) SetId(v int32) {
	o.Id = v
}

// GetPlannedCourseId returns the PlannedCourseId field value
func (o *TeacherEnrollment) GetPlannedCourseId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlannedCourseId
}

// GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field value
// and a boolean to check if the value has been set.
func (o *TeacherEnrollment) GetPlannedCourseIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlannedCourseId, true
}

// SetPlannedCourseId sets field value
func (o *TeacherEnrollment) SetPlannedCourseId(v int32) {
	o.PlannedCourseId = v
}

// GetTeacherId returns the TeacherId field value
func (o *TeacherEnrollment) GetTeacherId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TeacherId
}

// GetTeacherIdOk returns a tuple with the TeacherId field value
// and a boolean to check if the value has been set.
func (o *TeacherEnrollment) GetTeacherIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeacherId, true
}

// SetTeacherId sets field value
func (o *TeacherEnrollment) SetTeacherId(v int32) {
	o.TeacherId = v
}

// GetTeacherRoleId returns the TeacherRoleId field value if set, zero value otherwise.
func (o *TeacherEnrollment) GetTeacherRoleId() int32 {
	if o == nil || IsNil(o.TeacherRoleId) {
		var ret int32
		return ret
	}
	return *o.TeacherRoleId
}

// GetTeacherRoleIdOk returns a tuple with the TeacherRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeacherEnrollment) GetTeacherRoleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TeacherRoleId) {
		return nil, false
	}
	return o.TeacherRoleId, true
}

// HasTeacherRoleId returns a boolean if a field has been set.
func (o *TeacherEnrollment) HasTeacherRoleId() bool {
	if o != nil && !IsNil(o.TeacherRoleId) {
		return true
	}

	return false
}

// SetTeacherRoleId gets a reference to the given int32 and assigns it to the TeacherRoleId field.
func (o *TeacherEnrollment) SetTeacherRoleId(v int32) {
	o.TeacherRoleId = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TeacherEnrollment) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TeacherEnrollment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TeacherEnrollment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TeacherEnrollment) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TeacherEnrollment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TeacherEnrollment) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o TeacherEnrollment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeacherEnrollment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["planned_course_id"] = o.PlannedCourseId
	toSerialize["teacher_id"] = o.TeacherId
	if !IsNil(o.TeacherRoleId) {
		toSerialize["teacher_role_id"] = o.TeacherRoleId
	}
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *TeacherEnrollment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"planned_course_id",
		"teacher_id",
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

	varTeacherEnrollment := _TeacherEnrollment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeacherEnrollment)

	if err != nil {
		return err
	}

	*o = TeacherEnrollment(varTeacherEnrollment)

	return err
}

type NullableTeacherEnrollment struct {
	value *TeacherEnrollment
	isSet bool
}

func (v NullableTeacherEnrollment) Get() *TeacherEnrollment {
	return v.value
}

func (v *NullableTeacherEnrollment) Set(val *TeacherEnrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableTeacherEnrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableTeacherEnrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeacherEnrollment(val *TeacherEnrollment) *NullableTeacherEnrollment {
	return &NullableTeacherEnrollment{value: val, isSet: true}
}

func (v NullableTeacherEnrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeacherEnrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


