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

// checks if the CreateCommentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCommentRequest{}

// CreateCommentRequest struct for CreateCommentRequest
type CreateCommentRequest struct {
	// A string representing the content of a comment.
	Content string `json:"content"`
	// Identifier of the subject the comment is linked to.
	CommentableId int32 `json:"commentable_id"`
	CommentableType CommentType `json:"commentable_type"`
	AdditionalProperties map[string]interface{}
}

type _CreateCommentRequest CreateCommentRequest

// NewCreateCommentRequest instantiates a new CreateCommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCommentRequest(content string, commentableId int32, commentableType CommentType) *CreateCommentRequest {
	this := CreateCommentRequest{}
	this.Content = content
	this.CommentableId = commentableId
	this.CommentableType = commentableType
	return &this
}

// NewCreateCommentRequestWithDefaults instantiates a new CreateCommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCommentRequestWithDefaults() *CreateCommentRequest {
	this := CreateCommentRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *CreateCommentRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CreateCommentRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CreateCommentRequest) SetContent(v string) {
	o.Content = v
}

// GetCommentableId returns the CommentableId field value
func (o *CreateCommentRequest) GetCommentableId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CommentableId
}

// GetCommentableIdOk returns a tuple with the CommentableId field value
// and a boolean to check if the value has been set.
func (o *CreateCommentRequest) GetCommentableIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommentableId, true
}

// SetCommentableId sets field value
func (o *CreateCommentRequest) SetCommentableId(v int32) {
	o.CommentableId = v
}

// GetCommentableType returns the CommentableType field value
func (o *CreateCommentRequest) GetCommentableType() CommentType {
	if o == nil {
		var ret CommentType
		return ret
	}

	return o.CommentableType
}

// GetCommentableTypeOk returns a tuple with the CommentableType field value
// and a boolean to check if the value has been set.
func (o *CreateCommentRequest) GetCommentableTypeOk() (*CommentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommentableType, true
}

// SetCommentableType sets field value
func (o *CreateCommentRequest) SetCommentableType(v CommentType) {
	o.CommentableType = v
}

func (o CreateCommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCommentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	toSerialize["commentable_id"] = o.CommentableId
	toSerialize["commentable_type"] = o.CommentableType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateCommentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
		"commentable_id",
		"commentable_type",
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

	varCreateCommentRequest := _CreateCommentRequest{}

	err = json.Unmarshal(data, &varCreateCommentRequest)

	if err != nil {
		return err
	}

	*o = CreateCommentRequest(varCreateCommentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		delete(additionalProperties, "commentable_id")
		delete(additionalProperties, "commentable_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateCommentRequest struct {
	value *CreateCommentRequest
	isSet bool
}

func (v NullableCreateCommentRequest) Get() *CreateCommentRequest {
	return v.value
}

func (v *NullableCreateCommentRequest) Set(val *CreateCommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCommentRequest(val *CreateCommentRequest) *NullableCreateCommentRequest {
	return &NullableCreateCommentRequest{value: val, isSet: true}
}

func (v NullableCreateCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


