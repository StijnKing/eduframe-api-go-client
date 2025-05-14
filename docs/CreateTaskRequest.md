# CreateTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The title of the task. | 
**Description** | Pointer to **NullableString** | A string representing the description of the task. | [optional] 
**DueDate** | Pointer to **NullableString** | Date when the task must be completed. | [optional] 
**Starred** | Pointer to **bool** | Boolean if the task is starred. | [optional] 
**AssigneeId** | Pointer to **NullableInt32** | Unique identifier of the assigned user for the task. | [optional] 
**SubjectType** | Pointer to **NullableString** | Type of the subject. | [optional] 
**SubjectId** | Pointer to **NullableInt32** | Identifier of the subject. | [optional] 
**Completed** | Pointer to **bool** | Boolean representing the status of the task. The default value is false. | [optional] 

## Methods

### NewCreateTaskRequest

`func NewCreateTaskRequest(name string, ) *CreateTaskRequest`

NewCreateTaskRequest instantiates a new CreateTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaskRequestWithDefaults

`func NewCreateTaskRequestWithDefaults() *CreateTaskRequest`

NewCreateTaskRequestWithDefaults instantiates a new CreateTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTaskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTaskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTaskRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTaskRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTaskRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDueDate

`func (o *CreateTaskRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CreateTaskRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CreateTaskRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *CreateTaskRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *CreateTaskRequest) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *CreateTaskRequest) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetStarred

`func (o *CreateTaskRequest) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *CreateTaskRequest) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *CreateTaskRequest) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *CreateTaskRequest) HasStarred() bool`

HasStarred returns a boolean if a field has been set.

### GetAssigneeId

`func (o *CreateTaskRequest) GetAssigneeId() int32`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *CreateTaskRequest) GetAssigneeIdOk() (*int32, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *CreateTaskRequest) SetAssigneeId(v int32)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *CreateTaskRequest) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### SetAssigneeIdNil

`func (o *CreateTaskRequest) SetAssigneeIdNil(b bool)`

 SetAssigneeIdNil sets the value for AssigneeId to be an explicit nil

### UnsetAssigneeId
`func (o *CreateTaskRequest) UnsetAssigneeId()`

UnsetAssigneeId ensures that no value is present for AssigneeId, not even an explicit nil
### GetSubjectType

`func (o *CreateTaskRequest) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *CreateTaskRequest) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *CreateTaskRequest) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *CreateTaskRequest) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### SetSubjectTypeNil

`func (o *CreateTaskRequest) SetSubjectTypeNil(b bool)`

 SetSubjectTypeNil sets the value for SubjectType to be an explicit nil

### UnsetSubjectType
`func (o *CreateTaskRequest) UnsetSubjectType()`

UnsetSubjectType ensures that no value is present for SubjectType, not even an explicit nil
### GetSubjectId

`func (o *CreateTaskRequest) GetSubjectId() int32`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *CreateTaskRequest) GetSubjectIdOk() (*int32, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *CreateTaskRequest) SetSubjectId(v int32)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *CreateTaskRequest) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### SetSubjectIdNil

`func (o *CreateTaskRequest) SetSubjectIdNil(b bool)`

 SetSubjectIdNil sets the value for SubjectId to be an explicit nil

### UnsetSubjectId
`func (o *CreateTaskRequest) UnsetSubjectId()`

UnsetSubjectId ensures that no value is present for SubjectId, not even an explicit nil
### GetCompleted

`func (o *CreateTaskRequest) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *CreateTaskRequest) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *CreateTaskRequest) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *CreateTaskRequest) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


