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

// checks if the Program type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Program{}

// Program struct for Program
type Program struct {
	// Unique identifier of the program.
	Id int32 `json:"id"`
	// URL to the signup page for this program.
	SignupUrl string `json:"signup_url"`
	// Name of the program.
	Name string `json:"name"`
	// The price to be paid for this program.
	Cost NullableString `json:"cost"`
	CostScheme CostScheme `json:"cost_scheme"`
	// Boolean representing the publishable status of the program.
	IsPublished bool `json:"is_published"`
	// Identifier of the category of the program.
	CategoryId int32 `json:"category_id"`
	// Conditions for this program.
	Conditions NullableString `json:"conditions"`
	// Human readable identifier, unique per educator.
	Slug string `json:"slug"`
}

type _Program Program

// NewProgram instantiates a new Program object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgram(id int32, signupUrl string, name string, cost NullableString, costScheme CostScheme, isPublished bool, categoryId int32, conditions NullableString, slug string) *Program {
	this := Program{}
	this.Id = id
	this.SignupUrl = signupUrl
	this.Name = name
	this.Cost = cost
	this.CostScheme = costScheme
	this.IsPublished = isPublished
	this.CategoryId = categoryId
	this.Conditions = conditions
	this.Slug = slug
	return &this
}

// NewProgramWithDefaults instantiates a new Program object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramWithDefaults() *Program {
	this := Program{}
	return &this
}

// GetId returns the Id field value
func (o *Program) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Program) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Program) SetId(v int32) {
	o.Id = v
}

// GetSignupUrl returns the SignupUrl field value
func (o *Program) GetSignupUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignupUrl
}

// GetSignupUrlOk returns a tuple with the SignupUrl field value
// and a boolean to check if the value has been set.
func (o *Program) GetSignupUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignupUrl, true
}

// SetSignupUrl sets field value
func (o *Program) SetSignupUrl(v string) {
	o.SignupUrl = v
}

// GetName returns the Name field value
func (o *Program) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Program) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Program) SetName(v string) {
	o.Name = v
}

// GetCost returns the Cost field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Program) GetCost() string {
	if o == nil || o.Cost.Get() == nil {
		var ret string
		return ret
	}

	return *o.Cost.Get()
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Program) GetCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cost.Get(), o.Cost.IsSet()
}

// SetCost sets field value
func (o *Program) SetCost(v string) {
	o.Cost.Set(&v)
}

// GetCostScheme returns the CostScheme field value
func (o *Program) GetCostScheme() CostScheme {
	if o == nil {
		var ret CostScheme
		return ret
	}

	return o.CostScheme
}

// GetCostSchemeOk returns a tuple with the CostScheme field value
// and a boolean to check if the value has been set.
func (o *Program) GetCostSchemeOk() (*CostScheme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostScheme, true
}

// SetCostScheme sets field value
func (o *Program) SetCostScheme(v CostScheme) {
	o.CostScheme = v
}

// GetIsPublished returns the IsPublished field value
func (o *Program) GetIsPublished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value
// and a boolean to check if the value has been set.
func (o *Program) GetIsPublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublished, true
}

// SetIsPublished sets field value
func (o *Program) SetIsPublished(v bool) {
	o.IsPublished = v
}

// GetCategoryId returns the CategoryId field value
func (o *Program) GetCategoryId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *Program) GetCategoryIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *Program) SetCategoryId(v int32) {
	o.CategoryId = v
}

// GetConditions returns the Conditions field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Program) GetConditions() string {
	if o == nil || o.Conditions.Get() == nil {
		var ret string
		return ret
	}

	return *o.Conditions.Get()
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Program) GetConditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions.Get(), o.Conditions.IsSet()
}

// SetConditions sets field value
func (o *Program) SetConditions(v string) {
	o.Conditions.Set(&v)
}

// GetSlug returns the Slug field value
func (o *Program) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Program) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Program) SetSlug(v string) {
	o.Slug = v
}

func (o Program) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Program) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["signup_url"] = o.SignupUrl
	toSerialize["name"] = o.Name
	toSerialize["cost"] = o.Cost.Get()
	toSerialize["cost_scheme"] = o.CostScheme
	toSerialize["is_published"] = o.IsPublished
	toSerialize["category_id"] = o.CategoryId
	toSerialize["conditions"] = o.Conditions.Get()
	toSerialize["slug"] = o.Slug
	return toSerialize, nil
}

func (o *Program) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"signup_url",
		"name",
		"cost",
		"cost_scheme",
		"is_published",
		"category_id",
		"conditions",
		"slug",
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

	varProgram := _Program{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProgram)

	if err != nil {
		return err
	}

	*o = Program(varProgram)

	return err
}

type NullableProgram struct {
	value *Program
	isSet bool
}

func (v NullableProgram) Get() *Program {
	return v.value
}

func (v *NullableProgram) Set(val *Program) {
	v.value = val
	v.isSet = true
}

func (v NullableProgram) IsSet() bool {
	return v.isSet
}

func (v *NullableProgram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgram(val *Program) *NullableProgram {
	return &NullableProgram{value: val, isSet: true}
}

func (v NullableProgram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


