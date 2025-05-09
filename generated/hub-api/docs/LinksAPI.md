# \LinksAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchUserFollowers**](LinksAPI.md#FetchUserFollowers) | **Get** /v1/linksByTargetFid | To target FID
[**FetchUserFollowing**](LinksAPI.md#FetchUserFollowing) | **Get** /v1/linksByFid | From source FID
[**LookupUserRelation**](LinksAPI.md#LookupUserRelation) | **Get** /v1/linkById | By its FID and target FID



## FetchUserFollowers

> FetchUserFollowers200Response FetchUserFollowers(ctx).TargetFid(targetFid).LinkType(linkType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

To target FID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/neynar_hub_sdk"
)

func main() {
	targetFid := int32(56) // int32 | The FID of the target user for this link
	linkType := openapiclient.LinkType("follow") // LinkType |  (optional) (default to "follow")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.FetchUserFollowers(context.Background()).TargetFid(targetFid).LinkType(linkType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.FetchUserFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserFollowers`: FetchUserFollowers200Response
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.FetchUserFollowers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetFid** | **int32** | The FID of the target user for this link | 
 **linkType** | [**LinkType**](LinkType.md) |  | [default to &quot;follow&quot;]
 **pageSize** | **int32** | Maximum number of messages to return in a single response | 
 **reverse** | **bool** | Reverse the sort order, returning latest messages first | 
 **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FetchUserFollowers200Response**](FetchUserFollowers200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserFollowing

> FetchUserFollowing200Response FetchUserFollowing(ctx).Fid(fid).LinkType(linkType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

From source FID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/neynar_hub_sdk"
)

func main() {
	fid := int32(56) // int32 | The FID of the link's originator
	linkType := openapiclient.LinkType("follow") // LinkType |  (optional) (default to "follow")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.FetchUserFollowing(context.Background()).Fid(fid).LinkType(linkType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.FetchUserFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserFollowing`: FetchUserFollowing200Response
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.FetchUserFollowing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserFollowingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID of the link&#39;s originator | 
 **linkType** | [**LinkType**](LinkType.md) |  | [default to &quot;follow&quot;]
 **pageSize** | **int32** | Maximum number of messages to return in a single response | 
 **reverse** | **bool** | Reverse the sort order, returning latest messages first | 
 **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FetchUserFollowing200Response**](FetchUserFollowing200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupUserRelation

> LinkAdd LookupUserRelation(ctx).Fid(fid).TargetFid(targetFid).LinkType(linkType).Execute()

By its FID and target FID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/neynar_hub_sdk"
)

func main() {
	fid := int32(56) // int32 | The FID of the link's originator
	targetFid := int32(56) // int32 | The FID of the target user for this link
	linkType := openapiclient.LinkType("follow") // LinkType |  (default to "follow")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.LookupUserRelation(context.Background()).Fid(fid).TargetFid(targetFid).LinkType(linkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.LookupUserRelation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupUserRelation`: LinkAdd
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.LookupUserRelation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupUserRelationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID of the link&#39;s originator | 
 **targetFid** | **int32** | The FID of the target user for this link | 
 **linkType** | [**LinkType**](LinkType.md) |  | [default to &quot;follow&quot;]

### Return type

[**LinkAdd**](LinkAdd.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

