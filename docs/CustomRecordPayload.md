# CustomRecordPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the custom record is active. | [optional] [default to true]
**DisplayName** | **string** | The display name of the custom record. | 
**Properties** | [**map[string]CustomFieldValue**](CustomFieldValue.md) | The JSON properties of the custom record. | 

## Methods

### NewCustomRecordPayload

`func NewCustomRecordPayload(displayName string, properties map[string]CustomFieldValue, ) *CustomRecordPayload`

NewCustomRecordPayload instantiates a new CustomRecordPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRecordPayloadWithDefaults

`func NewCustomRecordPayloadWithDefaults() *CustomRecordPayload`

NewCustomRecordPayloadWithDefaults instantiates a new CustomRecordPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CustomRecordPayload) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CustomRecordPayload) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CustomRecordPayload) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CustomRecordPayload) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDisplayName

`func (o *CustomRecordPayload) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomRecordPayload) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomRecordPayload) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetProperties

`func (o *CustomRecordPayload) GetProperties() map[string]CustomFieldValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CustomRecordPayload) GetPropertiesOk() (*map[string]CustomFieldValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CustomRecordPayload) SetProperties(v map[string]CustomFieldValue)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


