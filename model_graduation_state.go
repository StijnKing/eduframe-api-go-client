/*
Eduframe - API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eduframe

import (
	"encoding/json"
	"fmt"
)

// GraduationState The graduation state of the enrollment.
type GraduationState string

// List of GraduationState
const (
	GRADUATIONSTATE_AWAITING_JUDGEMENT GraduationState = "awaiting_judgement"
	GRADUATIONSTATE_FAILED GraduationState = "failed"
	GRADUATIONSTATE_PASSED GraduationState = "passed"
)

// All allowed values of GraduationState enum
var AllowedGraduationStateEnumValues = []GraduationState{
	"awaiting_judgement",
	"failed",
	"passed",
}

func (v *GraduationState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GraduationState(value)
	for _, existing := range AllowedGraduationStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GraduationState", value)
}

// NewGraduationStateFromValue returns a pointer to a valid GraduationState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGraduationStateFromValue(v string) (*GraduationState, error) {
	ev := GraduationState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GraduationState: valid values are %v", v, AllowedGraduationStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GraduationState) IsValid() bool {
	for _, existing := range AllowedGraduationStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GraduationState value
func (v GraduationState) Ptr() *GraduationState {
	return &v
}

type NullableGraduationState struct {
	value *GraduationState
	isSet bool
}

func (v NullableGraduationState) Get() *GraduationState {
	return v.value
}

func (v *NullableGraduationState) Set(val *GraduationState) {
	v.value = val
	v.isSet = true
}

func (v NullableGraduationState) IsSet() bool {
	return v.isSet
}

func (v *NullableGraduationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraduationState(val *GraduationState) *NullableGraduationState {
	return &NullableGraduationState{value: val, isSet: true}
}

func (v NullableGraduationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraduationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

