# ChannelFollowReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer. &#x60;signer_uuid&#x60; is paired with API key, can&#39;t use a &#x60;uuid&#x60; made with a different API key.  | 
**ChannelId** | **string** | The unique identifier of a farcaster channel | 

## Methods

### NewChannelFollowReqBody

`func NewChannelFollowReqBody(signerUuid string, channelId string, ) *ChannelFollowReqBody`

NewChannelFollowReqBody instantiates a new ChannelFollowReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelFollowReqBodyWithDefaults

`func NewChannelFollowReqBodyWithDefaults() *ChannelFollowReqBody`

NewChannelFollowReqBodyWithDefaults instantiates a new ChannelFollowReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *ChannelFollowReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *ChannelFollowReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *ChannelFollowReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetChannelId

`func (o *ChannelFollowReqBody) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ChannelFollowReqBody) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ChannelFollowReqBody) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


