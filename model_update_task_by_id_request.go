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

// checks if the UpdateTaskByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTaskByIdRequest{}

// UpdateTaskByIdRequest struct for UpdateTaskByIdRequest
type UpdateTaskByIdRequest struct {
	// The title of the task.
	Name *string `json:"name,omitempty"`
	// A string representing the description of the task.
	Description NullableString `json:"description,omitempty"`
	// Date when the task must be completed.
	DueDate NullableString `json:"due_date,omitempty"`
	// Boolean if the task is starred.
	Starred *bool `json:"starred,omitempty"`
	// Unique identifier of the assigned user for the task.
	AssigneeId NullableInt32 `json:"assignee_id,omitempty"`
	// Type of the subject.
	SubjectType NullableString `json:"subject_type,omitempty"`
	// Identifier of the subject.
	SubjectId NullableInt32 `json:"subject_id,omitempty"`
	// Boolean representing the status of the task. The default value is false.
	Completed *bool `json:"completed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateTaskByIdRequest UpdateTaskByIdRequest

// NewUpdateTaskByIdRequest instantiates a new UpdateTaskByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTaskByIdRequest() *UpdateTaskByIdRequest {
	this := UpdateTaskByIdRequest{}
	return &this
}

// NewUpdateTaskByIdRequestWithDefaults instantiates a new UpdateTaskByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTaskByIdRequestWithDefaults() *UpdateTaskByIdRequest {
	this := UpdateTaskByIdRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateTaskByIdRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskByIdRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateTaskByIdRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateTaskByIdRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTaskByIdRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTaskByIdRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateTaskByIdRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateTaskByIdRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateTaskByIdRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateTaskByIdRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTaskByIdRequest) GetDueDate() string {
	if o == nil || IsNil(o.DueDate.Get()) {
		var ret string
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTaskByIdRequest) GetDueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *UpdateTaskByIdRequest) HasDueDate() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableString and assigns it to the DueDate field.
func (o *UpdateTaskByIdRequest) SetDueDate(v string) {
	o.DueDate.Set(&v)
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *UpdateTaskByIdRequest) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *UpdateTaskByIdRequest) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetStarred returns the Starred field value if set, zero value otherwise.
func (o *UpdateTaskByIdRequest) GetStarred() bool {
	if o == nil || IsNil(o.Starred) {
		var ret bool
		return ret
	}
	return *o.Starred
}

// GetStarredOk returns a tuple with the Starred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskByIdRequest) GetStarredOk() (*bool, bool) {
	if o == nil || IsNil(o.Starred) {
		return nil, false
	}
	return o.Starred, true
}

// HasStarred returns a boolean if a field has been set.
func (o *UpdateTaskByIdRequest) HasStarred() bool {
	if o != nil && !IsNil(o.Starred) {
		return true
	}

	return false
}

// SetStarred gets a reference to the given bool and assigns it to the Starred field.
func (o *UpdateTaskByIdRequest) SetStarred(v bool) {
	o.Starred = &v
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTaskByIdRequest) GetAssigneeId() int32 {
	if o == nil || IsNil(o.AssigneeId.Get()) {
		var ret int32
		return ret
	}
	return *o.AssigneeId.Get()
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTaskByIdRequest) GetAssigneeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssigneeId.Get(), o.AssigneeId.IsSet()
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *UpdateTaskByIdRequest) HasAssigneeId() bool {
	if o != nil && o.AssigneeId.IsSet() {
		return true
	}

	return false
}

// SetAssigneeId gets a reference to the given NullableInt32 and assigns it to the AssigneeId field.
func (o *UpdateTaskByIdRequest) SetAssigneeId(v int32) {
	o.AssigneeId.Set(&v)
}
// SetAssigneeIdNil sets the value for AssigneeId to be an explicit nil
func (o *UpdateTaskByIdRequest) SetAssigneeIdNil() {
	o.AssigneeId.Set(nil)
}

// UnsetAssigneeId ensures that no value is present for AssigneeId, not even an explicit nil
func (o *UpdateTaskByIdRequest) UnsetAssigneeId() {
	o.AssigneeId.Unset()
}

// GetSubjectType returns the SubjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTaskByIdRequest) GetSubjectType() string {
	if o == nil || IsNil(o.SubjectType.Get()) {
		var ret string
		return ret
	}
	return *o.SubjectType.Get()
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTaskByIdRequest) GetSubjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectType.Get(), o.SubjectType.IsSet()
}

// HasSubjectType returns a boolean if a field has been set.
func (o *UpdateTaskByIdRequest) HasSubjectType() bool {
	if o != nil && o.SubjectType.IsSet() {
		return true
	}

	return false
}

// SetSubjectType gets a reference to the given NullableString and assigns it to the SubjectType field.
func (o *UpdateTaskByIdRequest) SetSubjectType(v string) {
	o.SubjectType.Set(&v)
}
// SetSubjectTypeNil sets the value for SubjectType to be an explicit nil
func (o *UpdateTaskByIdRequest) SetSubjectTypeNil() {
	o.SubjectType.Set(nil)
}

// UnsetSubjectType ensures that no value is present for SubjectType, not even an explicit nil
func (o *UpdateTaskByIdRequest) UnsetSubjectType() {
	o.SubjectType.Unset()
}

// GetSubjectId returns the SubjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTaskByIdRequest) GetSubjectId() int32 {
	if o == nil || IsNil(o.SubjectId.Get()) {
		var ret int32
		return ret
	}
	return *o.SubjectId.Get()
}

// GetSubjectIdOk returns a tuple with the SubjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTaskByIdRequest) GetSubjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectId.Get(), o.SubjectId.IsSet()
}

// HasSubjectId returns a boolean if a field has been set.
func (o *UpdateTaskByIdRequest) HasSubjectId() bool {
	if o != nil && o.SubjectId.IsSet() {
		return true
	}

	return false
}

// SetSubjectId gets a reference to the given NullableInt32 and assigns it to the SubjectId field.
func (o *UpdateTaskByIdRequest) SetSubjectId(v int32) {
	o.SubjectId.Set(&v)
}
// SetSubjectIdNil sets the value for SubjectId to be an explicit nil
func (o *UpdateTaskByIdRequest) SetSubjectIdNil() {
	o.SubjectId.Set(nil)
}

// UnsetSubjectId ensures that no value is present for SubjectId, not even an explicit nil
func (o *UpdateTaskByIdRequest) UnsetSubjectId() {
	o.SubjectId.Unset()
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *UpdateTaskByIdRequest) GetCompleted() bool {
	if o == nil || IsNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskByIdRequest) GetCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *UpdateTaskByIdRequest) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *UpdateTaskByIdRequest) SetCompleted(v bool) {
	o.Completed = &v
}

func (o UpdateTaskByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTaskByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DueDate.IsSet() {
		toSerialize["due_date"] = o.DueDate.Get()
	}
	if !IsNil(o.Starred) {
		toSerialize["starred"] = o.Starred
	}
	if o.AssigneeId.IsSet() {
		toSerialize["assignee_id"] = o.AssigneeId.Get()
	}
	if o.SubjectType.IsSet() {
		toSerialize["subject_type"] = o.SubjectType.Get()
	}
	if o.SubjectId.IsSet() {
		toSerialize["subject_id"] = o.SubjectId.Get()
	}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateTaskByIdRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateTaskByIdRequest := _UpdateTaskByIdRequest{}

	err = json.Unmarshal(data, &varUpdateTaskByIdRequest)

	if err != nil {
		return err
	}

	*o = UpdateTaskByIdRequest(varUpdateTaskByIdRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "due_date")
		delete(additionalProperties, "starred")
		delete(additionalProperties, "assignee_id")
		delete(additionalProperties, "subject_type")
		delete(additionalProperties, "subject_id")
		delete(additionalProperties, "completed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateTaskByIdRequest struct {
	value *UpdateTaskByIdRequest
	isSet bool
}

func (v NullableUpdateTaskByIdRequest) Get() *UpdateTaskByIdRequest {
	return v.value
}

func (v *NullableUpdateTaskByIdRequest) Set(val *UpdateTaskByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTaskByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTaskByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTaskByIdRequest(val *UpdateTaskByIdRequest) *NullableUpdateTaskByIdRequest {
	return &NullableUpdateTaskByIdRequest{value: val, isSet: true}
}

func (v NullableUpdateTaskByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTaskByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


