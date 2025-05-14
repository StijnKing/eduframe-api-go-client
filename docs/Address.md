# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addressee** | **NullableString** | The addressee of the address. | 
**Address** | **string** | Concatenation of the street and house number. | 
**AddressLine2** | **NullableString** | A string representing the second line of the address. | 
**PostalCode** | **string** | A string representing the postal code. | 
**City** | **string** | A string representing the city. | 
**Country** | [**Country**](Country.md) |  | 
**StateProvince** | [**NullableUsaState**](UsaState.md) |  | 

## Methods

### NewAddress

`func NewAddress(addressee NullableString, address string, addressLine2 NullableString, postalCode string, city string, country Country, stateProvince NullableUsaState, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressee

`func (o *Address) GetAddressee() string`

GetAddressee returns the Addressee field if non-nil, zero value otherwise.

### GetAddresseeOk

`func (o *Address) GetAddresseeOk() (*string, bool)`

GetAddresseeOk returns a tuple with the Addressee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressee

`func (o *Address) SetAddressee(v string)`

SetAddressee sets Addressee field to given value.


### SetAddresseeNil

`func (o *Address) SetAddresseeNil(b bool)`

 SetAddresseeNil sets the value for Addressee to be an explicit nil

### UnsetAddressee
`func (o *Address) UnsetAddressee()`

UnsetAddressee ensures that no value is present for Addressee, not even an explicit nil
### GetAddress

`func (o *Address) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Address) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Address) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressLine2

`func (o *Address) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Address) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Address) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.


### SetAddressLine2Nil

`func (o *Address) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *Address) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *Address) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v Country)`

SetCountry sets Country field to given value.


### GetStateProvince

`func (o *Address) GetStateProvince() UsaState`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *Address) GetStateProvinceOk() (*UsaState, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *Address) SetStateProvince(v UsaState)`

SetStateProvince sets StateProvince field to given value.


### SetStateProvinceNil

`func (o *Address) SetStateProvinceNil(b bool)`

 SetStateProvinceNil sets the value for StateProvince to be an explicit nil

### UnsetStateProvince
`func (o *Address) UnsetStateProvince()`

UnsetStateProvince ensures that no value is present for StateProvince, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


