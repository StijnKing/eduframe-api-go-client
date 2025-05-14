# UserWithIncludes

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
**Address** | Pointer to [**NullableAddress**](Address.md) |  | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the user. | [optional] 

## Methods

### NewUserWithIncludes

`func NewUserWithIncludes(id int32, firstName string, middleName NullableString, lastName string, email string, avatarUrl string, roles []string, notesUser NullableString, description NullableString, labelIds []int32, employeeNumber NullableString, studentNumber NullableString, teacherHeadline NullableString, teacherDescription NullableString, locale NullableLocale, wantsNewsletter bool, updatedAt time.Time, createdAt time.Time, ) *UserWithIncludes`

NewUserWithIncludes instantiates a new UserWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithIncludesWithDefaults

`func NewUserWithIncludesWithDefaults() *UserWithIncludes`

NewUserWithIncludesWithDefaults instantiates a new UserWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetFirstName

`func (o *UserWithIncludes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserWithIncludes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserWithIncludes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *UserWithIncludes) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *UserWithIncludes) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *UserWithIncludes) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.


### SetMiddleNameNil

`func (o *UserWithIncludes) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *UserWithIncludes) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *UserWithIncludes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserWithIncludes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserWithIncludes) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *UserWithIncludes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserWithIncludes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserWithIncludes) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAvatarUrl

`func (o *UserWithIncludes) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserWithIncludes) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserWithIncludes) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetRoles

`func (o *UserWithIncludes) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserWithIncludes) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserWithIncludes) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetNotesUser

`func (o *UserWithIncludes) GetNotesUser() string`

GetNotesUser returns the NotesUser field if non-nil, zero value otherwise.

### GetNotesUserOk

`func (o *UserWithIncludes) GetNotesUserOk() (*string, bool)`

GetNotesUserOk returns a tuple with the NotesUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUser

`func (o *UserWithIncludes) SetNotesUser(v string)`

SetNotesUser sets NotesUser field to given value.


### SetNotesUserNil

`func (o *UserWithIncludes) SetNotesUserNil(b bool)`

 SetNotesUserNil sets the value for NotesUser to be an explicit nil

### UnsetNotesUser
`func (o *UserWithIncludes) UnsetNotesUser()`

UnsetNotesUser ensures that no value is present for NotesUser, not even an explicit nil
### GetDescription

`func (o *UserWithIncludes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserWithIncludes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserWithIncludes) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *UserWithIncludes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UserWithIncludes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabelIds

`func (o *UserWithIncludes) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *UserWithIncludes) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *UserWithIncludes) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.


### GetEmployeeNumber

`func (o *UserWithIncludes) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *UserWithIncludes) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *UserWithIncludes) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.


### SetEmployeeNumberNil

`func (o *UserWithIncludes) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *UserWithIncludes) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetStudentNumber

`func (o *UserWithIncludes) GetStudentNumber() string`

GetStudentNumber returns the StudentNumber field if non-nil, zero value otherwise.

### GetStudentNumberOk

`func (o *UserWithIncludes) GetStudentNumberOk() (*string, bool)`

GetStudentNumberOk returns a tuple with the StudentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentNumber

`func (o *UserWithIncludes) SetStudentNumber(v string)`

SetStudentNumber sets StudentNumber field to given value.


### SetStudentNumberNil

`func (o *UserWithIncludes) SetStudentNumberNil(b bool)`

 SetStudentNumberNil sets the value for StudentNumber to be an explicit nil

### UnsetStudentNumber
`func (o *UserWithIncludes) UnsetStudentNumber()`

UnsetStudentNumber ensures that no value is present for StudentNumber, not even an explicit nil
### GetTeacherHeadline

`func (o *UserWithIncludes) GetTeacherHeadline() string`

GetTeacherHeadline returns the TeacherHeadline field if non-nil, zero value otherwise.

### GetTeacherHeadlineOk

`func (o *UserWithIncludes) GetTeacherHeadlineOk() (*string, bool)`

GetTeacherHeadlineOk returns a tuple with the TeacherHeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherHeadline

`func (o *UserWithIncludes) SetTeacherHeadline(v string)`

SetTeacherHeadline sets TeacherHeadline field to given value.


### SetTeacherHeadlineNil

`func (o *UserWithIncludes) SetTeacherHeadlineNil(b bool)`

 SetTeacherHeadlineNil sets the value for TeacherHeadline to be an explicit nil

### UnsetTeacherHeadline
`func (o *UserWithIncludes) UnsetTeacherHeadline()`

UnsetTeacherHeadline ensures that no value is present for TeacherHeadline, not even an explicit nil
### GetTeacherDescription

`func (o *UserWithIncludes) GetTeacherDescription() string`

GetTeacherDescription returns the TeacherDescription field if non-nil, zero value otherwise.

### GetTeacherDescriptionOk

`func (o *UserWithIncludes) GetTeacherDescriptionOk() (*string, bool)`

GetTeacherDescriptionOk returns a tuple with the TeacherDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeacherDescription

`func (o *UserWithIncludes) SetTeacherDescription(v string)`

SetTeacherDescription sets TeacherDescription field to given value.


### SetTeacherDescriptionNil

`func (o *UserWithIncludes) SetTeacherDescriptionNil(b bool)`

 SetTeacherDescriptionNil sets the value for TeacherDescription to be an explicit nil

### UnsetTeacherDescription
`func (o *UserWithIncludes) UnsetTeacherDescription()`

UnsetTeacherDescription ensures that no value is present for TeacherDescription, not even an explicit nil
### GetLocale

`func (o *UserWithIncludes) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UserWithIncludes) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UserWithIncludes) SetLocale(v Locale)`

SetLocale sets Locale field to given value.


### SetLocaleNil

`func (o *UserWithIncludes) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *UserWithIncludes) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetWantsNewsletter

`func (o *UserWithIncludes) GetWantsNewsletter() bool`

GetWantsNewsletter returns the WantsNewsletter field if non-nil, zero value otherwise.

### GetWantsNewsletterOk

`func (o *UserWithIncludes) GetWantsNewsletterOk() (*bool, bool)`

GetWantsNewsletterOk returns a tuple with the WantsNewsletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantsNewsletter

`func (o *UserWithIncludes) SetWantsNewsletter(v bool)`

SetWantsNewsletter sets WantsNewsletter field to given value.


### GetUpdatedAt

`func (o *UserWithIncludes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserWithIncludes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserWithIncludes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *UserWithIncludes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserWithIncludes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserWithIncludes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetAddress

`func (o *UserWithIncludes) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UserWithIncludes) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UserWithIncludes) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UserWithIncludes) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *UserWithIncludes) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *UserWithIncludes) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCustom

`func (o *UserWithIncludes) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UserWithIncludes) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UserWithIncludes) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *UserWithIncludes) HasCustom() bool`

HasCustom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


