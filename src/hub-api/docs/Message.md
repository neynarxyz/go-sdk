# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**HashScheme** | [**HashScheme**](HashScheme.md) |  | [default to HASH_SCHEME_BLAKE3]
**Signature** | **string** |  | 
**SignatureScheme** | [**SignatureScheme**](SignatureScheme.md) |  | [default to ED25519]
**Signer** | **string** |  | 
**Data** | [**MessageAllOfData**](MessageAllOfData.md) |  | 

## Methods

### NewMessage

`func NewMessage(hash string, hashScheme HashScheme, signature string, signatureScheme SignatureScheme, signer string, data MessageAllOfData, ) *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *Message) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Message) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Message) SetHash(v string)`

SetHash sets Hash field to given value.


### GetHashScheme

`func (o *Message) GetHashScheme() HashScheme`

GetHashScheme returns the HashScheme field if non-nil, zero value otherwise.

### GetHashSchemeOk

`func (o *Message) GetHashSchemeOk() (*HashScheme, bool)`

GetHashSchemeOk returns a tuple with the HashScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashScheme

`func (o *Message) SetHashScheme(v HashScheme)`

SetHashScheme sets HashScheme field to given value.


### GetSignature

`func (o *Message) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Message) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Message) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetSignatureScheme

`func (o *Message) GetSignatureScheme() SignatureScheme`

GetSignatureScheme returns the SignatureScheme field if non-nil, zero value otherwise.

### GetSignatureSchemeOk

`func (o *Message) GetSignatureSchemeOk() (*SignatureScheme, bool)`

GetSignatureSchemeOk returns a tuple with the SignatureScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureScheme

`func (o *Message) SetSignatureScheme(v SignatureScheme)`

SetSignatureScheme sets SignatureScheme field to given value.


### GetSigner

`func (o *Message) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *Message) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *Message) SetSigner(v string)`

SetSigner sets Signer field to given value.


### GetData

`func (o *Message) GetData() MessageAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Message) GetDataOk() (*MessageAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Message) SetData(v MessageAllOfData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


