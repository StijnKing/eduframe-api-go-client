# CourseWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the course | [readonly] 
**Position** | **int32** | Sorting position of the course. Lower is higher. | 
**StartingPrice** | **string** | Lowest price of all of its planned courses. | 
**SignupUrl** | **string** | URL to the signup page for this course. | 
**Slug** | **string** | Human readable identifier, unique per educator. | 
**SlugHistory** | **[]string** | List of old slugs, old calls will be redirected. | [readonly] 
**Avatar** | **string** | URL to the original avatar image file. | 
**AvatarUrl** | **string** | URL to a resized avatar image (300x200^). | 
**AvatarThumbUrl** | **string** | URL to a resized avatar image (748x296^). | 
**ConditionsOrDefault** | **NullableString** | Conditions for this course with a fallback to the default course conditions of the educator. | 
**WebsiteUrl** | **NullableString** | Expected URL of the course on the educator website. | 
**CertificateTemplateId** | **NullableInt32** | Identifier of the optionally linked certificate template. | 
**CategoryId** | **int32** | Identifier of the category of the course. | 
**Name** | **string** | The name of the course. | 
**Code** | **string** | The code of the course. | 
**Duration** | Pointer to **NullableString** | The duration of the course. | [optional] 
**Level** | Pointer to **NullableString** | A string indicating the level of the course. | [optional] 
**MetaTitle** | **NullableString** | Meta title of the course for SEO purposes. | 
**MetaDescription** | **NullableString** | Meta description of the course for SEO purposes. | 
**Result** | Pointer to **NullableString** | The result of the course | [optional] 
**LabelIds** | **[]int32** | An array of integers representing unique identifier values associated with labels.  | 
**Cost** | **NullableString** | The price to be paid for this course. | 
**CostScheme** | [**CostScheme**](CostScheme.md) |  | 
**IsPublished** | **bool** | Boolean representing the publishable status of the course. | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 
**Custom** | **map[string]interface{}** | The custom properties of the program. | 

## Methods

### NewCourseWithIncludes

`func NewCourseWithIncludes(id int32, position int32, startingPrice string, signupUrl string, slug string, slugHistory []string, avatar string, avatarUrl string, avatarThumbUrl string, conditionsOrDefault NullableString, websiteUrl NullableString, certificateTemplateId NullableInt32, categoryId int32, name string, code string, metaTitle NullableString, metaDescription NullableString, labelIds []int32, cost NullableString, costScheme CostScheme, isPublished bool, updatedAt time.Time, createdAt time.Time, custom map[string]interface{}, ) *CourseWithIncludes`

NewCourseWithIncludes instantiates a new CourseWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseWithIncludesWithDefaults

`func NewCourseWithIncludesWithDefaults() *CourseWithIncludes`

NewCourseWithIncludesWithDefaults instantiates a new CourseWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetPosition

`func (o *CourseWithIncludes) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CourseWithIncludes) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CourseWithIncludes) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetStartingPrice

`func (o *CourseWithIncludes) GetStartingPrice() string`

GetStartingPrice returns the StartingPrice field if non-nil, zero value otherwise.

### GetStartingPriceOk

`func (o *CourseWithIncludes) GetStartingPriceOk() (*string, bool)`

GetStartingPriceOk returns a tuple with the StartingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingPrice

`func (o *CourseWithIncludes) SetStartingPrice(v string)`

SetStartingPrice sets StartingPrice field to given value.


### GetSignupUrl

`func (o *CourseWithIncludes) GetSignupUrl() string`

GetSignupUrl returns the SignupUrl field if non-nil, zero value otherwise.

### GetSignupUrlOk

`func (o *CourseWithIncludes) GetSignupUrlOk() (*string, bool)`

GetSignupUrlOk returns a tuple with the SignupUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupUrl

`func (o *CourseWithIncludes) SetSignupUrl(v string)`

SetSignupUrl sets SignupUrl field to given value.


### GetSlug

`func (o *CourseWithIncludes) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CourseWithIncludes) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CourseWithIncludes) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetSlugHistory

`func (o *CourseWithIncludes) GetSlugHistory() []string`

GetSlugHistory returns the SlugHistory field if non-nil, zero value otherwise.

### GetSlugHistoryOk

`func (o *CourseWithIncludes) GetSlugHistoryOk() (*[]string, bool)`

GetSlugHistoryOk returns a tuple with the SlugHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugHistory

`func (o *CourseWithIncludes) SetSlugHistory(v []string)`

SetSlugHistory sets SlugHistory field to given value.


### GetAvatar

`func (o *CourseWithIncludes) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CourseWithIncludes) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CourseWithIncludes) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetAvatarUrl

`func (o *CourseWithIncludes) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *CourseWithIncludes) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *CourseWithIncludes) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetAvatarThumbUrl

`func (o *CourseWithIncludes) GetAvatarThumbUrl() string`

GetAvatarThumbUrl returns the AvatarThumbUrl field if non-nil, zero value otherwise.

### GetAvatarThumbUrlOk

`func (o *CourseWithIncludes) GetAvatarThumbUrlOk() (*string, bool)`

GetAvatarThumbUrlOk returns a tuple with the AvatarThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarThumbUrl

`func (o *CourseWithIncludes) SetAvatarThumbUrl(v string)`

SetAvatarThumbUrl sets AvatarThumbUrl field to given value.


### GetConditionsOrDefault

`func (o *CourseWithIncludes) GetConditionsOrDefault() string`

GetConditionsOrDefault returns the ConditionsOrDefault field if non-nil, zero value otherwise.

### GetConditionsOrDefaultOk

`func (o *CourseWithIncludes) GetConditionsOrDefaultOk() (*string, bool)`

GetConditionsOrDefaultOk returns a tuple with the ConditionsOrDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionsOrDefault

`func (o *CourseWithIncludes) SetConditionsOrDefault(v string)`

SetConditionsOrDefault sets ConditionsOrDefault field to given value.


### SetConditionsOrDefaultNil

`func (o *CourseWithIncludes) SetConditionsOrDefaultNil(b bool)`

 SetConditionsOrDefaultNil sets the value for ConditionsOrDefault to be an explicit nil

### UnsetConditionsOrDefault
`func (o *CourseWithIncludes) UnsetConditionsOrDefault()`

UnsetConditionsOrDefault ensures that no value is present for ConditionsOrDefault, not even an explicit nil
### GetWebsiteUrl

`func (o *CourseWithIncludes) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *CourseWithIncludes) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *CourseWithIncludes) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.


### SetWebsiteUrlNil

`func (o *CourseWithIncludes) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *CourseWithIncludes) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetCertificateTemplateId

`func (o *CourseWithIncludes) GetCertificateTemplateId() int32`

GetCertificateTemplateId returns the CertificateTemplateId field if non-nil, zero value otherwise.

### GetCertificateTemplateIdOk

`func (o *CourseWithIncludes) GetCertificateTemplateIdOk() (*int32, bool)`

GetCertificateTemplateIdOk returns a tuple with the CertificateTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplateId

`func (o *CourseWithIncludes) SetCertificateTemplateId(v int32)`

SetCertificateTemplateId sets CertificateTemplateId field to given value.


### SetCertificateTemplateIdNil

`func (o *CourseWithIncludes) SetCertificateTemplateIdNil(b bool)`

 SetCertificateTemplateIdNil sets the value for CertificateTemplateId to be an explicit nil

### UnsetCertificateTemplateId
`func (o *CourseWithIncludes) UnsetCertificateTemplateId()`

UnsetCertificateTemplateId ensures that no value is present for CertificateTemplateId, not even an explicit nil
### GetCategoryId

`func (o *CourseWithIncludes) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CourseWithIncludes) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CourseWithIncludes) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetName

`func (o *CourseWithIncludes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CourseWithIncludes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CourseWithIncludes) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *CourseWithIncludes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CourseWithIncludes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CourseWithIncludes) SetCode(v string)`

SetCode sets Code field to given value.


### GetDuration

`func (o *CourseWithIncludes) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CourseWithIncludes) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CourseWithIncludes) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CourseWithIncludes) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *CourseWithIncludes) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *CourseWithIncludes) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetLevel

`func (o *CourseWithIncludes) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CourseWithIncludes) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CourseWithIncludes) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CourseWithIncludes) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *CourseWithIncludes) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *CourseWithIncludes) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetMetaTitle

`func (o *CourseWithIncludes) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *CourseWithIncludes) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *CourseWithIncludes) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.


### SetMetaTitleNil

`func (o *CourseWithIncludes) SetMetaTitleNil(b bool)`

 SetMetaTitleNil sets the value for MetaTitle to be an explicit nil

### UnsetMetaTitle
`func (o *CourseWithIncludes) UnsetMetaTitle()`

UnsetMetaTitle ensures that no value is present for MetaTitle, not even an explicit nil
### GetMetaDescription

`func (o *CourseWithIncludes) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *CourseWithIncludes) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *CourseWithIncludes) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.


### SetMetaDescriptionNil

`func (o *CourseWithIncludes) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *CourseWithIncludes) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
### GetResult

`func (o *CourseWithIncludes) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CourseWithIncludes) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CourseWithIncludes) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CourseWithIncludes) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *CourseWithIncludes) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *CourseWithIncludes) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetLabelIds

`func (o *CourseWithIncludes) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *CourseWithIncludes) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *CourseWithIncludes) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.


### GetCost

`func (o *CourseWithIncludes) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CourseWithIncludes) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CourseWithIncludes) SetCost(v string)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *CourseWithIncludes) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *CourseWithIncludes) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostScheme

`func (o *CourseWithIncludes) GetCostScheme() CostScheme`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *CourseWithIncludes) GetCostSchemeOk() (*CostScheme, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *CourseWithIncludes) SetCostScheme(v CostScheme)`

SetCostScheme sets CostScheme field to given value.


### GetIsPublished

`func (o *CourseWithIncludes) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *CourseWithIncludes) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *CourseWithIncludes) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.


### GetUpdatedAt

`func (o *CourseWithIncludes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CourseWithIncludes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CourseWithIncludes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *CourseWithIncludes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CourseWithIncludes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CourseWithIncludes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCustom

`func (o *CourseWithIncludes) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CourseWithIncludes) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CourseWithIncludes) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


