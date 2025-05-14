# CreateInvoiceRequestInvoiceItemsAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogVariantId** | Pointer to **NullableInt32** | Unique identifier of the catalog variant. | [optional] 
**Units** | **float32** | Integer representing the number of units of the invoice item. | 
**UnitPrice** | **string** | Decimal representing the price of an unit. | 
**Name** | **string** | The name of the invoice item. | 
**InvoiceVatId** | Pointer to **int32** | Identifier of the invoice vat. | [optional] 
**Destroy** | Pointer to **bool** | Set if you want to delete this item. | [optional] 

## Methods

### NewCreateInvoiceRequestInvoiceItemsAttributesInner

`func NewCreateInvoiceRequestInvoiceItemsAttributesInner(units float32, unitPrice string, name string, ) *CreateInvoiceRequestInvoiceItemsAttributesInner`

NewCreateInvoiceRequestInvoiceItemsAttributesInner instantiates a new CreateInvoiceRequestInvoiceItemsAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvoiceRequestInvoiceItemsAttributesInnerWithDefaults

`func NewCreateInvoiceRequestInvoiceItemsAttributesInnerWithDefaults() *CreateInvoiceRequestInvoiceItemsAttributesInner`

NewCreateInvoiceRequestInvoiceItemsAttributesInnerWithDefaults instantiates a new CreateInvoiceRequestInvoiceItemsAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogVariantId

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetCatalogVariantId() int32`

GetCatalogVariantId returns the CatalogVariantId field if non-nil, zero value otherwise.

### GetCatalogVariantIdOk

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetCatalogVariantIdOk() (*int32, bool)`

GetCatalogVariantIdOk returns a tuple with the CatalogVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogVariantId

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) SetCatalogVariantId(v int32)`

SetCatalogVariantId sets CatalogVariantId field to given value.

### HasCatalogVariantId

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) HasCatalogVariantId() bool`

HasCatalogVariantId returns a boolean if a field has been set.

### SetCatalogVariantIdNil

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) SetCatalogVariantIdNil(b bool)`

 SetCatalogVariantIdNil sets the value for CatalogVariantId to be an explicit nil

### UnsetCatalogVariantId
`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) UnsetCatalogVariantId()`

UnsetCatalogVariantId ensures that no value is present for CatalogVariantId, not even an explicit nil
### GetUnits

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetUnits() float32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetUnitsOk() (*float32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) SetUnits(v float32)`

SetUnits sets Units field to given value.


### GetUnitPrice

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.


### GetName

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) SetName(v string)`

SetName sets Name field to given value.


### GetInvoiceVatId

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetInvoiceVatId() int32`

GetInvoiceVatId returns the InvoiceVatId field if non-nil, zero value otherwise.

### GetInvoiceVatIdOk

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetInvoiceVatIdOk() (*int32, bool)`

GetInvoiceVatIdOk returns a tuple with the InvoiceVatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceVatId

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) SetInvoiceVatId(v int32)`

SetInvoiceVatId sets InvoiceVatId field to given value.

### HasInvoiceVatId

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) HasInvoiceVatId() bool`

HasInvoiceVatId returns a boolean if a field has been set.

### GetDestroy

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetDestroy() bool`

GetDestroy returns the Destroy field if non-nil, zero value otherwise.

### GetDestroyOk

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) GetDestroyOk() (*bool, bool)`

GetDestroyOk returns a tuple with the Destroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroy

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) SetDestroy(v bool)`

SetDestroy sets Destroy field to given value.

### HasDestroy

`func (o *CreateInvoiceRequestInvoiceItemsAttributesInner) HasDestroy() bool`

HasDestroy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


