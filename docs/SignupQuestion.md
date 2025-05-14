# SignupQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the question. | [readonly] 
**Position** | **int32** | Position of the question used for ordering. | 
**FieldType** | [**SignupQuestionFieldType**](SignupQuestionFieldType.md) |  | 
**Title** | **string** | The title and label of the question | 
**Slug** | **string** | The slug of the question | 
**Required** | **bool** | Define if this question is required on the Signup page. | 
**ForType** | **string** | Define to what type of models this should be linked. | 
**Visibility** | **[]string** | List of at what locations you want to show this field in the signup | 
**ForStudent** | **bool** | DEPRECATED: boolean if the question is visible for students. Please use the visibility attribute. | 
**ForCustomer** | **bool** | DEPRECATED: boolean if the question is visible for customers. Please use the visibility attribute. | 
**Choices** | **[]string** | Array of string with the choice options if the type of the field is select. | 
**SystemName** | **NullableString** | The internal system name used for this field. | 
**UseAsDuplicateIndicator** | **bool** | boolean if the question is use as duplicate indicator. | 

## Methods

### NewSignupQuestion

`func NewSignupQuestion(id int32, position int32, fieldType SignupQuestionFieldType, title string, slug string, required bool, forType string, visibility []string, forStudent bool, forCustomer bool, choices []string, systemName NullableString, useAsDuplicateIndicator bool, ) *SignupQuestion`

NewSignupQuestion instantiates a new SignupQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignupQuestionWithDefaults

`func NewSignupQuestionWithDefaults() *SignupQuestion`

NewSignupQuestionWithDefaults instantiates a new SignupQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignupQuestion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignupQuestion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignupQuestion) SetId(v int32)`

SetId sets Id field to given value.


### GetPosition

`func (o *SignupQuestion) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SignupQuestion) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SignupQuestion) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetFieldType

`func (o *SignupQuestion) GetFieldType() SignupQuestionFieldType`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *SignupQuestion) GetFieldTypeOk() (*SignupQuestionFieldType, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *SignupQuestion) SetFieldType(v SignupQuestionFieldType)`

SetFieldType sets FieldType field to given value.


### GetTitle

`func (o *SignupQuestion) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SignupQuestion) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SignupQuestion) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSlug

`func (o *SignupQuestion) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SignupQuestion) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SignupQuestion) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetRequired

`func (o *SignupQuestion) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SignupQuestion) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SignupQuestion) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetForType

`func (o *SignupQuestion) GetForType() string`

GetForType returns the ForType field if non-nil, zero value otherwise.

### GetForTypeOk

`func (o *SignupQuestion) GetForTypeOk() (*string, bool)`

GetForTypeOk returns a tuple with the ForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForType

`func (o *SignupQuestion) SetForType(v string)`

SetForType sets ForType field to given value.


### GetVisibility

`func (o *SignupQuestion) GetVisibility() []string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SignupQuestion) GetVisibilityOk() (*[]string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SignupQuestion) SetVisibility(v []string)`

SetVisibility sets Visibility field to given value.


### GetForStudent

`func (o *SignupQuestion) GetForStudent() bool`

GetForStudent returns the ForStudent field if non-nil, zero value otherwise.

### GetForStudentOk

`func (o *SignupQuestion) GetForStudentOk() (*bool, bool)`

GetForStudentOk returns a tuple with the ForStudent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForStudent

`func (o *SignupQuestion) SetForStudent(v bool)`

SetForStudent sets ForStudent field to given value.


### GetForCustomer

`func (o *SignupQuestion) GetForCustomer() bool`

GetForCustomer returns the ForCustomer field if non-nil, zero value otherwise.

### GetForCustomerOk

`func (o *SignupQuestion) GetForCustomerOk() (*bool, bool)`

GetForCustomerOk returns a tuple with the ForCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForCustomer

`func (o *SignupQuestion) SetForCustomer(v bool)`

SetForCustomer sets ForCustomer field to given value.


### GetChoices

`func (o *SignupQuestion) GetChoices() []string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *SignupQuestion) GetChoicesOk() (*[]string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *SignupQuestion) SetChoices(v []string)`

SetChoices sets Choices field to given value.


### GetSystemName

`func (o *SignupQuestion) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *SignupQuestion) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *SignupQuestion) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.


### SetSystemNameNil

`func (o *SignupQuestion) SetSystemNameNil(b bool)`

 SetSystemNameNil sets the value for SystemName to be an explicit nil

### UnsetSystemName
`func (o *SignupQuestion) UnsetSystemName()`

UnsetSystemName ensures that no value is present for SystemName, not even an explicit nil
### GetUseAsDuplicateIndicator

`func (o *SignupQuestion) GetUseAsDuplicateIndicator() bool`

GetUseAsDuplicateIndicator returns the UseAsDuplicateIndicator field if non-nil, zero value otherwise.

### GetUseAsDuplicateIndicatorOk

`func (o *SignupQuestion) GetUseAsDuplicateIndicatorOk() (*bool, bool)`

GetUseAsDuplicateIndicatorOk returns a tuple with the UseAsDuplicateIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAsDuplicateIndicator

`func (o *SignupQuestion) SetUseAsDuplicateIndicator(v bool)`

SetUseAsDuplicateIndicator sets UseAsDuplicateIndicator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


