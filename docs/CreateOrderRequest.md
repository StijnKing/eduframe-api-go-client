# CreateOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **NullableString** | Decimal representing the value of the order | [optional] 
**CostScheme** | Pointer to **string** | The cost schema that the payment will follow for the specified order. | [optional] 
**CatalogVariantId** | **int32** | Unique identifier of the orders CatalogVariant. | 
**CreatorId** | **int32** | Unique identifier of the orders Creator (User). | 
**AccountId** | Pointer to **int32** | The unique identifier associated with the orders Account. If not provided, the system will default to using the personal account.  | [optional] 
**PlannedCourseId** | Pointer to **int32** | *DEPRECATED*: Use catalog_variant_id instead. Unique identifier of the order&#39;s planned course.  | [optional] 
**PaymentMethodId** | Pointer to **int32** | Unique identifier of the orders PaymentMethod. | [optional] 
**StudentIds** | Pointer to **[]int32** | Array of student ids. A non-empty array is required if there are no student ids specified in the enrollments_attributes. | [optional] 
**PaymentOptionId** | Pointer to **int32** | Unique identifier of the orders PaymentOption. | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the order. | [optional] 
**Approve** | Pointer to **bool** | Optional: If the order should be approved or not. When omitted will default to false | [optional] 
**LabelIds** | Pointer to **[]int32** | Optional: Assign labels to the order. | [optional] 
**ReferralId** | Pointer to **int32** | Optional: Identifier of the referral. | [optional] 

## Methods

### NewCreateOrderRequest

`func NewCreateOrderRequest(catalogVariantId int32, creatorId int32, ) *CreateOrderRequest`

NewCreateOrderRequest instantiates a new CreateOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderRequestWithDefaults

`func NewCreateOrderRequestWithDefaults() *CreateOrderRequest`

NewCreateOrderRequestWithDefaults instantiates a new CreateOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *CreateOrderRequest) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CreateOrderRequest) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CreateOrderRequest) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CreateOrderRequest) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *CreateOrderRequest) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *CreateOrderRequest) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostScheme

`func (o *CreateOrderRequest) GetCostScheme() string`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *CreateOrderRequest) GetCostSchemeOk() (*string, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *CreateOrderRequest) SetCostScheme(v string)`

SetCostScheme sets CostScheme field to given value.

### HasCostScheme

`func (o *CreateOrderRequest) HasCostScheme() bool`

HasCostScheme returns a boolean if a field has been set.

### GetCatalogVariantId

`func (o *CreateOrderRequest) GetCatalogVariantId() int32`

GetCatalogVariantId returns the CatalogVariantId field if non-nil, zero value otherwise.

### GetCatalogVariantIdOk

`func (o *CreateOrderRequest) GetCatalogVariantIdOk() (*int32, bool)`

GetCatalogVariantIdOk returns a tuple with the CatalogVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogVariantId

`func (o *CreateOrderRequest) SetCatalogVariantId(v int32)`

SetCatalogVariantId sets CatalogVariantId field to given value.


### GetCreatorId

`func (o *CreateOrderRequest) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *CreateOrderRequest) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *CreateOrderRequest) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetAccountId

`func (o *CreateOrderRequest) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateOrderRequest) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateOrderRequest) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateOrderRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlannedCourseId

`func (o *CreateOrderRequest) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *CreateOrderRequest) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *CreateOrderRequest) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.

### HasPlannedCourseId

`func (o *CreateOrderRequest) HasPlannedCourseId() bool`

HasPlannedCourseId returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *CreateOrderRequest) GetPaymentMethodId() int32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *CreateOrderRequest) GetPaymentMethodIdOk() (*int32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *CreateOrderRequest) SetPaymentMethodId(v int32)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *CreateOrderRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetStudentIds

`func (o *CreateOrderRequest) GetStudentIds() []int32`

GetStudentIds returns the StudentIds field if non-nil, zero value otherwise.

### GetStudentIdsOk

`func (o *CreateOrderRequest) GetStudentIdsOk() (*[]int32, bool)`

GetStudentIdsOk returns a tuple with the StudentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudentIds

`func (o *CreateOrderRequest) SetStudentIds(v []int32)`

SetStudentIds sets StudentIds field to given value.

### HasStudentIds

`func (o *CreateOrderRequest) HasStudentIds() bool`

HasStudentIds returns a boolean if a field has been set.

### GetPaymentOptionId

`func (o *CreateOrderRequest) GetPaymentOptionId() int32`

GetPaymentOptionId returns the PaymentOptionId field if non-nil, zero value otherwise.

### GetPaymentOptionIdOk

`func (o *CreateOrderRequest) GetPaymentOptionIdOk() (*int32, bool)`

GetPaymentOptionIdOk returns a tuple with the PaymentOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOptionId

`func (o *CreateOrderRequest) SetPaymentOptionId(v int32)`

SetPaymentOptionId sets PaymentOptionId field to given value.

### HasPaymentOptionId

`func (o *CreateOrderRequest) HasPaymentOptionId() bool`

HasPaymentOptionId returns a boolean if a field has been set.

### GetCustom

`func (o *CreateOrderRequest) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CreateOrderRequest) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CreateOrderRequest) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *CreateOrderRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetApprove

`func (o *CreateOrderRequest) GetApprove() bool`

GetApprove returns the Approve field if non-nil, zero value otherwise.

### GetApproveOk

`func (o *CreateOrderRequest) GetApproveOk() (*bool, bool)`

GetApproveOk returns a tuple with the Approve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprove

`func (o *CreateOrderRequest) SetApprove(v bool)`

SetApprove sets Approve field to given value.

### HasApprove

`func (o *CreateOrderRequest) HasApprove() bool`

HasApprove returns a boolean if a field has been set.

### GetLabelIds

`func (o *CreateOrderRequest) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *CreateOrderRequest) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *CreateOrderRequest) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *CreateOrderRequest) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetReferralId

`func (o *CreateOrderRequest) GetReferralId() int32`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *CreateOrderRequest) GetReferralIdOk() (*int32, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralId

`func (o *CreateOrderRequest) SetReferralId(v int32)`

SetReferralId sets ReferralId field to given value.

### HasReferralId

`func (o *CreateOrderRequest) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


