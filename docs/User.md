# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the user. | [readonly] 
**FirstName** | **string** | First name of the user. | 
**MiddleName** | **NullableString** | Middle name of the user. | 
**LastName** | **string** | Last name of the user. | 
**Email** | **string** | The e-mail of the user. | 
**AvatarUrl** | **string** | The relative url path to the avatar. | 
**Roles** | **[]string** | List of roles the user has. | [readonly] 
**NotesUser** | **NullableString** | Short note about the user. | 
**Description** | **NullableString** | Long description of the user. (same as teacher_description) | 
**LabelIds** | **[]int32** | An array of integers representing unique identifier values associated with labels.  | 
**EmployeeNumber** | **NullableString** | The employee number of this user. | 
**StudentNumber** | **NullableString** | Unique identifier of the student. | 
**TeacherHeadline** | **NullableString** | Short description of the user. | 
**TeacherDescription** | **NullableString** | Long description of the user. (same as description) | 
**Locale** | [**NullableLocale**](Locale.md) |  | 
**WantsNewsletter** | **bool** | Boolean representing the possibility of the user to receive newsletters. | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewUser

`func NewUser(id int32, firstName string, middleName NullableString, lastName string, email string, avatarUrl string, roles []string, notesUser NullableString, description NullableString, labelIds []int32, employeeNumber NullableString, studentNumber NullableString, teacherHeadline NullableString, teacherDescription NullableString, locale NullableLocale, wantsNewsletter bool, updatedAt time.Time, createdAt time.Time, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.


### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *User) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *User) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *User) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.


### SetMiddleNameNil

`func (o *User) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *User) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAvatarUrl

`func (o *User) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *User) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *User) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetRoles

`func (o *User) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *User) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetNotesUser

`func (o *User) GetNotesUser() string`

GetNotesUser returns the NotesUser field if non-nil, zero value otherwise.

### GetNotesUserOk

`func (o *User) GetNotesUserOk() (*string, bool)`

GetNotesUserOk returns a tuple with the NotesUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUser

`func (o *User) SetNotesUser(v string)`

SetNotesUser sets NotesUser field to given value.


### SetNotesUserNil

`func (o *User) SetNotesUserNil(b bool)`

 SetNotesUserNil sets the value for NotesUser to be an explicit nil

### UnsetNotesUser
`func (o *User) UnsetNotesUser()`

UnsetNotesUser ensures that no value is present for NotesUser, not even an explicit nil
### GetDescription

`func (o *User) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *User) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *User) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *User) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *User) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabelIds

`func (o *User) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *User) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *User) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.


### GetEmployeeNumber

`func (o *User) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *User) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *User) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.


### SetEmployeeNumberNil

`func (o *User) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *User) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetStudentNumber

`func (o *User) GetStudentNumber() string`

GetStudentNumber returns the StudentNumber field if non-nil, zero value otherwise.

### GetStudentNumberOk

`func (o *User) GetStudentNumberOk() (*string, bool)`

GetStudentNumberOk returns a tuple with the StudentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentNumber

`func (o *User) SetStudentNumber(v string)`

SetStudentNumber sets StudentNumber field to given value.


### SetStudentNumberNil

`func (o *User) SetStudentNumberNil(b bool)`

 SetStudentNumberNil sets the value for StudentNumber to be an explicit nil

### UnsetStudentNumber
`func (o *User) UnsetStudentNumber()`

UnsetStudentNumber ensures that no value is present for StudentNumber, not even an explicit nil
### GetTeacherHeadline

`func (o *User) GetTeacherHeadline() string`

GetTeacherHeadline returns the TeacherHeadline field if non-nil, zero value otherwise.

### GetTeacherHeadlineOk

`func (o *User) GetTeacherHeadlineOk() (*string, bool)`

GetTeacherHeadlineOk returns a tuple with the TeacherHeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherHeadline

`func (o *User) SetTeacherHeadline(v string)`

SetTeacherHeadline sets TeacherHeadline field to given value.


### SetTeacherHeadlineNil

`func (o *User) SetTeacherHeadlineNil(b bool)`

 SetTeacherHeadlineNil sets the value for TeacherHeadline to be an explicit nil

### UnsetTeacherHeadline
`func (o *User) UnsetTeacherHeadline()`

UnsetTeacherHeadline ensures that no value is present for TeacherHeadline, not even an explicit nil
### GetTeacherDescription

`func (o *User) GetTeacherDescription() string`

GetTeacherDescription returns the TeacherDescription field if non-nil, zero value otherwise.

### GetTeacherDescriptionOk

`func (o *User) GetTeacherDescriptionOk() (*string, bool)`

GetTeacherDescriptionOk returns a tuple with the TeacherDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherDescription

`func (o *User) SetTeacherDescription(v string)`

SetTeacherDescription sets TeacherDescription field to given value.


### SetTeacherDescriptionNil

`func (o *User) SetTeacherDescriptionNil(b bool)`

 SetTeacherDescriptionNil sets the value for TeacherDescription to be an explicit nil

### UnsetTeacherDescription
`func (o *User) UnsetTeacherDescription()`

UnsetTeacherDescription ensures that no value is present for TeacherDescription, not even an explicit nil
### GetLocale

`func (o *User) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *User) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *User) SetLocale(v Locale)`

SetLocale sets Locale field to given value.


### SetLocaleNil

`func (o *User) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *User) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetWantsNewsletter

`func (o *User) GetWantsNewsletter() bool`

GetWantsNewsletter returns the WantsNewsletter field if non-nil, zero value otherwise.

### GetWantsNewsletterOk

`func (o *User) GetWantsNewsletterOk() (*bool, bool)`

GetWantsNewsletterOk returns a tuple with the WantsNewsletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantsNewsletter

`func (o *User) SetWantsNewsletter(v bool)`

SetWantsNewsletter sets WantsNewsletter field to given value.


### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


