# CastWithInteractions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Hash** | **string** |  | 
**ParentHash** | **NullableString** |  | 
**ParentUrl** | **NullableString** |  | 
**RootParentUrl** | **NullableString** |  | 
**ParentAuthor** | [**CastParentAuthor**](CastParentAuthor.md) |  | 
**Author** | [**User**](User.md) |  | 
**App** | Pointer to [**UserDehydrated**](UserDehydrated.md) |  | [optional] 
**Text** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Embeds** | [**[]Embed**](Embed.md) |  | 
**Type** | Pointer to [**CastNotificationType**](CastNotificationType.md) |  | [optional] 
**Frames** | Pointer to [**[]Frame**](Frame.md) |  | [optional] 
**Reactions** | [**CastWithInteractionsReactions**](CastWithInteractionsReactions.md) |  | 
**Replies** | [**CastWithInteractionsReplies**](CastWithInteractionsReplies.md) |  | 
**ThreadHash** | **NullableString** |  | 
**MentionedProfiles** | [**[]User**](User.md) |  | 
**MentionedProfilesRanges** | [**[]TextRange**](TextRange.md) | Positions within the text (inclusive start, exclusive end) where each mention occurs. Each index within this list corresponds to the same-numbered index in the mentioned_profiles list.  | 
**MentionedChannels** | [**[]ChannelDehydrated**](ChannelDehydrated.md) |  | 
**MentionedChannelsRanges** | [**[]TextRange**](TextRange.md) | Positions within the text (inclusive start, exclusive end) where each mention occurs. Each index within this list corresponds to the same-numbered index in the mentioned_channels list.  | 
**Channel** | [**ChannelOrChannelDehydrated**](ChannelOrChannelDehydrated.md) |  | 
**ViewerContext** | Pointer to [**CastViewerContext**](CastViewerContext.md) |  | [optional] 
**AuthorChannelContext** | Pointer to [**ChannelUserContext**](ChannelUserContext.md) |  | [optional] 

## Methods

### NewCastWithInteractions

`func NewCastWithInteractions(object string, hash string, parentHash NullableString, parentUrl NullableString, rootParentUrl NullableString, parentAuthor CastParentAuthor, author User, text string, timestamp time.Time, embeds []Embed, reactions CastWithInteractionsReactions, replies CastWithInteractionsReplies, threadHash NullableString, mentionedProfiles []User, mentionedProfilesRanges []TextRange, mentionedChannels []ChannelDehydrated, mentionedChannelsRanges []TextRange, channel ChannelOrChannelDehydrated, ) *CastWithInteractions`

NewCastWithInteractions instantiates a new CastWithInteractions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastWithInteractionsWithDefaults

`func NewCastWithInteractionsWithDefaults() *CastWithInteractions`

NewCastWithInteractionsWithDefaults instantiates a new CastWithInteractions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CastWithInteractions) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CastWithInteractions) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CastWithInteractions) SetObject(v string)`

SetObject sets Object field to given value.


### GetHash

`func (o *CastWithInteractions) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CastWithInteractions) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CastWithInteractions) SetHash(v string)`

SetHash sets Hash field to given value.


### GetParentHash

`func (o *CastWithInteractions) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *CastWithInteractions) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *CastWithInteractions) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### SetParentHashNil

`func (o *CastWithInteractions) SetParentHashNil(b bool)`

 SetParentHashNil sets the value for ParentHash to be an explicit nil

### UnsetParentHash
`func (o *CastWithInteractions) UnsetParentHash()`

UnsetParentHash ensures that no value is present for ParentHash, not even an explicit nil
### GetParentUrl

`func (o *CastWithInteractions) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *CastWithInteractions) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *CastWithInteractions) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.


### SetParentUrlNil

`func (o *CastWithInteractions) SetParentUrlNil(b bool)`

 SetParentUrlNil sets the value for ParentUrl to be an explicit nil

### UnsetParentUrl
`func (o *CastWithInteractions) UnsetParentUrl()`

UnsetParentUrl ensures that no value is present for ParentUrl, not even an explicit nil
### GetRootParentUrl

`func (o *CastWithInteractions) GetRootParentUrl() string`

GetRootParentUrl returns the RootParentUrl field if non-nil, zero value otherwise.

### GetRootParentUrlOk

`func (o *CastWithInteractions) GetRootParentUrlOk() (*string, bool)`

GetRootParentUrlOk returns a tuple with the RootParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootParentUrl

`func (o *CastWithInteractions) SetRootParentUrl(v string)`

SetRootParentUrl sets RootParentUrl field to given value.


### SetRootParentUrlNil

`func (o *CastWithInteractions) SetRootParentUrlNil(b bool)`

 SetRootParentUrlNil sets the value for RootParentUrl to be an explicit nil

### UnsetRootParentUrl
`func (o *CastWithInteractions) UnsetRootParentUrl()`

UnsetRootParentUrl ensures that no value is present for RootParentUrl, not even an explicit nil
### GetParentAuthor

`func (o *CastWithInteractions) GetParentAuthor() CastParentAuthor`

GetParentAuthor returns the ParentAuthor field if non-nil, zero value otherwise.

### GetParentAuthorOk

`func (o *CastWithInteractions) GetParentAuthorOk() (*CastParentAuthor, bool)`

GetParentAuthorOk returns a tuple with the ParentAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAuthor

`func (o *CastWithInteractions) SetParentAuthor(v CastParentAuthor)`

SetParentAuthor sets ParentAuthor field to given value.


### GetAuthor

`func (o *CastWithInteractions) GetAuthor() User`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CastWithInteractions) GetAuthorOk() (*User, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CastWithInteractions) SetAuthor(v User)`

SetAuthor sets Author field to given value.


### GetApp

`func (o *CastWithInteractions) GetApp() UserDehydrated`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CastWithInteractions) GetAppOk() (*UserDehydrated, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CastWithInteractions) SetApp(v UserDehydrated)`

SetApp sets App field to given value.

### HasApp

`func (o *CastWithInteractions) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetText

`func (o *CastWithInteractions) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CastWithInteractions) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CastWithInteractions) SetText(v string)`

SetText sets Text field to given value.


### GetTimestamp

`func (o *CastWithInteractions) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CastWithInteractions) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CastWithInteractions) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEmbeds

`func (o *CastWithInteractions) GetEmbeds() []Embed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *CastWithInteractions) GetEmbedsOk() (*[]Embed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *CastWithInteractions) SetEmbeds(v []Embed)`

SetEmbeds sets Embeds field to given value.


### GetType

`func (o *CastWithInteractions) GetType() CastNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CastWithInteractions) GetTypeOk() (*CastNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CastWithInteractions) SetType(v CastNotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *CastWithInteractions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFrames

`func (o *CastWithInteractions) GetFrames() []Frame`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *CastWithInteractions) GetFramesOk() (*[]Frame, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *CastWithInteractions) SetFrames(v []Frame)`

SetFrames sets Frames field to given value.

### HasFrames

`func (o *CastWithInteractions) HasFrames() bool`

HasFrames returns a boolean if a field has been set.

### GetReactions

`func (o *CastWithInteractions) GetReactions() CastWithInteractionsReactions`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *CastWithInteractions) GetReactionsOk() (*CastWithInteractionsReactions, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *CastWithInteractions) SetReactions(v CastWithInteractionsReactions)`

SetReactions sets Reactions field to given value.


### GetReplies

`func (o *CastWithInteractions) GetReplies() CastWithInteractionsReplies`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *CastWithInteractions) GetRepliesOk() (*CastWithInteractionsReplies, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *CastWithInteractions) SetReplies(v CastWithInteractionsReplies)`

SetReplies sets Replies field to given value.


### GetThreadHash

`func (o *CastWithInteractions) GetThreadHash() string`

GetThreadHash returns the ThreadHash field if non-nil, zero value otherwise.

### GetThreadHashOk

`func (o *CastWithInteractions) GetThreadHashOk() (*string, bool)`

GetThreadHashOk returns a tuple with the ThreadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadHash

`func (o *CastWithInteractions) SetThreadHash(v string)`

SetThreadHash sets ThreadHash field to given value.


### SetThreadHashNil

`func (o *CastWithInteractions) SetThreadHashNil(b bool)`

 SetThreadHashNil sets the value for ThreadHash to be an explicit nil

### UnsetThreadHash
`func (o *CastWithInteractions) UnsetThreadHash()`

UnsetThreadHash ensures that no value is present for ThreadHash, not even an explicit nil
### GetMentionedProfiles

`func (o *CastWithInteractions) GetMentionedProfiles() []User`

GetMentionedProfiles returns the MentionedProfiles field if non-nil, zero value otherwise.

### GetMentionedProfilesOk

`func (o *CastWithInteractions) GetMentionedProfilesOk() (*[]User, bool)`

GetMentionedProfilesOk returns a tuple with the MentionedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfiles

`func (o *CastWithInteractions) SetMentionedProfiles(v []User)`

SetMentionedProfiles sets MentionedProfiles field to given value.


### GetMentionedProfilesRanges

`func (o *CastWithInteractions) GetMentionedProfilesRanges() []TextRange`

GetMentionedProfilesRanges returns the MentionedProfilesRanges field if non-nil, zero value otherwise.

### GetMentionedProfilesRangesOk

`func (o *CastWithInteractions) GetMentionedProfilesRangesOk() (*[]TextRange, bool)`

GetMentionedProfilesRangesOk returns a tuple with the MentionedProfilesRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfilesRanges

`func (o *CastWithInteractions) SetMentionedProfilesRanges(v []TextRange)`

SetMentionedProfilesRanges sets MentionedProfilesRanges field to given value.


### GetMentionedChannels

`func (o *CastWithInteractions) GetMentionedChannels() []ChannelDehydrated`

GetMentionedChannels returns the MentionedChannels field if non-nil, zero value otherwise.

### GetMentionedChannelsOk

`func (o *CastWithInteractions) GetMentionedChannelsOk() (*[]ChannelDehydrated, bool)`

GetMentionedChannelsOk returns a tuple with the MentionedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedChannels

`func (o *CastWithInteractions) SetMentionedChannels(v []ChannelDehydrated)`

SetMentionedChannels sets MentionedChannels field to given value.


### GetMentionedChannelsRanges

`func (o *CastWithInteractions) GetMentionedChannelsRanges() []TextRange`

GetMentionedChannelsRanges returns the MentionedChannelsRanges field if non-nil, zero value otherwise.

### GetMentionedChannelsRangesOk

`func (o *CastWithInteractions) GetMentionedChannelsRangesOk() (*[]TextRange, bool)`

GetMentionedChannelsRangesOk returns a tuple with the MentionedChannelsRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedChannelsRanges

`func (o *CastWithInteractions) SetMentionedChannelsRanges(v []TextRange)`

SetMentionedChannelsRanges sets MentionedChannelsRanges field to given value.


### GetChannel

`func (o *CastWithInteractions) GetChannel() ChannelOrChannelDehydrated`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CastWithInteractions) GetChannelOk() (*ChannelOrChannelDehydrated, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CastWithInteractions) SetChannel(v ChannelOrChannelDehydrated)`

SetChannel sets Channel field to given value.


### GetViewerContext

`func (o *CastWithInteractions) GetViewerContext() CastViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *CastWithInteractions) GetViewerContextOk() (*CastViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *CastWithInteractions) SetViewerContext(v CastViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *CastWithInteractions) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.

### GetAuthorChannelContext

`func (o *CastWithInteractions) GetAuthorChannelContext() ChannelUserContext`

GetAuthorChannelContext returns the AuthorChannelContext field if non-nil, zero value otherwise.

### GetAuthorChannelContextOk

`func (o *CastWithInteractions) GetAuthorChannelContextOk() (*ChannelUserContext, bool)`

GetAuthorChannelContextOk returns a tuple with the AuthorChannelContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorChannelContext

`func (o *CastWithInteractions) SetAuthorChannelContext(v ChannelUserContext)`

SetAuthorChannelContext sets AuthorChannelContext field to given value.

### HasAuthorChannelContext

`func (o *CastWithInteractions) HasAuthorChannelContext() bool`

HasAuthorChannelContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


