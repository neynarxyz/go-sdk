# \OnchainAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployFungible**](OnchainAPI.md#DeployFungible) | **Post** /fungible | Deploy fungible
[**FetchRelevantFungibleOwners**](OnchainAPI.md#FetchRelevantFungibleOwners) | **Get** /farcaster/fungible/owner/relevant | Relevant owners
[**FetchUserBalance**](OnchainAPI.md#FetchUserBalance) | **Get** /farcaster/user/balance | Token balance
[**RegisterAccountOnChain**](OnchainAPI.md#RegisterAccountOnChain) | **Post** /farcaster/user/register | Register account on-chain



## DeployFungible

> DeployFungibleResponse DeployFungible(ctx).Owner(owner).Symbol(symbol).Name(name).MetadataMedia(metadataMedia).MetadataDescription(metadataDescription).MetadataNsfw(metadataNsfw).MetadataWebsiteLink(metadataWebsiteLink).MetadataTwitter(metadataTwitter).MetadataDiscord(metadataDiscord).MetadataTelegram(metadataTelegram).Network(network).Factory(factory).Execute()

Deploy fungible



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
	owner := "owner_example" // string | Ethereum address of the one who is creating the token
	symbol := "symbol_example" // string | Symbol/Ticker for the token
	name := "name_example" // string | Name of the token
	metadataMedia := os.NewFile(1234, "some_file") // *os.File | Media file associated with the token.  Supported formats are image/jpeg, image/gif and image/png  (optional)
	metadataDescription := "metadataDescription_example" // string | Description of the token (optional)
	metadataNsfw := "metadataNsfw_example" // string | Indicates if the token is NSFW (Not Safe For Work).  (optional)
	metadataWebsiteLink := "metadataWebsiteLink_example" // string | Website link related to the token (optional)
	metadataTwitter := "metadataTwitter_example" // string | Twitter profile link (optional)
	metadataDiscord := "metadataDiscord_example" // string | Discord server link (optional)
	metadataTelegram := "metadataTelegram_example" // string | Telegram link (optional)
	network := "network_example" // string | Network/Chain name (optional) (default to "base")
	factory := "factory_example" // string | Factory name - wow -> [wow.xyz](https://wow.xyz) - clanker -> [clanker.world](https://www.clanker.world)  (optional) (default to "wow")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnchainAPI.DeployFungible(context.Background()).Owner(owner).Symbol(symbol).Name(name).MetadataMedia(metadataMedia).MetadataDescription(metadataDescription).MetadataNsfw(metadataNsfw).MetadataWebsiteLink(metadataWebsiteLink).MetadataTwitter(metadataTwitter).MetadataDiscord(metadataDiscord).MetadataTelegram(metadataTelegram).Network(network).Factory(factory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnchainAPI.DeployFungible``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployFungible`: DeployFungibleResponse
	fmt.Fprintf(os.Stdout, "Response from `OnchainAPI.DeployFungible`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeployFungibleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Ethereum address of the one who is creating the token | 
 **symbol** | **string** | Symbol/Ticker for the token | 
 **name** | **string** | Name of the token | 
 **metadataMedia** | ***os.File** | Media file associated with the token.  Supported formats are image/jpeg, image/gif and image/png  | 
 **metadataDescription** | **string** | Description of the token | 
 **metadataNsfw** | **string** | Indicates if the token is NSFW (Not Safe For Work).  | 
 **metadataWebsiteLink** | **string** | Website link related to the token | 
 **metadataTwitter** | **string** | Twitter profile link | 
 **metadataDiscord** | **string** | Discord server link | 
 **metadataTelegram** | **string** | Telegram link | 
 **network** | **string** | Network/Chain name | [default to &quot;base&quot;]
 **factory** | **string** | Factory name - wow -&gt; [wow.xyz](https://wow.xyz) - clanker -&gt; [clanker.world](https://www.clanker.world)  | [default to &quot;wow&quot;]

### Return type

[**DeployFungibleResponse**](DeployFungibleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRelevantFungibleOwners

> RelevantFungibleOwnersResponse FetchRelevantFungibleOwners(ctx).ContractAddress(contractAddress).Networks(networks).ViewerFid(viewerFid).Execute()

Relevant owners



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
	contractAddress := "0x0db510e79909666d6dec7f5e49370838c16d950f" // string | Contract address of the fungible asset
	networks := []openapiclient.Networks{openapiclient.Networks("base")} // []Networks | Comma separated list of networks to fetch balances for. Currently, only \"base\" is supported.
	viewerFid := int32(194) // int32 | If you provide a viewer_fid, the response will include token holders from the user's network, respecting their mutes and blocks and including viewer_context; if not provided, the response will show top token holders across the network—both sets can be combined to generate a longer list if desired. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnchainAPI.FetchRelevantFungibleOwners(context.Background()).ContractAddress(contractAddress).Networks(networks).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnchainAPI.FetchRelevantFungibleOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchRelevantFungibleOwners`: RelevantFungibleOwnersResponse
	fmt.Fprintf(os.Stdout, "Response from `OnchainAPI.FetchRelevantFungibleOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchRelevantFungibleOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractAddress** | **string** | Contract address of the fungible asset | 
 **networks** | [**[]Networks**](Networks.md) | Comma separated list of networks to fetch balances for. Currently, only \&quot;base\&quot; is supported. | 
 **viewerFid** | **int32** | If you provide a viewer_fid, the response will include token holders from the user&#39;s network, respecting their mutes and blocks and including viewer_context; if not provided, the response will show top token holders across the network—both sets can be combined to generate a longer list if desired. | 

### Return type

[**RelevantFungibleOwnersResponse**](RelevantFungibleOwnersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserBalance

> BalanceResponse FetchUserBalance(ctx).Fid(fid).Networks(networks).Execute()

Token balance



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
	fid := int32(3) // int32 | FID of the user to fetch
	networks := []openapiclient.Networks{openapiclient.Networks("base")} // []Networks | Comma separated list of networks to fetch balances for. Currently, only \"base\" is supported.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnchainAPI.FetchUserBalance(context.Background()).Fid(fid).Networks(networks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnchainAPI.FetchUserBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserBalance`: BalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `OnchainAPI.FetchUserBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | FID of the user to fetch | 
 **networks** | [**[]Networks**](Networks.md) | Comma separated list of networks to fetch balances for. Currently, only \&quot;base\&quot; is supported. | 

### Return type

[**BalanceResponse**](BalanceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAccountOnChain

> RegisterUserOnChainResponse RegisterAccountOnChain(ctx).RegisterUserOnChainReqBody(registerUserOnChainReqBody).Execute()

Register account on-chain



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
	registerUserOnChainReqBody := *openapiclient.NewRegisterUserOnChainReqBody(*openapiclient.NewRegisterUserOnChainReqBodyRegistration(int32(1715190000), "Signature_example", "0x5a927ac639636e534b678e81768ca19e2c6280b7", "0x5a927ac639636e534b678e81768ca19e2c6280b7")) // RegisterUserOnChainReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnchainAPI.RegisterAccountOnChain(context.Background()).RegisterUserOnChainReqBody(registerUserOnChainReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnchainAPI.RegisterAccountOnChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAccountOnChain`: RegisterUserOnChainResponse
	fmt.Fprintf(os.Stdout, "Response from `OnchainAPI.RegisterAccountOnChain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAccountOnChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerUserOnChainReqBody** | [**RegisterUserOnChainReqBody**](RegisterUserOnChainReqBody.md) |  | 

### Return type

[**RegisterUserOnChainResponse**](RegisterUserOnChainResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

