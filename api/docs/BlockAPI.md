# \BlockAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBlock**](BlockAPI.md#DeleteBlock) | **Delete** /farcaster/block | Unblock FID
[**FetchBlockList**](BlockAPI.md#FetchBlockList) | **Get** /farcaster/block/list | Blocked / Blocked by FIDs
[**PublishBlock**](BlockAPI.md#PublishBlock) | **Post** /farcaster/block | Block FID



## DeleteBlock

> OperationResponse DeleteBlock(ctx).BlockReqBody(blockReqBody).Execute()

Unblock FID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/api"
)

func main() {
	blockReqBody := *openapiclient.NewBlockReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", int32(3)) // BlockReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockAPI.DeleteBlock(context.Background()).BlockReqBody(blockReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockAPI.DeleteBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlock`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `BlockAPI.DeleteBlock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockReqBody** | [**BlockReqBody**](BlockReqBody.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBlockList

> BlockListResponse FetchBlockList(ctx).BlockerFid(blockerFid).BlockedFid(blockedFid).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()

Blocked / Blocked by FIDs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/api"
)

func main() {
	blockerFid := int32(194) // int32 | Providing this will return the users that this user has blocked (optional)
	blockedFid := int32(194) // int32 | Providing this will return the users that have blocked this user (optional)
	limit := int32(20) // int32 | Number of results to fetch (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockAPI.FetchBlockList(context.Background()).BlockerFid(blockerFid).BlockedFid(blockedFid).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockAPI.FetchBlockList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchBlockList`: BlockListResponse
	fmt.Fprintf(os.Stdout, "Response from `BlockAPI.FetchBlockList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchBlockListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockerFid** | **int32** | Providing this will return the users that this user has blocked | 
 **blockedFid** | **int32** | Providing this will return the users that have blocked this user | 
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **cursor** | **string** | Pagination cursor. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**BlockListResponse**](BlockListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishBlock

> OperationResponse PublishBlock(ctx).BlockReqBody(blockReqBody).Execute()

Block FID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/api"
)

func main() {
	blockReqBody := *openapiclient.NewBlockReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", int32(3)) // BlockReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockAPI.PublishBlock(context.Background()).BlockReqBody(blockReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockAPI.PublishBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishBlock`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `BlockAPI.PublishBlock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockReqBody** | [**BlockReqBody**](BlockReqBody.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

