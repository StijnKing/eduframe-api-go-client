# PaymentOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the payment option. | [readonly] 
**Name** | **string** | Human readable name of the payment option. | 
**AvailableFrom** | **float32** | Minimum required price before this option is allowed. | 
**ExtraCost** | **string** | Extra cost on the invoice if this option is used. | 
**Percentage** | **NullableFloat32** | If the type of payment option is with a percentual deposit. | 
**Multiplier** | **NullableString** | The multiplier to get the total cost. | 

## Methods

### NewPaymentOption

`func NewPaymentOption(id int32, name string, availableFrom float32, extraCost string, percentage NullableFloat32, multiplier NullableString, ) *PaymentOption`

NewPaymentOption instantiates a new PaymentOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOptionWithDefaults

`func NewPaymentOptionWithDefaults() *PaymentOption`

NewPaymentOptionWithDefaults instantiates a new PaymentOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentOption) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentOption) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentOption) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PaymentOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentOption) SetName(v string)`

SetName sets Name field to given value.


### GetAvailableFrom

`func (o *PaymentOption) GetAvailableFrom() float32`

GetAvailableFrom returns the AvailableFrom field if non-nil, zero value otherwise.

### GetAvailableFromOk

`func (o *PaymentOption) GetAvailableFromOk() (*float32, bool)`

GetAvailableFromOk returns a tuple with the AvailableFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFrom

`func (o *PaymentOption) SetAvailableFrom(v float32)`

SetAvailableFrom sets AvailableFrom field to given value.


### GetExtraCost

`func (o *PaymentOption) GetExtraCost() string`

GetExtraCost returns the ExtraCost field if non-nil, zero value otherwise.

### GetExtraCostOk

`func (o *PaymentOption) GetExtraCostOk() (*string, bool)`

GetExtraCostOk returns a tuple with the ExtraCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraCost

`func (o *PaymentOption) SetExtraCost(v string)`

SetExtraCost sets ExtraCost field to given value.


### GetPercentage

`func (o *PaymentOption) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PaymentOption) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PaymentOption) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### SetPercentageNil

`func (o *PaymentOption) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *PaymentOption) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil
### GetMultiplier

`func (o *PaymentOption) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *PaymentOption) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *PaymentOption) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.


### SetMultiplierNil

`func (o *PaymentOption) SetMultiplierNil(b bool)`

 SetMultiplierNil sets the value for Multiplier to be an explicit nil

### UnsetMultiplier
`func (o *PaymentOption) UnsetMultiplier()`

UnsetMultiplier ensures that no value is present for Multiplier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


