# \ReactionAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteReaction**](ReactionAPI.md#DeleteReaction) | **Delete** /farcaster/reaction | Delete reaction
[**FetchCastReactions**](ReactionAPI.md#FetchCastReactions) | **Get** /farcaster/reactions/cast | Reactions for cast
[**FetchUserReactions**](ReactionAPI.md#FetchUserReactions) | **Get** /farcaster/reactions/user | Reactions for user
[**PublishReaction**](ReactionAPI.md#PublishReaction) | **Post** /farcaster/reaction | Post a reaction



## DeleteReaction

> OperationResponse DeleteReaction(ctx).ReactionReqBody(reactionReqBody).Execute()

Delete reaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	reactionReqBody := *openapiclient.NewReactionReqBody("SignerUuid_example", openapiclient.ReactionType("like"), "Target_example") // ReactionReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionAPI.DeleteReaction(context.Background()).ReactionReqBody(reactionReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionAPI.DeleteReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteReaction`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReactionAPI.DeleteReaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reactionReqBody** | [**ReactionReqBody**](ReactionReqBody.md) |  | 

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


## FetchCastReactions

> ReactionsCastResponse FetchCastReactions(ctx).Hash(hash).Types(types).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Reactions for cast



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	hash := "hash_example" // string |  (default to "0xfe90f9de682273e05b201629ad2338bdcd89b6be")
	types := []openapiclient.ReactionsType{openapiclient.ReactionsType("all")} // []ReactionsType | Customize which reaction types the request should search for. This is a comma-separated string that can include the following values: 'likes' and 'recasts'. By default api returns both. To select multiple types, use a comma-separated list of these values. 
	viewerFid := int32(56) // int32 | Providing this will return a list of reactions that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionAPI.FetchCastReactions(context.Background()).Hash(hash).Types(types).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionAPI.FetchCastReactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchCastReactions`: ReactionsCastResponse
	fmt.Fprintf(os.Stdout, "Response from `ReactionAPI.FetchCastReactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchCastReactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hash** | **string** |  | [default to &quot;0xfe90f9de682273e05b201629ad2338bdcd89b6be&quot;]
 **types** | [**[]ReactionsType**](ReactionsType.md) | Customize which reaction types the request should search for. This is a comma-separated string that can include the following values: &#39;likes&#39; and &#39;recasts&#39;. By default api returns both. To select multiple types, use a comma-separated list of these values.  | 
 **viewerFid** | **int32** | Providing this will return a list of reactions that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**ReactionsCastResponse**](ReactionsCastResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserReactions

> ReactionsResponse FetchUserReactions(ctx).Fid(fid).Type_(type_).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Reactions for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	fid := int32(3) // int32 | 
	type_ := openapiclient.ReactionsType("all") // ReactionsType | Type of reaction to fetch (likes or recasts or all)
	viewerFid := int32(3) // int32 | Providing this will return a list of reactions that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionAPI.FetchUserReactions(context.Background()).Fid(fid).Type_(type_).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionAPI.FetchUserReactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserReactions`: ReactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReactionAPI.FetchUserReactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserReactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** |  | 
 **type_** | [**ReactionsType**](ReactionsType.md) | Type of reaction to fetch (likes or recasts or all) | 
 **viewerFid** | **int32** | Providing this will return a list of reactions that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**ReactionsResponse**](ReactionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishReaction

> OperationResponse PublishReaction(ctx).ReactionReqBody(reactionReqBody).Execute()

Post a reaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	reactionReqBody := *openapiclient.NewReactionReqBody("SignerUuid_example", openapiclient.ReactionType("like"), "Target_example") // ReactionReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionAPI.PublishReaction(context.Background()).ReactionReqBody(reactionReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionAPI.PublishReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishReaction`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ReactionAPI.PublishReaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reactionReqBody** | [**ReactionReqBody**](ReactionReqBody.md) |  | 

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

