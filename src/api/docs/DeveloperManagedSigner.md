# DeveloperManagedSigner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | **string** | Ed25519 public key | 
**Status** | **string** |  | 
**SignerApprovalUrl** | Pointer to **string** |  | [optional] 
**Fid** | Pointer to **int32** | The unique identifier of a farcaster user (unsigned integer) | [optional] 

## Methods

### NewDeveloperManagedSigner

`func NewDeveloperManagedSigner(publicKey string, status string, ) *DeveloperManagedSigner`

NewDeveloperManagedSigner instantiates a new DeveloperManagedSigner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeveloperManagedSignerWithDefaults

`func NewDeveloperManagedSignerWithDefaults() *DeveloperManagedSigner`

NewDeveloperManagedSignerWithDefaults instantiates a new DeveloperManagedSigner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *DeveloperManagedSigner) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DeveloperManagedSigner) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DeveloperManagedSigner) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetStatus

`func (o *DeveloperManagedSigner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeveloperManagedSigner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeveloperManagedSigner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSignerApprovalUrl

`func (o *DeveloperManagedSigner) GetSignerApprovalUrl() string`

GetSignerApprovalUrl returns the SignerApprovalUrl field if non-nil, zero value otherwise.

### GetSignerApprovalUrlOk

`func (o *DeveloperManagedSigner) GetSignerApprovalUrlOk() (*string, bool)`

GetSignerApprovalUrlOk returns a tuple with the SignerApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerApprovalUrl

`func (o *DeveloperManagedSigner) SetSignerApprovalUrl(v string)`

SetSignerApprovalUrl sets SignerApprovalUrl field to given value.

### HasSignerApprovalUrl

`func (o *DeveloperManagedSigner) HasSignerApprovalUrl() bool`

HasSignerApprovalUrl returns a boolean if a field has been set.

### GetFid

`func (o *DeveloperManagedSigner) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *DeveloperManagedSigner) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *DeveloperManagedSigner) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *DeveloperManagedSigner) HasFid() bool`

HasFid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


