# MessageAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageType**](MessageType.md) |  | [default to MESSAGETYPE_MESSAGE_TYPE_CAST_ADD]
**Fid** | **int32** | The unique identifier (FID) of the user who created this message. FIDs are assigned sequentially when users register on the network and cannot be changed. | 
**Timestamp** | **int64** | Seconds since Farcaster Epoch (2021-01-01T00:00:00Z). Used to order messages chronologically and determine the most recent state. Must be within 10 minutes of the current time when the message is created. | 
**Network** | [**FarcasterNetwork**](FarcasterNetwork.md) |  | [default to FARCASTERNETWORK_FARCASTER_NETWORK_MAINNET]
**CastAddBody** | [**CastAddBody**](CastAddBody.md) | The content and metadata of the new cast, including the text, mentions, embeds, and any parent references for replies. | 
**TargetHash** | **string** | The unique hash identifier of the cast to be removed. Must be a cast that was previously created by the same FID specified in the message. | 
**ReactionBody** | [**ReactionBody**](ReactionBody.md) | Contains the type of reaction (like/recast) and the target content being reacted to. The target can be specified either by castId or URL. | 
**LinkBody** | [**LinkBody**](LinkBody.md) | Contains the details of the social connection, including the type of relationship and the target user. | 
**VerificationAddEthAddressBody** | [**VerificationAddEthAddressBody**](VerificationAddEthAddressBody.md) | Contains the blockchain address being verified, along with cryptographic proof of ownership through a signature. | 
**VerificationRemoveBody** | [**VerificationRemoveBody**](VerificationRemoveBody.md) | Contains the blockchain address for which the verification should be removed from the user&#39;s profile. | 
**UserDataBody** | [**UserDataBody**](UserDataBody.md) | Contains the type of profile metadata being updated and its new value. | 
**UsernameProofBody** | [**UserNameProof**](UserNameProof.md) |  | 
**FrameActionBody** | [**FrameActionBody**](FrameActionBody.md) | Contains the details of the frame interaction, including which button was pressed and the associated cast and URL. | 

## Methods

### NewMessageAllOfData

`func NewMessageAllOfData(type_ MessageType, fid int32, timestamp int64, network FarcasterNetwork, castAddBody CastAddBody, targetHash string, reactionBody ReactionBody, linkBody LinkBody, verificationAddEthAddressBody VerificationAddEthAddressBody, verificationRemoveBody VerificationRemoveBody, userDataBody UserDataBody, usernameProofBody UserNameProof, frameActionBody FrameActionBody, ) *MessageAllOfData`

NewMessageAllOfData instantiates a new MessageAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageAllOfDataWithDefaults

`func NewMessageAllOfDataWithDefaults() *MessageAllOfData`

NewMessageAllOfDataWithDefaults instantiates a new MessageAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageAllOfData) GetType() MessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageAllOfData) GetTypeOk() (*MessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageAllOfData) SetType(v MessageType)`

SetType sets Type field to given value.


### GetFid

`func (o *MessageAllOfData) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageAllOfData) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageAllOfData) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetTimestamp

`func (o *MessageAllOfData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageAllOfData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageAllOfData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNetwork

`func (o *MessageAllOfData) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageAllOfData) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageAllOfData) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.


### GetCastAddBody

`func (o *MessageAllOfData) GetCastAddBody() CastAddBody`

GetCastAddBody returns the CastAddBody field if non-nil, zero value otherwise.

### GetCastAddBodyOk

`func (o *MessageAllOfData) GetCastAddBodyOk() (*CastAddBody, bool)`

GetCastAddBodyOk returns a tuple with the CastAddBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastAddBody

`func (o *MessageAllOfData) SetCastAddBody(v CastAddBody)`

SetCastAddBody sets CastAddBody field to given value.


### GetTargetHash

`func (o *MessageAllOfData) GetTargetHash() string`

GetTargetHash returns the TargetHash field if non-nil, zero value otherwise.

### GetTargetHashOk

`func (o *MessageAllOfData) GetTargetHashOk() (*string, bool)`

GetTargetHashOk returns a tuple with the TargetHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHash

`func (o *MessageAllOfData) SetTargetHash(v string)`

SetTargetHash sets TargetHash field to given value.


### GetReactionBody

`func (o *MessageAllOfData) GetReactionBody() ReactionBody`

GetReactionBody returns the ReactionBody field if non-nil, zero value otherwise.

### GetReactionBodyOk

`func (o *MessageAllOfData) GetReactionBodyOk() (*ReactionBody, bool)`

GetReactionBodyOk returns a tuple with the ReactionBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionBody

`func (o *MessageAllOfData) SetReactionBody(v ReactionBody)`

SetReactionBody sets ReactionBody field to given value.


### GetLinkBody

`func (o *MessageAllOfData) GetLinkBody() LinkBody`

GetLinkBody returns the LinkBody field if non-nil, zero value otherwise.

### GetLinkBodyOk

`func (o *MessageAllOfData) GetLinkBodyOk() (*LinkBody, bool)`

GetLinkBodyOk returns a tuple with the LinkBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkBody

`func (o *MessageAllOfData) SetLinkBody(v LinkBody)`

SetLinkBody sets LinkBody field to given value.


### GetVerificationAddEthAddressBody

`func (o *MessageAllOfData) GetVerificationAddEthAddressBody() VerificationAddEthAddressBody`

GetVerificationAddEthAddressBody returns the VerificationAddEthAddressBody field if non-nil, zero value otherwise.

### GetVerificationAddEthAddressBodyOk

`func (o *MessageAllOfData) GetVerificationAddEthAddressBodyOk() (*VerificationAddEthAddressBody, bool)`

GetVerificationAddEthAddressBodyOk returns a tuple with the VerificationAddEthAddressBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationAddEthAddressBody

`func (o *MessageAllOfData) SetVerificationAddEthAddressBody(v VerificationAddEthAddressBody)`

SetVerificationAddEthAddressBody sets VerificationAddEthAddressBody field to given value.


### GetVerificationRemoveBody

`func (o *MessageAllOfData) GetVerificationRemoveBody() VerificationRemoveBody`

GetVerificationRemoveBody returns the VerificationRemoveBody field if non-nil, zero value otherwise.

### GetVerificationRemoveBodyOk

`func (o *MessageAllOfData) GetVerificationRemoveBodyOk() (*VerificationRemoveBody, bool)`

GetVerificationRemoveBodyOk returns a tuple with the VerificationRemoveBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRemoveBody

`func (o *MessageAllOfData) SetVerificationRemoveBody(v VerificationRemoveBody)`

SetVerificationRemoveBody sets VerificationRemoveBody field to given value.


### GetUserDataBody

`func (o *MessageAllOfData) GetUserDataBody() UserDataBody`

GetUserDataBody returns the UserDataBody field if non-nil, zero value otherwise.

### GetUserDataBodyOk

`func (o *MessageAllOfData) GetUserDataBodyOk() (*UserDataBody, bool)`

GetUserDataBodyOk returns a tuple with the UserDataBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataBody

`func (o *MessageAllOfData) SetUserDataBody(v UserDataBody)`

SetUserDataBody sets UserDataBody field to given value.


### GetUsernameProofBody

`func (o *MessageAllOfData) GetUsernameProofBody() UserNameProof`

GetUsernameProofBody returns the UsernameProofBody field if non-nil, zero value otherwise.

### GetUsernameProofBodyOk

`func (o *MessageAllOfData) GetUsernameProofBodyOk() (*UserNameProof, bool)`

GetUsernameProofBodyOk returns a tuple with the UsernameProofBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameProofBody

`func (o *MessageAllOfData) SetUsernameProofBody(v UserNameProof)`

SetUsernameProofBody sets UsernameProofBody field to given value.


### GetFrameActionBody

`func (o *MessageAllOfData) GetFrameActionBody() FrameActionBody`

GetFrameActionBody returns the FrameActionBody field if non-nil, zero value otherwise.

### GetFrameActionBodyOk

`func (o *MessageAllOfData) GetFrameActionBodyOk() (*FrameActionBody, bool)`

GetFrameActionBodyOk returns a tuple with the FrameActionBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameActionBody

`func (o *MessageAllOfData) SetFrameActionBody(v FrameActionBody)`

SetFrameActionBody sets FrameActionBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


