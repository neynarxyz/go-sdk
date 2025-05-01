# StorageLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreType** | [**StoreType**](StoreType.md) |  | [default to CASTS]
**Limit** | **int32** |  | 

## Methods

### NewStorageLimit

`func NewStorageLimit(storeType StoreType, limit int32, ) *StorageLimit`

NewStorageLimit instantiates a new StorageLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageLimitWithDefaults

`func NewStorageLimitWithDefaults() *StorageLimit`

NewStorageLimitWithDefaults instantiates a new StorageLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreType

`func (o *StorageLimit) GetStoreType() StoreType`

GetStoreType returns the StoreType field if non-nil, zero value otherwise.

### GetStoreTypeOk

`func (o *StorageLimit) GetStoreTypeOk() (*StoreType, bool)`

GetStoreTypeOk returns a tuple with the StoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreType

`func (o *StorageLimit) SetStoreType(v StoreType)`

SetStoreType sets StoreType field to given value.


### GetLimit

`func (o *StorageLimit) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *StorageLimit) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *StorageLimit) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


