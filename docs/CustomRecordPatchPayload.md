# CustomRecordPatchPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the custom record is active. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the custom record. | [optional] 
**Properties** | Pointer to [**map[string]CustomFieldValue**](CustomFieldValue.md) | The JSON properties of the custom record. | [optional] 

## Methods

### NewCustomRecordPatchPayload

`func NewCustomRecordPatchPayload() *CustomRecordPatchPayload`

NewCustomRecordPatchPayload instantiates a new CustomRecordPatchPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRecordPatchPayloadWithDefaults

`func NewCustomRecordPatchPayloadWithDefaults() *CustomRecordPatchPayload`

NewCustomRecordPatchPayloadWithDefaults instantiates a new CustomRecordPatchPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CustomRecordPatchPayload) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CustomRecordPatchPayload) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CustomRecordPatchPayload) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CustomRecordPatchPayload) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDisplayName

`func (o *CustomRecordPatchPayload) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomRecordPatchPayload) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomRecordPatchPayload) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CustomRecordPatchPayload) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProperties

`func (o *CustomRecordPatchPayload) GetProperties() map[string]CustomFieldValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CustomRecordPatchPayload) GetPropertiesOk() (*map[string]CustomFieldValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CustomRecordPatchPayload) SetProperties(v map[string]CustomFieldValue)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CustomRecordPatchPayload) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


