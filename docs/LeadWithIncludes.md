# LeadWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the lead | [readonly] 
**CreationMethod** | **string** | Indicates the lead was created | 
**ReferralText** | **NullableString** | Text supplied at referral | 
**Title** | **string** | Title of the lead | 
**AccountId** | **NullableInt32** | ID of the account linked to this lead | 
**UserId** | **NullableInt32** | ID of the user linked to this lead | 
**LabelIds** | **[]int32** | An array of integers representing unique identifier values associated with labels.  | 
**Value** | **NullableString** | Decimal representing the price of a lead | 
**CompanyName** | **NullableString** | Name of the company where this lead comes from | 
**FirstName** | **NullableString** | The first name of the lead | 
**MiddleName** | **NullableString** | The middle name of the lead | 
**LastName** | **NullableString** | The last name of the lead | 
**AdministratorId** | **NullableInt32** | ID of administrator that owns the lead | 
**Email** | **NullableString** | The email of the lead | 
**Phone** | **NullableString** | The phone number of the lead | 
**Status** | **string** | The status of the lead | 
**Quality** | **NullableString** | Star scoring for the lead | 
**WantsNewsletter** | **bool** | Indicates if lead wants to receive the newsletter or not | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 
**Address** | Pointer to [**NullableAddress**](Address.md) |  | [optional] 
**LeadProducts** | Pointer to [**[]LeadProduct**](LeadProduct.md) |  | [optional] 

## Methods

### NewLeadWithIncludes

`func NewLeadWithIncludes(id int32, creationMethod string, referralText NullableString, title string, accountId NullableInt32, userId NullableInt32, labelIds []int32, value NullableString, companyName NullableString, firstName NullableString, middleName NullableString, lastName NullableString, administratorId NullableInt32, email NullableString, phone NullableString, status string, quality NullableString, wantsNewsletter bool, updatedAt time.Time, createdAt time.Time, ) *LeadWithIncludes`

NewLeadWithIncludes instantiates a new LeadWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeadWithIncludesWithDefaults

`func NewLeadWithIncludesWithDefaults() *LeadWithIncludes`

NewLeadWithIncludesWithDefaults instantiates a new LeadWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LeadWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LeadWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LeadWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetCreationMethod

`func (o *LeadWithIncludes) GetCreationMethod() string`

GetCreationMethod returns the CreationMethod field if non-nil, zero value otherwise.

### GetCreationMethodOk

`func (o *LeadWithIncludes) GetCreationMethodOk() (*string, bool)`

GetCreationMethodOk returns a tuple with the CreationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMethod

`func (o *LeadWithIncludes) SetCreationMethod(v string)`

SetCreationMethod sets CreationMethod field to given value.


### GetReferralText

`func (o *LeadWithIncludes) GetReferralText() string`

GetReferralText returns the ReferralText field if non-nil, zero value otherwise.

### GetReferralTextOk

`func (o *LeadWithIncludes) GetReferralTextOk() (*string, bool)`

GetReferralTextOk returns a tuple with the ReferralText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralText

`func (o *LeadWithIncludes) SetReferralText(v string)`

SetReferralText sets ReferralText field to given value.


### SetReferralTextNil

`func (o *LeadWithIncludes) SetReferralTextNil(b bool)`

 SetReferralTextNil sets the value for ReferralText to be an explicit nil

### UnsetReferralText
`func (o *LeadWithIncludes) UnsetReferralText()`

UnsetReferralText ensures that no value is present for ReferralText, not even an explicit nil
### GetTitle

`func (o *LeadWithIncludes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LeadWithIncludes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LeadWithIncludes) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAccountId

`func (o *LeadWithIncludes) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LeadWithIncludes) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LeadWithIncludes) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *LeadWithIncludes) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *LeadWithIncludes) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetUserId

`func (o *LeadWithIncludes) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LeadWithIncludes) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LeadWithIncludes) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *LeadWithIncludes) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *LeadWithIncludes) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetLabelIds

`func (o *LeadWithIncludes) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *LeadWithIncludes) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *LeadWithIncludes) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.


### GetValue

`func (o *LeadWithIncludes) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LeadWithIncludes) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LeadWithIncludes) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *LeadWithIncludes) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *LeadWithIncludes) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetCompanyName

`func (o *LeadWithIncludes) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *LeadWithIncludes) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *LeadWithIncludes) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### SetCompanyNameNil

`func (o *LeadWithIncludes) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *LeadWithIncludes) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetFirstName

`func (o *LeadWithIncludes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *LeadWithIncludes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *LeadWithIncludes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *LeadWithIncludes) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *LeadWithIncludes) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetMiddleName

`func (o *LeadWithIncludes) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *LeadWithIncludes) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *LeadWithIncludes) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.


### SetMiddleNameNil

`func (o *LeadWithIncludes) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *LeadWithIncludes) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *LeadWithIncludes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *LeadWithIncludes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *LeadWithIncludes) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *LeadWithIncludes) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *LeadWithIncludes) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAdministratorId

`func (o *LeadWithIncludes) GetAdministratorId() int32`

GetAdministratorId returns the AdministratorId field if non-nil, zero value otherwise.

### GetAdministratorIdOk

`func (o *LeadWithIncludes) GetAdministratorIdOk() (*int32, bool)`

GetAdministratorIdOk returns a tuple with the AdministratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorId

`func (o *LeadWithIncludes) SetAdministratorId(v int32)`

SetAdministratorId sets AdministratorId field to given value.


### SetAdministratorIdNil

`func (o *LeadWithIncludes) SetAdministratorIdNil(b bool)`

 SetAdministratorIdNil sets the value for AdministratorId to be an explicit nil

### UnsetAdministratorId
`func (o *LeadWithIncludes) UnsetAdministratorId()`

UnsetAdministratorId ensures that no value is present for AdministratorId, not even an explicit nil
### GetEmail

`func (o *LeadWithIncludes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LeadWithIncludes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LeadWithIncludes) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *LeadWithIncludes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LeadWithIncludes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *LeadWithIncludes) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *LeadWithIncludes) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *LeadWithIncludes) SetPhone(v string)`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *LeadWithIncludes) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *LeadWithIncludes) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetStatus

`func (o *LeadWithIncludes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LeadWithIncludes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LeadWithIncludes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetQuality

`func (o *LeadWithIncludes) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *LeadWithIncludes) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *LeadWithIncludes) SetQuality(v string)`

SetQuality sets Quality field to given value.


### SetQualityNil

`func (o *LeadWithIncludes) SetQualityNil(b bool)`

 SetQualityNil sets the value for Quality to be an explicit nil

### UnsetQuality
`func (o *LeadWithIncludes) UnsetQuality()`

UnsetQuality ensures that no value is present for Quality, not even an explicit nil
### GetWantsNewsletter

`func (o *LeadWithIncludes) GetWantsNewsletter() bool`

GetWantsNewsletter returns the WantsNewsletter field if non-nil, zero value otherwise.

### GetWantsNewsletterOk

`func (o *LeadWithIncludes) GetWantsNewsletterOk() (*bool, bool)`

GetWantsNewsletterOk returns a tuple with the WantsNewsletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantsNewsletter

`func (o *LeadWithIncludes) SetWantsNewsletter(v bool)`

SetWantsNewsletter sets WantsNewsletter field to given value.


### GetUpdatedAt

`func (o *LeadWithIncludes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LeadWithIncludes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LeadWithIncludes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *LeadWithIncludes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LeadWithIncludes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LeadWithIncludes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetAddress

`func (o *LeadWithIncludes) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LeadWithIncludes) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LeadWithIncludes) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LeadWithIncludes) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *LeadWithIncludes) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *LeadWithIncludes) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetLeadProducts

`func (o *LeadWithIncludes) GetLeadProducts() []LeadProduct`

GetLeadProducts returns the LeadProducts field if non-nil, zero value otherwise.

### GetLeadProductsOk

`func (o *LeadWithIncludes) GetLeadProductsOk() (*[]LeadProduct, bool)`

GetLeadProductsOk returns a tuple with the LeadProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadProducts

`func (o *LeadWithIncludes) SetLeadProducts(v []LeadProduct)`

SetLeadProducts sets LeadProducts field to given value.

### HasLeadProducts

`func (o *LeadWithIncludes) HasLeadProducts() bool`

HasLeadProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


