# \UserAPI

All URIs are relative to *https://api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVerification**](UserAPI.md#DeleteVerification) | **Delete** /v2/farcaster/user/verification/ | Delete verification
[**FetchBulkUsers**](UserAPI.md#FetchBulkUsers) | **Get** /v2/farcaster/user/bulk/ | By FIDs
[**FetchBulkUsersByEthOrSolAddress**](UserAPI.md#FetchBulkUsersByEthOrSolAddress) | **Get** /v2/farcaster/user/bulk-by-address/ | By Eth or Sol addresses
[**FetchPowerUsers**](UserAPI.md#FetchPowerUsers) | **Get** /v2/farcaster/user/power/ | Power users
[**FetchPowerUsersLite**](UserAPI.md#FetchPowerUsersLite) | **Get** /v2/farcaster/user/power_lite/ | Power user FIDs
[**FetchUsersByLocation**](UserAPI.md#FetchUsersByLocation) | **Get** /v2/farcaster/user/by_location/ | By location
[**FollowUser**](UserAPI.md#FollowUser) | **Post** /v2/farcaster/user/follow/ | Follow user
[**GetFreshAccountFID**](UserAPI.md#GetFreshAccountFID) | **Get** /v2/farcaster/user/fid/ | Fetch fresh FID
[**LookupUserByCustodyAddress**](UserAPI.md#LookupUserByCustodyAddress) | **Get** /v2/farcaster/user/custody-address/ | By custody-address
[**LookupUserByUsername**](UserAPI.md#LookupUserByUsername) | **Get** /v2/farcaster/user/by_username/ | By username
[**LookupUsersByXUsername**](UserAPI.md#LookupUsersByXUsername) | **Get** /v2/farcaster/user/by_x_username/ | By X username
[**PublishVerification**](UserAPI.md#PublishVerification) | **Post** /v2/farcaster/user/verification/ | Add verification
[**RegisterAccount**](UserAPI.md#RegisterAccount) | **Post** /v2/farcaster/user/ | Register new account
[**SearchUser**](UserAPI.md#SearchUser) | **Get** /v2/farcaster/user/search/ | Search for Usernames
[**UnfollowUser**](UserAPI.md#UnfollowUser) | **Delete** /v2/farcaster/user/follow/ | Unfollow user
[**UpdateUser**](UserAPI.md#UpdateUser) | **Patch** /v2/farcaster/user/ | Update user profile



## DeleteVerification

> OperationResponse DeleteVerification(ctx).RemoveVerificationReqBody(removeVerificationReqBody).Execute()

Delete verification



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
	removeVerificationReqBody := *openapiclient.NewRemoveVerificationReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", "0x5a927ac639636e534b678e81768ca19e2c6280b7", "0x191905a9201170abb55f4c90a4cc968b44c1b71cdf3db2764b775c93e7e22b29", "0x2fc09da1f4dcb723fefb91f77932c249c418c0af00c66ed92ee1f35002c80d6a1145280c9f361d207d28447f8f7463366840d3a9309036cf6954afd1fd331beb1b") // RemoveVerificationReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.DeleteVerification(context.Background()).RemoveVerificationReqBody(removeVerificationReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVerification`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.DeleteVerification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeVerificationReqBody** | [**RemoveVerificationReqBody**](RemoveVerificationReqBody.md) |  | 

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


## FetchBulkUsers

> BulkUsersResponse FetchBulkUsers(ctx).Fids(fids).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Execute()

By FIDs



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
	fids := "194, 191, 6131" // string | Comma separated list of FIDs, up to 100 at a time
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)
	viewerFid := int32(3) // int32 | The unique identifier of a farcaster user or app (unsigned integer) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.FetchBulkUsers(context.Background()).Fids(fids).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FetchBulkUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchBulkUsers`: BulkUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.FetchBulkUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchBulkUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fids** | **string** | Comma separated list of FIDs, up to 100 at a time | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]
 **viewerFid** | **int32** | The unique identifier of a farcaster user or app (unsigned integer) | 

### Return type

[**BulkUsersResponse**](BulkUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBulkUsersByEthOrSolAddress

> BulkUsersByAddressResponse FetchBulkUsersByEthOrSolAddress(ctx).Addresses(addresses).XNeynarExperimental(xNeynarExperimental).AddressTypes(addressTypes).ViewerFid(viewerFid).Execute()

By Eth or Sol addresses



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
	addresses := "0xa6a8736f18f383f1cc2d938576933e5ea7df01a1,0x7cac817861e5c3384753403fb6c0c556c204b1ce" // string | Comma separated list of Ethereum addresses, up to 350 at a time
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)
	addressTypes := []string{"AddressTypes_example"} // []string | Customize which address types the request should search for. This is a comma-separated string that can include the following values: 'custody_address' and 'verified_address'. By default api returns both. To select multiple types, use a comma-separated list of these values. (optional)
	viewerFid := int32(3) // int32 | The unique identifier of a farcaster user or app (unsigned integer) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.FetchBulkUsersByEthOrSolAddress(context.Background()).Addresses(addresses).XNeynarExperimental(xNeynarExperimental).AddressTypes(addressTypes).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FetchBulkUsersByEthOrSolAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchBulkUsersByEthOrSolAddress`: BulkUsersByAddressResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.FetchBulkUsersByEthOrSolAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchBulkUsersByEthOrSolAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addresses** | **string** | Comma separated list of Ethereum addresses, up to 350 at a time | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]
 **addressTypes** | **[]string** | Customize which address types the request should search for. This is a comma-separated string that can include the following values: &#39;custody_address&#39; and &#39;verified_address&#39;. By default api returns both. To select multiple types, use a comma-separated list of these values. | 
 **viewerFid** | **int32** | The unique identifier of a farcaster user or app (unsigned integer) | 

### Return type

[**BulkUsersByAddressResponse**](BulkUsersByAddressResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPowerUsers

> UsersResponse FetchPowerUsers(ctx).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Power users



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
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)
	viewerFid := int32(3) // int32 | The unique identifier of a farcaster user or app (unsigned integer) (optional)
	limit := int32(10) // int32 | Number of power users to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.FetchPowerUsers(context.Background()).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FetchPowerUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPowerUsers`: UsersResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.FetchPowerUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchPowerUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]
 **viewerFid** | **int32** | The unique identifier of a farcaster user or app (unsigned integer) | 
 **limit** | **int32** | Number of power users to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor. | 

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


## FetchPowerUsersLite

> UserPowerLiteResponse FetchPowerUsersLite(ctx).XNeynarExperimental(xNeynarExperimental).Execute()

Power user FIDs



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
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.FetchPowerUsersLite(context.Background()).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FetchPowerUsersLite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPowerUsersLite`: UserPowerLiteResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.FetchPowerUsersLite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchPowerUsersLiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**UserPowerLiteResponse**](UserPowerLiteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsersByLocation

> UsersResponse FetchUsersByLocation(ctx).Latitude(latitude).Longitude(longitude).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

By location



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
	latitude := float32(37.77) // float32 | Latitude of the location
	longitude := float32(-122.41) // float32 | Longitude of the location
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)
	viewerFid := int32(3) // int32 | FID of the user viewing the feed. Providing this will return a list of users that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	limit := int32(25) // int32 | Number of results to fetch (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.FetchUsersByLocation(context.Background()).Latitude(latitude).Longitude(longitude).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FetchUsersByLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUsersByLocation`: UsersResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.FetchUsersByLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUsersByLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **latitude** | **float32** | Latitude of the location | 
 **longitude** | **float32** | Longitude of the location | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]
 **viewerFid** | **int32** | FID of the user viewing the feed. Providing this will return a list of users that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **limit** | **int32** | Number of results to fetch | [default to 25]
 **cursor** | **string** | Pagination cursor | 

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


## FollowUser

> BulkFollowResponse FollowUser(ctx).FollowReqBody(followReqBody).Execute()

Follow user



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
	followReqBody := *openapiclient.NewFollowReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", []int32{int32(3)}) // FollowReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.FollowUser(context.Background()).FollowReqBody(followReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FollowUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowUser`: BulkFollowResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.FollowUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFollowUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **followReqBody** | [**FollowReqBody**](FollowReqBody.md) |  | 

### Return type

[**BulkFollowResponse**](BulkFollowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFreshAccountFID

> UserFIDResponse GetFreshAccountFID(ctx).XNeynarExperimental(xNeynarExperimental).Execute()

Fetch fresh FID



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
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetFreshAccountFID(context.Background()).XNeynarExperimental(xNeynarExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetFreshAccountFID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFreshAccountFID`: UserFIDResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetFreshAccountFID`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFreshAccountFIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]

### Return type

[**UserFIDResponse**](UserFIDResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupUserByCustodyAddress

> UserResponse LookupUserByCustodyAddress(ctx).CustodyAddress(custodyAddress).Execute()

By custody-address



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
	custodyAddress := "0xd1b702203b1b3b641a699997746bd4a12d157909" // string | Custody Address associated with mnemonic

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.LookupUserByCustodyAddress(context.Background()).CustodyAddress(custodyAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.LookupUserByCustodyAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupUserByCustodyAddress`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.LookupUserByCustodyAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupUserByCustodyAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **custodyAddress** | **string** | Custody Address associated with mnemonic | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupUserByUsername

> UserResponse LookupUserByUsername(ctx).Username(username).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Execute()

By username



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
	username := "neynar" // string | Username of the user to fetch
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)
	viewerFid := int32(3) // int32 | The unique identifier of a farcaster user or app (unsigned integer) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.LookupUserByUsername(context.Background()).Username(username).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.LookupUserByUsername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupUserByUsername`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.LookupUserByUsername`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupUserByUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | Username of the user to fetch | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]
 **viewerFid** | **int32** | The unique identifier of a farcaster user or app (unsigned integer) | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupUsersByXUsername

> BulkUsersResponse LookupUsersByXUsername(ctx).XUsername(xUsername).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Execute()

By X username



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
	xUsername := "xUsername_example" // string | X (Twitter) username to search for, without the @ symbol
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)
	viewerFid := int32(3) // int32 | FID of the viewer for contextual information like follows and blocks (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.LookupUsersByXUsername(context.Background()).XUsername(xUsername).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.LookupUsersByXUsername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupUsersByXUsername`: BulkUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.LookupUsersByXUsername`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupUsersByXUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xUsername** | **string** | X (Twitter) username to search for, without the @ symbol | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]
 **viewerFid** | **int32** | FID of the viewer for contextual information like follows and blocks | 

### Return type

[**BulkUsersResponse**](BulkUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishVerification

> OperationResponse PublishVerification(ctx).AddVerificationReqBody(addVerificationReqBody).Execute()

Add verification



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
	addVerificationReqBody := *openapiclient.NewAddVerificationReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", "0x5a927ac639636e534b678e81768ca19e2c6280b7", "0x191905a9201170abb55f4c90a4cc968b44c1b71cdf3db2764b775c93e7e22b29", "0x2fc09da1f4dcb723fefb91f77932c249c418c0af00c66ed92ee1f35002c80d6a1145280c9f361d207d28447f8f7463366840d3a9309036cf6954afd1fd331beb1b") // AddVerificationReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.PublishVerification(context.Background()).AddVerificationReqBody(addVerificationReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PublishVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishVerification`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PublishVerification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addVerificationReqBody** | [**AddVerificationReqBody**](AddVerificationReqBody.md) |  | 

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


## RegisterAccount

> RegisterUserResponse RegisterAccount(ctx).RegisterUserReqBody(registerUserReqBody).Execute()

Register new account



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
	registerUserReqBody := *openapiclient.NewRegisterUserReqBody("Signature_example", float32(123), "RequestedUserCustodyAddress_example", float32(123)) // RegisterUserReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.RegisterAccount(context.Background()).RegisterUserReqBody(registerUserReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RegisterAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAccount`: RegisterUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.RegisterAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerUserReqBody** | [**RegisterUserReqBody**](RegisterUserReqBody.md) |  | 

### Return type

[**RegisterUserResponse**](RegisterUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUser

> UserSearchResponse SearchUser(ctx).Q(q).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Search for Usernames



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
	q := "r" // string | 
	xNeynarExperimental := true // bool | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. (optional) (default to false)
	viewerFid := int32(3) // int32 | Providing this will return search results that respects this user's mutes and blocks and includes `viewer_context`. (optional)
	limit := int32(10) // int32 | Number of users to fetch (optional) (default to 5)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.SearchUser(context.Background()).Q(q).XNeynarExperimental(xNeynarExperimental).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.SearchUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUser`: UserSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.SearchUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **xNeynarExperimental** | **bool** | Enables experimental features including filtering based on the Neynar score. See [docs](https://neynar.notion.site/Experimental-Features-1d2655195a8b80eb98b4d4ae7b76ae4a) for more details. | [default to false]
 **viewerFid** | **int32** | Providing this will return search results that respects this user&#39;s mutes and blocks and includes &#x60;viewer_context&#x60;. | 
 **limit** | **int32** | Number of users to fetch | [default to 5]
 **cursor** | **string** | Pagination cursor. | 

### Return type

[**UserSearchResponse**](UserSearchResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnfollowUser

> BulkFollowResponse UnfollowUser(ctx).FollowReqBody(followReqBody).Execute()

Unfollow user



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
	followReqBody := *openapiclient.NewFollowReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", []int32{int32(3)}) // FollowReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UnfollowUser(context.Background()).FollowReqBody(followReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UnfollowUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnfollowUser`: BulkFollowResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UnfollowUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **followReqBody** | [**FollowReqBody**](FollowReqBody.md) |  | 

### Return type

[**BulkFollowResponse**](BulkFollowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> OperationResponse UpdateUser(ctx).UpdateUserReqBody(updateUserReqBody).Execute()

Update user profile



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
	updateUserReqBody := *openapiclient.NewUpdateUserReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec") // UpdateUserReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UpdateUser(context.Background()).UpdateUserReqBody(updateUserReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserReqBody** | [**UpdateUserReqBody**](UpdateUserReqBody.md) |  | 

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

