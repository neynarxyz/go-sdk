# SubscriptionTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to [**SubscriptionTierPrice**](SubscriptionTierPrice.md) |  | [optional] 

## Methods

### NewSubscriptionTier

`func NewSubscriptionTier() *SubscriptionTier`

NewSubscriptionTier instantiates a new SubscriptionTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionTierWithDefaults

`func NewSubscriptionTierWithDefaults() *SubscriptionTier`

NewSubscriptionTierWithDefaults instantiates a new SubscriptionTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionTier) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionTier) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionTier) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionTier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrice

`func (o *SubscriptionTier) GetPrice() SubscriptionTierPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SubscriptionTier) GetPriceOk() (*SubscriptionTierPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SubscriptionTier) SetPrice(v SubscriptionTierPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SubscriptionTier) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


