# \OnChainEventsAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchUserOnChainEvents**](OnChainEventsAPI.md#FetchUserOnChainEvents) | **Get** /v1/onChainEventsByFid | Fetch a list of on-chain events provided by an FID
[**FetchUserOnChainSignersEvents**](OnChainEventsAPI.md#FetchUserOnChainSignersEvents) | **Get** /v1/onChainSignersByFid | Fetch a list of signers provided by an FID
[**LookupOnChainIdRegistryEventByAddress**](OnChainEventsAPI.md#LookupOnChainIdRegistryEventByAddress) | **Get** /v1/onChainIdRegistryEventByAddress | Fetch an on-chain ID Registry Event for a given Address



## FetchUserOnChainEvents

> FetchUserOnChainEvents200Response FetchUserOnChainEvents(ctx).Fid(fid).EventType(eventType).Execute()

Fetch a list of on-chain events provided by an FID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/hub"
)

func main() {
	fid := int32(616) // int32 | The FID being requested
	eventType := openapiclient.OnChainEventType("EVENT_TYPE_SIGNER") // OnChainEventType | The numeric or string value of the event type being requested (default to "EVENT_TYPE_SIGNER")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEventsAPI.FetchUserOnChainEvents(context.Background()).Fid(fid).EventType(eventType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEventsAPI.FetchUserOnChainEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserOnChainEvents`: FetchUserOnChainEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `OnChainEventsAPI.FetchUserOnChainEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserOnChainEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID being requested | 
 **eventType** | [**OnChainEventType**](OnChainEventType.md) | The numeric or string value of the event type being requested | [default to &quot;EVENT_TYPE_SIGNER&quot;]

### Return type

[**FetchUserOnChainEvents200Response**](FetchUserOnChainEvents200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserOnChainSignersEvents

> FetchUserOnChainSignersEvents200Response FetchUserOnChainSignersEvents(ctx).Fid(fid).Signer(signer).Execute()

Fetch a list of signers provided by an FID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/hub"
)

func main() {
	fid := int32(616) // int32 | The FID being requested
	signer := "0x0852c07b5695ff94138b025e3f9b4788e06133f04e254f0ea0eb85a06e999cdd" // string | The optional key of signer (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEventsAPI.FetchUserOnChainSignersEvents(context.Background()).Fid(fid).Signer(signer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEventsAPI.FetchUserOnChainSignersEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserOnChainSignersEvents`: FetchUserOnChainSignersEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `OnChainEventsAPI.FetchUserOnChainSignersEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserOnChainSignersEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID being requested | 
 **signer** | **string** | The optional key of signer | 

### Return type

[**FetchUserOnChainSignersEvents200Response**](FetchUserOnChainSignersEvents200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupOnChainIdRegistryEventByAddress

> OnChainEventIdRegister LookupOnChainIdRegistryEventByAddress(ctx).Address(address).Execute()

Fetch an on-chain ID Registry Event for a given Address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/hub"
)

func main() {
	address := "0x6b0bda3f2ffed5efc83fa8c024acff1dd45793f1" // string | The ETH address being requested

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEventsAPI.LookupOnChainIdRegistryEventByAddress(context.Background()).Address(address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEventsAPI.LookupOnChainIdRegistryEventByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupOnChainIdRegistryEventByAddress`: OnChainEventIdRegister
	fmt.Fprintf(os.Stdout, "Response from `OnChainEventsAPI.LookupOnChainIdRegistryEventByAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupOnChainIdRegistryEventByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | The ETH address being requested | 

### Return type

[**OnChainEventIdRegister**](OnChainEventIdRegister.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

