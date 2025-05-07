# \FidsAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchFids**](FidsAPI.md#FetchFids) | **Get** /v1/fids | Fetch a list of all the FIDs



## FetchFids

> FidsResponse FetchFids(ctx).ShardId(shardId).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Fetch a list of all the FIDs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/hub"
)

func main() {
	shardId := int32(1) // int32 | The shard ID to filter by
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FidsAPI.FetchFids(context.Background()).ShardId(shardId).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FidsAPI.FetchFids``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchFids`: FidsResponse
	fmt.Fprintf(os.Stdout, "Response from `FidsAPI.FetchFids`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shardId** | **int32** | The shard ID to filter by | 
 **pageSize** | **int32** | Maximum number of messages to return in a single response | 
 **reverse** | **bool** | Reverse the sort order, returning latest messages first | 
 **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FidsResponse**](FidsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

