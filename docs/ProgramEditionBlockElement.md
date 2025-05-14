# ProgramEditionBlockElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the element. | 
**Credits** | **NullableFloat64** | The credits of the element. | 
**ElementType** | **string** | The type of the element. | 
**Elements** | [**[]ProgramEditionBlockElementItem**](ProgramEditionBlockElementItem.md) | The elements of the block element. | 

## Methods

### NewProgramEditionBlockElement

`func NewProgramEditionBlockElement(name string, credits NullableFloat64, elementType string, elements []ProgramEditionBlockElementItem, ) *ProgramEditionBlockElement`

NewProgramEditionBlockElement instantiates a new ProgramEditionBlockElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramEditionBlockElementWithDefaults

`func NewProgramEditionBlockElementWithDefaults() *ProgramEditionBlockElement`

NewProgramEditionBlockElementWithDefaults instantiates a new ProgramEditionBlockElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProgramEditionBlockElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProgramEditionBlockElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProgramEditionBlockElement) SetName(v string)`

SetName sets Name field to given value.


### GetCredits

`func (o *ProgramEditionBlockElement) GetCredits() float64`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *ProgramEditionBlockElement) GetCreditsOk() (*float64, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *ProgramEditionBlockElement) SetCredits(v float64)`

SetCredits sets Credits field to given value.


### SetCreditsNil

`func (o *ProgramEditionBlockElement) SetCreditsNil(b bool)`

 SetCreditsNil sets the value for Credits to be an explicit nil

### UnsetCredits
`func (o *ProgramEditionBlockElement) UnsetCredits()`

UnsetCredits ensures that no value is present for Credits, not even an explicit nil
### GetElementType

`func (o *ProgramEditionBlockElement) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *ProgramEditionBlockElement) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *ProgramEditionBlockElement) SetElementType(v string)`

SetElementType sets ElementType field to given value.


### GetElements

`func (o *ProgramEditionBlockElement) GetElements() []ProgramEditionBlockElementItem`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *ProgramEditionBlockElement) GetElementsOk() (*[]ProgramEditionBlockElementItem, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *ProgramEditionBlockElement) SetElements(v []ProgramEditionBlockElementItem)`

SetElements sets Elements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


