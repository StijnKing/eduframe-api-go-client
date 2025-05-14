# Variant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the catalog variant. | [readonly] 
**ProductId** | **int32** | Unique identifier of the catalog product. | 
**Name** | **string** | Name of the catalog variant. | 
**Sku** | Pointer to **NullableString** | An optional unique identifier of the catalog variant. | [optional] 
**CostScheme** | [**CostScheme**](CostScheme.md) |  | 
**Cost** | **NullableString** | Cost of the variant. | 
**Currency** | [**Currency**](Currency.md) |  | 
**VariantableType** | [**VariantableType**](VariantableType.md) |  | 
**VariantableId** | **int32** | Unique identifier of the variantable. | 
**Availability** | **string** | Availability of the catalog variant. | 
**AvailablePlaces** | **int32** | Number of available places. | 
**IsPublished** | **bool** | Boolean showing if the variant is published or not. | 
**ShowOnWebsite** | **bool** | Boolean showing if the variant is shown on the website or not. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewVariant

`func NewVariant(id int32, productId int32, name string, costScheme CostScheme, cost NullableString, currency Currency, variantableType VariantableType, variantableId int32, availability string, availablePlaces int32, isPublished bool, showOnWebsite bool, updatedAt time.Time, createdAt time.Time, ) *Variant`

NewVariant instantiates a new Variant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantWithDefaults

`func NewVariantWithDefaults() *Variant`

NewVariantWithDefaults instantiates a new Variant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Variant) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Variant) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Variant) SetId(v int32)`

SetId sets Id field to given value.


### GetProductId

`func (o *Variant) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Variant) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Variant) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetName

`func (o *Variant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Variant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Variant) SetName(v string)`

SetName sets Name field to given value.


### GetSku

`func (o *Variant) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Variant) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Variant) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Variant) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *Variant) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *Variant) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetCostScheme

`func (o *Variant) GetCostScheme() CostScheme`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *Variant) GetCostSchemeOk() (*CostScheme, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *Variant) SetCostScheme(v CostScheme)`

SetCostScheme sets CostScheme field to given value.


### GetCost

`func (o *Variant) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Variant) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Variant) SetCost(v string)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *Variant) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *Variant) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCurrency

`func (o *Variant) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Variant) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Variant) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetVariantableType

`func (o *Variant) GetVariantableType() VariantableType`

GetVariantableType returns the VariantableType field if non-nil, zero value otherwise.

### GetVariantableTypeOk

`func (o *Variant) GetVariantableTypeOk() (*VariantableType, bool)`

GetVariantableTypeOk returns a tuple with the VariantableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantableType

`func (o *Variant) SetVariantableType(v VariantableType)`

SetVariantableType sets VariantableType field to given value.


### GetVariantableId

`func (o *Variant) GetVariantableId() int32`

GetVariantableId returns the VariantableId field if non-nil, zero value otherwise.

### GetVariantableIdOk

`func (o *Variant) GetVariantableIdOk() (*int32, bool)`

GetVariantableIdOk returns a tuple with the VariantableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantableId

`func (o *Variant) SetVariantableId(v int32)`

SetVariantableId sets VariantableId field to given value.


### GetAvailability

`func (o *Variant) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *Variant) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *Variant) SetAvailability(v string)`

SetAvailability sets Availability field to given value.


### GetAvailablePlaces

`func (o *Variant) GetAvailablePlaces() int32`

GetAvailablePlaces returns the AvailablePlaces field if non-nil, zero value otherwise.

### GetAvailablePlacesOk

`func (o *Variant) GetAvailablePlacesOk() (*int32, bool)`

GetAvailablePlacesOk returns a tuple with the AvailablePlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePlaces

`func (o *Variant) SetAvailablePlaces(v int32)`

SetAvailablePlaces sets AvailablePlaces field to given value.


### GetIsPublished

`func (o *Variant) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *Variant) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *Variant) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.


### GetShowOnWebsite

`func (o *Variant) GetShowOnWebsite() bool`

GetShowOnWebsite returns the ShowOnWebsite field if non-nil, zero value otherwise.

### GetShowOnWebsiteOk

`func (o *Variant) GetShowOnWebsiteOk() (*bool, bool)`

GetShowOnWebsiteOk returns a tuple with the ShowOnWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnWebsite

`func (o *Variant) SetShowOnWebsite(v bool)`

SetShowOnWebsite sets ShowOnWebsite field to given value.


### GetUpdatedAt

`func (o *Variant) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Variant) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Variant) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Variant) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Variant) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Variant) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


