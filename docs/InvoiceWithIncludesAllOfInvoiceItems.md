# InvoiceWithIncludesAllOfInvoiceItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the invoice item. | [readonly] 
**Name** | **string** | The name of the invoice item. | 
**Units** | **int32** | Integer representing the number of units of the invoice item. | 
**UnitPrice** | **string** | Decimal representing the price of an unit. | 
**InvoiceVatId** | **int32** | Identifier of the invoice vat. | 
**CatalogVariantId** | **NullableInt32** | Unique identifier of the catalog variant. | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewInvoiceWithIncludesAllOfInvoiceItems

`func NewInvoiceWithIncludesAllOfInvoiceItems(id int32, name string, units int32, unitPrice string, invoiceVatId int32, catalogVariantId NullableInt32, updatedAt time.Time, createdAt time.Time, ) *InvoiceWithIncludesAllOfInvoiceItems`

NewInvoiceWithIncludesAllOfInvoiceItems instantiates a new InvoiceWithIncludesAllOfInvoiceItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithIncludesAllOfInvoiceItemsWithDefaults

`func NewInvoiceWithIncludesAllOfInvoiceItemsWithDefaults() *InvoiceWithIncludesAllOfInvoiceItems`

NewInvoiceWithIncludesAllOfInvoiceItemsWithDefaults instantiates a new InvoiceWithIncludesAllOfInvoiceItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceWithIncludesAllOfInvoiceItems) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceWithIncludesAllOfInvoiceItems) SetName(v string)`

SetName sets Name field to given value.


### GetUnits

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *InvoiceWithIncludesAllOfInvoiceItems) SetUnits(v int32)`

SetUnits sets Units field to given value.


### GetUnitPrice

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InvoiceWithIncludesAllOfInvoiceItems) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.


### GetInvoiceVatId

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetInvoiceVatId() int32`

GetInvoiceVatId returns the InvoiceVatId field if non-nil, zero value otherwise.

### GetInvoiceVatIdOk

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetInvoiceVatIdOk() (*int32, bool)`

GetInvoiceVatIdOk returns a tuple with the InvoiceVatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceVatId

`func (o *InvoiceWithIncludesAllOfInvoiceItems) SetInvoiceVatId(v int32)`

SetInvoiceVatId sets InvoiceVatId field to given value.


### GetCatalogVariantId

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetCatalogVariantId() int32`

GetCatalogVariantId returns the CatalogVariantId field if non-nil, zero value otherwise.

### GetCatalogVariantIdOk

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetCatalogVariantIdOk() (*int32, bool)`

GetCatalogVariantIdOk returns a tuple with the CatalogVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogVariantId

`func (o *InvoiceWithIncludesAllOfInvoiceItems) SetCatalogVariantId(v int32)`

SetCatalogVariantId sets CatalogVariantId field to given value.


### SetCatalogVariantIdNil

`func (o *InvoiceWithIncludesAllOfInvoiceItems) SetCatalogVariantIdNil(b bool)`

 SetCatalogVariantIdNil sets the value for CatalogVariantId to be an explicit nil

### UnsetCatalogVariantId
`func (o *InvoiceWithIncludesAllOfInvoiceItems) UnsetCatalogVariantId()`

UnsetCatalogVariantId ensures that no value is present for CatalogVariantId, not even an explicit nil
### GetUpdatedAt

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceWithIncludesAllOfInvoiceItems) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceWithIncludesAllOfInvoiceItems) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceWithIncludesAllOfInvoiceItems) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


