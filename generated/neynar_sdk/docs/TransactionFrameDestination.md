# TransactionFrameDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Ethereum address | 
**Network** | [**Network**](Network.md) |  | 
**TokenContractAddress** | **string** | Token contract address for the payment (e.g. 0x833589fcd6edb6e08f4c7c32d4f71b54bda02913 is USDC on Base) | 
**Amount** | **float32** | Amount to send (must be greater than 0) | 

## Methods

### NewTransactionFrameDestination

`func NewTransactionFrameDestination(address string, network Network, tokenContractAddress string, amount float32, ) *TransactionFrameDestination`

NewTransactionFrameDestination instantiates a new TransactionFrameDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFrameDestinationWithDefaults

`func NewTransactionFrameDestinationWithDefaults() *TransactionFrameDestination`

NewTransactionFrameDestinationWithDefaults instantiates a new TransactionFrameDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TransactionFrameDestination) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionFrameDestination) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionFrameDestination) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNetwork

`func (o *TransactionFrameDestination) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionFrameDestination) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionFrameDestination) SetNetwork(v Network)`

SetNetwork sets Network field to given value.


### GetTokenContractAddress

`func (o *TransactionFrameDestination) GetTokenContractAddress() string`

GetTokenContractAddress returns the TokenContractAddress field if non-nil, zero value otherwise.

### GetTokenContractAddressOk

`func (o *TransactionFrameDestination) GetTokenContractAddressOk() (*string, bool)`

GetTokenContractAddressOk returns a tuple with the TokenContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenContractAddress

`func (o *TransactionFrameDestination) SetTokenContractAddress(v string)`

SetTokenContractAddress sets TokenContractAddress field to given value.


### GetAmount

`func (o *TransactionFrameDestination) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionFrameDestination) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionFrameDestination) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


