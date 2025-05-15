# \NotificationsAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAllNotifications**](NotificationsAPI.md#FetchAllNotifications) | **Get** /farcaster/notifications | For user
[**FetchChannelNotificationsForUser**](NotificationsAPI.md#FetchChannelNotificationsForUser) | **Get** /farcaster/notifications/channel | For user by channel
[**FetchNotificationsByParentUrlForUser**](NotificationsAPI.md#FetchNotificationsByParentUrlForUser) | **Get** /farcaster/notifications/parent_url | For user by parent_urls
[**MarkNotificationsAsSeen**](NotificationsAPI.md#MarkNotificationsAsSeen) | **Post** /farcaster/notifications/seen | Mark as seen



## FetchAllNotifications

> NotificationsResponse FetchAllNotifications(ctx).Fid(fid).Type_(type_).PriorityMode(priorityMode).Limit(limit).Cursor(cursor).Execute()

For user



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
	fid := int32(194) // int32 | FID of the user you you want to fetch notifications for. The response will respect this user's mutes and blocks.
	type_ := []openapiclient.NotificationType{openapiclient.NotificationType("follows")} // []NotificationType | Notification type to fetch. Comma separated values of follows, recasts, likes, mentions, replies. (optional)
	priorityMode := true // bool | When true, only returns notifications from power badge users and users that the user follows. (optional) (default to false)
	limit := int32(15) // int32 | Number of results to fetch (optional) (default to 15)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.FetchAllNotifications(context.Background()).Fid(fid).Type_(type_).PriorityMode(priorityMode).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.FetchAllNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchAllNotifications`: NotificationsResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.FetchAllNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | FID of the user you you want to fetch notifications for. The response will respect this user&#39;s mutes and blocks. | 
 **type_** | [**[]NotificationType**](NotificationType.md) | Notification type to fetch. Comma separated values of follows, recasts, likes, mentions, replies. | 
 **priorityMode** | **bool** | When true, only returns notifications from power badge users and users that the user follows. | [default to false]
 **limit** | **int32** | Number of results to fetch | [default to 15]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**NotificationsResponse**](NotificationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannelNotificationsForUser

> NotificationsResponse FetchChannelNotificationsForUser(ctx).Fid(fid).ChannelIds(channelIds).PriorityMode(priorityMode).Limit(limit).Cursor(cursor).Execute()

For user by channel



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
	fid := int32(194) // int32 | FID of the user you you want to fetch notifications for. The response will respect this user's mutes and blocks.
	channelIds := "neynar,farcaster" // string | Comma separated channel_ids (find list of all channels here - https://docs.neynar.com/reference/list-all-channels)
	priorityMode := true // bool | When true, only returns notifications from power badge users and users that the user follows. (optional) (default to false)
	limit := int32(15) // int32 | Number of results to fetch (optional) (default to 15)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.FetchChannelNotificationsForUser(context.Background()).Fid(fid).ChannelIds(channelIds).PriorityMode(priorityMode).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.FetchChannelNotificationsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchChannelNotificationsForUser`: NotificationsResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.FetchChannelNotificationsForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchChannelNotificationsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | FID of the user you you want to fetch notifications for. The response will respect this user&#39;s mutes and blocks. | 
 **channelIds** | **string** | Comma separated channel_ids (find list of all channels here - https://docs.neynar.com/reference/list-all-channels) | 
 **priorityMode** | **bool** | When true, only returns notifications from power badge users and users that the user follows. | [default to false]
 **limit** | **int32** | Number of results to fetch | [default to 15]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**NotificationsResponse**](NotificationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNotificationsByParentUrlForUser

> NotificationsResponse FetchNotificationsByParentUrlForUser(ctx).Fid(fid).ParentUrls(parentUrls).PriorityMode(priorityMode).Limit(limit).Cursor(cursor).Execute()

For user by parent_urls



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
	fid := int32(194) // int32 | FID of the user you you want to fetch notifications for. The response will respect this user's mutes and blocks.
	parentUrls := "chain://eip155:1/erc721:0xd4498134211baad5846ce70ce04e7c4da78931cc" // string | Comma separated parent_urls
	priorityMode := true // bool | When true, only returns notifications from power badge users and users that the user follows. (optional) (default to false)
	limit := int32(15) // int32 | Number of results to fetch (optional) (default to 15)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.FetchNotificationsByParentUrlForUser(context.Background()).Fid(fid).ParentUrls(parentUrls).PriorityMode(priorityMode).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.FetchNotificationsByParentUrlForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchNotificationsByParentUrlForUser`: NotificationsResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.FetchNotificationsByParentUrlForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchNotificationsByParentUrlForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | FID of the user you you want to fetch notifications for. The response will respect this user&#39;s mutes and blocks. | 
 **parentUrls** | **string** | Comma separated parent_urls | 
 **priorityMode** | **bool** | When true, only returns notifications from power badge users and users that the user follows. | [default to false]
 **limit** | **int32** | Number of results to fetch | [default to 15]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**NotificationsResponse**](NotificationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkNotificationsAsSeen

> OperationResponse MarkNotificationsAsSeen(ctx).MarkNotificationsAsSeenReqBody(markNotificationsAsSeenReqBody).Execute()

Mark as seen



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
	markNotificationsAsSeenReqBody := *openapiclient.NewMarkNotificationsAsSeenReqBody("SignerUuid_example") // MarkNotificationsAsSeenReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.MarkNotificationsAsSeen(context.Background()).MarkNotificationsAsSeenReqBody(markNotificationsAsSeenReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.MarkNotificationsAsSeen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkNotificationsAsSeen`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.MarkNotificationsAsSeen`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkNotificationsAsSeenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markNotificationsAsSeenReqBody** | [**MarkNotificationsAsSeenReqBody**](MarkNotificationsAsSeenReqBody.md) |  | 

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

