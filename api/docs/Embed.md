# Embed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Metadata** | Pointer to [**EmbedUrlMetadata**](EmbedUrlMetadata.md) |  | [optional] 
**CastId** | Pointer to [**CastId**](CastId.md) |  | [optional] 
**Cast** | [**CastEmbedded**](CastEmbedded.md) |  | 

## Methods

### NewEmbed

`func NewEmbed(url string, cast CastEmbedded, ) *Embed`

NewEmbed instantiates a new Embed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbedWithDefaults

`func NewEmbedWithDefaults() *Embed`

NewEmbedWithDefaults instantiates a new Embed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Embed) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Embed) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Embed) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMetadata

`func (o *Embed) GetMetadata() EmbedUrlMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Embed) GetMetadataOk() (*EmbedUrlMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Embed) SetMetadata(v EmbedUrlMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Embed) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCastId

`func (o *Embed) GetCastId() CastId`

GetCastId returns the CastId field if non-nil, zero value otherwise.

### GetCastIdOk

`func (o *Embed) GetCastIdOk() (*CastId, bool)`

GetCastIdOk returns a tuple with the CastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastId

`func (o *Embed) SetCastId(v CastId)`

SetCastId sets CastId field to given value.

### HasCastId

`func (o *Embed) HasCastId() bool`

HasCastId returns a boolean if a field has been set.

### GetCast

`func (o *Embed) GetCast() CastEmbedded`

GetCast returns the Cast field if non-nil, zero value otherwise.

### GetCastOk

`func (o *Embed) GetCastOk() (*CastEmbedded, bool)`

GetCastOk returns a tuple with the Cast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCast

`func (o *Embed) SetCast(v CastEmbedded)`

SetCast sets Cast field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


