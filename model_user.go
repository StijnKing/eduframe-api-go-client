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

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	// Unique identifier of the user.
	Id int32 `json:"id"`
	// First name of the user.
	FirstName string `json:"first_name"`
	// Middle name of the user.
	MiddleName NullableString `json:"middle_name"`
	// Last name of the user.
	LastName string `json:"last_name"`
	// The e-mail of the user.
	Email string `json:"email"`
	// The relative url path to the avatar.
	AvatarUrl string `json:"avatar_url"`
	// List of roles the user has.
	Roles []string `json:"roles"`
	// Short note about the user.
	NotesUser NullableString `json:"notes_user"`
	// Long description of the user. (same as teacher_description)
	Description NullableString `json:"description"`
	// An array of integers representing unique identifier values associated with labels. 
	LabelIds []int32 `json:"label_ids"`
	// The employee number of this user.
	EmployeeNumber NullableString `json:"employee_number"`
	// Unique identifier of the student.
	StudentNumber NullableString `json:"student_number"`
	// Short description of the user.
	TeacherHeadline NullableString `json:"teacher_headline"`
	// Long description of the user. (same as description)
	TeacherDescription NullableString `json:"teacher_description"`
	Locale NullableLocale `json:"locale"`
	// Boolean representing the possibility of the user to receive newsletters.
	WantsNewsletter bool `json:"wants_newsletter"`
	// Timestamp of last update.
	UpdatedAt time.Time `json:"updated_at"`
	// Timestamp of creation.
	CreatedAt time.Time `json:"created_at"`
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(id int32, firstName string, middleName NullableString, lastName string, email string, avatarUrl string, roles []string, notesUser NullableString, description NullableString, labelIds []int32, employeeNumber NullableString, studentNumber NullableString, teacherHeadline NullableString, teacherDescription NullableString, locale NullableLocale, wantsNewsletter bool, updatedAt time.Time, createdAt time.Time) *User {
	this := User{}
	this.Id = id
	this.FirstName = firstName
	this.MiddleName = middleName
	this.LastName = lastName
	this.Email = email
	this.AvatarUrl = avatarUrl
	this.Roles = roles
	this.NotesUser = notesUser
	this.Description = description
	this.LabelIds = labelIds
	this.EmployeeNumber = employeeNumber
	this.StudentNumber = studentNumber
	this.TeacherHeadline = teacherHeadline
	this.TeacherDescription = teacherDescription
	this.Locale = locale
	this.WantsNewsletter = wantsNewsletter
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value
func (o *User) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *User) SetId(v int32) {
	o.Id = v
}

// GetFirstName returns the FirstName field value
func (o *User) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *User) SetFirstName(v string) {
	o.FirstName = v
}

// GetMiddleName returns the MiddleName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetMiddleName() string {
	if o == nil || o.MiddleName.Get() == nil {
		var ret string
		return ret
	}

	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetMiddleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// SetMiddleName sets field value
func (o *User) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}

// GetLastName returns the LastName field value
func (o *User) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *User) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *User) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *User) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *User) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *User) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

// GetRoles returns the Roles field value
func (o *User) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *User) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *User) SetRoles(v []string) {
	o.Roles = v
}

// GetNotesUser returns the NotesUser field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetNotesUser() string {
	if o == nil || o.NotesUser.Get() == nil {
		var ret string
		return ret
	}

	return *o.NotesUser.Get()
}

// GetNotesUserOk returns a tuple with the NotesUser field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetNotesUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotesUser.Get(), o.NotesUser.IsSet()
}

// SetNotesUser sets field value
func (o *User) SetNotesUser(v string) {
	o.NotesUser.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *User) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetLabelIds returns the LabelIds field value
func (o *User) GetLabelIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.LabelIds
}

// GetLabelIdsOk returns a tuple with the LabelIds field value
// and a boolean to check if the value has been set.
func (o *User) GetLabelIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LabelIds, true
}

// SetLabelIds sets field value
func (o *User) SetLabelIds(v []int32) {
	o.LabelIds = v
}

// GetEmployeeNumber returns the EmployeeNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetEmployeeNumber() string {
	if o == nil || o.EmployeeNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmployeeNumber.Get()
}

// GetEmployeeNumberOk returns a tuple with the EmployeeNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetEmployeeNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmployeeNumber.Get(), o.EmployeeNumber.IsSet()
}

// SetEmployeeNumber sets field value
func (o *User) SetEmployeeNumber(v string) {
	o.EmployeeNumber.Set(&v)
}

// GetStudentNumber returns the StudentNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetStudentNumber() string {
	if o == nil || o.StudentNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.StudentNumber.Get()
}

// GetStudentNumberOk returns a tuple with the StudentNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetStudentNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StudentNumber.Get(), o.StudentNumber.IsSet()
}

// SetStudentNumber sets field value
func (o *User) SetStudentNumber(v string) {
	o.StudentNumber.Set(&v)
}

// GetTeacherHeadline returns the TeacherHeadline field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetTeacherHeadline() string {
	if o == nil || o.TeacherHeadline.Get() == nil {
		var ret string
		return ret
	}

	return *o.TeacherHeadline.Get()
}

// GetTeacherHeadlineOk returns a tuple with the TeacherHeadline field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetTeacherHeadlineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeacherHeadline.Get(), o.TeacherHeadline.IsSet()
}

// SetTeacherHeadline sets field value
func (o *User) SetTeacherHeadline(v string) {
	o.TeacherHeadline.Set(&v)
}

// GetTeacherDescription returns the TeacherDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *User) GetTeacherDescription() string {
	if o == nil || o.TeacherDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.TeacherDescription.Get()
}

// GetTeacherDescriptionOk returns a tuple with the TeacherDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetTeacherDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeacherDescription.Get(), o.TeacherDescription.IsSet()
}

// SetTeacherDescription sets field value
func (o *User) SetTeacherDescription(v string) {
	o.TeacherDescription.Set(&v)
}

// GetLocale returns the Locale field value
// If the value is explicit nil, the zero value for Locale will be returned
func (o *User) GetLocale() Locale {
	if o == nil || o.Locale.Get() == nil {
		var ret Locale
		return ret
	}

	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLocaleOk() (*Locale, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locale.Get(), o.Locale.IsSet()
}

// SetLocale sets field value
func (o *User) SetLocale(v Locale) {
	o.Locale.Set(&v)
}

// GetWantsNewsletter returns the WantsNewsletter field value
func (o *User) GetWantsNewsletter() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WantsNewsletter
}

// GetWantsNewsletterOk returns a tuple with the WantsNewsletter field value
// and a boolean to check if the value has been set.
func (o *User) GetWantsNewsletterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WantsNewsletter, true
}

// SetWantsNewsletter sets field value
func (o *User) SetWantsNewsletter(v bool) {
	o.WantsNewsletter = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *User) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *User) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *User) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *User) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["first_name"] = o.FirstName
	toSerialize["middle_name"] = o.MiddleName.Get()
	toSerialize["last_name"] = o.LastName
	toSerialize["email"] = o.Email
	toSerialize["avatar_url"] = o.AvatarUrl
	toSerialize["roles"] = o.Roles
	toSerialize["notes_user"] = o.NotesUser.Get()
	toSerialize["description"] = o.Description.Get()
	toSerialize["label_ids"] = o.LabelIds
	toSerialize["employee_number"] = o.EmployeeNumber.Get()
	toSerialize["student_number"] = o.StudentNumber.Get()
	toSerialize["teacher_headline"] = o.TeacherHeadline.Get()
	toSerialize["teacher_description"] = o.TeacherDescription.Get()
	toSerialize["locale"] = o.Locale.Get()
	toSerialize["wants_newsletter"] = o.WantsNewsletter
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *User) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"first_name",
		"middle_name",
		"last_name",
		"email",
		"avatar_url",
		"roles",
		"notes_user",
		"description",
		"label_ids",
		"employee_number",
		"student_number",
		"teacher_headline",
		"teacher_description",
		"locale",
		"wants_newsletter",
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

	varUser := _User{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUser)

	if err != nil {
		return err
	}

	*o = User(varUser)

	return err
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


