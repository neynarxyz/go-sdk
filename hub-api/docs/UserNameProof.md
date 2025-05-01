# UserNameProof

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int32** | Seconds since Unix Epoch which began on Jan 1, 1970 00:00:00 UTC | 
**Name** | **string** |  | 
**Owner** | **string** |  | 
**Signature** | **string** |  | 
**Fid** | **int32** | Farcaster ID (FID). A unique identifier assigned to each user in the Farcaster network. This number is permanent and cannot be changed. FIDs are assigned sequentially when users register on the network.  | 
**Type** | [**UserNameType**](UserNameType.md) |  | [default to USERNAME_TYPE_FNAME]

## Methods

### NewUserNameProof

`func NewUserNameProof(timestamp int32, name string, owner string, signature string, fid int32, type_ UserNameType, ) *UserNameProof`

NewUserNameProof instantiates a new UserNameProof object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNameProofWithDefaults

`func NewUserNameProofWithDefaults() *UserNameProof`

NewUserNameProofWithDefaults instantiates a new UserNameProof object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *UserNameProof) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserNameProof) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserNameProof) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetName

`func (o *UserNameProof) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserNameProof) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserNameProof) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *UserNameProof) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UserNameProof) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UserNameProof) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetSignature

`func (o *UserNameProof) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UserNameProof) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UserNameProof) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetFid

`func (o *UserNameProof) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *UserNameProof) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *UserNameProof) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetType

`func (o *UserNameProof) GetType() UserNameType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserNameProof) GetTypeOk() (*UserNameType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserNameProof) SetType(v UserNameType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


