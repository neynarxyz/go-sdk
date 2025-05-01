# SubscriptionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Status** | **bool** |  | 
**ExpiresAt** | **NullableInt64** |  | 
**SubscribedAt** | **NullableInt64** |  | 
**Tier** | [**NullableSubscriptionTier**](SubscriptionTier.md) |  | 

## Methods

### NewSubscriptionStatus

`func NewSubscriptionStatus(object string, status bool, expiresAt NullableInt64, subscribedAt NullableInt64, tier NullableSubscriptionTier, ) *SubscriptionStatus`

NewSubscriptionStatus instantiates a new SubscriptionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionStatusWithDefaults

`func NewSubscriptionStatusWithDefaults() *SubscriptionStatus`

NewSubscriptionStatusWithDefaults instantiates a new SubscriptionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *SubscriptionStatus) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionStatus) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionStatus) SetObject(v string)`

SetObject sets Object field to given value.


### GetStatus

`func (o *SubscriptionStatus) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionStatus) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionStatus) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetExpiresAt

`func (o *SubscriptionStatus) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SubscriptionStatus) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SubscriptionStatus) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *SubscriptionStatus) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *SubscriptionStatus) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetSubscribedAt

`func (o *SubscriptionStatus) GetSubscribedAt() int64`

GetSubscribedAt returns the SubscribedAt field if non-nil, zero value otherwise.

### GetSubscribedAtOk

`func (o *SubscriptionStatus) GetSubscribedAtOk() (*int64, bool)`

GetSubscribedAtOk returns a tuple with the SubscribedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedAt

`func (o *SubscriptionStatus) SetSubscribedAt(v int64)`

SetSubscribedAt sets SubscribedAt field to given value.


### SetSubscribedAtNil

`func (o *SubscriptionStatus) SetSubscribedAtNil(b bool)`

 SetSubscribedAtNil sets the value for SubscribedAt to be an explicit nil

### UnsetSubscribedAt
`func (o *SubscriptionStatus) UnsetSubscribedAt()`

UnsetSubscribedAt ensures that no value is present for SubscribedAt, not even an explicit nil
### GetTier

`func (o *SubscriptionStatus) GetTier() SubscriptionTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SubscriptionStatus) GetTierOk() (*SubscriptionTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SubscriptionStatus) SetTier(v SubscriptionTier)`

SetTier sets Tier field to given value.


### SetTierNil

`func (o *SubscriptionStatus) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *SubscriptionStatus) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


