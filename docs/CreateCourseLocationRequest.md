# CreateCourseLocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the location where the course is held. | 
**AddressAttributes** | Pointer to [**NullableAddressPayload**](AddressPayload.md) |  | [optional] 

## Methods

### NewCreateCourseLocationRequest

`func NewCreateCourseLocationRequest(name string, ) *CreateCourseLocationRequest`

NewCreateCourseLocationRequest instantiates a new CreateCourseLocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCourseLocationRequestWithDefaults

`func NewCreateCourseLocationRequestWithDefaults() *CreateCourseLocationRequest`

NewCreateCourseLocationRequestWithDefaults instantiates a new CreateCourseLocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCourseLocationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCourseLocationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCourseLocationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAddressAttributes

`func (o *CreateCourseLocationRequest) GetAddressAttributes() AddressPayload`

GetAddressAttributes returns the AddressAttributes field if non-nil, zero value otherwise.

### GetAddressAttributesOk

`func (o *CreateCourseLocationRequest) GetAddressAttributesOk() (*AddressPayload, bool)`

GetAddressAttributesOk returns a tuple with the AddressAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributes

`func (o *CreateCourseLocationRequest) SetAddressAttributes(v AddressPayload)`

SetAddressAttributes sets AddressAttributes field to given value.

### HasAddressAttributes

`func (o *CreateCourseLocationRequest) HasAddressAttributes() bool`

HasAddressAttributes returns a boolean if a field has been set.

### SetAddressAttributesNil

`func (o *CreateCourseLocationRequest) SetAddressAttributesNil(b bool)`

 SetAddressAttributesNil sets the value for AddressAttributes to be an explicit nil

### UnsetAddressAttributes
`func (o *CreateCourseLocationRequest) UnsetAddressAttributes()`

UnsetAddressAttributes ensures that no value is present for AddressAttributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


