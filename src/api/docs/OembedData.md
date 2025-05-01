# OembedData

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

### NewOembedData

`func NewOembedData(type_ string, version NullableString, ) *OembedData`

NewOembedData instantiates a new OembedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOembedDataWithDefaults

`func NewOembedDataWithDefaults() *OembedData`

NewOembedDataWithDefaults instantiates a new OembedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OembedData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OembedData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OembedData) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *OembedData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OembedData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OembedData) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *OembedData) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *OembedData) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetTitle

`func (o *OembedData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OembedData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OembedData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OembedData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *OembedData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *OembedData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthorName

`func (o *OembedData) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *OembedData) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *OembedData) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *OembedData) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### SetAuthorNameNil

`func (o *OembedData) SetAuthorNameNil(b bool)`

 SetAuthorNameNil sets the value for AuthorName to be an explicit nil

### UnsetAuthorName
`func (o *OembedData) UnsetAuthorName()`

UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
### GetAuthorUrl

`func (o *OembedData) GetAuthorUrl() string`

GetAuthorUrl returns the AuthorUrl field if non-nil, zero value otherwise.

### GetAuthorUrlOk

`func (o *OembedData) GetAuthorUrlOk() (*string, bool)`

GetAuthorUrlOk returns a tuple with the AuthorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUrl

`func (o *OembedData) SetAuthorUrl(v string)`

SetAuthorUrl sets AuthorUrl field to given value.

### HasAuthorUrl

`func (o *OembedData) HasAuthorUrl() bool`

HasAuthorUrl returns a boolean if a field has been set.

### SetAuthorUrlNil

`func (o *OembedData) SetAuthorUrlNil(b bool)`

 SetAuthorUrlNil sets the value for AuthorUrl to be an explicit nil

### UnsetAuthorUrl
`func (o *OembedData) UnsetAuthorUrl()`

UnsetAuthorUrl ensures that no value is present for AuthorUrl, not even an explicit nil
### GetProviderName

`func (o *OembedData) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *OembedData) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *OembedData) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *OembedData) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *OembedData) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *OembedData) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetProviderUrl

`func (o *OembedData) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *OembedData) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *OembedData) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *OembedData) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### SetProviderUrlNil

`func (o *OembedData) SetProviderUrlNil(b bool)`

 SetProviderUrlNil sets the value for ProviderUrl to be an explicit nil

### UnsetProviderUrl
`func (o *OembedData) UnsetProviderUrl()`

UnsetProviderUrl ensures that no value is present for ProviderUrl, not even an explicit nil
### GetCacheAge

`func (o *OembedData) GetCacheAge() string`

GetCacheAge returns the CacheAge field if non-nil, zero value otherwise.

### GetCacheAgeOk

`func (o *OembedData) GetCacheAgeOk() (*string, bool)`

GetCacheAgeOk returns a tuple with the CacheAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAge

`func (o *OembedData) SetCacheAge(v string)`

SetCacheAge sets CacheAge field to given value.

### HasCacheAge

`func (o *OembedData) HasCacheAge() bool`

HasCacheAge returns a boolean if a field has been set.

### SetCacheAgeNil

`func (o *OembedData) SetCacheAgeNil(b bool)`

 SetCacheAgeNil sets the value for CacheAge to be an explicit nil

### UnsetCacheAge
`func (o *OembedData) UnsetCacheAge()`

UnsetCacheAge ensures that no value is present for CacheAge, not even an explicit nil
### GetThumbnailUrl

`func (o *OembedData) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *OembedData) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *OembedData) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *OembedData) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *OembedData) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *OembedData) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetThumbnailWidth

`func (o *OembedData) GetThumbnailWidth() float32`

GetThumbnailWidth returns the ThumbnailWidth field if non-nil, zero value otherwise.

### GetThumbnailWidthOk

`func (o *OembedData) GetThumbnailWidthOk() (*float32, bool)`

GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailWidth

`func (o *OembedData) SetThumbnailWidth(v float32)`

SetThumbnailWidth sets ThumbnailWidth field to given value.

### HasThumbnailWidth

`func (o *OembedData) HasThumbnailWidth() bool`

HasThumbnailWidth returns a boolean if a field has been set.

### SetThumbnailWidthNil

`func (o *OembedData) SetThumbnailWidthNil(b bool)`

 SetThumbnailWidthNil sets the value for ThumbnailWidth to be an explicit nil

### UnsetThumbnailWidth
`func (o *OembedData) UnsetThumbnailWidth()`

UnsetThumbnailWidth ensures that no value is present for ThumbnailWidth, not even an explicit nil
### GetThumbnailHeight

`func (o *OembedData) GetThumbnailHeight() float32`

GetThumbnailHeight returns the ThumbnailHeight field if non-nil, zero value otherwise.

### GetThumbnailHeightOk

`func (o *OembedData) GetThumbnailHeightOk() (*float32, bool)`

GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailHeight

`func (o *OembedData) SetThumbnailHeight(v float32)`

SetThumbnailHeight sets ThumbnailHeight field to given value.

### HasThumbnailHeight

`func (o *OembedData) HasThumbnailHeight() bool`

HasThumbnailHeight returns a boolean if a field has been set.

### SetThumbnailHeightNil

`func (o *OembedData) SetThumbnailHeightNil(b bool)`

 SetThumbnailHeightNil sets the value for ThumbnailHeight to be an explicit nil

### UnsetThumbnailHeight
`func (o *OembedData) UnsetThumbnailHeight()`

UnsetThumbnailHeight ensures that no value is present for ThumbnailHeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


