# TransactionSendTxInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** |  | 
**TransactionHash** | **string** |  | 
**GasUsed** | **string** | Gas used for the transaction. | 
**ApprovalHash** | **string** | Hash of the transaction that approved the transfer. This is only present if the fungible token is not native token of the network. | 

## Methods

### NewTransactionSendTxInfo

`func NewTransactionSendTxInfo(network string, transactionHash string, gasUsed string, approvalHash string, ) *TransactionSendTxInfo`

NewTransactionSendTxInfo instantiates a new TransactionSendTxInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSendTxInfoWithDefaults

`func NewTransactionSendTxInfoWithDefaults() *TransactionSendTxInfo`

NewTransactionSendTxInfoWithDefaults instantiates a new TransactionSendTxInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *TransactionSendTxInfo) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionSendTxInfo) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionSendTxInfo) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTransactionHash

`func (o *TransactionSendTxInfo) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *TransactionSendTxInfo) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *TransactionSendTxInfo) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetGasUsed

`func (o *TransactionSendTxInfo) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *TransactionSendTxInfo) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *TransactionSendTxInfo) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetApprovalHash

`func (o *TransactionSendTxInfo) GetApprovalHash() string`

GetApprovalHash returns the ApprovalHash field if non-nil, zero value otherwise.

### GetApprovalHashOk

`func (o *TransactionSendTxInfo) GetApprovalHashOk() (*string, bool)`

GetApprovalHashOk returns a tuple with the ApprovalHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalHash

`func (o *TransactionSendTxInfo) SetApprovalHash(v string)`

SetApprovalHash sets ApprovalHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


