# Teacher

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

## Methods

### NewTeacher

`func NewTeacher(id int32, firstName string, middleName NullableString, lastName string, email string, labelIds []int32, active bool, avatarUrl string, employeeNumber NullableString, note NullableString, teacherHeadline NullableString, teacherDescription NullableString, locale NullableLocale, updatedAt time.Time, createdAt time.Time, ) *Teacher`

NewTeacher instantiates a new Teacher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeacherWithDefaults

`func NewTeacherWithDefaults() *Teacher`

NewTeacherWithDefaults instantiates a new Teacher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Teacher) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Teacher) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Teacher) SetId(v int32)`

SetId sets Id field to given value.


### GetFirstName

`func (o *Teacher) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Teacher) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Teacher) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *Teacher) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *Teacher) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *Teacher) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.


### SetMiddleNameNil

`func (o *Teacher) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *Teacher) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *Teacher) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Teacher) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Teacher) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *Teacher) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Teacher) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Teacher) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLabelIds

`func (o *Teacher) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *Teacher) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *Teacher) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.


### GetActive

`func (o *Teacher) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Teacher) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Teacher) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAvatarUrl

`func (o *Teacher) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Teacher) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Teacher) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetEmployeeNumber

`func (o *Teacher) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *Teacher) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *Teacher) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.


### SetEmployeeNumberNil

`func (o *Teacher) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *Teacher) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetNote

`func (o *Teacher) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Teacher) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Teacher) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *Teacher) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *Teacher) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetTeacherHeadline

`func (o *Teacher) GetTeacherHeadline() string`

GetTeacherHeadline returns the TeacherHeadline field if non-nil, zero value otherwise.

### GetTeacherHeadlineOk

`func (o *Teacher) GetTeacherHeadlineOk() (*string, bool)`

GetTeacherHeadlineOk returns a tuple with the TeacherHeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherHeadline

`func (o *Teacher) SetTeacherHeadline(v string)`

SetTeacherHeadline sets TeacherHeadline field to given value.


### SetTeacherHeadlineNil

`func (o *Teacher) SetTeacherHeadlineNil(b bool)`

 SetTeacherHeadlineNil sets the value for TeacherHeadline to be an explicit nil

### UnsetTeacherHeadline
`func (o *Teacher) UnsetTeacherHeadline()`

UnsetTeacherHeadline ensures that no value is present for TeacherHeadline, not even an explicit nil
### GetTeacherDescription

`func (o *Teacher) GetTeacherDescription() string`

GetTeacherDescription returns the TeacherDescription field if non-nil, zero value otherwise.

### GetTeacherDescriptionOk

`func (o *Teacher) GetTeacherDescriptionOk() (*string, bool)`

GetTeacherDescriptionOk returns a tuple with the TeacherDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherDescription

`func (o *Teacher) SetTeacherDescription(v string)`

SetTeacherDescription sets TeacherDescription field to given value.


### SetTeacherDescriptionNil

`func (o *Teacher) SetTeacherDescriptionNil(b bool)`

 SetTeacherDescriptionNil sets the value for TeacherDescription to be an explicit nil

### UnsetTeacherDescription
`func (o *Teacher) UnsetTeacherDescription()`

UnsetTeacherDescription ensures that no value is present for TeacherDescription, not even an explicit nil
### GetLocale

`func (o *Teacher) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Teacher) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Teacher) SetLocale(v Locale)`

SetLocale sets Locale field to given value.


### SetLocaleNil

`func (o *Teacher) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *Teacher) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetUpdatedAt

`func (o *Teacher) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Teacher) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Teacher) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Teacher) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Teacher) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Teacher) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


