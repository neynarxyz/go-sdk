# FrameNotificationTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationTokens** | [**[]FrameNotificationTokensNotificationTokensInner**](FrameNotificationTokensNotificationTokensInner.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewFrameNotificationTokens

`func NewFrameNotificationTokens(notificationTokens []FrameNotificationTokensNotificationTokensInner, next NextCursor, ) *FrameNotificationTokens`

NewFrameNotificationTokens instantiates a new FrameNotificationTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameNotificationTokensWithDefaults

`func NewFrameNotificationTokensWithDefaults() *FrameNotificationTokens`

NewFrameNotificationTokensWithDefaults instantiates a new FrameNotificationTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationTokens

`func (o *FrameNotificationTokens) GetNotificationTokens() []FrameNotificationTokensNotificationTokensInner`

GetNotificationTokens returns the NotificationTokens field if non-nil, zero value otherwise.

### GetNotificationTokensOk

`func (o *FrameNotificationTokens) GetNotificationTokensOk() (*[]FrameNotificationTokensNotificationTokensInner, bool)`

GetNotificationTokensOk returns a tuple with the NotificationTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTokens

`func (o *FrameNotificationTokens) SetNotificationTokens(v []FrameNotificationTokensNotificationTokensInner)`

SetNotificationTokens sets NotificationTokens field to given value.


### GetNext

`func (o *FrameNotificationTokens) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FrameNotificationTokens) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FrameNotificationTokens) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


