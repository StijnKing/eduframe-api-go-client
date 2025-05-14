# CustomFieldOptionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Enabled** | **bool** | Whether the option can be chosen or not | 

## Methods

### NewCustomFieldOptionPayload

`func NewCustomFieldOptionPayload(value string, enabled bool, ) *CustomFieldOptionPayload`

NewCustomFieldOptionPayload instantiates a new CustomFieldOptionPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldOptionPayloadWithDefaults

`func NewCustomFieldOptionPayloadWithDefaults() *CustomFieldOptionPayload`

NewCustomFieldOptionPayloadWithDefaults instantiates a new CustomFieldOptionPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CustomFieldOptionPayload) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldOptionPayload) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldOptionPayload) SetValue(v string)`

SetValue sets Value field to given value.


### GetEnabled

`func (o *CustomFieldOptionPayload) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomFieldOptionPayload) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomFieldOptionPayload) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


