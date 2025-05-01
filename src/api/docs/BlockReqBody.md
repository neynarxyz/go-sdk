# BlockReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer. &#x60;signer_uuid&#x60; is paired with API key, can&#39;t use a &#x60;uuid&#x60; made with a different API key.  | 
**BlockedFid** | **int32** | The unique identifier of a farcaster user (unsigned integer) | 

## Methods

### NewBlockReqBody

`func NewBlockReqBody(signerUuid string, blockedFid int32, ) *BlockReqBody`

NewBlockReqBody instantiates a new BlockReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockReqBodyWithDefaults

`func NewBlockReqBodyWithDefaults() *BlockReqBody`

NewBlockReqBodyWithDefaults instantiates a new BlockReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *BlockReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *BlockReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *BlockReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetBlockedFid

`func (o *BlockReqBody) GetBlockedFid() int32`

GetBlockedFid returns the BlockedFid field if non-nil, zero value otherwise.

### GetBlockedFidOk

`func (o *BlockReqBody) GetBlockedFidOk() (*int32, bool)`

GetBlockedFidOk returns a tuple with the BlockedFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedFid

`func (o *BlockReqBody) SetBlockedFid(v int32)`

SetBlockedFid sets BlockedFid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


