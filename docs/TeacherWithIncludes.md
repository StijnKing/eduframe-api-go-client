# TeacherWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the user. | [readonly] 
**FirstName** | **string** | First name of the user. | 
**MiddleName** | **NullableString** | Middle name of the user. | 
**LastName** | **string** | Last name of the user. | 
**Email** | **string** | The e-mail of the user. | 
**LabelIds** | **[]int32** | An array of integers representing unique identifier values associated with labels.  | 
**Active** | **bool** | Whether the teacher is currently active. | 
**AvatarUrl** | **string** | The relative url path to the avatar. | 
**EmployeeNumber** | **NullableString** | The employee number of this user. | 
**Note** | **NullableString** | Short note about the teacher. | 
**TeacherHeadline** | **NullableString** | Short description of the user. | 
**TeacherDescription** | **NullableString** | Long description of the user. (same as description) | 
**Locale** | [**NullableLocale**](Locale.md) |  | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the user. | [optional] 

## Methods

### NewTeacherWithIncludes

`func NewTeacherWithIncludes(id int32, firstName string, middleName NullableString, lastName string, email string, labelIds []int32, active bool, avatarUrl string, employeeNumber NullableString, note NullableString, teacherHeadline NullableString, teacherDescription NullableString, locale NullableLocale, updatedAt time.Time, createdAt time.Time, ) *TeacherWithIncludes`

NewTeacherWithIncludes instantiates a new TeacherWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeacherWithIncludesWithDefaults

`func NewTeacherWithIncludesWithDefaults() *TeacherWithIncludes`

NewTeacherWithIncludesWithDefaults instantiates a new TeacherWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeacherWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeacherWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeacherWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetFirstName

`func (o *TeacherWithIncludes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TeacherWithIncludes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TeacherWithIncludes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *TeacherWithIncludes) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *TeacherWithIncludes) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *TeacherWithIncludes) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.


### SetMiddleNameNil

`func (o *TeacherWithIncludes) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *TeacherWithIncludes) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *TeacherWithIncludes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TeacherWithIncludes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TeacherWithIncludes) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *TeacherWithIncludes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TeacherWithIncludes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TeacherWithIncludes) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLabelIds

`func (o *TeacherWithIncludes) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *TeacherWithIncludes) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *TeacherWithIncludes) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.


### GetActive

`func (o *TeacherWithIncludes) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TeacherWithIncludes) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TeacherWithIncludes) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAvatarUrl

`func (o *TeacherWithIncludes) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *TeacherWithIncludes) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *TeacherWithIncludes) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetEmployeeNumber

`func (o *TeacherWithIncludes) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *TeacherWithIncludes) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *TeacherWithIncludes) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.


### SetEmployeeNumberNil

`func (o *TeacherWithIncludes) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *TeacherWithIncludes) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetNote

`func (o *TeacherWithIncludes) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *TeacherWithIncludes) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *TeacherWithIncludes) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *TeacherWithIncludes) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *TeacherWithIncludes) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetTeacherHeadline

`func (o *TeacherWithIncludes) GetTeacherHeadline() string`

GetTeacherHeadline returns the TeacherHeadline field if non-nil, zero value otherwise.

### GetTeacherHeadlineOk

`func (o *TeacherWithIncludes) GetTeacherHeadlineOk() (*string, bool)`

GetTeacherHeadlineOk returns a tuple with the TeacherHeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherHeadline

`func (o *TeacherWithIncludes) SetTeacherHeadline(v string)`

SetTeacherHeadline sets TeacherHeadline field to given value.


### SetTeacherHeadlineNil

`func (o *TeacherWithIncludes) SetTeacherHeadlineNil(b bool)`

 SetTeacherHeadlineNil sets the value for TeacherHeadline to be an explicit nil

### UnsetTeacherHeadline
`func (o *TeacherWithIncludes) UnsetTeacherHeadline()`

UnsetTeacherHeadline ensures that no value is present for TeacherHeadline, not even an explicit nil
### GetTeacherDescription

`func (o *TeacherWithIncludes) GetTeacherDescription() string`

GetTeacherDescription returns the TeacherDescription field if non-nil, zero value otherwise.

### GetTeacherDescriptionOk

`func (o *TeacherWithIncludes) GetTeacherDescriptionOk() (*string, bool)`

GetTeacherDescriptionOk returns a tuple with the TeacherDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherDescription

`func (o *TeacherWithIncludes) SetTeacherDescription(v string)`

SetTeacherDescription sets TeacherDescription field to given value.


### SetTeacherDescriptionNil

`func (o *TeacherWithIncludes) SetTeacherDescriptionNil(b bool)`

 SetTeacherDescriptionNil sets the value for TeacherDescription to be an explicit nil

### UnsetTeacherDescription
`func (o *TeacherWithIncludes) UnsetTeacherDescription()`

UnsetTeacherDescription ensures that no value is present for TeacherDescription, not even an explicit nil
### GetLocale

`func (o *TeacherWithIncludes) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TeacherWithIncludes) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TeacherWithIncludes) SetLocale(v Locale)`

SetLocale sets Locale field to given value.


### SetLocaleNil

`func (o *TeacherWithIncludes) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *TeacherWithIncludes) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetUpdatedAt

`func (o *TeacherWithIncludes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeacherWithIncludes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeacherWithIncludes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *TeacherWithIncludes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeacherWithIncludes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeacherWithIncludes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCustom

`func (o *TeacherWithIncludes) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *TeacherWithIncludes) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *TeacherWithIncludes) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *TeacherWithIncludes) HasCustom() bool`

HasCustom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


