# EducatorWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the educator. | [readonly] 
**Slug** | **string** | Unique human readable identifier of the educator. | [readonly] 
**StandardPlanningText** | **NullableString** | Text that is shown by default on the planning page and signup. | 
**DefaultInvoiceVatMultiplier** | **NullableString** | Default VAT multiplier of the educator. | 
**Phone** | **NullableString** | Phone number of the educator. | 
**WebsiteUrl** | **NullableString** | Website URL of the educator. | 
**SignupDefaultAccountType** | **string** | What is de default selected account type on the signup page. | 
**SignupContactInfo** | **NullableString** | The contact info shown on the signup form (for directing questions to). | 
**Currency** | [**Currency**](Currency.md) |  | 
**Country** | [**Country**](Country.md) |  | 
**Email** | **NullableString** | The email (reply to) of the educator. | 
**TimeZone** | **string** | Time zone identifier of the educator. | 
**Name** | **string** | The name of the educator. | 
**Locale** | Pointer to [**NullableLocale**](Locale.md) |  | [optional] 
**TermsUrl** | Pointer to **string** | URL to the terms and conditions of the educator. | [optional] 
**InvoiceAddress** | Pointer to [**NullableAddress**](Address.md) |  | [optional] 

## Methods

### NewEducatorWithIncludes

`func NewEducatorWithIncludes(id int32, slug string, standardPlanningText NullableString, defaultInvoiceVatMultiplier NullableString, phone NullableString, websiteUrl NullableString, signupDefaultAccountType string, signupContactInfo NullableString, currency Currency, country Country, email NullableString, timeZone string, name string, ) *EducatorWithIncludes`

NewEducatorWithIncludes instantiates a new EducatorWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEducatorWithIncludesWithDefaults

`func NewEducatorWithIncludesWithDefaults() *EducatorWithIncludes`

NewEducatorWithIncludesWithDefaults instantiates a new EducatorWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EducatorWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EducatorWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EducatorWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetSlug

`func (o *EducatorWithIncludes) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EducatorWithIncludes) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EducatorWithIncludes) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetStandardPlanningText

`func (o *EducatorWithIncludes) GetStandardPlanningText() string`

GetStandardPlanningText returns the StandardPlanningText field if non-nil, zero value otherwise.

### GetStandardPlanningTextOk

`func (o *EducatorWithIncludes) GetStandardPlanningTextOk() (*string, bool)`

GetStandardPlanningTextOk returns a tuple with the StandardPlanningText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPlanningText

`func (o *EducatorWithIncludes) SetStandardPlanningText(v string)`

SetStandardPlanningText sets StandardPlanningText field to given value.


### SetStandardPlanningTextNil

`func (o *EducatorWithIncludes) SetStandardPlanningTextNil(b bool)`

 SetStandardPlanningTextNil sets the value for StandardPlanningText to be an explicit nil

### UnsetStandardPlanningText
`func (o *EducatorWithIncludes) UnsetStandardPlanningText()`

UnsetStandardPlanningText ensures that no value is present for StandardPlanningText, not even an explicit nil
### GetDefaultInvoiceVatMultiplier

`func (o *EducatorWithIncludes) GetDefaultInvoiceVatMultiplier() string`

GetDefaultInvoiceVatMultiplier returns the DefaultInvoiceVatMultiplier field if non-nil, zero value otherwise.

### GetDefaultInvoiceVatMultiplierOk

`func (o *EducatorWithIncludes) GetDefaultInvoiceVatMultiplierOk() (*string, bool)`

GetDefaultInvoiceVatMultiplierOk returns a tuple with the DefaultInvoiceVatMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInvoiceVatMultiplier

`func (o *EducatorWithIncludes) SetDefaultInvoiceVatMultiplier(v string)`

SetDefaultInvoiceVatMultiplier sets DefaultInvoiceVatMultiplier field to given value.


### SetDefaultInvoiceVatMultiplierNil

`func (o *EducatorWithIncludes) SetDefaultInvoiceVatMultiplierNil(b bool)`

 SetDefaultInvoiceVatMultiplierNil sets the value for DefaultInvoiceVatMultiplier to be an explicit nil

### UnsetDefaultInvoiceVatMultiplier
`func (o *EducatorWithIncludes) UnsetDefaultInvoiceVatMultiplier()`

UnsetDefaultInvoiceVatMultiplier ensures that no value is present for DefaultInvoiceVatMultiplier, not even an explicit nil
### GetPhone

`func (o *EducatorWithIncludes) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *EducatorWithIncludes) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *EducatorWithIncludes) SetPhone(v string)`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *EducatorWithIncludes) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *EducatorWithIncludes) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetWebsiteUrl

`func (o *EducatorWithIncludes) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *EducatorWithIncludes) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *EducatorWithIncludes) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.


### SetWebsiteUrlNil

`func (o *EducatorWithIncludes) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *EducatorWithIncludes) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetSignupDefaultAccountType

`func (o *EducatorWithIncludes) GetSignupDefaultAccountType() string`

GetSignupDefaultAccountType returns the SignupDefaultAccountType field if non-nil, zero value otherwise.

### GetSignupDefaultAccountTypeOk

`func (o *EducatorWithIncludes) GetSignupDefaultAccountTypeOk() (*string, bool)`

GetSignupDefaultAccountTypeOk returns a tuple with the SignupDefaultAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupDefaultAccountType

`func (o *EducatorWithIncludes) SetSignupDefaultAccountType(v string)`

SetSignupDefaultAccountType sets SignupDefaultAccountType field to given value.


### GetSignupContactInfo

`func (o *EducatorWithIncludes) GetSignupContactInfo() string`

GetSignupContactInfo returns the SignupContactInfo field if non-nil, zero value otherwise.

### GetSignupContactInfoOk

`func (o *EducatorWithIncludes) GetSignupContactInfoOk() (*string, bool)`

GetSignupContactInfoOk returns a tuple with the SignupContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupContactInfo

`func (o *EducatorWithIncludes) SetSignupContactInfo(v string)`

SetSignupContactInfo sets SignupContactInfo field to given value.


### SetSignupContactInfoNil

`func (o *EducatorWithIncludes) SetSignupContactInfoNil(b bool)`

 SetSignupContactInfoNil sets the value for SignupContactInfo to be an explicit nil

### UnsetSignupContactInfo
`func (o *EducatorWithIncludes) UnsetSignupContactInfo()`

UnsetSignupContactInfo ensures that no value is present for SignupContactInfo, not even an explicit nil
### GetCurrency

`func (o *EducatorWithIncludes) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EducatorWithIncludes) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EducatorWithIncludes) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetCountry

`func (o *EducatorWithIncludes) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *EducatorWithIncludes) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *EducatorWithIncludes) SetCountry(v Country)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *EducatorWithIncludes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EducatorWithIncludes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EducatorWithIncludes) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *EducatorWithIncludes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *EducatorWithIncludes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetTimeZone

`func (o *EducatorWithIncludes) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *EducatorWithIncludes) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *EducatorWithIncludes) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetName

`func (o *EducatorWithIncludes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EducatorWithIncludes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EducatorWithIncludes) SetName(v string)`

SetName sets Name field to given value.


### GetLocale

`func (o *EducatorWithIncludes) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *EducatorWithIncludes) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *EducatorWithIncludes) SetLocale(v Locale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *EducatorWithIncludes) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *EducatorWithIncludes) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *EducatorWithIncludes) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetTermsUrl

`func (o *EducatorWithIncludes) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *EducatorWithIncludes) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *EducatorWithIncludes) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *EducatorWithIncludes) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### GetInvoiceAddress

`func (o *EducatorWithIncludes) GetInvoiceAddress() Address`

GetInvoiceAddress returns the InvoiceAddress field if non-nil, zero value otherwise.

### GetInvoiceAddressOk

`func (o *EducatorWithIncludes) GetInvoiceAddressOk() (*Address, bool)`

GetInvoiceAddressOk returns a tuple with the InvoiceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAddress

`func (o *EducatorWithIncludes) SetInvoiceAddress(v Address)`

SetInvoiceAddress sets InvoiceAddress field to given value.

### HasInvoiceAddress

`func (o *EducatorWithIncludes) HasInvoiceAddress() bool`

HasInvoiceAddress returns a boolean if a field has been set.

### SetInvoiceAddressNil

`func (o *EducatorWithIncludes) SetInvoiceAddressNil(b bool)`

 SetInvoiceAddressNil sets the value for InvoiceAddress to be an explicit nil

### UnsetInvoiceAddress
`func (o *EducatorWithIncludes) UnsetInvoiceAddress()`

UnsetInvoiceAddress ensures that no value is present for InvoiceAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


