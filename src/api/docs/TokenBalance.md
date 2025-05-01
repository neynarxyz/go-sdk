# TokenBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Token** | [**TokenBalanceToken**](TokenBalanceToken.md) |  | 
**Balance** | [**TokenBalanceBalance**](TokenBalanceBalance.md) |  | 

## Methods

### NewTokenBalance

`func NewTokenBalance(object string, token TokenBalanceToken, balance TokenBalanceBalance, ) *TokenBalance`

NewTokenBalance instantiates a new TokenBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenBalanceWithDefaults

`func NewTokenBalanceWithDefaults() *TokenBalance`

NewTokenBalanceWithDefaults instantiates a new TokenBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *TokenBalance) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TokenBalance) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TokenBalance) SetObject(v string)`

SetObject sets Object field to given value.


### GetToken

`func (o *TokenBalance) GetToken() TokenBalanceToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenBalance) GetTokenOk() (*TokenBalanceToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenBalance) SetToken(v TokenBalanceToken)`

SetToken sets Token field to given value.


### GetBalance

`func (o *TokenBalance) GetBalance() TokenBalanceBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *TokenBalance) GetBalanceOk() (*TokenBalanceBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *TokenBalance) SetBalance(v TokenBalanceBalance)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


