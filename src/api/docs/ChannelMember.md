# ChannelMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Role** | [**ChannelMemberRole**](ChannelMemberRole.md) |  | 
**User** | [**ChannelMemberUser**](ChannelMemberUser.md) |  | 
**Channel** | [**ChannelMemberChannel**](ChannelMemberChannel.md) |  | 

## Methods

### NewChannelMember

`func NewChannelMember(object string, role ChannelMemberRole, user ChannelMemberUser, channel ChannelMemberChannel, ) *ChannelMember`

NewChannelMember instantiates a new ChannelMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelMemberWithDefaults

`func NewChannelMemberWithDefaults() *ChannelMember`

NewChannelMemberWithDefaults instantiates a new ChannelMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ChannelMember) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChannelMember) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChannelMember) SetObject(v string)`

SetObject sets Object field to given value.


### GetRole

`func (o *ChannelMember) GetRole() ChannelMemberRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChannelMember) GetRoleOk() (*ChannelMemberRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChannelMember) SetRole(v ChannelMemberRole)`

SetRole sets Role field to given value.


### GetUser

`func (o *ChannelMember) GetUser() ChannelMemberUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChannelMember) GetUserOk() (*ChannelMemberUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChannelMember) SetUser(v ChannelMemberUser)`

SetUser sets User field to given value.


### GetChannel

`func (o *ChannelMember) GetChannel() ChannelMemberChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChannelMember) GetChannelOk() (*ChannelMemberChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChannelMember) SetChannel(v ChannelMemberChannel)`

SetChannel sets Channel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


