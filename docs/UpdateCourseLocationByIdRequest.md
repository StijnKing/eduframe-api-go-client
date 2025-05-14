# UpdateCourseLocationByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the location where the course is held. | [optional] 
**AddressAttributes** | Pointer to [**NullableAddressPatchPayload**](AddressPatchPayload.md) |  | [optional] 

## Methods

### NewUpdateCourseLocationByIdRequest

`func NewUpdateCourseLocationByIdRequest() *UpdateCourseLocationByIdRequest`

NewUpdateCourseLocationByIdRequest instantiates a new UpdateCourseLocationByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCourseLocationByIdRequestWithDefaults

`func NewUpdateCourseLocationByIdRequestWithDefaults() *UpdateCourseLocationByIdRequest`

NewUpdateCourseLocationByIdRequestWithDefaults instantiates a new UpdateCourseLocationByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCourseLocationByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCourseLocationByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCourseLocationByIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCourseLocationByIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddressAttributes

`func (o *UpdateCourseLocationByIdRequest) GetAddressAttributes() AddressPatchPayload`

GetAddressAttributes returns the AddressAttributes field if non-nil, zero value otherwise.

### GetAddressAttributesOk

`func (o *UpdateCourseLocationByIdRequest) GetAddressAttributesOk() (*AddressPatchPayload, bool)`

GetAddressAttributesOk returns a tuple with the AddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributes

`func (o *UpdateCourseLocationByIdRequest) SetAddressAttributes(v AddressPatchPayload)`

SetAddressAttributes sets AddressAttributes field to given value.

### HasAddressAttributes

`func (o *UpdateCourseLocationByIdRequest) HasAddressAttributes() bool`

HasAddressAttributes returns a boolean if a field has been set.

### SetAddressAttributesNil

`func (o *UpdateCourseLocationByIdRequest) SetAddressAttributesNil(b bool)`

 SetAddressAttributesNil sets the value for AddressAttributes to be an explicit nil

### UnsetAddressAttributes
`func (o *UpdateCourseLocationByIdRequest) UnsetAddressAttributes()`

UnsetAddressAttributes ensures that no value is present for AddressAttributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


