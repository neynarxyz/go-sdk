# UserDataAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**HashScheme** | [**HashScheme**](HashScheme.md) |  | [default to HASHSCHEME_HASH_SCHEME_BLAKE3]
**Signature** | **string** |  | 
**SignatureScheme** | [**SignatureScheme**](SignatureScheme.md) |  | [default to SIGNATURESCHEME_SIGNATURE_SCHEME_ED25519]
**Signer** | **string** |  | 
**Data** | [**UserDataAddAllOfData**](UserDataAddAllOfData.md) |  | 

## Methods

### NewUserDataAdd

`func NewUserDataAdd(hash string, hashScheme HashScheme, signature string, signatureScheme SignatureScheme, signer string, data UserDataAddAllOfData, ) *UserDataAdd`

NewUserDataAdd instantiates a new UserDataAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataAddWithDefaults

`func NewUserDataAddWithDefaults() *UserDataAdd`

NewUserDataAddWithDefaults instantiates a new UserDataAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *UserDataAdd) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *UserDataAdd) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *UserDataAdd) SetHash(v string)`

SetHash sets Hash field to given value.


### GetHashScheme

`func (o *UserDataAdd) GetHashScheme() HashScheme`

GetHashScheme returns the HashScheme field if non-nil, zero value otherwise.

### GetHashSchemeOk

`func (o *UserDataAdd) GetHashSchemeOk() (*HashScheme, bool)`

GetHashSchemeOk returns a tuple with the HashScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashScheme

`func (o *UserDataAdd) SetHashScheme(v HashScheme)`

SetHashScheme sets HashScheme field to given value.


### GetSignature

`func (o *UserDataAdd) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UserDataAdd) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UserDataAdd) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetSignatureScheme

`func (o *UserDataAdd) GetSignatureScheme() SignatureScheme`

GetSignatureScheme returns the SignatureScheme field if non-nil, zero value otherwise.

### GetSignatureSchemeOk

`func (o *UserDataAdd) GetSignatureSchemeOk() (*SignatureScheme, bool)`

GetSignatureSchemeOk returns a tuple with the SignatureScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureScheme

`func (o *UserDataAdd) SetSignatureScheme(v SignatureScheme)`

SetSignatureScheme sets SignatureScheme field to given value.


### GetSigner

`func (o *UserDataAdd) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *UserDataAdd) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *UserDataAdd) SetSigner(v string)`

SetSigner sets Signer field to given value.


### GetData

`func (o *UserDataAdd) GetData() UserDataAddAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserDataAdd) GetDataOk() (*UserDataAddAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserDataAdd) SetData(v UserDataAddAllOfData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


