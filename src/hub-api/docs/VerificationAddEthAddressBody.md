# VerificationAddEthAddressBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The Ethereum (0x-prefixed) or Solana address that the user is claiming ownership of. Must match the address that produced the signature. | 
**EthSignature** | **string** | Base64-encoded signature produced by the blockchain address, proving ownership. For Ethereum, this is an ECDSA signature of a specific message format. | 
**BlockHash** | **string** | The hash of the most recent block when the signature was created. Used to verify the approximate time of signature creation. | 

## Methods

### NewVerificationAddEthAddressBody

`func NewVerificationAddEthAddressBody(address string, ethSignature string, blockHash string, ) *VerificationAddEthAddressBody`

NewVerificationAddEthAddressBody instantiates a new VerificationAddEthAddressBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationAddEthAddressBodyWithDefaults

`func NewVerificationAddEthAddressBodyWithDefaults() *VerificationAddEthAddressBody`

NewVerificationAddEthAddressBodyWithDefaults instantiates a new VerificationAddEthAddressBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VerificationAddEthAddressBody) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VerificationAddEthAddressBody) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VerificationAddEthAddressBody) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetEthSignature

`func (o *VerificationAddEthAddressBody) GetEthSignature() string`

GetEthSignature returns the EthSignature field if non-nil, zero value otherwise.

### GetEthSignatureOk

`func (o *VerificationAddEthAddressBody) GetEthSignatureOk() (*string, bool)`

GetEthSignatureOk returns a tuple with the EthSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthSignature

`func (o *VerificationAddEthAddressBody) SetEthSignature(v string)`

SetEthSignature sets EthSignature field to given value.


### GetBlockHash

`func (o *VerificationAddEthAddressBody) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *VerificationAddEthAddressBody) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *VerificationAddEthAddressBody) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


