# BanRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Banned** | Pointer to [**User**](User.md) |  | [optional] 
**BannedAt** | **time.Time** |  | 

## Methods

### NewBanRecord

`func NewBanRecord(object string, bannedAt time.Time, ) *BanRecord`

NewBanRecord instantiates a new BanRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBanRecordWithDefaults

`func NewBanRecordWithDefaults() *BanRecord`

NewBanRecordWithDefaults instantiates a new BanRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *BanRecord) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BanRecord) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BanRecord) SetObject(v string)`

SetObject sets Object field to given value.


### GetBanned

`func (o *BanRecord) GetBanned() User`

GetBanned returns the Banned field if non-nil, zero value otherwise.

### GetBannedOk

`func (o *BanRecord) GetBannedOk() (*User, bool)`

GetBannedOk returns a tuple with the Banned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanned

`func (o *BanRecord) SetBanned(v User)`

SetBanned sets Banned field to given value.

### HasBanned

`func (o *BanRecord) HasBanned() bool`

HasBanned returns a boolean if a field has been set.

### GetBannedAt

`func (o *BanRecord) GetBannedAt() time.Time`

GetBannedAt returns the BannedAt field if non-nil, zero value otherwise.

### GetBannedAtOk

`func (o *BanRecord) GetBannedAtOk() (*time.Time, bool)`

GetBannedAtOk returns a tuple with the BannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedAt

`func (o *BanRecord) SetBannedAt(v time.Time)`

SetBannedAt sets BannedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


