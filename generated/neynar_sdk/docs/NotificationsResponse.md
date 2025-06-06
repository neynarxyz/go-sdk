# NotificationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnseenNotificationsCount** | **int32** |  | 
**Notifications** | [**[]Notification**](Notification.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewNotificationsResponse

`func NewNotificationsResponse(unseenNotificationsCount int32, notifications []Notification, next NextCursor, ) *NotificationsResponse`

NewNotificationsResponse instantiates a new NotificationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsResponseWithDefaults

`func NewNotificationsResponseWithDefaults() *NotificationsResponse`

NewNotificationsResponseWithDefaults instantiates a new NotificationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnseenNotificationsCount

`func (o *NotificationsResponse) GetUnseenNotificationsCount() int32`

GetUnseenNotificationsCount returns the UnseenNotificationsCount field if non-nil, zero value otherwise.

### GetUnseenNotificationsCountOk

`func (o *NotificationsResponse) GetUnseenNotificationsCountOk() (*int32, bool)`

GetUnseenNotificationsCountOk returns a tuple with the UnseenNotificationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnseenNotificationsCount

`func (o *NotificationsResponse) SetUnseenNotificationsCount(v int32)`

SetUnseenNotificationsCount sets UnseenNotificationsCount field to given value.


### GetNotifications

`func (o *NotificationsResponse) GetNotifications() []Notification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *NotificationsResponse) GetNotificationsOk() (*[]Notification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *NotificationsResponse) SetNotifications(v []Notification)`

SetNotifications sets Notifications field to given value.


### GetNext

`func (o *NotificationsResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *NotificationsResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *NotificationsResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


