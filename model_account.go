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

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account struct for Account
type Account struct {
	// Unique identifier of the account.
	Id int32 `json:"id"`
	// Arbitrary string representing the name of the account. Is autogenerated for personal accounts.
	Name string `json:"name"`
	// A string representing the billing e-mail of the account
	Email NullableString `json:"email"`
	// A string representing the phone number of the account
	Phone NullableString `json:"phone"`
	AccountType AccountType `json:"account_type"`
	// An array of integers representing unique identifier values associated with labels. 
	LabelIds []int32 `json:"label_ids"`
	// Timestamp of last update.
	UpdatedAt time.Time `json:"updated_at"`
	// Timestamp of creation.
	CreatedAt time.Time `json:"created_at"`
}

type _Account Account

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(id int32, name string, email NullableString, phone NullableString, accountType AccountType, labelIds []int32, updatedAt time.Time, createdAt time.Time) *Account {
	this := Account{}
	this.Id = id
	this.Name = name
	this.Email = email
	this.Phone = phone
	this.AccountType = accountType
	this.LabelIds = labelIds
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetId returns the Id field value
func (o *Account) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Account) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Account) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Account) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Account) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Account) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Account) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *Account) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetPhone returns the Phone field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Account) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}

	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// SetPhone sets field value
func (o *Account) SetPhone(v string) {
	o.Phone.Set(&v)
}

// GetAccountType returns the AccountType field value
func (o *Account) GetAccountType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *Account) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *Account) SetAccountType(v AccountType) {
	o.AccountType = v
}

// GetLabelIds returns the LabelIds field value
func (o *Account) GetLabelIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.LabelIds
}

// GetLabelIdsOk returns a tuple with the LabelIds field value
// and a boolean to check if the value has been set.
func (o *Account) GetLabelIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LabelIds, true
}

// SetLabelIds sets field value
func (o *Account) SetLabelIds(v []int32) {
	o.LabelIds = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Account) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Account) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Account) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Account) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Account) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Account) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email.Get()
	toSerialize["phone"] = o.Phone.Get()
	toSerialize["account_type"] = o.AccountType
	toSerialize["label_ids"] = o.LabelIds
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *Account) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"email",
		"phone",
		"account_type",
		"label_ids",
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

	varAccount := _Account{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccount)

	if err != nil {
		return err
	}

	*o = Account(varAccount)

	return err
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


