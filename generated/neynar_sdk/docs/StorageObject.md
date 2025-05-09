# StorageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Used** | Pointer to **int32** |  | [optional] 
**Capacity** | Pointer to **int32** |  | [optional] 

## Methods

### NewStorageObject

`func NewStorageObject() *StorageObject`

NewStorageObject instantiates a new StorageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageObjectWithDefaults

`func NewStorageObjectWithDefaults() *StorageObject`

NewStorageObjectWithDefaults instantiates a new StorageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *StorageObject) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *StorageObject) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *StorageObject) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *StorageObject) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUsed

`func (o *StorageObject) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *StorageObject) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *StorageObject) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *StorageObject) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetCapacity

`func (o *StorageObject) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *StorageObject) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *StorageObject) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *StorageObject) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


