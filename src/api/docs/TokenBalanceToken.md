# TokenBalanceToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Name** | **string** | The token name e.g. \&quot;Ethereum\&quot; | 
**Symbol** | **string** | The token symbol e.g. \&quot;ETH\&quot; | 
**Address** | Pointer to **string** | The contract address of the token (omitted for native token) | [optional] 
**Decimals** | Pointer to **int32** | The number of decimals the token uses | [optional] 

## Methods

### NewTokenBalanceToken

`func NewTokenBalanceToken(object string, name string, symbol string, ) *TokenBalanceToken`

NewTokenBalanceToken instantiates a new TokenBalanceToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenBalanceTokenWithDefaults

`func NewTokenBalanceTokenWithDefaults() *TokenBalanceToken`

NewTokenBalanceTokenWithDefaults instantiates a new TokenBalanceToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *TokenBalanceToken) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TokenBalanceToken) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TokenBalanceToken) SetObject(v string)`

SetObject sets Object field to given value.


### GetName

`func (o *TokenBalanceToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenBalanceToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenBalanceToken) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *TokenBalanceToken) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenBalanceToken) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenBalanceToken) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetAddress

`func (o *TokenBalanceToken) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenBalanceToken) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenBalanceToken) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TokenBalanceToken) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDecimals

`func (o *TokenBalanceToken) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenBalanceToken) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenBalanceToken) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *TokenBalanceToken) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


