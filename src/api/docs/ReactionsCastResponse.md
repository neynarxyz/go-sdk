# ReactionsCastResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reactions** | [**[]ReactionForCast**](ReactionForCast.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewReactionsCastResponse

`func NewReactionsCastResponse(reactions []ReactionForCast, next NextCursor, ) *ReactionsCastResponse`

NewReactionsCastResponse instantiates a new ReactionsCastResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionsCastResponseWithDefaults

`func NewReactionsCastResponseWithDefaults() *ReactionsCastResponse`

NewReactionsCastResponseWithDefaults instantiates a new ReactionsCastResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReactions

`func (o *ReactionsCastResponse) GetReactions() []ReactionForCast`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *ReactionsCastResponse) GetReactionsOk() (*[]ReactionForCast, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *ReactionsCastResponse) SetReactions(v []ReactionForCast)`

SetReactions sets Reactions field to given value.


### GetNext

`func (o *ReactionsCastResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ReactionsCastResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ReactionsCastResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


