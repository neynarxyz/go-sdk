# RegisterUserReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | **string** |  | 
**Fid** | **float32** |  | 
**RequestedUserCustodyAddress** | **string** |  | 
**Deadline** | **float32** |  | 
**Fname** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**RegisterUserReqBodyMetadata**](RegisterUserReqBodyMetadata.md) |  | [optional] 

## Methods

### NewRegisterUserReqBody

`func NewRegisterUserReqBody(signature string, fid float32, requestedUserCustodyAddress string, deadline float32, ) *RegisterUserReqBody`

NewRegisterUserReqBody instantiates a new RegisterUserReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserReqBodyWithDefaults

`func NewRegisterUserReqBodyWithDefaults() *RegisterUserReqBody`

NewRegisterUserReqBodyWithDefaults instantiates a new RegisterUserReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *RegisterUserReqBody) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *RegisterUserReqBody) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *RegisterUserReqBody) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetFid

`func (o *RegisterUserReqBody) GetFid() float32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *RegisterUserReqBody) GetFidOk() (*float32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *RegisterUserReqBody) SetFid(v float32)`

SetFid sets Fid field to given value.


### GetRequestedUserCustodyAddress

`func (o *RegisterUserReqBody) GetRequestedUserCustodyAddress() string`

GetRequestedUserCustodyAddress returns the RequestedUserCustodyAddress field if non-nil, zero value otherwise.

### GetRequestedUserCustodyAddressOk

`func (o *RegisterUserReqBody) GetRequestedUserCustodyAddressOk() (*string, bool)`

GetRequestedUserCustodyAddressOk returns a tuple with the RequestedUserCustodyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedUserCustodyAddress

`func (o *RegisterUserReqBody) SetRequestedUserCustodyAddress(v string)`

SetRequestedUserCustodyAddress sets RequestedUserCustodyAddress field to given value.


### GetDeadline

`func (o *RegisterUserReqBody) GetDeadline() float32`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *RegisterUserReqBody) GetDeadlineOk() (*float32, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *RegisterUserReqBody) SetDeadline(v float32)`

SetDeadline sets Deadline field to given value.


### GetFname

`func (o *RegisterUserReqBody) GetFname() string`

GetFname returns the Fname field if non-nil, zero value otherwise.

### GetFnameOk

`func (o *RegisterUserReqBody) GetFnameOk() (*string, bool)`

GetFnameOk returns a tuple with the Fname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFname

`func (o *RegisterUserReqBody) SetFname(v string)`

SetFname sets Fname field to given value.

### HasFname

`func (o *RegisterUserReqBody) HasFname() bool`

HasFname returns a boolean if a field has been set.

### GetMetadata

`func (o *RegisterUserReqBody) GetMetadata() RegisterUserReqBodyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RegisterUserReqBody) GetMetadataOk() (*RegisterUserReqBodyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RegisterUserReqBody) SetMetadata(v RegisterUserReqBodyMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RegisterUserReqBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


