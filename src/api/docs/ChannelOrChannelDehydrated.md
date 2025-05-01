# ChannelOrChannelDehydrated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Object** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DescriptionMentionedProfiles** | Pointer to [**[]UserDehydrated**](UserDehydrated.md) |  | [optional] 
**DescriptionMentionedProfilesRanges** | Pointer to [**[]TextRange**](TextRange.md) | Positions within the text (inclusive start, exclusive end) where each mention occurs. | [optional] 
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

## Methods

### NewChannelOrChannelDehydrated

`func NewChannelOrChannelDehydrated(id string, url string, object string, name string, ) *ChannelOrChannelDehydrated`

NewChannelOrChannelDehydrated instantiates a new ChannelOrChannelDehydrated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelOrChannelDehydratedWithDefaults

`func NewChannelOrChannelDehydratedWithDefaults() *ChannelOrChannelDehydrated`

NewChannelOrChannelDehydratedWithDefaults instantiates a new ChannelOrChannelDehydrated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelOrChannelDehydrated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelOrChannelDehydrated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelOrChannelDehydrated) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ChannelOrChannelDehydrated) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChannelOrChannelDehydrated) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChannelOrChannelDehydrated) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetObject

`func (o *ChannelOrChannelDehydrated) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChannelOrChannelDehydrated) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChannelOrChannelDehydrated) SetObject(v string)`

SetObject sets Object field to given value.


### GetName

`func (o *ChannelOrChannelDehydrated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelOrChannelDehydrated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelOrChannelDehydrated) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ChannelOrChannelDehydrated) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelOrChannelDehydrated) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelOrChannelDehydrated) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelOrChannelDehydrated) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionMentionedProfiles

`func (o *ChannelOrChannelDehydrated) GetDescriptionMentionedProfiles() []UserDehydrated`

GetDescriptionMentionedProfiles returns the DescriptionMentionedProfiles field if non-nil, zero value otherwise.

### GetDescriptionMentionedProfilesOk

`func (o *ChannelOrChannelDehydrated) GetDescriptionMentionedProfilesOk() (*[]UserDehydrated, bool)`

GetDescriptionMentionedProfilesOk returns a tuple with the DescriptionMentionedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionMentionedProfiles

`func (o *ChannelOrChannelDehydrated) SetDescriptionMentionedProfiles(v []UserDehydrated)`

SetDescriptionMentionedProfiles sets DescriptionMentionedProfiles field to given value.

### HasDescriptionMentionedProfiles

`func (o *ChannelOrChannelDehydrated) HasDescriptionMentionedProfiles() bool`

HasDescriptionMentionedProfiles returns a boolean if a field has been set.

### GetDescriptionMentionedProfilesRanges

`func (o *ChannelOrChannelDehydrated) GetDescriptionMentionedProfilesRanges() []TextRange`

GetDescriptionMentionedProfilesRanges returns the DescriptionMentionedProfilesRanges field if non-nil, zero value otherwise.

### GetDescriptionMentionedProfilesRangesOk

`func (o *ChannelOrChannelDehydrated) GetDescriptionMentionedProfilesRangesOk() (*[]TextRange, bool)`

GetDescriptionMentionedProfilesRangesOk returns a tuple with the DescriptionMentionedProfilesRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionMentionedProfilesRanges

`func (o *ChannelOrChannelDehydrated) SetDescriptionMentionedProfilesRanges(v []TextRange)`

SetDescriptionMentionedProfilesRanges sets DescriptionMentionedProfilesRanges field to given value.

### HasDescriptionMentionedProfilesRanges

`func (o *ChannelOrChannelDehydrated) HasDescriptionMentionedProfilesRanges() bool`

HasDescriptionMentionedProfilesRanges returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChannelOrChannelDehydrated) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChannelOrChannelDehydrated) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChannelOrChannelDehydrated) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChannelOrChannelDehydrated) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFollowerCount

`func (o *ChannelOrChannelDehydrated) GetFollowerCount() float32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *ChannelOrChannelDehydrated) GetFollowerCountOk() (*float32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *ChannelOrChannelDehydrated) SetFollowerCount(v float32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *ChannelOrChannelDehydrated) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetExternalLink

`func (o *ChannelOrChannelDehydrated) GetExternalLink() ChannelExternalLink`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *ChannelOrChannelDehydrated) GetExternalLinkOk() (*ChannelExternalLink, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *ChannelOrChannelDehydrated) SetExternalLink(v ChannelExternalLink)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *ChannelOrChannelDehydrated) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetImageUrl

`func (o *ChannelOrChannelDehydrated) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ChannelOrChannelDehydrated) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ChannelOrChannelDehydrated) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ChannelOrChannelDehydrated) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetParentUrl

`func (o *ChannelOrChannelDehydrated) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *ChannelOrChannelDehydrated) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *ChannelOrChannelDehydrated) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.

### HasParentUrl

`func (o *ChannelOrChannelDehydrated) HasParentUrl() bool`

HasParentUrl returns a boolean if a field has been set.

### GetLead

`func (o *ChannelOrChannelDehydrated) GetLead() User`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *ChannelOrChannelDehydrated) GetLeadOk() (*User, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *ChannelOrChannelDehydrated) SetLead(v User)`

SetLead sets Lead field to given value.

### HasLead

`func (o *ChannelOrChannelDehydrated) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetModeratorFids

`func (o *ChannelOrChannelDehydrated) GetModeratorFids() []int32`

GetModeratorFids returns the ModeratorFids field if non-nil, zero value otherwise.

### GetModeratorFidsOk

`func (o *ChannelOrChannelDehydrated) GetModeratorFidsOk() (*[]int32, bool)`

GetModeratorFidsOk returns a tuple with the ModeratorFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeratorFids

`func (o *ChannelOrChannelDehydrated) SetModeratorFids(v []int32)`

SetModeratorFids sets ModeratorFids field to given value.

### HasModeratorFids

`func (o *ChannelOrChannelDehydrated) HasModeratorFids() bool`

HasModeratorFids returns a boolean if a field has been set.

### GetMemberCount

`func (o *ChannelOrChannelDehydrated) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *ChannelOrChannelDehydrated) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *ChannelOrChannelDehydrated) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *ChannelOrChannelDehydrated) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetModerator

`func (o *ChannelOrChannelDehydrated) GetModerator() User`

GetModerator returns the Moderator field if non-nil, zero value otherwise.

### GetModeratorOk

`func (o *ChannelOrChannelDehydrated) GetModeratorOk() (*User, bool)`

GetModeratorOk returns a tuple with the Moderator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerator

`func (o *ChannelOrChannelDehydrated) SetModerator(v User)`

SetModerator sets Moderator field to given value.

### HasModerator

`func (o *ChannelOrChannelDehydrated) HasModerator() bool`

HasModerator returns a boolean if a field has been set.

### GetPinnedCastHash

`func (o *ChannelOrChannelDehydrated) GetPinnedCastHash() string`

GetPinnedCastHash returns the PinnedCastHash field if non-nil, zero value otherwise.

### GetPinnedCastHashOk

`func (o *ChannelOrChannelDehydrated) GetPinnedCastHashOk() (*string, bool)`

GetPinnedCastHashOk returns a tuple with the PinnedCastHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedCastHash

`func (o *ChannelOrChannelDehydrated) SetPinnedCastHash(v string)`

SetPinnedCastHash sets PinnedCastHash field to given value.

### HasPinnedCastHash

`func (o *ChannelOrChannelDehydrated) HasPinnedCastHash() bool`

HasPinnedCastHash returns a boolean if a field has been set.

### GetHosts

`func (o *ChannelOrChannelDehydrated) GetHosts() []User`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ChannelOrChannelDehydrated) GetHostsOk() (*[]User, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ChannelOrChannelDehydrated) SetHosts(v []User)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ChannelOrChannelDehydrated) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetViewerContext

`func (o *ChannelOrChannelDehydrated) GetViewerContext() ChannelUserContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *ChannelOrChannelDehydrated) GetViewerContextOk() (*ChannelUserContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *ChannelOrChannelDehydrated) SetViewerContext(v ChannelUserContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *ChannelOrChannelDehydrated) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


