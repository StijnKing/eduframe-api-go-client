# UpdateAffiliationByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyContact** | Pointer to **bool** | Boolean indicating if this user is a key contact of the account. | [optional] 
**UserId** | Pointer to **int32** | Unique identifier of the associated user | [optional] 
**AccountId** | Pointer to **int32** | Unique identifier of the associated account | [optional] 

## Methods

### NewUpdateAffiliationByIdRequest

`func NewUpdateAffiliationByIdRequest() *UpdateAffiliationByIdRequest`

NewUpdateAffiliationByIdRequest instantiates a new UpdateAffiliationByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAffiliationByIdRequestWithDefaults

`func NewUpdateAffiliationByIdRequestWithDefaults() *UpdateAffiliationByIdRequest`

NewUpdateAffiliationByIdRequestWithDefaults instantiates a new UpdateAffiliationByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyContact

`func (o *UpdateAffiliationByIdRequest) GetKeyContact() bool`

GetKeyContact returns the KeyContact field if non-nil, zero value otherwise.

### GetKeyContactOk

`func (o *UpdateAffiliationByIdRequest) GetKeyContactOk() (*bool, bool)`

GetKeyContactOk returns a tuple with the KeyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyContact

`func (o *UpdateAffiliationByIdRequest) SetKeyContact(v bool)`

SetKeyContact sets KeyContact field to given value.

### HasKeyContact

`func (o *UpdateAffiliationByIdRequest) HasKeyContact() bool`

HasKeyContact returns a boolean if a field has been set.

### GetUserId

`func (o *UpdateAffiliationByIdRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateAffiliationByIdRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateAffiliationByIdRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UpdateAffiliationByIdRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateAffiliationByIdRequest) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateAffiliationByIdRequest) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateAffiliationByIdRequest) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateAffiliationByIdRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


