# PlannedCourse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the user. | [readonly] 
**Status** | [**PlannedCourseStatus**](PlannedCourseStatus.md) |  | 
**DurationInDays** | **NullableInt32** | The period of time of the planned course. For flexible planned courses this equals the provided +duration+. For fixed planned courses this equals the number of meetings if any, if the fixed planned course has no meetings, it&#39;s the number of days between the start and end date.  | 
**AvailabilityState** | [**AvailabilityState**](AvailabilityState.md) |  | 
**Payable** | **bool** | Boolean wether there are payments involved for this course. Basically its true if the cost_scheme is +student+ or +order+.  | 
**CurrentParticipants** | **int32** | The current amount of participants. | 
**ConfirmedActiveAndCompletedEnrollmentsCount** | **int32** | The amount of confirmed active and completed enrollments. | 
**RequestedEnrollmentsCount** | **int32** | The amount of requested enrollments. ie. the number of reserved seats | 
**AvailablePlaces** | **bool** | Boolean if there are any places available. | 
**CanvasLink** | Pointer to **NullableString** | URL to the course in canvas. Is only returned if available. | [optional] 
**Currency** | [**Currency**](Currency.md) |  | 
**CostMultiplier** | **NullableString** | A positive float representing the multiplier for the VAT. Say you have 21% VAT, this will return +1.21+.  | 
**IsPublished** | **bool** | Boolean if is published on the website. | 
**CourseId** | **int32** | Unique identifier of the course. | 
**Type** | **string** | The type of the course. | 
**StartDate** | **string** | Date at which the planned course starts. Only needed for fixed planned courses. | 
**EndDate** | **NullableString** | Date at which the planned course ends. Only needed for fixed planned courses. | 
**MinParticipants** | **NullableInt32** | A number representing the minimum number of participants that can enroll for the planned course. | 
**MaxParticipants** | **NullableInt32** | A number representing the maximum number of participants that can enroll for the planned course. | 
**CostScheme** | [**CostScheme**](CostScheme.md) |  | 
**Cost** | **NullableString** | A positive float representing the price of the planned course. | 
**CourseVariantId** | **NullableInt32** | Unique identifier of the course variant. | 
**CourseLocationId** | **NullableInt32** | Unique identifier of the course location. | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewPlannedCourse

`func NewPlannedCourse(id int32, status PlannedCourseStatus, durationInDays NullableInt32, availabilityState AvailabilityState, payable bool, currentParticipants int32, confirmedActiveAndCompletedEnrollmentsCount int32, requestedEnrollmentsCount int32, availablePlaces bool, currency Currency, costMultiplier NullableString, isPublished bool, courseId int32, type_ string, startDate string, endDate NullableString, minParticipants NullableInt32, maxParticipants NullableInt32, costScheme CostScheme, cost NullableString, courseVariantId NullableInt32, courseLocationId NullableInt32, updatedAt time.Time, createdAt time.Time, ) *PlannedCourse`

NewPlannedCourse instantiates a new PlannedCourse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannedCourseWithDefaults

`func NewPlannedCourseWithDefaults() *PlannedCourse`

NewPlannedCourseWithDefaults instantiates a new PlannedCourse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlannedCourse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlannedCourse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlannedCourse) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *PlannedCourse) GetStatus() PlannedCourseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlannedCourse) GetStatusOk() (*PlannedCourseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlannedCourse) SetStatus(v PlannedCourseStatus)`

SetStatus sets Status field to given value.


### GetDurationInDays

`func (o *PlannedCourse) GetDurationInDays() int32`

GetDurationInDays returns the DurationInDays field if non-nil, zero value otherwise.

### GetDurationInDaysOk

`func (o *PlannedCourse) GetDurationInDaysOk() (*int32, bool)`

GetDurationInDaysOk returns a tuple with the DurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInDays

`func (o *PlannedCourse) SetDurationInDays(v int32)`

SetDurationInDays sets DurationInDays field to given value.


### SetDurationInDaysNil

`func (o *PlannedCourse) SetDurationInDaysNil(b bool)`

 SetDurationInDaysNil sets the value for DurationInDays to be an explicit nil

### UnsetDurationInDays
`func (o *PlannedCourse) UnsetDurationInDays()`

UnsetDurationInDays ensures that no value is present for DurationInDays, not even an explicit nil
### GetAvailabilityState

`func (o *PlannedCourse) GetAvailabilityState() AvailabilityState`

GetAvailabilityState returns the AvailabilityState field if non-nil, zero value otherwise.

### GetAvailabilityStateOk

`func (o *PlannedCourse) GetAvailabilityStateOk() (*AvailabilityState, bool)`

GetAvailabilityStateOk returns a tuple with the AvailabilityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityState

`func (o *PlannedCourse) SetAvailabilityState(v AvailabilityState)`

SetAvailabilityState sets AvailabilityState field to given value.


### GetPayable

`func (o *PlannedCourse) GetPayable() bool`

GetPayable returns the Payable field if non-nil, zero value otherwise.

### GetPayableOk

`func (o *PlannedCourse) GetPayableOk() (*bool, bool)`

GetPayableOk returns a tuple with the Payable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayable

`func (o *PlannedCourse) SetPayable(v bool)`

SetPayable sets Payable field to given value.


### GetCurrentParticipants

`func (o *PlannedCourse) GetCurrentParticipants() int32`

GetCurrentParticipants returns the CurrentParticipants field if non-nil, zero value otherwise.

### GetCurrentParticipantsOk

`func (o *PlannedCourse) GetCurrentParticipantsOk() (*int32, bool)`

GetCurrentParticipantsOk returns a tuple with the CurrentParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentParticipants

`func (o *PlannedCourse) SetCurrentParticipants(v int32)`

SetCurrentParticipants sets CurrentParticipants field to given value.


### GetConfirmedActiveAndCompletedEnrollmentsCount

`func (o *PlannedCourse) GetConfirmedActiveAndCompletedEnrollmentsCount() int32`

GetConfirmedActiveAndCompletedEnrollmentsCount returns the ConfirmedActiveAndCompletedEnrollmentsCount field if non-nil, zero value otherwise.

### GetConfirmedActiveAndCompletedEnrollmentsCountOk

`func (o *PlannedCourse) GetConfirmedActiveAndCompletedEnrollmentsCountOk() (*int32, bool)`

GetConfirmedActiveAndCompletedEnrollmentsCountOk returns a tuple with the ConfirmedActiveAndCompletedEnrollmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedActiveAndCompletedEnrollmentsCount

`func (o *PlannedCourse) SetConfirmedActiveAndCompletedEnrollmentsCount(v int32)`

SetConfirmedActiveAndCompletedEnrollmentsCount sets ConfirmedActiveAndCompletedEnrollmentsCount field to given value.


### GetRequestedEnrollmentsCount

`func (o *PlannedCourse) GetRequestedEnrollmentsCount() int32`

GetRequestedEnrollmentsCount returns the RequestedEnrollmentsCount field if non-nil, zero value otherwise.

### GetRequestedEnrollmentsCountOk

`func (o *PlannedCourse) GetRequestedEnrollmentsCountOk() (*int32, bool)`

GetRequestedEnrollmentsCountOk returns a tuple with the RequestedEnrollmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedEnrollmentsCount

`func (o *PlannedCourse) SetRequestedEnrollmentsCount(v int32)`

SetRequestedEnrollmentsCount sets RequestedEnrollmentsCount field to given value.


### GetAvailablePlaces

`func (o *PlannedCourse) GetAvailablePlaces() bool`

GetAvailablePlaces returns the AvailablePlaces field if non-nil, zero value otherwise.

### GetAvailablePlacesOk

`func (o *PlannedCourse) GetAvailablePlacesOk() (*bool, bool)`

GetAvailablePlacesOk returns a tuple with the AvailablePlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePlaces

`func (o *PlannedCourse) SetAvailablePlaces(v bool)`

SetAvailablePlaces sets AvailablePlaces field to given value.


### GetCanvasLink

`func (o *PlannedCourse) GetCanvasLink() string`

GetCanvasLink returns the CanvasLink field if non-nil, zero value otherwise.

### GetCanvasLinkOk

`func (o *PlannedCourse) GetCanvasLinkOk() (*string, bool)`

GetCanvasLinkOk returns a tuple with the CanvasLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvasLink

`func (o *PlannedCourse) SetCanvasLink(v string)`

SetCanvasLink sets CanvasLink field to given value.

### HasCanvasLink

`func (o *PlannedCourse) HasCanvasLink() bool`

HasCanvasLink returns a boolean if a field has been set.

### SetCanvasLinkNil

`func (o *PlannedCourse) SetCanvasLinkNil(b bool)`

 SetCanvasLinkNil sets the value for CanvasLink to be an explicit nil

### UnsetCanvasLink
`func (o *PlannedCourse) UnsetCanvasLink()`

UnsetCanvasLink ensures that no value is present for CanvasLink, not even an explicit nil
### GetCurrency

`func (o *PlannedCourse) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PlannedCourse) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PlannedCourse) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetCostMultiplier

`func (o *PlannedCourse) GetCostMultiplier() string`

GetCostMultiplier returns the CostMultiplier field if non-nil, zero value otherwise.

### GetCostMultiplierOk

`func (o *PlannedCourse) GetCostMultiplierOk() (*string, bool)`

GetCostMultiplierOk returns a tuple with the CostMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostMultiplier

`func (o *PlannedCourse) SetCostMultiplier(v string)`

SetCostMultiplier sets CostMultiplier field to given value.


### SetCostMultiplierNil

`func (o *PlannedCourse) SetCostMultiplierNil(b bool)`

 SetCostMultiplierNil sets the value for CostMultiplier to be an explicit nil

### UnsetCostMultiplier
`func (o *PlannedCourse) UnsetCostMultiplier()`

UnsetCostMultiplier ensures that no value is present for CostMultiplier, not even an explicit nil
### GetIsPublished

`func (o *PlannedCourse) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *PlannedCourse) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *PlannedCourse) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.


### GetCourseId

`func (o *PlannedCourse) GetCourseId() int32`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *PlannedCourse) GetCourseIdOk() (*int32, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *PlannedCourse) SetCourseId(v int32)`

SetCourseId sets CourseId field to given value.


### GetType

`func (o *PlannedCourse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlannedCourse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlannedCourse) SetType(v string)`

SetType sets Type field to given value.


### GetStartDate

`func (o *PlannedCourse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PlannedCourse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PlannedCourse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *PlannedCourse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PlannedCourse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PlannedCourse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *PlannedCourse) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *PlannedCourse) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetMinParticipants

`func (o *PlannedCourse) GetMinParticipants() int32`

GetMinParticipants returns the MinParticipants field if non-nil, zero value otherwise.

### GetMinParticipantsOk

`func (o *PlannedCourse) GetMinParticipantsOk() (*int32, bool)`

GetMinParticipantsOk returns a tuple with the MinParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinParticipants

`func (o *PlannedCourse) SetMinParticipants(v int32)`

SetMinParticipants sets MinParticipants field to given value.


### SetMinParticipantsNil

`func (o *PlannedCourse) SetMinParticipantsNil(b bool)`

 SetMinParticipantsNil sets the value for MinParticipants to be an explicit nil

### UnsetMinParticipants
`func (o *PlannedCourse) UnsetMinParticipants()`

UnsetMinParticipants ensures that no value is present for MinParticipants, not even an explicit nil
### GetMaxParticipants

`func (o *PlannedCourse) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *PlannedCourse) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *PlannedCourse) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.


### SetMaxParticipantsNil

`func (o *PlannedCourse) SetMaxParticipantsNil(b bool)`

 SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil

### UnsetMaxParticipants
`func (o *PlannedCourse) UnsetMaxParticipants()`

UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
### GetCostScheme

`func (o *PlannedCourse) GetCostScheme() CostScheme`

GetCostScheme returns the CostScheme field if non-nil, zero value otherwise.

### GetCostSchemeOk

`func (o *PlannedCourse) GetCostSchemeOk() (*CostScheme, bool)`

GetCostSchemeOk returns a tuple with the CostScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostScheme

`func (o *PlannedCourse) SetCostScheme(v CostScheme)`

SetCostScheme sets CostScheme field to given value.


### GetCost

`func (o *PlannedCourse) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *PlannedCourse) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *PlannedCourse) SetCost(v string)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *PlannedCourse) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *PlannedCourse) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCourseVariantId

`func (o *PlannedCourse) GetCourseVariantId() int32`

GetCourseVariantId returns the CourseVariantId field if non-nil, zero value otherwise.

### GetCourseVariantIdOk

`func (o *PlannedCourse) GetCourseVariantIdOk() (*int32, bool)`

GetCourseVariantIdOk returns a tuple with the CourseVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseVariantId

`func (o *PlannedCourse) SetCourseVariantId(v int32)`

SetCourseVariantId sets CourseVariantId field to given value.


### SetCourseVariantIdNil

`func (o *PlannedCourse) SetCourseVariantIdNil(b bool)`

 SetCourseVariantIdNil sets the value for CourseVariantId to be an explicit nil

### UnsetCourseVariantId
`func (o *PlannedCourse) UnsetCourseVariantId()`

UnsetCourseVariantId ensures that no value is present for CourseVariantId, not even an explicit nil
### GetCourseLocationId

`func (o *PlannedCourse) GetCourseLocationId() int32`

GetCourseLocationId returns the CourseLocationId field if non-nil, zero value otherwise.

### GetCourseLocationIdOk

`func (o *PlannedCourse) GetCourseLocationIdOk() (*int32, bool)`

GetCourseLocationIdOk returns a tuple with the CourseLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseLocationId

`func (o *PlannedCourse) SetCourseLocationId(v int32)`

SetCourseLocationId sets CourseLocationId field to given value.


### SetCourseLocationIdNil

`func (o *PlannedCourse) SetCourseLocationIdNil(b bool)`

 SetCourseLocationIdNil sets the value for CourseLocationId to be an explicit nil

### UnsetCourseLocationId
`func (o *PlannedCourse) UnsetCourseLocationId()`

UnsetCourseLocationId ensures that no value is present for CourseLocationId, not even an explicit nil
### GetUpdatedAt

`func (o *PlannedCourse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlannedCourse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlannedCourse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *PlannedCourse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlannedCourse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlannedCourse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


