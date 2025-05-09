# EmbedCast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CastId** | Pointer to [**CastId**](CastId.md) | [DEPRECATED: Use \&quot;cast\&quot; key instead] | [optional] 
**Cast** | [**CastEmbedded**](CastEmbedded.md) |  | 

## Methods

### NewEmbedCast

`func NewEmbedCast(cast CastEmbedded, ) *EmbedCast`

NewEmbedCast instantiates a new EmbedCast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbedCastWithDefaults

`func NewEmbedCastWithDefaults() *EmbedCast`

NewEmbedCastWithDefaults instantiates a new EmbedCast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCastId

`func (o *EmbedCast) GetCastId() CastId`

GetCastId returns the CastId field if non-nil, zero value otherwise.

### GetCastIdOk

`func (o *EmbedCast) GetCastIdOk() (*CastId, bool)`

GetCastIdOk returns a tuple with the CastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastId

`func (o *EmbedCast) SetCastId(v CastId)`

SetCastId sets CastId field to given value.

### HasCastId

`func (o *EmbedCast) HasCastId() bool`

HasCastId returns a boolean if a field has been set.

### GetCast

`func (o *EmbedCast) GetCast() CastEmbedded`

GetCast returns the Cast field if non-nil, zero value otherwise.

### GetCastOk

`func (o *EmbedCast) GetCastOk() (*CastEmbedded, bool)`

GetCastOk returns a tuple with the Cast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCast

`func (o *EmbedCast) SetCast(v CastEmbedded)`

SetCast sets Cast field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


