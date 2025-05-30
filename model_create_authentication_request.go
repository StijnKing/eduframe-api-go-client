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

// checks if the CreateAuthenticationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAuthenticationRequest{}

// CreateAuthenticationRequest struct for CreateAuthenticationRequest
type CreateAuthenticationRequest struct {
	// Login identifier.
	Uid string `json:"uid"`
	// Identifier of the associated User.
	UserId int32 `json:"user_id"`
	AuthenticationProviderType AuthenticationProviderType `json:"authentication_provider_type"`
	AdditionalProperties map[string]interface{}
}

type _CreateAuthenticationRequest CreateAuthenticationRequest

// NewCreateAuthenticationRequest instantiates a new CreateAuthenticationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthenticationRequest(uid string, userId int32, authenticationProviderType AuthenticationProviderType) *CreateAuthenticationRequest {
	this := CreateAuthenticationRequest{}
	this.Uid = uid
	this.UserId = userId
	this.AuthenticationProviderType = authenticationProviderType
	return &this
}

// NewCreateAuthenticationRequestWithDefaults instantiates a new CreateAuthenticationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthenticationRequestWithDefaults() *CreateAuthenticationRequest {
	this := CreateAuthenticationRequest{}
	return &this
}

// GetUid returns the Uid field value
func (o *CreateAuthenticationRequest) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *CreateAuthenticationRequest) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *CreateAuthenticationRequest) SetUid(v string) {
	o.Uid = v
}

// GetUserId returns the UserId field value
func (o *CreateAuthenticationRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CreateAuthenticationRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CreateAuthenticationRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetAuthenticationProviderType returns the AuthenticationProviderType field value
func (o *CreateAuthenticationRequest) GetAuthenticationProviderType() AuthenticationProviderType {
	if o == nil {
		var ret AuthenticationProviderType
		return ret
	}

	return o.AuthenticationProviderType
}

// GetAuthenticationProviderTypeOk returns a tuple with the AuthenticationProviderType field value
// and a boolean to check if the value has been set.
func (o *CreateAuthenticationRequest) GetAuthenticationProviderTypeOk() (*AuthenticationProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationProviderType, true
}

// SetAuthenticationProviderType sets field value
func (o *CreateAuthenticationRequest) SetAuthenticationProviderType(v AuthenticationProviderType) {
	o.AuthenticationProviderType = v
}

func (o CreateAuthenticationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAuthenticationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uid"] = o.Uid
	toSerialize["user_id"] = o.UserId
	toSerialize["authentication_provider_type"] = o.AuthenticationProviderType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateAuthenticationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uid",
		"user_id",
		"authentication_provider_type",
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

	varCreateAuthenticationRequest := _CreateAuthenticationRequest{}

	err = json.Unmarshal(data, &varCreateAuthenticationRequest)

	if err != nil {
		return err
	}

	*o = CreateAuthenticationRequest(varCreateAuthenticationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uid")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "authentication_provider_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAuthenticationRequest struct {
	value *CreateAuthenticationRequest
	isSet bool
}

func (v NullableCreateAuthenticationRequest) Get() *CreateAuthenticationRequest {
	return v.value
}

func (v *NullableCreateAuthenticationRequest) Set(val *CreateAuthenticationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthenticationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthenticationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthenticationRequest(val *CreateAuthenticationRequest) *NullableCreateAuthenticationRequest {
	return &NullableCreateAuthenticationRequest{value: val, isSet: true}
}

func (v NullableCreateAuthenticationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthenticationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


