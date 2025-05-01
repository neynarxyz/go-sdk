# UsersActiveChannelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]Channel**](Channel.md) |  | [optional] 
**Next** | Pointer to [**NextCursor**](NextCursor.md) |  | [optional] 

## Methods

### NewUsersActiveChannelsResponse

`func NewUsersActiveChannelsResponse() *UsersActiveChannelsResponse`

NewUsersActiveChannelsResponse instantiates a new UsersActiveChannelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersActiveChannelsResponseWithDefaults

`func NewUsersActiveChannelsResponseWithDefaults() *UsersActiveChannelsResponse`

NewUsersActiveChannelsResponseWithDefaults instantiates a new UsersActiveChannelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *UsersActiveChannelsResponse) GetChannels() []Channel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *UsersActiveChannelsResponse) GetChannelsOk() (*[]Channel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *UsersActiveChannelsResponse) SetChannels(v []Channel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *UsersActiveChannelsResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetNext

`func (o *UsersActiveChannelsResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *UsersActiveChannelsResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *UsersActiveChannelsResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.

### HasNext

`func (o *UsersActiveChannelsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


