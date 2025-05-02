# FetchUserFollowing200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | **string** | Base64-encoded pagination token for fetching the next page of results. An empty value indicates there are no more pages to return. Used in conjunction with the pageSize parameter to implement pagination across large result sets. | 
**Messages** | [**[]LinkAdd**](LinkAdd.md) |  | 

## Methods

### NewFetchUserFollowing200Response

`func NewFetchUserFollowing200Response(nextPageToken string, messages []LinkAdd, ) *FetchUserFollowing200Response`

NewFetchUserFollowing200Response instantiates a new FetchUserFollowing200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchUserFollowing200ResponseWithDefaults

`func NewFetchUserFollowing200ResponseWithDefaults() *FetchUserFollowing200Response`

NewFetchUserFollowing200ResponseWithDefaults instantiates a new FetchUserFollowing200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *FetchUserFollowing200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FetchUserFollowing200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FetchUserFollowing200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.


### GetMessages

`func (o *FetchUserFollowing200Response) GetMessages() []LinkAdd`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *FetchUserFollowing200Response) GetMessagesOk() (*[]LinkAdd, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *FetchUserFollowing200Response) SetMessages(v []LinkAdd)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


