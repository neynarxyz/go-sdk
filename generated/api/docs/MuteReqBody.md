# MuteReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | The unique identifier of a farcaster user (unsigned integer) | 
**MutedFid** | **int32** | The unique identifier of a farcaster user (unsigned integer) | 

## Methods

### NewMuteReqBody

`func NewMuteReqBody(fid int32, mutedFid int32, ) *MuteReqBody`

NewMuteReqBody instantiates a new MuteReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMuteReqBodyWithDefaults

`func NewMuteReqBodyWithDefaults() *MuteReqBody`

NewMuteReqBodyWithDefaults instantiates a new MuteReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MuteReqBody) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MuteReqBody) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MuteReqBody) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetMutedFid

`func (o *MuteReqBody) GetMutedFid() int32`

GetMutedFid returns the MutedFid field if non-nil, zero value otherwise.

### GetMutedFidOk

`func (o *MuteReqBody) GetMutedFidOk() (*int32, bool)`

GetMutedFidOk returns a tuple with the MutedFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedFid

`func (o *MuteReqBody) SetMutedFid(v int32)`

SetMutedFid sets MutedFid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


