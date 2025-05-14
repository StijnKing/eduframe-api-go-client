# VariantWithIncludes

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
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the product. | [optional] 
**CustomAssociations** | Pointer to [**[]CustomAssociationWithRecords**](CustomAssociationWithRecords.md) | The custom associations linked to the catalog variant.  ![Custom Objects](https://img.shields.io/badge/Feature-Custom_objects-1d8127) ![Beta](https://img.shields.io/badge/Beta-7d15a3)  | [optional] 

## Methods

### NewVariantWithIncludes

`func NewVariantWithIncludes(id int32, productId int32, name string, costScheme CostScheme, cost NullableString, currency Currency, variantableType VariantableType, variantableId int32, availability string, availablePlaces int32, isPublished bool, showOnWebsite bool, updatedAt time.Time, createdAt time.Time, ) *VariantWithIncludes`

NewVariantWithIncludes instantiates a new VariantWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantWithIncludesWithDefaults

`func NewVariantWithIncludesWithDefaults() *VariantWithIncludes`

NewVariantWithIncludesWithDefaults instantiates a new VariantWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VariantWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariantWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariantWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetProductId

`func (o *VariantWithIncludes) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *VariantWithIncludes) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *VariantWithIncludes) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetName

`func (o *VariantWithIncludes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariantWithIncludes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariantWithIncludes) SetName(v string)`

SetName sets Name field to given value.


### GetSku

`func (o *VariantWithIncludes) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *VariantWithIncludes) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *VariantWithIncludes) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *VariantWithIncludes) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *VariantWithIncludes) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *VariantWithIncludes) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetCostScheme

`func (o *VariantWithIncludes) GetCostScheme() CostScheme`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *VariantWithIncludes) GetCostSchemeOk() (*CostScheme, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *VariantWithIncludes) SetCostScheme(v CostScheme)`

SetCostScheme sets CostScheme field to given value.


### GetCost

`func (o *VariantWithIncludes) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *VariantWithIncludes) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *VariantWithIncludes) SetCost(v string)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *VariantWithIncludes) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *VariantWithIncludes) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCurrency

`func (o *VariantWithIncludes) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VariantWithIncludes) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VariantWithIncludes) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetVariantableType

`func (o *VariantWithIncludes) GetVariantableType() VariantableType`

GetVariantableType returns the VariantableType field if non-nil, zero value otherwise.

### GetVariantableTypeOk

`func (o *VariantWithIncludes) GetVariantableTypeOk() (*VariantableType, bool)`

GetVariantableTypeOk returns a tuple with the VariantableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantableType

`func (o *VariantWithIncludes) SetVariantableType(v VariantableType)`

SetVariantableType sets VariantableType field to given value.


### GetVariantableId

`func (o *VariantWithIncludes) GetVariantableId() int32`

GetVariantableId returns the VariantableId field if non-nil, zero value otherwise.

### GetVariantableIdOk

`func (o *VariantWithIncludes) GetVariantableIdOk() (*int32, bool)`

GetVariantableIdOk returns a tuple with the VariantableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantableId

`func (o *VariantWithIncludes) SetVariantableId(v int32)`

SetVariantableId sets VariantableId field to given value.


### GetAvailability

`func (o *VariantWithIncludes) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *VariantWithIncludes) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *VariantWithIncludes) SetAvailability(v string)`

SetAvailability sets Availability field to given value.


### GetAvailablePlaces

`func (o *VariantWithIncludes) GetAvailablePlaces() int32`

GetAvailablePlaces returns the AvailablePlaces field if non-nil, zero value otherwise.

### GetAvailablePlacesOk

`func (o *VariantWithIncludes) GetAvailablePlacesOk() (*int32, bool)`

GetAvailablePlacesOk returns a tuple with the AvailablePlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePlaces

`func (o *VariantWithIncludes) SetAvailablePlaces(v int32)`

SetAvailablePlaces sets AvailablePlaces field to given value.


### GetIsPublished

`func (o *VariantWithIncludes) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *VariantWithIncludes) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *VariantWithIncludes) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.


### GetShowOnWebsite

`func (o *VariantWithIncludes) GetShowOnWebsite() bool`

GetShowOnWebsite returns the ShowOnWebsite field if non-nil, zero value otherwise.

### GetShowOnWebsiteOk

`func (o *VariantWithIncludes) GetShowOnWebsiteOk() (*bool, bool)`

GetShowOnWebsiteOk returns a tuple with the ShowOnWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnWebsite

`func (o *VariantWithIncludes) SetShowOnWebsite(v bool)`

SetShowOnWebsite sets ShowOnWebsite field to given value.


### GetUpdatedAt

`func (o *VariantWithIncludes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VariantWithIncludes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VariantWithIncludes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *VariantWithIncludes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VariantWithIncludes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VariantWithIncludes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCustom

`func (o *VariantWithIncludes) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *VariantWithIncludes) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *VariantWithIncludes) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *VariantWithIncludes) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomAssociations

`func (o *VariantWithIncludes) GetCustomAssociations() []CustomAssociationWithRecords`

GetCustomAssociations returns the CustomAssociations field if non-nil, zero value otherwise.

### GetCustomAssociationsOk

`func (o *VariantWithIncludes) GetCustomAssociationsOk() (*[]CustomAssociationWithRecords, bool)`

GetCustomAssociationsOk returns a tuple with the CustomAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssociations

`func (o *VariantWithIncludes) SetCustomAssociations(v []CustomAssociationWithRecords)`

SetCustomAssociations sets CustomAssociations field to given value.

### HasCustomAssociations

`func (o *VariantWithIncludes) HasCustomAssociations() bool`

HasCustomAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


