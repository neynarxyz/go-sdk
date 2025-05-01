# \MuteAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMute**](MuteAPI.md#DeleteMute) | **Delete** /farcaster/mute | Unmute FID
[**FetchMuteList**](MuteAPI.md#FetchMuteList) | **Get** /farcaster/mute/list | Muted FIDs of user
[**PublishMute**](MuteAPI.md#PublishMute) | **Post** /farcaster/mute | Mute FID



## DeleteMute

> MuteResponse DeleteMute(ctx).MuteReqBody(muteReqBody).Execute()

Unmute FID



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
	muteReqBody := *openapiclient.NewMuteReqBody(int32(123), int32(123)) // MuteReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MuteAPI.DeleteMute(context.Background()).MuteReqBody(muteReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MuteAPI.DeleteMute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMute`: MuteResponse
	fmt.Fprintf(os.Stdout, "Response from `MuteAPI.DeleteMute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **muteReqBody** | [**MuteReqBody**](MuteReqBody.md) |  | 

### Return type

[**MuteResponse**](MuteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMuteList

> MuteListResponse FetchMuteList(ctx).Fid(fid).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()

Muted FIDs of user



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
	fid := int32(194) // int32 | The user's FID (identifier)
	limit := int32(20) // int32 | Number of results to fetch (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MuteAPI.FetchMuteList(context.Background()).Fid(fid).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MuteAPI.FetchMuteList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchMuteList`: MuteListResponse
	fmt.Fprintf(os.Stdout, "Response from `MuteAPI.FetchMuteList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchMuteListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The user&#39;s FID (identifier) | 
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **cursor** | **string** | Pagination cursor. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**MuteListResponse**](MuteListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishMute

> MuteResponse PublishMute(ctx).MuteReqBody(muteReqBody).Execute()

Mute FID



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
	muteReqBody := *openapiclient.NewMuteReqBody(int32(123), int32(123)) // MuteReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MuteAPI.PublishMute(context.Background()).MuteReqBody(muteReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MuteAPI.PublishMute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishMute`: MuteResponse
	fmt.Fprintf(os.Stdout, "Response from `MuteAPI.PublishMute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishMuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **muteReqBody** | [**MuteReqBody**](MuteReqBody.md) |  | 

### Return type

[**MuteResponse**](MuteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

