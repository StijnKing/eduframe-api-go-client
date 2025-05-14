# Order

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

## Methods

### NewOrder

`func NewOrder(id int32, number float32, status OrderStatus, numberOfStudents int32, origin string, customerComment NullableString, totalCostExcl NullableString, totalCostIncl NullableString, cost NullableString, costScheme CostScheme, catalogVariantId int32, creatorId NullableInt32, accountId int32, paymentOptionId NullableInt32, labelIds []int32, referralId NullableInt32, updatedAt time.Time, createdAt time.Time, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v int32)`

SetId sets Id field to given value.


### GetNumber

`func (o *Order) GetNumber() float32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Order) GetNumberOk() (*float32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Order) SetNumber(v float32)`

SetNumber sets Number field to given value.


### GetStatus

`func (o *Order) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.


### GetNumberOfStudents

`func (o *Order) GetNumberOfStudents() int32`

GetNumberOfStudents returns the NumberOfStudents field if non-nil, zero value otherwise.

### GetNumberOfStudentsOk

`func (o *Order) GetNumberOfStudentsOk() (*int32, bool)`

GetNumberOfStudentsOk returns a tuple with the NumberOfStudents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfStudents

`func (o *Order) SetNumberOfStudents(v int32)`

SetNumberOfStudents sets NumberOfStudents field to given value.


### GetOrigin

`func (o *Order) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Order) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Order) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetCustomerComment

`func (o *Order) GetCustomerComment() string`

GetCustomerComment returns the CustomerComment field if non-nil, zero value otherwise.

### GetCustomerCommentOk

`func (o *Order) GetCustomerCommentOk() (*string, bool)`

GetCustomerCommentOk returns a tuple with the CustomerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerComment

`func (o *Order) SetCustomerComment(v string)`

SetCustomerComment sets CustomerComment field to given value.


### SetCustomerCommentNil

`func (o *Order) SetCustomerCommentNil(b bool)`

 SetCustomerCommentNil sets the value for CustomerComment to be an explicit nil

### UnsetCustomerComment
`func (o *Order) UnsetCustomerComment()`

UnsetCustomerComment ensures that no value is present for CustomerComment, not even an explicit nil
### GetTotalCostExcl

`func (o *Order) GetTotalCostExcl() string`

GetTotalCostExcl returns the TotalCostExcl field if non-nil, zero value otherwise.

### GetTotalCostExclOk

`func (o *Order) GetTotalCostExclOk() (*string, bool)`

GetTotalCostExclOk returns a tuple with the TotalCostExcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostExcl

`func (o *Order) SetTotalCostExcl(v string)`

SetTotalCostExcl sets TotalCostExcl field to given value.


### SetTotalCostExclNil

`func (o *Order) SetTotalCostExclNil(b bool)`

 SetTotalCostExclNil sets the value for TotalCostExcl to be an explicit nil

### UnsetTotalCostExcl
`func (o *Order) UnsetTotalCostExcl()`

UnsetTotalCostExcl ensures that no value is present for TotalCostExcl, not even an explicit nil
### GetTotalCostIncl

`func (o *Order) GetTotalCostIncl() string`

GetTotalCostIncl returns the TotalCostIncl field if non-nil, zero value otherwise.

### GetTotalCostInclOk

`func (o *Order) GetTotalCostInclOk() (*string, bool)`

GetTotalCostInclOk returns a tuple with the TotalCostIncl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostIncl

`func (o *Order) SetTotalCostIncl(v string)`

SetTotalCostIncl sets TotalCostIncl field to given value.


### SetTotalCostInclNil

`func (o *Order) SetTotalCostInclNil(b bool)`

 SetTotalCostInclNil sets the value for TotalCostIncl to be an explicit nil

### UnsetTotalCostIncl
`func (o *Order) UnsetTotalCostIncl()`

UnsetTotalCostIncl ensures that no value is present for TotalCostIncl, not even an explicit nil
### GetCost

`func (o *Order) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Order) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Order) SetCost(v string)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *Order) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *Order) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostScheme

`func (o *Order) GetCostScheme() CostScheme`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *Order) GetCostSchemeOk() (*CostScheme, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *Order) SetCostScheme(v CostScheme)`

SetCostScheme sets CostScheme field to given value.


### GetCatalogVariantId

`func (o *Order) GetCatalogVariantId() int32`

GetCatalogVariantId returns the CatalogVariantId field if non-nil, zero value otherwise.

### GetCatalogVariantIdOk

`func (o *Order) GetCatalogVariantIdOk() (*int32, bool)`

GetCatalogVariantIdOk returns a tuple with the CatalogVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogVariantId

`func (o *Order) SetCatalogVariantId(v int32)`

SetCatalogVariantId sets CatalogVariantId field to given value.


### GetCreatorId

`func (o *Order) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Order) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Order) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### SetCreatorIdNil

`func (o *Order) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *Order) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetAccountId

`func (o *Order) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Order) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Order) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetPaymentMethodId

`func (o *Order) GetPaymentMethodId() int32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *Order) GetPaymentMethodIdOk() (*int32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *Order) SetPaymentMethodId(v int32)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *Order) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### SetPaymentMethodIdNil

`func (o *Order) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *Order) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetPaymentOptionId

`func (o *Order) GetPaymentOptionId() int32`

GetPaymentOptionId returns the PaymentOptionId field if non-nil, zero value otherwise.

### GetPaymentOptionIdOk

`func (o *Order) GetPaymentOptionIdOk() (*int32, bool)`

GetPaymentOptionIdOk returns a tuple with the PaymentOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOptionId

`func (o *Order) SetPaymentOptionId(v int32)`

SetPaymentOptionId sets PaymentOptionId field to given value.


### SetPaymentOptionIdNil

`func (o *Order) SetPaymentOptionIdNil(b bool)`

 SetPaymentOptionIdNil sets the value for PaymentOptionId to be an explicit nil

### UnsetPaymentOptionId
`func (o *Order) UnsetPaymentOptionId()`

UnsetPaymentOptionId ensures that no value is present for PaymentOptionId, not even an explicit nil
### GetLabelIds

`func (o *Order) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *Order) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *Order) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.


### GetDiscountCode

`func (o *Order) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *Order) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *Order) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *Order) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### SetDiscountCodeNil

`func (o *Order) SetDiscountCodeNil(b bool)`

 SetDiscountCodeNil sets the value for DiscountCode to be an explicit nil

### UnsetDiscountCode
`func (o *Order) UnsetDiscountCode()`

UnsetDiscountCode ensures that no value is present for DiscountCode, not even an explicit nil
### GetTotalDiscount

`func (o *Order) GetTotalDiscount() string`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *Order) GetTotalDiscountOk() (*string, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *Order) SetTotalDiscount(v string)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *Order) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### SetTotalDiscountNil

`func (o *Order) SetTotalDiscountNil(b bool)`

 SetTotalDiscountNil sets the value for TotalDiscount to be an explicit nil

### UnsetTotalDiscount
`func (o *Order) UnsetTotalDiscount()`

UnsetTotalDiscount ensures that no value is present for TotalDiscount, not even an explicit nil
### GetReferralId

`func (o *Order) GetReferralId() int32`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *Order) GetReferralIdOk() (*int32, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralId

`func (o *Order) SetReferralId(v int32)`

SetReferralId sets ReferralId field to given value.


### SetReferralIdNil

`func (o *Order) SetReferralIdNil(b bool)`

 SetReferralIdNil sets the value for ReferralId to be an explicit nil

### UnsetReferralId
`func (o *Order) UnsetReferralId()`

UnsetReferralId ensures that no value is present for ReferralId, not even an explicit nil
### GetUpdatedAt

`func (o *Order) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Order) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Order) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Order) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Order) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Order) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


