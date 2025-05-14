# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the payment method. | [readonly] 
**Name** | **string** | Human readable name of the payment method. | 
**Gateway** | **NullableString** | The type of the integration. | 
**PluginId** | **NullableInt32** | Unique identifier of the optionally linked plugin. | 

## Methods

### NewPaymentMethod

`func NewPaymentMethod(id int32, name string, gateway NullableString, pluginId NullableInt32, ) *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PaymentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethod) SetName(v string)`

SetName sets Name field to given value.


### GetGateway

`func (o *PaymentMethod) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PaymentMethod) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PaymentMethod) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### SetGatewayNil

`func (o *PaymentMethod) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *PaymentMethod) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetPluginId

`func (o *PaymentMethod) GetPluginId() int32`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *PaymentMethod) GetPluginIdOk() (*int32, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *PaymentMethod) SetPluginId(v int32)`

SetPluginId sets PluginId field to given value.


### SetPluginIdNil

`func (o *PaymentMethod) SetPluginIdNil(b bool)`

 SetPluginIdNil sets the value for PluginId to be an explicit nil

### UnsetPluginId
`func (o *PaymentMethod) UnsetPluginId()`

UnsetPluginId ensures that no value is present for PluginId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


