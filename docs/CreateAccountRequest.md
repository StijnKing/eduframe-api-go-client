# CreateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Arbitrary string representing the name of the account. Is autogenerated for personal accounts. | 
**Email** | Pointer to **NullableString** | A string representing the billing e-mail of the account | [optional] 
**Phone** | Pointer to **NullableString** | A string representing the phone number of the account | [optional] 
**LabelIds** | Pointer to **[]int32** | IDs of the labels | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the account. | [optional] 
**AddressAttributes** | Pointer to [**NullableAddressPayload**](AddressPayload.md) |  | [optional] 
**SignupAnswersAttributes** | Pointer to [**[]CreateAccountRequestSignupAnswersAttributesInner**](CreateAccountRequestSignupAnswersAttributesInner.md) |  | [optional] 

## Methods

### NewCreateAccountRequest

`func NewCreateAccountRequest(name string, ) *CreateAccountRequest`

NewCreateAccountRequest instantiates a new CreateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountRequestWithDefaults

`func NewCreateAccountRequestWithDefaults() *CreateAccountRequest`

NewCreateAccountRequestWithDefaults instantiates a new CreateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccountRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CreateAccountRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateAccountRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateAccountRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateAccountRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CreateAccountRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateAccountRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *CreateAccountRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateAccountRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateAccountRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CreateAccountRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *CreateAccountRequest) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *CreateAccountRequest) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetLabelIds

`func (o *CreateAccountRequest) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *CreateAccountRequest) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *CreateAccountRequest) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *CreateAccountRequest) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetCustom

`func (o *CreateAccountRequest) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CreateAccountRequest) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CreateAccountRequest) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *CreateAccountRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetAddressAttributes

`func (o *CreateAccountRequest) GetAddressAttributes() AddressPayload`

GetAddressAttributes returns the AddressAttributes field if non-nil, zero value otherwise.

### GetAddressAttributesOk

`func (o *CreateAccountRequest) GetAddressAttributesOk() (*AddressPayload, bool)`

GetAddressAttributesOk returns a tuple with the AddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributes

`func (o *CreateAccountRequest) SetAddressAttributes(v AddressPayload)`

SetAddressAttributes sets AddressAttributes field to given value.

### HasAddressAttributes

`func (o *CreateAccountRequest) HasAddressAttributes() bool`

HasAddressAttributes returns a boolean if a field has been set.

### SetAddressAttributesNil

`func (o *CreateAccountRequest) SetAddressAttributesNil(b bool)`

 SetAddressAttributesNil sets the value for AddressAttributes to be an explicit nil

### UnsetAddressAttributes
`func (o *CreateAccountRequest) UnsetAddressAttributes()`

UnsetAddressAttributes ensures that no value is present for AddressAttributes, not even an explicit nil
### GetSignupAnswersAttributes

`func (o *CreateAccountRequest) GetSignupAnswersAttributes() []CreateAccountRequestSignupAnswersAttributesInner`

GetSignupAnswersAttributes returns the SignupAnswersAttributes field if non-nil, zero value otherwise.

### GetSignupAnswersAttributesOk

`func (o *CreateAccountRequest) GetSignupAnswersAttributesOk() (*[]CreateAccountRequestSignupAnswersAttributesInner, bool)`

GetSignupAnswersAttributesOk returns a tuple with the SignupAnswersAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupAnswersAttributes

`func (o *CreateAccountRequest) SetSignupAnswersAttributes(v []CreateAccountRequestSignupAnswersAttributesInner)`

SetSignupAnswersAttributes sets SignupAnswersAttributes field to given value.

### HasSignupAnswersAttributes

`func (o *CreateAccountRequest) HasSignupAnswersAttributes() bool`

HasSignupAnswersAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


