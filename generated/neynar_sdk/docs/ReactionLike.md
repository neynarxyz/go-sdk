# ReactionLike

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | The unique identifier of a farcaster user or app (unsigned integer) | 
**Fname** | **string** |  | 

## Methods

### NewReactionLike

`func NewReactionLike(fid int32, fname string, ) *ReactionLike`

NewReactionLike instantiates a new ReactionLike object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionLikeWithDefaults

`func NewReactionLikeWithDefaults() *ReactionLike`

NewReactionLikeWithDefaults instantiates a new ReactionLike object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *ReactionLike) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *ReactionLike) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *ReactionLike) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetFname

`func (o *ReactionLike) GetFname() string`

GetFname returns the Fname field if non-nil, zero value otherwise.

### GetFnameOk

`func (o *ReactionLike) GetFnameOk() (*string, bool)`

GetFnameOk returns a tuple with the Fname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFname

`func (o *ReactionLike) SetFname(v string)`

SetFname sets Fname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


