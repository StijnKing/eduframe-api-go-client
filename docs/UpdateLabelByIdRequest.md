# UpdateLabelByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the label | [optional] 
**Color** | Pointer to **string** | Hex code of the color of the label | [optional] 
**ModelType** | Pointer to **string** | The model type for which this label is made available | [optional] 

## Methods

### NewUpdateLabelByIdRequest

`func NewUpdateLabelByIdRequest() *UpdateLabelByIdRequest`

NewUpdateLabelByIdRequest instantiates a new UpdateLabelByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLabelByIdRequestWithDefaults

`func NewUpdateLabelByIdRequestWithDefaults() *UpdateLabelByIdRequest`

NewUpdateLabelByIdRequestWithDefaults instantiates a new UpdateLabelByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateLabelByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLabelByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLabelByIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLabelByIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColor

`func (o *UpdateLabelByIdRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateLabelByIdRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateLabelByIdRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *UpdateLabelByIdRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetModelType

`func (o *UpdateLabelByIdRequest) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *UpdateLabelByIdRequest) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *UpdateLabelByIdRequest) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *UpdateLabelByIdRequest) HasModelType() bool`

HasModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


