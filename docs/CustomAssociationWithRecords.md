# CustomAssociationWithRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociableType** | **string** | The type of the associated system model in slug format. | [readonly] 
**CustomAssociationId** | **int32** | Unique identifier for the custom association. | [readonly] 
**CustomObjectId** | Pointer to **int32** | Unique identifier for the custom object. | [optional] [readonly] 
**CustomRecordIds** | **[]int32** | List of custom record identifiers linked. | [readonly] 

## Methods

### NewCustomAssociationWithRecords

`func NewCustomAssociationWithRecords(associableType string, customAssociationId int32, customRecordIds []int32, ) *CustomAssociationWithRecords`

NewCustomAssociationWithRecords instantiates a new CustomAssociationWithRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAssociationWithRecordsWithDefaults

`func NewCustomAssociationWithRecordsWithDefaults() *CustomAssociationWithRecords`

NewCustomAssociationWithRecordsWithDefaults instantiates a new CustomAssociationWithRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociableType

`func (o *CustomAssociationWithRecords) GetAssociableType() string`

GetAssociableType returns the AssociableType field if non-nil, zero value otherwise.

### GetAssociableTypeOk

`func (o *CustomAssociationWithRecords) GetAssociableTypeOk() (*string, bool)`

GetAssociableTypeOk returns a tuple with the AssociableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociableType

`func (o *CustomAssociationWithRecords) SetAssociableType(v string)`

SetAssociableType sets AssociableType field to given value.


### GetCustomAssociationId

`func (o *CustomAssociationWithRecords) GetCustomAssociationId() int32`

GetCustomAssociationId returns the CustomAssociationId field if non-nil, zero value otherwise.

### GetCustomAssociationIdOk

`func (o *CustomAssociationWithRecords) GetCustomAssociationIdOk() (*int32, bool)`

GetCustomAssociationIdOk returns a tuple with the CustomAssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssociationId

`func (o *CustomAssociationWithRecords) SetCustomAssociationId(v int32)`

SetCustomAssociationId sets CustomAssociationId field to given value.


### GetCustomObjectId

`func (o *CustomAssociationWithRecords) GetCustomObjectId() int32`

GetCustomObjectId returns the CustomObjectId field if non-nil, zero value otherwise.

### GetCustomObjectIdOk

`func (o *CustomAssociationWithRecords) GetCustomObjectIdOk() (*int32, bool)`

GetCustomObjectIdOk returns a tuple with the CustomObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomObjectId

`func (o *CustomAssociationWithRecords) SetCustomObjectId(v int32)`

SetCustomObjectId sets CustomObjectId field to given value.

### HasCustomObjectId

`func (o *CustomAssociationWithRecords) HasCustomObjectId() bool`

HasCustomObjectId returns a boolean if a field has been set.

### GetCustomRecordIds

`func (o *CustomAssociationWithRecords) GetCustomRecordIds() []int32`

GetCustomRecordIds returns the CustomRecordIds field if non-nil, zero value otherwise.

### GetCustomRecordIdsOk

`func (o *CustomAssociationWithRecords) GetCustomRecordIdsOk() (*[]int32, bool)`

GetCustomRecordIdsOk returns a tuple with the CustomRecordIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRecordIds

`func (o *CustomAssociationWithRecords) SetCustomRecordIds(v []int32)`

SetCustomRecordIds sets CustomRecordIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


