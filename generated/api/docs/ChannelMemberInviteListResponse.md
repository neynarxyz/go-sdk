# ChannelMemberInviteListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invites** | [**[]ChannelMemberInvite**](ChannelMemberInvite.md) |  | 
**Next** | Pointer to [**NextCursor**](NextCursor.md) |  | [optional] 

## Methods

### NewChannelMemberInviteListResponse

`func NewChannelMemberInviteListResponse(invites []ChannelMemberInvite, ) *ChannelMemberInviteListResponse`

NewChannelMemberInviteListResponse instantiates a new ChannelMemberInviteListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelMemberInviteListResponseWithDefaults

`func NewChannelMemberInviteListResponseWithDefaults() *ChannelMemberInviteListResponse`

NewChannelMemberInviteListResponseWithDefaults instantiates a new ChannelMemberInviteListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvites

`func (o *ChannelMemberInviteListResponse) GetInvites() []ChannelMemberInvite`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *ChannelMemberInviteListResponse) GetInvitesOk() (*[]ChannelMemberInvite, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *ChannelMemberInviteListResponse) SetInvites(v []ChannelMemberInvite)`

SetInvites sets Invites field to given value.


### GetNext

`func (o *ChannelMemberInviteListResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ChannelMemberInviteListResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ChannelMemberInviteListResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.

### HasNext

`func (o *ChannelMemberInviteListResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


