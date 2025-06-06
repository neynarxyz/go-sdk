# GetNotificationCampaignStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationCampaigns** | [**[]GetNotificationCampaignStats200ResponseNotificationCampaignsInner**](GetNotificationCampaignStats200ResponseNotificationCampaignsInner.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewGetNotificationCampaignStats200Response

`func NewGetNotificationCampaignStats200Response(notificationCampaigns []GetNotificationCampaignStats200ResponseNotificationCampaignsInner, next NextCursor, ) *GetNotificationCampaignStats200Response`

NewGetNotificationCampaignStats200Response instantiates a new GetNotificationCampaignStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationCampaignStats200ResponseWithDefaults

`func NewGetNotificationCampaignStats200ResponseWithDefaults() *GetNotificationCampaignStats200Response`

NewGetNotificationCampaignStats200ResponseWithDefaults instantiates a new GetNotificationCampaignStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationCampaigns

`func (o *GetNotificationCampaignStats200Response) GetNotificationCampaigns() []GetNotificationCampaignStats200ResponseNotificationCampaignsInner`

GetNotificationCampaigns returns the NotificationCampaigns field if non-nil, zero value otherwise.

### GetNotificationCampaignsOk

`func (o *GetNotificationCampaignStats200Response) GetNotificationCampaignsOk() (*[]GetNotificationCampaignStats200ResponseNotificationCampaignsInner, bool)`

GetNotificationCampaignsOk returns a tuple with the NotificationCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCampaigns

`func (o *GetNotificationCampaignStats200Response) SetNotificationCampaigns(v []GetNotificationCampaignStats200ResponseNotificationCampaignsInner)`

SetNotificationCampaigns sets NotificationCampaigns field to given value.


### GetNext

`func (o *GetNotificationCampaignStats200Response) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetNotificationCampaignStats200Response) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetNotificationCampaignStats200Response) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


