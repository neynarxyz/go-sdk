# BanListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bans** | [**[]BanRecord**](BanRecord.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewBanListResponse

`func NewBanListResponse(bans []BanRecord, next NextCursor, ) *BanListResponse`

NewBanListResponse instantiates a new BanListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBanListResponseWithDefaults

`func NewBanListResponseWithDefaults() *BanListResponse`

NewBanListResponseWithDefaults instantiates a new BanListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBans

`func (o *BanListResponse) GetBans() []BanRecord`

GetBans returns the Bans field if non-nil, zero value otherwise.

### GetBansOk

`func (o *BanListResponse) GetBansOk() (*[]BanRecord, bool)`

GetBansOk returns a tuple with the Bans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBans

`func (o *BanListResponse) SetBans(v []BanRecord)`

SetBans sets Bans field to given value.


### GetNext

`func (o *BanListResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *BanListResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *BanListResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


