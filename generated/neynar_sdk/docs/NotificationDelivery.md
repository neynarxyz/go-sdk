# NotificationDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Fid** | **int32** | The unique identifier of a farcaster user (unsigned integer) | 
**Status** | **string** |  | 

## Methods

### NewNotificationDelivery

`func NewNotificationDelivery(object string, fid int32, status string, ) *NotificationDelivery`

NewNotificationDelivery instantiates a new NotificationDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDeliveryWithDefaults

`func NewNotificationDeliveryWithDefaults() *NotificationDelivery`

NewNotificationDeliveryWithDefaults instantiates a new NotificationDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *NotificationDelivery) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *NotificationDelivery) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *NotificationDelivery) SetObject(v string)`

SetObject sets Object field to given value.


### GetFid

`func (o *NotificationDelivery) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *NotificationDelivery) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *NotificationDelivery) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetStatus

`func (o *NotificationDelivery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationDelivery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationDelivery) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


