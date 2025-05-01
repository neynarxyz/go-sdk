# \StorageAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LookupUserStorageLimit**](StorageAPI.md#LookupUserStorageLimit) | **Get** /v1/storageLimitsByFid | FID&#39;s limits



## LookupUserStorageLimit

> StorageLimitsResponse LookupUserStorageLimit(ctx).Fid(fid).Execute()

FID's limits



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
	fid := int32(3) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.LookupUserStorageLimit(context.Background()).Fid(fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.LookupUserStorageLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupUserStorageLimit`: StorageLimitsResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.LookupUserStorageLimit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupUserStorageLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** |  | 

### Return type

[**StorageLimitsResponse**](StorageLimitsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

