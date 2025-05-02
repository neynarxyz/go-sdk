# \CastsAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCastsByParent**](CastsAPI.md#FetchCastsByParent) | **Get** /v1/castsByParent | By parent cast
[**FetchCastsMentioningUser**](CastsAPI.md#FetchCastsMentioningUser) | **Get** /v1/castsByMention | Mentioning an FID
[**FetchUsersCasts**](CastsAPI.md#FetchUsersCasts) | **Get** /v1/castsByFid | By FID
[**LookupCastByHashAndFid**](CastsAPI.md#LookupCastByHashAndFid) | **Get** /v1/castById | By FID and Hash



## FetchCastsByParent

> FetchCastsByParent200Response FetchCastsByParent(ctx).Fid(fid).Hash(hash).Url(url).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

By parent cast



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/hub"
)

func main() {
	fid := int32(194) // int32 | The Farcaster ID (FID) of the parent cast's creator. This parameter must be used together with the 'hash' parameter to uniquely identify a parent cast. Required only when using hash-based lookup instead of URL-based lookup. The FID is a unique identifier assigned to each Farcaster user. (optional)
	hash := "0x776593353e47dc4e7f4df3199a9b04cc8efa30d9" // string | The unique hash identifier of the parent cast. Must be used together with the 'fid' parameter when doing hash-based lookup. This is a 40-character hexadecimal string prefixed with '0x' that uniquely identifies the cast within the creator's posts. Not required if using URL-based lookup. (optional)
	url := "chain://eip155:1/erc721:0x39d89b649ffa044383333d297e325d42d31329b2" // string | Cast URL starting with 'chain://' (optional)
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastsAPI.FetchCastsByParent(context.Background()).Fid(fid).Hash(hash).Url(url).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastsAPI.FetchCastsByParent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchCastsByParent`: FetchCastsByParent200Response
	fmt.Fprintf(os.Stdout, "Response from `CastsAPI.FetchCastsByParent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchCastsByParentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The Farcaster ID (FID) of the parent cast&#39;s creator. This parameter must be used together with the &#39;hash&#39; parameter to uniquely identify a parent cast. Required only when using hash-based lookup instead of URL-based lookup. The FID is a unique identifier assigned to each Farcaster user. | 
 **hash** | **string** | The unique hash identifier of the parent cast. Must be used together with the &#39;fid&#39; parameter when doing hash-based lookup. This is a 40-character hexadecimal string prefixed with &#39;0x&#39; that uniquely identifies the cast within the creator&#39;s posts. Not required if using URL-based lookup. | 
 **url** | **string** | Cast URL starting with &#39;chain://&#39; | 
 **pageSize** | **int32** | Maximum number of messages to return in a single response | 
 **reverse** | **bool** | Reverse the sort order, returning latest messages first | 
 **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FetchCastsByParent200Response**](FetchCastsByParent200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCastsMentioningUser

> FetchCastsMentioningUser200Response FetchCastsMentioningUser(ctx).Fid(fid).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Mentioning an FID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/hub"
)

func main() {
	fid := int32(3) // int32 | The FID that is mentioned in a cast
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastsAPI.FetchCastsMentioningUser(context.Background()).Fid(fid).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastsAPI.FetchCastsMentioningUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchCastsMentioningUser`: FetchCastsMentioningUser200Response
	fmt.Fprintf(os.Stdout, "Response from `CastsAPI.FetchCastsMentioningUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchCastsMentioningUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID that is mentioned in a cast | 
 **pageSize** | **int32** | Maximum number of messages to return in a single response | 
 **reverse** | **bool** | Reverse the sort order, returning latest messages first | 
 **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FetchCastsMentioningUser200Response**](FetchCastsMentioningUser200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsersCasts

> FetchUsersCasts200Response FetchUsersCasts(ctx).Fid(fid).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

By FID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/hub"
)

func main() {
	fid := int32(56) // int32 | The FID of the casts' creator
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastsAPI.FetchUsersCasts(context.Background()).Fid(fid).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastsAPI.FetchUsersCasts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUsersCasts`: FetchUsersCasts200Response
	fmt.Fprintf(os.Stdout, "Response from `CastsAPI.FetchUsersCasts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUsersCastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID of the casts&#39; creator | 
 **pageSize** | **int32** | Maximum number of messages to return in a single response | 
 **reverse** | **bool** | Reverse the sort order, returning latest messages first | 
 **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FetchUsersCasts200Response**](FetchUsersCasts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupCastByHashAndFid

> CastAdd LookupCastByHashAndFid(ctx).Fid(fid).Hash(hash).Execute()

By FID and Hash



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/neynarxyz/go-sdk/hub"
)

func main() {
	fid := int32(56) // int32 | The FID of the cast's creator
	hash := "0x03aff391a6eb1772b20b4ead9a89f732be75fe27" // string | The unique hash identifier of the cast. This is a 40-character hexadecimal string prefixed with '0x' that uniquely identifies a specific cast in the Farcaster network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastsAPI.LookupCastByHashAndFid(context.Background()).Fid(fid).Hash(hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastsAPI.LookupCastByHashAndFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupCastByHashAndFid`: CastAdd
	fmt.Fprintf(os.Stdout, "Response from `CastsAPI.LookupCastByHashAndFid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupCastByHashAndFidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID of the cast&#39;s creator | 
 **hash** | **string** | The unique hash identifier of the cast. This is a 40-character hexadecimal string prefixed with &#39;0x&#39; that uniquely identifies a specific cast in the Farcaster network. | 

### Return type

[**CastAdd**](CastAdd.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

