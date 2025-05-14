# UpdateEnrollmentByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **NullableString** | If it is an enrollment of a fixed course, it equals the end date. For a flexible course, it returns the enrollment specific end date. | [optional] 

## Methods

### NewUpdateEnrollmentByIdRequest

`func NewUpdateEnrollmentByIdRequest() *UpdateEnrollmentByIdRequest`

NewUpdateEnrollmentByIdRequest instantiates a new UpdateEnrollmentByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEnrollmentByIdRequestWithDefaults

`func NewUpdateEnrollmentByIdRequestWithDefaults() *UpdateEnrollmentByIdRequest`

NewUpdateEnrollmentByIdRequestWithDefaults instantiates a new UpdateEnrollmentByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *UpdateEnrollmentByIdRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateEnrollmentByIdRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateEnrollmentByIdRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateEnrollmentByIdRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *UpdateEnrollmentByIdRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *UpdateEnrollmentByIdRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


