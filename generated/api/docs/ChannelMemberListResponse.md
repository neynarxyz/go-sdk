# ChannelMemberListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]ChannelMember**](ChannelMember.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewChannelMemberListResponse

`func NewChannelMemberListResponse(members []ChannelMember, next NextCursor, ) *ChannelMemberListResponse`

NewChannelMemberListResponse instantiates a new ChannelMemberListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelMemberListResponseWithDefaults

`func NewChannelMemberListResponseWithDefaults() *ChannelMemberListResponse`

NewChannelMemberListResponseWithDefaults instantiates a new ChannelMemberListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ChannelMemberListResponse) GetMembers() []ChannelMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ChannelMemberListResponse) GetMembersOk() (*[]ChannelMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ChannelMemberListResponse) SetMembers(v []ChannelMember)`

SetMembers sets Members field to given value.


### GetNext

`func (o *ChannelMemberListResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ChannelMemberListResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ChannelMemberListResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


