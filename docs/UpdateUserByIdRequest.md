# UpdateUserByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First name of the user. | [optional] 
**MiddleName** | Pointer to **NullableString** | Middle name of the user. | [optional] 
**LastName** | Pointer to **string** | Last name of the user. | [optional] 
**Email** | Pointer to **string** | The e-mail of the user. | [optional] 
**Locale** | Pointer to [**NullableLocale**](Locale.md) |  | [optional] 
**WantsNewsletter** | Pointer to **bool** | Boolean representing the possibility of the user to receive newsletters. | [optional] 
**WithAuthentication** | Pointer to **bool** | If the user should be able to login and thus receive login details by mail. Only relevant when creating the user. | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the user. | [optional] 
**AddressAttributes** | Pointer to [**NullableAddressPatchPayload**](AddressPatchPayload.md) |  | [optional] 
**InvoiceAddressAttributes** | Pointer to [**NullableAddressPatchPayload**](AddressPatchPayload.md) |  | [optional] 
**LabelIds** | Pointer to **[]int32** | IDs of the labels | [optional] 

## Methods

### NewUpdateUserByIdRequest

`func NewUpdateUserByIdRequest() *UpdateUserByIdRequest`

NewUpdateUserByIdRequest instantiates a new UpdateUserByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserByIdRequestWithDefaults

`func NewUpdateUserByIdRequestWithDefaults() *UpdateUserByIdRequest`

NewUpdateUserByIdRequestWithDefaults instantiates a new UpdateUserByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UpdateUserByIdRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateUserByIdRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateUserByIdRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateUserByIdRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *UpdateUserByIdRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *UpdateUserByIdRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *UpdateUserByIdRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *UpdateUserByIdRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *UpdateUserByIdRequest) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *UpdateUserByIdRequest) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *UpdateUserByIdRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateUserByIdRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateUserByIdRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateUserByIdRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateUserByIdRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserByIdRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserByIdRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUserByIdRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLocale

`func (o *UpdateUserByIdRequest) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UpdateUserByIdRequest) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UpdateUserByIdRequest) SetLocale(v Locale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UpdateUserByIdRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *UpdateUserByIdRequest) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *UpdateUserByIdRequest) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetWantsNewsletter

`func (o *UpdateUserByIdRequest) GetWantsNewsletter() bool`

GetWantsNewsletter returns the WantsNewsletter field if non-nil, zero value otherwise.

### GetWantsNewsletterOk

`func (o *UpdateUserByIdRequest) GetWantsNewsletterOk() (*bool, bool)`

GetWantsNewsletterOk returns a tuple with the WantsNewsletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantsNewsletter

`func (o *UpdateUserByIdRequest) SetWantsNewsletter(v bool)`

SetWantsNewsletter sets WantsNewsletter field to given value.

### HasWantsNewsletter

`func (o *UpdateUserByIdRequest) HasWantsNewsletter() bool`

HasWantsNewsletter returns a boolean if a field has been set.

### GetWithAuthentication

`func (o *UpdateUserByIdRequest) GetWithAuthentication() bool`

GetWithAuthentication returns the WithAuthentication field if non-nil, zero value otherwise.

### GetWithAuthenticationOk

`func (o *UpdateUserByIdRequest) GetWithAuthenticationOk() (*bool, bool)`

GetWithAuthenticationOk returns a tuple with the WithAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAuthentication

`func (o *UpdateUserByIdRequest) SetWithAuthentication(v bool)`

SetWithAuthentication sets WithAuthentication field to given value.

### HasWithAuthentication

`func (o *UpdateUserByIdRequest) HasWithAuthentication() bool`

HasWithAuthentication returns a boolean if a field has been set.

### GetCustom

`func (o *UpdateUserByIdRequest) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UpdateUserByIdRequest) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UpdateUserByIdRequest) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *UpdateUserByIdRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetAddressAttributes

`func (o *UpdateUserByIdRequest) GetAddressAttributes() AddressPatchPayload`

GetAddressAttributes returns the AddressAttributes field if non-nil, zero value otherwise.

### GetAddressAttributesOk

`func (o *UpdateUserByIdRequest) GetAddressAttributesOk() (*AddressPatchPayload, bool)`

GetAddressAttributesOk returns a tuple with the AddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributes

`func (o *UpdateUserByIdRequest) SetAddressAttributes(v AddressPatchPayload)`

SetAddressAttributes sets AddressAttributes field to given value.

### HasAddressAttributes

`func (o *UpdateUserByIdRequest) HasAddressAttributes() bool`

HasAddressAttributes returns a boolean if a field has been set.

### SetAddressAttributesNil

`func (o *UpdateUserByIdRequest) SetAddressAttributesNil(b bool)`

 SetAddressAttributesNil sets the value for AddressAttributes to be an explicit nil

### UnsetAddressAttributes
`func (o *UpdateUserByIdRequest) UnsetAddressAttributes()`

UnsetAddressAttributes ensures that no value is present for AddressAttributes, not even an explicit nil
### GetInvoiceAddressAttributes

`func (o *UpdateUserByIdRequest) GetInvoiceAddressAttributes() AddressPatchPayload`

GetInvoiceAddressAttributes returns the InvoiceAddressAttributes field if non-nil, zero value otherwise.

### GetInvoiceAddressAttributesOk

`func (o *UpdateUserByIdRequest) GetInvoiceAddressAttributesOk() (*AddressPatchPayload, bool)`

GetInvoiceAddressAttributesOk returns a tuple with the InvoiceAddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAddressAttributes

`func (o *UpdateUserByIdRequest) SetInvoiceAddressAttributes(v AddressPatchPayload)`

SetInvoiceAddressAttributes sets InvoiceAddressAttributes field to given value.

### HasInvoiceAddressAttributes

`func (o *UpdateUserByIdRequest) HasInvoiceAddressAttributes() bool`

HasInvoiceAddressAttributes returns a boolean if a field has been set.

### SetInvoiceAddressAttributesNil

`func (o *UpdateUserByIdRequest) SetInvoiceAddressAttributesNil(b bool)`

 SetInvoiceAddressAttributesNil sets the value for InvoiceAddressAttributes to be an explicit nil

### UnsetInvoiceAddressAttributes
`func (o *UpdateUserByIdRequest) UnsetInvoiceAddressAttributes()`

UnsetInvoiceAddressAttributes ensures that no value is present for InvoiceAddressAttributes, not even an explicit nil
### GetLabelIds

`func (o *UpdateUserByIdRequest) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *UpdateUserByIdRequest) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *UpdateUserByIdRequest) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *UpdateUserByIdRequest) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


