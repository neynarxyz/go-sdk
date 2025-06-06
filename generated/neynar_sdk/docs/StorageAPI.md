# \StorageAPI

All URIs are relative to *https://api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuyStorage**](StorageAPI.md#BuyStorage) | **Post** /v2/farcaster/storage/buy/ | Buy storage
[**LookupUserStorageAllocations**](StorageAPI.md#LookupUserStorageAllocations) | **Get** /v2/farcaster/storage/allocations/ | Allocation of user
[**LookupUserStorageUsage**](StorageAPI.md#LookupUserStorageUsage) | **Get** /v2/farcaster/storage/usage/ | Usage of user



## BuyStorage

> StorageAllocationsResponse BuyStorage(ctx).BuyStorageReqBody(buyStorageReqBody).Execute()

Buy storage



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
	buyStorageReqBody := *openapiclient.NewBuyStorageReqBody(int32(1)) // BuyStorageReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.BuyStorage(context.Background()).BuyStorageReqBody(buyStorageReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.BuyStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuyStorage`: StorageAllocationsResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.BuyStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuyStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyStorageReqBody** | [**BuyStorageReqBody**](BuyStorageReqBody.md) |  | 

### Return type

[**StorageAllocationsResponse**](StorageAllocationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupUserStorageAllocations

> StorageAllocationsResponse LookupUserStorageAllocations(ctx).Fid(fid).Execute()

Allocation of user



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
	fid := int32(3) // int32 | The unique identifier of a farcaster user or app (unsigned integer)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.LookupUserStorageAllocations(context.Background()).Fid(fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.LookupUserStorageAllocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupUserStorageAllocations`: StorageAllocationsResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.LookupUserStorageAllocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupUserStorageAllocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The unique identifier of a farcaster user or app (unsigned integer) | 

### Return type

[**StorageAllocationsResponse**](StorageAllocationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupUserStorageUsage

> StorageUsageResponse LookupUserStorageUsage(ctx).Fid(fid).Execute()

Usage of user



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
	fid := int32(3) // int32 | The unique identifier of a farcaster user or app (unsigned integer)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.LookupUserStorageUsage(context.Background()).Fid(fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.LookupUserStorageUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupUserStorageUsage`: StorageUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.LookupUserStorageUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupUserStorageUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The unique identifier of a farcaster user or app (unsigned integer) | 

### Return type

[**StorageUsageResponse**](StorageUsageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

