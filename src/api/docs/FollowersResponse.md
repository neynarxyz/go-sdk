# FollowersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]Follower**](Follower.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewFollowersResponse

`func NewFollowersResponse(users []Follower, next NextCursor, ) *FollowersResponse`

NewFollowersResponse instantiates a new FollowersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowersResponseWithDefaults

`func NewFollowersResponseWithDefaults() *FollowersResponse`

NewFollowersResponseWithDefaults instantiates a new FollowersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *FollowersResponse) GetUsers() []Follower`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FollowersResponse) GetUsersOk() (*[]Follower, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FollowersResponse) SetUsers(v []Follower)`

SetUsers sets Users field to given value.


### GetNext

`func (o *FollowersResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FollowersResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FollowersResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


