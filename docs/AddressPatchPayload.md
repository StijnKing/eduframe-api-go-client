# AddressPatchPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addressee** | Pointer to **NullableString** | The addressee of the address. | [optional] 
**Address** | Pointer to **string** | Concatenation of the street and house number. | [optional] 
**AddressLine2** | Pointer to **NullableString** | A string representing the second line of the address. | [optional] 
**PostalCode** | Pointer to **string** | A string representing the postal code. | [optional] 
**City** | Pointer to **string** | A string representing the city. | [optional] 
**StateProvince** | Pointer to [**NullableUsaState**](UsaState.md) |  | [optional] 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 

## Methods

### NewAddressPatchPayload

`func NewAddressPatchPayload() *AddressPatchPayload`

NewAddressPatchPayload instantiates a new AddressPatchPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressPatchPayloadWithDefaults

`func NewAddressPatchPayloadWithDefaults() *AddressPatchPayload`

NewAddressPatchPayloadWithDefaults instantiates a new AddressPatchPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressee

`func (o *AddressPatchPayload) GetAddressee() string`

GetAddressee returns the Addressee field if non-nil, zero value otherwise.

### GetAddresseeOk

`func (o *AddressPatchPayload) GetAddresseeOk() (*string, bool)`

GetAddresseeOk returns a tuple with the Addressee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressee

`func (o *AddressPatchPayload) SetAddressee(v string)`

SetAddressee sets Addressee field to given value.

### HasAddressee

`func (o *AddressPatchPayload) HasAddressee() bool`

HasAddressee returns a boolean if a field has been set.

### SetAddresseeNil

`func (o *AddressPatchPayload) SetAddresseeNil(b bool)`

 SetAddresseeNil sets the value for Addressee to be an explicit nil

### UnsetAddressee
`func (o *AddressPatchPayload) UnsetAddressee()`

UnsetAddressee ensures that no value is present for Addressee, not even an explicit nil
### GetAddress

`func (o *AddressPatchPayload) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressPatchPayload) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressPatchPayload) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddressPatchPayload) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressLine2

`func (o *AddressPatchPayload) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AddressPatchPayload) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AddressPatchPayload) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AddressPatchPayload) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *AddressPatchPayload) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *AddressPatchPayload) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *AddressPatchPayload) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressPatchPayload) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressPatchPayload) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AddressPatchPayload) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCity

`func (o *AddressPatchPayload) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressPatchPayload) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressPatchPayload) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressPatchPayload) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStateProvince

`func (o *AddressPatchPayload) GetStateProvince() UsaState`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *AddressPatchPayload) GetStateProvinceOk() (*UsaState, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *AddressPatchPayload) SetStateProvince(v UsaState)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *AddressPatchPayload) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.

### SetStateProvinceNil

`func (o *AddressPatchPayload) SetStateProvinceNil(b bool)`

 SetStateProvinceNil sets the value for StateProvince to be an explicit nil

### UnsetStateProvince
`func (o *AddressPatchPayload) UnsetStateProvince()`

UnsetStateProvince ensures that no value is present for StateProvince, not even an explicit nil
### GetCountry

`func (o *AddressPatchPayload) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressPatchPayload) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressPatchPayload) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AddressPatchPayload) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


