# \ChannelAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAllChannels**](ChannelAPI.md#FetchAllChannels) | **Get** /farcaster/channel/list | Fetch all channels with their details
[**FetchBulkChannels**](ChannelAPI.md#FetchBulkChannels) | **Get** /farcaster/channel/bulk | Bulk fetch
[**FetchChannelInvites**](ChannelAPI.md#FetchChannelInvites) | **Get** /farcaster/channel/member/invite/list | Open invites
[**FetchChannelMembers**](ChannelAPI.md#FetchChannelMembers) | **Get** /farcaster/channel/member/list | Fetch members
[**FetchFollowersForAChannel**](ChannelAPI.md#FetchFollowersForAChannel) | **Get** /farcaster/channel/followers | For channel
[**FetchRelevantFollowersForAChannel**](ChannelAPI.md#FetchRelevantFollowersForAChannel) | **Get** /farcaster/channel/followers/relevant | Relevant followers
[**FetchTrendingChannels**](ChannelAPI.md#FetchTrendingChannels) | **Get** /farcaster/channel/trending | Channels by activity
[**FetchUserChannelMemberships**](ChannelAPI.md#FetchUserChannelMemberships) | **Get** /farcaster/user/memberships/list | Member of
[**FetchUserChannels**](ChannelAPI.md#FetchUserChannels) | **Get** /farcaster/user/channels | Following
[**FetchUsersActiveChannels**](ChannelAPI.md#FetchUsersActiveChannels) | **Get** /farcaster/channel/user | Fetch channels that user is active in
[**FollowChannel**](ChannelAPI.md#FollowChannel) | **Post** /farcaster/channel/follow | Follow a channel
[**InviteChannelMember**](ChannelAPI.md#InviteChannelMember) | **Post** /farcaster/channel/member/invite | Invite
[**LookupChannel**](ChannelAPI.md#LookupChannel) | **Get** /farcaster/channel | By ID or parent_url
[**RemoveChannelMember**](ChannelAPI.md#RemoveChannelMember) | **Delete** /farcaster/channel/member | Remove user
[**RespondChannelInvite**](ChannelAPI.md#RespondChannelInvite) | **Put** /farcaster/channel/member/invite | Accept or reject an invite
[**SearchChannels**](ChannelAPI.md#SearchChannels) | **Get** /farcaster/channel/search | Search by ID or name
[**UnfollowChannel**](ChannelAPI.md#UnfollowChannel) | **Delete** /farcaster/channel/follow | Unfollow a channel



## FetchAllChannels

> ChannelListResponse FetchAllChannels(ctx).Limit(limit).Cursor(cursor).Execute()

Fetch all channels with their details



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
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.FetchAllChannels(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.FetchAllChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchAllChannels`: ChannelListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.FetchAllChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**ChannelListResponse**](ChannelListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBulkChannels

> ChannelResponseBulk FetchBulkChannels(ctx).Ids(ids).Type_(type_).ViewerFid(viewerFid).Execute()

Bulk fetch



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
	ids := "neynar,warpcast" // string | Comma separated list of channel IDs or parent_urls, up to 100 at a time
	type_ := openapiclient.ChannelType("id") // ChannelType | Type of identifier being used to query the channels. Defaults to ID. (optional)
	viewerFid := int32(194) // int32 | FID of the user viewing the channels. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.FetchBulkChannels(context.Background()).Ids(ids).Type_(type_).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.FetchBulkChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchBulkChannels`: ChannelResponseBulk
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.FetchBulkChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchBulkChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma separated list of channel IDs or parent_urls, up to 100 at a time | 
 **type_** | [**ChannelType**](ChannelType.md) | Type of identifier being used to query the channels. Defaults to ID. | 
 **viewerFid** | **int32** | FID of the user viewing the channels. | 

### Return type

[**ChannelResponseBulk**](ChannelResponseBulk.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannelInvites

> ChannelMemberInviteListResponse FetchChannelInvites(ctx).ChannelId(channelId).InvitedFid(invitedFid).Limit(limit).Cursor(cursor).Execute()

Open invites



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
	channelId := "neynar" // string | Channel ID for the channel being queried (optional)
	invitedFid := int32(194) // int32 | FID of the user being invited (optional)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.FetchChannelInvites(context.Background()).ChannelId(channelId).InvitedFid(invitedFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.FetchChannelInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchChannelInvites`: ChannelMemberInviteListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.FetchChannelInvites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchChannelInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** | Channel ID for the channel being queried | 
 **invitedFid** | **int32** | FID of the user being invited | 
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**ChannelMemberInviteListResponse**](ChannelMemberInviteListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannelMembers

> ChannelMemberListResponse FetchChannelMembers(ctx).ChannelId(channelId).Fid(fid).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()

Fetch members



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
	channelId := "neynar" // string | Channel ID for the channel being queried
	fid := int32(194) // int32 | FID of the user being queried. Specify this to check if a user is a member of the channel without paginating through all members. (optional)
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.FetchChannelMembers(context.Background()).ChannelId(channelId).Fid(fid).Limit(limit).Cursor(cursor).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.FetchChannelMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchChannelMembers`: ChannelMemberListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.FetchChannelMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchChannelMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** | Channel ID for the channel being queried | 
 **fid** | **int32** | FID of the user being queried. Specify this to check if a user is a member of the channel without paginating through all members. | 
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **cursor** | **string** | Pagination cursor. | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**ChannelMemberListResponse**](ChannelMemberListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFollowersForAChannel

> UsersResponse FetchFollowersForAChannel(ctx).Id(id).ViewerFid(viewerFid).Cursor(cursor).Limit(limit).XNeynarExperimental(xNeynarExperimental).Execute()

For channel



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
	id := "founders" // string | Channel ID for the channel being queried
	viewerFid := int32(56) // int32 | Providing this will return a list of followers that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	limit := int32(30) // int32 | Number of followers to fetch (optional) (default to 25)
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.FetchFollowersForAChannel(context.Background()).Id(id).ViewerFid(viewerFid).Cursor(cursor).Limit(limit).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.FetchFollowersForAChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchFollowersForAChannel`: UsersResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.FetchFollowersForAChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFollowersForAChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Channel ID for the channel being queried | 
 **viewerFid** | **int32** | Providing this will return a list of followers that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **cursor** | **string** | Pagination cursor. | 
 **limit** | **int32** | Number of followers to fetch | [default to 25]
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


## FetchRelevantFollowersForAChannel

> RelevantFollowersResponse FetchRelevantFollowersForAChannel(ctx).Id(id).ViewerFid(viewerFid).XNeynarExperimental(xNeynarExperimental).Execute()

Relevant followers



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
	id := "neynar" // string | Channel ID being queried
	viewerFid := int32(3) // int32 | The FID of the user to customize this response for. Providing this will also return a list of followers that respects this user's mutes and blocks and includes `viewer_context`.
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.FetchRelevantFollowersForAChannel(context.Background()).Id(id).ViewerFid(viewerFid).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.FetchRelevantFollowersForAChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchRelevantFollowersForAChannel`: RelevantFollowersResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.FetchRelevantFollowersForAChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchRelevantFollowersForAChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Channel ID being queried | 
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


## FetchTrendingChannels

> TrendingChannelResponse FetchTrendingChannels(ctx).TimeWindow(timeWindow).Limit(limit).Cursor(cursor).Execute()

Channels by activity



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
	timeWindow := "timeWindow_example" // string |  (optional)
	limit := int32(10) // int32 | Number of results to fetch (optional) (default to 10)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.FetchTrendingChannels(context.Background()).TimeWindow(timeWindow).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.FetchTrendingChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchTrendingChannels`: TrendingChannelResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.FetchTrendingChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchTrendingChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeWindow** | **string** |  | 
 **limit** | **int32** | Number of results to fetch | [default to 10]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**TrendingChannelResponse**](TrendingChannelResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserChannelMemberships

> ChannelMemberListResponse FetchUserChannelMemberships(ctx).Fid(fid).Limit(limit).Cursor(cursor).Execute()

Member of



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
	fid := int32(3) // int32 | The FID of the user.
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.FetchUserChannelMemberships(context.Background()).Fid(fid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.FetchUserChannelMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserChannelMemberships`: ChannelMemberListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.FetchUserChannelMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserChannelMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID of the user. | 
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**ChannelMemberListResponse**](ChannelMemberListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserChannels

> ChannelListResponse FetchUserChannels(ctx).Fid(fid).Limit(limit).Cursor(cursor).Execute()

Following



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
	fid := int32(56) // int32 | The FID of the user.
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.FetchUserChannels(context.Background()).Fid(fid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.FetchUserChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserChannels`: ChannelListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.FetchUserChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID of the user. | 
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**ChannelListResponse**](ChannelListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsersActiveChannels

> UsersActiveChannelsResponse FetchUsersActiveChannels(ctx).Fid(fid).Limit(limit).Cursor(cursor).Execute()

Fetch channels that user is active in



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
	fid := int32(194) // int32 | The user's FID (identifier)
	limit := int32(20) // int32 | Number of results to fetch (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.FetchUsersActiveChannels(context.Background()).Fid(fid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.FetchUsersActiveChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUsersActiveChannels`: UsersActiveChannelsResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.FetchUsersActiveChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUsersActiveChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The user&#39;s FID (identifier) | 
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**UsersActiveChannelsResponse**](UsersActiveChannelsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowChannel

> OperationResponse FollowChannel(ctx).ChannelFollowReqBody(channelFollowReqBody).Execute()

Follow a channel



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
	channelFollowReqBody := *openapiclient.NewChannelFollowReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", "neynar") // ChannelFollowReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.FollowChannel(context.Background()).ChannelFollowReqBody(channelFollowReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.FollowChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowChannel`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.FollowChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFollowChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelFollowReqBody** | [**ChannelFollowReqBody**](ChannelFollowReqBody.md) |  | 

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


## InviteChannelMember

> OperationResponse InviteChannelMember(ctx).InviteChannelMemberReqBody(inviteChannelMemberReqBody).Execute()

Invite



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
	inviteChannelMemberReqBody := *openapiclient.NewInviteChannelMemberReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", "neynar", int32(3), openapiclient.ChannelMemberRole("member")) // InviteChannelMemberReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.InviteChannelMember(context.Background()).InviteChannelMemberReqBody(inviteChannelMemberReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.InviteChannelMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteChannelMember`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.InviteChannelMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteChannelMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteChannelMemberReqBody** | [**InviteChannelMemberReqBody**](InviteChannelMemberReqBody.md) |  | 

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


## LookupChannel

> ChannelResponse LookupChannel(ctx).Id(id).Type_(type_).ViewerFid(viewerFid).Execute()

By ID or parent_url



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
	id := "neynar" // string | Channel ID for the channel being queried
	type_ := openapiclient.ChannelType("id") // ChannelType | Type of identifier being used to query the channel. Defaults to ID. (optional)
	viewerFid := int32(194) // int32 | FID of the user viewing the channel. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.LookupChannel(context.Background()).Id(id).Type_(type_).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.LookupChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupChannel`: ChannelResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.LookupChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Channel ID for the channel being queried | 
 **type_** | [**ChannelType**](ChannelType.md) | Type of identifier being used to query the channel. Defaults to ID. | 
 **viewerFid** | **int32** | FID of the user viewing the channel. | 

### Return type

[**ChannelResponse**](ChannelResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveChannelMember

> OperationResponse RemoveChannelMember(ctx).RemoveChannelMemberReqBody(removeChannelMemberReqBody).Execute()

Remove user



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
	removeChannelMemberReqBody := *openapiclient.NewRemoveChannelMemberReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", "neynar", int32(3), openapiclient.ChannelMemberRole("member")) // RemoveChannelMemberReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.RemoveChannelMember(context.Background()).RemoveChannelMemberReqBody(removeChannelMemberReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.RemoveChannelMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveChannelMember`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.RemoveChannelMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveChannelMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeChannelMemberReqBody** | [**RemoveChannelMemberReqBody**](RemoveChannelMemberReqBody.md) |  | 

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


## RespondChannelInvite

> OperationResponse RespondChannelInvite(ctx).RespondChannelInviteReqBody(respondChannelInviteReqBody).Execute()

Accept or reject an invite



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
	respondChannelInviteReqBody := *openapiclient.NewRespondChannelInviteReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", "neynar", openapiclient.ChannelMemberRole("member"), false) // RespondChannelInviteReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.RespondChannelInvite(context.Background()).RespondChannelInviteReqBody(respondChannelInviteReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.RespondChannelInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RespondChannelInvite`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.RespondChannelInvite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRespondChannelInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **respondChannelInviteReqBody** | [**RespondChannelInviteReqBody**](RespondChannelInviteReqBody.md) |  | 

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


## SearchChannels

> ChannelSearchResponse SearchChannels(ctx).Q(q).Limit(limit).Cursor(cursor).Execute()

Search by ID or name



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
	q := "neynar" // string | Channel ID or name for the channel being queried
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.SearchChannels(context.Background()).Q(q).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.SearchChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchChannels`: ChannelSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.SearchChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Channel ID or name for the channel being queried | 
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**ChannelSearchResponse**](ChannelSearchResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnfollowChannel

> OperationResponse UnfollowChannel(ctx).ChannelFollowReqBody(channelFollowReqBody).Execute()

Unfollow a channel



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
	channelFollowReqBody := *openapiclient.NewChannelFollowReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", "neynar") // ChannelFollowReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.UnfollowChannel(context.Background()).ChannelFollowReqBody(channelFollowReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.UnfollowChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnfollowChannel`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.UnfollowChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelFollowReqBody** | [**ChannelFollowReqBody**](ChannelFollowReqBody.md) |  | 

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

