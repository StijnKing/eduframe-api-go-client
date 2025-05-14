# CreateAuthenticationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Login identifier. | 
**UserId** | **int32** | Identifier of the associated User. | 
**AuthenticationProviderType** | [**AuthenticationProviderType**](AuthenticationProviderType.md) |  | 

## Methods

### NewCreateAuthenticationRequest

`func NewCreateAuthenticationRequest(uid string, userId int32, authenticationProviderType AuthenticationProviderType, ) *CreateAuthenticationRequest`

NewCreateAuthenticationRequest instantiates a new CreateAuthenticationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthenticationRequestWithDefaults

`func NewCreateAuthenticationRequestWithDefaults() *CreateAuthenticationRequest`

NewCreateAuthenticationRequestWithDefaults instantiates a new CreateAuthenticationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *CreateAuthenticationRequest) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CreateAuthenticationRequest) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CreateAuthenticationRequest) SetUid(v string)`

SetUid sets Uid field to given value.


### GetUserId

`func (o *CreateAuthenticationRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateAuthenticationRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateAuthenticationRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetAuthenticationProviderType

`func (o *CreateAuthenticationRequest) GetAuthenticationProviderType() AuthenticationProviderType`

GetAuthenticationProviderType returns the AuthenticationProviderType field if non-nil, zero value otherwise.

### GetAuthenticationProviderTypeOk

`func (o *CreateAuthenticationRequest) GetAuthenticationProviderTypeOk() (*AuthenticationProviderType, bool)`

GetAuthenticationProviderTypeOk returns a tuple with the AuthenticationProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProviderType

`func (o *CreateAuthenticationRequest) SetAuthenticationProviderType(v AuthenticationProviderType)`

SetAuthenticationProviderType sets AuthenticationProviderType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


