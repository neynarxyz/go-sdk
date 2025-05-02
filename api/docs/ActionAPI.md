# \ActionAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublishFarcasterAction**](ActionAPI.md#PublishFarcasterAction) | **Post** /farcaster/action | User actions across apps



## PublishFarcasterAction

> map[string]interface{} PublishFarcasterAction(ctx).FarcasterActionReqBody(farcasterActionReqBody).Execute()

User actions across apps



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/api"
)

func main() {
	farcasterActionReqBody := *openapiclient.NewFarcasterActionReqBody("123e4567-e89b-12d3-a456-426614174000", "https://example.com", *openapiclient.NewFarcasterActionReqBodyAction("create.user")) // FarcasterActionReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionAPI.PublishFarcasterAction(context.Background()).FarcasterActionReqBody(farcasterActionReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionAPI.PublishFarcasterAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishFarcasterAction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActionAPI.PublishFarcasterAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishFarcasterActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **farcasterActionReqBody** | [**FarcasterActionReqBody**](FarcasterActionReqBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

