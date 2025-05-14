# Educator

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

## Methods

### NewEducator

`func NewEducator(id int32, slug string, standardPlanningText NullableString, defaultInvoiceVatMultiplier NullableString, phone NullableString, websiteUrl NullableString, signupDefaultAccountType string, signupContactInfo NullableString, currency Currency, country Country, email NullableString, timeZone string, name string, ) *Educator`

NewEducator instantiates a new Educator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEducatorWithDefaults

`func NewEducatorWithDefaults() *Educator`

NewEducatorWithDefaults instantiates a new Educator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Educator) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Educator) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Educator) SetId(v int32)`

SetId sets Id field to given value.


### GetSlug

`func (o *Educator) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Educator) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Educator) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetStandardPlanningText

`func (o *Educator) GetStandardPlanningText() string`

GetStandardPlanningText returns the StandardPlanningText field if non-nil, zero value otherwise.

### GetStandardPlanningTextOk

`func (o *Educator) GetStandardPlanningTextOk() (*string, bool)`

GetStandardPlanningTextOk returns a tuple with the StandardPlanningText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPlanningText

`func (o *Educator) SetStandardPlanningText(v string)`

SetStandardPlanningText sets StandardPlanningText field to given value.


### SetStandardPlanningTextNil

`func (o *Educator) SetStandardPlanningTextNil(b bool)`

 SetStandardPlanningTextNil sets the value for StandardPlanningText to be an explicit nil

### UnsetStandardPlanningText
`func (o *Educator) UnsetStandardPlanningText()`

UnsetStandardPlanningText ensures that no value is present for StandardPlanningText, not even an explicit nil
### GetDefaultInvoiceVatMultiplier

`func (o *Educator) GetDefaultInvoiceVatMultiplier() string`

GetDefaultInvoiceVatMultiplier returns the DefaultInvoiceVatMultiplier field if non-nil, zero value otherwise.

### GetDefaultInvoiceVatMultiplierOk

`func (o *Educator) GetDefaultInvoiceVatMultiplierOk() (*string, bool)`

GetDefaultInvoiceVatMultiplierOk returns a tuple with the DefaultInvoiceVatMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInvoiceVatMultiplier

`func (o *Educator) SetDefaultInvoiceVatMultiplier(v string)`

SetDefaultInvoiceVatMultiplier sets DefaultInvoiceVatMultiplier field to given value.


### SetDefaultInvoiceVatMultiplierNil

`func (o *Educator) SetDefaultInvoiceVatMultiplierNil(b bool)`

 SetDefaultInvoiceVatMultiplierNil sets the value for DefaultInvoiceVatMultiplier to be an explicit nil

### UnsetDefaultInvoiceVatMultiplier
`func (o *Educator) UnsetDefaultInvoiceVatMultiplier()`

UnsetDefaultInvoiceVatMultiplier ensures that no value is present for DefaultInvoiceVatMultiplier, not even an explicit nil
### GetPhone

`func (o *Educator) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Educator) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Educator) SetPhone(v string)`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *Educator) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *Educator) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetWebsiteUrl

`func (o *Educator) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *Educator) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *Educator) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.


### SetWebsiteUrlNil

`func (o *Educator) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *Educator) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetSignupDefaultAccountType

`func (o *Educator) GetSignupDefaultAccountType() string`

GetSignupDefaultAccountType returns the SignupDefaultAccountType field if non-nil, zero value otherwise.

### GetSignupDefaultAccountTypeOk

`func (o *Educator) GetSignupDefaultAccountTypeOk() (*string, bool)`

GetSignupDefaultAccountTypeOk returns a tuple with the SignupDefaultAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupDefaultAccountType

`func (o *Educator) SetSignupDefaultAccountType(v string)`

SetSignupDefaultAccountType sets SignupDefaultAccountType field to given value.


### GetSignupContactInfo

`func (o *Educator) GetSignupContactInfo() string`

GetSignupContactInfo returns the SignupContactInfo field if non-nil, zero value otherwise.

### GetSignupContactInfoOk

`func (o *Educator) GetSignupContactInfoOk() (*string, bool)`

GetSignupContactInfoOk returns a tuple with the SignupContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupContactInfo

`func (o *Educator) SetSignupContactInfo(v string)`

SetSignupContactInfo sets SignupContactInfo field to given value.


### SetSignupContactInfoNil

`func (o *Educator) SetSignupContactInfoNil(b bool)`

 SetSignupContactInfoNil sets the value for SignupContactInfo to be an explicit nil

### UnsetSignupContactInfo
`func (o *Educator) UnsetSignupContactInfo()`

UnsetSignupContactInfo ensures that no value is present for SignupContactInfo, not even an explicit nil
### GetCurrency

`func (o *Educator) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Educator) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Educator) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetCountry

`func (o *Educator) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Educator) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Educator) SetCountry(v Country)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *Educator) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Educator) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Educator) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *Educator) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Educator) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetTimeZone

`func (o *Educator) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Educator) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Educator) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetName

`func (o *Educator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Educator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Educator) SetName(v string)`

SetName sets Name field to given value.


### GetLocale

`func (o *Educator) GetLocale() Locale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Educator) GetLocaleOk() (*Locale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Educator) SetLocale(v Locale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *Educator) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *Educator) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *Educator) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetTermsUrl

`func (o *Educator) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *Educator) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *Educator) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *Educator) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


