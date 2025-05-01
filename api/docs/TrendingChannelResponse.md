# TrendingChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | [**[]ChannelActivity**](ChannelActivity.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewTrendingChannelResponse

`func NewTrendingChannelResponse(channels []ChannelActivity, next NextCursor, ) *TrendingChannelResponse`

NewTrendingChannelResponse instantiates a new TrendingChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendingChannelResponseWithDefaults

`func NewTrendingChannelResponseWithDefaults() *TrendingChannelResponse`

NewTrendingChannelResponseWithDefaults instantiates a new TrendingChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *TrendingChannelResponse) GetChannels() []ChannelActivity`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *TrendingChannelResponse) GetChannelsOk() (*[]ChannelActivity, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *TrendingChannelResponse) SetChannels(v []ChannelActivity)`

SetChannels sets Channels field to given value.


### GetNext

`func (o *TrendingChannelResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *TrendingChannelResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *TrendingChannelResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


