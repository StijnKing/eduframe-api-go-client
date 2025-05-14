# LeadProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogProductId** | **int32** | ID of the catalog product | 
**CatalogVariantId** | **NullableInt32** | ID of the catalog variant | 

## Methods

### NewLeadProduct

`func NewLeadProduct(catalogProductId int32, catalogVariantId NullableInt32, ) *LeadProduct`

NewLeadProduct instantiates a new LeadProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeadProductWithDefaults

`func NewLeadProductWithDefaults() *LeadProduct`

NewLeadProductWithDefaults instantiates a new LeadProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogProductId

`func (o *LeadProduct) GetCatalogProductId() int32`

GetCatalogProductId returns the CatalogProductId field if non-nil, zero value otherwise.

### GetCatalogProductIdOk

`func (o *LeadProduct) GetCatalogProductIdOk() (*int32, bool)`

GetCatalogProductIdOk returns a tuple with the CatalogProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogProductId

`func (o *LeadProduct) SetCatalogProductId(v int32)`

SetCatalogProductId sets CatalogProductId field to given value.


### GetCatalogVariantId

`func (o *LeadProduct) GetCatalogVariantId() int32`

GetCatalogVariantId returns the CatalogVariantId field if non-nil, zero value otherwise.

### GetCatalogVariantIdOk

`func (o *LeadProduct) GetCatalogVariantIdOk() (*int32, bool)`

GetCatalogVariantIdOk returns a tuple with the CatalogVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogVariantId

`func (o *LeadProduct) SetCatalogVariantId(v int32)`

SetCatalogVariantId sets CatalogVariantId field to given value.


### SetCatalogVariantIdNil

`func (o *LeadProduct) SetCatalogVariantIdNil(b bool)`

 SetCatalogVariantIdNil sets the value for CatalogVariantId to be an explicit nil

### UnsetCatalogVariantId
`func (o *LeadProduct) UnsetCatalogVariantId()`

UnsetCatalogVariantId ensures that no value is present for CatalogVariantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


