# EmbedUrlMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**ContentLength** | Pointer to **NullableInt32** |  | [optional] 
**Image** | Pointer to [**EmbedUrlMetadataImage**](EmbedUrlMetadataImage.md) |  | [optional] 
**Video** | Pointer to [**EmbedUrlMetadataVideo**](EmbedUrlMetadataVideo.md) |  | [optional] 
**Html** | Pointer to [**HtmlMetadata**](HtmlMetadata.md) |  | [optional] 
**Frame** | Pointer to [**Frame**](Frame.md) |  | [optional] 

## Methods

### NewEmbedUrlMetadata

`func NewEmbedUrlMetadata(status string, ) *EmbedUrlMetadata`

NewEmbedUrlMetadata instantiates a new EmbedUrlMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbedUrlMetadataWithDefaults

`func NewEmbedUrlMetadataWithDefaults() *EmbedUrlMetadata`

NewEmbedUrlMetadataWithDefaults instantiates a new EmbedUrlMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EmbedUrlMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmbedUrlMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmbedUrlMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetContentType

`func (o *EmbedUrlMetadata) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *EmbedUrlMetadata) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *EmbedUrlMetadata) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *EmbedUrlMetadata) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *EmbedUrlMetadata) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *EmbedUrlMetadata) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetContentLength

`func (o *EmbedUrlMetadata) GetContentLength() int32`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *EmbedUrlMetadata) GetContentLengthOk() (*int32, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *EmbedUrlMetadata) SetContentLength(v int32)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *EmbedUrlMetadata) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### SetContentLengthNil

`func (o *EmbedUrlMetadata) SetContentLengthNil(b bool)`

 SetContentLengthNil sets the value for ContentLength to be an explicit nil

### UnsetContentLength
`func (o *EmbedUrlMetadata) UnsetContentLength()`

UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
### GetImage

`func (o *EmbedUrlMetadata) GetImage() EmbedUrlMetadataImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *EmbedUrlMetadata) GetImageOk() (*EmbedUrlMetadataImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *EmbedUrlMetadata) SetImage(v EmbedUrlMetadataImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *EmbedUrlMetadata) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVideo

`func (o *EmbedUrlMetadata) GetVideo() EmbedUrlMetadataVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *EmbedUrlMetadata) GetVideoOk() (*EmbedUrlMetadataVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *EmbedUrlMetadata) SetVideo(v EmbedUrlMetadataVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *EmbedUrlMetadata) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetHtml

`func (o *EmbedUrlMetadata) GetHtml() HtmlMetadata`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *EmbedUrlMetadata) GetHtmlOk() (*HtmlMetadata, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *EmbedUrlMetadata) SetHtml(v HtmlMetadata)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *EmbedUrlMetadata) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetFrame

`func (o *EmbedUrlMetadata) GetFrame() Frame`

GetFrame returns the Frame field if non-nil, zero value otherwise.

### GetFrameOk

`func (o *EmbedUrlMetadata) GetFrameOk() (*Frame, bool)`

GetFrameOk returns a tuple with the Frame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrame

`func (o *EmbedUrlMetadata) SetFrame(v Frame)`

SetFrame sets Frame field to given value.

### HasFrame

`func (o *EmbedUrlMetadata) HasFrame() bool`

HasFrame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


