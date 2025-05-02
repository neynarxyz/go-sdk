# CastWithInteractionsAndConversationsRef

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
**DirectReplies** | **[]map[string]interface{}** | note: This is recursive. It contains the direct replies to the cast and their direct replies up to n reply_depth. | 

## Methods

### NewCastWithInteractionsAndConversationsRef

`func NewCastWithInteractionsAndConversationsRef(object string, hash string, parentHash NullableString, parentUrl NullableString, rootParentUrl NullableString, parentAuthor CastParentAuthor, author User, text string, timestamp time.Time, embeds []Embed, reactions CastWithInteractionsReactions, replies CastWithInteractionsReplies, threadHash NullableString, mentionedProfiles []User, mentionedProfilesRanges []TextRange, mentionedChannels []ChannelDehydrated, mentionedChannelsRanges []TextRange, channel ChannelOrChannelDehydrated, directReplies []map[string]interface{}, ) *CastWithInteractionsAndConversationsRef`

NewCastWithInteractionsAndConversationsRef instantiates a new CastWithInteractionsAndConversationsRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastWithInteractionsAndConversationsRefWithDefaults

`func NewCastWithInteractionsAndConversationsRefWithDefaults() *CastWithInteractionsAndConversationsRef`

NewCastWithInteractionsAndConversationsRefWithDefaults instantiates a new CastWithInteractionsAndConversationsRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CastWithInteractionsAndConversationsRef) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CastWithInteractionsAndConversationsRef) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CastWithInteractionsAndConversationsRef) SetObject(v string)`

SetObject sets Object field to given value.


### GetHash

`func (o *CastWithInteractionsAndConversationsRef) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CastWithInteractionsAndConversationsRef) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CastWithInteractionsAndConversationsRef) SetHash(v string)`

SetHash sets Hash field to given value.


### GetParentHash

`func (o *CastWithInteractionsAndConversationsRef) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *CastWithInteractionsAndConversationsRef) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *CastWithInteractionsAndConversationsRef) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### SetParentHashNil

`func (o *CastWithInteractionsAndConversationsRef) SetParentHashNil(b bool)`

 SetParentHashNil sets the value for ParentHash to be an explicit nil

### UnsetParentHash
`func (o *CastWithInteractionsAndConversationsRef) UnsetParentHash()`

UnsetParentHash ensures that no value is present for ParentHash, not even an explicit nil
### GetParentUrl

`func (o *CastWithInteractionsAndConversationsRef) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *CastWithInteractionsAndConversationsRef) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *CastWithInteractionsAndConversationsRef) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.


### SetParentUrlNil

`func (o *CastWithInteractionsAndConversationsRef) SetParentUrlNil(b bool)`

 SetParentUrlNil sets the value for ParentUrl to be an explicit nil

### UnsetParentUrl
`func (o *CastWithInteractionsAndConversationsRef) UnsetParentUrl()`

UnsetParentUrl ensures that no value is present for ParentUrl, not even an explicit nil
### GetRootParentUrl

`func (o *CastWithInteractionsAndConversationsRef) GetRootParentUrl() string`

GetRootParentUrl returns the RootParentUrl field if non-nil, zero value otherwise.

### GetRootParentUrlOk

`func (o *CastWithInteractionsAndConversationsRef) GetRootParentUrlOk() (*string, bool)`

GetRootParentUrlOk returns a tuple with the RootParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootParentUrl

`func (o *CastWithInteractionsAndConversationsRef) SetRootParentUrl(v string)`

SetRootParentUrl sets RootParentUrl field to given value.


### SetRootParentUrlNil

`func (o *CastWithInteractionsAndConversationsRef) SetRootParentUrlNil(b bool)`

 SetRootParentUrlNil sets the value for RootParentUrl to be an explicit nil

### UnsetRootParentUrl
`func (o *CastWithInteractionsAndConversationsRef) UnsetRootParentUrl()`

UnsetRootParentUrl ensures that no value is present for RootParentUrl, not even an explicit nil
### GetParentAuthor

`func (o *CastWithInteractionsAndConversationsRef) GetParentAuthor() CastParentAuthor`

GetParentAuthor returns the ParentAuthor field if non-nil, zero value otherwise.

### GetParentAuthorOk

`func (o *CastWithInteractionsAndConversationsRef) GetParentAuthorOk() (*CastParentAuthor, bool)`

GetParentAuthorOk returns a tuple with the ParentAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAuthor

`func (o *CastWithInteractionsAndConversationsRef) SetParentAuthor(v CastParentAuthor)`

SetParentAuthor sets ParentAuthor field to given value.


### GetAuthor

`func (o *CastWithInteractionsAndConversationsRef) GetAuthor() User`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CastWithInteractionsAndConversationsRef) GetAuthorOk() (*User, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CastWithInteractionsAndConversationsRef) SetAuthor(v User)`

SetAuthor sets Author field to given value.


### GetApp

`func (o *CastWithInteractionsAndConversationsRef) GetApp() UserDehydrated`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CastWithInteractionsAndConversationsRef) GetAppOk() (*UserDehydrated, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CastWithInteractionsAndConversationsRef) SetApp(v UserDehydrated)`

SetApp sets App field to given value.

### HasApp

`func (o *CastWithInteractionsAndConversationsRef) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetText

`func (o *CastWithInteractionsAndConversationsRef) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CastWithInteractionsAndConversationsRef) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CastWithInteractionsAndConversationsRef) SetText(v string)`

SetText sets Text field to given value.


### GetTimestamp

`func (o *CastWithInteractionsAndConversationsRef) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CastWithInteractionsAndConversationsRef) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CastWithInteractionsAndConversationsRef) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEmbeds

`func (o *CastWithInteractionsAndConversationsRef) GetEmbeds() []Embed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *CastWithInteractionsAndConversationsRef) GetEmbedsOk() (*[]Embed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *CastWithInteractionsAndConversationsRef) SetEmbeds(v []Embed)`

SetEmbeds sets Embeds field to given value.


### GetType

`func (o *CastWithInteractionsAndConversationsRef) GetType() CastNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CastWithInteractionsAndConversationsRef) GetTypeOk() (*CastNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CastWithInteractionsAndConversationsRef) SetType(v CastNotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *CastWithInteractionsAndConversationsRef) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFrames

`func (o *CastWithInteractionsAndConversationsRef) GetFrames() []Frame`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *CastWithInteractionsAndConversationsRef) GetFramesOk() (*[]Frame, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *CastWithInteractionsAndConversationsRef) SetFrames(v []Frame)`

SetFrames sets Frames field to given value.

### HasFrames

`func (o *CastWithInteractionsAndConversationsRef) HasFrames() bool`

HasFrames returns a boolean if a field has been set.

### GetReactions

`func (o *CastWithInteractionsAndConversationsRef) GetReactions() CastWithInteractionsReactions`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *CastWithInteractionsAndConversationsRef) GetReactionsOk() (*CastWithInteractionsReactions, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *CastWithInteractionsAndConversationsRef) SetReactions(v CastWithInteractionsReactions)`

SetReactions sets Reactions field to given value.


### GetReplies

`func (o *CastWithInteractionsAndConversationsRef) GetReplies() CastWithInteractionsReplies`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *CastWithInteractionsAndConversationsRef) GetRepliesOk() (*CastWithInteractionsReplies, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *CastWithInteractionsAndConversationsRef) SetReplies(v CastWithInteractionsReplies)`

SetReplies sets Replies field to given value.


### GetThreadHash

`func (o *CastWithInteractionsAndConversationsRef) GetThreadHash() string`

GetThreadHash returns the ThreadHash field if non-nil, zero value otherwise.

### GetThreadHashOk

`func (o *CastWithInteractionsAndConversationsRef) GetThreadHashOk() (*string, bool)`

GetThreadHashOk returns a tuple with the ThreadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadHash

`func (o *CastWithInteractionsAndConversationsRef) SetThreadHash(v string)`

SetThreadHash sets ThreadHash field to given value.


### SetThreadHashNil

`func (o *CastWithInteractionsAndConversationsRef) SetThreadHashNil(b bool)`

 SetThreadHashNil sets the value for ThreadHash to be an explicit nil

### UnsetThreadHash
`func (o *CastWithInteractionsAndConversationsRef) UnsetThreadHash()`

UnsetThreadHash ensures that no value is present for ThreadHash, not even an explicit nil
### GetMentionedProfiles

`func (o *CastWithInteractionsAndConversationsRef) GetMentionedProfiles() []User`

GetMentionedProfiles returns the MentionedProfiles field if non-nil, zero value otherwise.

### GetMentionedProfilesOk

`func (o *CastWithInteractionsAndConversationsRef) GetMentionedProfilesOk() (*[]User, bool)`

GetMentionedProfilesOk returns a tuple with the MentionedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfiles

`func (o *CastWithInteractionsAndConversationsRef) SetMentionedProfiles(v []User)`

SetMentionedProfiles sets MentionedProfiles field to given value.


### GetMentionedProfilesRanges

`func (o *CastWithInteractionsAndConversationsRef) GetMentionedProfilesRanges() []TextRange`

GetMentionedProfilesRanges returns the MentionedProfilesRanges field if non-nil, zero value otherwise.

### GetMentionedProfilesRangesOk

`func (o *CastWithInteractionsAndConversationsRef) GetMentionedProfilesRangesOk() (*[]TextRange, bool)`

GetMentionedProfilesRangesOk returns a tuple with the MentionedProfilesRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfilesRanges

`func (o *CastWithInteractionsAndConversationsRef) SetMentionedProfilesRanges(v []TextRange)`

SetMentionedProfilesRanges sets MentionedProfilesRanges field to given value.


### GetMentionedChannels

`func (o *CastWithInteractionsAndConversationsRef) GetMentionedChannels() []ChannelDehydrated`

GetMentionedChannels returns the MentionedChannels field if non-nil, zero value otherwise.

### GetMentionedChannelsOk

`func (o *CastWithInteractionsAndConversationsRef) GetMentionedChannelsOk() (*[]ChannelDehydrated, bool)`

GetMentionedChannelsOk returns a tuple with the MentionedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedChannels

`func (o *CastWithInteractionsAndConversationsRef) SetMentionedChannels(v []ChannelDehydrated)`

SetMentionedChannels sets MentionedChannels field to given value.


### GetMentionedChannelsRanges

`func (o *CastWithInteractionsAndConversationsRef) GetMentionedChannelsRanges() []TextRange`

GetMentionedChannelsRanges returns the MentionedChannelsRanges field if non-nil, zero value otherwise.

### GetMentionedChannelsRangesOk

`func (o *CastWithInteractionsAndConversationsRef) GetMentionedChannelsRangesOk() (*[]TextRange, bool)`

GetMentionedChannelsRangesOk returns a tuple with the MentionedChannelsRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedChannelsRanges

`func (o *CastWithInteractionsAndConversationsRef) SetMentionedChannelsRanges(v []TextRange)`

SetMentionedChannelsRanges sets MentionedChannelsRanges field to given value.


### GetChannel

`func (o *CastWithInteractionsAndConversationsRef) GetChannel() ChannelOrChannelDehydrated`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CastWithInteractionsAndConversationsRef) GetChannelOk() (*ChannelOrChannelDehydrated, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CastWithInteractionsAndConversationsRef) SetChannel(v ChannelOrChannelDehydrated)`

SetChannel sets Channel field to given value.


### GetViewerContext

`func (o *CastWithInteractionsAndConversationsRef) GetViewerContext() CastViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *CastWithInteractionsAndConversationsRef) GetViewerContextOk() (*CastViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *CastWithInteractionsAndConversationsRef) SetViewerContext(v CastViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *CastWithInteractionsAndConversationsRef) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.

### GetAuthorChannelContext

`func (o *CastWithInteractionsAndConversationsRef) GetAuthorChannelContext() ChannelUserContext`

GetAuthorChannelContext returns the AuthorChannelContext field if non-nil, zero value otherwise.

### GetAuthorChannelContextOk

`func (o *CastWithInteractionsAndConversationsRef) GetAuthorChannelContextOk() (*ChannelUserContext, bool)`

GetAuthorChannelContextOk returns a tuple with the AuthorChannelContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorChannelContext

`func (o *CastWithInteractionsAndConversationsRef) SetAuthorChannelContext(v ChannelUserContext)`

SetAuthorChannelContext sets AuthorChannelContext field to given value.

### HasAuthorChannelContext

`func (o *CastWithInteractionsAndConversationsRef) HasAuthorChannelContext() bool`

HasAuthorChannelContext returns a boolean if a field has been set.

### GetDirectReplies

`func (o *CastWithInteractionsAndConversationsRef) GetDirectReplies() []map[string]interface{}`

GetDirectReplies returns the DirectReplies field if non-nil, zero value otherwise.

### GetDirectRepliesOk

`func (o *CastWithInteractionsAndConversationsRef) GetDirectRepliesOk() (*[]map[string]interface{}, bool)`

GetDirectRepliesOk returns a tuple with the DirectReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectReplies

`func (o *CastWithInteractionsAndConversationsRef) SetDirectReplies(v []map[string]interface{})`

SetDirectReplies sets DirectReplies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


