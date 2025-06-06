# SubscriptionCheckResponseAnyOfValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Status** | **bool** |  | 
**ExpiresAt** | **NullableInt64** |  | 
**SubscribedAt** | **NullableInt64** |  | 
**Tier** | [**NullableSubscriptionTier**](SubscriptionTier.md) |  | 

## Methods

### NewSubscriptionCheckResponseAnyOfValue

`func NewSubscriptionCheckResponseAnyOfValue(object string, status bool, expiresAt NullableInt64, subscribedAt NullableInt64, tier NullableSubscriptionTier, ) *SubscriptionCheckResponseAnyOfValue`

NewSubscriptionCheckResponseAnyOfValue instantiates a new SubscriptionCheckResponseAnyOfValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCheckResponseAnyOfValueWithDefaults

`func NewSubscriptionCheckResponseAnyOfValueWithDefaults() *SubscriptionCheckResponseAnyOfValue`

NewSubscriptionCheckResponseAnyOfValueWithDefaults instantiates a new SubscriptionCheckResponseAnyOfValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *SubscriptionCheckResponseAnyOfValue) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionCheckResponseAnyOfValue) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionCheckResponseAnyOfValue) SetObject(v string)`

SetObject sets Object field to given value.


### GetStatus

`func (o *SubscriptionCheckResponseAnyOfValue) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionCheckResponseAnyOfValue) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionCheckResponseAnyOfValue) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetExpiresAt

`func (o *SubscriptionCheckResponseAnyOfValue) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SubscriptionCheckResponseAnyOfValue) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SubscriptionCheckResponseAnyOfValue) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *SubscriptionCheckResponseAnyOfValue) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *SubscriptionCheckResponseAnyOfValue) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetSubscribedAt

`func (o *SubscriptionCheckResponseAnyOfValue) GetSubscribedAt() int64`

GetSubscribedAt returns the SubscribedAt field if non-nil, zero value otherwise.

### GetSubscribedAtOk

`func (o *SubscriptionCheckResponseAnyOfValue) GetSubscribedAtOk() (*int64, bool)`

GetSubscribedAtOk returns a tuple with the SubscribedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedAt

`func (o *SubscriptionCheckResponseAnyOfValue) SetSubscribedAt(v int64)`

SetSubscribedAt sets SubscribedAt field to given value.


### SetSubscribedAtNil

`func (o *SubscriptionCheckResponseAnyOfValue) SetSubscribedAtNil(b bool)`

 SetSubscribedAtNil sets the value for SubscribedAt to be an explicit nil

### UnsetSubscribedAt
`func (o *SubscriptionCheckResponseAnyOfValue) UnsetSubscribedAt()`

UnsetSubscribedAt ensures that no value is present for SubscribedAt, not even an explicit nil
### GetTier

`func (o *SubscriptionCheckResponseAnyOfValue) GetTier() SubscriptionTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SubscriptionCheckResponseAnyOfValue) GetTierOk() (*SubscriptionTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SubscriptionCheckResponseAnyOfValue) SetTier(v SubscriptionTier)`

SetTier sets Tier field to given value.


### SetTierNil

`func (o *SubscriptionCheckResponseAnyOfValue) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *SubscriptionCheckResponseAnyOfValue) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


