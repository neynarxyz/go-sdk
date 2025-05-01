# HtmlMetadataAllOfOembed

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
**Html** | **string** | The HTML required to embed a video player. The HTML should have no padding or margins. Consumers may wish to load the HTML in an off-domain iframe to avoid XSS vulnerabilities. | 
**Width** | **float32** | The width in pixels of the image specified in the url parameter. | 
**Height** | **float32** | The height in pixels of the image specified in the url parameter. | 
**Url** | **string** | The source URL of the image. Consumers should be able to insert this URL into an &lt;img&gt; element. Only HTTP and HTTPS URLs are valid. | 

## Methods

### NewHtmlMetadataAllOfOembed

`func NewHtmlMetadataAllOfOembed(type_ string, version string, html string, width float32, height float32, url string, ) *HtmlMetadataAllOfOembed`

NewHtmlMetadataAllOfOembed instantiates a new HtmlMetadataAllOfOembed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHtmlMetadataAllOfOembedWithDefaults

`func NewHtmlMetadataAllOfOembedWithDefaults() *HtmlMetadataAllOfOembed`

NewHtmlMetadataAllOfOembedWithDefaults instantiates a new HtmlMetadataAllOfOembed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HtmlMetadataAllOfOembed) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HtmlMetadataAllOfOembed) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HtmlMetadataAllOfOembed) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *HtmlMetadataAllOfOembed) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HtmlMetadataAllOfOembed) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HtmlMetadataAllOfOembed) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetTitle

`func (o *HtmlMetadataAllOfOembed) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HtmlMetadataAllOfOembed) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HtmlMetadataAllOfOembed) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HtmlMetadataAllOfOembed) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAuthorName

`func (o *HtmlMetadataAllOfOembed) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *HtmlMetadataAllOfOembed) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *HtmlMetadataAllOfOembed) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *HtmlMetadataAllOfOembed) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetAuthorUrl

`func (o *HtmlMetadataAllOfOembed) GetAuthorUrl() string`

GetAuthorUrl returns the AuthorUrl field if non-nil, zero value otherwise.

### GetAuthorUrlOk

`func (o *HtmlMetadataAllOfOembed) GetAuthorUrlOk() (*string, bool)`

GetAuthorUrlOk returns a tuple with the AuthorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUrl

`func (o *HtmlMetadataAllOfOembed) SetAuthorUrl(v string)`

SetAuthorUrl sets AuthorUrl field to given value.

### HasAuthorUrl

`func (o *HtmlMetadataAllOfOembed) HasAuthorUrl() bool`

HasAuthorUrl returns a boolean if a field has been set.

### GetProviderName

`func (o *HtmlMetadataAllOfOembed) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *HtmlMetadataAllOfOembed) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *HtmlMetadataAllOfOembed) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *HtmlMetadataAllOfOembed) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetProviderUrl

`func (o *HtmlMetadataAllOfOembed) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *HtmlMetadataAllOfOembed) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *HtmlMetadataAllOfOembed) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *HtmlMetadataAllOfOembed) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### GetCacheAge

`func (o *HtmlMetadataAllOfOembed) GetCacheAge() string`

GetCacheAge returns the CacheAge field if non-nil, zero value otherwise.

### GetCacheAgeOk

`func (o *HtmlMetadataAllOfOembed) GetCacheAgeOk() (*string, bool)`

GetCacheAgeOk returns a tuple with the CacheAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAge

`func (o *HtmlMetadataAllOfOembed) SetCacheAge(v string)`

SetCacheAge sets CacheAge field to given value.

### HasCacheAge

`func (o *HtmlMetadataAllOfOembed) HasCacheAge() bool`

HasCacheAge returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *HtmlMetadataAllOfOembed) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *HtmlMetadataAllOfOembed) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *HtmlMetadataAllOfOembed) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *HtmlMetadataAllOfOembed) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetThumbnailWidth

`func (o *HtmlMetadataAllOfOembed) GetThumbnailWidth() float32`

GetThumbnailWidth returns the ThumbnailWidth field if non-nil, zero value otherwise.

### GetThumbnailWidthOk

`func (o *HtmlMetadataAllOfOembed) GetThumbnailWidthOk() (*float32, bool)`

GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailWidth

`func (o *HtmlMetadataAllOfOembed) SetThumbnailWidth(v float32)`

SetThumbnailWidth sets ThumbnailWidth field to given value.

### HasThumbnailWidth

`func (o *HtmlMetadataAllOfOembed) HasThumbnailWidth() bool`

HasThumbnailWidth returns a boolean if a field has been set.

### GetThumbnailHeight

`func (o *HtmlMetadataAllOfOembed) GetThumbnailHeight() float32`

GetThumbnailHeight returns the ThumbnailHeight field if non-nil, zero value otherwise.

### GetThumbnailHeightOk

`func (o *HtmlMetadataAllOfOembed) GetThumbnailHeightOk() (*float32, bool)`

GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailHeight

`func (o *HtmlMetadataAllOfOembed) SetThumbnailHeight(v float32)`

SetThumbnailHeight sets ThumbnailHeight field to given value.

### HasThumbnailHeight

`func (o *HtmlMetadataAllOfOembed) HasThumbnailHeight() bool`

HasThumbnailHeight returns a boolean if a field has been set.

### GetHtml

`func (o *HtmlMetadataAllOfOembed) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *HtmlMetadataAllOfOembed) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *HtmlMetadataAllOfOembed) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetWidth

`func (o *HtmlMetadataAllOfOembed) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *HtmlMetadataAllOfOembed) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *HtmlMetadataAllOfOembed) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *HtmlMetadataAllOfOembed) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *HtmlMetadataAllOfOembed) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *HtmlMetadataAllOfOembed) SetHeight(v float32)`

SetHeight sets Height field to given value.


### GetUrl

`func (o *HtmlMetadataAllOfOembed) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HtmlMetadataAllOfOembed) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HtmlMetadataAllOfOembed) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


