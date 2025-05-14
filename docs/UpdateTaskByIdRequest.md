# UpdateTaskByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The title of the task. | [optional] 
**Description** | Pointer to **NullableString** | A string representing the description of the task. | [optional] 
**DueDate** | Pointer to **NullableString** | Date when the task must be completed. | [optional] 
**Starred** | Pointer to **bool** | Boolean if the task is starred. | [optional] 
**AssigneeId** | Pointer to **NullableInt32** | Unique identifier of the assigned user for the task. | [optional] 
**SubjectType** | Pointer to **NullableString** | Type of the subject. | [optional] 
**SubjectId** | Pointer to **NullableInt32** | Identifier of the subject. | [optional] 
**Completed** | Pointer to **bool** | Boolean representing the status of the task. The default value is false. | [optional] 

## Methods

### NewUpdateTaskByIdRequest

`func NewUpdateTaskByIdRequest() *UpdateTaskByIdRequest`

NewUpdateTaskByIdRequest instantiates a new UpdateTaskByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskByIdRequestWithDefaults

`func NewUpdateTaskByIdRequestWithDefaults() *UpdateTaskByIdRequest`

NewUpdateTaskByIdRequestWithDefaults instantiates a new UpdateTaskByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTaskByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTaskByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTaskByIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTaskByIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateTaskByIdRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTaskByIdRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTaskByIdRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTaskByIdRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateTaskByIdRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateTaskByIdRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDueDate

`func (o *UpdateTaskByIdRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *UpdateTaskByIdRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *UpdateTaskByIdRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *UpdateTaskByIdRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *UpdateTaskByIdRequest) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *UpdateTaskByIdRequest) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetStarred

`func (o *UpdateTaskByIdRequest) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *UpdateTaskByIdRequest) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *UpdateTaskByIdRequest) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *UpdateTaskByIdRequest) HasStarred() bool`

HasStarred returns a boolean if a field has been set.

### GetAssigneeId

`func (o *UpdateTaskByIdRequest) GetAssigneeId() int32`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *UpdateTaskByIdRequest) GetAssigneeIdOk() (*int32, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *UpdateTaskByIdRequest) SetAssigneeId(v int32)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *UpdateTaskByIdRequest) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### SetAssigneeIdNil

`func (o *UpdateTaskByIdRequest) SetAssigneeIdNil(b bool)`

 SetAssigneeIdNil sets the value for AssigneeId to be an explicit nil

### UnsetAssigneeId
`func (o *UpdateTaskByIdRequest) UnsetAssigneeId()`

UnsetAssigneeId ensures that no value is present for AssigneeId, not even an explicit nil
### GetSubjectType

`func (o *UpdateTaskByIdRequest) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *UpdateTaskByIdRequest) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *UpdateTaskByIdRequest) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *UpdateTaskByIdRequest) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### SetSubjectTypeNil

`func (o *UpdateTaskByIdRequest) SetSubjectTypeNil(b bool)`

 SetSubjectTypeNil sets the value for SubjectType to be an explicit nil

### UnsetSubjectType
`func (o *UpdateTaskByIdRequest) UnsetSubjectType()`

UnsetSubjectType ensures that no value is present for SubjectType, not even an explicit nil
### GetSubjectId

`func (o *UpdateTaskByIdRequest) GetSubjectId() int32`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *UpdateTaskByIdRequest) GetSubjectIdOk() (*int32, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *UpdateTaskByIdRequest) SetSubjectId(v int32)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *UpdateTaskByIdRequest) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### SetSubjectIdNil

`func (o *UpdateTaskByIdRequest) SetSubjectIdNil(b bool)`

 SetSubjectIdNil sets the value for SubjectId to be an explicit nil

### UnsetSubjectId
`func (o *UpdateTaskByIdRequest) UnsetSubjectId()`

UnsetSubjectId ensures that no value is present for SubjectId, not even an explicit nil
### GetCompleted

`func (o *UpdateTaskByIdRequest) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *UpdateTaskByIdRequest) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *UpdateTaskByIdRequest) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *UpdateTaskByIdRequest) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


