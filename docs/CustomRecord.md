# CustomRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the custom record. | 
**Active** | **bool** | Whether the custom record is active. | 
**DisplayName** | **string** | The display name of the custom record. | 
**Properties** | [**map[string]CustomFieldValue**](CustomFieldValue.md) | The JSON properties of the custom record. | 

## Methods

### NewCustomRecord

`func NewCustomRecord(id int32, active bool, displayName string, properties map[string]CustomFieldValue, ) *CustomRecord`

NewCustomRecord instantiates a new CustomRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRecordWithDefaults

`func NewCustomRecordWithDefaults() *CustomRecord`

NewCustomRecordWithDefaults instantiates a new CustomRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomRecord) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomRecord) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomRecord) SetId(v int32)`

SetId sets Id field to given value.


### GetActive

`func (o *CustomRecord) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CustomRecord) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CustomRecord) SetActive(v bool)`

SetActive sets Active field to given value.


### GetDisplayName

`func (o *CustomRecord) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomRecord) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomRecord) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetProperties

`func (o *CustomRecord) GetProperties() map[string]CustomFieldValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CustomRecord) GetPropertiesOk() (*map[string]CustomFieldValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CustomRecord) SetProperties(v map[string]CustomFieldValue)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


