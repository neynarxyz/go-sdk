# OnChainEventStorageRent

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
**StorageRentEventBody** | [**StorageRentEventBody**](StorageRentEventBody.md) |  | 

## Methods

### NewOnChainEventStorageRent

`func NewOnChainEventStorageRent(type_ string, chainId int32, blockNumber int32, blockHash string, blockTimestamp int32, transactionHash string, logIndex int32, txIndex int32, fid int32, storageRentEventBody StorageRentEventBody, ) *OnChainEventStorageRent`

NewOnChainEventStorageRent instantiates a new OnChainEventStorageRent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnChainEventStorageRentWithDefaults

`func NewOnChainEventStorageRentWithDefaults() *OnChainEventStorageRent`

NewOnChainEventStorageRentWithDefaults instantiates a new OnChainEventStorageRent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OnChainEventStorageRent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnChainEventStorageRent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnChainEventStorageRent) SetType(v string)`

SetType sets Type field to given value.


### GetChainId

`func (o *OnChainEventStorageRent) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *OnChainEventStorageRent) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *OnChainEventStorageRent) SetChainId(v int32)`

SetChainId sets ChainId field to given value.


### GetBlockNumber

`func (o *OnChainEventStorageRent) GetBlockNumber() int32`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *OnChainEventStorageRent) GetBlockNumberOk() (*int32, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *OnChainEventStorageRent) SetBlockNumber(v int32)`

SetBlockNumber sets BlockNumber field to given value.


### GetBlockHash

`func (o *OnChainEventStorageRent) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *OnChainEventStorageRent) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *OnChainEventStorageRent) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockTimestamp

`func (o *OnChainEventStorageRent) GetBlockTimestamp() int32`

GetBlockTimestamp returns the BlockTimestamp field if non-nil, zero value otherwise.

### GetBlockTimestampOk

`func (o *OnChainEventStorageRent) GetBlockTimestampOk() (*int32, bool)`

GetBlockTimestampOk returns a tuple with the BlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimestamp

`func (o *OnChainEventStorageRent) SetBlockTimestamp(v int32)`

SetBlockTimestamp sets BlockTimestamp field to given value.


### GetTransactionHash

`func (o *OnChainEventStorageRent) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *OnChainEventStorageRent) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *OnChainEventStorageRent) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetLogIndex

`func (o *OnChainEventStorageRent) GetLogIndex() int32`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *OnChainEventStorageRent) GetLogIndexOk() (*int32, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *OnChainEventStorageRent) SetLogIndex(v int32)`

SetLogIndex sets LogIndex field to given value.


### GetTxIndex

`func (o *OnChainEventStorageRent) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *OnChainEventStorageRent) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *OnChainEventStorageRent) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.


### GetFid

`func (o *OnChainEventStorageRent) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *OnChainEventStorageRent) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *OnChainEventStorageRent) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetStorageRentEventBody

`func (o *OnChainEventStorageRent) GetStorageRentEventBody() StorageRentEventBody`

GetStorageRentEventBody returns the StorageRentEventBody field if non-nil, zero value otherwise.

### GetStorageRentEventBodyOk

`func (o *OnChainEventStorageRent) GetStorageRentEventBodyOk() (*StorageRentEventBody, bool)`

GetStorageRentEventBodyOk returns a tuple with the StorageRentEventBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRentEventBody

`func (o *OnChainEventStorageRent) SetStorageRentEventBody(v StorageRentEventBody)`

SetStorageRentEventBody sets StorageRentEventBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


