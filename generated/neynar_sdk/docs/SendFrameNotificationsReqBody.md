# SendFrameNotificationsReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetFids** | **[]int32** | An array of target FIDs to whom the notifications should be sent. Each FID must be a positive integer. Pass an empty array to send notifications to all FIDs with notifications enabled for the mini app. | 
**Notification** | [**SendFrameNotificationsReqBodyNotification**](SendFrameNotificationsReqBodyNotification.md) |  | 
**Filters** | Pointer to [**SendFrameNotificationsReqBodyFilters**](SendFrameNotificationsReqBodyFilters.md) |  | [optional] 

## Methods

### NewSendFrameNotificationsReqBody

`func NewSendFrameNotificationsReqBody(targetFids []int32, notification SendFrameNotificationsReqBodyNotification, ) *SendFrameNotificationsReqBody`

NewSendFrameNotificationsReqBody instantiates a new SendFrameNotificationsReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendFrameNotificationsReqBodyWithDefaults

`func NewSendFrameNotificationsReqBodyWithDefaults() *SendFrameNotificationsReqBody`

NewSendFrameNotificationsReqBodyWithDefaults instantiates a new SendFrameNotificationsReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetFids

`func (o *SendFrameNotificationsReqBody) GetTargetFids() []int32`

GetTargetFids returns the TargetFids field if non-nil, zero value otherwise.

### GetTargetFidsOk

`func (o *SendFrameNotificationsReqBody) GetTargetFidsOk() (*[]int32, bool)`

GetTargetFidsOk returns a tuple with the TargetFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFids

`func (o *SendFrameNotificationsReqBody) SetTargetFids(v []int32)`

SetTargetFids sets TargetFids field to given value.


### GetNotification

`func (o *SendFrameNotificationsReqBody) GetNotification() SendFrameNotificationsReqBodyNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *SendFrameNotificationsReqBody) GetNotificationOk() (*SendFrameNotificationsReqBodyNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *SendFrameNotificationsReqBody) SetNotification(v SendFrameNotificationsReqBodyNotification)`

SetNotification sets Notification field to given value.


### GetFilters

`func (o *SendFrameNotificationsReqBody) GetFilters() SendFrameNotificationsReqBodyFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SendFrameNotificationsReqBody) GetFiltersOk() (*SendFrameNotificationsReqBodyFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SendFrameNotificationsReqBody) SetFilters(v SendFrameNotificationsReqBodyFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SendFrameNotificationsReqBody) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


