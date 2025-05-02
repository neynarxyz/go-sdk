# \FollowsAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchFollowSuggestions**](FollowsAPI.md#FetchFollowSuggestions) | **Get** /farcaster/following/suggested | Suggest Follows
[**FetchRelevantFollowers**](FollowsAPI.md#FetchRelevantFollowers) | **Get** /farcaster/followers/relevant | Relevant followers
[**FetchUserFollowers**](FollowsAPI.md#FetchUserFollowers) | **Get** /farcaster/followers | Followers
[**FetchUserFollowing**](FollowsAPI.md#FetchUserFollowing) | **Get** /farcaster/following | Following



## FetchFollowSuggestions

> UsersResponse FetchFollowSuggestions(ctx).Fid(fid).ViewerFid(viewerFid).Limit(limit).XNeynarExperimental(xNeynarExperimental).Execute()

Suggest Follows



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/api"
)

func main() {
	fid := int32(2) // int32 | FID of the user whose following you want to fetch.
	viewerFid := int32(3) // int32 | Providing this will return a list of users that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FollowsAPI.FetchFollowSuggestions(context.Background()).Fid(fid).ViewerFid(viewerFid).Limit(limit).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FollowsAPI.FetchFollowSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchFollowSuggestions`: UsersResponse
	fmt.Fprintf(os.Stdout, "Response from `FollowsAPI.FetchFollowSuggestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFollowSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | FID of the user whose following you want to fetch. | 
 **viewerFid** | **int32** | Providing this will return a list of users that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRelevantFollowers

> RelevantFollowersResponse FetchRelevantFollowers(ctx).TargetFid(targetFid).ViewerFid(viewerFid).XNeynarExperimental(xNeynarExperimental).Execute()

Relevant followers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/api"
)

func main() {
	targetFid := int32(56) // int32 | User who's profile you are looking at
	viewerFid := int32(56) // int32 | The FID of the user to customize this response for. Providing this will also return a list of followers that respects this user's mutes and blocks and includes `viewer_context`.
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FollowsAPI.FetchRelevantFollowers(context.Background()).TargetFid(targetFid).ViewerFid(viewerFid).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FollowsAPI.FetchRelevantFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchRelevantFollowers`: RelevantFollowersResponse
	fmt.Fprintf(os.Stdout, "Response from `FollowsAPI.FetchRelevantFollowers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchRelevantFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetFid** | **int32** | User who&#39;s profile you are looking at | 
 **viewerFid** | **int32** | The FID of the user to customize this response for. Providing this will also return a list of followers that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**RelevantFollowersResponse**](RelevantFollowersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserFollowers

> FollowersResponse FetchUserFollowers(ctx).Fid(fid).ViewerFid(viewerFid).SortType(sortType).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()

Followers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/api"
)

func main() {
	fid := int32(56) // int32 | User who's profile you are looking at
	viewerFid := int32(56) // int32 | Providing this will return a list of followers that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	sortType := openapiclient.FollowSortType("desc_chron") // FollowSortType | Sort type for fetch followers. Default is `desc_chron` (optional)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FollowsAPI.FetchUserFollowers(context.Background()).Fid(fid).ViewerFid(viewerFid).SortType(sortType).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FollowsAPI.FetchUserFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserFollowers`: FollowersResponse
	fmt.Fprintf(os.Stdout, "Response from `FollowsAPI.FetchUserFollowers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | User who&#39;s profile you are looking at | 
 **viewerFid** | **int32** | Providing this will return a list of followers that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **sortType** | [**FollowSortType**](FollowSortType.md) | Sort type for fetch followers. Default is &#x60;desc_chron&#x60; | 
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **cursor** | **string** | Pagination cursor. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**FollowersResponse**](FollowersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserFollowing

> FollowersResponse FetchUserFollowing(ctx).Fid(fid).ViewerFid(viewerFid).SortType(sortType).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()

Following



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/api"
)

func main() {
	fid := int32(2) // int32 | FID of the user whose following you want to fetch.
	viewerFid := int32(3) // int32 | Providing this will return a list of users that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	sortType := openapiclient.FollowSortType("desc_chron") // FollowSortType | Optional parameter to sort the users based on different criteria. (optional)
	limit := int32(25) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FollowsAPI.FetchUserFollowing(context.Background()).Fid(fid).ViewerFid(viewerFid).SortType(sortType).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FollowsAPI.FetchUserFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserFollowing`: FollowersResponse
	fmt.Fprintf(os.Stdout, "Response from `FollowsAPI.FetchUserFollowing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserFollowingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | FID of the user whose following you want to fetch. | 
 **viewerFid** | **int32** | Providing this will return a list of users that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **sortType** | [**FollowSortType**](FollowSortType.md) | Optional parameter to sort the users based on different criteria. | 
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**FollowersResponse**](FollowersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

