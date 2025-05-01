# FetchEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageEventId** | **int32** |  | 
**Events** | [**[]HubEvent**](HubEvent.md) |  | 

## Methods

### NewFetchEvents200Response

`func NewFetchEvents200Response(nextPageEventId int32, events []HubEvent, ) *FetchEvents200Response`

NewFetchEvents200Response instantiates a new FetchEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchEvents200ResponseWithDefaults

`func NewFetchEvents200ResponseWithDefaults() *FetchEvents200Response`

NewFetchEvents200ResponseWithDefaults instantiates a new FetchEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageEventId

`func (o *FetchEvents200Response) GetNextPageEventId() int32`

GetNextPageEventId returns the NextPageEventId field if non-nil, zero value otherwise.

### GetNextPageEventIdOk

`func (o *FetchEvents200Response) GetNextPageEventIdOk() (*int32, bool)`

GetNextPageEventIdOk returns a tuple with the NextPageEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageEventId

`func (o *FetchEvents200Response) SetNextPageEventId(v int32)`

SetNextPageEventId sets NextPageEventId field to given value.


### GetEvents

`func (o *FetchEvents200Response) GetEvents() []HubEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FetchEvents200Response) GetEventsOk() (*[]HubEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FetchEvents200Response) SetEvents(v []HubEvent)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


