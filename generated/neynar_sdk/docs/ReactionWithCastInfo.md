# ReactionWithCastInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReactionType** | **string** |  | 
**Cast** | [**CastWithInteractions**](CastWithInteractions.md) |  | 
**ReactionTimestamp** | **time.Time** |  | 
**Object** | **string** |  | 
**User** | [**UserDehydrated**](UserDehydrated.md) |  | 

## Methods

### NewReactionWithCastInfo

`func NewReactionWithCastInfo(reactionType string, cast CastWithInteractions, reactionTimestamp time.Time, object string, user UserDehydrated, ) *ReactionWithCastInfo`

NewReactionWithCastInfo instantiates a new ReactionWithCastInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionWithCastInfoWithDefaults

`func NewReactionWithCastInfoWithDefaults() *ReactionWithCastInfo`

NewReactionWithCastInfoWithDefaults instantiates a new ReactionWithCastInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReactionType

`func (o *ReactionWithCastInfo) GetReactionType() string`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *ReactionWithCastInfo) GetReactionTypeOk() (*string, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *ReactionWithCastInfo) SetReactionType(v string)`

SetReactionType sets ReactionType field to given value.


### GetCast

`func (o *ReactionWithCastInfo) GetCast() CastWithInteractions`

GetCast returns the Cast field if non-nil, zero value otherwise.

### GetCastOk

`func (o *ReactionWithCastInfo) GetCastOk() (*CastWithInteractions, bool)`

GetCastOk returns a tuple with the Cast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCast

`func (o *ReactionWithCastInfo) SetCast(v CastWithInteractions)`

SetCast sets Cast field to given value.


### GetReactionTimestamp

`func (o *ReactionWithCastInfo) GetReactionTimestamp() time.Time`

GetReactionTimestamp returns the ReactionTimestamp field if non-nil, zero value otherwise.

### GetReactionTimestampOk

`func (o *ReactionWithCastInfo) GetReactionTimestampOk() (*time.Time, bool)`

GetReactionTimestampOk returns a tuple with the ReactionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionTimestamp

`func (o *ReactionWithCastInfo) SetReactionTimestamp(v time.Time)`

SetReactionTimestamp sets ReactionTimestamp field to given value.


### GetObject

`func (o *ReactionWithCastInfo) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ReactionWithCastInfo) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ReactionWithCastInfo) SetObject(v string)`

SetObject sets Object field to given value.


### GetUser

`func (o *ReactionWithCastInfo) GetUser() UserDehydrated`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ReactionWithCastInfo) GetUserOk() (*UserDehydrated, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ReactionWithCastInfo) SetUser(v UserDehydrated)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


