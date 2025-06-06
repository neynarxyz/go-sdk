# EmbedDeep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CastId** | Pointer to [**CastId**](CastId.md) | [DEPRECATED: Use \&quot;cast\&quot; key instead] | [optional] 
**Cast** | [**CastDehydrated**](CastDehydrated.md) |  | 
**Url** | **string** |  | 
**Metadata** | Pointer to [**EmbedUrlMetadata**](EmbedUrlMetadata.md) |  | [optional] 

## Methods

### NewEmbedDeep

`func NewEmbedDeep(cast CastDehydrated, url string, ) *EmbedDeep`

NewEmbedDeep instantiates a new EmbedDeep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbedDeepWithDefaults

`func NewEmbedDeepWithDefaults() *EmbedDeep`

NewEmbedDeepWithDefaults instantiates a new EmbedDeep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCastId

`func (o *EmbedDeep) GetCastId() CastId`

GetCastId returns the CastId field if non-nil, zero value otherwise.

### GetCastIdOk

`func (o *EmbedDeep) GetCastIdOk() (*CastId, bool)`

GetCastIdOk returns a tuple with the CastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastId

`func (o *EmbedDeep) SetCastId(v CastId)`

SetCastId sets CastId field to given value.

### HasCastId

`func (o *EmbedDeep) HasCastId() bool`

HasCastId returns a boolean if a field has been set.

### GetCast

`func (o *EmbedDeep) GetCast() CastDehydrated`

GetCast returns the Cast field if non-nil, zero value otherwise.

### GetCastOk

`func (o *EmbedDeep) GetCastOk() (*CastDehydrated, bool)`

GetCastOk returns a tuple with the Cast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCast

`func (o *EmbedDeep) SetCast(v CastDehydrated)`

SetCast sets Cast field to given value.


### GetUrl

`func (o *EmbedDeep) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EmbedDeep) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EmbedDeep) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMetadata

`func (o *EmbedDeep) GetMetadata() EmbedUrlMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EmbedDeep) GetMetadataOk() (*EmbedUrlMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EmbedDeep) SetMetadata(v EmbedUrlMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EmbedDeep) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


