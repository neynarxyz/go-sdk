# MessageCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**HashScheme** | [**HashScheme**](HashScheme.md) |  | [default to HASHSCHEME_HASH_SCHEME_BLAKE3]
**Signature** | **string** |  | 
**SignatureScheme** | [**SignatureScheme**](SignatureScheme.md) |  | [default to SIGNATURESCHEME_SIGNATURE_SCHEME_ED25519]
**Signer** | **string** |  | 

## Methods

### NewMessageCommon

`func NewMessageCommon(hash string, hashScheme HashScheme, signature string, signatureScheme SignatureScheme, signer string, ) *MessageCommon`

NewMessageCommon instantiates a new MessageCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageCommonWithDefaults

`func NewMessageCommonWithDefaults() *MessageCommon`

NewMessageCommonWithDefaults instantiates a new MessageCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *MessageCommon) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *MessageCommon) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *MessageCommon) SetHash(v string)`

SetHash sets Hash field to given value.


### GetHashScheme

`func (o *MessageCommon) GetHashScheme() HashScheme`

GetHashScheme returns the HashScheme field if non-nil, zero value otherwise.

### GetHashSchemeOk

`func (o *MessageCommon) GetHashSchemeOk() (*HashScheme, bool)`

GetHashSchemeOk returns a tuple with the HashScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashScheme

`func (o *MessageCommon) SetHashScheme(v HashScheme)`

SetHashScheme sets HashScheme field to given value.


### GetSignature

`func (o *MessageCommon) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *MessageCommon) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *MessageCommon) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetSignatureScheme

`func (o *MessageCommon) GetSignatureScheme() SignatureScheme`

GetSignatureScheme returns the SignatureScheme field if non-nil, zero value otherwise.

### GetSignatureSchemeOk

`func (o *MessageCommon) GetSignatureSchemeOk() (*SignatureScheme, bool)`

GetSignatureSchemeOk returns a tuple with the SignatureScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureScheme

`func (o *MessageCommon) SetSignatureScheme(v SignatureScheme)`

SetSignatureScheme sets SignatureScheme field to given value.


### GetSigner

`func (o *MessageCommon) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *MessageCommon) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *MessageCommon) SetSigner(v string)`

SetSigner sets Signer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


