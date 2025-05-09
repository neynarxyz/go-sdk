# LinkBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**LinkType**](LinkType.md) |  | [default to LINKTYPE_FOLLOW]
**DisplayTimestamp** | Pointer to **int64** | User-defined timestamp that preserves the original creation time when message.data.timestamp needs to be updated for compaction. Useful for maintaining chronological order in user feeds. | [optional] 
**TargetFid** | **int32** | Farcaster ID (FID). A unique identifier assigned to each user in the Farcaster network. This number is permanent and cannot be changed. FIDs are assigned sequentially when users register on the network.  | 

## Methods

### NewLinkBody

`func NewLinkBody(type_ LinkType, targetFid int32, ) *LinkBody`

NewLinkBody instantiates a new LinkBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkBodyWithDefaults

`func NewLinkBodyWithDefaults() *LinkBody`

NewLinkBodyWithDefaults instantiates a new LinkBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LinkBody) GetType() LinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkBody) GetTypeOk() (*LinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkBody) SetType(v LinkType)`

SetType sets Type field to given value.


### GetDisplayTimestamp

`func (o *LinkBody) GetDisplayTimestamp() int64`

GetDisplayTimestamp returns the DisplayTimestamp field if non-nil, zero value otherwise.

### GetDisplayTimestampOk

`func (o *LinkBody) GetDisplayTimestampOk() (*int64, bool)`

GetDisplayTimestampOk returns a tuple with the DisplayTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTimestamp

`func (o *LinkBody) SetDisplayTimestamp(v int64)`

SetDisplayTimestamp sets DisplayTimestamp field to given value.

### HasDisplayTimestamp

`func (o *LinkBody) HasDisplayTimestamp() bool`

HasDisplayTimestamp returns a boolean if a field has been set.

### GetTargetFid

`func (o *LinkBody) GetTargetFid() int32`

GetTargetFid returns the TargetFid field if non-nil, zero value otherwise.

### GetTargetFidOk

`func (o *LinkBody) GetTargetFidOk() (*int32, bool)`

GetTargetFidOk returns a tuple with the TargetFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFid

`func (o *LinkBody) SetTargetFid(v int32)`

SetTargetFid sets TargetFid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


