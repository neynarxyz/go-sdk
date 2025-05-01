# MessageDataVerificationRemove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | The unique identifier (FID) of the user who created this message. FIDs are assigned sequentially when users register on the network and cannot be changed. | 
**Timestamp** | **int64** | Seconds since Farcaster Epoch (2021-01-01T00:00:00Z). Used to order messages chronologically and determine the most recent state. Must be within 10 minutes of the current time when the message is created. | 
**Network** | [**FarcasterNetwork**](FarcasterNetwork.md) |  | [default to MAINNET]
**VerificationRemoveBody** | [**VerificationRemoveBody**](VerificationRemoveBody.md) | Contains the blockchain address for which the verification should be removed from the user&#39;s profile. | 

## Methods

### NewMessageDataVerificationRemove

`func NewMessageDataVerificationRemove(fid int32, timestamp int64, network FarcasterNetwork, verificationRemoveBody VerificationRemoveBody, ) *MessageDataVerificationRemove`

NewMessageDataVerificationRemove instantiates a new MessageDataVerificationRemove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataVerificationRemoveWithDefaults

`func NewMessageDataVerificationRemoveWithDefaults() *MessageDataVerificationRemove`

NewMessageDataVerificationRemoveWithDefaults instantiates a new MessageDataVerificationRemove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageDataVerificationRemove) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataVerificationRemove) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataVerificationRemove) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetTimestamp

`func (o *MessageDataVerificationRemove) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataVerificationRemove) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataVerificationRemove) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNetwork

`func (o *MessageDataVerificationRemove) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataVerificationRemove) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataVerificationRemove) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.


### GetVerificationRemoveBody

`func (o *MessageDataVerificationRemove) GetVerificationRemoveBody() VerificationRemoveBody`

GetVerificationRemoveBody returns the VerificationRemoveBody field if non-nil, zero value otherwise.

### GetVerificationRemoveBodyOk

`func (o *MessageDataVerificationRemove) GetVerificationRemoveBodyOk() (*VerificationRemoveBody, bool)`

GetVerificationRemoveBodyOk returns a tuple with the VerificationRemoveBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRemoveBody

`func (o *MessageDataVerificationRemove) SetVerificationRemoveBody(v VerificationRemoveBody)`

SetVerificationRemoveBody sets VerificationRemoveBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


