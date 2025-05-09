# RegisterUserOnChainReqBodySignersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | **string** | Hex encoded signature | 
**Metadata** | **string** | Hex encoded signature | 
**PublicKey** | **string** | Ed25519 public key | 
**KeyType** | Pointer to **int32** |  | [optional] [default to 1]
**MetadataType** | Pointer to **int32** |  | [optional] [default to 1]
**Deadline** | **int32** |  | 

## Methods

### NewRegisterUserOnChainReqBodySignersInner

`func NewRegisterUserOnChainReqBodySignersInner(signature string, metadata string, publicKey string, deadline int32, ) *RegisterUserOnChainReqBodySignersInner`

NewRegisterUserOnChainReqBodySignersInner instantiates a new RegisterUserOnChainReqBodySignersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserOnChainReqBodySignersInnerWithDefaults

`func NewRegisterUserOnChainReqBodySignersInnerWithDefaults() *RegisterUserOnChainReqBodySignersInner`

NewRegisterUserOnChainReqBodySignersInnerWithDefaults instantiates a new RegisterUserOnChainReqBodySignersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *RegisterUserOnChainReqBodySignersInner) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *RegisterUserOnChainReqBodySignersInner) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *RegisterUserOnChainReqBodySignersInner) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetMetadata

`func (o *RegisterUserOnChainReqBodySignersInner) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RegisterUserOnChainReqBodySignersInner) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RegisterUserOnChainReqBodySignersInner) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetPublicKey

`func (o *RegisterUserOnChainReqBodySignersInner) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RegisterUserOnChainReqBodySignersInner) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RegisterUserOnChainReqBodySignersInner) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetKeyType

`func (o *RegisterUserOnChainReqBodySignersInner) GetKeyType() int32`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *RegisterUserOnChainReqBodySignersInner) GetKeyTypeOk() (*int32, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *RegisterUserOnChainReqBodySignersInner) SetKeyType(v int32)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *RegisterUserOnChainReqBodySignersInner) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetMetadataType

`func (o *RegisterUserOnChainReqBodySignersInner) GetMetadataType() int32`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *RegisterUserOnChainReqBodySignersInner) GetMetadataTypeOk() (*int32, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *RegisterUserOnChainReqBodySignersInner) SetMetadataType(v int32)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *RegisterUserOnChainReqBodySignersInner) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.

### GetDeadline

`func (o *RegisterUserOnChainReqBodySignersInner) GetDeadline() int32`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *RegisterUserOnChainReqBodySignersInner) GetDeadlineOk() (*int32, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *RegisterUserOnChainReqBodySignersInner) SetDeadline(v int32)`

SetDeadline sets Deadline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


