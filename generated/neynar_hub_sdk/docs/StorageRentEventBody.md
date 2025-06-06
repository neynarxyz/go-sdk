# StorageRentEventBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payer** | **string** |  | 
**Units** | **int64** |  | 
**Expiry** | **int64** |  | 

## Methods

### NewStorageRentEventBody

`func NewStorageRentEventBody(payer string, units int64, expiry int64, ) *StorageRentEventBody`

NewStorageRentEventBody instantiates a new StorageRentEventBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageRentEventBodyWithDefaults

`func NewStorageRentEventBodyWithDefaults() *StorageRentEventBody`

NewStorageRentEventBodyWithDefaults instantiates a new StorageRentEventBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayer

`func (o *StorageRentEventBody) GetPayer() string`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *StorageRentEventBody) GetPayerOk() (*string, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *StorageRentEventBody) SetPayer(v string)`

SetPayer sets Payer field to given value.


### GetUnits

`func (o *StorageRentEventBody) GetUnits() int64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *StorageRentEventBody) GetUnitsOk() (*int64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *StorageRentEventBody) SetUnits(v int64)`

SetUnits sets Units field to given value.


### GetExpiry

`func (o *StorageRentEventBody) GetExpiry() int64`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *StorageRentEventBody) GetExpiryOk() (*int64, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *StorageRentEventBody) SetExpiry(v int64)`

SetExpiry sets Expiry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


