# CustomAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique identifier of the association. | [readonly] 
**CustomObject** | [**CustomObject**](CustomObject.md) |  | 

## Methods

### NewCustomAssociation

`func NewCustomAssociation(id int32, customObject CustomObject, ) *CustomAssociation`

NewCustomAssociation instantiates a new CustomAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAssociationWithDefaults

`func NewCustomAssociationWithDefaults() *CustomAssociation`

NewCustomAssociationWithDefaults instantiates a new CustomAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomAssociation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAssociation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAssociation) SetId(v int32)`

SetId sets Id field to given value.


### GetCustomObject

`func (o *CustomAssociation) GetCustomObject() CustomObject`

GetCustomObject returns the CustomObject field if non-nil, zero value otherwise.

### GetCustomObjectOk

`func (o *CustomAssociation) GetCustomObjectOk() (*CustomObject, bool)`

GetCustomObjectOk returns a tuple with the CustomObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomObject

`func (o *CustomAssociation) SetCustomObject(v CustomObject)`

SetCustomObject sets CustomObject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


