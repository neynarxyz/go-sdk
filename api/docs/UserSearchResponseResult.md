# UserSearchResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]SearchedUser**](SearchedUser.md) |  | 
**Next** | Pointer to [**NextCursor**](NextCursor.md) |  | [optional] 

## Methods

### NewUserSearchResponseResult

`func NewUserSearchResponseResult(users []SearchedUser, ) *UserSearchResponseResult`

NewUserSearchResponseResult instantiates a new UserSearchResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSearchResponseResultWithDefaults

`func NewUserSearchResponseResultWithDefaults() *UserSearchResponseResult`

NewUserSearchResponseResultWithDefaults instantiates a new UserSearchResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UserSearchResponseResult) GetUsers() []SearchedUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserSearchResponseResult) GetUsersOk() (*[]SearchedUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserSearchResponseResult) SetUsers(v []SearchedUser)`

SetUsers sets Users field to given value.


### GetNext

`func (o *UserSearchResponseResult) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *UserSearchResponseResult) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *UserSearchResponseResult) SetNext(v NextCursor)`

SetNext sets Next field to given value.

### HasNext

`func (o *UserSearchResponseResult) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


