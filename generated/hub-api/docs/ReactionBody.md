# ReactionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ReactionType**](ReactionType.md) |  | [default to REACTIONTYPE_LIKE]
**TargetCastId** | Pointer to [**CastId**](CastId.md) | The unique identifier of the cast being reacted to. Must be a valid CastId containing the creator&#39;s FID and the cast&#39;s hash. | [optional] 
**TargetUrl** | Pointer to **string** | The URL being reacted to. Used when reacting to external content or when the cast ID is not available. | [optional] 

## Methods

### NewReactionBody

`func NewReactionBody(type_ ReactionType, ) *ReactionBody`

NewReactionBody instantiates a new ReactionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionBodyWithDefaults

`func NewReactionBodyWithDefaults() *ReactionBody`

NewReactionBodyWithDefaults instantiates a new ReactionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReactionBody) GetType() ReactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReactionBody) GetTypeOk() (*ReactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReactionBody) SetType(v ReactionType)`

SetType sets Type field to given value.


### GetTargetCastId

`func (o *ReactionBody) GetTargetCastId() CastId`

GetTargetCastId returns the TargetCastId field if non-nil, zero value otherwise.

### GetTargetCastIdOk

`func (o *ReactionBody) GetTargetCastIdOk() (*CastId, bool)`

GetTargetCastIdOk returns a tuple with the TargetCastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCastId

`func (o *ReactionBody) SetTargetCastId(v CastId)`

SetTargetCastId sets TargetCastId field to given value.

### HasTargetCastId

`func (o *ReactionBody) HasTargetCastId() bool`

HasTargetCastId returns a boolean if a field has been set.

### GetTargetUrl

`func (o *ReactionBody) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *ReactionBody) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *ReactionBody) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *ReactionBody) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


