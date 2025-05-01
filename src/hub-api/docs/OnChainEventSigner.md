# OnChainEventSigner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ChainId** | **int32** |  | 
**BlockNumber** | **int32** |  | 
**BlockHash** | **string** |  | 
**BlockTimestamp** | **int32** |  | 
**TransactionHash** | **string** |  | 
**LogIndex** | **int32** |  | 
**TxIndex** | **int32** |  | 
**Fid** | **int32** |  | 
**SignerEventBody** | [**SignerEventBody**](SignerEventBody.md) |  | 

## Methods

### NewOnChainEventSigner

`func NewOnChainEventSigner(type_ string, chainId int32, blockNumber int32, blockHash string, blockTimestamp int32, transactionHash string, logIndex int32, txIndex int32, fid int32, signerEventBody SignerEventBody, ) *OnChainEventSigner`

NewOnChainEventSigner instantiates a new OnChainEventSigner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnChainEventSignerWithDefaults

`func NewOnChainEventSignerWithDefaults() *OnChainEventSigner`

NewOnChainEventSignerWithDefaults instantiates a new OnChainEventSigner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OnChainEventSigner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnChainEventSigner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnChainEventSigner) SetType(v string)`

SetType sets Type field to given value.


### GetChainId

`func (o *OnChainEventSigner) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *OnChainEventSigner) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *OnChainEventSigner) SetChainId(v int32)`

SetChainId sets ChainId field to given value.


### GetBlockNumber

`func (o *OnChainEventSigner) GetBlockNumber() int32`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *OnChainEventSigner) GetBlockNumberOk() (*int32, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *OnChainEventSigner) SetBlockNumber(v int32)`

SetBlockNumber sets BlockNumber field to given value.


### GetBlockHash

`func (o *OnChainEventSigner) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *OnChainEventSigner) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *OnChainEventSigner) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockTimestamp

`func (o *OnChainEventSigner) GetBlockTimestamp() int32`

GetBlockTimestamp returns the BlockTimestamp field if non-nil, zero value otherwise.

### GetBlockTimestampOk

`func (o *OnChainEventSigner) GetBlockTimestampOk() (*int32, bool)`

GetBlockTimestampOk returns a tuple with the BlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimestamp

`func (o *OnChainEventSigner) SetBlockTimestamp(v int32)`

SetBlockTimestamp sets BlockTimestamp field to given value.


### GetTransactionHash

`func (o *OnChainEventSigner) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *OnChainEventSigner) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *OnChainEventSigner) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetLogIndex

`func (o *OnChainEventSigner) GetLogIndex() int32`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *OnChainEventSigner) GetLogIndexOk() (*int32, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *OnChainEventSigner) SetLogIndex(v int32)`

SetLogIndex sets LogIndex field to given value.


### GetTxIndex

`func (o *OnChainEventSigner) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *OnChainEventSigner) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *OnChainEventSigner) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.


### GetFid

`func (o *OnChainEventSigner) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *OnChainEventSigner) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *OnChainEventSigner) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetSignerEventBody

`func (o *OnChainEventSigner) GetSignerEventBody() SignerEventBody`

GetSignerEventBody returns the SignerEventBody field if non-nil, zero value otherwise.

### GetSignerEventBodyOk

`func (o *OnChainEventSigner) GetSignerEventBodyOk() (*SignerEventBody, bool)`

GetSignerEventBodyOk returns a tuple with the SignerEventBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerEventBody

`func (o *OnChainEventSigner) SetSignerEventBody(v SignerEventBody)`

SetSignerEventBody sets SignerEventBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


