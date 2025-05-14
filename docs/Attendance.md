# Attendance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the attendance record. | [readonly] 
**MeetingId** | **int32** | Unique identifier of the meeting. | 
**EnrollmentId** | **int32** | Unique identifier of the enrollment. | 
**State** | [**AttendanceState**](AttendanceState.md) |  | 
**Comment** | **NullableString** | Comment about this attendance. | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewAttendance

`func NewAttendance(id int32, meetingId int32, enrollmentId int32, state AttendanceState, comment NullableString, updatedAt time.Time, createdAt time.Time, ) *Attendance`

NewAttendance instantiates a new Attendance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttendanceWithDefaults

`func NewAttendanceWithDefaults() *Attendance`

NewAttendanceWithDefaults instantiates a new Attendance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attendance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attendance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attendance) SetId(v int32)`

SetId sets Id field to given value.


### GetMeetingId

`func (o *Attendance) GetMeetingId() int32`

GetMeetingId returns the MeetingId field if non-nil, zero value otherwise.

### GetMeetingIdOk

`func (o *Attendance) GetMeetingIdOk() (*int32, bool)`

GetMeetingIdOk returns a tuple with the MeetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingId

`func (o *Attendance) SetMeetingId(v int32)`

SetMeetingId sets MeetingId field to given value.


### GetEnrollmentId

`func (o *Attendance) GetEnrollmentId() int32`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *Attendance) GetEnrollmentIdOk() (*int32, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *Attendance) SetEnrollmentId(v int32)`

SetEnrollmentId sets EnrollmentId field to given value.


### GetState

`func (o *Attendance) GetState() AttendanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Attendance) GetStateOk() (*AttendanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Attendance) SetState(v AttendanceState)`

SetState sets State field to given value.


### GetComment

`func (o *Attendance) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Attendance) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Attendance) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *Attendance) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Attendance) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetUpdatedAt

`func (o *Attendance) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Attendance) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Attendance) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Attendance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Attendance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Attendance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


