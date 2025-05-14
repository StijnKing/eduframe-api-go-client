# CreateTeacher201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | The result of the operation. | [optional] 
**Location** | **string** | The URI of the created resource. | 

## Methods

### NewCreateTeacher201Response

`func NewCreateTeacher201Response(location string, ) *CreateTeacher201Response`

NewCreateTeacher201Response instantiates a new CreateTeacher201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTeacher201ResponseWithDefaults

`func NewCreateTeacher201ResponseWithDefaults() *CreateTeacher201Response`

NewCreateTeacher201ResponseWithDefaults instantiates a new CreateTeacher201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CreateTeacher201Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateTeacher201Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateTeacher201Response) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CreateTeacher201Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetLocation

`func (o *CreateTeacher201Response) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateTeacher201Response) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateTeacher201Response) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


