# AddressBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**VerifiedAddress** | [**AddressBalanceVerifiedAddress**](AddressBalanceVerifiedAddress.md) |  | 
**TokenBalances** | [**[]TokenBalance**](TokenBalance.md) |  | 

## Methods

### NewAddressBalance

`func NewAddressBalance(object string, verifiedAddress AddressBalanceVerifiedAddress, tokenBalances []TokenBalance, ) *AddressBalance`

NewAddressBalance instantiates a new AddressBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressBalanceWithDefaults

`func NewAddressBalanceWithDefaults() *AddressBalance`

NewAddressBalanceWithDefaults instantiates a new AddressBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *AddressBalance) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AddressBalance) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AddressBalance) SetObject(v string)`

SetObject sets Object field to given value.


### GetVerifiedAddress

`func (o *AddressBalance) GetVerifiedAddress() AddressBalanceVerifiedAddress`

GetVerifiedAddress returns the VerifiedAddress field if non-nil, zero value otherwise.

### GetVerifiedAddressOk

`func (o *AddressBalance) GetVerifiedAddressOk() (*AddressBalanceVerifiedAddress, bool)`

GetVerifiedAddressOk returns a tuple with the VerifiedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAddress

`func (o *AddressBalance) SetVerifiedAddress(v AddressBalanceVerifiedAddress)`

SetVerifiedAddress sets VerifiedAddress field to given value.


### GetTokenBalances

`func (o *AddressBalance) GetTokenBalances() []TokenBalance`

GetTokenBalances returns the TokenBalances field if non-nil, zero value otherwise.

### GetTokenBalancesOk

`func (o *AddressBalance) GetTokenBalancesOk() (*[]TokenBalance, bool)`

GetTokenBalancesOk returns a tuple with the TokenBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBalances

`func (o *AddressBalance) SetTokenBalances(v []TokenBalance)`

SetTokenBalances sets TokenBalances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


