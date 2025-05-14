# DiscountCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the discount code. | 
**Name** | **string** | The name of the discount code. | 
**DiscountCode** | **string** | The discount code. | 
**DiscountType** | **string** | The type of discount, which can be a fixed amount or a percentage. | 
**Amount** | **float32** | The amount of the discount. | 
**MaxUsage** | **NullableInt32** | The maximum number of times the discount code can be used. | 
**UsageCount** | **int32** | The number of times the discount code has been used. | 
**StartDate** | **string** | The start date of the discount code. | 
**ExpirationDate** | **string** | The expiration date of the discount code. | 

## Methods

### NewDiscountCode

`func NewDiscountCode(id int32, name string, discountCode string, discountType string, amount float32, maxUsage NullableInt32, usageCount int32, startDate string, expirationDate string, ) *DiscountCode`

NewDiscountCode instantiates a new DiscountCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountCodeWithDefaults

`func NewDiscountCodeWithDefaults() *DiscountCode`

NewDiscountCodeWithDefaults instantiates a new DiscountCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscountCode) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscountCode) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscountCode) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *DiscountCode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscountCode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscountCode) SetName(v string)`

SetName sets Name field to given value.


### GetDiscountCode

`func (o *DiscountCode) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *DiscountCode) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *DiscountCode) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.


### GetDiscountType

`func (o *DiscountCode) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *DiscountCode) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *DiscountCode) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.


### GetAmount

`func (o *DiscountCode) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DiscountCode) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DiscountCode) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetMaxUsage

`func (o *DiscountCode) GetMaxUsage() int32`

GetMaxUsage returns the MaxUsage field if non-nil, zero value otherwise.

### GetMaxUsageOk

`func (o *DiscountCode) GetMaxUsageOk() (*int32, bool)`

GetMaxUsageOk returns a tuple with the MaxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsage

`func (o *DiscountCode) SetMaxUsage(v int32)`

SetMaxUsage sets MaxUsage field to given value.


### SetMaxUsageNil

`func (o *DiscountCode) SetMaxUsageNil(b bool)`

 SetMaxUsageNil sets the value for MaxUsage to be an explicit nil

### UnsetMaxUsage
`func (o *DiscountCode) UnsetMaxUsage()`

UnsetMaxUsage ensures that no value is present for MaxUsage, not even an explicit nil
### GetUsageCount

`func (o *DiscountCode) GetUsageCount() int32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *DiscountCode) GetUsageCountOk() (*int32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *DiscountCode) SetUsageCount(v int32)`

SetUsageCount sets UsageCount field to given value.


### GetStartDate

`func (o *DiscountCode) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DiscountCode) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DiscountCode) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetExpirationDate

`func (o *DiscountCode) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *DiscountCode) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *DiscountCode) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


