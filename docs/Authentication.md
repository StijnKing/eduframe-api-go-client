# Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the authentication. | [readonly] 
**UserId** | **int32** | Identifier of the associated User. | 
**AuthenticationProviderId** | **int32** | Unique identifier of the authentication provider | 
**Uid** | **string** | Login identifier. | 
**OtpEnabled** | **bool** | Whether or not the user has enabled two factor authentication. | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewAuthentication

`func NewAuthentication(id int32, userId int32, authenticationProviderId int32, uid string, otpEnabled bool, updatedAt time.Time, createdAt time.Time, ) *Authentication`

NewAuthentication instantiates a new Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationWithDefaults

`func NewAuthenticationWithDefaults() *Authentication`

NewAuthenticationWithDefaults instantiates a new Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Authentication) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Authentication) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Authentication) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *Authentication) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Authentication) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Authentication) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetAuthenticationProviderId

`func (o *Authentication) GetAuthenticationProviderId() int32`

GetAuthenticationProviderId returns the AuthenticationProviderId field if non-nil, zero value otherwise.

### GetAuthenticationProviderIdOk

`func (o *Authentication) GetAuthenticationProviderIdOk() (*int32, bool)`

GetAuthenticationProviderIdOk returns a tuple with the AuthenticationProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProviderId

`func (o *Authentication) SetAuthenticationProviderId(v int32)`

SetAuthenticationProviderId sets AuthenticationProviderId field to given value.


### GetUid

`func (o *Authentication) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Authentication) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Authentication) SetUid(v string)`

SetUid sets Uid field to given value.


### GetOtpEnabled

`func (o *Authentication) GetOtpEnabled() bool`

GetOtpEnabled returns the OtpEnabled field if non-nil, zero value otherwise.

### GetOtpEnabledOk

`func (o *Authentication) GetOtpEnabledOk() (*bool, bool)`

GetOtpEnabledOk returns a tuple with the OtpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpEnabled

`func (o *Authentication) SetOtpEnabled(v bool)`

SetOtpEnabled sets OtpEnabled field to given value.


### GetUpdatedAt

`func (o *Authentication) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Authentication) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Authentication) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Authentication) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Authentication) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Authentication) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


