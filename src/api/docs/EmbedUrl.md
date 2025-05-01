# EmbedUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Metadata** | Pointer to [**EmbedUrlMetadata**](EmbedUrlMetadata.md) |  | [optional] 

## Methods

### NewEmbedUrl

`func NewEmbedUrl(url string, ) *EmbedUrl`

NewEmbedUrl instantiates a new EmbedUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbedUrlWithDefaults

`func NewEmbedUrlWithDefaults() *EmbedUrl`

NewEmbedUrlWithDefaults instantiates a new EmbedUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *EmbedUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EmbedUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EmbedUrl) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMetadata

`func (o *EmbedUrl) GetMetadata() EmbedUrlMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EmbedUrl) GetMetadataOk() (*EmbedUrlMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EmbedUrl) SetMetadata(v EmbedUrlMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EmbedUrl) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


