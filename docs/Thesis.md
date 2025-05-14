# Thesis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the thesis. | [readonly] 
**Title** | **NullableString** | The title of the thesis. | 
**Subtitle** | **NullableString** | The subtitle of the thesis. | 
**Keywords** | **[]string** | The keywords of the thesis. | 
**Abstract** | **NullableString** | The abstract of the thesis. | 
**Authors** | **[]string** | Array of authors of the thesis. | 
**Date** | **NullableString** | Date of the program thesis. | 
**Supervisors** | **[]string** | Array of supervisors of the thesis. | 
**Confidentiality** | **NullableString** | Confidentiality level of the program thesis defaults to no. | 
**Locale** | **NullableString** | Locale of the program thesis. | 
**FileUrl** | **NullableString** | URL to the file of the thesis. | 

## Methods

### NewThesis

`func NewThesis(id int32, title NullableString, subtitle NullableString, keywords []string, abstract NullableString, authors []string, date NullableString, supervisors []string, confidentiality NullableString, locale NullableString, fileUrl NullableString, ) *Thesis`

NewThesis instantiates a new Thesis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThesisWithDefaults

`func NewThesisWithDefaults() *Thesis`

NewThesisWithDefaults instantiates a new Thesis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Thesis) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Thesis) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Thesis) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *Thesis) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Thesis) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Thesis) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *Thesis) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Thesis) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSubtitle

`func (o *Thesis) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *Thesis) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *Thesis) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.


### SetSubtitleNil

`func (o *Thesis) SetSubtitleNil(b bool)`

 SetSubtitleNil sets the value for Subtitle to be an explicit nil

### UnsetSubtitle
`func (o *Thesis) UnsetSubtitle()`

UnsetSubtitle ensures that no value is present for Subtitle, not even an explicit nil
### GetKeywords

`func (o *Thesis) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *Thesis) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *Thesis) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetAbstract

`func (o *Thesis) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *Thesis) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *Thesis) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.


### SetAbstractNil

`func (o *Thesis) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *Thesis) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetAuthors

`func (o *Thesis) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *Thesis) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *Thesis) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.


### GetDate

`func (o *Thesis) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Thesis) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Thesis) SetDate(v string)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *Thesis) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *Thesis) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetSupervisors

`func (o *Thesis) GetSupervisors() []string`

GetSupervisors returns the Supervisors field if non-nil, zero value otherwise.

### GetSupervisorsOk

`func (o *Thesis) GetSupervisorsOk() (*[]string, bool)`

GetSupervisorsOk returns a tuple with the Supervisors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisors

`func (o *Thesis) SetSupervisors(v []string)`

SetSupervisors sets Supervisors field to given value.


### GetConfidentiality

`func (o *Thesis) GetConfidentiality() string`

GetConfidentiality returns the Confidentiality field if non-nil, zero value otherwise.

### GetConfidentialityOk

`func (o *Thesis) GetConfidentialityOk() (*string, bool)`

GetConfidentialityOk returns a tuple with the Confidentiality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentiality

`func (o *Thesis) SetConfidentiality(v string)`

SetConfidentiality sets Confidentiality field to given value.


### SetConfidentialityNil

`func (o *Thesis) SetConfidentialityNil(b bool)`

 SetConfidentialityNil sets the value for Confidentiality to be an explicit nil

### UnsetConfidentiality
`func (o *Thesis) UnsetConfidentiality()`

UnsetConfidentiality ensures that no value is present for Confidentiality, not even an explicit nil
### GetLocale

`func (o *Thesis) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Thesis) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Thesis) SetLocale(v string)`

SetLocale sets Locale field to given value.


### SetLocaleNil

`func (o *Thesis) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *Thesis) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetFileUrl

`func (o *Thesis) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *Thesis) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *Thesis) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### SetFileUrlNil

`func (o *Thesis) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *Thesis) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


