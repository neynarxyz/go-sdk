# StorageAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**UserDehydrated**](UserDehydrated.md) |  | [optional] 
**Units** | Pointer to **int32** |  | [optional] 
**Expiry** | Pointer to **time.Time** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStorageAllocation

`func NewStorageAllocation() *StorageAllocation`

NewStorageAllocation instantiates a new StorageAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageAllocationWithDefaults

`func NewStorageAllocationWithDefaults() *StorageAllocation`

NewStorageAllocationWithDefaults instantiates a new StorageAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *StorageAllocation) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *StorageAllocation) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *StorageAllocation) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *StorageAllocation) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUser

`func (o *StorageAllocation) GetUser() UserDehydrated`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *StorageAllocation) GetUserOk() (*UserDehydrated, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *StorageAllocation) SetUser(v UserDehydrated)`

SetUser sets User field to given value.

### HasUser

`func (o *StorageAllocation) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUnits

`func (o *StorageAllocation) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *StorageAllocation) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *StorageAllocation) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *StorageAllocation) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetExpiry

`func (o *StorageAllocation) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *StorageAllocation) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *StorageAllocation) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *StorageAllocation) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetTimestamp

`func (o *StorageAllocation) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StorageAllocation) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StorageAllocation) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *StorageAllocation) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


