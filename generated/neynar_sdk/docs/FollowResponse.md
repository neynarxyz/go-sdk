# FollowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**TargetFid** | **int32** | The unique identifier of a farcaster user or app (unsigned integer) | 
**Hash** | **string** |  | 

## Methods

### NewFollowResponse

`func NewFollowResponse(success bool, targetFid int32, hash string, ) *FollowResponse`

NewFollowResponse instantiates a new FollowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowResponseWithDefaults

`func NewFollowResponseWithDefaults() *FollowResponse`

NewFollowResponseWithDefaults instantiates a new FollowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *FollowResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FollowResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FollowResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetTargetFid

`func (o *FollowResponse) GetTargetFid() int32`

GetTargetFid returns the TargetFid field if non-nil, zero value otherwise.

### GetTargetFidOk

`func (o *FollowResponse) GetTargetFidOk() (*int32, bool)`

GetTargetFidOk returns a tuple with the TargetFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFid

`func (o *FollowResponse) SetTargetFid(v int32)`

SetTargetFid sets TargetFid field to given value.


### GetHash

`func (o *FollowResponse) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *FollowResponse) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *FollowResponse) SetHash(v string)`

SetHash sets Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


