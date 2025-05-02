# \FeedAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCastsForUser**](FeedAPI.md#FetchCastsForUser) | **Get** /farcaster/feed/user/casts | Chronologically
[**FetchFeed**](FeedAPI.md#FetchFeed) | **Get** /farcaster/feed | By filters
[**FetchFeedByChannelIds**](FeedAPI.md#FetchFeedByChannelIds) | **Get** /farcaster/feed/channels | By channel IDs
[**FetchFeedByParentUrls**](FeedAPI.md#FetchFeedByParentUrls) | **Get** /farcaster/feed/parent_urls | By parent URLs
[**FetchFeedForYou**](FeedAPI.md#FetchFeedForYou) | **Get** /farcaster/feed/for_you | For you
[**FetchFramesOnlyFeed**](FeedAPI.md#FetchFramesOnlyFeed) | **Get** /farcaster/feed/frames | Casts with mini apps
[**FetchPopularCastsByUser**](FeedAPI.md#FetchPopularCastsByUser) | **Get** /farcaster/feed/user/popular | 10 most popular casts
[**FetchRepliesAndRecastsForUser**](FeedAPI.md#FetchRepliesAndRecastsForUser) | **Get** /farcaster/feed/user/replies_and_recasts | Replies and recasts
[**FetchTrendingFeed**](FeedAPI.md#FetchTrendingFeed) | **Get** /farcaster/feed/trending | Trending feeds
[**FetchUserFollowingFeed**](FeedAPI.md#FetchUserFollowingFeed) | **Get** /farcaster/feed/following | Following



## FetchCastsForUser

> FeedResponse FetchCastsForUser(ctx).Fid(fid).AppFid(appFid).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).IncludeReplies(includeReplies).ParentUrl(parentUrl).ChannelId(channelId).XNeynarExperimental(xNeynarExperimental).Execute()

Chronologically



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
	fid := int32(194) // int32 | FID of user whose recent casts you want to fetch
	appFid := int32(9152) // int32 | Optionally filter to casts created via a specific app FID, e.g. 9152 for Warpcast (optional)
	viewerFid := int32(3) // int32 | FID of the user viewing the feed (optional)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor (optional)
	includeReplies := true // bool | Include reply casts by the author in the response, true by default (optional) (default to true)
	parentUrl := "parentUrl_example" // string | Parent URL to filter the feed; mutually exclusive with channel_id (optional)
	channelId := "channelId_example" // string | Channel ID to filter the feed; mutually exclusive with parent_url (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FetchCastsForUser(context.Background()).Fid(fid).AppFid(appFid).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).IncludeReplies(includeReplies).ParentUrl(parentUrl).ChannelId(channelId).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FetchCastsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchCastsForUser`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FetchCastsForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchCastsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | FID of user whose recent casts you want to fetch | 
 **appFid** | **int32** | Optionally filter to casts created via a specific app FID, e.g. 9152 for Warpcast | 
 **viewerFid** | **int32** | FID of the user viewing the feed | 
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor | 
 **includeReplies** | **bool** | Include reply casts by the author in the response, true by default | [default to true]
 **parentUrl** | **string** | Parent URL to filter the feed; mutually exclusive with channel_id | 
 **channelId** | **string** | Channel ID to filter the feed; mutually exclusive with parent_url | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFeed

> FeedResponse FetchFeed(ctx).FeedType(feedType).FilterType(filterType).Fid(fid).Fids(fids).ParentUrl(parentUrl).ChannelId(channelId).MembersOnly(membersOnly).EmbedUrl(embedUrl).EmbedTypes(embedTypes).WithRecasts(withRecasts).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).XNeynarExperimental(xNeynarExperimental).Execute()

By filters



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
	feedType := openapiclient.FeedType("following") // FeedType | Defaults to following (requires FID or address). If set to filter (requires filter_type) (default to "following")
	filterType := openapiclient.FilterType("fids") // FilterType | Used when feed_type=filter. Can be set to FIDs (requires FIDs) or parent_url (requires parent_url) or channel_id (requires channel_id) (optional)
	fid := int32(56) // int32 | (Optional) FID of user whose feed you want to create. By default, the API expects this field, except if you pass a filter_type (optional)
	fids := "3,2,194" // string | Used when filter_type=FIDs . Create a feed based on a list of FIDs. Max array size is 100. Requires feed_type and filter_type. (optional)
	parentUrl := "parentUrl_example" // string | Used when filter_type=parent_url can be used to fetch content under any parent url e.g. FIP-2 channels on Warpcast. Requires feed_type and filter_type. (optional)
	channelId := "channelId_example" // string | Used when filter_type=channel_id can be used to fetch casts under a channel. Requires feed_type and filter_type. (optional)
	membersOnly := true // bool | Used when filter_type=channel_id. Only include casts from members of the channel. True by default. (optional) (default to true)
	embedUrl := "embedUrl_example" // string | Used when filter_type=embed_url. Casts with embedded URLs prefixed by this embed_url param will be returned. We normalize your given URL prefix and prepend 'https://' if no protocol is included. Requires feed_type and filter_type. (optional)
	embedTypes := []openapiclient.EmbedType{openapiclient.EmbedType("text")} // []EmbedType | Used when filter_type=embed_types can be used to fetch all casts with matching content types. Requires feed_type and filter_type. (optional)
	withRecasts := true // bool | Include recasts in the response, true by default (optional) (default to true)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	viewerFid := int32(3) // int32 | Providing this will return a feed that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FetchFeed(context.Background()).FeedType(feedType).FilterType(filterType).Fid(fid).Fids(fids).ParentUrl(parentUrl).ChannelId(channelId).MembersOnly(membersOnly).EmbedUrl(embedUrl).EmbedTypes(embedTypes).WithRecasts(withRecasts).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FetchFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchFeed`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FetchFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedType** | [**FeedType**](FeedType.md) | Defaults to following (requires FID or address). If set to filter (requires filter_type) | [default to &quot;following&quot;]
 **filterType** | [**FilterType**](FilterType.md) | Used when feed_type&#x3D;filter. Can be set to FIDs (requires FIDs) or parent_url (requires parent_url) or channel_id (requires channel_id) | 
 **fid** | **int32** | (Optional) FID of user whose feed you want to create. By default, the API expects this field, except if you pass a filter_type | 
 **fids** | **string** | Used when filter_type&#x3D;FIDs . Create a feed based on a list of FIDs. Max array size is 100. Requires feed_type and filter_type. | 
 **parentUrl** | **string** | Used when filter_type&#x3D;parent_url can be used to fetch content under any parent url e.g. FIP-2 channels on Warpcast. Requires feed_type and filter_type. | 
 **channelId** | **string** | Used when filter_type&#x3D;channel_id can be used to fetch casts under a channel. Requires feed_type and filter_type. | 
 **membersOnly** | **bool** | Used when filter_type&#x3D;channel_id. Only include casts from members of the channel. True by default. | [default to true]
 **embedUrl** | **string** | Used when filter_type&#x3D;embed_url. Casts with embedded URLs prefixed by this embed_url param will be returned. We normalize your given URL prefix and prepend &#39;https://&#39; if no protocol is included. Requires feed_type and filter_type. | 
 **embedTypes** | [**[]EmbedType**](EmbedType.md) | Used when filter_type&#x3D;embed_types can be used to fetch all casts with matching content types. Requires feed_type and filter_type. | 
 **withRecasts** | **bool** | Include recasts in the response, true by default | [default to true]
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 
 **viewerFid** | **int32** | Providing this will return a feed that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFeedByChannelIds

> FeedResponse FetchFeedByChannelIds(ctx).ChannelIds(channelIds).WithRecasts(withRecasts).ViewerFid(viewerFid).WithReplies(withReplies).MembersOnly(membersOnly).Fids(fids).Limit(limit).Cursor(cursor).ShouldModerate(shouldModerate).XNeynarExperimental(xNeynarExperimental).Execute()

By channel IDs



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
	channelIds := "neynar,farcaster" // string | Comma separated list of up to 10 channel IDs e.g. neynar,farcaster
	withRecasts := true // bool | Include recasts in the response, true by default (optional) (default to true)
	viewerFid := int32(3) // int32 | Providing this will return a feed that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	withReplies := true // bool | Include replies in the response, false by default (optional) (default to false)
	membersOnly := true // bool | Only include casts from members of the channel. True by default. (optional) (default to true)
	fids := "fids_example" // string | Comma separated list of FIDs to filter the feed by, up to 10 at a time (optional)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	shouldModerate := true // bool | If true, only casts that have been liked by the moderator (if one exists) will be returned. (optional) (default to false)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FetchFeedByChannelIds(context.Background()).ChannelIds(channelIds).WithRecasts(withRecasts).ViewerFid(viewerFid).WithReplies(withReplies).MembersOnly(membersOnly).Fids(fids).Limit(limit).Cursor(cursor).ShouldModerate(shouldModerate).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FetchFeedByChannelIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchFeedByChannelIds`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FetchFeedByChannelIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFeedByChannelIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelIds** | **string** | Comma separated list of up to 10 channel IDs e.g. neynar,farcaster | 
 **withRecasts** | **bool** | Include recasts in the response, true by default | [default to true]
 **viewerFid** | **int32** | Providing this will return a feed that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **withReplies** | **bool** | Include replies in the response, false by default | [default to false]
 **membersOnly** | **bool** | Only include casts from members of the channel. True by default. | [default to true]
 **fids** | **string** | Comma separated list of FIDs to filter the feed by, up to 10 at a time | 
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 
 **shouldModerate** | **bool** | If true, only casts that have been liked by the moderator (if one exists) will be returned. | [default to false]
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFeedByParentUrls

> FeedResponse FetchFeedByParentUrls(ctx).ParentUrls(parentUrls).WithRecasts(withRecasts).ViewerFid(viewerFid).WithReplies(withReplies).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()

By parent URLs



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
	parentUrls := "chain://eip155:1/erc721:0xd4498134211baad5846ce70ce04e7c4da78931cc" // string | Comma separated list of parent_urls
	withRecasts := true // bool | Include recasts in the response, true by default (optional) (default to true)
	viewerFid := int32(3) // int32 | Providing this will return a feed that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	withReplies := true // bool | Include replies in the response, false by default (optional) (default to false)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FetchFeedByParentUrls(context.Background()).ParentUrls(parentUrls).WithRecasts(withRecasts).ViewerFid(viewerFid).WithReplies(withReplies).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FetchFeedByParentUrls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchFeedByParentUrls`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FetchFeedByParentUrls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFeedByParentUrlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentUrls** | **string** | Comma separated list of parent_urls | 
 **withRecasts** | **bool** | Include recasts in the response, true by default | [default to true]
 **viewerFid** | **int32** | Providing this will return a feed that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **withReplies** | **bool** | Include replies in the response, false by default | [default to false]
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFeedForYou

> FeedResponse FetchFeedForYou(ctx).Fid(fid).ViewerFid(viewerFid).Provider(provider).Limit(limit).Cursor(cursor).ProviderMetadata(providerMetadata).XNeynarExperimental(xNeynarExperimental).Execute()

For you



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
	fid := int32(194) // int32 | FID of user whose feed you want to create
	viewerFid := int32(3) // int32 | Providing this will return a feed that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	provider := openapiclient.ForYouProvider("neynar") // ForYouProvider |  (optional) (default to "neynar")
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	providerMetadata := "providerMetadata_example" // string | provider_metadata is a URI-encoded stringified JSON object that can be used to pass additional metadata to the provider. Only available for mbd provider right now. See [here](https://docs.neynar.com/docs/feed-for-you-w-external-providers) on how to use.  (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FetchFeedForYou(context.Background()).Fid(fid).ViewerFid(viewerFid).Provider(provider).Limit(limit).Cursor(cursor).ProviderMetadata(providerMetadata).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FetchFeedForYou``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchFeedForYou`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FetchFeedForYou`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFeedForYouRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | FID of user whose feed you want to create | 
 **viewerFid** | **int32** | Providing this will return a feed that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **provider** | [**ForYouProvider**](ForYouProvider.md) |  | [default to &quot;neynar&quot;]
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 
 **providerMetadata** | **string** | provider_metadata is a URI-encoded stringified JSON object that can be used to pass additional metadata to the provider. Only available for mbd provider right now. See [here](https://docs.neynar.com/docs/feed-for-you-w-external-providers) on how to use.  | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFramesOnlyFeed

> FeedResponse FetchFramesOnlyFeed(ctx).Limit(limit).ViewerFid(viewerFid).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()

Casts with mini apps



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
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	viewerFid := int32(3) // int32 | Providing this will return a feed that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FetchFramesOnlyFeed(context.Background()).Limit(limit).ViewerFid(viewerFid).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FetchFramesOnlyFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchFramesOnlyFeed`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FetchFramesOnlyFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFramesOnlyFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **viewerFid** | **int32** | Providing this will return a feed that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **cursor** | **string** | Pagination cursor. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPopularCastsByUser

> BulkCastsResponse FetchPopularCastsByUser(ctx).Fid(fid).ViewerFid(viewerFid).Execute()

10 most popular casts



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
	fid := int32(194) // int32 | FID of user whose feed you want to create
	viewerFid := int32(3) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FetchPopularCastsByUser(context.Background()).Fid(fid).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FetchPopularCastsByUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPopularCastsByUser`: BulkCastsResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FetchPopularCastsByUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchPopularCastsByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | FID of user whose feed you want to create | 
 **viewerFid** | **int32** |  | 

### Return type

[**BulkCastsResponse**](BulkCastsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRepliesAndRecastsForUser

> FeedResponse FetchRepliesAndRecastsForUser(ctx).Fid(fid).Filter(filter).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).Execute()

Replies and recasts



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
	fid := int32(194) // int32 | FID of user whose replies and recasts you want to fetch
	filter := "replies" // string | filter to fetch only replies or recasts (optional) (default to "all")
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	viewerFid := int32(3) // int32 | Providing this will return a feed that respects this user's mutes and blocks and includes `viewer_context`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FetchRepliesAndRecastsForUser(context.Background()).Fid(fid).Filter(filter).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FetchRepliesAndRecastsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchRepliesAndRecastsForUser`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FetchRepliesAndRecastsForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchRepliesAndRecastsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | FID of user whose replies and recasts you want to fetch | 
 **filter** | **string** | filter to fetch only replies or recasts | [default to &quot;all&quot;]
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 
 **viewerFid** | **int32** | Providing this will return a feed that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTrendingFeed

> FeedResponse FetchTrendingFeed(ctx).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).TimeWindow(timeWindow).ChannelId(channelId).ParentUrl(parentUrl).Provider(provider).ProviderMetadata(providerMetadata).XNeynarExperimental(xNeynarExperimental).Execute()

Trending feeds



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
	limit := int32(10) // int32 | Number of results to fetch (optional) (default to 10)
	cursor := "cursor_example" // string | Pagination cursor (optional)
	viewerFid := int32(3) // int32 | Providing this will return a feed that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	timeWindow := openapiclient.TrendingTimeWindow("1h") // TrendingTimeWindow | Time window for trending casts (7d window for channel feeds only) (optional)
	channelId := "neynar" // string | Channel ID to filter trending casts. Less active channels might have no casts in the time window selected. Provide either `channel_id` or `parent_url`, not both. (optional)
	parentUrl := "chain://eip155:1/erc721:0xd4498134211baad5846ce70ce04e7c4da78931cc" // string | Parent URL to filter trending casts. Less active channels might have no casts in the time window selected. Provide either `channel_id` or `parent_url`, not both. (optional)
	provider := openapiclient.FeedTrendingProvider("neynar") // FeedTrendingProvider | The provider of the trending casts feed. (optional) (default to "neynar")
	providerMetadata := "providerMetadata_example" // string | provider_metadata is a URI-encoded stringified JSON object that can be used to pass additional metadata to the provider. Only available for mbd provider right now. See [here](https://docs.neynar.com/docs/feed-for-you-w-external-providers) on how to use.  (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FetchTrendingFeed(context.Background()).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).TimeWindow(timeWindow).ChannelId(channelId).ParentUrl(parentUrl).Provider(provider).ProviderMetadata(providerMetadata).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FetchTrendingFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchTrendingFeed`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FetchTrendingFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchTrendingFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to fetch | [default to 10]
 **cursor** | **string** | Pagination cursor | 
 **viewerFid** | **int32** | Providing this will return a feed that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **timeWindow** | [**TrendingTimeWindow**](TrendingTimeWindow.md) | Time window for trending casts (7d window for channel feeds only) | 
 **channelId** | **string** | Channel ID to filter trending casts. Less active channels might have no casts in the time window selected. Provide either &#x60;channel_id&#x60; or &#x60;parent_url&#x60;, not both. | 
 **parentUrl** | **string** | Parent URL to filter trending casts. Less active channels might have no casts in the time window selected. Provide either &#x60;channel_id&#x60; or &#x60;parent_url&#x60;, not both. | 
 **provider** | [**FeedTrendingProvider**](FeedTrendingProvider.md) | The provider of the trending casts feed. | [default to &quot;neynar&quot;]
 **providerMetadata** | **string** | provider_metadata is a URI-encoded stringified JSON object that can be used to pass additional metadata to the provider. Only available for mbd provider right now. See [here](https://docs.neynar.com/docs/feed-for-you-w-external-providers) on how to use.  | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserFollowingFeed

> FeedResponse FetchUserFollowingFeed(ctx).Fid(fid).ViewerFid(viewerFid).WithRecasts(withRecasts).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()

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
	fid := int32(3) // int32 | FID of user whose feed you want to create
	viewerFid := int32(3) // int32 | Providing this will return a feed that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	withRecasts := true // bool | Include recasts in the response, true by default (optional) (default to true)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FetchUserFollowingFeed(context.Background()).Fid(fid).ViewerFid(viewerFid).WithRecasts(withRecasts).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FetchUserFollowingFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserFollowingFeed`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FetchUserFollowingFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserFollowingFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | FID of user whose feed you want to create | 
 **viewerFid** | **int32** | Providing this will return a feed that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **withRecasts** | **bool** | Include recasts in the response, true by default | [default to true]
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

