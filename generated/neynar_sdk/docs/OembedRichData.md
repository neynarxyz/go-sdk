# OembedRichData

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
**Html** | **NullableString** | The HTML required to display the resource. The HTML should have no padding or margins. Consumers may wish to load the HTML in an off-domain iframe to avoid XSS vulnerabilities. The markup should be valid XHTML 1.0 Basic. | 
**Width** | **NullableFloat32** | The width in pixels required to display the HTML. | 
**Height** | **NullableFloat32** | The height in pixels required to display the HTML. | 

## Methods

### NewOembedRichData

`func NewOembedRichData(type_ string, version NullableString, html NullableString, width NullableFloat32, height NullableFloat32, ) *OembedRichData`

NewOembedRichData instantiates a new OembedRichData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOembedRichDataWithDefaults

`func NewOembedRichDataWithDefaults() *OembedRichData`

NewOembedRichDataWithDefaults instantiates a new OembedRichData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OembedRichData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OembedRichData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OembedRichData) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *OembedRichData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OembedRichData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OembedRichData) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *OembedRichData) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *OembedRichData) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetTitle

`func (o *OembedRichData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OembedRichData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OembedRichData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OembedRichData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *OembedRichData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *OembedRichData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthorName

`func (o *OembedRichData) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *OembedRichData) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *OembedRichData) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *OembedRichData) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### SetAuthorNameNil

`func (o *OembedRichData) SetAuthorNameNil(b bool)`

 SetAuthorNameNil sets the value for AuthorName to be an explicit nil

### UnsetAuthorName
`func (o *OembedRichData) UnsetAuthorName()`

UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
### GetAuthorUrl

`func (o *OembedRichData) GetAuthorUrl() string`

GetAuthorUrl returns the AuthorUrl field if non-nil, zero value otherwise.

### GetAuthorUrlOk

`func (o *OembedRichData) GetAuthorUrlOk() (*string, bool)`

GetAuthorUrlOk returns a tuple with the AuthorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUrl

`func (o *OembedRichData) SetAuthorUrl(v string)`

SetAuthorUrl sets AuthorUrl field to given value.

### HasAuthorUrl

`func (o *OembedRichData) HasAuthorUrl() bool`

HasAuthorUrl returns a boolean if a field has been set.

### SetAuthorUrlNil

`func (o *OembedRichData) SetAuthorUrlNil(b bool)`

 SetAuthorUrlNil sets the value for AuthorUrl to be an explicit nil

### UnsetAuthorUrl
`func (o *OembedRichData) UnsetAuthorUrl()`

UnsetAuthorUrl ensures that no value is present for AuthorUrl, not even an explicit nil
### GetProviderName

`func (o *OembedRichData) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *OembedRichData) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *OembedRichData) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *OembedRichData) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *OembedRichData) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *OembedRichData) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetProviderUrl

`func (o *OembedRichData) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *OembedRichData) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *OembedRichData) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *OembedRichData) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### SetProviderUrlNil

`func (o *OembedRichData) SetProviderUrlNil(b bool)`

 SetProviderUrlNil sets the value for ProviderUrl to be an explicit nil

### UnsetProviderUrl
`func (o *OembedRichData) UnsetProviderUrl()`

UnsetProviderUrl ensures that no value is present for ProviderUrl, not even an explicit nil
### GetCacheAge

`func (o *OembedRichData) GetCacheAge() string`

GetCacheAge returns the CacheAge field if non-nil, zero value otherwise.

### GetCacheAgeOk

`func (o *OembedRichData) GetCacheAgeOk() (*string, bool)`

GetCacheAgeOk returns a tuple with the CacheAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAge

`func (o *OembedRichData) SetCacheAge(v string)`

SetCacheAge sets CacheAge field to given value.

### HasCacheAge

`func (o *OembedRichData) HasCacheAge() bool`

HasCacheAge returns a boolean if a field has been set.

### SetCacheAgeNil

`func (o *OembedRichData) SetCacheAgeNil(b bool)`

 SetCacheAgeNil sets the value for CacheAge to be an explicit nil

### UnsetCacheAge
`func (o *OembedRichData) UnsetCacheAge()`

UnsetCacheAge ensures that no value is present for CacheAge, not even an explicit nil
### GetThumbnailUrl

`func (o *OembedRichData) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *OembedRichData) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *OembedRichData) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *OembedRichData) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *OembedRichData) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *OembedRichData) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetThumbnailWidth

`func (o *OembedRichData) GetThumbnailWidth() float32`

GetThumbnailWidth returns the ThumbnailWidth field if non-nil, zero value otherwise.

### GetThumbnailWidthOk

`func (o *OembedRichData) GetThumbnailWidthOk() (*float32, bool)`

GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailWidth

`func (o *OembedRichData) SetThumbnailWidth(v float32)`

SetThumbnailWidth sets ThumbnailWidth field to given value.

### HasThumbnailWidth

`func (o *OembedRichData) HasThumbnailWidth() bool`

HasThumbnailWidth returns a boolean if a field has been set.

### SetThumbnailWidthNil

`func (o *OembedRichData) SetThumbnailWidthNil(b bool)`

 SetThumbnailWidthNil sets the value for ThumbnailWidth to be an explicit nil

### UnsetThumbnailWidth
`func (o *OembedRichData) UnsetThumbnailWidth()`

UnsetThumbnailWidth ensures that no value is present for ThumbnailWidth, not even an explicit nil
### GetThumbnailHeight

`func (o *OembedRichData) GetThumbnailHeight() float32`

GetThumbnailHeight returns the ThumbnailHeight field if non-nil, zero value otherwise.

### GetThumbnailHeightOk

`func (o *OembedRichData) GetThumbnailHeightOk() (*float32, bool)`

GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailHeight

`func (o *OembedRichData) SetThumbnailHeight(v float32)`

SetThumbnailHeight sets ThumbnailHeight field to given value.

### HasThumbnailHeight

`func (o *OembedRichData) HasThumbnailHeight() bool`

HasThumbnailHeight returns a boolean if a field has been set.

### SetThumbnailHeightNil

`func (o *OembedRichData) SetThumbnailHeightNil(b bool)`

 SetThumbnailHeightNil sets the value for ThumbnailHeight to be an explicit nil

### UnsetThumbnailHeight
`func (o *OembedRichData) UnsetThumbnailHeight()`

UnsetThumbnailHeight ensures that no value is present for ThumbnailHeight, not even an explicit nil
### GetHtml

`func (o *OembedRichData) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *OembedRichData) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *OembedRichData) SetHtml(v string)`

SetHtml sets Html field to given value.


### SetHtmlNil

`func (o *OembedRichData) SetHtmlNil(b bool)`

 SetHtmlNil sets the value for Html to be an explicit nil

### UnsetHtml
`func (o *OembedRichData) UnsetHtml()`

UnsetHtml ensures that no value is present for Html, not even an explicit nil
### GetWidth

`func (o *OembedRichData) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *OembedRichData) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *OembedRichData) SetWidth(v float32)`

SetWidth sets Width field to given value.


### SetWidthNil

`func (o *OembedRichData) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *OembedRichData) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *OembedRichData) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *OembedRichData) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *OembedRichData) SetHeight(v float32)`

SetHeight sets Height field to given value.


### SetHeightNil

`func (o *OembedRichData) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *OembedRichData) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


