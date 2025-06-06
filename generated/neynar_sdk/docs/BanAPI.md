# \BanAPI

All URIs are relative to *https://api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBans**](BanAPI.md#DeleteBans) | **Delete** /v2/farcaster/ban/ | Unban FIDs from app
[**FetchBanList**](BanAPI.md#FetchBanList) | **Get** /v2/farcaster/ban/list/ | Banned FIDs of app
[**PublishBans**](BanAPI.md#PublishBans) | **Post** /v2/farcaster/ban/ | Ban FIDs from app



## DeleteBans

> BanResponse DeleteBans(ctx).BanReqBody(banReqBody).Execute()

Unban FIDs from app



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
	banReqBody := *openapiclient.NewBanReqBody([]int32{int32(3)}) // BanReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BanAPI.DeleteBans(context.Background()).BanReqBody(banReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BanAPI.DeleteBans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBans`: BanResponse
	fmt.Fprintf(os.Stdout, "Response from `BanAPI.DeleteBans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **banReqBody** | [**BanReqBody**](BanReqBody.md) |  | 

### Return type

[**BanResponse**](BanResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBanList

> BanListResponse FetchBanList(ctx).XNeynarExperimental(xNeynarExperimental).Limit(limit).Cursor(cursor).Execute()

Banned FIDs of app



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
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)
	limit := int32(20) // int32 | Number of results to fetch (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BanAPI.FetchBanList(context.Background()).XNeynarExperimental(xNeynarExperimental).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BanAPI.FetchBanList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchBanList`: BanListResponse
	fmt.Fprintf(os.Stdout, "Response from `BanAPI.FetchBanList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchBanListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**BanListResponse**](BanListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishBans

> BanResponse PublishBans(ctx).BanReqBody(banReqBody).Execute()

Ban FIDs from app



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
	banReqBody := *openapiclient.NewBanReqBody([]int32{int32(3)}) // BanReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BanAPI.PublishBans(context.Background()).BanReqBody(banReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BanAPI.PublishBans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishBans`: BanResponse
	fmt.Fprintf(os.Stdout, "Response from `BanAPI.PublishBans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishBansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **banReqBody** | [**BanReqBody**](BanReqBody.md) |  | 

### Return type

[**BanResponse**](BanResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

