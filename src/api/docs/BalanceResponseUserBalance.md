# BalanceResponseUserBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**User** | [**UserDehydrated**](UserDehydrated.md) |  | 
**AddressBalances** | [**[]AddressBalance**](AddressBalance.md) |  | 

## Methods

### NewBalanceResponseUserBalance

`func NewBalanceResponseUserBalance(object string, user UserDehydrated, addressBalances []AddressBalance, ) *BalanceResponseUserBalance`

NewBalanceResponseUserBalance instantiates a new BalanceResponseUserBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceResponseUserBalanceWithDefaults

`func NewBalanceResponseUserBalanceWithDefaults() *BalanceResponseUserBalance`

NewBalanceResponseUserBalanceWithDefaults instantiates a new BalanceResponseUserBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *BalanceResponseUserBalance) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BalanceResponseUserBalance) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BalanceResponseUserBalance) SetObject(v string)`

SetObject sets Object field to given value.


### GetUser

`func (o *BalanceResponseUserBalance) GetUser() UserDehydrated`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BalanceResponseUserBalance) GetUserOk() (*UserDehydrated, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BalanceResponseUserBalance) SetUser(v UserDehydrated)`

SetUser sets User field to given value.


### GetAddressBalances

`func (o *BalanceResponseUserBalance) GetAddressBalances() []AddressBalance`

GetAddressBalances returns the AddressBalances field if non-nil, zero value otherwise.

### GetAddressBalancesOk

`func (o *BalanceResponseUserBalance) GetAddressBalancesOk() (*[]AddressBalance, bool)`

GetAddressBalancesOk returns a tuple with the AddressBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBalances

`func (o *BalanceResponseUserBalance) SetAddressBalances(v []AddressBalance)`

SetAddressBalances sets AddressBalances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


