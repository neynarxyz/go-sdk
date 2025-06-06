# RespondChannelInviteReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer. &#x60;signer_uuid&#x60; is paired with API key, can&#39;t use a &#x60;uuid&#x60; made with a different API key. | 
**ChannelId** | **string** | The unique identifier of a farcaster channel | 
**Role** | [**ChannelMemberRole**](ChannelMemberRole.md) |  | 
**Accept** | **bool** | Accept or reject the invite | 

## Methods

### NewRespondChannelInviteReqBody

`func NewRespondChannelInviteReqBody(signerUuid string, channelId string, role ChannelMemberRole, accept bool, ) *RespondChannelInviteReqBody`

NewRespondChannelInviteReqBody instantiates a new RespondChannelInviteReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespondChannelInviteReqBodyWithDefaults

`func NewRespondChannelInviteReqBodyWithDefaults() *RespondChannelInviteReqBody`

NewRespondChannelInviteReqBodyWithDefaults instantiates a new RespondChannelInviteReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *RespondChannelInviteReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *RespondChannelInviteReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *RespondChannelInviteReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetChannelId

`func (o *RespondChannelInviteReqBody) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *RespondChannelInviteReqBody) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *RespondChannelInviteReqBody) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetRole

`func (o *RespondChannelInviteReqBody) GetRole() ChannelMemberRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RespondChannelInviteReqBody) GetRoleOk() (*ChannelMemberRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RespondChannelInviteReqBody) SetRole(v ChannelMemberRole)`

SetRole sets Role field to given value.


### GetAccept

`func (o *RespondChannelInviteReqBody) GetAccept() bool`

GetAccept returns the Accept field if non-nil, zero value otherwise.

### GetAcceptOk

`func (o *RespondChannelInviteReqBody) GetAcceptOk() (*bool, bool)`

GetAcceptOk returns a tuple with the Accept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccept

`func (o *RespondChannelInviteReqBody) SetAccept(v bool)`

SetAccept sets Accept field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


