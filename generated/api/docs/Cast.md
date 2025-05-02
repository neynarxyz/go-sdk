# Cast

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

## Methods

### NewCast

`func NewCast(object string, hash string, parentHash NullableString, parentUrl NullableString, rootParentUrl NullableString, parentAuthor CastParentAuthor, author User, text string, timestamp time.Time, embeds []Embed, ) *Cast`

NewCast instantiates a new Cast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastWithDefaults

`func NewCastWithDefaults() *Cast`

NewCastWithDefaults instantiates a new Cast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Cast) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Cast) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Cast) SetObject(v string)`

SetObject sets Object field to given value.


### GetHash

`func (o *Cast) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Cast) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Cast) SetHash(v string)`

SetHash sets Hash field to given value.


### GetParentHash

`func (o *Cast) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *Cast) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *Cast) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### SetParentHashNil

`func (o *Cast) SetParentHashNil(b bool)`

 SetParentHashNil sets the value for ParentHash to be an explicit nil

### UnsetParentHash
`func (o *Cast) UnsetParentHash()`

UnsetParentHash ensures that no value is present for ParentHash, not even an explicit nil
### GetParentUrl

`func (o *Cast) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *Cast) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *Cast) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.


### SetParentUrlNil

`func (o *Cast) SetParentUrlNil(b bool)`

 SetParentUrlNil sets the value for ParentUrl to be an explicit nil

### UnsetParentUrl
`func (o *Cast) UnsetParentUrl()`

UnsetParentUrl ensures that no value is present for ParentUrl, not even an explicit nil
### GetRootParentUrl

`func (o *Cast) GetRootParentUrl() string`

GetRootParentUrl returns the RootParentUrl field if non-nil, zero value otherwise.

### GetRootParentUrlOk

`func (o *Cast) GetRootParentUrlOk() (*string, bool)`

GetRootParentUrlOk returns a tuple with the RootParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootParentUrl

`func (o *Cast) SetRootParentUrl(v string)`

SetRootParentUrl sets RootParentUrl field to given value.


### SetRootParentUrlNil

`func (o *Cast) SetRootParentUrlNil(b bool)`

 SetRootParentUrlNil sets the value for RootParentUrl to be an explicit nil

### UnsetRootParentUrl
`func (o *Cast) UnsetRootParentUrl()`

UnsetRootParentUrl ensures that no value is present for RootParentUrl, not even an explicit nil
### GetParentAuthor

`func (o *Cast) GetParentAuthor() CastParentAuthor`

GetParentAuthor returns the ParentAuthor field if non-nil, zero value otherwise.

### GetParentAuthorOk

`func (o *Cast) GetParentAuthorOk() (*CastParentAuthor, bool)`

GetParentAuthorOk returns a tuple with the ParentAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAuthor

`func (o *Cast) SetParentAuthor(v CastParentAuthor)`

SetParentAuthor sets ParentAuthor field to given value.


### GetAuthor

`func (o *Cast) GetAuthor() User`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Cast) GetAuthorOk() (*User, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Cast) SetAuthor(v User)`

SetAuthor sets Author field to given value.


### GetApp

`func (o *Cast) GetApp() UserDehydrated`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *Cast) GetAppOk() (*UserDehydrated, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *Cast) SetApp(v UserDehydrated)`

SetApp sets App field to given value.

### HasApp

`func (o *Cast) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetText

`func (o *Cast) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Cast) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Cast) SetText(v string)`

SetText sets Text field to given value.


### GetTimestamp

`func (o *Cast) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Cast) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Cast) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEmbeds

`func (o *Cast) GetEmbeds() []Embed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *Cast) GetEmbedsOk() (*[]Embed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *Cast) SetEmbeds(v []Embed)`

SetEmbeds sets Embeds field to given value.


### GetType

`func (o *Cast) GetType() CastNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cast) GetTypeOk() (*CastNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cast) SetType(v CastNotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *Cast) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


