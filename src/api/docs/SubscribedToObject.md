# SubscribedToObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**ProviderName** | **string** |  | 
**ContractAddress** | Pointer to **string** |  | [optional] 
**ProtocolVersion** | Pointer to **int32** |  | [optional] 
**Chain** | Pointer to **int32** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**SubscribedAt** | Pointer to **time.Time** |  | [optional] 
**TierId** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscribedToObject

`func NewSubscribedToObject(object string, providerName string, ) *SubscribedToObject`

NewSubscribedToObject instantiates a new SubscribedToObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribedToObjectWithDefaults

`func NewSubscribedToObjectWithDefaults() *SubscribedToObject`

NewSubscribedToObjectWithDefaults instantiates a new SubscribedToObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *SubscribedToObject) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscribedToObject) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscribedToObject) SetObject(v string)`

SetObject sets Object field to given value.


### GetProviderName

`func (o *SubscribedToObject) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *SubscribedToObject) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *SubscribedToObject) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetContractAddress

`func (o *SubscribedToObject) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *SubscribedToObject) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *SubscribedToObject) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *SubscribedToObject) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *SubscribedToObject) GetProtocolVersion() int32`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *SubscribedToObject) GetProtocolVersionOk() (*int32, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *SubscribedToObject) SetProtocolVersion(v int32)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *SubscribedToObject) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### GetChain

`func (o *SubscribedToObject) GetChain() int32`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *SubscribedToObject) GetChainOk() (*int32, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *SubscribedToObject) SetChain(v int32)`

SetChain sets Chain field to given value.

### HasChain

`func (o *SubscribedToObject) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetExpiresAt

`func (o *SubscribedToObject) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SubscribedToObject) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SubscribedToObject) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SubscribedToObject) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSubscribedAt

`func (o *SubscribedToObject) GetSubscribedAt() time.Time`

GetSubscribedAt returns the SubscribedAt field if non-nil, zero value otherwise.

### GetSubscribedAtOk

`func (o *SubscribedToObject) GetSubscribedAtOk() (*time.Time, bool)`

GetSubscribedAtOk returns a tuple with the SubscribedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedAt

`func (o *SubscribedToObject) SetSubscribedAt(v time.Time)`

SetSubscribedAt sets SubscribedAt field to given value.

### HasSubscribedAt

`func (o *SubscribedToObject) HasSubscribedAt() bool`

HasSubscribedAt returns a boolean if a field has been set.

### GetTierId

`func (o *SubscribedToObject) GetTierId() string`

GetTierId returns the TierId field if non-nil, zero value otherwise.

### GetTierIdOk

`func (o *SubscribedToObject) GetTierIdOk() (*string, bool)`

GetTierIdOk returns a tuple with the TierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierId

`func (o *SubscribedToObject) SetTierId(v string)`

SetTierId sets TierId field to given value.

### HasTierId

`func (o *SubscribedToObject) HasTierId() bool`

HasTierId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


