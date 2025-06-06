# RemoveChannelMemberReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer. &#x60;signer_uuid&#x60; is paired with API key, can&#39;t use a &#x60;uuid&#x60; made with a different API key.  | 
**ChannelId** | **string** | The unique identifier of a farcaster channel | 
**Fid** | **int32** | The unique identifier of a farcaster user (unsigned integer) | 
**Role** | [**ChannelMemberRole**](ChannelMemberRole.md) |  | 

## Methods

### NewRemoveChannelMemberReqBody

`func NewRemoveChannelMemberReqBody(signerUuid string, channelId string, fid int32, role ChannelMemberRole, ) *RemoveChannelMemberReqBody`

NewRemoveChannelMemberReqBody instantiates a new RemoveChannelMemberReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveChannelMemberReqBodyWithDefaults

`func NewRemoveChannelMemberReqBodyWithDefaults() *RemoveChannelMemberReqBody`

NewRemoveChannelMemberReqBodyWithDefaults instantiates a new RemoveChannelMemberReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *RemoveChannelMemberReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *RemoveChannelMemberReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *RemoveChannelMemberReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetChannelId

`func (o *RemoveChannelMemberReqBody) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *RemoveChannelMemberReqBody) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *RemoveChannelMemberReqBody) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetFid

`func (o *RemoveChannelMemberReqBody) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *RemoveChannelMemberReqBody) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *RemoveChannelMemberReqBody) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetRole

`func (o *RemoveChannelMemberReqBody) GetRole() ChannelMemberRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RemoveChannelMemberReqBody) GetRoleOk() (*ChannelMemberRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RemoveChannelMemberReqBody) SetRole(v ChannelMemberRole)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


