# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the invoice. | 
**ReferenceId** | **string** | Reference id of the Invoice. Is used to find the invoice. | 
**NumberInt** | **NullableFloat32** | The invoice number converted to an integer value. | 
**OrderId** | **NullableInt32** | Unique identifier of the order the invoice belongs to. | 
**OrderNumber** | **NullableFloat32** | The order number of the invoice. | 
**Status** | [**InvoiceStatus**](InvoiceStatus.md) |  | 
**ExpirationDate** | **NullableString** | The expiration date of the invoice. Is set when the invoice is set to open. | 
**OpenedAt** | **NullableString** | The date when the invoice status was changed to open. | 
**Description** | **string** | Human readable description of the invoice. | 
**AccountName** | **string** | The name of the account that is paying. Is copied to the invoice and is thus not automatically updated if the account name changes. | 
**Currency** | [**Currency**](Currency.md) |  | 
**TotalIncl** | **string** | The total cost of the invoice including VAT. | 
**TotalExcl** | **string** | The total cost of the invoice excluding VAT. | 
**TotalOpen** | **string** | The open cost of the invoice including VAT. | 
**PdfUrl** | **string** | Url to the download path of the invoice in PDF format. | 
**XmlUrl** | **string** | Url to the download path of the invoice in UML format. | 
**Number** | **string** | The invoice number which is unique per educator. If left empty, it autoincrements. | 
**AccountId** | **int32** | Identifier of the account. | 
**Feature** | **NullableString** | Some description of the invoice which is displayed on the invoice. | 
**Footnote** | **NullableString** | The note displayed at the bottom of the invoice. | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewInvoice

`func NewInvoice(id int32, referenceId string, numberInt NullableFloat32, orderId NullableInt32, orderNumber NullableFloat32, status InvoiceStatus, expirationDate NullableString, openedAt NullableString, description string, accountName string, currency Currency, totalIncl string, totalExcl string, totalOpen string, pdfUrl string, xmlUrl string, number string, accountId int32, feature NullableString, footnote NullableString, updatedAt time.Time, createdAt time.Time, ) *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v int32)`

SetId sets Id field to given value.


### GetReferenceId

`func (o *Invoice) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Invoice) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Invoice) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetNumberInt

`func (o *Invoice) GetNumberInt() float32`

GetNumberInt returns the NumberInt field if non-nil, zero value otherwise.

### GetNumberIntOk

`func (o *Invoice) GetNumberIntOk() (*float32, bool)`

GetNumberIntOk returns a tuple with the NumberInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberInt

`func (o *Invoice) SetNumberInt(v float32)`

SetNumberInt sets NumberInt field to given value.


### SetNumberIntNil

`func (o *Invoice) SetNumberIntNil(b bool)`

 SetNumberIntNil sets the value for NumberInt to be an explicit nil

### UnsetNumberInt
`func (o *Invoice) UnsetNumberInt()`

UnsetNumberInt ensures that no value is present for NumberInt, not even an explicit nil
### GetOrderId

`func (o *Invoice) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Invoice) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Invoice) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### SetOrderIdNil

`func (o *Invoice) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *Invoice) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetOrderNumber

`func (o *Invoice) GetOrderNumber() float32`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *Invoice) GetOrderNumberOk() (*float32, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *Invoice) SetOrderNumber(v float32)`

SetOrderNumber sets OrderNumber field to given value.


### SetOrderNumberNil

`func (o *Invoice) SetOrderNumberNil(b bool)`

 SetOrderNumberNil sets the value for OrderNumber to be an explicit nil

### UnsetOrderNumber
`func (o *Invoice) UnsetOrderNumber()`

UnsetOrderNumber ensures that no value is present for OrderNumber, not even an explicit nil
### GetStatus

`func (o *Invoice) GetStatus() InvoiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*InvoiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v InvoiceStatus)`

SetStatus sets Status field to given value.


### GetExpirationDate

`func (o *Invoice) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Invoice) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Invoice) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### SetExpirationDateNil

`func (o *Invoice) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *Invoice) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetOpenedAt

`func (o *Invoice) GetOpenedAt() string`

GetOpenedAt returns the OpenedAt field if non-nil, zero value otherwise.

### GetOpenedAtOk

`func (o *Invoice) GetOpenedAtOk() (*string, bool)`

GetOpenedAtOk returns a tuple with the OpenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAt

`func (o *Invoice) SetOpenedAt(v string)`

SetOpenedAt sets OpenedAt field to given value.


### SetOpenedAtNil

`func (o *Invoice) SetOpenedAtNil(b bool)`

 SetOpenedAtNil sets the value for OpenedAt to be an explicit nil

### UnsetOpenedAt
`func (o *Invoice) UnsetOpenedAt()`

UnsetOpenedAt ensures that no value is present for OpenedAt, not even an explicit nil
### GetDescription

`func (o *Invoice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Invoice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Invoice) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAccountName

`func (o *Invoice) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Invoice) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Invoice) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetCurrency

`func (o *Invoice) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetTotalIncl

`func (o *Invoice) GetTotalIncl() string`

GetTotalIncl returns the TotalIncl field if non-nil, zero value otherwise.

### GetTotalInclOk

`func (o *Invoice) GetTotalInclOk() (*string, bool)`

GetTotalInclOk returns a tuple with the TotalIncl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIncl

`func (o *Invoice) SetTotalIncl(v string)`

SetTotalIncl sets TotalIncl field to given value.


### GetTotalExcl

`func (o *Invoice) GetTotalExcl() string`

GetTotalExcl returns the TotalExcl field if non-nil, zero value otherwise.

### GetTotalExclOk

`func (o *Invoice) GetTotalExclOk() (*string, bool)`

GetTotalExclOk returns a tuple with the TotalExcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExcl

`func (o *Invoice) SetTotalExcl(v string)`

SetTotalExcl sets TotalExcl field to given value.


### GetTotalOpen

`func (o *Invoice) GetTotalOpen() string`

GetTotalOpen returns the TotalOpen field if non-nil, zero value otherwise.

### GetTotalOpenOk

`func (o *Invoice) GetTotalOpenOk() (*string, bool)`

GetTotalOpenOk returns a tuple with the TotalOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpen

`func (o *Invoice) SetTotalOpen(v string)`

SetTotalOpen sets TotalOpen field to given value.


### GetPdfUrl

`func (o *Invoice) GetPdfUrl() string`

GetPdfUrl returns the PdfUrl field if non-nil, zero value otherwise.

### GetPdfUrlOk

`func (o *Invoice) GetPdfUrlOk() (*string, bool)`

GetPdfUrlOk returns a tuple with the PdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfUrl

`func (o *Invoice) SetPdfUrl(v string)`

SetPdfUrl sets PdfUrl field to given value.


### GetXmlUrl

`func (o *Invoice) GetXmlUrl() string`

GetXmlUrl returns the XmlUrl field if non-nil, zero value otherwise.

### GetXmlUrlOk

`func (o *Invoice) GetXmlUrlOk() (*string, bool)`

GetXmlUrlOk returns a tuple with the XmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlUrl

`func (o *Invoice) SetXmlUrl(v string)`

SetXmlUrl sets XmlUrl field to given value.


### GetNumber

`func (o *Invoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Invoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Invoice) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetAccountId

`func (o *Invoice) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Invoice) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Invoice) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetFeature

`func (o *Invoice) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *Invoice) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *Invoice) SetFeature(v string)`

SetFeature sets Feature field to given value.


### SetFeatureNil

`func (o *Invoice) SetFeatureNil(b bool)`

 SetFeatureNil sets the value for Feature to be an explicit nil

### UnsetFeature
`func (o *Invoice) UnsetFeature()`

UnsetFeature ensures that no value is present for Feature, not even an explicit nil
### GetFootnote

`func (o *Invoice) GetFootnote() string`

GetFootnote returns the Footnote field if non-nil, zero value otherwise.

### GetFootnoteOk

`func (o *Invoice) GetFootnoteOk() (*string, bool)`

GetFootnoteOk returns a tuple with the Footnote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFootnote

`func (o *Invoice) SetFootnote(v string)`

SetFootnote sets Footnote field to given value.


### SetFootnoteNil

`func (o *Invoice) SetFootnoteNil(b bool)`

 SetFootnoteNil sets the value for Footnote to be an explicit nil

### UnsetFootnote
`func (o *Invoice) UnsetFootnote()`

UnsetFootnote ensures that no value is present for Footnote, not even an explicit nil
### GetUpdatedAt

`func (o *Invoice) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Invoice) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Invoice) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Invoice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Invoice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Invoice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


