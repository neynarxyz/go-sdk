# CastEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**ParentHash** | **NullableString** |  | 
**ParentUrl** | **NullableString** |  | 
**RootParentUrl** | **NullableString** |  | 
**ParentAuthor** | [**CastEmbeddedParentAuthor**](CastEmbeddedParentAuthor.md) |  | 
**Author** | [**UserDehydrated**](UserDehydrated.md) |  | 
**App** | Pointer to [**NullableUserDehydrated**](UserDehydrated.md) |  | [optional] 
**Text** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Embeds** | [**[]EmbedDeep**](EmbedDeep.md) |  | 
**Channel** | [**NullableChannelDehydrated**](ChannelDehydrated.md) |  | 

## Methods

### NewCastEmbedded

`func NewCastEmbedded(hash string, parentHash NullableString, parentUrl NullableString, rootParentUrl NullableString, parentAuthor CastEmbeddedParentAuthor, author UserDehydrated, text string, timestamp time.Time, embeds []EmbedDeep, channel NullableChannelDehydrated, ) *CastEmbedded`

NewCastEmbedded instantiates a new CastEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastEmbeddedWithDefaults

`func NewCastEmbeddedWithDefaults() *CastEmbedded`

NewCastEmbeddedWithDefaults instantiates a new CastEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *CastEmbedded) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CastEmbedded) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CastEmbedded) SetHash(v string)`

SetHash sets Hash field to given value.


### GetParentHash

`func (o *CastEmbedded) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *CastEmbedded) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *CastEmbedded) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### SetParentHashNil

`func (o *CastEmbedded) SetParentHashNil(b bool)`

 SetParentHashNil sets the value for ParentHash to be an explicit nil

### UnsetParentHash
`func (o *CastEmbedded) UnsetParentHash()`

UnsetParentHash ensures that no value is present for ParentHash, not even an explicit nil
### GetParentUrl

`func (o *CastEmbedded) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *CastEmbedded) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *CastEmbedded) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.


### SetParentUrlNil

`func (o *CastEmbedded) SetParentUrlNil(b bool)`

 SetParentUrlNil sets the value for ParentUrl to be an explicit nil

### UnsetParentUrl
`func (o *CastEmbedded) UnsetParentUrl()`

UnsetParentUrl ensures that no value is present for ParentUrl, not even an explicit nil
### GetRootParentUrl

`func (o *CastEmbedded) GetRootParentUrl() string`

GetRootParentUrl returns the RootParentUrl field if non-nil, zero value otherwise.

### GetRootParentUrlOk

`func (o *CastEmbedded) GetRootParentUrlOk() (*string, bool)`

GetRootParentUrlOk returns a tuple with the RootParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootParentUrl

`func (o *CastEmbedded) SetRootParentUrl(v string)`

SetRootParentUrl sets RootParentUrl field to given value.


### SetRootParentUrlNil

`func (o *CastEmbedded) SetRootParentUrlNil(b bool)`

 SetRootParentUrlNil sets the value for RootParentUrl to be an explicit nil

### UnsetRootParentUrl
`func (o *CastEmbedded) UnsetRootParentUrl()`

UnsetRootParentUrl ensures that no value is present for RootParentUrl, not even an explicit nil
### GetParentAuthor

`func (o *CastEmbedded) GetParentAuthor() CastEmbeddedParentAuthor`

GetParentAuthor returns the ParentAuthor field if non-nil, zero value otherwise.

### GetParentAuthorOk

`func (o *CastEmbedded) GetParentAuthorOk() (*CastEmbeddedParentAuthor, bool)`

GetParentAuthorOk returns a tuple with the ParentAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAuthor

`func (o *CastEmbedded) SetParentAuthor(v CastEmbeddedParentAuthor)`

SetParentAuthor sets ParentAuthor field to given value.


### GetAuthor

`func (o *CastEmbedded) GetAuthor() UserDehydrated`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CastEmbedded) GetAuthorOk() (*UserDehydrated, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CastEmbedded) SetAuthor(v UserDehydrated)`

SetAuthor sets Author field to given value.


### GetApp

`func (o *CastEmbedded) GetApp() UserDehydrated`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CastEmbedded) GetAppOk() (*UserDehydrated, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CastEmbedded) SetApp(v UserDehydrated)`

SetApp sets App field to given value.

### HasApp

`func (o *CastEmbedded) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetAppNil

`func (o *CastEmbedded) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *CastEmbedded) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetText

`func (o *CastEmbedded) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CastEmbedded) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CastEmbedded) SetText(v string)`

SetText sets Text field to given value.


### GetTimestamp

`func (o *CastEmbedded) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CastEmbedded) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CastEmbedded) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEmbeds

`func (o *CastEmbedded) GetEmbeds() []EmbedDeep`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *CastEmbedded) GetEmbedsOk() (*[]EmbedDeep, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *CastEmbedded) SetEmbeds(v []EmbedDeep)`

SetEmbeds sets Embeds field to given value.


### GetChannel

`func (o *CastEmbedded) GetChannel() ChannelDehydrated`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CastEmbedded) GetChannelOk() (*ChannelDehydrated, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CastEmbedded) SetChannel(v ChannelDehydrated)`

SetChannel sets Channel field to given value.


### SetChannelNil

`func (o *CastEmbedded) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *CastEmbedded) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


