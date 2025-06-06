# AppHostGetEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | Legacy event type corresponding to the requested event type: - frame_added: User adds a mini app to their account - frame_removed: User removes a mini app from their account - notifications_enabled: User enables notifications for a mini app - notifications_disabled: User disables notifications for a mini app | 
**NotificationDetails** | Pointer to [**AppHostGetEventResponseNotificationDetails**](AppHostGetEventResponseNotificationDetails.md) |  | [optional] 

## Methods

### NewAppHostGetEventResponse

`func NewAppHostGetEventResponse(event string, ) *AppHostGetEventResponse`

NewAppHostGetEventResponse instantiates a new AppHostGetEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppHostGetEventResponseWithDefaults

`func NewAppHostGetEventResponseWithDefaults() *AppHostGetEventResponse`

NewAppHostGetEventResponseWithDefaults instantiates a new AppHostGetEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AppHostGetEventResponse) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AppHostGetEventResponse) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AppHostGetEventResponse) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetNotificationDetails

`func (o *AppHostGetEventResponse) GetNotificationDetails() AppHostGetEventResponseNotificationDetails`

GetNotificationDetails returns the NotificationDetails field if non-nil, zero value otherwise.

### GetNotificationDetailsOk

`func (o *AppHostGetEventResponse) GetNotificationDetailsOk() (*AppHostGetEventResponseNotificationDetails, bool)`

GetNotificationDetailsOk returns a tuple with the NotificationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDetails

`func (o *AppHostGetEventResponse) SetNotificationDetails(v AppHostGetEventResponseNotificationDetails)`

SetNotificationDetails sets NotificationDetails field to given value.

### HasNotificationDetails

`func (o *AppHostGetEventResponse) HasNotificationDetails() bool`

HasNotificationDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


