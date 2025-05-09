# FetchVerificationsByFid200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]Verification**](Verification.md) |  | 
**NextPageToken** | **string** |  | 

## Methods

### NewFetchVerificationsByFid200Response

`func NewFetchVerificationsByFid200Response(messages []Verification, nextPageToken string, ) *FetchVerificationsByFid200Response`

NewFetchVerificationsByFid200Response instantiates a new FetchVerificationsByFid200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchVerificationsByFid200ResponseWithDefaults

`func NewFetchVerificationsByFid200ResponseWithDefaults() *FetchVerificationsByFid200Response`

NewFetchVerificationsByFid200ResponseWithDefaults instantiates a new FetchVerificationsByFid200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *FetchVerificationsByFid200Response) GetMessages() []Verification`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *FetchVerificationsByFid200Response) GetMessagesOk() (*[]Verification, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *FetchVerificationsByFid200Response) SetMessages(v []Verification)`

SetMessages sets Messages field to given value.


### GetNextPageToken

`func (o *FetchVerificationsByFid200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FetchVerificationsByFid200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FetchVerificationsByFid200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


