# UpdateMeetingLocationByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the meeting location. | [optional] 
**CourseLocationId** | Pointer to **int32** | Unique identifier of the course location. | [optional] 
**Capacity** | Pointer to **NullableFloat32** | Capacity of the meeting location | [optional] 
**AddressAttributes** | Pointer to [**NullableAddressPatchPayload**](AddressPatchPayload.md) |  | [optional] 

## Methods

### NewUpdateMeetingLocationByIdRequest

`func NewUpdateMeetingLocationByIdRequest() *UpdateMeetingLocationByIdRequest`

NewUpdateMeetingLocationByIdRequest instantiates a new UpdateMeetingLocationByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMeetingLocationByIdRequestWithDefaults

`func NewUpdateMeetingLocationByIdRequestWithDefaults() *UpdateMeetingLocationByIdRequest`

NewUpdateMeetingLocationByIdRequestWithDefaults instantiates a new UpdateMeetingLocationByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMeetingLocationByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMeetingLocationByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMeetingLocationByIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMeetingLocationByIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCourseLocationId

`func (o *UpdateMeetingLocationByIdRequest) GetCourseLocationId() int32`

GetCourseLocationId returns the CourseLocationId field if non-nil, zero value otherwise.

### GetCourseLocationIdOk

`func (o *UpdateMeetingLocationByIdRequest) GetCourseLocationIdOk() (*int32, bool)`

GetCourseLocationIdOk returns a tuple with the CourseLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseLocationId

`func (o *UpdateMeetingLocationByIdRequest) SetCourseLocationId(v int32)`

SetCourseLocationId sets CourseLocationId field to given value.

### HasCourseLocationId

`func (o *UpdateMeetingLocationByIdRequest) HasCourseLocationId() bool`

HasCourseLocationId returns a boolean if a field has been set.

### GetCapacity

`func (o *UpdateMeetingLocationByIdRequest) GetCapacity() float32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *UpdateMeetingLocationByIdRequest) GetCapacityOk() (*float32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *UpdateMeetingLocationByIdRequest) SetCapacity(v float32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *UpdateMeetingLocationByIdRequest) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### SetCapacityNil

`func (o *UpdateMeetingLocationByIdRequest) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *UpdateMeetingLocationByIdRequest) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetAddressAttributes

`func (o *UpdateMeetingLocationByIdRequest) GetAddressAttributes() AddressPatchPayload`

GetAddressAttributes returns the AddressAttributes field if non-nil, zero value otherwise.

### GetAddressAttributesOk

`func (o *UpdateMeetingLocationByIdRequest) GetAddressAttributesOk() (*AddressPatchPayload, bool)`

GetAddressAttributesOk returns a tuple with the AddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributes

`func (o *UpdateMeetingLocationByIdRequest) SetAddressAttributes(v AddressPatchPayload)`

SetAddressAttributes sets AddressAttributes field to given value.

### HasAddressAttributes

`func (o *UpdateMeetingLocationByIdRequest) HasAddressAttributes() bool`

HasAddressAttributes returns a boolean if a field has been set.

### SetAddressAttributesNil

`func (o *UpdateMeetingLocationByIdRequest) SetAddressAttributesNil(b bool)`

 SetAddressAttributesNil sets the value for AddressAttributes to be an explicit nil

### UnsetAddressAttributes
`func (o *UpdateMeetingLocationByIdRequest) UnsetAddressAttributes()`

UnsetAddressAttributes ensures that no value is present for AddressAttributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


