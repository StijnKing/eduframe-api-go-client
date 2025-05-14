# SetAttendanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeetingId** | **int32** | Unique identifier of the meeting. | 
**EnrollmentId** | **int32** | Unique identifier of the enrollment. | 
**State** | Pointer to [**AttendanceState**](AttendanceState.md) |  | [optional] 
**Comment** | Pointer to **NullableString** | Comment about this attendance. | [optional] 

## Methods

### NewSetAttendanceRequest

`func NewSetAttendanceRequest(meetingId int32, enrollmentId int32, ) *SetAttendanceRequest`

NewSetAttendanceRequest instantiates a new SetAttendanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAttendanceRequestWithDefaults

`func NewSetAttendanceRequestWithDefaults() *SetAttendanceRequest`

NewSetAttendanceRequestWithDefaults instantiates a new SetAttendanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeetingId

`func (o *SetAttendanceRequest) GetMeetingId() int32`

GetMeetingId returns the MeetingId field if non-nil, zero value otherwise.

### GetMeetingIdOk

`func (o *SetAttendanceRequest) GetMeetingIdOk() (*int32, bool)`

GetMeetingIdOk returns a tuple with the MeetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingId

`func (o *SetAttendanceRequest) SetMeetingId(v int32)`

SetMeetingId sets MeetingId field to given value.


### GetEnrollmentId

`func (o *SetAttendanceRequest) GetEnrollmentId() int32`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SetAttendanceRequest) GetEnrollmentIdOk() (*int32, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SetAttendanceRequest) SetEnrollmentId(v int32)`

SetEnrollmentId sets EnrollmentId field to given value.


### GetState

`func (o *SetAttendanceRequest) GetState() AttendanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SetAttendanceRequest) GetStateOk() (*AttendanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SetAttendanceRequest) SetState(v AttendanceState)`

SetState sets State field to given value.

### HasState

`func (o *SetAttendanceRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetComment

`func (o *SetAttendanceRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SetAttendanceRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SetAttendanceRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SetAttendanceRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *SetAttendanceRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *SetAttendanceRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


