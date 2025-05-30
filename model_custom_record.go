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

// checks if the CustomRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomRecord{}

// CustomRecord struct for CustomRecord
type CustomRecord struct {
	// Unique identifier of the custom record.
	Id int32 `json:"id"`
	// Whether the custom record is active.
	Active bool `json:"active"`
	// The display name of the custom record.
	DisplayName string `json:"display_name"`
	// The JSON properties of the custom record.
	Properties map[string]CustomFieldValue `json:"properties"`
}

type _CustomRecord CustomRecord

// NewCustomRecord instantiates a new CustomRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomRecord(id int32, active bool, displayName string, properties map[string]CustomFieldValue) *CustomRecord {
	this := CustomRecord{}
	this.Id = id
	this.Active = active
	this.DisplayName = displayName
	this.Properties = properties
	return &this
}

// NewCustomRecordWithDefaults instantiates a new CustomRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomRecordWithDefaults() *CustomRecord {
	this := CustomRecord{}
	return &this
}

// GetId returns the Id field value
func (o *CustomRecord) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomRecord) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomRecord) SetId(v int32) {
	o.Id = v
}

// GetActive returns the Active field value
func (o *CustomRecord) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *CustomRecord) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *CustomRecord) SetActive(v bool) {
	o.Active = v
}

// GetDisplayName returns the DisplayName field value
func (o *CustomRecord) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CustomRecord) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CustomRecord) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetProperties returns the Properties field value
func (o *CustomRecord) GetProperties() map[string]CustomFieldValue {
	if o == nil {
		var ret map[string]CustomFieldValue
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *CustomRecord) GetPropertiesOk() (map[string]CustomFieldValue, bool) {
	if o == nil {
		return map[string]CustomFieldValue{}, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *CustomRecord) SetProperties(v map[string]CustomFieldValue) {
	o.Properties = v
}

func (o CustomRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["active"] = o.Active
	toSerialize["display_name"] = o.DisplayName
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

func (o *CustomRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"active",
		"display_name",
		"properties",
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

	varCustomRecord := _CustomRecord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomRecord)

	if err != nil {
		return err
	}

	*o = CustomRecord(varCustomRecord)

	return err
}

type NullableCustomRecord struct {
	value *CustomRecord
	isSet bool
}

func (v NullableCustomRecord) Get() *CustomRecord {
	return v.value
}

func (v *NullableCustomRecord) Set(val *CustomRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomRecord(val *CustomRecord) *NullableCustomRecord {
	return &NullableCustomRecord{value: val, isSet: true}
}

func (v NullableCustomRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


