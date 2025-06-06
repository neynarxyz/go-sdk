# AppHostUserStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationsEnabled** | [**[]AppHostUserStateResponseNotificationsEnabledInner**](AppHostUserStateResponseNotificationsEnabledInner.md) | List of domains for which notifications are enabled for this user | 

## Methods

### NewAppHostUserStateResponse

`func NewAppHostUserStateResponse(notificationsEnabled []AppHostUserStateResponseNotificationsEnabledInner, ) *AppHostUserStateResponse`

NewAppHostUserStateResponse instantiates a new AppHostUserStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppHostUserStateResponseWithDefaults

`func NewAppHostUserStateResponseWithDefaults() *AppHostUserStateResponse`

NewAppHostUserStateResponseWithDefaults instantiates a new AppHostUserStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationsEnabled

`func (o *AppHostUserStateResponse) GetNotificationsEnabled() []AppHostUserStateResponseNotificationsEnabledInner`

GetNotificationsEnabled returns the NotificationsEnabled field if non-nil, zero value otherwise.

### GetNotificationsEnabledOk

`func (o *AppHostUserStateResponse) GetNotificationsEnabledOk() (*[]AppHostUserStateResponseNotificationsEnabledInner, bool)`

GetNotificationsEnabledOk returns a tuple with the NotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEnabled

`func (o *AppHostUserStateResponse) SetNotificationsEnabled(v []AppHostUserStateResponseNotificationsEnabledInner)`

SetNotificationsEnabled sets NotificationsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


