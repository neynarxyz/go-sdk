# OembedLinkData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Version** | **NullableString** |  | 
**Title** | Pointer to **NullableString** | A text title, describing the resource. | [optional] 
**AuthorName** | Pointer to **NullableString** | The name of the author/owner of the resource. | [optional] 
**AuthorUrl** | Pointer to **NullableString** | A URL for the author/owner of the resource. | [optional] 
**ProviderName** | Pointer to **NullableString** | The name of the resource provider. | [optional] 
**ProviderUrl** | Pointer to **NullableString** | The url of the resource provider. | [optional] 
**CacheAge** | Pointer to **NullableString** | The suggested cache lifetime for this resource, in seconds. Consumers may choose to use this value or not. | [optional] 
**ThumbnailUrl** | Pointer to **NullableString** | A URL to a thumbnail image representing the resource. The thumbnail must respect any maxwidth and maxheight parameters. If this parameter is present, thumbnail_width and thumbnail_height must also be present. | [optional] 
**ThumbnailWidth** | Pointer to **NullableFloat32** | The width of the optional thumbnail. If this parameter is present, thumbnail_url and thumbnail_height must also be present. | [optional] 
**ThumbnailHeight** | Pointer to **NullableFloat32** | The height of the optional thumbnail. If this parameter is present, thumbnail_url and thumbnail_width must also be present. | [optional] 

## Methods

### NewOembedLinkData

`func NewOembedLinkData(type_ string, version NullableString, ) *OembedLinkData`

NewOembedLinkData instantiates a new OembedLinkData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOembedLinkDataWithDefaults

`func NewOembedLinkDataWithDefaults() *OembedLinkData`

NewOembedLinkDataWithDefaults instantiates a new OembedLinkData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OembedLinkData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OembedLinkData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OembedLinkData) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *OembedLinkData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OembedLinkData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OembedLinkData) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *OembedLinkData) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *OembedLinkData) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetTitle

`func (o *OembedLinkData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OembedLinkData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OembedLinkData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OembedLinkData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *OembedLinkData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *OembedLinkData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthorName

`func (o *OembedLinkData) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *OembedLinkData) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *OembedLinkData) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *OembedLinkData) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### SetAuthorNameNil

`func (o *OembedLinkData) SetAuthorNameNil(b bool)`

 SetAuthorNameNil sets the value for AuthorName to be an explicit nil

### UnsetAuthorName
`func (o *OembedLinkData) UnsetAuthorName()`

UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
### GetAuthorUrl

`func (o *OembedLinkData) GetAuthorUrl() string`

GetAuthorUrl returns the AuthorUrl field if non-nil, zero value otherwise.

### GetAuthorUrlOk

`func (o *OembedLinkData) GetAuthorUrlOk() (*string, bool)`

GetAuthorUrlOk returns a tuple with the AuthorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUrl

`func (o *OembedLinkData) SetAuthorUrl(v string)`

SetAuthorUrl sets AuthorUrl field to given value.

### HasAuthorUrl

`func (o *OembedLinkData) HasAuthorUrl() bool`

HasAuthorUrl returns a boolean if a field has been set.

### SetAuthorUrlNil

`func (o *OembedLinkData) SetAuthorUrlNil(b bool)`

 SetAuthorUrlNil sets the value for AuthorUrl to be an explicit nil

### UnsetAuthorUrl
`func (o *OembedLinkData) UnsetAuthorUrl()`

UnsetAuthorUrl ensures that no value is present for AuthorUrl, not even an explicit nil
### GetProviderName

`func (o *OembedLinkData) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *OembedLinkData) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *OembedLinkData) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *OembedLinkData) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *OembedLinkData) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *OembedLinkData) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetProviderUrl

`func (o *OembedLinkData) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *OembedLinkData) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *OembedLinkData) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *OembedLinkData) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### SetProviderUrlNil

`func (o *OembedLinkData) SetProviderUrlNil(b bool)`

 SetProviderUrlNil sets the value for ProviderUrl to be an explicit nil

### UnsetProviderUrl
`func (o *OembedLinkData) UnsetProviderUrl()`

UnsetProviderUrl ensures that no value is present for ProviderUrl, not even an explicit nil
### GetCacheAge

`func (o *OembedLinkData) GetCacheAge() string`

GetCacheAge returns the CacheAge field if non-nil, zero value otherwise.

### GetCacheAgeOk

`func (o *OembedLinkData) GetCacheAgeOk() (*string, bool)`

GetCacheAgeOk returns a tuple with the CacheAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAge

`func (o *OembedLinkData) SetCacheAge(v string)`

SetCacheAge sets CacheAge field to given value.

### HasCacheAge

`func (o *OembedLinkData) HasCacheAge() bool`

HasCacheAge returns a boolean if a field has been set.

### SetCacheAgeNil

`func (o *OembedLinkData) SetCacheAgeNil(b bool)`

 SetCacheAgeNil sets the value for CacheAge to be an explicit nil

### UnsetCacheAge
`func (o *OembedLinkData) UnsetCacheAge()`

UnsetCacheAge ensures that no value is present for CacheAge, not even an explicit nil
### GetThumbnailUrl

`func (o *OembedLinkData) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *OembedLinkData) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *OembedLinkData) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *OembedLinkData) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *OembedLinkData) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *OembedLinkData) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetThumbnailWidth

`func (o *OembedLinkData) GetThumbnailWidth() float32`

GetThumbnailWidth returns the ThumbnailWidth field if non-nil, zero value otherwise.

### GetThumbnailWidthOk

`func (o *OembedLinkData) GetThumbnailWidthOk() (*float32, bool)`

GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailWidth

`func (o *OembedLinkData) SetThumbnailWidth(v float32)`

SetThumbnailWidth sets ThumbnailWidth field to given value.

### HasThumbnailWidth

`func (o *OembedLinkData) HasThumbnailWidth() bool`

HasThumbnailWidth returns a boolean if a field has been set.

### SetThumbnailWidthNil

`func (o *OembedLinkData) SetThumbnailWidthNil(b bool)`

 SetThumbnailWidthNil sets the value for ThumbnailWidth to be an explicit nil

### UnsetThumbnailWidth
`func (o *OembedLinkData) UnsetThumbnailWidth()`

UnsetThumbnailWidth ensures that no value is present for ThumbnailWidth, not even an explicit nil
### GetThumbnailHeight

`func (o *OembedLinkData) GetThumbnailHeight() float32`

GetThumbnailHeight returns the ThumbnailHeight field if non-nil, zero value otherwise.

### GetThumbnailHeightOk

`func (o *OembedLinkData) GetThumbnailHeightOk() (*float32, bool)`

GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailHeight

`func (o *OembedLinkData) SetThumbnailHeight(v float32)`

SetThumbnailHeight sets ThumbnailHeight field to given value.

### HasThumbnailHeight

`func (o *OembedLinkData) HasThumbnailHeight() bool`

HasThumbnailHeight returns a boolean if a field has been set.

### SetThumbnailHeightNil

`func (o *OembedLinkData) SetThumbnailHeightNil(b bool)`

 SetThumbnailHeightNil sets the value for ThumbnailHeight to be an explicit nil

### UnsetThumbnailHeight
`func (o *OembedLinkData) UnsetThumbnailHeight()`

UnsetThumbnailHeight ensures that no value is present for ThumbnailHeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


