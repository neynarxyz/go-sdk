# GetNotificationCampaignStats200ResponseNotificationCampaignsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the notification campaign. | 
**Title** | **string** | The title of the notification campaign. | 
**Body** | **string** | The body text of the notification. | 
**CreatedAt** | **time.Time** |  | 
**Stats** | [**GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats**](GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats.md) |  | 

## Methods

### NewGetNotificationCampaignStats200ResponseNotificationCampaignsInner

`func NewGetNotificationCampaignStats200ResponseNotificationCampaignsInner(id string, title string, body string, createdAt time.Time, stats GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats, ) *GetNotificationCampaignStats200ResponseNotificationCampaignsInner`

NewGetNotificationCampaignStats200ResponseNotificationCampaignsInner instantiates a new GetNotificationCampaignStats200ResponseNotificationCampaignsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationCampaignStats200ResponseNotificationCampaignsInnerWithDefaults

`func NewGetNotificationCampaignStats200ResponseNotificationCampaignsInnerWithDefaults() *GetNotificationCampaignStats200ResponseNotificationCampaignsInner`

NewGetNotificationCampaignStats200ResponseNotificationCampaignsInnerWithDefaults instantiates a new GetNotificationCampaignStats200ResponseNotificationCampaignsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStats

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) GetStats() GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) GetStatsOk() (*GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetNotificationCampaignStats200ResponseNotificationCampaignsInner) SetStats(v GetNotificationCampaignStats200ResponseNotificationCampaignsInnerStats)`

SetStats sets Stats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


