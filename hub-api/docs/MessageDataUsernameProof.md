# MessageDataUsernameProof

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | The unique identifier (FID) of the user who created this message. FIDs are assigned sequentially when users register on the network and cannot be changed. | 
**Timestamp** | **int64** | Seconds since Farcaster Epoch (2021-01-01T00:00:00Z). Used to order messages chronologically and determine the most recent state. Must be within 10 minutes of the current time when the message is created. | 
**Network** | [**FarcasterNetwork**](FarcasterNetwork.md) |  | [default to FARCASTER_NETWORK_MAINNET]
**UsernameProofBody** | [**UserNameProof**](UserNameProof.md) |  | 

## Methods

### NewMessageDataUsernameProof

`func NewMessageDataUsernameProof(fid int32, timestamp int64, network FarcasterNetwork, usernameProofBody UserNameProof, ) *MessageDataUsernameProof`

NewMessageDataUsernameProof instantiates a new MessageDataUsernameProof object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataUsernameProofWithDefaults

`func NewMessageDataUsernameProofWithDefaults() *MessageDataUsernameProof`

NewMessageDataUsernameProofWithDefaults instantiates a new MessageDataUsernameProof object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageDataUsernameProof) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataUsernameProof) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataUsernameProof) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetTimestamp

`func (o *MessageDataUsernameProof) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataUsernameProof) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataUsernameProof) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNetwork

`func (o *MessageDataUsernameProof) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataUsernameProof) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataUsernameProof) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.


### GetUsernameProofBody

`func (o *MessageDataUsernameProof) GetUsernameProofBody() UserNameProof`

GetUsernameProofBody returns the UsernameProofBody field if non-nil, zero value otherwise.

### GetUsernameProofBodyOk

`func (o *MessageDataUsernameProof) GetUsernameProofBodyOk() (*UserNameProof, bool)`

GetUsernameProofBodyOk returns a tuple with the UsernameProofBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameProofBody

`func (o *MessageDataUsernameProof) SetUsernameProofBody(v UserNameProof)`

SetUsernameProofBody sets UsernameProofBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


