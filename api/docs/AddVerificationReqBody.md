# AddVerificationReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer. &#x60;signer_uuid&#x60; is paired with API key, can&#39;t use a &#x60;uuid&#x60; made with a different API key.  | 
**Address** | **string** | Ethereum address | 
**BlockHash** | **string** |  | 
**EthSignature** | **string** |  | 
**VerificationType** | Pointer to [**VerificationType**](VerificationType.md) |  | [optional] [default to VERIFICATIONTYPE__0]
**ChainId** | Pointer to [**VerificationChainId**](VerificationChainId.md) |  | [optional] [default to VERIFICATIONCHAINID__0]

## Methods

### NewAddVerificationReqBody

`func NewAddVerificationReqBody(signerUuid string, address string, blockHash string, ethSignature string, ) *AddVerificationReqBody`

NewAddVerificationReqBody instantiates a new AddVerificationReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVerificationReqBodyWithDefaults

`func NewAddVerificationReqBodyWithDefaults() *AddVerificationReqBody`

NewAddVerificationReqBodyWithDefaults instantiates a new AddVerificationReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *AddVerificationReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *AddVerificationReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *AddVerificationReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetAddress

`func (o *AddVerificationReqBody) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddVerificationReqBody) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddVerificationReqBody) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBlockHash

`func (o *AddVerificationReqBody) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *AddVerificationReqBody) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *AddVerificationReqBody) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetEthSignature

`func (o *AddVerificationReqBody) GetEthSignature() string`

GetEthSignature returns the EthSignature field if non-nil, zero value otherwise.

### GetEthSignatureOk

`func (o *AddVerificationReqBody) GetEthSignatureOk() (*string, bool)`

GetEthSignatureOk returns a tuple with the EthSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthSignature

`func (o *AddVerificationReqBody) SetEthSignature(v string)`

SetEthSignature sets EthSignature field to given value.


### GetVerificationType

`func (o *AddVerificationReqBody) GetVerificationType() VerificationType`

GetVerificationType returns the VerificationType field if non-nil, zero value otherwise.

### GetVerificationTypeOk

`func (o *AddVerificationReqBody) GetVerificationTypeOk() (*VerificationType, bool)`

GetVerificationTypeOk returns a tuple with the VerificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationType

`func (o *AddVerificationReqBody) SetVerificationType(v VerificationType)`

SetVerificationType sets VerificationType field to given value.

### HasVerificationType

`func (o *AddVerificationReqBody) HasVerificationType() bool`

HasVerificationType returns a boolean if a field has been set.

### GetChainId

`func (o *AddVerificationReqBody) GetChainId() VerificationChainId`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *AddVerificationReqBody) GetChainIdOk() (*VerificationChainId, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *AddVerificationReqBody) SetChainId(v VerificationChainId)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *AddVerificationReqBody) HasChainId() bool`

HasChainId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


