# ChannelMemberChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Object** | **string** |  | 
**CreatedAt** | Pointer to **float32** | Epoch timestamp in seconds. | [optional] 
**FollowerCount** | Pointer to **float32** | Number of followers the channel has. | [optional] 
**ExternalLink** | Pointer to [**ChannelExternalLink**](ChannelExternalLink.md) |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**ParentUrl** | Pointer to **string** |  | [optional] 
**Lead** | Pointer to [**User**](User.md) |  | [optional] 
**ModeratorFids** | Pointer to **[]int32** |  | [optional] 
**MemberCount** | Pointer to **int32** |  | [optional] 
**Moderator** | Pointer to [**User**](User.md) | Use &#x60;lead&#x60; instead. | [optional] 
**PinnedCastHash** | Pointer to **string** | Cast Hash | [optional] [default to "0xfe90f9de682273e05b201629ad2338bdcd89b6be"]
**Hosts** | Pointer to [**[]User**](User.md) |  | [optional] 
**ViewerContext** | Pointer to [**ChannelUserContext**](ChannelUserContext.md) |  | [optional] 
**DescriptionMentionedProfiles** | Pointer to [**[]UserDehydrated**](UserDehydrated.md) |  | [optional] 
**DescriptionMentionedProfilesRanges** | Pointer to [**[]TextRange**](TextRange.md) | Positions within the text (inclusive start, exclusive end) where each mention occurs. | [optional] 

## Methods

### NewChannelMemberChannel

`func NewChannelMemberChannel(id string, url string, name string, object string, ) *ChannelMemberChannel`

NewChannelMemberChannel instantiates a new ChannelMemberChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelMemberChannelWithDefaults

`func NewChannelMemberChannelWithDefaults() *ChannelMemberChannel`

NewChannelMemberChannelWithDefaults instantiates a new ChannelMemberChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelMemberChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelMemberChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelMemberChannel) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ChannelMemberChannel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChannelMemberChannel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChannelMemberChannel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *ChannelMemberChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelMemberChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelMemberChannel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ChannelMemberChannel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelMemberChannel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelMemberChannel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelMemberChannel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetObject

`func (o *ChannelMemberChannel) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChannelMemberChannel) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChannelMemberChannel) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *ChannelMemberChannel) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChannelMemberChannel) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChannelMemberChannel) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChannelMemberChannel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFollowerCount

`func (o *ChannelMemberChannel) GetFollowerCount() float32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *ChannelMemberChannel) GetFollowerCountOk() (*float32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *ChannelMemberChannel) SetFollowerCount(v float32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *ChannelMemberChannel) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetExternalLink

`func (o *ChannelMemberChannel) GetExternalLink() ChannelExternalLink`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *ChannelMemberChannel) GetExternalLinkOk() (*ChannelExternalLink, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *ChannelMemberChannel) SetExternalLink(v ChannelExternalLink)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *ChannelMemberChannel) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetImageUrl

`func (o *ChannelMemberChannel) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ChannelMemberChannel) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ChannelMemberChannel) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ChannelMemberChannel) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetParentUrl

`func (o *ChannelMemberChannel) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *ChannelMemberChannel) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *ChannelMemberChannel) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.

### HasParentUrl

`func (o *ChannelMemberChannel) HasParentUrl() bool`

HasParentUrl returns a boolean if a field has been set.

### GetLead

`func (o *ChannelMemberChannel) GetLead() User`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *ChannelMemberChannel) GetLeadOk() (*User, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *ChannelMemberChannel) SetLead(v User)`

SetLead sets Lead field to given value.

### HasLead

`func (o *ChannelMemberChannel) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetModeratorFids

`func (o *ChannelMemberChannel) GetModeratorFids() []int32`

GetModeratorFids returns the ModeratorFids field if non-nil, zero value otherwise.

### GetModeratorFidsOk

`func (o *ChannelMemberChannel) GetModeratorFidsOk() (*[]int32, bool)`

GetModeratorFidsOk returns a tuple with the ModeratorFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeratorFids

`func (o *ChannelMemberChannel) SetModeratorFids(v []int32)`

SetModeratorFids sets ModeratorFids field to given value.

### HasModeratorFids

`func (o *ChannelMemberChannel) HasModeratorFids() bool`

HasModeratorFids returns a boolean if a field has been set.

### GetMemberCount

`func (o *ChannelMemberChannel) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *ChannelMemberChannel) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *ChannelMemberChannel) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *ChannelMemberChannel) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetModerator

`func (o *ChannelMemberChannel) GetModerator() User`

GetModerator returns the Moderator field if non-nil, zero value otherwise.

### GetModeratorOk

`func (o *ChannelMemberChannel) GetModeratorOk() (*User, bool)`

GetModeratorOk returns a tuple with the Moderator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerator

`func (o *ChannelMemberChannel) SetModerator(v User)`

SetModerator sets Moderator field to given value.

### HasModerator

`func (o *ChannelMemberChannel) HasModerator() bool`

HasModerator returns a boolean if a field has been set.

### GetPinnedCastHash

`func (o *ChannelMemberChannel) GetPinnedCastHash() string`

GetPinnedCastHash returns the PinnedCastHash field if non-nil, zero value otherwise.

### GetPinnedCastHashOk

`func (o *ChannelMemberChannel) GetPinnedCastHashOk() (*string, bool)`

GetPinnedCastHashOk returns a tuple with the PinnedCastHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedCastHash

`func (o *ChannelMemberChannel) SetPinnedCastHash(v string)`

SetPinnedCastHash sets PinnedCastHash field to given value.

### HasPinnedCastHash

`func (o *ChannelMemberChannel) HasPinnedCastHash() bool`

HasPinnedCastHash returns a boolean if a field has been set.

### GetHosts

`func (o *ChannelMemberChannel) GetHosts() []User`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ChannelMemberChannel) GetHostsOk() (*[]User, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ChannelMemberChannel) SetHosts(v []User)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ChannelMemberChannel) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetViewerContext

`func (o *ChannelMemberChannel) GetViewerContext() ChannelUserContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *ChannelMemberChannel) GetViewerContextOk() (*ChannelUserContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *ChannelMemberChannel) SetViewerContext(v ChannelUserContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *ChannelMemberChannel) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.

### GetDescriptionMentionedProfiles

`func (o *ChannelMemberChannel) GetDescriptionMentionedProfiles() []UserDehydrated`

GetDescriptionMentionedProfiles returns the DescriptionMentionedProfiles field if non-nil, zero value otherwise.

### GetDescriptionMentionedProfilesOk

`func (o *ChannelMemberChannel) GetDescriptionMentionedProfilesOk() (*[]UserDehydrated, bool)`

GetDescriptionMentionedProfilesOk returns a tuple with the DescriptionMentionedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionMentionedProfiles

`func (o *ChannelMemberChannel) SetDescriptionMentionedProfiles(v []UserDehydrated)`

SetDescriptionMentionedProfiles sets DescriptionMentionedProfiles field to given value.

### HasDescriptionMentionedProfiles

`func (o *ChannelMemberChannel) HasDescriptionMentionedProfiles() bool`

HasDescriptionMentionedProfiles returns a boolean if a field has been set.

### GetDescriptionMentionedProfilesRanges

`func (o *ChannelMemberChannel) GetDescriptionMentionedProfilesRanges() []TextRange`

GetDescriptionMentionedProfilesRanges returns the DescriptionMentionedProfilesRanges field if non-nil, zero value otherwise.

### GetDescriptionMentionedProfilesRangesOk

`func (o *ChannelMemberChannel) GetDescriptionMentionedProfilesRangesOk() (*[]TextRange, bool)`

GetDescriptionMentionedProfilesRangesOk returns a tuple with the DescriptionMentionedProfilesRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionMentionedProfilesRanges

`func (o *ChannelMemberChannel) SetDescriptionMentionedProfilesRanges(v []TextRange)`

SetDescriptionMentionedProfilesRanges sets DescriptionMentionedProfilesRanges field to given value.

### HasDescriptionMentionedProfilesRanges

`func (o *ChannelMemberChannel) HasDescriptionMentionedProfilesRanges() bool`

HasDescriptionMentionedProfilesRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


