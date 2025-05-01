# MuteRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Muted** | [**User**](User.md) |  | 
**MutedAt** | **time.Time** |  | 

## Methods

### NewMuteRecord

`func NewMuteRecord(object string, muted User, mutedAt time.Time, ) *MuteRecord`

NewMuteRecord instantiates a new MuteRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMuteRecordWithDefaults

`func NewMuteRecordWithDefaults() *MuteRecord`

NewMuteRecordWithDefaults instantiates a new MuteRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *MuteRecord) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *MuteRecord) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *MuteRecord) SetObject(v string)`

SetObject sets Object field to given value.


### GetMuted

`func (o *MuteRecord) GetMuted() User`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *MuteRecord) GetMutedOk() (*User, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *MuteRecord) SetMuted(v User)`

SetMuted sets Muted field to given value.


### GetMutedAt

`func (o *MuteRecord) GetMutedAt() time.Time`

GetMutedAt returns the MutedAt field if non-nil, zero value otherwise.

### GetMutedAtOk

`func (o *MuteRecord) GetMutedAtOk() (*time.Time, bool)`

GetMutedAtOk returns a tuple with the MutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedAt

`func (o *MuteRecord) SetMutedAt(v time.Time)`

SetMutedAt sets MutedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


