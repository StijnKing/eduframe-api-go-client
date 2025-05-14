# AddressPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addressee** | Pointer to **NullableString** | The addressee of the address. | [optional] 
**Address** | **string** | Concatenation of the street and house number. | 
**AddressLine2** | Pointer to **NullableString** | A string representing the second line of the address. | [optional] 
**PostalCode** | **string** | A string representing the postal code. | 
**City** | **string** | A string representing the city. | 
**StateProvince** | Pointer to [**NullableUsaState**](UsaState.md) |  | [optional] 
**Country** | [**Country**](Country.md) |  | 

## Methods

### NewAddressPayload

`func NewAddressPayload(address string, postalCode string, city string, country Country, ) *AddressPayload`

NewAddressPayload instantiates a new AddressPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressPayloadWithDefaults

`func NewAddressPayloadWithDefaults() *AddressPayload`

NewAddressPayloadWithDefaults instantiates a new AddressPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressee

`func (o *AddressPayload) GetAddressee() string`

GetAddressee returns the Addressee field if non-nil, zero value otherwise.

### GetAddresseeOk

`func (o *AddressPayload) GetAddresseeOk() (*string, bool)`

GetAddresseeOk returns a tuple with the Addressee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressee

`func (o *AddressPayload) SetAddressee(v string)`

SetAddressee sets Addressee field to given value.

### HasAddressee

`func (o *AddressPayload) HasAddressee() bool`

HasAddressee returns a boolean if a field has been set.

### SetAddresseeNil

`func (o *AddressPayload) SetAddresseeNil(b bool)`

 SetAddresseeNil sets the value for Addressee to be an explicit nil

### UnsetAddressee
`func (o *AddressPayload) UnsetAddressee()`

UnsetAddressee ensures that no value is present for Addressee, not even an explicit nil
### GetAddress

`func (o *AddressPayload) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressPayload) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressPayload) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressLine2

`func (o *AddressPayload) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AddressPayload) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AddressPayload) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AddressPayload) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *AddressPayload) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *AddressPayload) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *AddressPayload) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressPayload) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressPayload) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCity

`func (o *AddressPayload) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressPayload) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressPayload) SetCity(v string)`

SetCity sets City field to given value.


### GetStateProvince

`func (o *AddressPayload) GetStateProvince() UsaState`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *AddressPayload) GetStateProvinceOk() (*UsaState, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *AddressPayload) SetStateProvince(v UsaState)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *AddressPayload) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.

### SetStateProvinceNil

`func (o *AddressPayload) SetStateProvinceNil(b bool)`

 SetStateProvinceNil sets the value for StateProvince to be an explicit nil

### UnsetStateProvince
`func (o *AddressPayload) UnsetStateProvince()`

UnsetStateProvince ensures that no value is present for StateProvince, not even an explicit nil
### GetCountry

`func (o *AddressPayload) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressPayload) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressPayload) SetCountry(v Country)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


