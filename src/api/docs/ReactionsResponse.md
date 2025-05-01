# ReactionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reactions** | [**[]ReactionWithCastInfo**](ReactionWithCastInfo.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewReactionsResponse

`func NewReactionsResponse(reactions []ReactionWithCastInfo, next NextCursor, ) *ReactionsResponse`

NewReactionsResponse instantiates a new ReactionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionsResponseWithDefaults

`func NewReactionsResponseWithDefaults() *ReactionsResponse`

NewReactionsResponseWithDefaults instantiates a new ReactionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReactions

`func (o *ReactionsResponse) GetReactions() []ReactionWithCastInfo`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *ReactionsResponse) GetReactionsOk() (*[]ReactionWithCastInfo, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *ReactionsResponse) SetReactions(v []ReactionWithCastInfo)`

SetReactions sets Reactions field to given value.


### GetNext

`func (o *ReactionsResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ReactionsResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ReactionsResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


