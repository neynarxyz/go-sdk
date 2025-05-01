# SendFrameNotificationsReqBodyNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The title of the notification. Must be between 1 and 32 characters. | 
**Body** | **string** | The body of the notification. Must be between 1 and 128 characters. | 
**TargetUrl** | **string** | The target URL to open when the user clicks the notification. Must be a valid URL. | 
**Uuid** | Pointer to **string** | An optional UUID for the notification, used as an idempotency key. | [optional] 

## Methods

### NewSendFrameNotificationsReqBodyNotification

`func NewSendFrameNotificationsReqBodyNotification(title string, body string, targetUrl string, ) *SendFrameNotificationsReqBodyNotification`

NewSendFrameNotificationsReqBodyNotification instantiates a new SendFrameNotificationsReqBodyNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendFrameNotificationsReqBodyNotificationWithDefaults

`func NewSendFrameNotificationsReqBodyNotificationWithDefaults() *SendFrameNotificationsReqBodyNotification`

NewSendFrameNotificationsReqBodyNotificationWithDefaults instantiates a new SendFrameNotificationsReqBodyNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SendFrameNotificationsReqBodyNotification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendFrameNotificationsReqBodyNotification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendFrameNotificationsReqBodyNotification) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *SendFrameNotificationsReqBodyNotification) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendFrameNotificationsReqBodyNotification) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendFrameNotificationsReqBodyNotification) SetBody(v string)`

SetBody sets Body field to given value.


### GetTargetUrl

`func (o *SendFrameNotificationsReqBodyNotification) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *SendFrameNotificationsReqBodyNotification) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *SendFrameNotificationsReqBodyNotification) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetUuid

`func (o *SendFrameNotificationsReqBodyNotification) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SendFrameNotificationsReqBodyNotification) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SendFrameNotificationsReqBodyNotification) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SendFrameNotificationsReqBodyNotification) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


