# TransactionSendFungiblesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** |  | 
**FungibleContractAddress** | Pointer to **string** | Ethereum address | [optional] 
**Recipients** | [**[]TransactionSendFungiblesRecipient**](TransactionSendFungiblesRecipient.md) |  | 

## Methods

### NewTransactionSendFungiblesRequest

`func NewTransactionSendFungiblesRequest(network string, recipients []TransactionSendFungiblesRecipient, ) *TransactionSendFungiblesRequest`

NewTransactionSendFungiblesRequest instantiates a new TransactionSendFungiblesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSendFungiblesRequestWithDefaults

`func NewTransactionSendFungiblesRequestWithDefaults() *TransactionSendFungiblesRequest`

NewTransactionSendFungiblesRequestWithDefaults instantiates a new TransactionSendFungiblesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *TransactionSendFungiblesRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionSendFungiblesRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionSendFungiblesRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetFungibleContractAddress

`func (o *TransactionSendFungiblesRequest) GetFungibleContractAddress() string`

GetFungibleContractAddress returns the FungibleContractAddress field if non-nil, zero value otherwise.

### GetFungibleContractAddressOk

`func (o *TransactionSendFungiblesRequest) GetFungibleContractAddressOk() (*string, bool)`

GetFungibleContractAddressOk returns a tuple with the FungibleContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFungibleContractAddress

`func (o *TransactionSendFungiblesRequest) SetFungibleContractAddress(v string)`

SetFungibleContractAddress sets FungibleContractAddress field to given value.

### HasFungibleContractAddress

`func (o *TransactionSendFungiblesRequest) HasFungibleContractAddress() bool`

HasFungibleContractAddress returns a boolean if a field has been set.

### GetRecipients

`func (o *TransactionSendFungiblesRequest) GetRecipients() []TransactionSendFungiblesRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *TransactionSendFungiblesRequest) GetRecipientsOk() (*[]TransactionSendFungiblesRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *TransactionSendFungiblesRequest) SetRecipients(v []TransactionSendFungiblesRecipient)`

SetRecipients sets Recipients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


