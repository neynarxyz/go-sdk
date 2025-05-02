# CastWithInteractionsAndConversations

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
**DirectReplies** | [**[]CastWithInteractionsAndConversationsRef**](CastWithInteractionsAndConversationsRef.md) | note: This is recursive. It contains the direct replies to the cast and their direct replies up to n reply_depth. | 

## Methods

### NewCastWithInteractionsAndConversations

`func NewCastWithInteractionsAndConversations(object string, hash string, parentHash NullableString, parentUrl NullableString, rootParentUrl NullableString, parentAuthor CastParentAuthor, author User, text string, timestamp time.Time, embeds []Embed, reactions CastWithInteractionsReactions, replies CastWithInteractionsReplies, threadHash NullableString, mentionedProfiles []User, mentionedProfilesRanges []TextRange, mentionedChannels []ChannelDehydrated, mentionedChannelsRanges []TextRange, channel ChannelOrChannelDehydrated, directReplies []CastWithInteractionsAndConversationsRef, ) *CastWithInteractionsAndConversations`

NewCastWithInteractionsAndConversations instantiates a new CastWithInteractionsAndConversations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastWithInteractionsAndConversationsWithDefaults

`func NewCastWithInteractionsAndConversationsWithDefaults() *CastWithInteractionsAndConversations`

NewCastWithInteractionsAndConversationsWithDefaults instantiates a new CastWithInteractionsAndConversations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CastWithInteractionsAndConversations) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CastWithInteractionsAndConversations) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CastWithInteractionsAndConversations) SetObject(v string)`

SetObject sets Object field to given value.


### GetHash

`func (o *CastWithInteractionsAndConversations) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CastWithInteractionsAndConversations) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CastWithInteractionsAndConversations) SetHash(v string)`

SetHash sets Hash field to given value.


### GetParentHash

`func (o *CastWithInteractionsAndConversations) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *CastWithInteractionsAndConversations) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *CastWithInteractionsAndConversations) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### SetParentHashNil

`func (o *CastWithInteractionsAndConversations) SetParentHashNil(b bool)`

 SetParentHashNil sets the value for ParentHash to be an explicit nil

### UnsetParentHash
`func (o *CastWithInteractionsAndConversations) UnsetParentHash()`

UnsetParentHash ensures that no value is present for ParentHash, not even an explicit nil
### GetParentUrl

`func (o *CastWithInteractionsAndConversations) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *CastWithInteractionsAndConversations) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *CastWithInteractionsAndConversations) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.


### SetParentUrlNil

`func (o *CastWithInteractionsAndConversations) SetParentUrlNil(b bool)`

 SetParentUrlNil sets the value for ParentUrl to be an explicit nil

### UnsetParentUrl
`func (o *CastWithInteractionsAndConversations) UnsetParentUrl()`

UnsetParentUrl ensures that no value is present for ParentUrl, not even an explicit nil
### GetRootParentUrl

`func (o *CastWithInteractionsAndConversations) GetRootParentUrl() string`

GetRootParentUrl returns the RootParentUrl field if non-nil, zero value otherwise.

### GetRootParentUrlOk

`func (o *CastWithInteractionsAndConversations) GetRootParentUrlOk() (*string, bool)`

GetRootParentUrlOk returns a tuple with the RootParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootParentUrl

`func (o *CastWithInteractionsAndConversations) SetRootParentUrl(v string)`

SetRootParentUrl sets RootParentUrl field to given value.


### SetRootParentUrlNil

`func (o *CastWithInteractionsAndConversations) SetRootParentUrlNil(b bool)`

 SetRootParentUrlNil sets the value for RootParentUrl to be an explicit nil

### UnsetRootParentUrl
`func (o *CastWithInteractionsAndConversations) UnsetRootParentUrl()`

UnsetRootParentUrl ensures that no value is present for RootParentUrl, not even an explicit nil
### GetParentAuthor

`func (o *CastWithInteractionsAndConversations) GetParentAuthor() CastParentAuthor`

GetParentAuthor returns the ParentAuthor field if non-nil, zero value otherwise.

### GetParentAuthorOk

`func (o *CastWithInteractionsAndConversations) GetParentAuthorOk() (*CastParentAuthor, bool)`

GetParentAuthorOk returns a tuple with the ParentAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAuthor

`func (o *CastWithInteractionsAndConversations) SetParentAuthor(v CastParentAuthor)`

SetParentAuthor sets ParentAuthor field to given value.


### GetAuthor

`func (o *CastWithInteractionsAndConversations) GetAuthor() User`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CastWithInteractionsAndConversations) GetAuthorOk() (*User, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CastWithInteractionsAndConversations) SetAuthor(v User)`

SetAuthor sets Author field to given value.


### GetApp

`func (o *CastWithInteractionsAndConversations) GetApp() UserDehydrated`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CastWithInteractionsAndConversations) GetAppOk() (*UserDehydrated, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CastWithInteractionsAndConversations) SetApp(v UserDehydrated)`

SetApp sets App field to given value.

### HasApp

`func (o *CastWithInteractionsAndConversations) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetText

`func (o *CastWithInteractionsAndConversations) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CastWithInteractionsAndConversations) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CastWithInteractionsAndConversations) SetText(v string)`

SetText sets Text field to given value.


### GetTimestamp

`func (o *CastWithInteractionsAndConversations) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CastWithInteractionsAndConversations) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CastWithInteractionsAndConversations) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEmbeds

`func (o *CastWithInteractionsAndConversations) GetEmbeds() []Embed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *CastWithInteractionsAndConversations) GetEmbedsOk() (*[]Embed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *CastWithInteractionsAndConversations) SetEmbeds(v []Embed)`

SetEmbeds sets Embeds field to given value.


### GetType

`func (o *CastWithInteractionsAndConversations) GetType() CastNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CastWithInteractionsAndConversations) GetTypeOk() (*CastNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CastWithInteractionsAndConversations) SetType(v CastNotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *CastWithInteractionsAndConversations) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFrames

`func (o *CastWithInteractionsAndConversations) GetFrames() []Frame`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *CastWithInteractionsAndConversations) GetFramesOk() (*[]Frame, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *CastWithInteractionsAndConversations) SetFrames(v []Frame)`

SetFrames sets Frames field to given value.

### HasFrames

`func (o *CastWithInteractionsAndConversations) HasFrames() bool`

HasFrames returns a boolean if a field has been set.

### GetReactions

`func (o *CastWithInteractionsAndConversations) GetReactions() CastWithInteractionsReactions`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *CastWithInteractionsAndConversations) GetReactionsOk() (*CastWithInteractionsReactions, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *CastWithInteractionsAndConversations) SetReactions(v CastWithInteractionsReactions)`

SetReactions sets Reactions field to given value.


### GetReplies

`func (o *CastWithInteractionsAndConversations) GetReplies() CastWithInteractionsReplies`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *CastWithInteractionsAndConversations) GetRepliesOk() (*CastWithInteractionsReplies, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *CastWithInteractionsAndConversations) SetReplies(v CastWithInteractionsReplies)`

SetReplies sets Replies field to given value.


### GetThreadHash

`func (o *CastWithInteractionsAndConversations) GetThreadHash() string`

GetThreadHash returns the ThreadHash field if non-nil, zero value otherwise.

### GetThreadHashOk

`func (o *CastWithInteractionsAndConversations) GetThreadHashOk() (*string, bool)`

GetThreadHashOk returns a tuple with the ThreadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadHash

`func (o *CastWithInteractionsAndConversations) SetThreadHash(v string)`

SetThreadHash sets ThreadHash field to given value.


### SetThreadHashNil

`func (o *CastWithInteractionsAndConversations) SetThreadHashNil(b bool)`

 SetThreadHashNil sets the value for ThreadHash to be an explicit nil

### UnsetThreadHash
`func (o *CastWithInteractionsAndConversations) UnsetThreadHash()`

UnsetThreadHash ensures that no value is present for ThreadHash, not even an explicit nil
### GetMentionedProfiles

`func (o *CastWithInteractionsAndConversations) GetMentionedProfiles() []User`

GetMentionedProfiles returns the MentionedProfiles field if non-nil, zero value otherwise.

### GetMentionedProfilesOk

`func (o *CastWithInteractionsAndConversations) GetMentionedProfilesOk() (*[]User, bool)`

GetMentionedProfilesOk returns a tuple with the MentionedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfiles

`func (o *CastWithInteractionsAndConversations) SetMentionedProfiles(v []User)`

SetMentionedProfiles sets MentionedProfiles field to given value.


### GetMentionedProfilesRanges

`func (o *CastWithInteractionsAndConversations) GetMentionedProfilesRanges() []TextRange`

GetMentionedProfilesRanges returns the MentionedProfilesRanges field if non-nil, zero value otherwise.

### GetMentionedProfilesRangesOk

`func (o *CastWithInteractionsAndConversations) GetMentionedProfilesRangesOk() (*[]TextRange, bool)`

GetMentionedProfilesRangesOk returns a tuple with the MentionedProfilesRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfilesRanges

`func (o *CastWithInteractionsAndConversations) SetMentionedProfilesRanges(v []TextRange)`

SetMentionedProfilesRanges sets MentionedProfilesRanges field to given value.


### GetMentionedChannels

`func (o *CastWithInteractionsAndConversations) GetMentionedChannels() []ChannelDehydrated`

GetMentionedChannels returns the MentionedChannels field if non-nil, zero value otherwise.

### GetMentionedChannelsOk

`func (o *CastWithInteractionsAndConversations) GetMentionedChannelsOk() (*[]ChannelDehydrated, bool)`

GetMentionedChannelsOk returns a tuple with the MentionedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedChannels

`func (o *CastWithInteractionsAndConversations) SetMentionedChannels(v []ChannelDehydrated)`

SetMentionedChannels sets MentionedChannels field to given value.


### GetMentionedChannelsRanges

`func (o *CastWithInteractionsAndConversations) GetMentionedChannelsRanges() []TextRange`

GetMentionedChannelsRanges returns the MentionedChannelsRanges field if non-nil, zero value otherwise.

### GetMentionedChannelsRangesOk

`func (o *CastWithInteractionsAndConversations) GetMentionedChannelsRangesOk() (*[]TextRange, bool)`

GetMentionedChannelsRangesOk returns a tuple with the MentionedChannelsRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedChannelsRanges

`func (o *CastWithInteractionsAndConversations) SetMentionedChannelsRanges(v []TextRange)`

SetMentionedChannelsRanges sets MentionedChannelsRanges field to given value.


### GetChannel

`func (o *CastWithInteractionsAndConversations) GetChannel() ChannelOrChannelDehydrated`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CastWithInteractionsAndConversations) GetChannelOk() (*ChannelOrChannelDehydrated, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CastWithInteractionsAndConversations) SetChannel(v ChannelOrChannelDehydrated)`

SetChannel sets Channel field to given value.


### GetViewerContext

`func (o *CastWithInteractionsAndConversations) GetViewerContext() CastViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *CastWithInteractionsAndConversations) GetViewerContextOk() (*CastViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *CastWithInteractionsAndConversations) SetViewerContext(v CastViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *CastWithInteractionsAndConversations) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.

### GetAuthorChannelContext

`func (o *CastWithInteractionsAndConversations) GetAuthorChannelContext() ChannelUserContext`

GetAuthorChannelContext returns the AuthorChannelContext field if non-nil, zero value otherwise.

### GetAuthorChannelContextOk

`func (o *CastWithInteractionsAndConversations) GetAuthorChannelContextOk() (*ChannelUserContext, bool)`

GetAuthorChannelContextOk returns a tuple with the AuthorChannelContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorChannelContext

`func (o *CastWithInteractionsAndConversations) SetAuthorChannelContext(v ChannelUserContext)`

SetAuthorChannelContext sets AuthorChannelContext field to given value.

### HasAuthorChannelContext

`func (o *CastWithInteractionsAndConversations) HasAuthorChannelContext() bool`

HasAuthorChannelContext returns a boolean if a field has been set.

### GetDirectReplies

`func (o *CastWithInteractionsAndConversations) GetDirectReplies() []CastWithInteractionsAndConversationsRef`

GetDirectReplies returns the DirectReplies field if non-nil, zero value otherwise.

### GetDirectRepliesOk

`func (o *CastWithInteractionsAndConversations) GetDirectRepliesOk() (*[]CastWithInteractionsAndConversationsRef, bool)`

GetDirectRepliesOk returns a tuple with the DirectReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectReplies

`func (o *CastWithInteractionsAndConversations) SetDirectReplies(v []CastWithInteractionsAndConversationsRef)`

SetDirectReplies sets DirectReplies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


