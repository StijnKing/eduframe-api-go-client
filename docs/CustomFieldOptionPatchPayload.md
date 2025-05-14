# CustomFieldOptionPatchPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Whether the option can be chosen or not | [optional] 

## Methods

### NewCustomFieldOptionPatchPayload

`func NewCustomFieldOptionPatchPayload() *CustomFieldOptionPatchPayload`

NewCustomFieldOptionPatchPayload instantiates a new CustomFieldOptionPatchPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldOptionPatchPayloadWithDefaults

`func NewCustomFieldOptionPatchPayloadWithDefaults() *CustomFieldOptionPatchPayload`

NewCustomFieldOptionPatchPayloadWithDefaults instantiates a new CustomFieldOptionPatchPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CustomFieldOptionPatchPayload) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldOptionPatchPayload) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldOptionPatchPayload) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldOptionPatchPayload) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEnabled

`func (o *CustomFieldOptionPatchPayload) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomFieldOptionPatchPayload) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomFieldOptionPatchPayload) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustomFieldOptionPatchPayload) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


