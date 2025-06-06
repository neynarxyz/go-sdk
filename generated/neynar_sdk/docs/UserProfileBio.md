# UserProfileBio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**MentionedProfiles** | Pointer to [**[]UserDehydrated**](UserDehydrated.md) |  | [optional] 
**MentionedProfilesRanges** | Pointer to [**[]TextRange**](TextRange.md) | Positions within the text (inclusive start, exclusive end) where each mention occurs. Each index within this list corresponds to the same-numbered index in the mentioned_profiles list. | [optional] 
**MentionedChannels** | Pointer to [**[]ChannelDehydrated**](ChannelDehydrated.md) |  | [optional] 
**MentionedChannelsRanges** | Pointer to [**[]TextRange**](TextRange.md) | Positions within the text (inclusive start, exclusive end) where each mention occurs. Each index within this list corresponds to the same-numbered index in the mentioned_channels list. | [optional] 

## Methods

### NewUserProfileBio

`func NewUserProfileBio(text string, ) *UserProfileBio`

NewUserProfileBio instantiates a new UserProfileBio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileBioWithDefaults

`func NewUserProfileBioWithDefaults() *UserProfileBio`

NewUserProfileBioWithDefaults instantiates a new UserProfileBio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *UserProfileBio) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *UserProfileBio) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *UserProfileBio) SetText(v string)`

SetText sets Text field to given value.


### GetMentionedProfiles

`func (o *UserProfileBio) GetMentionedProfiles() []UserDehydrated`

GetMentionedProfiles returns the MentionedProfiles field if non-nil, zero value otherwise.

### GetMentionedProfilesOk

`func (o *UserProfileBio) GetMentionedProfilesOk() (*[]UserDehydrated, bool)`

GetMentionedProfilesOk returns a tuple with the MentionedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfiles

`func (o *UserProfileBio) SetMentionedProfiles(v []UserDehydrated)`

SetMentionedProfiles sets MentionedProfiles field to given value.

### HasMentionedProfiles

`func (o *UserProfileBio) HasMentionedProfiles() bool`

HasMentionedProfiles returns a boolean if a field has been set.

### GetMentionedProfilesRanges

`func (o *UserProfileBio) GetMentionedProfilesRanges() []TextRange`

GetMentionedProfilesRanges returns the MentionedProfilesRanges field if non-nil, zero value otherwise.

### GetMentionedProfilesRangesOk

`func (o *UserProfileBio) GetMentionedProfilesRangesOk() (*[]TextRange, bool)`

GetMentionedProfilesRangesOk returns a tuple with the MentionedProfilesRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfilesRanges

`func (o *UserProfileBio) SetMentionedProfilesRanges(v []TextRange)`

SetMentionedProfilesRanges sets MentionedProfilesRanges field to given value.

### HasMentionedProfilesRanges

`func (o *UserProfileBio) HasMentionedProfilesRanges() bool`

HasMentionedProfilesRanges returns a boolean if a field has been set.

### GetMentionedChannels

`func (o *UserProfileBio) GetMentionedChannels() []ChannelDehydrated`

GetMentionedChannels returns the MentionedChannels field if non-nil, zero value otherwise.

### GetMentionedChannelsOk

`func (o *UserProfileBio) GetMentionedChannelsOk() (*[]ChannelDehydrated, bool)`

GetMentionedChannelsOk returns a tuple with the MentionedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedChannels

`func (o *UserProfileBio) SetMentionedChannels(v []ChannelDehydrated)`

SetMentionedChannels sets MentionedChannels field to given value.

### HasMentionedChannels

`func (o *UserProfileBio) HasMentionedChannels() bool`

HasMentionedChannels returns a boolean if a field has been set.

### GetMentionedChannelsRanges

`func (o *UserProfileBio) GetMentionedChannelsRanges() []TextRange`

GetMentionedChannelsRanges returns the MentionedChannelsRanges field if non-nil, zero value otherwise.

### GetMentionedChannelsRangesOk

`func (o *UserProfileBio) GetMentionedChannelsRangesOk() (*[]TextRange, bool)`

GetMentionedChannelsRangesOk returns a tuple with the MentionedChannelsRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedChannelsRanges

`func (o *UserProfileBio) SetMentionedChannelsRanges(v []TextRange)`

SetMentionedChannelsRanges sets MentionedChannelsRanges field to given value.

### HasMentionedChannelsRanges

`func (o *UserProfileBio) HasMentionedChannelsRanges() bool`

HasMentionedChannelsRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


