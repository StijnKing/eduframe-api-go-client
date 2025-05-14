# MeetingLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the meeting location | [readonly] 
**CourseLocationId** | **int32** | Unique identifier of the course location. | 
**Name** | **string** |  | 
**Capacity** | **NullableInt32** | Capacity of the meeting location | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewMeetingLocation

`func NewMeetingLocation(id int32, courseLocationId int32, name string, capacity NullableInt32, updatedAt time.Time, createdAt time.Time, ) *MeetingLocation`

NewMeetingLocation instantiates a new MeetingLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingLocationWithDefaults

`func NewMeetingLocationWithDefaults() *MeetingLocation`

NewMeetingLocationWithDefaults instantiates a new MeetingLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MeetingLocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeetingLocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeetingLocation) SetId(v int32)`

SetId sets Id field to given value.


### GetCourseLocationId

`func (o *MeetingLocation) GetCourseLocationId() int32`

GetCourseLocationId returns the CourseLocationId field if non-nil, zero value otherwise.

### GetCourseLocationIdOk

`func (o *MeetingLocation) GetCourseLocationIdOk() (*int32, bool)`

GetCourseLocationIdOk returns a tuple with the CourseLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseLocationId

`func (o *MeetingLocation) SetCourseLocationId(v int32)`

SetCourseLocationId sets CourseLocationId field to given value.


### GetName

`func (o *MeetingLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeetingLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeetingLocation) SetName(v string)`

SetName sets Name field to given value.


### GetCapacity

`func (o *MeetingLocation) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *MeetingLocation) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *MeetingLocation) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### SetCapacityNil

`func (o *MeetingLocation) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *MeetingLocation) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetUpdatedAt

`func (o *MeetingLocation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MeetingLocation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MeetingLocation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *MeetingLocation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MeetingLocation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MeetingLocation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


