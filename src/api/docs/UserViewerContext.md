# UserViewerContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Following** | **bool** | Indicates if the viewer is following the user. | 
**FollowedBy** | **bool** | Indicates if the viewer is followed by the user. | 
**Blocking** | **bool** | Indicates if the viewer is blocking the user. | 
**BlockedBy** | **bool** | Indicates if the viewer is blocked by the user. | 

## Methods

### NewUserViewerContext

`func NewUserViewerContext(following bool, followedBy bool, blocking bool, blockedBy bool, ) *UserViewerContext`

NewUserViewerContext instantiates a new UserViewerContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserViewerContextWithDefaults

`func NewUserViewerContextWithDefaults() *UserViewerContext`

NewUserViewerContextWithDefaults instantiates a new UserViewerContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowing

`func (o *UserViewerContext) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *UserViewerContext) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *UserViewerContext) SetFollowing(v bool)`

SetFollowing sets Following field to given value.


### GetFollowedBy

`func (o *UserViewerContext) GetFollowedBy() bool`

GetFollowedBy returns the FollowedBy field if non-nil, zero value otherwise.

### GetFollowedByOk

`func (o *UserViewerContext) GetFollowedByOk() (*bool, bool)`

GetFollowedByOk returns a tuple with the FollowedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowedBy

`func (o *UserViewerContext) SetFollowedBy(v bool)`

SetFollowedBy sets FollowedBy field to given value.


### GetBlocking

`func (o *UserViewerContext) GetBlocking() bool`

GetBlocking returns the Blocking field if non-nil, zero value otherwise.

### GetBlockingOk

`func (o *UserViewerContext) GetBlockingOk() (*bool, bool)`

GetBlockingOk returns a tuple with the Blocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocking

`func (o *UserViewerContext) SetBlocking(v bool)`

SetBlocking sets Blocking field to given value.


### GetBlockedBy

`func (o *UserViewerContext) GetBlockedBy() bool`

GetBlockedBy returns the BlockedBy field if non-nil, zero value otherwise.

### GetBlockedByOk

`func (o *UserViewerContext) GetBlockedByOk() (*bool, bool)`

GetBlockedByOk returns a tuple with the BlockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBy

`func (o *UserViewerContext) SetBlockedBy(v bool)`

SetBlockedBy sets BlockedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


