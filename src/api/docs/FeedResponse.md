# FeedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Casts** | [**[]CastWithInteractions**](CastWithInteractions.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewFeedResponse

`func NewFeedResponse(casts []CastWithInteractions, next NextCursor, ) *FeedResponse`

NewFeedResponse instantiates a new FeedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedResponseWithDefaults

`func NewFeedResponseWithDefaults() *FeedResponse`

NewFeedResponseWithDefaults instantiates a new FeedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCasts

`func (o *FeedResponse) GetCasts() []CastWithInteractions`

GetCasts returns the Casts field if non-nil, zero value otherwise.

### GetCastsOk

`func (o *FeedResponse) GetCastsOk() (*[]CastWithInteractions, bool)`

GetCastsOk returns a tuple with the Casts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasts

`func (o *FeedResponse) SetCasts(v []CastWithInteractions)`

SetCasts sets Casts field to given value.


### GetNext

`func (o *FeedResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FeedResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FeedResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


