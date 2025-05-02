# \FnameAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IsFnameAvailable**](FnameAPI.md#IsFnameAvailable) | **Get** /farcaster/fname/availability | Check fname availability



## IsFnameAvailable

> FnameAvailabilityResponse IsFnameAvailable(ctx).Fname(fname).Execute()

Check fname availability



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/generated/api"
)

func main() {
	fname := "farcaster" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FnameAPI.IsFnameAvailable(context.Background()).Fname(fname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FnameAPI.IsFnameAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsFnameAvailable`: FnameAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `FnameAPI.IsFnameAvailable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIsFnameAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fname** | **string** |  | 

### Return type

[**FnameAvailabilityResponse**](FnameAvailabilityResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

