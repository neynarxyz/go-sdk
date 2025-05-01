# CastId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | The Farcaster ID (FID) of the user who created the cast. Required to uniquely identify the cast&#39;s author in the network. | 
**Hash** | **string** | A unique hash that identifies a specific cast within the creator&#39;s posts. Generated using HASH_SCHEME_BLAKE3 of the cast&#39;s content. | 

## Methods

### NewCastId

`func NewCastId(fid int32, hash string, ) *CastId`

NewCastId instantiates a new CastId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastIdWithDefaults

`func NewCastIdWithDefaults() *CastId`

NewCastIdWithDefaults instantiates a new CastId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *CastId) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *CastId) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *CastId) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetHash

`func (o *CastId) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CastId) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CastId) SetHash(v string)`

SetHash sets Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


