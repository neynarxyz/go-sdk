# BalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserBalance** | Pointer to [**BalanceResponseUserBalance**](BalanceResponseUserBalance.md) |  | [optional] 

## Methods

### NewBalanceResponse

`func NewBalanceResponse() *BalanceResponse`

NewBalanceResponse instantiates a new BalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceResponseWithDefaults

`func NewBalanceResponseWithDefaults() *BalanceResponse`

NewBalanceResponseWithDefaults instantiates a new BalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserBalance

`func (o *BalanceResponse) GetUserBalance() BalanceResponseUserBalance`

GetUserBalance returns the UserBalance field if non-nil, zero value otherwise.

### GetUserBalanceOk

`func (o *BalanceResponse) GetUserBalanceOk() (*BalanceResponseUserBalance, bool)`

GetUserBalanceOk returns a tuple with the UserBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBalance

`func (o *BalanceResponse) SetUserBalance(v BalanceResponseUserBalance)`

SetUserBalance sets UserBalance field to given value.

### HasUserBalance

`func (o *BalanceResponse) HasUserBalance() bool`

HasUserBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


