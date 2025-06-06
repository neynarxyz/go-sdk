# \AppHostAPI

All URIs are relative to *https://api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppHostGetEvent**](AppHostAPI.md#AppHostGetEvent) | **Get** /v2/farcaster/app_host/user/event | Get app host event
[**AppHostGetUserState**](AppHostAPI.md#AppHostGetUserState) | **Get** /v2/farcaster/app_host/user/state/ | Get the user&#39;s notification subscriptions
[**AppHostPostEvent**](AppHostAPI.md#AppHostPostEvent) | **Post** /v2/farcaster/app_host/user/event | Process app host event



## AppHostGetEvent

> AppHostGetEventResponse AppHostGetEvent(ctx).AppDomain(appDomain).Fid(fid).Event(event).Execute()

Get app host event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/neynar_sdk"
)

func main() {
	appDomain := "demo.neynar.com" // string | The domain of the mini app
	fid := int32(56) // int32 | The FID of the user who initiated the event
	event := "event_example" // string | The type of event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppHostAPI.AppHostGetEvent(context.Background()).AppDomain(appDomain).Fid(fid).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppHostAPI.AppHostGetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppHostGetEvent`: AppHostGetEventResponse
	fmt.Fprintf(os.Stdout, "Response from `AppHostAPI.AppHostGetEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppHostGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appDomain** | **string** | The domain of the mini app | 
 **fid** | **int32** | The FID of the user who initiated the event | 
 **event** | **string** | The type of event | 

### Return type

[**AppHostGetEventResponse**](AppHostGetEventResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppHostGetUserState

> AppHostUserStateResponse AppHostGetUserState(ctx).Fid(fid).Execute()

Get the user's notification subscriptions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/neynar_sdk"
)

func main() {
	fid := int32(56) // int32 | The FID of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppHostAPI.AppHostGetUserState(context.Background()).Fid(fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppHostAPI.AppHostGetUserState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppHostGetUserState`: AppHostUserStateResponse
	fmt.Fprintf(os.Stdout, "Response from `AppHostAPI.AppHostGetUserState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppHostGetUserStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID of the user | 

### Return type

[**AppHostUserStateResponse**](AppHostUserStateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppHostPostEvent

> AppHostPostEventResponse AppHostPostEvent(ctx).AppHostPostEventBody(appHostPostEventBody).Execute()

Process app host event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/neynar_sdk"
)

func main() {
	appHostPostEventBody := openapiclient.AppHostPostEventBody{AppHostPostEventBodyOneOf: openapiclient.NewAppHostPostEventBodyOneOf("SignedMessage_example", "demo.neynar.com")} // AppHostPostEventBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppHostAPI.AppHostPostEvent(context.Background()).AppHostPostEventBody(appHostPostEventBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppHostAPI.AppHostPostEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppHostPostEvent`: AppHostPostEventResponse
	fmt.Fprintf(os.Stdout, "Response from `AppHostAPI.AppHostPostEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppHostPostEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appHostPostEventBody** | [**AppHostPostEventBody**](AppHostPostEventBody.md) |  | 

### Return type

[**AppHostPostEventResponse**](AppHostPostEventResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

