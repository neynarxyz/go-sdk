# \MetricsAPI

All URIs are relative to *https://api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCastMetrics**](MetricsAPI.md#FetchCastMetrics) | **Get** /v2/farcaster/cast/metrics/ | Metrics for casts



## FetchCastMetrics

> CastsMetricsResponse FetchCastMetrics(ctx).Q(q).XNeynarExperimental(xNeynarExperimental).Interval(interval).AuthorFid(authorFid).ChannelId(channelId).Execute()

Metrics for casts



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
	q := "star wars" // string | Query string to search for casts
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)
	interval := "interval_example" // string | Interval of time for which to fetch metrics. Default is 30d. (optional)
	authorFid := int32(3) // int32 | Fid of the user whose casts you want to search (optional)
	channelId := "channelId_example" // string | Channel ID of the casts you want to search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.FetchCastMetrics(context.Background()).Q(q).XNeynarExperimental(xNeynarExperimental).Interval(interval).AuthorFid(authorFid).ChannelId(channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.FetchCastMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchCastMetrics`: CastsMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.FetchCastMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchCastMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Query string to search for casts | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]
 **interval** | **string** | Interval of time for which to fetch metrics. Default is 30d. | 
 **authorFid** | **int32** | Fid of the user whose casts you want to search | 
 **channelId** | **string** | Channel ID of the casts you want to search | 

### Return type

[**CastsMetricsResponse**](CastsMetricsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

