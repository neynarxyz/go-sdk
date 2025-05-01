# HubInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**IsSyncing** | **bool** |  | 
**Nickname** | **string** |  | 
**RootHash** | **string** |  | 
**DbStats** | Pointer to [**DbStats**](DbStats.md) |  | [optional] 
**PeerId** | **string** |  | 
**HubOperatorFid** | **int32** |  | 

## Methods

### NewHubInfoResponse

`func NewHubInfoResponse(version string, isSyncing bool, nickname string, rootHash string, peerId string, hubOperatorFid int32, ) *HubInfoResponse`

NewHubInfoResponse instantiates a new HubInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubInfoResponseWithDefaults

`func NewHubInfoResponseWithDefaults() *HubInfoResponse`

NewHubInfoResponseWithDefaults instantiates a new HubInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *HubInfoResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HubInfoResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HubInfoResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetIsSyncing

`func (o *HubInfoResponse) GetIsSyncing() bool`

GetIsSyncing returns the IsSyncing field if non-nil, zero value otherwise.

### GetIsSyncingOk

`func (o *HubInfoResponse) GetIsSyncingOk() (*bool, bool)`

GetIsSyncingOk returns a tuple with the IsSyncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncing

`func (o *HubInfoResponse) SetIsSyncing(v bool)`

SetIsSyncing sets IsSyncing field to given value.


### GetNickname

`func (o *HubInfoResponse) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *HubInfoResponse) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *HubInfoResponse) SetNickname(v string)`

SetNickname sets Nickname field to given value.


### GetRootHash

`func (o *HubInfoResponse) GetRootHash() string`

GetRootHash returns the RootHash field if non-nil, zero value otherwise.

### GetRootHashOk

`func (o *HubInfoResponse) GetRootHashOk() (*string, bool)`

GetRootHashOk returns a tuple with the RootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootHash

`func (o *HubInfoResponse) SetRootHash(v string)`

SetRootHash sets RootHash field to given value.


### GetDbStats

`func (o *HubInfoResponse) GetDbStats() DbStats`

GetDbStats returns the DbStats field if non-nil, zero value otherwise.

### GetDbStatsOk

`func (o *HubInfoResponse) GetDbStatsOk() (*DbStats, bool)`

GetDbStatsOk returns a tuple with the DbStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbStats

`func (o *HubInfoResponse) SetDbStats(v DbStats)`

SetDbStats sets DbStats field to given value.

### HasDbStats

`func (o *HubInfoResponse) HasDbStats() bool`

HasDbStats returns a boolean if a field has been set.

### GetPeerId

`func (o *HubInfoResponse) GetPeerId() string`

GetPeerId returns the PeerId field if non-nil, zero value otherwise.

### GetPeerIdOk

`func (o *HubInfoResponse) GetPeerIdOk() (*string, bool)`

GetPeerIdOk returns a tuple with the PeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerId

`func (o *HubInfoResponse) SetPeerId(v string)`

SetPeerId sets PeerId field to given value.


### GetHubOperatorFid

`func (o *HubInfoResponse) GetHubOperatorFid() int32`

GetHubOperatorFid returns the HubOperatorFid field if non-nil, zero value otherwise.

### GetHubOperatorFidOk

`func (o *HubInfoResponse) GetHubOperatorFidOk() (*int32, bool)`

GetHubOperatorFidOk returns a tuple with the HubOperatorFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubOperatorFid

`func (o *HubInfoResponse) SetHubOperatorFid(v int32)`

SetHubOperatorFid sets HubOperatorFid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


