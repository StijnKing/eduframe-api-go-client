# CreateLeadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the lead | [optional] 
**AccountId** | Pointer to **NullableInt32** | ID of the account linked to this lead | [optional] 
**UserId** | Pointer to **NullableInt32** | ID of the user linked to this lead | [optional] 
**Value** | Pointer to **NullableString** | Decimal representing the price of a lead | [optional] 
**CompanyName** | Pointer to **NullableString** | Name of the company where this lead comes from | [optional] 
**FirstName** | Pointer to **NullableString** | The first name of the lead | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name of the lead | [optional] 
**LastName** | Pointer to **NullableString** | The last name of the lead | [optional] 
**AdministratorId** | Pointer to **NullableInt32** | ID of administrator that owns the lead | [optional] 
**Email** | Pointer to **NullableString** | The email of the lead | [optional] 
**Phone** | Pointer to **NullableString** | The phone number of the lead **Note** : Use an international phone format unless the phone number is from the educator configured country.  | [optional] 
**Status** | Pointer to **string** | The status of the lead | [optional] 
**Quality** | Pointer to **NullableFloat32** | Star scoring for the lead | [optional] 
**WantsNewsletter** | Pointer to **bool** | Indicates if lead wants to receive the newsletter or not | [optional] 
**Comment** | Pointer to **NullableString** | Comment for a lead | [optional] 
**LabelIds** | Pointer to **[]int32** | IDs of the labels | [optional] 
**AddressAttributes** | Pointer to [**NullableAddressPayload**](AddressPayload.md) |  | [optional] 
**LeadProducts** | Pointer to [**[]CreateLeadRequestLeadProductsInner**](CreateLeadRequestLeadProductsInner.md) | Array of products and variants the lead is interested in.  | [optional] 

## Methods

### NewCreateLeadRequest

`func NewCreateLeadRequest() *CreateLeadRequest`

NewCreateLeadRequest instantiates a new CreateLeadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLeadRequestWithDefaults

`func NewCreateLeadRequestWithDefaults() *CreateLeadRequest`

NewCreateLeadRequestWithDefaults instantiates a new CreateLeadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateLeadRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateLeadRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateLeadRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateLeadRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateLeadRequest) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateLeadRequest) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateLeadRequest) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateLeadRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *CreateLeadRequest) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *CreateLeadRequest) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetUserId

`func (o *CreateLeadRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateLeadRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateLeadRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateLeadRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CreateLeadRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CreateLeadRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetValue

`func (o *CreateLeadRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateLeadRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateLeadRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateLeadRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CreateLeadRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CreateLeadRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetCompanyName

`func (o *CreateLeadRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CreateLeadRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CreateLeadRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CreateLeadRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CreateLeadRequest) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CreateLeadRequest) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetFirstName

`func (o *CreateLeadRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateLeadRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateLeadRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CreateLeadRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *CreateLeadRequest) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *CreateLeadRequest) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetMiddleName

`func (o *CreateLeadRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *CreateLeadRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *CreateLeadRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *CreateLeadRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *CreateLeadRequest) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *CreateLeadRequest) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *CreateLeadRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateLeadRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateLeadRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CreateLeadRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *CreateLeadRequest) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *CreateLeadRequest) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAdministratorId

`func (o *CreateLeadRequest) GetAdministratorId() int32`

GetAdministratorId returns the AdministratorId field if non-nil, zero value otherwise.

### GetAdministratorIdOk

`func (o *CreateLeadRequest) GetAdministratorIdOk() (*int32, bool)`

GetAdministratorIdOk returns a tuple with the AdministratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorId

`func (o *CreateLeadRequest) SetAdministratorId(v int32)`

SetAdministratorId sets AdministratorId field to given value.

### HasAdministratorId

`func (o *CreateLeadRequest) HasAdministratorId() bool`

HasAdministratorId returns a boolean if a field has been set.

### SetAdministratorIdNil

`func (o *CreateLeadRequest) SetAdministratorIdNil(b bool)`

 SetAdministratorIdNil sets the value for AdministratorId to be an explicit nil

### UnsetAdministratorId
`func (o *CreateLeadRequest) UnsetAdministratorId()`

UnsetAdministratorId ensures that no value is present for AdministratorId, not even an explicit nil
### GetEmail

`func (o *CreateLeadRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateLeadRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateLeadRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateLeadRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CreateLeadRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateLeadRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *CreateLeadRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateLeadRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateLeadRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CreateLeadRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *CreateLeadRequest) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *CreateLeadRequest) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetStatus

`func (o *CreateLeadRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateLeadRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateLeadRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateLeadRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetQuality

`func (o *CreateLeadRequest) GetQuality() float32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *CreateLeadRequest) GetQualityOk() (*float32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *CreateLeadRequest) SetQuality(v float32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *CreateLeadRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### SetQualityNil

`func (o *CreateLeadRequest) SetQualityNil(b bool)`

 SetQualityNil sets the value for Quality to be an explicit nil

### UnsetQuality
`func (o *CreateLeadRequest) UnsetQuality()`

UnsetQuality ensures that no value is present for Quality, not even an explicit nil
### GetWantsNewsletter

`func (o *CreateLeadRequest) GetWantsNewsletter() bool`

GetWantsNewsletter returns the WantsNewsletter field if non-nil, zero value otherwise.

### GetWantsNewsletterOk

`func (o *CreateLeadRequest) GetWantsNewsletterOk() (*bool, bool)`

GetWantsNewsletterOk returns a tuple with the WantsNewsletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantsNewsletter

`func (o *CreateLeadRequest) SetWantsNewsletter(v bool)`

SetWantsNewsletter sets WantsNewsletter field to given value.

### HasWantsNewsletter

`func (o *CreateLeadRequest) HasWantsNewsletter() bool`

HasWantsNewsletter returns a boolean if a field has been set.

### GetComment

`func (o *CreateLeadRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateLeadRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateLeadRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateLeadRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *CreateLeadRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *CreateLeadRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLabelIds

`func (o *CreateLeadRequest) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *CreateLeadRequest) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *CreateLeadRequest) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *CreateLeadRequest) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetAddressAttributes

`func (o *CreateLeadRequest) GetAddressAttributes() AddressPayload`

GetAddressAttributes returns the AddressAttributes field if non-nil, zero value otherwise.

### GetAddressAttributesOk

`func (o *CreateLeadRequest) GetAddressAttributesOk() (*AddressPayload, bool)`

GetAddressAttributesOk returns a tuple with the AddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributes

`func (o *CreateLeadRequest) SetAddressAttributes(v AddressPayload)`

SetAddressAttributes sets AddressAttributes field to given value.

### HasAddressAttributes

`func (o *CreateLeadRequest) HasAddressAttributes() bool`

HasAddressAttributes returns a boolean if a field has been set.

### SetAddressAttributesNil

`func (o *CreateLeadRequest) SetAddressAttributesNil(b bool)`

 SetAddressAttributesNil sets the value for AddressAttributes to be an explicit nil

### UnsetAddressAttributes
`func (o *CreateLeadRequest) UnsetAddressAttributes()`

UnsetAddressAttributes ensures that no value is present for AddressAttributes, not even an explicit nil
### GetLeadProducts

`func (o *CreateLeadRequest) GetLeadProducts() []CreateLeadRequestLeadProductsInner`

GetLeadProducts returns the LeadProducts field if non-nil, zero value otherwise.

### GetLeadProductsOk

`func (o *CreateLeadRequest) GetLeadProductsOk() (*[]CreateLeadRequestLeadProductsInner, bool)`

GetLeadProductsOk returns a tuple with the LeadProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadProducts

`func (o *CreateLeadRequest) SetLeadProducts(v []CreateLeadRequestLeadProductsInner)`

SetLeadProducts sets LeadProducts field to given value.

### HasLeadProducts

`func (o *CreateLeadRequest) HasLeadProducts() bool`

HasLeadProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


