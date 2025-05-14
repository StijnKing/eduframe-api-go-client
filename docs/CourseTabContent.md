# CourseTabContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the course tab content. | [readonly] 
**CourseTabId** | **int32** | Unique identifier of the course tab. | [readonly] 
**Content** | **NullableString** | The HTML content of the course tab. | 

## Methods

### NewCourseTabContent

`func NewCourseTabContent(id int32, courseTabId int32, content NullableString, ) *CourseTabContent`

NewCourseTabContent instantiates a new CourseTabContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseTabContentWithDefaults

`func NewCourseTabContentWithDefaults() *CourseTabContent`

NewCourseTabContentWithDefaults instantiates a new CourseTabContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseTabContent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseTabContent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseTabContent) SetId(v int32)`

SetId sets Id field to given value.


### GetCourseTabId

`func (o *CourseTabContent) GetCourseTabId() int32`

GetCourseTabId returns the CourseTabId field if non-nil, zero value otherwise.

### GetCourseTabIdOk

`func (o *CourseTabContent) GetCourseTabIdOk() (*int32, bool)`

GetCourseTabIdOk returns a tuple with the CourseTabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseTabId

`func (o *CourseTabContent) SetCourseTabId(v int32)`

SetCourseTabId sets CourseTabId field to given value.


### GetContent

`func (o *CourseTabContent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CourseTabContent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CourseTabContent) SetContent(v string)`

SetContent sets Content field to given value.


### SetContentNil

`func (o *CourseTabContent) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CourseTabContent) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


