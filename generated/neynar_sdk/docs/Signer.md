# Signer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**SignerUuid** | **string** | UUID of the signer. &#x60;signer_uuid&#x60; is paired with API key, can&#39;t use a &#x60;uuid&#x60; made with a different API key. | 
**PublicKey** | **string** | Ed25519 public key | 
**Status** | **string** |  | 
**SignerApprovalUrl** | Pointer to **string** |  | [optional] 
**Fid** | Pointer to **int32** | The unique identifier of a farcaster user or app (unsigned integer) | [optional] 
**Permissions** | Pointer to [**[]SharedSignerPermission**](SharedSignerPermission.md) |  | [optional] 

## Methods

### NewSigner

`func NewSigner(signerUuid string, publicKey string, status string, ) *Signer`

NewSigner instantiates a new Signer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignerWithDefaults

`func NewSignerWithDefaults() *Signer`

NewSignerWithDefaults instantiates a new Signer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Signer) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Signer) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Signer) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Signer) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetSignerUuid

`func (o *Signer) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *Signer) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *Signer) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetPublicKey

`func (o *Signer) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Signer) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Signer) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetStatus

`func (o *Signer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Signer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Signer) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSignerApprovalUrl

`func (o *Signer) GetSignerApprovalUrl() string`

GetSignerApprovalUrl returns the SignerApprovalUrl field if non-nil, zero value otherwise.

### GetSignerApprovalUrlOk

`func (o *Signer) GetSignerApprovalUrlOk() (*string, bool)`

GetSignerApprovalUrlOk returns a tuple with the SignerApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerApprovalUrl

`func (o *Signer) SetSignerApprovalUrl(v string)`

SetSignerApprovalUrl sets SignerApprovalUrl field to given value.

### HasSignerApprovalUrl

`func (o *Signer) HasSignerApprovalUrl() bool`

HasSignerApprovalUrl returns a boolean if a field has been set.

### GetFid

`func (o *Signer) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *Signer) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *Signer) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *Signer) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetPermissions

`func (o *Signer) GetPermissions() []SharedSignerPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Signer) GetPermissionsOk() (*[]SharedSignerPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Signer) SetPermissions(v []SharedSignerPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Signer) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


