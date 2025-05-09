# MessageDataReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**MessageType**](MessageType.md) |  | [optional] [default to MESSAGETYPE_MESSAGE_TYPE_CAST_ADD]
**Fid** | **int32** | The unique identifier (FID) of the user who created this message. FIDs are assigned sequentially when users register on the network and cannot be changed. | 
**Timestamp** | **int64** | Seconds since Farcaster Epoch (2021-01-01T00:00:00Z). Used to order messages chronologically and determine the most recent state. Must be within 10 minutes of the current time when the message is created. | 
**Network** | [**FarcasterNetwork**](FarcasterNetwork.md) |  | [default to FARCASTERNETWORK_FARCASTER_NETWORK_MAINNET]
**ReactionBody** | [**ReactionBody**](ReactionBody.md) | Contains the type of reaction (like/recast) and the target content being reacted to. The target can be specified either by castId or URL. | 

## Methods

### NewMessageDataReaction

`func NewMessageDataReaction(fid int32, timestamp int64, network FarcasterNetwork, reactionBody ReactionBody, ) *MessageDataReaction`

NewMessageDataReaction instantiates a new MessageDataReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataReactionWithDefaults

`func NewMessageDataReactionWithDefaults() *MessageDataReaction`

NewMessageDataReactionWithDefaults instantiates a new MessageDataReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageDataReaction) GetType() MessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageDataReaction) GetTypeOk() (*MessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageDataReaction) SetType(v MessageType)`

SetType sets Type field to given value.

### HasType

`func (o *MessageDataReaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFid

`func (o *MessageDataReaction) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataReaction) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataReaction) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetTimestamp

`func (o *MessageDataReaction) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataReaction) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataReaction) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNetwork

`func (o *MessageDataReaction) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataReaction) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataReaction) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.


### GetReactionBody

`func (o *MessageDataReaction) GetReactionBody() ReactionBody`

GetReactionBody returns the ReactionBody field if non-nil, zero value otherwise.

### GetReactionBodyOk

`func (o *MessageDataReaction) GetReactionBodyOk() (*ReactionBody, bool)`

GetReactionBodyOk returns a tuple with the ReactionBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionBody

`func (o *MessageDataReaction) SetReactionBody(v ReactionBody)`

SetReactionBody sets ReactionBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


