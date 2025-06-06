# OnChainEvent

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
**SignerMigratedEventBody** | [**SignerMigratedEventBody**](SignerMigratedEventBody.md) |  | 
**IdRegisterEventBody** | [**IdRegisterEventBody**](IdRegisterEventBody.md) |  | 
**StorageRentEventBody** | [**StorageRentEventBody**](StorageRentEventBody.md) |  | 

## Methods

### NewOnChainEvent

`func NewOnChainEvent(type_ string, chainId int32, blockNumber int32, blockHash string, blockTimestamp int32, transactionHash string, logIndex int32, txIndex int32, fid int32, signerEventBody SignerEventBody, signerMigratedEventBody SignerMigratedEventBody, idRegisterEventBody IdRegisterEventBody, storageRentEventBody StorageRentEventBody, ) *OnChainEvent`

NewOnChainEvent instantiates a new OnChainEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnChainEventWithDefaults

`func NewOnChainEventWithDefaults() *OnChainEvent`

NewOnChainEventWithDefaults instantiates a new OnChainEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OnChainEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnChainEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnChainEvent) SetType(v string)`

SetType sets Type field to given value.


### GetChainId

`func (o *OnChainEvent) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *OnChainEvent) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *OnChainEvent) SetChainId(v int32)`

SetChainId sets ChainId field to given value.


### GetBlockNumber

`func (o *OnChainEvent) GetBlockNumber() int32`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *OnChainEvent) GetBlockNumberOk() (*int32, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *OnChainEvent) SetBlockNumber(v int32)`

SetBlockNumber sets BlockNumber field to given value.


### GetBlockHash

`func (o *OnChainEvent) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *OnChainEvent) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *OnChainEvent) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockTimestamp

`func (o *OnChainEvent) GetBlockTimestamp() int32`

GetBlockTimestamp returns the BlockTimestamp field if non-nil, zero value otherwise.

### GetBlockTimestampOk

`func (o *OnChainEvent) GetBlockTimestampOk() (*int32, bool)`

GetBlockTimestampOk returns a tuple with the BlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimestamp

`func (o *OnChainEvent) SetBlockTimestamp(v int32)`

SetBlockTimestamp sets BlockTimestamp field to given value.


### GetTransactionHash

`func (o *OnChainEvent) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *OnChainEvent) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *OnChainEvent) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetLogIndex

`func (o *OnChainEvent) GetLogIndex() int32`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *OnChainEvent) GetLogIndexOk() (*int32, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *OnChainEvent) SetLogIndex(v int32)`

SetLogIndex sets LogIndex field to given value.


### GetTxIndex

`func (o *OnChainEvent) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *OnChainEvent) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *OnChainEvent) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.


### GetFid

`func (o *OnChainEvent) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *OnChainEvent) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *OnChainEvent) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetSignerEventBody

`func (o *OnChainEvent) GetSignerEventBody() SignerEventBody`

GetSignerEventBody returns the SignerEventBody field if non-nil, zero value otherwise.

### GetSignerEventBodyOk

`func (o *OnChainEvent) GetSignerEventBodyOk() (*SignerEventBody, bool)`

GetSignerEventBodyOk returns a tuple with the SignerEventBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerEventBody

`func (o *OnChainEvent) SetSignerEventBody(v SignerEventBody)`

SetSignerEventBody sets SignerEventBody field to given value.


### GetSignerMigratedEventBody

`func (o *OnChainEvent) GetSignerMigratedEventBody() SignerMigratedEventBody`

GetSignerMigratedEventBody returns the SignerMigratedEventBody field if non-nil, zero value otherwise.

### GetSignerMigratedEventBodyOk

`func (o *OnChainEvent) GetSignerMigratedEventBodyOk() (*SignerMigratedEventBody, bool)`

GetSignerMigratedEventBodyOk returns a tuple with the SignerMigratedEventBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerMigratedEventBody

`func (o *OnChainEvent) SetSignerMigratedEventBody(v SignerMigratedEventBody)`

SetSignerMigratedEventBody sets SignerMigratedEventBody field to given value.


### GetIdRegisterEventBody

`func (o *OnChainEvent) GetIdRegisterEventBody() IdRegisterEventBody`

GetIdRegisterEventBody returns the IdRegisterEventBody field if non-nil, zero value otherwise.

### GetIdRegisterEventBodyOk

`func (o *OnChainEvent) GetIdRegisterEventBodyOk() (*IdRegisterEventBody, bool)`

GetIdRegisterEventBodyOk returns a tuple with the IdRegisterEventBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdRegisterEventBody

`func (o *OnChainEvent) SetIdRegisterEventBody(v IdRegisterEventBody)`

SetIdRegisterEventBody sets IdRegisterEventBody field to given value.


### GetStorageRentEventBody

`func (o *OnChainEvent) GetStorageRentEventBody() StorageRentEventBody`

GetStorageRentEventBody returns the StorageRentEventBody field if non-nil, zero value otherwise.

### GetStorageRentEventBodyOk

`func (o *OnChainEvent) GetStorageRentEventBodyOk() (*StorageRentEventBody, bool)`

GetStorageRentEventBodyOk returns a tuple with the StorageRentEventBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRentEventBody

`func (o *OnChainEvent) SetStorageRentEventBody(v StorageRentEventBody)`

SetStorageRentEventBody sets StorageRentEventBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


