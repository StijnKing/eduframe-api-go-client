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

// Locale the model 'Locale'
type Locale string

// List of Locale
const (
	LOCALE_DE Locale = "de"
	LOCALE_EN Locale = "en"
	LOCALE_EN_GB Locale = "en-GB"
	LOCALE_EN_US Locale = "en-US"
	LOCALE_ES Locale = "es"
	LOCALE_IS Locale = "is"
	LOCALE_NL Locale = "nl"
)

// All allowed values of Locale enum
var AllowedLocaleEnumValues = []Locale{
	"de",
	"en",
	"en-GB",
	"en-US",
	"es",
	"is",
	"nl",
}

func (v *Locale) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Locale(value)
	for _, existing := range AllowedLocaleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Locale", value)
}

// NewLocaleFromValue returns a pointer to a valid Locale
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocaleFromValue(v string) (*Locale, error) {
	ev := Locale(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Locale: valid values are %v", v, AllowedLocaleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Locale) IsValid() bool {
	for _, existing := range AllowedLocaleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Locale value
func (v Locale) Ptr() *Locale {
	return &v
}

type NullableLocale struct {
	value *Locale
	isSet bool
}

func (v NullableLocale) Get() *Locale {
	return v.value
}

func (v *NullableLocale) Set(val *Locale) {
	v.value = val
	v.isSet = true
}

func (v NullableLocale) IsSet() bool {
	return v.isSet
}

func (v *NullableLocale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocale(val *Locale) *NullableLocale {
	return &NullableLocale{value: val, isSet: true}
}

func (v NullableLocale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

