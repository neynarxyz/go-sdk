# \MessageAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublishMessage**](MessageAPI.md#PublishMessage) | **Post** /v1/submitMessage | Submit signed message
[**ValidateMessage**](MessageAPI.md#ValidateMessage) | **Post** /v1/validateMessage | Validate signed message



## PublishMessage

> Message PublishMessage(ctx).Body(body).Execute()

Submit signed message



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
	body := os.NewFile(1234, "some_file") // []byte | A Message is a delta operation on the Farcaster network. The message protobuf is an envelope that wraps a MessageData object and contains a hash and signature which can verify its authenticity. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.PublishMessage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.PublishMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishMessage`: Message
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.PublishMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]byte** | A Message is a delta operation on the Farcaster network. The message protobuf is an envelope that wraps a MessageData object and contains a hash and signature which can verify its authenticity.  | 

### Return type

[**Message**](Message.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateMessage

> ValidateMessageResponse ValidateMessage(ctx).Body(body).Execute()

Validate signed message



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
	body := os.NewFile(1234, "some_file") // []byte | A Message is a delta operation on the Farcaster network. The message protobuf is an envelope that wraps a MessageData object and contains a hash and signature which can verify its authenticity. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.ValidateMessage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.ValidateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateMessage`: ValidateMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.ValidateMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]byte** | A Message is a delta operation on the Farcaster network. The message protobuf is an envelope that wraps a MessageData object and contains a hash and signature which can verify its authenticity.  | 

### Return type

[**ValidateMessageResponse**](ValidateMessageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

