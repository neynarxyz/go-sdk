# SubscribedTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**ProviderName** | Pointer to **string** |  | [optional] 
**ContractAddress** | **string** |  | 
**Chain** | **int32** |  | 
**Metadata** | [**SubscriptionMetadata**](SubscriptionMetadata.md) |  | 
**OwnerAddress** | **string** |  | 
**Price** | [**SubscriptionPrice**](SubscriptionPrice.md) |  | 
**Tiers** | Pointer to [**[]SubscriptionTier**](SubscriptionTier.md) |  | [optional] 
**ProtocolVersion** | **int32** |  | 
**Token** | [**SubscriptionToken**](SubscriptionToken.md) |  | 
**ExpiresAt** | **time.Time** |  | 
**SubscribedAt** | **time.Time** |  | 
**Tier** | [**SubscriptionTier**](SubscriptionTier.md) |  | 
**Creator** | [**User**](User.md) |  | 

## Methods

### NewSubscribedTo

`func NewSubscribedTo(object string, contractAddress string, chain int32, metadata SubscriptionMetadata, ownerAddress string, price SubscriptionPrice, protocolVersion int32, token SubscriptionToken, expiresAt time.Time, subscribedAt time.Time, tier SubscriptionTier, creator User, ) *SubscribedTo`

NewSubscribedTo instantiates a new SubscribedTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribedToWithDefaults

`func NewSubscribedToWithDefaults() *SubscribedTo`

NewSubscribedToWithDefaults instantiates a new SubscribedTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *SubscribedTo) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscribedTo) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscribedTo) SetObject(v string)`

SetObject sets Object field to given value.


### GetProviderName

`func (o *SubscribedTo) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *SubscribedTo) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *SubscribedTo) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *SubscribedTo) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetContractAddress

`func (o *SubscribedTo) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *SubscribedTo) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *SubscribedTo) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetChain

`func (o *SubscribedTo) GetChain() int32`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *SubscribedTo) GetChainOk() (*int32, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *SubscribedTo) SetChain(v int32)`

SetChain sets Chain field to given value.


### GetMetadata

`func (o *SubscribedTo) GetMetadata() SubscriptionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscribedTo) GetMetadataOk() (*SubscriptionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscribedTo) SetMetadata(v SubscriptionMetadata)`

SetMetadata sets Metadata field to given value.


### GetOwnerAddress

`func (o *SubscribedTo) GetOwnerAddress() string`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *SubscribedTo) GetOwnerAddressOk() (*string, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *SubscribedTo) SetOwnerAddress(v string)`

SetOwnerAddress sets OwnerAddress field to given value.


### GetPrice

`func (o *SubscribedTo) GetPrice() SubscriptionPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SubscribedTo) GetPriceOk() (*SubscriptionPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SubscribedTo) SetPrice(v SubscriptionPrice)`

SetPrice sets Price field to given value.


### GetTiers

`func (o *SubscribedTo) GetTiers() []SubscriptionTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *SubscribedTo) GetTiersOk() (*[]SubscriptionTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *SubscribedTo) SetTiers(v []SubscriptionTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *SubscribedTo) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *SubscribedTo) GetProtocolVersion() int32`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *SubscribedTo) GetProtocolVersionOk() (*int32, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *SubscribedTo) SetProtocolVersion(v int32)`

SetProtocolVersion sets ProtocolVersion field to given value.


### GetToken

`func (o *SubscribedTo) GetToken() SubscriptionToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SubscribedTo) GetTokenOk() (*SubscriptionToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SubscribedTo) SetToken(v SubscriptionToken)`

SetToken sets Token field to given value.


### GetExpiresAt

`func (o *SubscribedTo) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SubscribedTo) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SubscribedTo) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetSubscribedAt

`func (o *SubscribedTo) GetSubscribedAt() time.Time`

GetSubscribedAt returns the SubscribedAt field if non-nil, zero value otherwise.

### GetSubscribedAtOk

`func (o *SubscribedTo) GetSubscribedAtOk() (*time.Time, bool)`

GetSubscribedAtOk returns a tuple with the SubscribedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedAt

`func (o *SubscribedTo) SetSubscribedAt(v time.Time)`

SetSubscribedAt sets SubscribedAt field to given value.


### GetTier

`func (o *SubscribedTo) GetTier() SubscriptionTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SubscribedTo) GetTierOk() (*SubscriptionTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SubscribedTo) SetTier(v SubscriptionTier)`

SetTier sets Tier field to given value.


### GetCreator

`func (o *SubscribedTo) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SubscribedTo) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SubscribedTo) SetCreator(v User)`

SetCreator sets Creator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


