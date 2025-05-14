# ProgramEditionThesisElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the element. | 
**Credits** | **NullableFloat64** | The credits of the element. | 
**ElementType** | **string** | The type of the element. | 
**CourseId** | **int32** | The identifier of the associated course. | 
**PlannedCourseId** | **NullableInt32** | The identifier of the associated planned course. | 

## Methods

### NewProgramEditionThesisElement

`func NewProgramEditionThesisElement(name string, credits NullableFloat64, elementType string, courseId int32, plannedCourseId NullableInt32, ) *ProgramEditionThesisElement`

NewProgramEditionThesisElement instantiates a new ProgramEditionThesisElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramEditionThesisElementWithDefaults

`func NewProgramEditionThesisElementWithDefaults() *ProgramEditionThesisElement`

NewProgramEditionThesisElementWithDefaults instantiates a new ProgramEditionThesisElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProgramEditionThesisElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProgramEditionThesisElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProgramEditionThesisElement) SetName(v string)`

SetName sets Name field to given value.


### GetCredits

`func (o *ProgramEditionThesisElement) GetCredits() float64`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *ProgramEditionThesisElement) GetCreditsOk() (*float64, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *ProgramEditionThesisElement) SetCredits(v float64)`

SetCredits sets Credits field to given value.


### SetCreditsNil

`func (o *ProgramEditionThesisElement) SetCreditsNil(b bool)`

 SetCreditsNil sets the value for Credits to be an explicit nil

### UnsetCredits
`func (o *ProgramEditionThesisElement) UnsetCredits()`

UnsetCredits ensures that no value is present for Credits, not even an explicit nil
### GetElementType

`func (o *ProgramEditionThesisElement) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *ProgramEditionThesisElement) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *ProgramEditionThesisElement) SetElementType(v string)`

SetElementType sets ElementType field to given value.


### GetCourseId

`func (o *ProgramEditionThesisElement) GetCourseId() int32`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *ProgramEditionThesisElement) GetCourseIdOk() (*int32, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *ProgramEditionThesisElement) SetCourseId(v int32)`

SetCourseId sets CourseId field to given value.


### GetPlannedCourseId

`func (o *ProgramEditionThesisElement) GetPlannedCourseId() int32`

GetPlannedCourseId returns the PlannedCourseId field if non-nil, zero value otherwise.

### GetPlannedCourseIdOk

`func (o *ProgramEditionThesisElement) GetPlannedCourseIdOk() (*int32, bool)`

GetPlannedCourseIdOk returns a tuple with the PlannedCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCourseId

`func (o *ProgramEditionThesisElement) SetPlannedCourseId(v int32)`

SetPlannedCourseId sets PlannedCourseId field to given value.


### SetPlannedCourseIdNil

`func (o *ProgramEditionThesisElement) SetPlannedCourseIdNil(b bool)`

 SetPlannedCourseIdNil sets the value for PlannedCourseId to be an explicit nil

### UnsetPlannedCourseId
`func (o *ProgramEditionThesisElement) UnsetPlannedCourseId()`

UnsetPlannedCourseId ensures that no value is present for PlannedCourseId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


