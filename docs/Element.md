# Element

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the element. | [readonly] 
**EditionId** | **int32** | The identifier of the associated course. | 
**Position** | **int32** | Sorting position of the element. Lower is higher. | 
**CourseId** | **int32** | The identifier of the associated course. | 
**PlannedCourseId** | **NullableInt32** | The identifier of the associated course. | 

## Methods

### NewElement

`func NewElement(id int32, editionId int32, position int32, courseId int32, plannedCourseId NullableInt32, ) *Element`

NewElement instantiates a new Element object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElementWithDefaults

`func NewElementWithDefaults() *Element`

NewElementWithDefaults instantiates a new Element object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Element) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Element) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Element) SetId(v int32)`

SetId sets Id field to given value.


### GetEditionId

`func (o *Element) GetEditionId() int32`

GetEditionId returns the EditionId field if non-nil, zero value otherwise.

### GetEditionIdOk

`func (o *Element) GetEditionIdOk() (*int32, bool)`

GetEditionIdOk returns a tuple with the EditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionId

`func (o *Element) SetEditionId(v int32)`

SetEditionId sets EditionId field to given value.


### GetPosition

`func (o *Element) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Element) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Element) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetCourseId

`func (o *Element) GetCourseId() int32`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *Element) GetCourseIdOk() (*int32, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *Element) SetCourseId(v int32)`

SetCourseId sets CourseId field to given value.


### GetPlannedCourseId

`func (o *Element) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *Element) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *Element) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.


### SetPlannedCourseIdNil

`func (o *Element) SetPlannedCourseIdNil(b bool)`

 SetPlannedCourseIdNil sets the value for PlannedCourseId to be an explicit nil

### UnsetPlannedCourseId
`func (o *Element) UnsetPlannedCourseId()`

UnsetPlannedCourseId ensures that no value is present for PlannedCourseId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


