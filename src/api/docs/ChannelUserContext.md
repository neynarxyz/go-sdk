# ChannelUserContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Following** | **bool** | Indicates if the user is following the channel. | 
**Role** | Pointer to [**ChannelMemberRole**](ChannelMemberRole.md) |  | [optional] 

## Methods

### NewChannelUserContext

`func NewChannelUserContext(following bool, ) *ChannelUserContext`

NewChannelUserContext instantiates a new ChannelUserContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelUserContextWithDefaults

`func NewChannelUserContextWithDefaults() *ChannelUserContext`

NewChannelUserContextWithDefaults instantiates a new ChannelUserContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowing

`func (o *ChannelUserContext) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *ChannelUserContext) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *ChannelUserContext) SetFollowing(v bool)`

SetFollowing sets Following field to given value.


### GetRole

`func (o *ChannelUserContext) GetRole() ChannelMemberRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChannelUserContext) GetRoleOk() (*ChannelMemberRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChannelUserContext) SetRole(v ChannelMemberRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ChannelUserContext) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


