# ValidateFrameActionReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageBytesInHex** | **string** | Hexadecimal string of message bytes. | 
**CastReactionContext** | Pointer to **bool** | Adds viewer_context inside the cast object to indicate whether the interactor reacted to the cast housing the mini app. | [optional] [default to true]
**FollowContext** | Pointer to **bool** | Adds viewer_context inside the user (interactor) object to indicate whether the interactor follows or is followed by the cast author. | [optional] [default to false]
**SignerContext** | Pointer to **bool** | Adds context about the app used by the user inside &#x60;frame.action&#x60;. | [optional] [default to false]
**ChannelFollowContext** | Pointer to **bool** | Adds context about the channel that the cast belongs to inside of the cast object. | [optional] [default to false]

## Methods

### NewValidateFrameActionReqBody

`func NewValidateFrameActionReqBody(messageBytesInHex string, ) *ValidateFrameActionReqBody`

NewValidateFrameActionReqBody instantiates a new ValidateFrameActionReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateFrameActionReqBodyWithDefaults

`func NewValidateFrameActionReqBodyWithDefaults() *ValidateFrameActionReqBody`

NewValidateFrameActionReqBodyWithDefaults instantiates a new ValidateFrameActionReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageBytesInHex

`func (o *ValidateFrameActionReqBody) GetMessageBytesInHex() string`

GetMessageBytesInHex returns the MessageBytesInHex field if non-nil, zero value otherwise.

### GetMessageBytesInHexOk

`func (o *ValidateFrameActionReqBody) GetMessageBytesInHexOk() (*string, bool)`

GetMessageBytesInHexOk returns a tuple with the MessageBytesInHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBytesInHex

`func (o *ValidateFrameActionReqBody) SetMessageBytesInHex(v string)`

SetMessageBytesInHex sets MessageBytesInHex field to given value.


### GetCastReactionContext

`func (o *ValidateFrameActionReqBody) GetCastReactionContext() bool`

GetCastReactionContext returns the CastReactionContext field if non-nil, zero value otherwise.

### GetCastReactionContextOk

`func (o *ValidateFrameActionReqBody) GetCastReactionContextOk() (*bool, bool)`

GetCastReactionContextOk returns a tuple with the CastReactionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastReactionContext

`func (o *ValidateFrameActionReqBody) SetCastReactionContext(v bool)`

SetCastReactionContext sets CastReactionContext field to given value.

### HasCastReactionContext

`func (o *ValidateFrameActionReqBody) HasCastReactionContext() bool`

HasCastReactionContext returns a boolean if a field has been set.

### GetFollowContext

`func (o *ValidateFrameActionReqBody) GetFollowContext() bool`

GetFollowContext returns the FollowContext field if non-nil, zero value otherwise.

### GetFollowContextOk

`func (o *ValidateFrameActionReqBody) GetFollowContextOk() (*bool, bool)`

GetFollowContextOk returns a tuple with the FollowContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowContext

`func (o *ValidateFrameActionReqBody) SetFollowContext(v bool)`

SetFollowContext sets FollowContext field to given value.

### HasFollowContext

`func (o *ValidateFrameActionReqBody) HasFollowContext() bool`

HasFollowContext returns a boolean if a field has been set.

### GetSignerContext

`func (o *ValidateFrameActionReqBody) GetSignerContext() bool`

GetSignerContext returns the SignerContext field if non-nil, zero value otherwise.

### GetSignerContextOk

`func (o *ValidateFrameActionReqBody) GetSignerContextOk() (*bool, bool)`

GetSignerContextOk returns a tuple with the SignerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerContext

`func (o *ValidateFrameActionReqBody) SetSignerContext(v bool)`

SetSignerContext sets SignerContext field to given value.

### HasSignerContext

`func (o *ValidateFrameActionReqBody) HasSignerContext() bool`

HasSignerContext returns a boolean if a field has been set.

### GetChannelFollowContext

`func (o *ValidateFrameActionReqBody) GetChannelFollowContext() bool`

GetChannelFollowContext returns the ChannelFollowContext field if non-nil, zero value otherwise.

### GetChannelFollowContextOk

`func (o *ValidateFrameActionReqBody) GetChannelFollowContextOk() (*bool, bool)`

GetChannelFollowContextOk returns a tuple with the ChannelFollowContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelFollowContext

`func (o *ValidateFrameActionReqBody) SetChannelFollowContext(v bool)`

SetChannelFollowContext sets ChannelFollowContext field to given value.

### HasChannelFollowContext

`func (o *ValidateFrameActionReqBody) HasChannelFollowContext() bool`

HasChannelFollowContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


