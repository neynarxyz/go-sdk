# StorageAllocationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalActiveUnits** | Pointer to **int32** |  | [optional] 
**Allocations** | Pointer to [**[]StorageAllocation**](StorageAllocation.md) |  | [optional] 

## Methods

### NewStorageAllocationsResponse

`func NewStorageAllocationsResponse() *StorageAllocationsResponse`

NewStorageAllocationsResponse instantiates a new StorageAllocationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageAllocationsResponseWithDefaults

`func NewStorageAllocationsResponseWithDefaults() *StorageAllocationsResponse`

NewStorageAllocationsResponseWithDefaults instantiates a new StorageAllocationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalActiveUnits

`func (o *StorageAllocationsResponse) GetTotalActiveUnits() int32`

GetTotalActiveUnits returns the TotalActiveUnits field if non-nil, zero value otherwise.

### GetTotalActiveUnitsOk

`func (o *StorageAllocationsResponse) GetTotalActiveUnitsOk() (*int32, bool)`

GetTotalActiveUnitsOk returns a tuple with the TotalActiveUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActiveUnits

`func (o *StorageAllocationsResponse) SetTotalActiveUnits(v int32)`

SetTotalActiveUnits sets TotalActiveUnits field to given value.

### HasTotalActiveUnits

`func (o *StorageAllocationsResponse) HasTotalActiveUnits() bool`

HasTotalActiveUnits returns a boolean if a field has been set.

### GetAllocations

`func (o *StorageAllocationsResponse) GetAllocations() []StorageAllocation`

GetAllocations returns the Allocations field if non-nil, zero value otherwise.

### GetAllocationsOk

`func (o *StorageAllocationsResponse) GetAllocationsOk() (*[]StorageAllocation, bool)`

GetAllocationsOk returns a tuple with the Allocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocations

`func (o *StorageAllocationsResponse) SetAllocations(v []StorageAllocation)`

SetAllocations sets Allocations field to given value.

### HasAllocations

`func (o *StorageAllocationsResponse) HasAllocations() bool`

HasAllocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


