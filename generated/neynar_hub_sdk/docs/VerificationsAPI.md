# \VerificationsAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVerificationsByFid**](VerificationsAPI.md#FetchVerificationsByFid) | **Get** /v1/verificationsByFid | Provided by an FID



## FetchVerificationsByFid

> FetchVerificationsByFid200Response FetchVerificationsByFid(ctx).Fid(fid).Address(address).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Provided by an FID



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
	fid := int32(616) // int32 | The FID being requested
	address := "address_example" // string | The optional ETH address to filter by (optional)
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.FetchVerificationsByFid(context.Background()).Fid(fid).Address(address).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.FetchVerificationsByFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchVerificationsByFid`: FetchVerificationsByFid200Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.FetchVerificationsByFid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchVerificationsByFidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID being requested | 
 **address** | **string** | The optional ETH address to filter by | 
 **pageSize** | **int32** | Maximum number of messages to return in a single response | 
 **reverse** | **bool** | Reverse the sort order, returning latest messages first | 
 **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FetchVerificationsByFid200Response**](FetchVerificationsByFid200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

