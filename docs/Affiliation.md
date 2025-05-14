# Affiliation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the affiliation record. | [readonly] 
**AccountId** | **int32** | Unique identifier of the associated account | 
**UserId** | **int32** | Unique identifier of the associated user | 
**KeyContact** | **bool** | Boolean indicating if this user is a key contact of the account. | [default to false]
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewAffiliation

`func NewAffiliation(id int32, accountId int32, userId int32, keyContact bool, updatedAt time.Time, createdAt time.Time, ) *Affiliation`

NewAffiliation instantiates a new Affiliation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliationWithDefaults

`func NewAffiliationWithDefaults() *Affiliation`

NewAffiliationWithDefaults instantiates a new Affiliation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Affiliation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Affiliation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Affiliation) SetId(v int32)`

SetId sets Id field to given value.


### GetAccountId

`func (o *Affiliation) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Affiliation) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Affiliation) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetUserId

`func (o *Affiliation) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Affiliation) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Affiliation) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetKeyContact

`func (o *Affiliation) GetKeyContact() bool`

GetKeyContact returns the KeyContact field if non-nil, zero value otherwise.

### GetKeyContactOk

`func (o *Affiliation) GetKeyContactOk() (*bool, bool)`

GetKeyContactOk returns a tuple with the KeyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyContact

`func (o *Affiliation) SetKeyContact(v bool)`

SetKeyContact sets KeyContact field to given value.


### GetUpdatedAt

`func (o *Affiliation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Affiliation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Affiliation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Affiliation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Affiliation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Affiliation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


