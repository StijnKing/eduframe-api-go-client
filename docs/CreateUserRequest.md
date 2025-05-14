# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | First name of the user. | 
**MiddleName** | Pointer to **NullableString** | Middle name of the user. | [optional] 
**LastName** | **string** | Last name of the user. | 
**Email** | **string** | The e-mail of the user. | 
**Locale** | Pointer to [**NullableLocale**](Locale.md) |  | [optional] 
**WantsNewsletter** | Pointer to **bool** | Boolean representing the possibility of the user to receive newsletters. | [optional] 
**WithAuthentication** | Pointer to **bool** | If the user should be able to login and thus receive login details by mail. Only relevant when creating the user. | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the user. | [optional] 
**AddressAttributes** | Pointer to [**NullableAddressPayload**](AddressPayload.md) |  | [optional] 
**InvoiceAddressAttributes** | Pointer to [**NullableAddressPayload**](AddressPayload.md) |  | [optional] 
**LabelIds** | Pointer to **[]int32** | An array containing the identifiers of the labels associated with the user. When updating this array, the existing labels are replaced with the new ones provided.  | [optional] 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest(firstName string, lastName string, email string, ) *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *CreateUserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateUserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateUserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *CreateUserRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *CreateUserRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *CreateUserRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *CreateUserRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *CreateUserRequest) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *CreateUserRequest) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *CreateUserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateUserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateUserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *CreateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLocale

`func (o *CreateUserRequest) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateUserRequest) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateUserRequest) SetLocale(v Locale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CreateUserRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *CreateUserRequest) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *CreateUserRequest) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetWantsNewsletter

`func (o *CreateUserRequest) GetWantsNewsletter() bool`

GetWantsNewsletter returns the WantsNewsletter field if non-nil, zero value otherwise.

### GetWantsNewsletterOk

`func (o *CreateUserRequest) GetWantsNewsletterOk() (*bool, bool)`

GetWantsNewsletterOk returns a tuple with the WantsNewsletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantsNewsletter

`func (o *CreateUserRequest) SetWantsNewsletter(v bool)`

SetWantsNewsletter sets WantsNewsletter field to given value.

### HasWantsNewsletter

`func (o *CreateUserRequest) HasWantsNewsletter() bool`

HasWantsNewsletter returns a boolean if a field has been set.

### GetWithAuthentication

`func (o *CreateUserRequest) GetWithAuthentication() bool`

GetWithAuthentication returns the WithAuthentication field if non-nil, zero value otherwise.

### GetWithAuthenticationOk

`func (o *CreateUserRequest) GetWithAuthenticationOk() (*bool, bool)`

GetWithAuthenticationOk returns a tuple with the WithAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAuthentication

`func (o *CreateUserRequest) SetWithAuthentication(v bool)`

SetWithAuthentication sets WithAuthentication field to given value.

### HasWithAuthentication

`func (o *CreateUserRequest) HasWithAuthentication() bool`

HasWithAuthentication returns a boolean if a field has been set.

### GetCustom

`func (o *CreateUserRequest) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CreateUserRequest) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CreateUserRequest) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *CreateUserRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetAddressAttributes

`func (o *CreateUserRequest) GetAddressAttributes() AddressPayload`

GetAddressAttributes returns the AddressAttributes field if non-nil, zero value otherwise.

### GetAddressAttributesOk

`func (o *CreateUserRequest) GetAddressAttributesOk() (*AddressPayload, bool)`

GetAddressAttributesOk returns a tuple with the AddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributes

`func (o *CreateUserRequest) SetAddressAttributes(v AddressPayload)`

SetAddressAttributes sets AddressAttributes field to given value.

### HasAddressAttributes

`func (o *CreateUserRequest) HasAddressAttributes() bool`

HasAddressAttributes returns a boolean if a field has been set.

### SetAddressAttributesNil

`func (o *CreateUserRequest) SetAddressAttributesNil(b bool)`

 SetAddressAttributesNil sets the value for AddressAttributes to be an explicit nil

### UnsetAddressAttributes
`func (o *CreateUserRequest) UnsetAddressAttributes()`

UnsetAddressAttributes ensures that no value is present for AddressAttributes, not even an explicit nil
### GetInvoiceAddressAttributes

`func (o *CreateUserRequest) GetInvoiceAddressAttributes() AddressPayload`

GetInvoiceAddressAttributes returns the InvoiceAddressAttributes field if non-nil, zero value otherwise.

### GetInvoiceAddressAttributesOk

`func (o *CreateUserRequest) GetInvoiceAddressAttributesOk() (*AddressPayload, bool)`

GetInvoiceAddressAttributesOk returns a tuple with the InvoiceAddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAddressAttributes

`func (o *CreateUserRequest) SetInvoiceAddressAttributes(v AddressPayload)`

SetInvoiceAddressAttributes sets InvoiceAddressAttributes field to given value.

### HasInvoiceAddressAttributes

`func (o *CreateUserRequest) HasInvoiceAddressAttributes() bool`

HasInvoiceAddressAttributes returns a boolean if a field has been set.

### SetInvoiceAddressAttributesNil

`func (o *CreateUserRequest) SetInvoiceAddressAttributesNil(b bool)`

 SetInvoiceAddressAttributesNil sets the value for InvoiceAddressAttributes to be an explicit nil

### UnsetInvoiceAddressAttributes
`func (o *CreateUserRequest) UnsetInvoiceAddressAttributes()`

UnsetInvoiceAddressAttributes ensures that no value is present for InvoiceAddressAttributes, not even an explicit nil
### GetLabelIds

`func (o *CreateUserRequest) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *CreateUserRequest) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *CreateUserRequest) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *CreateUserRequest) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


