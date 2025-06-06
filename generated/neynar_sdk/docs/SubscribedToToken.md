# SubscribedToToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**Address** | **NullableString** |  | 
**Decimals** | **int32** |  | 
**Erc20** | **bool** |  | 

## Methods

### NewSubscribedToToken

`func NewSubscribedToToken(symbol string, address NullableString, decimals int32, erc20 bool, ) *SubscribedToToken`

NewSubscribedToToken instantiates a new SubscribedToToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribedToTokenWithDefaults

`func NewSubscribedToTokenWithDefaults() *SubscribedToToken`

NewSubscribedToTokenWithDefaults instantiates a new SubscribedToToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *SubscribedToToken) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SubscribedToToken) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SubscribedToToken) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetAddress

`func (o *SubscribedToToken) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SubscribedToToken) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SubscribedToToken) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *SubscribedToToken) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *SubscribedToToken) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetDecimals

`func (o *SubscribedToToken) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *SubscribedToToken) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *SubscribedToToken) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetErc20

`func (o *SubscribedToToken) GetErc20() bool`

GetErc20 returns the Erc20 field if non-nil, zero value otherwise.

### GetErc20Ok

`func (o *SubscribedToToken) GetErc20Ok() (*bool, bool)`

GetErc20Ok returns a tuple with the Erc20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErc20

`func (o *SubscribedToToken) SetErc20(v bool)`

SetErc20 sets Erc20 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


