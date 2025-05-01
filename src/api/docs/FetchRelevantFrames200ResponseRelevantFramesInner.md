# FetchRelevantFrames200ResponseRelevantFramesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frames** | [**[]FrameV2WithFullAuthor**](FrameV2WithFullAuthor.md) | Array of FrameV2 objects | 
**TopRelevantUsers** | [**[]User**](User.md) | Array of the most relevant users | 
**RemainingRelevantUsers** | [**[]UserDehydrated**](UserDehydrated.md) | Array of remaining relevant users in dehydrated form | 

## Methods

### NewFetchRelevantFrames200ResponseRelevantFramesInner

`func NewFetchRelevantFrames200ResponseRelevantFramesInner(frames []FrameV2WithFullAuthor, topRelevantUsers []User, remainingRelevantUsers []UserDehydrated, ) *FetchRelevantFrames200ResponseRelevantFramesInner`

NewFetchRelevantFrames200ResponseRelevantFramesInner instantiates a new FetchRelevantFrames200ResponseRelevantFramesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchRelevantFrames200ResponseRelevantFramesInnerWithDefaults

`func NewFetchRelevantFrames200ResponseRelevantFramesInnerWithDefaults() *FetchRelevantFrames200ResponseRelevantFramesInner`

NewFetchRelevantFrames200ResponseRelevantFramesInnerWithDefaults instantiates a new FetchRelevantFrames200ResponseRelevantFramesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrames

`func (o *FetchRelevantFrames200ResponseRelevantFramesInner) GetFrames() []FrameV2WithFullAuthor`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *FetchRelevantFrames200ResponseRelevantFramesInner) GetFramesOk() (*[]FrameV2WithFullAuthor, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *FetchRelevantFrames200ResponseRelevantFramesInner) SetFrames(v []FrameV2WithFullAuthor)`

SetFrames sets Frames field to given value.


### GetTopRelevantUsers

`func (o *FetchRelevantFrames200ResponseRelevantFramesInner) GetTopRelevantUsers() []User`

GetTopRelevantUsers returns the TopRelevantUsers field if non-nil, zero value otherwise.

### GetTopRelevantUsersOk

`func (o *FetchRelevantFrames200ResponseRelevantFramesInner) GetTopRelevantUsersOk() (*[]User, bool)`

GetTopRelevantUsersOk returns a tuple with the TopRelevantUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopRelevantUsers

`func (o *FetchRelevantFrames200ResponseRelevantFramesInner) SetTopRelevantUsers(v []User)`

SetTopRelevantUsers sets TopRelevantUsers field to given value.


### GetRemainingRelevantUsers

`func (o *FetchRelevantFrames200ResponseRelevantFramesInner) GetRemainingRelevantUsers() []UserDehydrated`

GetRemainingRelevantUsers returns the RemainingRelevantUsers field if non-nil, zero value otherwise.

### GetRemainingRelevantUsersOk

`func (o *FetchRelevantFrames200ResponseRelevantFramesInner) GetRemainingRelevantUsersOk() (*[]UserDehydrated, bool)`

GetRemainingRelevantUsersOk returns a tuple with the RemainingRelevantUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingRelevantUsers

`func (o *FetchRelevantFrames200ResponseRelevantFramesInner) SetRemainingRelevantUsers(v []UserDehydrated)`

SetRemainingRelevantUsers sets RemainingRelevantUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


