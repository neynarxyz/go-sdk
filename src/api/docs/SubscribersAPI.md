# \SubscribersAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSubscribedToForFid**](SubscribersAPI.md#FetchSubscribedToForFid) | **Get** /farcaster/user/subscribed_to | Subscribed to
[**FetchSubscribersForFid**](SubscribersAPI.md#FetchSubscribersForFid) | **Get** /farcaster/user/subscribers | Subscribers of a user
[**FetchSubscriptionCheck**](SubscribersAPI.md#FetchSubscriptionCheck) | **Get** /stp/subscription_check | Hypersub subscription check
[**FetchSubscriptionsForFid**](SubscribersAPI.md#FetchSubscriptionsForFid) | **Get** /farcaster/user/subscriptions_created | Subscriptions created by FID



## FetchSubscribedToForFid

> SubscribedToResponse FetchSubscribedToForFid(ctx).Fid(fid).SubscriptionProvider(subscriptionProvider).ViewerFid(viewerFid).Execute()

Subscribed to



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	fid := int32(3206) // int32 | 
	subscriptionProvider := openapiclient.SubscriptionProvider("fabric_stp") // SubscriptionProvider | 
	viewerFid := int32(3) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscribersAPI.FetchSubscribedToForFid(context.Background()).Fid(fid).SubscriptionProvider(subscriptionProvider).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscribersAPI.FetchSubscribedToForFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSubscribedToForFid`: SubscribedToResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscribersAPI.FetchSubscribedToForFid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchSubscribedToForFidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** |  | 
 **subscriptionProvider** | [**SubscriptionProvider**](SubscriptionProvider.md) |  | 
 **viewerFid** | **int32** |  | 

### Return type

[**SubscribedToResponse**](SubscribedToResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSubscribersForFid

> SubscribersResponse FetchSubscribersForFid(ctx).Fid(fid).SubscriptionProvider(subscriptionProvider).ViewerFid(viewerFid).Execute()

Subscribers of a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	fid := int32(3206) // int32 | 
	subscriptionProvider := openapiclient.SubscriptionProviders("fabric_stp") // SubscriptionProviders | 
	viewerFid := int32(3) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscribersAPI.FetchSubscribersForFid(context.Background()).Fid(fid).SubscriptionProvider(subscriptionProvider).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscribersAPI.FetchSubscribersForFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSubscribersForFid`: SubscribersResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscribersAPI.FetchSubscribersForFid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchSubscribersForFidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** |  | 
 **subscriptionProvider** | [**SubscriptionProviders**](SubscriptionProviders.md) |  | 
 **viewerFid** | **int32** |  | 

### Return type

[**SubscribersResponse**](SubscribersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSubscriptionCheck

> SubscriptionCheckResponse FetchSubscriptionCheck(ctx).Addresses(addresses).ContractAddress(contractAddress).ChainId(chainId).Execute()

Hypersub subscription check



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	addresses := "0xedd3783e8c7c52b80cfbd026a63c207edc9cbee7,0x5a927ac639636e534b678e81768ca19e2c6280b7" // string | Comma separated list of Ethereum addresses, up to 350 at a time
	contractAddress := "0x76ad4cb9ac51c09f4d9c2cadcea75c9fa9074e5b" // string | Ethereum address of the STP contract
	chainId := "8453" // string | Chain ID of the STP contract

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscribersAPI.FetchSubscriptionCheck(context.Background()).Addresses(addresses).ContractAddress(contractAddress).ChainId(chainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscribersAPI.FetchSubscriptionCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSubscriptionCheck`: SubscriptionCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscribersAPI.FetchSubscriptionCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchSubscriptionCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addresses** | **string** | Comma separated list of Ethereum addresses, up to 350 at a time | 
 **contractAddress** | **string** | Ethereum address of the STP contract | 
 **chainId** | **string** | Chain ID of the STP contract | 

### Return type

[**SubscriptionCheckResponse**](SubscriptionCheckResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSubscriptionsForFid

> SubscriptionsResponse FetchSubscriptionsForFid(ctx).Fid(fid).SubscriptionProvider(subscriptionProvider).Execute()

Subscriptions created by FID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	fid := int32(528) // int32 | 
	subscriptionProvider := openapiclient.SubscriptionProvider("fabric_stp") // SubscriptionProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscribersAPI.FetchSubscriptionsForFid(context.Background()).Fid(fid).SubscriptionProvider(subscriptionProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscribersAPI.FetchSubscriptionsForFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSubscriptionsForFid`: SubscriptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscribersAPI.FetchSubscriptionsForFid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchSubscriptionsForFidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** |  | 
 **subscriptionProvider** | [**SubscriptionProvider**](SubscriptionProvider.md) |  | 

### Return type

[**SubscriptionsResponse**](SubscriptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

