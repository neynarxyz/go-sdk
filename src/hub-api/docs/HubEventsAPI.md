# \HubEventsAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchEvents**](HubEventsAPI.md#FetchEvents) | **Get** /v1/events | Page of events
[**LookupEvent**](HubEventsAPI.md#LookupEvent) | **Get** /v1/eventById | Event by ID



## FetchEvents

> FetchEvents200Response FetchEvents(ctx).FromEventId(fromEventId).Execute()

Page of events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/hub"
)

func main() {
	fromEventId := int32(0) // int32 | An optional Hub Id to start getting events from. This is also returned from the API as nextPageEventId, which can be used to page through all the Hub events. Set it to 0 to start from the first event.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HubEventsAPI.FetchEvents(context.Background()).FromEventId(fromEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HubEventsAPI.FetchEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchEvents`: FetchEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `HubEventsAPI.FetchEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromEventId** | **int32** | An optional Hub Id to start getting events from. This is also returned from the API as nextPageEventId, which can be used to page through all the Hub events. Set it to 0 to start from the first event.  | 

### Return type

[**FetchEvents200Response**](FetchEvents200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupEvent

> HubEvent LookupEvent(ctx).EventId(eventId).Execute()

Event by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/hub"
)

func main() {
	eventId := int32(56) // int32 | The Hub Id of the event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HubEventsAPI.LookupEvent(context.Background()).EventId(eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HubEventsAPI.LookupEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupEvent`: HubEvent
	fmt.Fprintf(os.Stdout, "Response from `HubEventsAPI.LookupEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventId** | **int32** | The Hub Id of the event | 

### Return type

[**HubEvent**](HubEvent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

