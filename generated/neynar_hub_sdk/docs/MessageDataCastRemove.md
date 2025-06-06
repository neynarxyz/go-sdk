# MessageDataCastRemove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**MessageType**](MessageType.md) |  | [optional] [default to MESSAGETYPE_MESSAGE_TYPE_CAST_ADD]
**Fid** | **int32** | The unique identifier (FID) of the user who created this message. FIDs are assigned sequentially when users register on the network and cannot be changed. | 
**Timestamp** | **int64** | Seconds since Farcaster Epoch (2021-01-01T00:00:00Z). Used to order messages chronologically and determine the most recent state. Must be within 10 minutes of the current time when the message is created. | 
**Network** | [**FarcasterNetwork**](FarcasterNetwork.md) |  | [default to FARCASTERNETWORK_FARCASTER_NETWORK_MAINNET]
**TargetHash** | **string** | The unique hash identifier of the cast to be removed. Must be a cast that was previously created by the same FID specified in the message. | 

## Methods

### NewMessageDataCastRemove

`func NewMessageDataCastRemove(fid int32, timestamp int64, network FarcasterNetwork, targetHash string, ) *MessageDataCastRemove`

NewMessageDataCastRemove instantiates a new MessageDataCastRemove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataCastRemoveWithDefaults

`func NewMessageDataCastRemoveWithDefaults() *MessageDataCastRemove`

NewMessageDataCastRemoveWithDefaults instantiates a new MessageDataCastRemove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageDataCastRemove) GetType() MessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageDataCastRemove) GetTypeOk() (*MessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageDataCastRemove) SetType(v MessageType)`

SetType sets Type field to given value.

### HasType

`func (o *MessageDataCastRemove) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFid

`func (o *MessageDataCastRemove) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataCastRemove) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataCastRemove) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetTimestamp

`func (o *MessageDataCastRemove) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataCastRemove) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataCastRemove) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNetwork

`func (o *MessageDataCastRemove) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataCastRemove) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataCastRemove) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.


### GetTargetHash

`func (o *MessageDataCastRemove) GetTargetHash() string`

GetTargetHash returns the TargetHash field if non-nil, zero value otherwise.

### GetTargetHashOk

`func (o *MessageDataCastRemove) GetTargetHashOk() (*string, bool)`

GetTargetHashOk returns a tuple with the TargetHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHash

`func (o *MessageDataCastRemove) SetTargetHash(v string)`

SetTargetHash sets TargetHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


