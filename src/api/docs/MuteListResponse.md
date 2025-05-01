# MuteListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mutes** | [**[]MuteRecord**](MuteRecord.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewMuteListResponse

`func NewMuteListResponse(mutes []MuteRecord, next NextCursor, ) *MuteListResponse`

NewMuteListResponse instantiates a new MuteListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMuteListResponseWithDefaults

`func NewMuteListResponseWithDefaults() *MuteListResponse`

NewMuteListResponseWithDefaults instantiates a new MuteListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMutes

`func (o *MuteListResponse) GetMutes() []MuteRecord`

GetMutes returns the Mutes field if non-nil, zero value otherwise.

### GetMutesOk

`func (o *MuteListResponse) GetMutesOk() (*[]MuteRecord, bool)`

GetMutesOk returns a tuple with the Mutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutes

`func (o *MuteListResponse) SetMutes(v []MuteRecord)`

SetMutes sets Mutes field to given value.


### GetNext

`func (o *MuteListResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *MuteListResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *MuteListResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


