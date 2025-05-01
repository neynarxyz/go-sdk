# FidsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | **string** | Base64-encoded pagination token for fetching the next page of results. An empty value indicates there are no more pages to return. Used in conjunction with the pageSize parameter to implement pagination across large result sets. | 
**Fids** | **[]int32** |  | 

## Methods

### NewFidsResponse

`func NewFidsResponse(nextPageToken string, fids []int32, ) *FidsResponse`

NewFidsResponse instantiates a new FidsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFidsResponseWithDefaults

`func NewFidsResponseWithDefaults() *FidsResponse`

NewFidsResponseWithDefaults instantiates a new FidsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *FidsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FidsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FidsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.


### GetFids

`func (o *FidsResponse) GetFids() []int32`

GetFids returns the Fids field if non-nil, zero value otherwise.

### GetFidsOk

`func (o *FidsResponse) GetFidsOk() (*[]int32, bool)`

GetFidsOk returns a tuple with the Fids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFids

`func (o *FidsResponse) SetFids(v []int32)`

SetFids sets Fids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


