# BlockRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Blocked** | Pointer to [**User**](User.md) |  | [optional] 
**Blocker** | Pointer to [**User**](User.md) |  | [optional] 
**BlockedAt** | **time.Time** |  | 

## Methods

### NewBlockRecord

`func NewBlockRecord(object string, blockedAt time.Time, ) *BlockRecord`

NewBlockRecord instantiates a new BlockRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockRecordWithDefaults

`func NewBlockRecordWithDefaults() *BlockRecord`

NewBlockRecordWithDefaults instantiates a new BlockRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *BlockRecord) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BlockRecord) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BlockRecord) SetObject(v string)`

SetObject sets Object field to given value.


### GetBlocked

`func (o *BlockRecord) GetBlocked() User`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *BlockRecord) GetBlockedOk() (*User, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *BlockRecord) SetBlocked(v User)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *BlockRecord) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBlocker

`func (o *BlockRecord) GetBlocker() User`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *BlockRecord) GetBlockerOk() (*User, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *BlockRecord) SetBlocker(v User)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *BlockRecord) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetBlockedAt

`func (o *BlockRecord) GetBlockedAt() time.Time`

GetBlockedAt returns the BlockedAt field if non-nil, zero value otherwise.

### GetBlockedAtOk

`func (o *BlockRecord) GetBlockedAtOk() (*time.Time, bool)`

GetBlockedAtOk returns a tuple with the BlockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedAt

`func (o *BlockRecord) SetBlockedAt(v time.Time)`

SetBlockedAt sets BlockedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


