# StorageUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**UserDehydrated**](UserDehydrated.md) |  | [optional] 
**Casts** | Pointer to [**StorageObject**](StorageObject.md) |  | [optional] 
**Reactions** | Pointer to [**StorageObject**](StorageObject.md) |  | [optional] 
**Links** | Pointer to [**StorageObject**](StorageObject.md) |  | [optional] 
**VerifiedAddresses** | Pointer to [**StorageObject**](StorageObject.md) |  | [optional] 
**UsernameProofs** | Pointer to [**StorageObject**](StorageObject.md) |  | [optional] 
**Signers** | Pointer to [**StorageObject**](StorageObject.md) |  | [optional] 
**TotalActiveUnits** | Pointer to **int32** |  | [optional] 

## Methods

### NewStorageUsageResponse

`func NewStorageUsageResponse() *StorageUsageResponse`

NewStorageUsageResponse instantiates a new StorageUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageUsageResponseWithDefaults

`func NewStorageUsageResponseWithDefaults() *StorageUsageResponse`

NewStorageUsageResponseWithDefaults instantiates a new StorageUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *StorageUsageResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *StorageUsageResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *StorageUsageResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *StorageUsageResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUser

`func (o *StorageUsageResponse) GetUser() UserDehydrated`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *StorageUsageResponse) GetUserOk() (*UserDehydrated, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *StorageUsageResponse) SetUser(v UserDehydrated)`

SetUser sets User field to given value.

### HasUser

`func (o *StorageUsageResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCasts

`func (o *StorageUsageResponse) GetCasts() StorageObject`

GetCasts returns the Casts field if non-nil, zero value otherwise.

### GetCastsOk

`func (o *StorageUsageResponse) GetCastsOk() (*StorageObject, bool)`

GetCastsOk returns a tuple with the Casts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasts

`func (o *StorageUsageResponse) SetCasts(v StorageObject)`

SetCasts sets Casts field to given value.

### HasCasts

`func (o *StorageUsageResponse) HasCasts() bool`

HasCasts returns a boolean if a field has been set.

### GetReactions

`func (o *StorageUsageResponse) GetReactions() StorageObject`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *StorageUsageResponse) GetReactionsOk() (*StorageObject, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *StorageUsageResponse) SetReactions(v StorageObject)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *StorageUsageResponse) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetLinks

`func (o *StorageUsageResponse) GetLinks() StorageObject`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StorageUsageResponse) GetLinksOk() (*StorageObject, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StorageUsageResponse) SetLinks(v StorageObject)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StorageUsageResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetVerifiedAddresses

`func (o *StorageUsageResponse) GetVerifiedAddresses() StorageObject`

GetVerifiedAddresses returns the VerifiedAddresses field if non-nil, zero value otherwise.

### GetVerifiedAddressesOk

`func (o *StorageUsageResponse) GetVerifiedAddressesOk() (*StorageObject, bool)`

GetVerifiedAddressesOk returns a tuple with the VerifiedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAddresses

`func (o *StorageUsageResponse) SetVerifiedAddresses(v StorageObject)`

SetVerifiedAddresses sets VerifiedAddresses field to given value.

### HasVerifiedAddresses

`func (o *StorageUsageResponse) HasVerifiedAddresses() bool`

HasVerifiedAddresses returns a boolean if a field has been set.

### GetUsernameProofs

`func (o *StorageUsageResponse) GetUsernameProofs() StorageObject`

GetUsernameProofs returns the UsernameProofs field if non-nil, zero value otherwise.

### GetUsernameProofsOk

`func (o *StorageUsageResponse) GetUsernameProofsOk() (*StorageObject, bool)`

GetUsernameProofsOk returns a tuple with the UsernameProofs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameProofs

`func (o *StorageUsageResponse) SetUsernameProofs(v StorageObject)`

SetUsernameProofs sets UsernameProofs field to given value.

### HasUsernameProofs

`func (o *StorageUsageResponse) HasUsernameProofs() bool`

HasUsernameProofs returns a boolean if a field has been set.

### GetSigners

`func (o *StorageUsageResponse) GetSigners() StorageObject`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *StorageUsageResponse) GetSignersOk() (*StorageObject, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *StorageUsageResponse) SetSigners(v StorageObject)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *StorageUsageResponse) HasSigners() bool`

HasSigners returns a boolean if a field has been set.

### GetTotalActiveUnits

`func (o *StorageUsageResponse) GetTotalActiveUnits() int32`

GetTotalActiveUnits returns the TotalActiveUnits field if non-nil, zero value otherwise.

### GetTotalActiveUnitsOk

`func (o *StorageUsageResponse) GetTotalActiveUnitsOk() (*int32, bool)`

GetTotalActiveUnitsOk returns a tuple with the TotalActiveUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActiveUnits

`func (o *StorageUsageResponse) SetTotalActiveUnits(v int32)`

SetTotalActiveUnits sets TotalActiveUnits field to given value.

### HasTotalActiveUnits

`func (o *StorageUsageResponse) HasTotalActiveUnits() bool`

HasTotalActiveUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


