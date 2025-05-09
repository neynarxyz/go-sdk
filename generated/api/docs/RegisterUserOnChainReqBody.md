# RegisterUserOnChainReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registration** | [**RegisterUserOnChainReqBodyRegistration**](RegisterUserOnChainReqBodyRegistration.md) |  | 
**StorageUnits** | Pointer to **int32** |  | [optional] 
**Signers** | Pointer to [**[]RegisterUserOnChainReqBodySignersInner**](RegisterUserOnChainReqBodySignersInner.md) |  | [optional] 

## Methods

### NewRegisterUserOnChainReqBody

`func NewRegisterUserOnChainReqBody(registration RegisterUserOnChainReqBodyRegistration, ) *RegisterUserOnChainReqBody`

NewRegisterUserOnChainReqBody instantiates a new RegisterUserOnChainReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserOnChainReqBodyWithDefaults

`func NewRegisterUserOnChainReqBodyWithDefaults() *RegisterUserOnChainReqBody`

NewRegisterUserOnChainReqBodyWithDefaults instantiates a new RegisterUserOnChainReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistration

`func (o *RegisterUserOnChainReqBody) GetRegistration() RegisterUserOnChainReqBodyRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *RegisterUserOnChainReqBody) GetRegistrationOk() (*RegisterUserOnChainReqBodyRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *RegisterUserOnChainReqBody) SetRegistration(v RegisterUserOnChainReqBodyRegistration)`

SetRegistration sets Registration field to given value.


### GetStorageUnits

`func (o *RegisterUserOnChainReqBody) GetStorageUnits() int32`

GetStorageUnits returns the StorageUnits field if non-nil, zero value otherwise.

### GetStorageUnitsOk

`func (o *RegisterUserOnChainReqBody) GetStorageUnitsOk() (*int32, bool)`

GetStorageUnitsOk returns a tuple with the StorageUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUnits

`func (o *RegisterUserOnChainReqBody) SetStorageUnits(v int32)`

SetStorageUnits sets StorageUnits field to given value.

### HasStorageUnits

`func (o *RegisterUserOnChainReqBody) HasStorageUnits() bool`

HasStorageUnits returns a boolean if a field has been set.

### GetSigners

`func (o *RegisterUserOnChainReqBody) GetSigners() []RegisterUserOnChainReqBodySignersInner`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *RegisterUserOnChainReqBody) GetSignersOk() (*[]RegisterUserOnChainReqBodySignersInner, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *RegisterUserOnChainReqBody) SetSigners(v []RegisterUserOnChainReqBodySignersInner)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *RegisterUserOnChainReqBody) HasSigners() bool`

HasSigners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


