# CreateAffiliationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Unique identifier of the associated user | 
**AccountId** | **int32** | Unique identifier of the associated account | 
**KeyContact** | Pointer to **bool** | Boolean indicating if this user is a key contact of the account. | [optional] [default to false]

## Methods

### NewCreateAffiliationRequest

`func NewCreateAffiliationRequest(userId int32, accountId int32, ) *CreateAffiliationRequest`

NewCreateAffiliationRequest instantiates a new CreateAffiliationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAffiliationRequestWithDefaults

`func NewCreateAffiliationRequestWithDefaults() *CreateAffiliationRequest`

NewCreateAffiliationRequestWithDefaults instantiates a new CreateAffiliationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateAffiliationRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateAffiliationRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateAffiliationRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetAccountId

`func (o *CreateAffiliationRequest) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateAffiliationRequest) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateAffiliationRequest) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetKeyContact

`func (o *CreateAffiliationRequest) GetKeyContact() bool`

GetKeyContact returns the KeyContact field if non-nil, zero value otherwise.

### GetKeyContactOk

`func (o *CreateAffiliationRequest) GetKeyContactOk() (*bool, bool)`

GetKeyContactOk returns a tuple with the KeyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyContact

`func (o *CreateAffiliationRequest) SetKeyContact(v bool)`

SetKeyContact sets KeyContact field to given value.

### HasKeyContact

`func (o *CreateAffiliationRequest) HasKeyContact() bool`

HasKeyContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


