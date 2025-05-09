# ChannelMemberInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** | The unique identifier of a farcaster channel | 
**Role** | [**ChannelMemberRole**](ChannelMemberRole.md) |  | 
**Inviter** | [**User**](User.md) |  | 
**Invited** | [**User**](User.md) |  | 

## Methods

### NewChannelMemberInvite

`func NewChannelMemberInvite(channelId string, role ChannelMemberRole, inviter User, invited User, ) *ChannelMemberInvite`

NewChannelMemberInvite instantiates a new ChannelMemberInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelMemberInviteWithDefaults

`func NewChannelMemberInviteWithDefaults() *ChannelMemberInvite`

NewChannelMemberInviteWithDefaults instantiates a new ChannelMemberInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *ChannelMemberInvite) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ChannelMemberInvite) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ChannelMemberInvite) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetRole

`func (o *ChannelMemberInvite) GetRole() ChannelMemberRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChannelMemberInvite) GetRoleOk() (*ChannelMemberRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChannelMemberInvite) SetRole(v ChannelMemberRole)`

SetRole sets Role field to given value.


### GetInviter

`func (o *ChannelMemberInvite) GetInviter() User`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *ChannelMemberInvite) GetInviterOk() (*User, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *ChannelMemberInvite) SetInviter(v User)`

SetInviter sets Inviter field to given value.


### GetInvited

`func (o *ChannelMemberInvite) GetInvited() User`

GetInvited returns the Invited field if non-nil, zero value otherwise.

### GetInvitedOk

`func (o *ChannelMemberInvite) GetInvitedOk() (*User, bool)`

GetInvitedOk returns a tuple with the Invited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvited

`func (o *ChannelMemberInvite) SetInvited(v User)`

SetInvited sets Invited field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


