# GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntendedRecipientNotificationTokenCount** | **int32** | The total number of notification tokens for intended recipients. | 
**IntendedRecipientAppFids** | **[]int32** | An array of Farcaster FIDs of intended recipient applications. | 
**SuccessfulSends** | **int32** | The number of notifications successfully sent. | 
**SuccessfulSendsByAppFid** | **map[string]int32** | A record mapping app FIDs (as strings) to the number of successful sends for that app. | 
**TotalOpens** | **int32** | The total number of times notifications from this campaign have been opened. | 
**TotalOpensByAppFid** | **map[string]int32** | A record mapping app FIDs (as strings) to the number of opens for that app. | 
**UniqueOpens** | **int32** | The number of unique recipients who opened a notification from this campaign. | 
**UniqueOpensByAppFid** | **map[string]int32** | A record mapping app FIDs (as strings) to the number of unique opens for that app. | 

## Methods

### NewGetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats

`func NewGetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats(intendedRecipientNotificationTokenCount int32, intendedRecipientAppFids []int32, successfulSends int32, successfulSendsByAppFid map[string]int32, totalOpens int32, totalOpensByAppFid map[string]int32, uniqueOpens int32, uniqueOpensByAppFid map[string]int32, ) *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats`

NewGetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats instantiates a new GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationCampaignStats200ResponseNotificationCampaignsInnerStatsWithDefaults

`func NewGetNotificationCampaignStats200ResponseNotificationCampaignsInnerStatsWithDefaults() *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats`

NewGetNotificationCampaignStats200ResponseNotificationCampaignsInnerStatsWithDefaults instantiates a new GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntendedRecipientNotificationTokenCount

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetIntendedRecipientNotificationTokenCount() int32`

GetIntendedRecipientNotificationTokenCount returns the IntendedRecipientNotificationTokenCount field if non-nil, zero value otherwise.

### GetIntendedRecipientNotificationTokenCountOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetIntendedRecipientNotificationTokenCountOk() (*int32, bool)`

GetIntendedRecipientNotificationTokenCountOk returns a tuple with the IntendedRecipientNotificationTokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedRecipientNotificationTokenCount

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) SetIntendedRecipientNotificationTokenCount(v int32)`

SetIntendedRecipientNotificationTokenCount sets IntendedRecipientNotificationTokenCount field to given value.


### GetIntendedRecipientAppFids

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetIntendedRecipientAppFids() []int32`

GetIntendedRecipientAppFids returns the IntendedRecipientAppFids field if non-nil, zero value otherwise.

### GetIntendedRecipientAppFidsOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetIntendedRecipientAppFidsOk() (*[]int32, bool)`

GetIntendedRecipientAppFidsOk returns a tuple with the IntendedRecipientAppFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedRecipientAppFids

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) SetIntendedRecipientAppFids(v []int32)`

SetIntendedRecipientAppFids sets IntendedRecipientAppFids field to given value.


### GetSuccessfulSends

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetSuccessfulSends() int32`

GetSuccessfulSends returns the SuccessfulSends field if non-nil, zero value otherwise.

### GetSuccessfulSendsOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetSuccessfulSendsOk() (*int32, bool)`

GetSuccessfulSendsOk returns a tuple with the SuccessfulSends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulSends

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) SetSuccessfulSends(v int32)`

SetSuccessfulSends sets SuccessfulSends field to given value.


### GetSuccessfulSendsByAppFid

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetSuccessfulSendsByAppFid() map[string]int32`

GetSuccessfulSendsByAppFid returns the SuccessfulSendsByAppFid field if non-nil, zero value otherwise.

### GetSuccessfulSendsByAppFidOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetSuccessfulSendsByAppFidOk() (*map[string]int32, bool)`

GetSuccessfulSendsByAppFidOk returns a tuple with the SuccessfulSendsByAppFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulSendsByAppFid

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) SetSuccessfulSendsByAppFid(v map[string]int32)`

SetSuccessfulSendsByAppFid sets SuccessfulSendsByAppFid field to given value.


### GetTotalOpens

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetTotalOpens() int32`

GetTotalOpens returns the TotalOpens field if non-nil, zero value otherwise.

### GetTotalOpensOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetTotalOpensOk() (*int32, bool)`

GetTotalOpensOk returns a tuple with the TotalOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpens

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) SetTotalOpens(v int32)`

SetTotalOpens sets TotalOpens field to given value.


### GetTotalOpensByAppFid

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetTotalOpensByAppFid() map[string]int32`

GetTotalOpensByAppFid returns the TotalOpensByAppFid field if non-nil, zero value otherwise.

### GetTotalOpensByAppFidOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetTotalOpensByAppFidOk() (*map[string]int32, bool)`

GetTotalOpensByAppFidOk returns a tuple with the TotalOpensByAppFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpensByAppFid

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) SetTotalOpensByAppFid(v map[string]int32)`

SetTotalOpensByAppFid sets TotalOpensByAppFid field to given value.


### GetUniqueOpens

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetUniqueOpens() int32`

GetUniqueOpens returns the UniqueOpens field if non-nil, zero value otherwise.

### GetUniqueOpensOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetUniqueOpensOk() (*int32, bool)`

GetUniqueOpensOk returns a tuple with the UniqueOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueOpens

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) SetUniqueOpens(v int32)`

SetUniqueOpens sets UniqueOpens field to given value.


### GetUniqueOpensByAppFid

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetUniqueOpensByAppFid() map[string]int32`

GetUniqueOpensByAppFid returns the UniqueOpensByAppFid field if non-nil, zero value otherwise.

### GetUniqueOpensByAppFidOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) GetUniqueOpensByAppFidOk() (*map[string]int32, bool)`

GetUniqueOpensByAppFidOk returns a tuple with the UniqueOpensByAppFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueOpensByAppFid

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats) SetUniqueOpensByAppFid(v map[string]int32)`

SetUniqueOpensByAppFid sets UniqueOpensByAppFid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


