# \OnchainAPI

All URIs are relative to *https://api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployFungible**](OnchainAPI.md#DeployFungible) | **Post** /v2/fungible/ | Deploy fungible
[**FetchRelevantFungibleOwners**](OnchainAPI.md#FetchRelevantFungibleOwners) | **Get** /v2/farcaster/fungible/owner/relevant/ | Relevant owners
[**FetchUserBalance**](OnchainAPI.md#FetchUserBalance) | **Get** /v2/farcaster/user/balance/ | Token balance
[**RegisterAccountOnchain**](OnchainAPI.md#RegisterAccountOnchain) | **Post** /v2/farcaster/user/register/ | Register Farcaster account onchain
[**SendFungiblesToUsers**](OnchainAPI.md#SendFungiblesToUsers) | **Post** /v2/farcaster/fungible/send/ | Send fungibles



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
	openapiclient "github.com/neynarxyz/go-sdk/generated/neynar_sdk"
)

func main() {
	owner := "owner_example" // string | Ethereum address of the one who is creating the token
	symbol := "symbol_example" // string | Symbol/Ticker for the token
	name := "name_example" // string | Name of the token
	metadataMedia := os.NewFile(1234, "some_file") // *os.File | Media file associated with the token.  Supported formats are image/jpeg, image/gif and image/png (optional)
	metadataDescription := "metadataDescription_example" // string | Description of the token (optional)
	metadataNsfw := "metadataNsfw_example" // string | Indicates if the token is NSFW (Not Safe For Work). (optional)
	metadataWebsiteLink := "metadataWebsiteLink_example" // string | Website link related to the token (optional)
	metadataTwitter := "metadataTwitter_example" // string | Twitter profile link (optional)
	metadataDiscord := "metadataDiscord_example" // string | Discord server link (optional)
	metadataTelegram := "metadataTelegram_example" // string | Telegram link (optional)
	network := "network_example" // string | Network/Chain name (optional) (default to "base")
	factory := "factory_example" // string | Factory name - wow -> [wow.xyz](https://wow.xyz) - clanker -> [clanker.world](https://www.clanker.world) (optional) (default to "wow")

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
 **metadataMedia** | ***os.File** | Media file associated with the token.  Supported formats are image/jpeg, image/gif and image/png | 
 **metadataDescription** | **string** | Description of the token | 
 **metadataNsfw** | **string** | Indicates if the token is NSFW (Not Safe For Work). | 
 **metadataWebsiteLink** | **string** | Website link related to the token | 
 **metadataTwitter** | **string** | Twitter profile link | 
 **metadataDiscord** | **string** | Discord server link | 
 **metadataTelegram** | **string** | Telegram link | 
 **network** | **string** | Network/Chain name | [default to &quot;base&quot;]
 **factory** | **string** | Factory name - wow -&gt; [wow.xyz](https://wow.xyz) - clanker -&gt; [clanker.world](https://www.clanker.world) | [default to &quot;wow&quot;]

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

> RelevantFungibleOwnersResponse FetchRelevantFungibleOwners(ctx).ContractAddress(contractAddress).Network(network).ViewerFid(viewerFid).Execute()

Relevant owners



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
	contractAddress := "0x0db510e79909666d6dec7f5e49370838c16d950f (eth) --OR-- EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v (solana)" // string | Contract address of the fungible asset
	network := "network_example" // string | Network of the fungible asset.
	viewerFid := int32(3) // int32 | If you provide a viewer_fid, the response will include token holders from the user's network, respecting their mutes and blocks and including viewer_context; if not provided, the response will show top token holders across the network—both sets can be combined to generate a longer list if desired. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnchainAPI.FetchRelevantFungibleOwners(context.Background()).ContractAddress(contractAddress).Network(network).ViewerFid(viewerFid).Execute()
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
 **network** | **string** | Network of the fungible asset. | 
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
	openapiclient "github.com/neynarxyz/go-sdk/generated/neynar_sdk"
)

func main() {
	fid := int32(56) // int32 | FID of the user to fetch
	networks := []string{"Networks_example"} // []string | Comma separated list of networks to fetch balances for

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
 **networks** | **[]string** | Comma separated list of networks to fetch balances for | 

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


## RegisterAccountOnchain

> RegisterUserOnChainResponse RegisterAccountOnchain(ctx).RegisterUserOnChainReqBody(registerUserOnChainReqBody).Execute()

Register Farcaster account onchain



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
	registerUserOnChainReqBody := *openapiclient.NewRegisterUserOnChainReqBody(*openapiclient.NewRegisterUserOnChainReqBodyRegistration(int32(1715190000), "Signature_example", "0x5a927ac639636e534b678e81768ca19e2c6280b7", "0x5a927ac639636e534b678e81768ca19e2c6280b7")) // RegisterUserOnChainReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnchainAPI.RegisterAccountOnchain(context.Background()).RegisterUserOnChainReqBody(registerUserOnChainReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnchainAPI.RegisterAccountOnchain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAccountOnchain`: RegisterUserOnChainResponse
	fmt.Fprintf(os.Stdout, "Response from `OnchainAPI.RegisterAccountOnchain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAccountOnchainRequest struct via the builder pattern


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


## SendFungiblesToUsers

> TransactionSendFungiblesResponse SendFungiblesToUsers(ctx).XWalletId(xWalletId).TransactionSendFungiblesRequest(transactionSendFungiblesRequest).Execute()

Send fungibles



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
	xWalletId := "xWalletId_example" // string | Wallet ID to use for transactions
	transactionSendFungiblesRequest := *openapiclient.NewTransactionSendFungiblesRequest("Network_example", []openapiclient.TransactionSendFungiblesRecipient{*openapiclient.NewTransactionSendFungiblesRecipient(int32(3), float32(123))}) // TransactionSendFungiblesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnchainAPI.SendFungiblesToUsers(context.Background()).XWalletId(xWalletId).TransactionSendFungiblesRequest(transactionSendFungiblesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnchainAPI.SendFungiblesToUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendFungiblesToUsers`: TransactionSendFungiblesResponse
	fmt.Fprintf(os.Stdout, "Response from `OnchainAPI.SendFungiblesToUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendFungiblesToUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xWalletId** | **string** | Wallet ID to use for transactions | 
 **transactionSendFungiblesRequest** | [**TransactionSendFungiblesRequest**](TransactionSendFungiblesRequest.md) |  | 

### Return type

[**TransactionSendFungiblesResponse**](TransactionSendFungiblesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

