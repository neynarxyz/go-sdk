# \ReactionsAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCastReactions**](ReactionsAPI.md#FetchCastReactions) | **Get** /v1/reactionsByCast | On cast
[**FetchReactionsByTarget**](ReactionsAPI.md#FetchReactionsByTarget) | **Get** /v1/reactionsByTarget | To a target URL
[**FetchUserReactions**](ReactionsAPI.md#FetchUserReactions) | **Get** /v1/reactionsByFid | By FID
[**LookupReactionById**](ReactionsAPI.md#LookupReactionById) | **Get** /v1/reactionById | By FID or cast



## FetchCastReactions

> FetchCastReactions200Response FetchCastReactions(ctx).TargetFid(targetFid).TargetHash(targetHash).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

On cast



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	targetFid := int32(56) // int32 | The FID of the cast's creator. Required to uniquely identify the cast that received the reactions. Must be used in conjunction with target_hash.
	targetHash := "targetHash_example" // string | The unique hash identifier of the cast that received the reactions. This is a 40-character hexadecimal string prefixed with '0x' that uniquely identifies the cast within the creator's posts. Must be used with target_fid.
	reactionType := openapiclient.ReactionType("Like") // ReactionType |  (default to "Like")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionsAPI.FetchCastReactions(context.Background()).TargetFid(targetFid).TargetHash(targetHash).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionsAPI.FetchCastReactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchCastReactions`: FetchCastReactions200Response
	fmt.Fprintf(os.Stdout, "Response from `ReactionsAPI.FetchCastReactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchCastReactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetFid** | **int32** | The FID of the cast&#39;s creator. Required to uniquely identify the cast that received the reactions. Must be used in conjunction with target_hash. | 
 **targetHash** | **string** | The unique hash identifier of the cast that received the reactions. This is a 40-character hexadecimal string prefixed with &#39;0x&#39; that uniquely identifies the cast within the creator&#39;s posts. Must be used with target_fid. | 
 **reactionType** | [**ReactionType**](ReactionType.md) |  | [default to &quot;Like&quot;]
 **pageSize** | **int32** | Maximum number of messages to return in a single response | 
 **reverse** | **bool** | Reverse the sort order, returning latest messages first | 
 **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FetchCastReactions200Response**](FetchCastReactions200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchReactionsByTarget

> FetchCastReactions200Response FetchReactionsByTarget(ctx).Url(url).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

To a target URL



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	url := "url_example" // string | Target URL starting with 'chain://'.
	reactionType := openapiclient.ReactionType("Like") // ReactionType |  (optional) (default to "Like")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionsAPI.FetchReactionsByTarget(context.Background()).Url(url).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionsAPI.FetchReactionsByTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchReactionsByTarget`: FetchCastReactions200Response
	fmt.Fprintf(os.Stdout, "Response from `ReactionsAPI.FetchReactionsByTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchReactionsByTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | Target URL starting with &#39;chain://&#39;. | 
 **reactionType** | [**ReactionType**](ReactionType.md) |  | [default to &quot;Like&quot;]
 **pageSize** | **int32** | Maximum number of messages to return in a single response | 
 **reverse** | **bool** | Reverse the sort order, returning latest messages first | 
 **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FetchCastReactions200Response**](FetchCastReactions200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserReactions

> FetchCastReactions200Response FetchUserReactions(ctx).Fid(fid).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

By FID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	fid := int32(56) // int32 | The FID of the reaction's creator
	reactionType := openapiclient.ReactionType("Like") // ReactionType |  (default to "Like")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionsAPI.FetchUserReactions(context.Background()).Fid(fid).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionsAPI.FetchUserReactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserReactions`: FetchCastReactions200Response
	fmt.Fprintf(os.Stdout, "Response from `ReactionsAPI.FetchUserReactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserReactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID of the reaction&#39;s creator | 
 **reactionType** | [**ReactionType**](ReactionType.md) |  | [default to &quot;Like&quot;]
 **pageSize** | **int32** | Maximum number of messages to return in a single response | 
 **reverse** | **bool** | Reverse the sort order, returning latest messages first | 
 **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FetchCastReactions200Response**](FetchCastReactions200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupReactionById

> Reaction LookupReactionById(ctx).Fid(fid).TargetFid(targetFid).TargetHash(targetHash).ReactionType(reactionType).Execute()

By FID or cast



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	fid := int32(616) // int32 | The FID of the reaction's creator
	targetFid := int32(616) // int32 | The FID of the cast's creator
	targetHash := "0xfec8fd3546e1f46cf5ad10a4ff9f5d53a3bbe9be" // string | The cast's hash
	reactionType := openapiclient.ReactionType("Like") // ReactionType |  (default to "Like")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionsAPI.LookupReactionById(context.Background()).Fid(fid).TargetFid(targetFid).TargetHash(targetHash).ReactionType(reactionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionsAPI.LookupReactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupReactionById`: Reaction
	fmt.Fprintf(os.Stdout, "Response from `ReactionsAPI.LookupReactionById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupReactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID of the reaction&#39;s creator | 
 **targetFid** | **int32** | The FID of the cast&#39;s creator | 
 **targetHash** | **string** | The cast&#39;s hash | 
 **reactionType** | [**ReactionType**](ReactionType.md) |  | [default to &quot;Like&quot;]

### Return type

[**Reaction**](Reaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

