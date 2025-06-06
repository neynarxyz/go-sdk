# LocationAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**StateCode** | Pointer to **string** |  | [optional] 
**Country** | **string** |  | 
**CountryCode** | Pointer to **string** |  | [optional] 

## Methods

### NewLocationAddress

`func NewLocationAddress(city string, country string, ) *LocationAddress`

NewLocationAddress instantiates a new LocationAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAddressWithDefaults

`func NewLocationAddressWithDefaults() *LocationAddress`

NewLocationAddressWithDefaults instantiates a new LocationAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *LocationAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LocationAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LocationAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *LocationAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LocationAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LocationAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LocationAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateCode

`func (o *LocationAddress) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *LocationAddress) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *LocationAddress) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *LocationAddress) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### GetCountry

`func (o *LocationAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LocationAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LocationAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCountryCode

`func (o *LocationAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *LocationAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *LocationAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *LocationAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


