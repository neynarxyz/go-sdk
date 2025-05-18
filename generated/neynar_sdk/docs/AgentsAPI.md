# \AgentsAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransactionPayFrame**](AgentsAPI.md#CreateTransactionPayFrame) | **Post** /farcaster/frame/transaction/pay | Create transaction pay mini app
[**FetchUserInteractions**](AgentsAPI.md#FetchUserInteractions) | **Get** /farcaster/user/interactions | User interactions
[**LookupCastConversationSummary**](AgentsAPI.md#LookupCastConversationSummary) | **Get** /farcaster/cast/conversation/summary | Cast conversation summary



## CreateTransactionPayFrame

> TransactionFrameResponse CreateTransactionPayFrame(ctx).FramePayTransactionRequestBody(framePayTransactionRequestBody).Execute()

Create transaction pay mini app



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
	framePayTransactionRequestBody := *openapiclient.NewFramePayTransactionRequestBody(*openapiclient.NewFramePayTransactionRequestBodyTransaction(*openapiclient.NewTransactionFrameDestination("0x5a927ac639636e534b678e81768ca19e2c6280b7", openapiclient.Networks("ethereum"), "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913", float32(0.01))), *openapiclient.NewTransactionFrameConfig([]openapiclient.TransactionFrameLineItem{*openapiclient.NewTransactionFrameLineItem("Payment", "Payment for goods")})) // FramePayTransactionRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.CreateTransactionPayFrame(context.Background()).FramePayTransactionRequestBody(framePayTransactionRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.CreateTransactionPayFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTransactionPayFrame`: TransactionFrameResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.CreateTransactionPayFrame`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionPayFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **framePayTransactionRequestBody** | [**FramePayTransactionRequestBody**](FramePayTransactionRequestBody.md) |  | 

### Return type

[**TransactionFrameResponse**](TransactionFrameResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserInteractions

> FetchUserInteractions200Response FetchUserInteractions(ctx).Fids(fids).Type_(type_).Execute()

User interactions



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
	fids := "194, 191" // string | Comma separated list of two FIDs
	type_ := []openapiclient.NotificationType{openapiclient.NotificationType("follows")} // []NotificationType | Comma seperated list of Interaction type to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.FetchUserInteractions(context.Background()).Fids(fids).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.FetchUserInteractions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserInteractions`: FetchUserInteractions200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.FetchUserInteractions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserInteractionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fids** | **string** | Comma separated list of two FIDs | 
 **type_** | [**[]NotificationType**](NotificationType.md) | Comma seperated list of Interaction type to fetch | 

### Return type

[**FetchUserInteractions200Response**](FetchUserInteractions200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupCastConversationSummary

> ConversationSummary LookupCastConversationSummary(ctx).Identifier(identifier).Limit(limit).Prompt(prompt).Execute()

Cast conversation summary



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
	identifier := "0x9288c1e862aa72bd69d0e383a28b9a76b63cbdb4" // string | Cast identifier (Its either a url or a hash)
	limit := int32(50) // int32 | Number of casts to consider in a summary up to a point of target cast (optional) (default to 20)
	prompt := "be unreasonably dramatic" // string | Additional prompt used to generate a summary (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.LookupCastConversationSummary(context.Background()).Identifier(identifier).Limit(limit).Prompt(prompt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.LookupCastConversationSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupCastConversationSummary`: ConversationSummary
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.LookupCastConversationSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupCastConversationSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string** | Cast identifier (Its either a url or a hash) | 
 **limit** | **int32** | Number of casts to consider in a summary up to a point of target cast | [default to 20]
 **prompt** | **string** | Additional prompt used to generate a summary | 

### Return type

[**ConversationSummary**](ConversationSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

