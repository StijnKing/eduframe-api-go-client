# OrderWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the order. | [readonly] 
**Number** | **float32** | Educator specific identifier of the order which autoincrements. | 
**Status** | [**OrderStatus**](OrderStatus.md) |  | 
**NumberOfStudents** | **int32** | Number of students in this order. | 
**Origin** | **string** | Where the order originated from ex. Unknown, backend, signup. | 
**CustomerComment** | **NullableString** | Comment from the customer. | 
**TotalCostExcl** | **NullableString** | Decimal representing the cost of the order excluding VAT | [readonly] 
**TotalCostIncl** | **NullableString** | Decimal representing the cost of the order including VAT | [readonly] 
**Cost** | **NullableString** | Decimal representing the value of the order | 
**CostScheme** | [**CostScheme**](CostScheme.md) |  | 
**CatalogVariantId** | **int32** | Unique identifier of the orders CatalogVariant. | 
**CreatorId** | **NullableInt32** | Unique identifier of the orders Creator (User). | 
**AccountId** | **int32** | Unique identifier of the orders Account. | 
**PaymentMethodId** | Pointer to **NullableInt32** | Unique identifier of the orders PaymentMethod. | [optional] 
**PaymentOptionId** | **NullableInt32** | Unique identifier of the orders PaymentOption. | 
**LabelIds** | **[]int32** | An array of integers representing unique identifier values associated with labels.  | 
**DiscountCode** | Pointer to **NullableString** | The discount code used for this order. | [optional] 
**TotalDiscount** | Pointer to **NullableString** | Decimal representing the total discount applied to this order. | [optional] 
**ReferralId** | **NullableInt32** | Identifier of the referral. | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the order. | [optional] 
**OrderItems** | Pointer to [**[]OrderItem**](OrderItem.md) |  | [optional] 

## Methods

### NewOrderWithIncludes

`func NewOrderWithIncludes(id int32, number float32, status OrderStatus, numberOfStudents int32, origin string, customerComment NullableString, totalCostExcl NullableString, totalCostIncl NullableString, cost NullableString, costScheme CostScheme, catalogVariantId int32, creatorId NullableInt32, accountId int32, paymentOptionId NullableInt32, labelIds []int32, referralId NullableInt32, updatedAt time.Time, createdAt time.Time, ) *OrderWithIncludes`

NewOrderWithIncludes instantiates a new OrderWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithIncludesWithDefaults

`func NewOrderWithIncludesWithDefaults() *OrderWithIncludes`

NewOrderWithIncludesWithDefaults instantiates a new OrderWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetNumber

`func (o *OrderWithIncludes) GetNumber() float32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OrderWithIncludes) GetNumberOk() (*float32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OrderWithIncludes) SetNumber(v float32)`

SetNumber sets Number field to given value.


### GetStatus

`func (o *OrderWithIncludes) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderWithIncludes) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderWithIncludes) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.


### GetNumberOfStudents

`func (o *OrderWithIncludes) GetNumberOfStudents() int32`

GetNumberOfStudents returns the NumberOfStudents field if non-nil, zero value otherwise.

### GetNumberOfStudentsOk

`func (o *OrderWithIncludes) GetNumberOfStudentsOk() (*int32, bool)`

GetNumberOfStudentsOk returns a tuple with the NumberOfStudents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfStudents

`func (o *OrderWithIncludes) SetNumberOfStudents(v int32)`

SetNumberOfStudents sets NumberOfStudents field to given value.


### GetOrigin

`func (o *OrderWithIncludes) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *OrderWithIncludes) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *OrderWithIncludes) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetCustomerComment

`func (o *OrderWithIncludes) GetCustomerComment() string`

GetCustomerComment returns the CustomerComment field if non-nil, zero value otherwise.

### GetCustomerCommentOk

`func (o *OrderWithIncludes) GetCustomerCommentOk() (*string, bool)`

GetCustomerCommentOk returns a tuple with the CustomerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerComment

`func (o *OrderWithIncludes) SetCustomerComment(v string)`

SetCustomerComment sets CustomerComment field to given value.


### SetCustomerCommentNil

`func (o *OrderWithIncludes) SetCustomerCommentNil(b bool)`

 SetCustomerCommentNil sets the value for CustomerComment to be an explicit nil

### UnsetCustomerComment
`func (o *OrderWithIncludes) UnsetCustomerComment()`

UnsetCustomerComment ensures that no value is present for CustomerComment, not even an explicit nil
### GetTotalCostExcl

`func (o *OrderWithIncludes) GetTotalCostExcl() string`

GetTotalCostExcl returns the TotalCostExcl field if non-nil, zero value otherwise.

### GetTotalCostExclOk

`func (o *OrderWithIncludes) GetTotalCostExclOk() (*string, bool)`

GetTotalCostExclOk returns a tuple with the TotalCostExcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostExcl

`func (o *OrderWithIncludes) SetTotalCostExcl(v string)`

SetTotalCostExcl sets TotalCostExcl field to given value.


### SetTotalCostExclNil

`func (o *OrderWithIncludes) SetTotalCostExclNil(b bool)`

 SetTotalCostExclNil sets the value for TotalCostExcl to be an explicit nil

### UnsetTotalCostExcl
`func (o *OrderWithIncludes) UnsetTotalCostExcl()`

UnsetTotalCostExcl ensures that no value is present for TotalCostExcl, not even an explicit nil
### GetTotalCostIncl

`func (o *OrderWithIncludes) GetTotalCostIncl() string`

GetTotalCostIncl returns the TotalCostIncl field if non-nil, zero value otherwise.

### GetTotalCostInclOk

`func (o *OrderWithIncludes) GetTotalCostInclOk() (*string, bool)`

GetTotalCostInclOk returns a tuple with the TotalCostIncl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostIncl

`func (o *OrderWithIncludes) SetTotalCostIncl(v string)`

SetTotalCostIncl sets TotalCostIncl field to given value.


### SetTotalCostInclNil

`func (o *OrderWithIncludes) SetTotalCostInclNil(b bool)`

 SetTotalCostInclNil sets the value for TotalCostIncl to be an explicit nil

### UnsetTotalCostIncl
`func (o *OrderWithIncludes) UnsetTotalCostIncl()`

UnsetTotalCostIncl ensures that no value is present for TotalCostIncl, not even an explicit nil
### GetCost

`func (o *OrderWithIncludes) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OrderWithIncludes) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OrderWithIncludes) SetCost(v string)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *OrderWithIncludes) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *OrderWithIncludes) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostScheme

`func (o *OrderWithIncludes) GetCostScheme() CostScheme`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *OrderWithIncludes) GetCostSchemeOk() (*CostScheme, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *OrderWithIncludes) SetCostScheme(v CostScheme)`

SetCostScheme sets CostScheme field to given value.


### GetCatalogVariantId

`func (o *OrderWithIncludes) GetCatalogVariantId() int32`

GetCatalogVariantId returns the CatalogVariantId field if non-nil, zero value otherwise.

### GetCatalogVariantIdOk

`func (o *OrderWithIncludes) GetCatalogVariantIdOk() (*int32, bool)`

GetCatalogVariantIdOk returns a tuple with the CatalogVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogVariantId

`func (o *OrderWithIncludes) SetCatalogVariantId(v int32)`

SetCatalogVariantId sets CatalogVariantId field to given value.


### GetCreatorId

`func (o *OrderWithIncludes) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *OrderWithIncludes) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *OrderWithIncludes) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### SetCreatorIdNil

`func (o *OrderWithIncludes) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *OrderWithIncludes) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetAccountId

`func (o *OrderWithIncludes) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OrderWithIncludes) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OrderWithIncludes) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetPaymentMethodId

`func (o *OrderWithIncludes) GetPaymentMethodId() int32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *OrderWithIncludes) GetPaymentMethodIdOk() (*int32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *OrderWithIncludes) SetPaymentMethodId(v int32)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *OrderWithIncludes) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### SetPaymentMethodIdNil

`func (o *OrderWithIncludes) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *OrderWithIncludes) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetPaymentOptionId

`func (o *OrderWithIncludes) GetPaymentOptionId() int32`

GetPaymentOptionId returns the PaymentOptionId field if non-nil, zero value otherwise.

### GetPaymentOptionIdOk

`func (o *OrderWithIncludes) GetPaymentOptionIdOk() (*int32, bool)`

GetPaymentOptionIdOk returns a tuple with the PaymentOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOptionId

`func (o *OrderWithIncludes) SetPaymentOptionId(v int32)`

SetPaymentOptionId sets PaymentOptionId field to given value.


### SetPaymentOptionIdNil

`func (o *OrderWithIncludes) SetPaymentOptionIdNil(b bool)`

 SetPaymentOptionIdNil sets the value for PaymentOptionId to be an explicit nil

### UnsetPaymentOptionId
`func (o *OrderWithIncludes) UnsetPaymentOptionId()`

UnsetPaymentOptionId ensures that no value is present for PaymentOptionId, not even an explicit nil
### GetLabelIds

`func (o *OrderWithIncludes) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *OrderWithIncludes) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *OrderWithIncludes) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.


### GetDiscountCode

`func (o *OrderWithIncludes) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *OrderWithIncludes) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *OrderWithIncludes) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *OrderWithIncludes) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### SetDiscountCodeNil

`func (o *OrderWithIncludes) SetDiscountCodeNil(b bool)`

 SetDiscountCodeNil sets the value for DiscountCode to be an explicit nil

### UnsetDiscountCode
`func (o *OrderWithIncludes) UnsetDiscountCode()`

UnsetDiscountCode ensures that no value is present for DiscountCode, not even an explicit nil
### GetTotalDiscount

`func (o *OrderWithIncludes) GetTotalDiscount() string`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *OrderWithIncludes) GetTotalDiscountOk() (*string, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *OrderWithIncludes) SetTotalDiscount(v string)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *OrderWithIncludes) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### SetTotalDiscountNil

`func (o *OrderWithIncludes) SetTotalDiscountNil(b bool)`

 SetTotalDiscountNil sets the value for TotalDiscount to be an explicit nil

### UnsetTotalDiscount
`func (o *OrderWithIncludes) UnsetTotalDiscount()`

UnsetTotalDiscount ensures that no value is present for TotalDiscount, not even an explicit nil
### GetReferralId

`func (o *OrderWithIncludes) GetReferralId() int32`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *OrderWithIncludes) GetReferralIdOk() (*int32, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralId

`func (o *OrderWithIncludes) SetReferralId(v int32)`

SetReferralId sets ReferralId field to given value.


### SetReferralIdNil

`func (o *OrderWithIncludes) SetReferralIdNil(b bool)`

 SetReferralIdNil sets the value for ReferralId to be an explicit nil

### UnsetReferralId
`func (o *OrderWithIncludes) UnsetReferralId()`

UnsetReferralId ensures that no value is present for ReferralId, not even an explicit nil
### GetUpdatedAt

`func (o *OrderWithIncludes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrderWithIncludes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrderWithIncludes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *OrderWithIncludes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderWithIncludes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderWithIncludes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCustom

`func (o *OrderWithIncludes) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *OrderWithIncludes) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *OrderWithIncludes) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *OrderWithIncludes) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetOrderItems

`func (o *OrderWithIncludes) GetOrderItems() []OrderItem`

GetOrderItems returns the OrderItems field if non-nil, zero value otherwise.

### GetOrderItemsOk

`func (o *OrderWithIncludes) GetOrderItemsOk() (*[]OrderItem, bool)`

GetOrderItemsOk returns a tuple with the OrderItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItems

`func (o *OrderWithIncludes) SetOrderItems(v []OrderItem)`

SetOrderItems sets OrderItems field to given value.

### HasOrderItems

`func (o *OrderWithIncludes) HasOrderItems() bool`

HasOrderItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


