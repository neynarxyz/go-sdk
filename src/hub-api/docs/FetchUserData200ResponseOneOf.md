# FetchUserData200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | **string** | Base64-encoded pagination token for fetching the next page of results. An empty value indicates there are no more pages to return. Used in conjunction with the pageSize parameter to implement pagination across large result sets. | 
**Messages** | [**[]UserDataAdd**](UserDataAdd.md) |  | 

## Methods

### NewFetchUserData200ResponseOneOf

`func NewFetchUserData200ResponseOneOf(nextPageToken string, messages []UserDataAdd, ) *FetchUserData200ResponseOneOf`

NewFetchUserData200ResponseOneOf instantiates a new FetchUserData200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchUserData200ResponseOneOfWithDefaults

`func NewFetchUserData200ResponseOneOfWithDefaults() *FetchUserData200ResponseOneOf`

NewFetchUserData200ResponseOneOfWithDefaults instantiates a new FetchUserData200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *FetchUserData200ResponseOneOf) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FetchUserData200ResponseOneOf) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FetchUserData200ResponseOneOf) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.


### GetMessages

`func (o *FetchUserData200ResponseOneOf) GetMessages() []UserDataAdd`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *FetchUserData200ResponseOneOf) GetMessagesOk() (*[]UserDataAdd, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *FetchUserData200ResponseOneOf) SetMessages(v []UserDataAdd)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


