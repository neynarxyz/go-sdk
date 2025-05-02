# ReactionForCast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReactionType** | **string** |  | 
**ReactionTimestamp** | **time.Time** |  | 
**Object** | **string** |  | 
**User** | [**User**](User.md) |  | 

## Methods

### NewReactionForCast

`func NewReactionForCast(reactionType string, reactionTimestamp time.Time, object string, user User, ) *ReactionForCast`

NewReactionForCast instantiates a new ReactionForCast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionForCastWithDefaults

`func NewReactionForCastWithDefaults() *ReactionForCast`

NewReactionForCastWithDefaults instantiates a new ReactionForCast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReactionType

`func (o *ReactionForCast) GetReactionType() string`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *ReactionForCast) GetReactionTypeOk() (*string, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *ReactionForCast) SetReactionType(v string)`

SetReactionType sets ReactionType field to given value.


### GetReactionTimestamp

`func (o *ReactionForCast) GetReactionTimestamp() time.Time`

GetReactionTimestamp returns the ReactionTimestamp field if non-nil, zero value otherwise.

### GetReactionTimestampOk

`func (o *ReactionForCast) GetReactionTimestampOk() (*time.Time, bool)`

GetReactionTimestampOk returns a tuple with the ReactionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionTimestamp

`func (o *ReactionForCast) SetReactionTimestamp(v time.Time)`

SetReactionTimestamp sets ReactionTimestamp field to given value.


### GetObject

`func (o *ReactionForCast) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ReactionForCast) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ReactionForCast) SetObject(v string)`

SetObject sets Object field to given value.


### GetUser

`func (o *ReactionForCast) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ReactionForCast) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ReactionForCast) SetUser(v User)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


