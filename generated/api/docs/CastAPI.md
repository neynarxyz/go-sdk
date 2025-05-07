# \CastAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCast**](CastAPI.md#DeleteCast) | **Delete** /farcaster/cast | Delete a cast
[**FetchBulkCasts**](CastAPI.md#FetchBulkCasts) | **Get** /farcaster/casts | Bulk fetch casts
[**FetchComposerActions**](CastAPI.md#FetchComposerActions) | **Get** /farcaster/cast/composer_actions/list | Fetch composer actions
[**FetchEmbeddedUrlMetadata**](CastAPI.md#FetchEmbeddedUrlMetadata) | **Get** /farcaster/cast/embed/crawl | Embedded URL metadata
[**LookupCastByHashOrWarpcastUrl**](CastAPI.md#LookupCastByHashOrWarpcastUrl) | **Get** /farcaster/cast | By hash or URL
[**LookupCastConversation**](CastAPI.md#LookupCastConversation) | **Get** /farcaster/cast/conversation | Conversation for a cast
[**PublishCast**](CastAPI.md#PublishCast) | **Post** /farcaster/cast | Post a cast
[**SearchCasts**](CastAPI.md#SearchCasts) | **Get** /farcaster/cast/search | Search for casts



## DeleteCast

> OperationResponse DeleteCast(ctx).DeleteCastReqBody(deleteCastReqBody).Execute()

Delete a cast



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
	deleteCastReqBody := *openapiclient.NewDeleteCastReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", "TargetHash_example") // DeleteCastReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.DeleteCast(context.Background()).DeleteCastReqBody(deleteCastReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.DeleteCast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCast`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.DeleteCast`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteCastReqBody** | [**DeleteCastReqBody**](DeleteCastReqBody.md) |  | 

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


## FetchBulkCasts

> CastsResponse FetchBulkCasts(ctx).Casts(casts).ViewerFid(viewerFid).SortType(sortType).XNeynarExperimental(xNeynarExperimental).Execute()

Bulk fetch casts



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
	casts := "0xa896906a5e397b4fec247c3ee0e9e4d4990b8004,0x27ff810f7f718afd8c40be236411f017982e0994" // string | Hashes of the cast to be retrived (Comma separated, no spaces)
	viewerFid := int32(3) // int32 | adds viewer_context to cast object to show whether viewer has liked or recasted the cast. (optional)
	sortType := "sortType_example" // string | Optional parameter to sort the casts based on different criteria (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.FetchBulkCasts(context.Background()).Casts(casts).ViewerFid(viewerFid).SortType(sortType).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.FetchBulkCasts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchBulkCasts`: CastsResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.FetchBulkCasts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchBulkCastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **casts** | **string** | Hashes of the cast to be retrived (Comma separated, no spaces) | 
 **viewerFid** | **int32** | adds viewer_context to cast object to show whether viewer has liked or recasted the cast. | 
 **sortType** | **string** | Optional parameter to sort the casts based on different criteria | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**CastsResponse**](CastsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchComposerActions

> CastComposerActionsListResponse FetchComposerActions(ctx).List(list).Limit(limit).Cursor(cursor).Execute()

Fetch composer actions



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
	list := openapiclient.CastComposerType("top") // CastComposerType | Type of list to fetch.
	limit := int32(25) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.FetchComposerActions(context.Background()).List(list).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.FetchComposerActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchComposerActions`: CastComposerActionsListResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.FetchComposerActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchComposerActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | [**CastComposerType**](CastComposerType.md) | Type of list to fetch. | 
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**CastComposerActionsListResponse**](CastComposerActionsListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEmbeddedUrlMetadata

> CastEmbedCrawlResponse FetchEmbeddedUrlMetadata(ctx).Url(url).Execute()

Embedded URL metadata



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
	url := "url_example" // string | URL to crawl metadata of

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.FetchEmbeddedUrlMetadata(context.Background()).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.FetchEmbeddedUrlMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchEmbeddedUrlMetadata`: CastEmbedCrawlResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.FetchEmbeddedUrlMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchEmbeddedUrlMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | URL to crawl metadata of | 

### Return type

[**CastEmbedCrawlResponse**](CastEmbedCrawlResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupCastByHashOrWarpcastUrl

> CastResponse LookupCastByHashOrWarpcastUrl(ctx).Identifier(identifier).Type_(type_).ViewerFid(viewerFid).XNeynarExperimental(xNeynarExperimental).Execute()

By hash or URL



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
	identifier := "https://warpcast.com/rish/0x9288c1" // string | Cast identifier (Its either a url or a hash)
	type_ := openapiclient.CastParamType("url") // CastParamType | 
	viewerFid := int32(3) // int32 | adds viewer_context to cast object to show whether viewer has liked or recasted the cast. (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.LookupCastByHashOrWarpcastUrl(context.Background()).Identifier(identifier).Type_(type_).ViewerFid(viewerFid).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.LookupCastByHashOrWarpcastUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupCastByHashOrWarpcastUrl`: CastResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.LookupCastByHashOrWarpcastUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupCastByHashOrWarpcastUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string** | Cast identifier (Its either a url or a hash) | 
 **type_** | [**CastParamType**](CastParamType.md) |  | 
 **viewerFid** | **int32** | adds viewer_context to cast object to show whether viewer has liked or recasted the cast. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**CastResponse**](CastResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupCastConversation

> Conversation LookupCastConversation(ctx).Identifier(identifier).Type_(type_).ReplyDepth(replyDepth).IncludeChronologicalParentCasts(includeChronologicalParentCasts).ViewerFid(viewerFid).SortType(sortType).Fold(fold).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()

Conversation for a cast



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
	identifier := "https://warpcast.com/rish/0x9288c1" // string | Cast identifier (Its either a url or a hash)
	type_ := openapiclient.CastParamType("url") // CastParamType | 
	replyDepth := int32(56) // int32 | The depth of replies in the conversation that will be returned (default 2) (optional) (default to 2)
	includeChronologicalParentCasts := true // bool | Include all parent casts in chronological order (optional) (default to false)
	viewerFid := int32(3) // int32 | Providing this will return a conversation that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	sortType := openapiclient.CastConversationSortType("chron") // CastConversationSortType | Sort type for the ordering of descendants. Default is `chron` (optional)
	fold := "fold_example" // string | Show conversation above or below the fold. Lower quality responses are hidden below the fold. Not passing in a value shows the full conversation without any folding. (optional)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.LookupCastConversation(context.Background()).Identifier(identifier).Type_(type_).ReplyDepth(replyDepth).IncludeChronologicalParentCasts(includeChronologicalParentCasts).ViewerFid(viewerFid).SortType(sortType).Fold(fold).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.LookupCastConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupCastConversation`: Conversation
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.LookupCastConversation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupCastConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string** | Cast identifier (Its either a url or a hash) | 
 **type_** | [**CastParamType**](CastParamType.md) |  | 
 **replyDepth** | **int32** | The depth of replies in the conversation that will be returned (default 2) | [default to 2]
 **includeChronologicalParentCasts** | **bool** | Include all parent casts in chronological order | [default to false]
 **viewerFid** | **int32** | Providing this will return a conversation that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **sortType** | [**CastConversationSortType**](CastConversationSortType.md) | Sort type for the ordering of descendants. Default is &#x60;chron&#x60; | 
 **fold** | **string** | Show conversation above or below the fold. Lower quality responses are hidden below the fold. Not passing in a value shows the full conversation without any folding. | 
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **cursor** | **string** | Pagination cursor. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**Conversation**](Conversation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishCast

> PostCastResponse PublishCast(ctx).PostCastReqBody(postCastReqBody).Execute()

Post a cast



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
	postCastReqBody := *openapiclient.NewPostCastReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec") // PostCastReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.PublishCast(context.Background()).PostCastReqBody(postCastReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.PublishCast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishCast`: PostCastResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.PublishCast`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishCastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postCastReqBody** | [**PostCastReqBody**](PostCastReqBody.md) |  | 

### Return type

[**PostCastResponse**](PostCastResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCasts

> CastsSearchResponse SearchCasts(ctx).Q(q).Mode(mode).SortType(sortType).AuthorFid(authorFid).ViewerFid(viewerFid).ParentUrl(parentUrl).ChannelId(channelId).PriorityMode(priorityMode).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()

Search for casts



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
	q := "star (wars | trek) "space battle" after:2024-05-04" // string | Query string to search for casts. Supported operators:  | Operator  | Description                                                                                              | | --------- | -------------------------------------------------------------------------------------------------------- | | `+`       | Acts as the AND operator. This is the default operator between terms and can usually be omitted.         | | `\\|`      | Acts as the OR operator.                                                                                 | | `*`       | When used at the end of a term, signifies a prefix query.                                                  | | `\"`       | Wraps several terms into a phrase (for example, `\"star wars\"`).                                          | | `(`, `)`  | Wrap a clause for precedence (for example, `star + (wars \\| trek)`).                                     | | `~n`      | When used after a term (for example, `satr~3`), sets `fuzziness`. When used after a phrase, sets `slop`. | | `-`       | Negates the term.                                                                                        | | `before:` | Search for casts before a specific date. (e.g. `before:2025-04-20`)                                       | | `after:`  | Search for casts after a specific date. (e.g. `after:2025-04-20`)                                         | 
	mode := "literal" // string | Choices are: - `literal` - Searches for the words in the query string (default) - `semantic` - Searches for the meaning of the query string - `hybrid` - Combines both literal and semantic results  (optional)
	sortType := openapiclient.SearchSortType("desc_chron") // SearchSortType | Choices are: - `desc_chron` - All casts sorted by time (default) - `algorithmic` - Casts sorted by engagement and time  (optional)
	authorFid := int32(194) // int32 | Fid of the user whose casts you want to search (optional)
	viewerFid := int32(3) // int32 | Providing this will return search results that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	parentUrl := "parentUrl_example" // string | Parent URL of the casts you want to search (optional)
	channelId := "channelId_example" // string | Channel ID of the casts you want to search (optional)
	priorityMode := true // bool | When true, only returns search results from power badge users and users that the viewer follows (if viewer_fid is provided). (optional) (default to false)
	limit := int32(25) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.SearchCasts(context.Background()).Q(q).Mode(mode).SortType(sortType).AuthorFid(authorFid).ViewerFid(viewerFid).ParentUrl(parentUrl).ChannelId(channelId).PriorityMode(priorityMode).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.SearchCasts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCasts`: CastsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.SearchCasts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Query string to search for casts. Supported operators:  | Operator  | Description                                                                                              | | --------- | -------------------------------------------------------------------------------------------------------- | | &#x60;+&#x60;       | Acts as the AND operator. This is the default operator between terms and can usually be omitted.         | | &#x60;\\|&#x60;      | Acts as the OR operator.                                                                                 | | &#x60;*&#x60;       | When used at the end of a term, signifies a prefix query.                                                  | | &#x60;\&quot;&#x60;       | Wraps several terms into a phrase (for example, &#x60;\&quot;star wars\&quot;&#x60;).                                          | | &#x60;(&#x60;, &#x60;)&#x60;  | Wrap a clause for precedence (for example, &#x60;star + (wars \\| trek)&#x60;).                                     | | &#x60;~n&#x60;      | When used after a term (for example, &#x60;satr~3&#x60;), sets &#x60;fuzziness&#x60;. When used after a phrase, sets &#x60;slop&#x60;. | | &#x60;-&#x60;       | Negates the term.                                                                                        | | &#x60;before:&#x60; | Search for casts before a specific date. (e.g. &#x60;before:2025-04-20&#x60;)                                       | | &#x60;after:&#x60;  | Search for casts after a specific date. (e.g. &#x60;after:2025-04-20&#x60;)                                         |  | 
 **mode** | **string** | Choices are: - &#x60;literal&#x60; - Searches for the words in the query string (default) - &#x60;semantic&#x60; - Searches for the meaning of the query string - &#x60;hybrid&#x60; - Combines both literal and semantic results  | 
 **sortType** | [**SearchSortType**](SearchSortType.md) | Choices are: - &#x60;desc_chron&#x60; - All casts sorted by time (default) - &#x60;algorithmic&#x60; - Casts sorted by engagement and time  | 
 **authorFid** | **int32** | Fid of the user whose casts you want to search | 
 **viewerFid** | **int32** | Providing this will return search results that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **parentUrl** | **string** | Parent URL of the casts you want to search | 
 **channelId** | **string** | Channel ID of the casts you want to search | 
 **priorityMode** | **bool** | When true, only returns search results from power badge users and users that the viewer follows (if viewer_fid is provided). | [default to false]
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**CastsSearchResponse**](CastsSearchResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

