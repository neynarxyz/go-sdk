# \InfoAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LookupHubInfo**](InfoAPI.md#LookupHubInfo) | **Get** /v1/info | Sync Methods



## LookupHubInfo

> HubInfoResponse LookupHubInfo(ctx).Dbstats(dbstats).Execute()

Sync Methods



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/hub"
)

func main() {
	dbstats := true // bool | Controls whether the response includes database statistics. When true, the response includes information about the hub's database state, storage usage, and performance metrics.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfoAPI.LookupHubInfo(context.Background()).Dbstats(dbstats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.LookupHubInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupHubInfo`: HubInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.LookupHubInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupHubInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbstats** | **bool** | Controls whether the response includes database statistics. When true, the response includes information about the hub&#39;s database state, storage usage, and performance metrics. | 

### Return type

[**HubInfoResponse**](HubInfoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

