# \UsernamesAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchUsernameProofByName**](UsernamesAPI.md#FetchUsernameProofByName) | **Get** /v1/userNameProofByName | Proof for a username
[**FetchUsernameProofsByFid**](UsernamesAPI.md#FetchUsernameProofsByFid) | **Get** /v1/userNameProofsByFid | Proofs provided by an FID



## FetchUsernameProofByName

> UserNameProof FetchUsernameProofByName(ctx).Name(name).Execute()

Proof for a username



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
	name := "rish" // string | The Farcaster username or ENS address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsernamesAPI.FetchUsernameProofByName(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsernamesAPI.FetchUsernameProofByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUsernameProofByName`: UserNameProof
	fmt.Fprintf(os.Stdout, "Response from `UsernamesAPI.FetchUsernameProofByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUsernameProofByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The Farcaster username or ENS address | 

### Return type

[**UserNameProof**](UserNameProof.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsernameProofsByFid

> UsernameProofsResponse FetchUsernameProofsByFid(ctx).Fid(fid).Execute()

Proofs provided by an FID



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
	fid := int32(616) // int32 | The FID being requested

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsernamesAPI.FetchUsernameProofsByFid(context.Background()).Fid(fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsernamesAPI.FetchUsernameProofsByFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUsernameProofsByFid`: UsernameProofsResponse
	fmt.Fprintf(os.Stdout, "Response from `UsernamesAPI.FetchUsernameProofsByFid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUsernameProofsByFidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID being requested | 

### Return type

[**UsernameProofsResponse**](UsernameProofsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

