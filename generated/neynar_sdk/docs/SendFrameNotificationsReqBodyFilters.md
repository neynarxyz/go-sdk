# SendFrameNotificationsReqBodyFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeFids** | Pointer to **[]int32** | Only send notifications to users who are not in the given FIDs. | [optional] 
**FollowingFid** | Pointer to **int32** | Only send notifications to users who follow the given FID. | [optional] 
**MinimumUserScore** | Pointer to **float32** | Only send notifications to users with a score greater than or equal to this value. | [optional] 
**NearLocation** | Pointer to [**Location**](Location.md) |  | [optional] 

## Methods

### NewSendFrameNotificationsReqBodyFilters

`func NewSendFrameNotificationsReqBodyFilters() *SendFrameNotificationsReqBodyFilters`

NewSendFrameNotificationsReqBodyFilters instantiates a new SendFrameNotificationsReqBodyFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendFrameNotificationsReqBodyFiltersWithDefaults

`func NewSendFrameNotificationsReqBodyFiltersWithDefaults() *SendFrameNotificationsReqBodyFilters`

NewSendFrameNotificationsReqBodyFiltersWithDefaults instantiates a new SendFrameNotificationsReqBodyFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeFids

`func (o *SendFrameNotificationsReqBodyFilters) GetExcludeFids() []int32`

GetExcludeFids returns the ExcludeFids field if non-nil, zero value otherwise.

### GetExcludeFidsOk

`func (o *SendFrameNotificationsReqBodyFilters) GetExcludeFidsOk() (*[]int32, bool)`

GetExcludeFidsOk returns a tuple with the ExcludeFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFids

`func (o *SendFrameNotificationsReqBodyFilters) SetExcludeFids(v []int32)`

SetExcludeFids sets ExcludeFids field to given value.

### HasExcludeFids

`func (o *SendFrameNotificationsReqBodyFilters) HasExcludeFids() bool`

HasExcludeFids returns a boolean if a field has been set.

### GetFollowingFid

`func (o *SendFrameNotificationsReqBodyFilters) GetFollowingFid() int32`

GetFollowingFid returns the FollowingFid field if non-nil, zero value otherwise.

### GetFollowingFidOk

`func (o *SendFrameNotificationsReqBodyFilters) GetFollowingFidOk() (*int32, bool)`

GetFollowingFidOk returns a tuple with the FollowingFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingFid

`func (o *SendFrameNotificationsReqBodyFilters) SetFollowingFid(v int32)`

SetFollowingFid sets FollowingFid field to given value.

### HasFollowingFid

`func (o *SendFrameNotificationsReqBodyFilters) HasFollowingFid() bool`

HasFollowingFid returns a boolean if a field has been set.

### GetMinimumUserScore

`func (o *SendFrameNotificationsReqBodyFilters) GetMinimumUserScore() float32`

GetMinimumUserScore returns the MinimumUserScore field if non-nil, zero value otherwise.

### GetMinimumUserScoreOk

`func (o *SendFrameNotificationsReqBodyFilters) GetMinimumUserScoreOk() (*float32, bool)`

GetMinimumUserScoreOk returns a tuple with the MinimumUserScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumUserScore

`func (o *SendFrameNotificationsReqBodyFilters) SetMinimumUserScore(v float32)`

SetMinimumUserScore sets MinimumUserScore field to given value.

### HasMinimumUserScore

`func (o *SendFrameNotificationsReqBodyFilters) HasMinimumUserScore() bool`

HasMinimumUserScore returns a boolean if a field has been set.

### GetNearLocation

`func (o *SendFrameNotificationsReqBodyFilters) GetNearLocation() Location`

GetNearLocation returns the NearLocation field if non-nil, zero value otherwise.

### GetNearLocationOk

`func (o *SendFrameNotificationsReqBodyFilters) GetNearLocationOk() (*Location, bool)`

GetNearLocationOk returns a tuple with the NearLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearLocation

`func (o *SendFrameNotificationsReqBodyFilters) SetNearLocation(v Location)`

SetNearLocation sets NearLocation field to given value.

### HasNearLocation

`func (o *SendFrameNotificationsReqBodyFilters) HasNearLocation() bool`

HasNearLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


