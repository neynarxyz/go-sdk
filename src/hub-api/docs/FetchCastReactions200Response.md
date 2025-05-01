# FetchCastReactions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | **string** | Base64-encoded pagination token for fetching the next page of results. An empty value indicates there are no more pages to return. Used in conjunction with the pageSize parameter to implement pagination across large result sets. | 
**Messages** | [**[]Reaction**](Reaction.md) |  | 

## Methods

### NewFetchCastReactions200Response

`func NewFetchCastReactions200Response(nextPageToken string, messages []Reaction, ) *FetchCastReactions200Response`

NewFetchCastReactions200Response instantiates a new FetchCastReactions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchCastReactions200ResponseWithDefaults

`func NewFetchCastReactions200ResponseWithDefaults() *FetchCastReactions200Response`

NewFetchCastReactions200ResponseWithDefaults instantiates a new FetchCastReactions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *FetchCastReactions200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FetchCastReactions200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FetchCastReactions200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.


### GetMessages

`func (o *FetchCastReactions200Response) GetMessages() []Reaction`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *FetchCastReactions200Response) GetMessagesOk() (*[]Reaction, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *FetchCastReactions200Response) SetMessages(v []Reaction)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


