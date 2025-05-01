# OembedLinkData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Version** | **string** |  | 
**Title** | Pointer to **string** | A text title, describing the resource. | [optional] 
**AuthorName** | Pointer to **string** | The name of the author/owner of the resource. | [optional] 
**AuthorUrl** | Pointer to **string** | A URL for the author/owner of the resource. | [optional] 
**ProviderName** | Pointer to **string** | The name of the resource provider. | [optional] 
**ProviderUrl** | Pointer to **string** | The url of the resource provider. | [optional] 
**CacheAge** | Pointer to **string** | The suggested cache lifetime for this resource, in seconds. Consumers may choose to use this value or not. | [optional] 
**ThumbnailUrl** | Pointer to **string** | A URL to a thumbnail image representing the resource. The thumbnail must respect any maxwidth and maxheight parameters. If this parameter is present, thumbnail_width and thumbnail_height must also be present. | [optional] 
**ThumbnailWidth** | Pointer to **float32** | The width of the optional thumbnail. If this parameter is present, thumbnail_url and thumbnail_height must also be present. | [optional] 
**ThumbnailHeight** | Pointer to **float32** | The height of the optional thumbnail. If this parameter is present, thumbnail_url and thumbnail_width must also be present. | [optional] 

## Methods

### NewOembedLinkData

`func NewOembedLinkData(type_ string, version string, ) *OembedLinkData`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


