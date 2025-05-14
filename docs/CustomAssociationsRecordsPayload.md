# CustomAssociationsRecordsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomAssociationId** | **int32** | The unique identifier of the association. | 
**CustomRecordIds** | **[]int32** | List of custom record identifiers linked. | 

## Methods

### NewCustomAssociationsRecordsPayload

`func NewCustomAssociationsRecordsPayload(customAssociationId int32, customRecordIds []int32, ) *CustomAssociationsRecordsPayload`

NewCustomAssociationsRecordsPayload instantiates a new CustomAssociationsRecordsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAssociationsRecordsPayloadWithDefaults

`func NewCustomAssociationsRecordsPayloadWithDefaults() *CustomAssociationsRecordsPayload`

NewCustomAssociationsRecordsPayloadWithDefaults instantiates a new CustomAssociationsRecordsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomAssociationId

`func (o *CustomAssociationsRecordsPayload) GetCustomAssociationId() int32`

GetCustomAssociationId returns the CustomAssociationId field if non-nil, zero value otherwise.

### GetCustomAssociationIdOk

`func (o *CustomAssociationsRecordsPayload) GetCustomAssociationIdOk() (*int32, bool)`

GetCustomAssociationIdOk returns a tuple with the CustomAssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssociationId

`func (o *CustomAssociationsRecordsPayload) SetCustomAssociationId(v int32)`

SetCustomAssociationId sets CustomAssociationId field to given value.


### GetCustomRecordIds

`func (o *CustomAssociationsRecordsPayload) GetCustomRecordIds() []int32`

GetCustomRecordIds returns the CustomRecordIds field if non-nil, zero value otherwise.

### GetCustomRecordIdsOk

`func (o *CustomAssociationsRecordsPayload) GetCustomRecordIdsOk() (*[]int32, bool)`

GetCustomRecordIdsOk returns a tuple with the CustomRecordIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRecordIds

`func (o *CustomAssociationsRecordsPayload) SetCustomRecordIds(v []int32)`

SetCustomRecordIds sets CustomRecordIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


