# MessageDataFrameAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | The unique identifier (FID) of the user who created this message. FIDs are assigned sequentially when users register on the network and cannot be changed. | 
**Timestamp** | **int64** | Seconds since Farcaster Epoch (2021-01-01T00:00:00Z). Used to order messages chronologically and determine the most recent state. Must be within 10 minutes of the current time when the message is created. | 
**Network** | [**FarcasterNetwork**](FarcasterNetwork.md) |  | [default to FARCASTERNETWORK_FARCASTER_NETWORK_MAINNET]
**FrameActionBody** | [**FrameActionBody**](FrameActionBody.md) | Contains the details of the frame interaction, including which button was pressed and the associated cast and URL. | 

## Methods

### NewMessageDataFrameAction

`func NewMessageDataFrameAction(fid int32, timestamp int64, network FarcasterNetwork, frameActionBody FrameActionBody, ) *MessageDataFrameAction`

NewMessageDataFrameAction instantiates a new MessageDataFrameAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataFrameActionWithDefaults

`func NewMessageDataFrameActionWithDefaults() *MessageDataFrameAction`

NewMessageDataFrameActionWithDefaults instantiates a new MessageDataFrameAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageDataFrameAction) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataFrameAction) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataFrameAction) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetTimestamp

`func (o *MessageDataFrameAction) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataFrameAction) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataFrameAction) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNetwork

`func (o *MessageDataFrameAction) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataFrameAction) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataFrameAction) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.


### GetFrameActionBody

`func (o *MessageDataFrameAction) GetFrameActionBody() FrameActionBody`

GetFrameActionBody returns the FrameActionBody field if non-nil, zero value otherwise.

### GetFrameActionBodyOk

`func (o *MessageDataFrameAction) GetFrameActionBodyOk() (*FrameActionBody, bool)`

GetFrameActionBodyOk returns a tuple with the FrameActionBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameActionBody

`func (o *MessageDataFrameAction) SetFrameActionBody(v FrameActionBody)`

SetFrameActionBody sets FrameActionBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


