# CreateMeetingLocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the meeting location. | 
**CourseLocationId** | **int32** | Unique identifier of the course location. | 
**Capacity** | Pointer to **NullableInt32** | Capacity of the meeting location | [optional] 
**AddressAttributes** | Pointer to [**NullableAddressPayload**](AddressPayload.md) |  | [optional] 

## Methods

### NewCreateMeetingLocationRequest

`func NewCreateMeetingLocationRequest(name string, courseLocationId int32, ) *CreateMeetingLocationRequest`

NewCreateMeetingLocationRequest instantiates a new CreateMeetingLocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMeetingLocationRequestWithDefaults

`func NewCreateMeetingLocationRequestWithDefaults() *CreateMeetingLocationRequest`

NewCreateMeetingLocationRequestWithDefaults instantiates a new CreateMeetingLocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMeetingLocationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMeetingLocationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMeetingLocationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCourseLocationId

`func (o *CreateMeetingLocationRequest) GetCourseLocationId() int32`

GetCourseLocationId returns the CourseLocationId field if non-nil, zero value otherwise.

### GetCourseLocationIdOk

`func (o *CreateMeetingLocationRequest) GetCourseLocationIdOk() (*int32, bool)`

GetCourseLocationIdOk returns a tuple with the CourseLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseLocationId

`func (o *CreateMeetingLocationRequest) SetCourseLocationId(v int32)`

SetCourseLocationId sets CourseLocationId field to given value.


### GetCapacity

`func (o *CreateMeetingLocationRequest) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *CreateMeetingLocationRequest) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *CreateMeetingLocationRequest) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *CreateMeetingLocationRequest) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### SetCapacityNil

`func (o *CreateMeetingLocationRequest) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *CreateMeetingLocationRequest) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetAddressAttributes

`func (o *CreateMeetingLocationRequest) GetAddressAttributes() AddressPayload`

GetAddressAttributes returns the AddressAttributes field if non-nil, zero value otherwise.

### GetAddressAttributesOk

`func (o *CreateMeetingLocationRequest) GetAddressAttributesOk() (*AddressPayload, bool)`

GetAddressAttributesOk returns a tuple with the AddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributes

`func (o *CreateMeetingLocationRequest) SetAddressAttributes(v AddressPayload)`

SetAddressAttributes sets AddressAttributes field to given value.

### HasAddressAttributes

`func (o *CreateMeetingLocationRequest) HasAddressAttributes() bool`

HasAddressAttributes returns a boolean if a field has been set.

### SetAddressAttributesNil

`func (o *CreateMeetingLocationRequest) SetAddressAttributesNil(b bool)`

 SetAddressAttributesNil sets the value for AddressAttributes to be an explicit nil

### UnsetAddressAttributes
`func (o *CreateMeetingLocationRequest) UnsetAddressAttributes()`

UnsetAddressAttributes ensures that no value is present for AddressAttributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


