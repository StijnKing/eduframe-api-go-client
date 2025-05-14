# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the catalog product. | [readonly] 
**Name** | **string** | Name of the product. | 
**Slug** | **string** | Slug of the product. | 
**LabelIds** | **[]int32** | An array of integers representing unique identifier values associated with labels.  | 
**CostScheme** | [**CostScheme**](CostScheme.md) |  | 
**Cost** | **NullableString** | Cost of the product. | 
**Currency** | [**Currency**](Currency.md) |  | 
**ProductableType** | [**ProductableType**](ProductableType.md) |  | 
**ProductableId** | **int32** | Unique identifier of the productable. | 
**Avatar** | **string** | URL to the original avatar image file. | 
**Position** | **int32** | Sorting position of the element. Lower is higher. | 
**ConditionsOrDefault** | **NullableString** | Conditions for this product with a fallback to the default course conditions of the educator. | 
**CategoryId** | **int32** | Identifier of the category of the course. | 
**IsPublished** | **bool** | Boolean showing if the product is published or not. | 
**ShowOnWebsite** | **bool** | Boolean showing if the product is shown on the website or not. | [readonly] 
**SignupUrl** | **string** | URL to the signup page for this course. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewProduct

`func NewProduct(id int32, name string, slug string, labelIds []int32, costScheme CostScheme, cost NullableString, currency Currency, productableType ProductableType, productableId int32, avatar string, position int32, conditionsOrDefault NullableString, categoryId int32, isPublished bool, showOnWebsite bool, signupUrl string, updatedAt time.Time, createdAt time.Time, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Product) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Product) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Product) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetLabelIds

`func (o *Product) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *Product) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *Product) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.


### GetCostScheme

`func (o *Product) GetCostScheme() CostScheme`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *Product) GetCostSchemeOk() (*CostScheme, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *Product) SetCostScheme(v CostScheme)`

SetCostScheme sets CostScheme field to given value.


### GetCost

`func (o *Product) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Product) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Product) SetCost(v string)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *Product) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *Product) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCurrency

`func (o *Product) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Product) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Product) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetProductableType

`func (o *Product) GetProductableType() ProductableType`

GetProductableType returns the ProductableType field if non-nil, zero value otherwise.

### GetProductableTypeOk

`func (o *Product) GetProductableTypeOk() (*ProductableType, bool)`

GetProductableTypeOk returns a tuple with the ProductableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductableType

`func (o *Product) SetProductableType(v ProductableType)`

SetProductableType sets ProductableType field to given value.


### GetProductableId

`func (o *Product) GetProductableId() int32`

GetProductableId returns the ProductableId field if non-nil, zero value otherwise.

### GetProductableIdOk

`func (o *Product) GetProductableIdOk() (*int32, bool)`

GetProductableIdOk returns a tuple with the ProductableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductableId

`func (o *Product) SetProductableId(v int32)`

SetProductableId sets ProductableId field to given value.


### GetAvatar

`func (o *Product) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Product) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Product) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetPosition

`func (o *Product) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Product) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Product) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetConditionsOrDefault

`func (o *Product) GetConditionsOrDefault() string`

GetConditionsOrDefault returns the ConditionsOrDefault field if non-nil, zero value otherwise.

### GetConditionsOrDefaultOk

`func (o *Product) GetConditionsOrDefaultOk() (*string, bool)`

GetConditionsOrDefaultOk returns a tuple with the ConditionsOrDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionsOrDefault

`func (o *Product) SetConditionsOrDefault(v string)`

SetConditionsOrDefault sets ConditionsOrDefault field to given value.


### SetConditionsOrDefaultNil

`func (o *Product) SetConditionsOrDefaultNil(b bool)`

 SetConditionsOrDefaultNil sets the value for ConditionsOrDefault to be an explicit nil

### UnsetConditionsOrDefault
`func (o *Product) UnsetConditionsOrDefault()`

UnsetConditionsOrDefault ensures that no value is present for ConditionsOrDefault, not even an explicit nil
### GetCategoryId

`func (o *Product) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *Product) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *Product) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetIsPublished

`func (o *Product) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *Product) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *Product) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.


### GetShowOnWebsite

`func (o *Product) GetShowOnWebsite() bool`

GetShowOnWebsite returns the ShowOnWebsite field if non-nil, zero value otherwise.

### GetShowOnWebsiteOk

`func (o *Product) GetShowOnWebsiteOk() (*bool, bool)`

GetShowOnWebsiteOk returns a tuple with the ShowOnWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnWebsite

`func (o *Product) SetShowOnWebsite(v bool)`

SetShowOnWebsite sets ShowOnWebsite field to given value.


### GetSignupUrl

`func (o *Product) GetSignupUrl() string`

GetSignupUrl returns the SignupUrl field if non-nil, zero value otherwise.

### GetSignupUrlOk

`func (o *Product) GetSignupUrlOk() (*string, bool)`

GetSignupUrlOk returns a tuple with the SignupUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupUrl

`func (o *Product) SetSignupUrl(v string)`

SetSignupUrl sets SignupUrl field to given value.


### GetUpdatedAt

`func (o *Product) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Product) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Product) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Product) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Product) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Product) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


