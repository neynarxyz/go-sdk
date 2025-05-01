# MessageDataCastAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | The unique identifier (FID) of the user who created this message. FIDs are assigned sequentially when users register on the network and cannot be changed. | 
**Timestamp** | **int64** | Seconds since Farcaster Epoch (2021-01-01T00:00:00Z). Used to order messages chronologically and determine the most recent state. Must be within 10 minutes of the current time when the message is created. | 
**Network** | [**FarcasterNetwork**](FarcasterNetwork.md) |  | [default to MAINNET]
**CastAddBody** | [**CastAddBody**](CastAddBody.md) | The content and metadata of the new cast, including the text, mentions, embeds, and any parent references for replies. | 

## Methods

### NewMessageDataCastAdd

`func NewMessageDataCastAdd(fid int32, timestamp int64, network FarcasterNetwork, castAddBody CastAddBody, ) *MessageDataCastAdd`

NewMessageDataCastAdd instantiates a new MessageDataCastAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataCastAddWithDefaults

`func NewMessageDataCastAddWithDefaults() *MessageDataCastAdd`

NewMessageDataCastAddWithDefaults instantiates a new MessageDataCastAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageDataCastAdd) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataCastAdd) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataCastAdd) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetTimestamp

`func (o *MessageDataCastAdd) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataCastAdd) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataCastAdd) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNetwork

`func (o *MessageDataCastAdd) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataCastAdd) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataCastAdd) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.


### GetCastAddBody

`func (o *MessageDataCastAdd) GetCastAddBody() CastAddBody`

GetCastAddBody returns the CastAddBody field if non-nil, zero value otherwise.

### GetCastAddBodyOk

`func (o *MessageDataCastAdd) GetCastAddBodyOk() (*CastAddBody, bool)`

GetCastAddBodyOk returns a tuple with the CastAddBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastAddBody

`func (o *MessageDataCastAdd) SetCastAddBody(v CastAddBody)`

SetCastAddBody sets CastAddBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


