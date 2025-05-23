/*
Farcaster API V2

The Farcaster API allows you to interact with the Farcaster protocol. See the [Neynar docs](https://docs.neynar.com/reference) for more details.

API version: 2.43.0
Contact: team@neynar.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neynar_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the CastWithInteractions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastWithInteractions{}

// CastWithInteractions struct for CastWithInteractions
type CastWithInteractions struct {
	Object            string                        `json:"object"`
	Hash              string                        `json:"hash"`
	ParentHash        NullableString                `json:"parent_hash"`
	ParentUrl         NullableString                `json:"parent_url"`
	RootParentUrl     NullableString                `json:"root_parent_url"`
	ParentAuthor      CastParentAuthor              `json:"parent_author"`
	Author            User                          `json:"author"`
	App               NullableUserDehydrated        `json:"app,omitempty"`
	Text              string                        `json:"text"`
	Timestamp         time.Time                     `json:"timestamp"`
	Embeds            []Embed                       `json:"embeds"`
	Type              *CastNotificationType         `json:"type,omitempty"`
	Frames            []Frame                       `json:"frames,omitempty"`
	Reactions         CastWithInteractionsReactions `json:"reactions"`
	Replies           CastWithInteractionsReplies   `json:"replies"`
	ThreadHash        NullableString                `json:"thread_hash"`
	MentionedProfiles []User                        `json:"mentioned_profiles"`
	// Positions within the text (inclusive start, exclusive end) where each mention occurs. Each index within this list corresponds to the same-numbered index in the mentioned_profiles list.
	MentionedProfilesRanges []TextRange         `json:"mentioned_profiles_ranges"`
	MentionedChannels       []ChannelDehydrated `json:"mentioned_channels"`
	// Positions within the text (inclusive start, exclusive end) where each mention occurs. Each index within this list corresponds to the same-numbered index in the mentioned_channels list.
	MentionedChannelsRanges []TextRange                        `json:"mentioned_channels_ranges"`
	Channel                 NullableChannelOrChannelDehydrated `json:"channel"`
	ViewerContext           *CastViewerContext                 `json:"viewer_context,omitempty"`
	AuthorChannelContext    *ChannelUserContext                `json:"author_channel_context,omitempty"`
}

type _CastWithInteractions CastWithInteractions

// NewCastWithInteractions instantiates a new CastWithInteractions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastWithInteractions(object string, hash string, parentHash NullableString, parentUrl NullableString, rootParentUrl NullableString, parentAuthor CastParentAuthor, author User, text string, timestamp time.Time, embeds []Embed, reactions CastWithInteractionsReactions, replies CastWithInteractionsReplies, threadHash NullableString, mentionedProfiles []User, mentionedProfilesRanges []TextRange, mentionedChannels []ChannelDehydrated, mentionedChannelsRanges []TextRange, channel NullableChannelOrChannelDehydrated) *CastWithInteractions {
	this := CastWithInteractions{}
	this.Object = object
	this.Hash = hash
	this.ParentHash = parentHash
	this.ParentUrl = parentUrl
	this.RootParentUrl = rootParentUrl
	this.ParentAuthor = parentAuthor
	this.Author = author
	this.Text = text
	this.Timestamp = timestamp
	this.Embeds = embeds
	this.Reactions = reactions
	this.Replies = replies
	this.ThreadHash = threadHash
	this.MentionedProfiles = mentionedProfiles
	this.MentionedProfilesRanges = mentionedProfilesRanges
	this.MentionedChannels = mentionedChannels
	this.MentionedChannelsRanges = mentionedChannelsRanges
	this.Channel = channel
	return &this
}

// NewCastWithInteractionsWithDefaults instantiates a new CastWithInteractions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastWithInteractionsWithDefaults() *CastWithInteractions {
	this := CastWithInteractions{}
	return &this
}

// GetObject returns the Object field value
func (o *CastWithInteractions) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CastWithInteractions) SetObject(v string) {
	o.Object = v
}

// GetHash returns the Hash field value
func (o *CastWithInteractions) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *CastWithInteractions) SetHash(v string) {
	o.Hash = v
}

// GetParentHash returns the ParentHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CastWithInteractions) GetParentHash() string {
	if o == nil || o.ParentHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentHash.Get()
}

// GetParentHashOk returns a tuple with the ParentHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractions) GetParentHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentHash.Get(), o.ParentHash.IsSet()
}

// SetParentHash sets field value
func (o *CastWithInteractions) SetParentHash(v string) {
	o.ParentHash.Set(&v)
}

// GetParentUrl returns the ParentUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CastWithInteractions) GetParentUrl() string {
	if o == nil || o.ParentUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentUrl.Get()
}

// GetParentUrlOk returns a tuple with the ParentUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractions) GetParentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentUrl.Get(), o.ParentUrl.IsSet()
}

// SetParentUrl sets field value
func (o *CastWithInteractions) SetParentUrl(v string) {
	o.ParentUrl.Set(&v)
}

// GetRootParentUrl returns the RootParentUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CastWithInteractions) GetRootParentUrl() string {
	if o == nil || o.RootParentUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.RootParentUrl.Get()
}

// GetRootParentUrlOk returns a tuple with the RootParentUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractions) GetRootParentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootParentUrl.Get(), o.RootParentUrl.IsSet()
}

// SetRootParentUrl sets field value
func (o *CastWithInteractions) SetRootParentUrl(v string) {
	o.RootParentUrl.Set(&v)
}

// GetParentAuthor returns the ParentAuthor field value
func (o *CastWithInteractions) GetParentAuthor() CastParentAuthor {
	if o == nil {
		var ret CastParentAuthor
		return ret
	}

	return o.ParentAuthor
}

// GetParentAuthorOk returns a tuple with the ParentAuthor field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetParentAuthorOk() (*CastParentAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentAuthor, true
}

// SetParentAuthor sets field value
func (o *CastWithInteractions) SetParentAuthor(v CastParentAuthor) {
	o.ParentAuthor = v
}

// GetAuthor returns the Author field value
func (o *CastWithInteractions) GetAuthor() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetAuthorOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *CastWithInteractions) SetAuthor(v User) {
	o.Author = v
}

// GetApp returns the App field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CastWithInteractions) GetApp() UserDehydrated {
	if o == nil || IsNil(o.App.Get()) {
		var ret UserDehydrated
		return ret
	}
	return *o.App.Get()
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractions) GetAppOk() (*UserDehydrated, bool) {
	if o == nil {
		return nil, false
	}
	return o.App.Get(), o.App.IsSet()
}

// HasApp returns a boolean if a field has been set.
func (o *CastWithInteractions) HasApp() bool {
	if o != nil && o.App.IsSet() {
		return true
	}

	return false
}

// SetApp gets a reference to the given NullableUserDehydrated and assigns it to the App field.
func (o *CastWithInteractions) SetApp(v UserDehydrated) {
	o.App.Set(&v)
}

// SetAppNil sets the value for App to be an explicit nil
func (o *CastWithInteractions) SetAppNil() {
	o.App.Set(nil)
}

// UnsetApp ensures that no value is present for App, not even an explicit nil
func (o *CastWithInteractions) UnsetApp() {
	o.App.Unset()
}

// GetText returns the Text field value
func (o *CastWithInteractions) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *CastWithInteractions) SetText(v string) {
	o.Text = v
}

// GetTimestamp returns the Timestamp field value
func (o *CastWithInteractions) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *CastWithInteractions) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetEmbeds returns the Embeds field value
func (o *CastWithInteractions) GetEmbeds() []Embed {
	if o == nil {
		var ret []Embed
		return ret
	}

	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetEmbedsOk() ([]Embed, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embeds, true
}

// SetEmbeds sets field value
func (o *CastWithInteractions) SetEmbeds(v []Embed) {
	o.Embeds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CastWithInteractions) GetType() CastNotificationType {
	if o == nil || IsNil(o.Type) {
		var ret CastNotificationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetTypeOk() (*CastNotificationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CastWithInteractions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CastNotificationType and assigns it to the Type field.
func (o *CastWithInteractions) SetType(v CastNotificationType) {
	o.Type = &v
}

// GetFrames returns the Frames field value if set, zero value otherwise.
func (o *CastWithInteractions) GetFrames() []Frame {
	if o == nil || IsNil(o.Frames) {
		var ret []Frame
		return ret
	}
	return o.Frames
}

// GetFramesOk returns a tuple with the Frames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetFramesOk() ([]Frame, bool) {
	if o == nil || IsNil(o.Frames) {
		return nil, false
	}
	return o.Frames, true
}

// HasFrames returns a boolean if a field has been set.
func (o *CastWithInteractions) HasFrames() bool {
	if o != nil && !IsNil(o.Frames) {
		return true
	}

	return false
}

// SetFrames gets a reference to the given []Frame and assigns it to the Frames field.
func (o *CastWithInteractions) SetFrames(v []Frame) {
	o.Frames = v
}

// GetReactions returns the Reactions field value
func (o *CastWithInteractions) GetReactions() CastWithInteractionsReactions {
	if o == nil {
		var ret CastWithInteractionsReactions
		return ret
	}

	return o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetReactionsOk() (*CastWithInteractionsReactions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reactions, true
}

// SetReactions sets field value
func (o *CastWithInteractions) SetReactions(v CastWithInteractionsReactions) {
	o.Reactions = v
}

// GetReplies returns the Replies field value
func (o *CastWithInteractions) GetReplies() CastWithInteractionsReplies {
	if o == nil {
		var ret CastWithInteractionsReplies
		return ret
	}

	return o.Replies
}

// GetRepliesOk returns a tuple with the Replies field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetRepliesOk() (*CastWithInteractionsReplies, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replies, true
}

// SetReplies sets field value
func (o *CastWithInteractions) SetReplies(v CastWithInteractionsReplies) {
	o.Replies = v
}

// GetThreadHash returns the ThreadHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CastWithInteractions) GetThreadHash() string {
	if o == nil || o.ThreadHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.ThreadHash.Get()
}

// GetThreadHashOk returns a tuple with the ThreadHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractions) GetThreadHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreadHash.Get(), o.ThreadHash.IsSet()
}

// SetThreadHash sets field value
func (o *CastWithInteractions) SetThreadHash(v string) {
	o.ThreadHash.Set(&v)
}

// GetMentionedProfiles returns the MentionedProfiles field value
func (o *CastWithInteractions) GetMentionedProfiles() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.MentionedProfiles
}

// GetMentionedProfilesOk returns a tuple with the MentionedProfiles field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetMentionedProfilesOk() ([]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionedProfiles, true
}

// SetMentionedProfiles sets field value
func (o *CastWithInteractions) SetMentionedProfiles(v []User) {
	o.MentionedProfiles = v
}

// GetMentionedProfilesRanges returns the MentionedProfilesRanges field value
func (o *CastWithInteractions) GetMentionedProfilesRanges() []TextRange {
	if o == nil {
		var ret []TextRange
		return ret
	}

	return o.MentionedProfilesRanges
}

// GetMentionedProfilesRangesOk returns a tuple with the MentionedProfilesRanges field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetMentionedProfilesRangesOk() ([]TextRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionedProfilesRanges, true
}

// SetMentionedProfilesRanges sets field value
func (o *CastWithInteractions) SetMentionedProfilesRanges(v []TextRange) {
	o.MentionedProfilesRanges = v
}

// GetMentionedChannels returns the MentionedChannels field value
func (o *CastWithInteractions) GetMentionedChannels() []ChannelDehydrated {
	if o == nil {
		var ret []ChannelDehydrated
		return ret
	}

	return o.MentionedChannels
}

// GetMentionedChannelsOk returns a tuple with the MentionedChannels field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetMentionedChannelsOk() ([]ChannelDehydrated, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionedChannels, true
}

// SetMentionedChannels sets field value
func (o *CastWithInteractions) SetMentionedChannels(v []ChannelDehydrated) {
	o.MentionedChannels = v
}

// GetMentionedChannelsRanges returns the MentionedChannelsRanges field value
func (o *CastWithInteractions) GetMentionedChannelsRanges() []TextRange {
	if o == nil {
		var ret []TextRange
		return ret
	}

	return o.MentionedChannelsRanges
}

// GetMentionedChannelsRangesOk returns a tuple with the MentionedChannelsRanges field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetMentionedChannelsRangesOk() ([]TextRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionedChannelsRanges, true
}

// SetMentionedChannelsRanges sets field value
func (o *CastWithInteractions) SetMentionedChannelsRanges(v []TextRange) {
	o.MentionedChannelsRanges = v
}

// GetChannel returns the Channel field value
// If the value is explicit nil, the zero value for ChannelOrChannelDehydrated will be returned
func (o *CastWithInteractions) GetChannel() ChannelOrChannelDehydrated {
	if o == nil || o.Channel.Get() == nil {
		var ret ChannelOrChannelDehydrated
		return ret
	}

	return *o.Channel.Get()
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractions) GetChannelOk() (*ChannelOrChannelDehydrated, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channel.Get(), o.Channel.IsSet()
}

// SetChannel sets field value
func (o *CastWithInteractions) SetChannel(v ChannelOrChannelDehydrated) {
	o.Channel.Set(&v)
}

// GetViewerContext returns the ViewerContext field value if set, zero value otherwise.
func (o *CastWithInteractions) GetViewerContext() CastViewerContext {
	if o == nil || IsNil(o.ViewerContext) {
		var ret CastViewerContext
		return ret
	}
	return *o.ViewerContext
}

// GetViewerContextOk returns a tuple with the ViewerContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetViewerContextOk() (*CastViewerContext, bool) {
	if o == nil || IsNil(o.ViewerContext) {
		return nil, false
	}
	return o.ViewerContext, true
}

// HasViewerContext returns a boolean if a field has been set.
func (o *CastWithInteractions) HasViewerContext() bool {
	if o != nil && !IsNil(o.ViewerContext) {
		return true
	}

	return false
}

// SetViewerContext gets a reference to the given CastViewerContext and assigns it to the ViewerContext field.
func (o *CastWithInteractions) SetViewerContext(v CastViewerContext) {
	o.ViewerContext = &v
}

// GetAuthorChannelContext returns the AuthorChannelContext field value if set, zero value otherwise.
func (o *CastWithInteractions) GetAuthorChannelContext() ChannelUserContext {
	if o == nil || IsNil(o.AuthorChannelContext) {
		var ret ChannelUserContext
		return ret
	}
	return *o.AuthorChannelContext
}

// GetAuthorChannelContextOk returns a tuple with the AuthorChannelContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetAuthorChannelContextOk() (*ChannelUserContext, bool) {
	if o == nil || IsNil(o.AuthorChannelContext) {
		return nil, false
	}
	return o.AuthorChannelContext, true
}

// HasAuthorChannelContext returns a boolean if a field has been set.
func (o *CastWithInteractions) HasAuthorChannelContext() bool {
	if o != nil && !IsNil(o.AuthorChannelContext) {
		return true
	}

	return false
}

// SetAuthorChannelContext gets a reference to the given ChannelUserContext and assigns it to the AuthorChannelContext field.
func (o *CastWithInteractions) SetAuthorChannelContext(v ChannelUserContext) {
	o.AuthorChannelContext = &v
}

func (o CastWithInteractions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastWithInteractions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["hash"] = o.Hash
	toSerialize["parent_hash"] = o.ParentHash.Get()
	toSerialize["parent_url"] = o.ParentUrl.Get()
	toSerialize["root_parent_url"] = o.RootParentUrl.Get()
	toSerialize["parent_author"] = o.ParentAuthor
	toSerialize["author"] = o.Author
	if o.App.IsSet() {
		toSerialize["app"] = o.App.Get()
	}
	toSerialize["text"] = o.Text
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["embeds"] = o.Embeds
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Frames) {
		toSerialize["frames"] = o.Frames
	}
	toSerialize["reactions"] = o.Reactions
	toSerialize["replies"] = o.Replies
	toSerialize["thread_hash"] = o.ThreadHash.Get()
	toSerialize["mentioned_profiles"] = o.MentionedProfiles
	toSerialize["mentioned_profiles_ranges"] = o.MentionedProfilesRanges
	toSerialize["mentioned_channels"] = o.MentionedChannels
	toSerialize["mentioned_channels_ranges"] = o.MentionedChannelsRanges
	toSerialize["channel"] = o.Channel.Get()
	if !IsNil(o.ViewerContext) {
		toSerialize["viewer_context"] = o.ViewerContext
	}
	if !IsNil(o.AuthorChannelContext) {
		toSerialize["author_channel_context"] = o.AuthorChannelContext
	}
	return toSerialize, nil
}

func (o *CastWithInteractions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"hash",
		"parent_hash",
		"parent_url",
		"root_parent_url",
		"parent_author",
		"author",
		"text",
		"timestamp",
		"embeds",
		"reactions",
		"replies",
		"thread_hash",
		"mentioned_profiles",
		"mentioned_profiles_ranges",
		"mentioned_channels",
		"mentioned_channels_ranges",
		"channel",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	varCastWithInteractions := _CastWithInteractions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCastWithInteractions)

	if err != nil {
		return err
	}

	*o = CastWithInteractions(varCastWithInteractions)

	return err
}

type NullableCastWithInteractions struct {
	value *CastWithInteractions
	isSet bool
}

func (v NullableCastWithInteractions) Get() *CastWithInteractions {
	return v.value
}

func (v *NullableCastWithInteractions) Set(val *CastWithInteractions) {
	v.value = val
	v.isSet = true
}

func (v NullableCastWithInteractions) IsSet() bool {
	return v.isSet
}

func (v *NullableCastWithInteractions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastWithInteractions(val *CastWithInteractions) *NullableCastWithInteractions {
	return &NullableCastWithInteractions{value: val, isSet: true}
}

func (v NullableCastWithInteractions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastWithInteractions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
