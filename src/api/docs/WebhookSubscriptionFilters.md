# WebhookSubscriptionFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CastCreated** | Pointer to [**WebhookSubscriptionFiltersCast**](WebhookSubscriptionFiltersCast.md) |  | [optional] 
**CastDeleted** | Pointer to [**WebhookSubscriptionFiltersCast**](WebhookSubscriptionFiltersCast.md) |  | [optional] 
**UserCreated** | Pointer to **map[string]interface{}** |  | [optional] 
**UserUpdated** | Pointer to [**WebhookSubscriptionFiltersUserUpdated**](WebhookSubscriptionFiltersUserUpdated.md) |  | [optional] 
**FollowCreated** | Pointer to [**WebhookSubscriptionFiltersFollow**](WebhookSubscriptionFiltersFollow.md) |  | [optional] 
**FollowDeleted** | Pointer to [**WebhookSubscriptionFiltersFollow**](WebhookSubscriptionFiltersFollow.md) |  | [optional] 
**ReactionCreated** | Pointer to [**WebhookSubscriptionFiltersReaction**](WebhookSubscriptionFiltersReaction.md) |  | [optional] 
**ReactionDeleted** | Pointer to [**WebhookSubscriptionFiltersReaction**](WebhookSubscriptionFiltersReaction.md) |  | [optional] 

## Methods

### NewWebhookSubscriptionFilters

`func NewWebhookSubscriptionFilters() *WebhookSubscriptionFilters`

NewWebhookSubscriptionFilters instantiates a new WebhookSubscriptionFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionFiltersWithDefaults

`func NewWebhookSubscriptionFiltersWithDefaults() *WebhookSubscriptionFilters`

NewWebhookSubscriptionFiltersWithDefaults instantiates a new WebhookSubscriptionFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCastCreated

`func (o *WebhookSubscriptionFilters) GetCastCreated() WebhookSubscriptionFiltersCast`

GetCastCreated returns the CastCreated field if non-nil, zero value otherwise.

### GetCastCreatedOk

`func (o *WebhookSubscriptionFilters) GetCastCreatedOk() (*WebhookSubscriptionFiltersCast, bool)`

GetCastCreatedOk returns a tuple with the CastCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastCreated

`func (o *WebhookSubscriptionFilters) SetCastCreated(v WebhookSubscriptionFiltersCast)`

SetCastCreated sets CastCreated field to given value.

### HasCastCreated

`func (o *WebhookSubscriptionFilters) HasCastCreated() bool`

HasCastCreated returns a boolean if a field has been set.

### GetCastDeleted

`func (o *WebhookSubscriptionFilters) GetCastDeleted() WebhookSubscriptionFiltersCast`

GetCastDeleted returns the CastDeleted field if non-nil, zero value otherwise.

### GetCastDeletedOk

`func (o *WebhookSubscriptionFilters) GetCastDeletedOk() (*WebhookSubscriptionFiltersCast, bool)`

GetCastDeletedOk returns a tuple with the CastDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastDeleted

`func (o *WebhookSubscriptionFilters) SetCastDeleted(v WebhookSubscriptionFiltersCast)`

SetCastDeleted sets CastDeleted field to given value.

### HasCastDeleted

`func (o *WebhookSubscriptionFilters) HasCastDeleted() bool`

HasCastDeleted returns a boolean if a field has been set.

### GetUserCreated

`func (o *WebhookSubscriptionFilters) GetUserCreated() map[string]interface{}`

GetUserCreated returns the UserCreated field if non-nil, zero value otherwise.

### GetUserCreatedOk

`func (o *WebhookSubscriptionFilters) GetUserCreatedOk() (*map[string]interface{}, bool)`

GetUserCreatedOk returns a tuple with the UserCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCreated

`func (o *WebhookSubscriptionFilters) SetUserCreated(v map[string]interface{})`

SetUserCreated sets UserCreated field to given value.

### HasUserCreated

`func (o *WebhookSubscriptionFilters) HasUserCreated() bool`

HasUserCreated returns a boolean if a field has been set.

### GetUserUpdated

`func (o *WebhookSubscriptionFilters) GetUserUpdated() WebhookSubscriptionFiltersUserUpdated`

GetUserUpdated returns the UserUpdated field if non-nil, zero value otherwise.

### GetUserUpdatedOk

`func (o *WebhookSubscriptionFilters) GetUserUpdatedOk() (*WebhookSubscriptionFiltersUserUpdated, bool)`

GetUserUpdatedOk returns a tuple with the UserUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUpdated

`func (o *WebhookSubscriptionFilters) SetUserUpdated(v WebhookSubscriptionFiltersUserUpdated)`

SetUserUpdated sets UserUpdated field to given value.

### HasUserUpdated

`func (o *WebhookSubscriptionFilters) HasUserUpdated() bool`

HasUserUpdated returns a boolean if a field has been set.

### GetFollowCreated

`func (o *WebhookSubscriptionFilters) GetFollowCreated() WebhookSubscriptionFiltersFollow`

GetFollowCreated returns the FollowCreated field if non-nil, zero value otherwise.

### GetFollowCreatedOk

`func (o *WebhookSubscriptionFilters) GetFollowCreatedOk() (*WebhookSubscriptionFiltersFollow, bool)`

GetFollowCreatedOk returns a tuple with the FollowCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowCreated

`func (o *WebhookSubscriptionFilters) SetFollowCreated(v WebhookSubscriptionFiltersFollow)`

SetFollowCreated sets FollowCreated field to given value.

### HasFollowCreated

`func (o *WebhookSubscriptionFilters) HasFollowCreated() bool`

HasFollowCreated returns a boolean if a field has been set.

### GetFollowDeleted

`func (o *WebhookSubscriptionFilters) GetFollowDeleted() WebhookSubscriptionFiltersFollow`

GetFollowDeleted returns the FollowDeleted field if non-nil, zero value otherwise.

### GetFollowDeletedOk

`func (o *WebhookSubscriptionFilters) GetFollowDeletedOk() (*WebhookSubscriptionFiltersFollow, bool)`

GetFollowDeletedOk returns a tuple with the FollowDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowDeleted

`func (o *WebhookSubscriptionFilters) SetFollowDeleted(v WebhookSubscriptionFiltersFollow)`

SetFollowDeleted sets FollowDeleted field to given value.

### HasFollowDeleted

`func (o *WebhookSubscriptionFilters) HasFollowDeleted() bool`

HasFollowDeleted returns a boolean if a field has been set.

### GetReactionCreated

`func (o *WebhookSubscriptionFilters) GetReactionCreated() WebhookSubscriptionFiltersReaction`

GetReactionCreated returns the ReactionCreated field if non-nil, zero value otherwise.

### GetReactionCreatedOk

`func (o *WebhookSubscriptionFilters) GetReactionCreatedOk() (*WebhookSubscriptionFiltersReaction, bool)`

GetReactionCreatedOk returns a tuple with the ReactionCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionCreated

`func (o *WebhookSubscriptionFilters) SetReactionCreated(v WebhookSubscriptionFiltersReaction)`

SetReactionCreated sets ReactionCreated field to given value.

### HasReactionCreated

`func (o *WebhookSubscriptionFilters) HasReactionCreated() bool`

HasReactionCreated returns a boolean if a field has been set.

### GetReactionDeleted

`func (o *WebhookSubscriptionFilters) GetReactionDeleted() WebhookSubscriptionFiltersReaction`

GetReactionDeleted returns the ReactionDeleted field if non-nil, zero value otherwise.

### GetReactionDeletedOk

`func (o *WebhookSubscriptionFilters) GetReactionDeletedOk() (*WebhookSubscriptionFiltersReaction, bool)`

GetReactionDeletedOk returns a tuple with the ReactionDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionDeleted

`func (o *WebhookSubscriptionFilters) SetReactionDeleted(v WebhookSubscriptionFiltersReaction)`

SetReactionDeleted sets ReactionDeleted field to given value.

### HasReactionDeleted

`func (o *WebhookSubscriptionFilters) HasReactionDeleted() bool`

HasReactionDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


