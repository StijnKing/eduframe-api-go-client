# MeetingLocationWithIncludes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the meeting location | [readonly] 
**CourseLocationId** | **int32** | Unique identifier of the course location. | 
**Name** | **string** |  | 
**Capacity** | **NullableInt32** | Capacity of the meeting location | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 
**Address** | Pointer to [**NullableAddress**](Address.md) |  | [optional] 

## Methods

### NewMeetingLocationWithIncludes

`func NewMeetingLocationWithIncludes(id int32, courseLocationId int32, name string, capacity NullableInt32, updatedAt time.Time, createdAt time.Time, ) *MeetingLocationWithIncludes`

NewMeetingLocationWithIncludes instantiates a new MeetingLocationWithIncludes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingLocationWithIncludesWithDefaults

`func NewMeetingLocationWithIncludesWithDefaults() *MeetingLocationWithIncludes`

NewMeetingLocationWithIncludesWithDefaults instantiates a new MeetingLocationWithIncludes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MeetingLocationWithIncludes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeetingLocationWithIncludes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeetingLocationWithIncludes) SetId(v int32)`

SetId sets Id field to given value.


### GetCourseLocationId

`func (o *MeetingLocationWithIncludes) GetCourseLocationId() int32`

GetCourseLocationId returns the CourseLocationId field if non-nil, zero value otherwise.

### GetCourseLocationIdOk

`func (o *MeetingLocationWithIncludes) GetCourseLocationIdOk() (*int32, bool)`

GetCourseLocationIdOk returns a tuple with the CourseLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseLocationId

`func (o *MeetingLocationWithIncludes) SetCourseLocationId(v int32)`

SetCourseLocationId sets CourseLocationId field to given value.


### GetName

`func (o *MeetingLocationWithIncludes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeetingLocationWithIncludes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeetingLocationWithIncludes) SetName(v string)`

SetName sets Name field to given value.


### GetCapacity

`func (o *MeetingLocationWithIncludes) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *MeetingLocationWithIncludes) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *MeetingLocationWithIncludes) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### SetCapacityNil

`func (o *MeetingLocationWithIncludes) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *MeetingLocationWithIncludes) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetUpdatedAt

`func (o *MeetingLocationWithIncludes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MeetingLocationWithIncludes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MeetingLocationWithIncludes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *MeetingLocationWithIncludes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MeetingLocationWithIncludes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MeetingLocationWithIncludes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetAddress

`func (o *MeetingLocationWithIncludes) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MeetingLocationWithIncludes) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MeetingLocationWithIncludes) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MeetingLocationWithIncludes) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *MeetingLocationWithIncludes) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MeetingLocationWithIncludes) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


