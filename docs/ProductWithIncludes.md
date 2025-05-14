# ProductWithIncludes

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
**CourseTabContents** | Pointer to [**[]CourseTabContent**](CourseTabContent.md) |  | [optional] 
**Custom** | Pointer to **map[string]interface{}** | The custom properties of the product. | [optional] 
**CustomAssociations** | Pointer to [**[]CustomAssociationWithRecords**](CustomAssociationWithRecords.md) | The custom associations linked to the catalog product.  ![Custom Objects](https://img.shields.io/badge/Feature-Custom_objects-1d8127) ![Beta](https://img.shields.io/badge/Beta-7d15a3)  | [optional] 

## Methods

### NewProductWithIncludes

`func NewProductWithIncludes(id int32, name string, slug string, labelIds []int32, costScheme CostScheme, cost NullableString, currency Currency, productableType ProductableType, productableId int32, avatar string, position int32, conditionsOrDefault NullableString, categoryId int32, isPublished bool, showOnWebsite bool, signupUrl string, updatedAt time.Time, createdAt time.Time, ) *ProductWithIncludes`

NewProductWithIncludes instantiates a new ProductWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithIncludesWithDefaults

`func NewProductWithIncludesWithDefaults() *ProductWithIncludes`

NewProductWithIncludesWithDefaults instantiates a new ProductWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProductWithIncludes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductWithIncludes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductWithIncludes) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *ProductWithIncludes) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProductWithIncludes) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProductWithIncludes) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetLabelIds

`func (o *ProductWithIncludes) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *ProductWithIncludes) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *ProductWithIncludes) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.


### GetCostScheme

`func (o *ProductWithIncludes) GetCostScheme() CostScheme`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *ProductWithIncludes) GetCostSchemeOk() (*CostScheme, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *ProductWithIncludes) SetCostScheme(v CostScheme)`

SetCostScheme sets CostScheme field to given value.


### GetCost

`func (o *ProductWithIncludes) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProductWithIncludes) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProductWithIncludes) SetCost(v string)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *ProductWithIncludes) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *ProductWithIncludes) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCurrency

`func (o *ProductWithIncludes) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductWithIncludes) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductWithIncludes) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetProductableType

`func (o *ProductWithIncludes) GetProductableType() ProductableType`

GetProductableType returns the ProductableType field if non-nil, zero value otherwise.

### GetProductableTypeOk

`func (o *ProductWithIncludes) GetProductableTypeOk() (*ProductableType, bool)`

GetProductableTypeOk returns a tuple with the ProductableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductableType

`func (o *ProductWithIncludes) SetProductableType(v ProductableType)`

SetProductableType sets ProductableType field to given value.


### GetProductableId

`func (o *ProductWithIncludes) GetProductableId() int32`

GetProductableId returns the ProductableId field if non-nil, zero value otherwise.

### GetProductableIdOk

`func (o *ProductWithIncludes) GetProductableIdOk() (*int32, bool)`

GetProductableIdOk returns a tuple with the ProductableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductableId

`func (o *ProductWithIncludes) SetProductableId(v int32)`

SetProductableId sets ProductableId field to given value.


### GetAvatar

`func (o *ProductWithIncludes) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ProductWithIncludes) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ProductWithIncludes) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetPosition

`func (o *ProductWithIncludes) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProductWithIncludes) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProductWithIncludes) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetConditionsOrDefault

`func (o *ProductWithIncludes) GetConditionsOrDefault() string`

GetConditionsOrDefault returns the ConditionsOrDefault field if non-nil, zero value otherwise.

### GetConditionsOrDefaultOk

`func (o *ProductWithIncludes) GetConditionsOrDefaultOk() (*string, bool)`

GetConditionsOrDefaultOk returns a tuple with the ConditionsOrDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionsOrDefault

`func (o *ProductWithIncludes) SetConditionsOrDefault(v string)`

SetConditionsOrDefault sets ConditionsOrDefault field to given value.


### SetConditionsOrDefaultNil

`func (o *ProductWithIncludes) SetConditionsOrDefaultNil(b bool)`

 SetConditionsOrDefaultNil sets the value for ConditionsOrDefault to be an explicit nil

### UnsetConditionsOrDefault
`func (o *ProductWithIncludes) UnsetConditionsOrDefault()`

UnsetConditionsOrDefault ensures that no value is present for ConditionsOrDefault, not even an explicit nil
### GetCategoryId

`func (o *ProductWithIncludes) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ProductWithIncludes) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ProductWithIncludes) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetIsPublished

`func (o *ProductWithIncludes) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *ProductWithIncludes) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *ProductWithIncludes) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.


### GetShowOnWebsite

`func (o *ProductWithIncludes) GetShowOnWebsite() bool`

GetShowOnWebsite returns the ShowOnWebsite field if non-nil, zero value otherwise.

### GetShowOnWebsiteOk

`func (o *ProductWithIncludes) GetShowOnWebsiteOk() (*bool, bool)`

GetShowOnWebsiteOk returns a tuple with the ShowOnWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnWebsite

`func (o *ProductWithIncludes) SetShowOnWebsite(v bool)`

SetShowOnWebsite sets ShowOnWebsite field to given value.


### GetSignupUrl

`func (o *ProductWithIncludes) GetSignupUrl() string`

GetSignupUrl returns the SignupUrl field if non-nil, zero value otherwise.

### GetSignupUrlOk

`func (o *ProductWithIncludes) GetSignupUrlOk() (*string, bool)`

GetSignupUrlOk returns a tuple with the SignupUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupUrl

`func (o *ProductWithIncludes) SetSignupUrl(v string)`

SetSignupUrl sets SignupUrl field to given value.


### GetUpdatedAt

`func (o *ProductWithIncludes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProductWithIncludes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProductWithIncludes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *ProductWithIncludes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductWithIncludes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductWithIncludes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCourseTabContents

`func (o *ProductWithIncludes) GetCourseTabContents() []CourseTabContent`

GetCourseTabContents returns the CourseTabContents field if non-nil, zero value otherwise.

### GetCourseTabContentsOk

`func (o *ProductWithIncludes) GetCourseTabContentsOk() (*[]CourseTabContent, bool)`

GetCourseTabContentsOk returns a tuple with the CourseTabContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseTabContents

`func (o *ProductWithIncludes) SetCourseTabContents(v []CourseTabContent)`

SetCourseTabContents sets CourseTabContents field to given value.

### HasCourseTabContents

`func (o *ProductWithIncludes) HasCourseTabContents() bool`

HasCourseTabContents returns a boolean if a field has been set.

### GetCustom

`func (o *ProductWithIncludes) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ProductWithIncludes) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ProductWithIncludes) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *ProductWithIncludes) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomAssociations

`func (o *ProductWithIncludes) GetCustomAssociations() []CustomAssociationWithRecords`

GetCustomAssociations returns the CustomAssociations field if non-nil, zero value otherwise.

### GetCustomAssociationsOk

`func (o *ProductWithIncludes) GetCustomAssociationsOk() (*[]CustomAssociationWithRecords, bool)`

GetCustomAssociationsOk returns a tuple with the CustomAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssociations

`func (o *ProductWithIncludes) SetCustomAssociations(v []CustomAssociationWithRecords)`

SetCustomAssociations sets CustomAssociations field to given value.

### HasCustomAssociations

`func (o *ProductWithIncludes) HasCustomAssociations() bool`

HasCustomAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


