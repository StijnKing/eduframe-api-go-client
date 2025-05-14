# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the task. | [readonly] 
**CreatorId** | **int32** | Unique identifier of the user who created the task. | 
**CompletedAt** | **NullableString** | Timestamp of when the task was completed. | 
**CompletedById** | **NullableInt32** | Unique identifier of the user who completed the task. | 
**AssignedById** | **NullableInt32** | Unique identifier of the user who assigned the task. | 
**Name** | **string** | The title of the task. | 
**Description** | **NullableString** | A string representing the description of the task. | 
**DueDate** | **NullableString** | Date when the task must be completed. | 
**Starred** | **bool** | Boolean if the task is starred. | 
**AssigneeId** | **NullableInt32** | Unique identifier of the assigned user for the task. | 
**SubjectType** | **NullableString** | Type of the subject. | 
**SubjectId** | **NullableInt32** | Identifier of the subject. | 
**UpdatedAt** | **time.Time** | Timestamp of last update. | [readonly] 
**CreatedAt** | **time.Time** | Timestamp of creation. | [readonly] 

## Methods

### NewTask

`func NewTask(id int32, creatorId int32, completedAt NullableString, completedById NullableInt32, assignedById NullableInt32, name string, description NullableString, dueDate NullableString, starred bool, assigneeId NullableInt32, subjectType NullableString, subjectId NullableInt32, updatedAt time.Time, createdAt time.Time, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Task) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatorId

`func (o *Task) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Task) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Task) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetCompletedAt

`func (o *Task) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Task) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Task) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *Task) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *Task) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCompletedById

`func (o *Task) GetCompletedById() int32`

GetCompletedById returns the CompletedById field if non-nil, zero value otherwise.

### GetCompletedByIdOk

`func (o *Task) GetCompletedByIdOk() (*int32, bool)`

GetCompletedByIdOk returns a tuple with the CompletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedById

`func (o *Task) SetCompletedById(v int32)`

SetCompletedById sets CompletedById field to given value.


### SetCompletedByIdNil

`func (o *Task) SetCompletedByIdNil(b bool)`

 SetCompletedByIdNil sets the value for CompletedById to be an explicit nil

### UnsetCompletedById
`func (o *Task) UnsetCompletedById()`

UnsetCompletedById ensures that no value is present for CompletedById, not even an explicit nil
### GetAssignedById

`func (o *Task) GetAssignedById() int32`

GetAssignedById returns the AssignedById field if non-nil, zero value otherwise.

### GetAssignedByIdOk

`func (o *Task) GetAssignedByIdOk() (*int32, bool)`

GetAssignedByIdOk returns a tuple with the AssignedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedById

`func (o *Task) SetAssignedById(v int32)`

SetAssignedById sets AssignedById field to given value.


### SetAssignedByIdNil

`func (o *Task) SetAssignedByIdNil(b bool)`

 SetAssignedByIdNil sets the value for AssignedById to be an explicit nil

### UnsetAssignedById
`func (o *Task) UnsetAssignedById()`

UnsetAssignedById ensures that no value is present for AssignedById, not even an explicit nil
### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Task) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Task) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Task) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Task) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Task) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDueDate

`func (o *Task) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Task) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Task) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### SetDueDateNil

`func (o *Task) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *Task) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetStarred

`func (o *Task) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *Task) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *Task) SetStarred(v bool)`

SetStarred sets Starred field to given value.


### GetAssigneeId

`func (o *Task) GetAssigneeId() int32`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *Task) GetAssigneeIdOk() (*int32, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *Task) SetAssigneeId(v int32)`

SetAssigneeId sets AssigneeId field to given value.


### SetAssigneeIdNil

`func (o *Task) SetAssigneeIdNil(b bool)`

 SetAssigneeIdNil sets the value for AssigneeId to be an explicit nil

### UnsetAssigneeId
`func (o *Task) UnsetAssigneeId()`

UnsetAssigneeId ensures that no value is present for AssigneeId, not even an explicit nil
### GetSubjectType

`func (o *Task) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *Task) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *Task) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.


### SetSubjectTypeNil

`func (o *Task) SetSubjectTypeNil(b bool)`

 SetSubjectTypeNil sets the value for SubjectType to be an explicit nil

### UnsetSubjectType
`func (o *Task) UnsetSubjectType()`

UnsetSubjectType ensures that no value is present for SubjectType, not even an explicit nil
### GetSubjectId

`func (o *Task) GetSubjectId() int32`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *Task) GetSubjectIdOk() (*int32, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *Task) SetSubjectId(v int32)`

SetSubjectId sets SubjectId field to given value.


### SetSubjectIdNil

`func (o *Task) SetSubjectIdNil(b bool)`

 SetSubjectIdNil sets the value for SubjectId to be an explicit nil

### UnsetSubjectId
`func (o *Task) UnsetSubjectId()`

UnsetSubjectId ensures that no value is present for SubjectId, not even an explicit nil
### GetUpdatedAt

`func (o *Task) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Task) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Task) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Task) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


