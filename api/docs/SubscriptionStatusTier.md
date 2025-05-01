# SubscriptionStatusTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to [**SubscriptionTierPrice**](SubscriptionTierPrice.md) |  | [optional] 

## Methods

### NewSubscriptionStatusTier

`func NewSubscriptionStatusTier() *SubscriptionStatusTier`

NewSubscriptionStatusTier instantiates a new SubscriptionStatusTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionStatusTierWithDefaults

`func NewSubscriptionStatusTierWithDefaults() *SubscriptionStatusTier`

NewSubscriptionStatusTierWithDefaults instantiates a new SubscriptionStatusTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionStatusTier) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionStatusTier) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionStatusTier) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionStatusTier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrice

`func (o *SubscriptionStatusTier) GetPrice() SubscriptionTierPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SubscriptionStatusTier) GetPriceOk() (*SubscriptionTierPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SubscriptionStatusTier) SetPrice(v SubscriptionTierPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SubscriptionStatusTier) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


