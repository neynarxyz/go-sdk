# \FrameAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNeynarFrame**](FrameAPI.md#DeleteNeynarFrame) | **Delete** /farcaster/frame | Delete mini app
[**FetchFrameCatalog**](FrameAPI.md#FetchFrameCatalog) | **Get** /farcaster/frame/catalog | Mini apps catalog
[**FetchFrameMetaTagsFromUrl**](FrameAPI.md#FetchFrameMetaTagsFromUrl) | **Get** /farcaster/frame/crawl | Meta tags from URL
[**FetchNeynarFrames**](FrameAPI.md#FetchNeynarFrames) | **Get** /farcaster/frame/list | List of mini apps
[**FetchNotificationTokens**](FrameAPI.md#FetchNotificationTokens) | **Get** /farcaster/frame/notification_tokens | List of mini app notification tokens 
[**FetchRelevantFrames**](FrameAPI.md#FetchRelevantFrames) | **Get** /farcaster/frame/relevant | Relevant mini apps
[**FetchValidateFrameAnalytics**](FrameAPI.md#FetchValidateFrameAnalytics) | **Get** /farcaster/frame/validate/analytics | Analytics for the mini app
[**FetchValidateFrameList**](FrameAPI.md#FetchValidateFrameList) | **Get** /farcaster/frame/validate/list | All mini apps validated by user
[**GetTransactionPayFrame**](FrameAPI.md#GetTransactionPayFrame) | **Get** /farcaster/frame/transaction/pay | Get transaction pay mini app
[**LookupNeynarFrame**](FrameAPI.md#LookupNeynarFrame) | **Get** /farcaster/frame | Mini app by UUID or URL
[**PostFrameAction**](FrameAPI.md#PostFrameAction) | **Post** /farcaster/frame/action | Post a mini app action, cast action or a cast composer action
[**PostFrameActionDeveloperManaged**](FrameAPI.md#PostFrameActionDeveloperManaged) | **Post** /farcaster/frame/developer_managed/action | Signature packet
[**PublishFrameNotifications**](FrameAPI.md#PublishFrameNotifications) | **Post** /farcaster/frame/notifications | Send notifications
[**PublishNeynarFrame**](FrameAPI.md#PublishNeynarFrame) | **Post** /farcaster/frame | Create mini app
[**UpdateNeynarFrame**](FrameAPI.md#UpdateNeynarFrame) | **Put** /farcaster/frame | Update mini app
[**ValidateFrameAction**](FrameAPI.md#ValidateFrameAction) | **Post** /farcaster/frame/validate | Validate mini app action



## DeleteNeynarFrame

> DeleteFrameResponse DeleteNeynarFrame(ctx).DeleteFrameReqBody(deleteFrameReqBody).Execute()

Delete mini app



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
	deleteFrameReqBody := *openapiclient.NewDeleteFrameReqBody() // DeleteFrameReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.DeleteNeynarFrame(context.Background()).DeleteFrameReqBody(deleteFrameReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.DeleteNeynarFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNeynarFrame`: DeleteFrameResponse
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.DeleteNeynarFrame`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNeynarFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteFrameReqBody** | [**DeleteFrameReqBody**](DeleteFrameReqBody.md) |  | 

### Return type

[**DeleteFrameResponse**](DeleteFrameResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFrameCatalog

> FrameCatalogResponse FetchFrameCatalog(ctx).Limit(limit).Cursor(cursor).TimeWindow(timeWindow).Execute()

Mini apps catalog



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
	limit := int32(56) // int32 | Number of results to fetch (optional) (default to 100)
	cursor := "cursor_example" // string | Pagination cursor (optional)
	timeWindow := openapiclient.TrendingTimeWindow("1h") // TrendingTimeWindow | Time window used to calculate the change in trending score for each mini app, used to sort mini app results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.FetchFrameCatalog(context.Background()).Limit(limit).Cursor(cursor).TimeWindow(timeWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.FetchFrameCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchFrameCatalog`: FrameCatalogResponse
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.FetchFrameCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFrameCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to fetch | [default to 100]
 **cursor** | **string** | Pagination cursor | 
 **timeWindow** | [**TrendingTimeWindow**](TrendingTimeWindow.md) | Time window used to calculate the change in trending score for each mini app, used to sort mini app results | 

### Return type

[**FrameCatalogResponse**](FrameCatalogResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFrameMetaTagsFromUrl

> FetchFrameMetaTagsFromUrl200Response FetchFrameMetaTagsFromUrl(ctx).Url(url).Execute()

Meta tags from URL



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
	url := "https://frames.neynar.com/f/862277df/ff7be6a4" // string | The mini app URL to crawl

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.FetchFrameMetaTagsFromUrl(context.Background()).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.FetchFrameMetaTagsFromUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchFrameMetaTagsFromUrl`: FetchFrameMetaTagsFromUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.FetchFrameMetaTagsFromUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchFrameMetaTagsFromUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | The mini app URL to crawl | 

### Return type

[**FetchFrameMetaTagsFromUrl200Response**](FetchFrameMetaTagsFromUrl200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNeynarFrames

> []NeynarFrame FetchNeynarFrames(ctx).Execute()

List of mini apps



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.FetchNeynarFrames(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.FetchNeynarFrames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchNeynarFrames`: []NeynarFrame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.FetchNeynarFrames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchNeynarFramesRequest struct via the builder pattern


### Return type

[**[]NeynarFrame**](NeynarFrame.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNotificationTokens

> FrameNotificationTokens FetchNotificationTokens(ctx).Limit(limit).Fids(fids).Cursor(cursor).Execute()

List of mini app notification tokens 



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
	limit := int32(30) // int32 | Number of results to fetch (optional) (default to 20)
	fids := "194, 191, 6131" // string | Comma separated list of FIDs, up to 100 at a time (optional)
	cursor := "cursor_example" // string | Pagination cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.FetchNotificationTokens(context.Background()).Limit(limit).Fids(fids).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.FetchNotificationTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchNotificationTokens`: FrameNotificationTokens
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.FetchNotificationTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchNotificationTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to fetch | [default to 20]
 **fids** | **string** | Comma separated list of FIDs, up to 100 at a time | 
 **cursor** | **string** | Pagination cursor | 

### Return type

[**FrameNotificationTokens**](FrameNotificationTokens.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRelevantFrames

> FetchRelevantFrames200Response FetchRelevantFrames(ctx).ViewerFid(viewerFid).TimeWindow(timeWindow).Execute()

Relevant mini apps



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
	viewerFid := int32(56) // int32 | FID of the user to fetch relevant mini apps for
	timeWindow := openapiclient.TrendingTimeWindow("1h") // TrendingTimeWindow | Time window used to limit statistics used to calculate mini app relevance (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.FetchRelevantFrames(context.Background()).ViewerFid(viewerFid).TimeWindow(timeWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.FetchRelevantFrames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchRelevantFrames`: FetchRelevantFrames200Response
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.FetchRelevantFrames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchRelevantFramesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewerFid** | **int32** | FID of the user to fetch relevant mini apps for | 
 **timeWindow** | [**TrendingTimeWindow**](TrendingTimeWindow.md) | Time window used to limit statistics used to calculate mini app relevance | 

### Return type

[**FetchRelevantFrames200Response**](FetchRelevantFrames200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchValidateFrameAnalytics

> FrameValidateAnalyticsResponse FetchValidateFrameAnalytics(ctx).FrameUrl(frameUrl).AnalyticsType(analyticsType).Start(start).Stop(stop).AggregateWindow(aggregateWindow).Execute()

Analytics for the mini app



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	frameUrl := "https://shorturl.at/bDRY9" // string | 
	analyticsType := openapiclient.ValidateFrameAnalyticsType("total-interactors") // ValidateFrameAnalyticsType | 
	start := time.Now() // time.Time |  (default to "2024-04-06T06:44:56.811Z")
	stop := time.Now() // time.Time |  (default to "2024-04-08T06:44:56.811Z")
	aggregateWindow := openapiclient.ValidateFrameAggregateWindow("10s") // ValidateFrameAggregateWindow | Required for `analytics_type=interactions-per-cast` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.FetchValidateFrameAnalytics(context.Background()).FrameUrl(frameUrl).AnalyticsType(analyticsType).Start(start).Stop(stop).AggregateWindow(aggregateWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.FetchValidateFrameAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchValidateFrameAnalytics`: FrameValidateAnalyticsResponse
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.FetchValidateFrameAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchValidateFrameAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frameUrl** | **string** |  | 
 **analyticsType** | [**ValidateFrameAnalyticsType**](ValidateFrameAnalyticsType.md) |  | 
 **start** | **time.Time** |  | [default to &quot;2024-04-06T06:44:56.811Z&quot;]
 **stop** | **time.Time** |  | [default to &quot;2024-04-08T06:44:56.811Z&quot;]
 **aggregateWindow** | [**ValidateFrameAggregateWindow**](ValidateFrameAggregateWindow.md) | Required for &#x60;analytics_type&#x3D;interactions-per-cast&#x60; | 

### Return type

[**FrameValidateAnalyticsResponse**](FrameValidateAnalyticsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchValidateFrameList

> FrameValidateListResponse FetchValidateFrameList(ctx).Execute()

All mini apps validated by user



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.FetchValidateFrameList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.FetchValidateFrameList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchValidateFrameList`: FrameValidateListResponse
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.FetchValidateFrameList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchValidateFrameListRequest struct via the builder pattern


### Return type

[**FrameValidateListResponse**](FrameValidateListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionPayFrame

> TransactionFrameResponse GetTransactionPayFrame(ctx).Id(id).Execute()

Get transaction pay mini app



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
	id := "id_example" // string | ID of the transaction mini app to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.GetTransactionPayFrame(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.GetTransactionPayFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionPayFrame`: TransactionFrameResponse
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.GetTransactionPayFrame`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionPayFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | ID of the transaction mini app to retrieve | 

### Return type

[**TransactionFrameResponse**](TransactionFrameResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupNeynarFrame

> NeynarFrame LookupNeynarFrame(ctx).Type_(type_).Uuid(uuid).Url(url).Execute()

Mini app by UUID or URL



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
	type_ := openapiclient.FrameType("uuid") // FrameType | 
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the mini app to fetch (optional)
	url := "url_example" // string | URL of the Neynar mini app to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.LookupNeynarFrame(context.Background()).Type_(type_).Uuid(uuid).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.LookupNeynarFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupNeynarFrame`: NeynarFrame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.LookupNeynarFrame`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupNeynarFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**FrameType**](FrameType.md) |  | 
 **uuid** | **string** | UUID of the mini app to fetch | 
 **url** | **string** | URL of the Neynar mini app to fetch | 

### Return type

[**NeynarFrame**](NeynarFrame.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFrameAction

> Frame PostFrameAction(ctx).FrameActionReqBody(frameActionReqBody).Execute()

Post a mini app action, cast action or a cast composer action



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
	frameActionReqBody := *openapiclient.NewFrameActionReqBody("SignerUuid_example", *openapiclient.NewFrameAction(*openapiclient.NewFrameActionButton(int32(123), openapiclient.FrameButtonActionType("post")), "FramesUrl_example", "PostUrl_example")) // FrameActionReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.PostFrameAction(context.Background()).FrameActionReqBody(frameActionReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.PostFrameAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFrameAction`: Frame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.PostFrameAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFrameActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frameActionReqBody** | [**FrameActionReqBody**](FrameActionReqBody.md) |  | 

### Return type

[**Frame**](Frame.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFrameActionDeveloperManaged

> Frame PostFrameActionDeveloperManaged(ctx).FrameDeveloperManagedActionReqBody(frameDeveloperManagedActionReqBody).Execute()

Signature packet



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
	frameDeveloperManagedActionReqBody := *openapiclient.NewFrameDeveloperManagedActionReqBody(*openapiclient.NewFrameAction(*openapiclient.NewFrameActionButton(int32(123), openapiclient.FrameButtonActionType("post")), "FramesUrl_example", "PostUrl_example"), *openapiclient.NewFrameSignaturePacket(*openapiclient.NewFrameSignaturePacketUntrustedData(), *openapiclient.NewFrameSignaturePacketTrustedData())) // FrameDeveloperManagedActionReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.PostFrameActionDeveloperManaged(context.Background()).FrameDeveloperManagedActionReqBody(frameDeveloperManagedActionReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.PostFrameActionDeveloperManaged``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFrameActionDeveloperManaged`: Frame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.PostFrameActionDeveloperManaged`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFrameActionDeveloperManagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frameDeveloperManagedActionReqBody** | [**FrameDeveloperManagedActionReqBody**](FrameDeveloperManagedActionReqBody.md) |  | 

### Return type

[**Frame**](Frame.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishFrameNotifications

> SendFrameNotificationsResponse PublishFrameNotifications(ctx).SendFrameNotificationsReqBody(sendFrameNotificationsReqBody).Execute()

Send notifications



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
	sendFrameNotificationsReqBody := *openapiclient.NewSendFrameNotificationsReqBody([]int32{int32(123)}, *openapiclient.NewSendFrameNotificationsReqBodyNotification("New Message", "You have received a new message in your inbox.", "https://example.com/notifications")) // SendFrameNotificationsReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.PublishFrameNotifications(context.Background()).SendFrameNotificationsReqBody(sendFrameNotificationsReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.PublishFrameNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishFrameNotifications`: SendFrameNotificationsResponse
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.PublishFrameNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishFrameNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendFrameNotificationsReqBody** | [**SendFrameNotificationsReqBody**](SendFrameNotificationsReqBody.md) |  | 

### Return type

[**SendFrameNotificationsResponse**](SendFrameNotificationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishNeynarFrame

> NeynarFrame PublishNeynarFrame(ctx).NeynarFrameCreationReqBody(neynarFrameCreationReqBody).Execute()

Create mini app



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
	neynarFrameCreationReqBody := *openapiclient.NewNeynarFrameCreationReqBody("Name_example", []openapiclient.NeynarFramePage{*openapiclient.NewNeynarFramePage("Uuid_example", "Version_example", "Title_example", *openapiclient.NewNeynarPageImage("Url_example", "AspectRatio_example"))}) // NeynarFrameCreationReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.PublishNeynarFrame(context.Background()).NeynarFrameCreationReqBody(neynarFrameCreationReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.PublishNeynarFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishNeynarFrame`: NeynarFrame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.PublishNeynarFrame`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishNeynarFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **neynarFrameCreationReqBody** | [**NeynarFrameCreationReqBody**](NeynarFrameCreationReqBody.md) |  | 

### Return type

[**NeynarFrame**](NeynarFrame.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNeynarFrame

> NeynarFrame UpdateNeynarFrame(ctx).NeynarFrameUpdateReqBody(neynarFrameUpdateReqBody).Execute()

Update mini app



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
	neynarFrameUpdateReqBody := *openapiclient.NewNeynarFrameUpdateReqBody("Uuid_example", []openapiclient.NeynarFramePage{*openapiclient.NewNeynarFramePage("Uuid_example", "Version_example", "Title_example", *openapiclient.NewNeynarPageImage("Url_example", "AspectRatio_example"))}) // NeynarFrameUpdateReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.UpdateNeynarFrame(context.Background()).NeynarFrameUpdateReqBody(neynarFrameUpdateReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.UpdateNeynarFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNeynarFrame`: NeynarFrame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.UpdateNeynarFrame`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNeynarFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **neynarFrameUpdateReqBody** | [**NeynarFrameUpdateReqBody**](NeynarFrameUpdateReqBody.md) |  | 

### Return type

[**NeynarFrame**](NeynarFrame.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateFrameAction

> ValidateFrameActionResponse ValidateFrameAction(ctx).ValidateFrameActionReqBody(validateFrameActionReqBody).Execute()

Validate mini app action



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
	validateFrameActionReqBody := *openapiclient.NewValidateFrameActionReqBody("MessageBytesInHex_example") // ValidateFrameActionReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.ValidateFrameAction(context.Background()).ValidateFrameActionReqBody(validateFrameActionReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.ValidateFrameAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateFrameAction`: ValidateFrameActionResponse
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.ValidateFrameAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateFrameActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateFrameActionReqBody** | [**ValidateFrameActionReqBody**](ValidateFrameActionReqBody.md) |  | 

### Return type

[**ValidateFrameActionResponse**](ValidateFrameActionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

