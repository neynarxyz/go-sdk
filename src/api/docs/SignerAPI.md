# \SignerAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSigner**](SignerAPI.md#CreateSigner) | **Post** /farcaster/signer | Create signer
[**FetchAuthorizationUrl**](SignerAPI.md#FetchAuthorizationUrl) | **Get** /farcaster/login/authorize | Fetch authorization url
[**FetchSigners**](SignerAPI.md#FetchSigners) | **Get** /farcaster/signer/list | List signers
[**LookupDeveloperManagedSigner**](SignerAPI.md#LookupDeveloperManagedSigner) | **Get** /farcaster/signer/developer_managed | Status by public key
[**LookupSigner**](SignerAPI.md#LookupSigner) | **Get** /farcaster/signer | Status
[**PublishMessageToFarcaster**](SignerAPI.md#PublishMessageToFarcaster) | **Post** /farcaster/message | Publish message
[**RegisterSignedKey**](SignerAPI.md#RegisterSignedKey) | **Post** /farcaster/signer/signed_key | Register Signed Key
[**RegisterSignedKeyForDeveloperManagedSigner**](SignerAPI.md#RegisterSignedKeyForDeveloperManagedSigner) | **Post** /farcaster/signer/developer_managed/signed_key | Register Signed Key



## CreateSigner

> Signer CreateSigner(ctx).Execute()

Create signer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.CreateSigner(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.CreateSigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSigner`: Signer
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.CreateSigner`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSignerRequest struct via the builder pattern


### Return type

[**Signer**](Signer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAuthorizationUrl

> AuthorizationUrlResponse FetchAuthorizationUrl(ctx).ClientId(clientId).ResponseType(responseType).Execute()

Fetch authorization url



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	clientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	responseType := openapiclient.AuthorizationUrlResponseType("code") // AuthorizationUrlResponseType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.FetchAuthorizationUrl(context.Background()).ClientId(clientId).ResponseType(responseType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.FetchAuthorizationUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchAuthorizationUrl`: AuthorizationUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.FetchAuthorizationUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchAuthorizationUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **responseType** | [**AuthorizationUrlResponseType**](AuthorizationUrlResponseType.md) |  | 

### Return type

[**AuthorizationUrlResponse**](AuthorizationUrlResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSigners

> SignerListResponse FetchSigners(ctx).Message(message).Signature(signature).Execute()

List signers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	message := "example.com%20wants%20you%20to%20sign%20in%20with%20your%20Ethereum%20account%3A%5Cn0x23A...F232%5Cn%5CnSign%20in%20to%20continue.%5Cn%5CnURI%3A%20example.com%5CnVersion%3A%201%5CnChain%20ID%3A%201%5CnNonce%3A%20xyz123%5CnIssued%20At%3A%202021-09-01T14%3A52%3A07Z" // string | A Sign-In with Ethereum (SIWE) message that the user's Ethereum wallet signs. This message includes details such as the domain, address, statement, URI, nonce, and other relevant information following the EIP-4361 standard. It should be structured and URL-encoded.  example:  example.com wants you to sign in with your Ethereum account:\\\\n0x23A...F232\\\\n\\\\nSign in to continue.\\\\n\\\\nURI: example.com\\\\nVersion: 1\\\\nChain ID: 1\\\\nNonce: xyz123\\\\nIssued At: 2021-09-01T14:52:07Z  Note: This is just an example message (So, message is invalid, since we don't want any signers related to NEYNAR_API_DOCS to be exposed).   [Checkout fetch-signers API documentation for more details.](https://docs.neynar.com/docs/fetch-signers-1) 
	signature := "0x25f8...1cf" // string | The digital signature produced by signing the provided SIWE message with the user's Ethereum private key. This signature is used to verify the authenticity of the message and the identity of the signer. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.FetchSigners(context.Background()).Message(message).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.FetchSigners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSigners`: SignerListResponse
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.FetchSigners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchSignersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **string** | A Sign-In with Ethereum (SIWE) message that the user&#39;s Ethereum wallet signs. This message includes details such as the domain, address, statement, URI, nonce, and other relevant information following the EIP-4361 standard. It should be structured and URL-encoded.  example:  example.com wants you to sign in with your Ethereum account:\\\\n0x23A...F232\\\\n\\\\nSign in to continue.\\\\n\\\\nURI: example.com\\\\nVersion: 1\\\\nChain ID: 1\\\\nNonce: xyz123\\\\nIssued At: 2021-09-01T14:52:07Z  Note: This is just an example message (So, message is invalid, since we don&#39;t want any signers related to NEYNAR_API_DOCS to be exposed).   [Checkout fetch-signers API documentation for more details.](https://docs.neynar.com/docs/fetch-signers-1)  | 
 **signature** | **string** | The digital signature produced by signing the provided SIWE message with the user&#39;s Ethereum private key. This signature is used to verify the authenticity of the message and the identity of the signer.  | 

### Return type

[**SignerListResponse**](SignerListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupDeveloperManagedSigner

> DeveloperManagedSigner LookupDeveloperManagedSigner(ctx).PublicKey(publicKey).Execute()

Status by public key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	publicKey := "0x3daa8f99c5f760688a3c9f95716ed93dee5ed5d7722d776b7c4deac957755f22" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.LookupDeveloperManagedSigner(context.Background()).PublicKey(publicKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.LookupDeveloperManagedSigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupDeveloperManagedSigner`: DeveloperManagedSigner
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.LookupDeveloperManagedSigner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupDeveloperManagedSignerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicKey** | **string** |  | 

### Return type

[**DeveloperManagedSigner**](DeveloperManagedSigner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupSigner

> Signer LookupSigner(ctx).SignerUuid(signerUuid).Execute()

Status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	signerUuid := "19d0c5fd-9b33-4a48-a0e2-bc7b0555baec" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.LookupSigner(context.Background()).SignerUuid(signerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.LookupSigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupSigner`: Signer
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.LookupSigner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupSignerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signerUuid** | **string** |  | 

### Return type

[**Signer**](Signer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishMessageToFarcaster

> map[string]interface{} PublishMessageToFarcaster(ctx).Body(body).Execute()

Publish message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.PublishMessageToFarcaster(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.PublishMessageToFarcaster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishMessageToFarcaster`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.PublishMessageToFarcaster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishMessageToFarcasterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterSignedKey

> Signer RegisterSignedKey(ctx).RegisterSignerKeyReqBody(registerSignerKeyReqBody).Execute()

Register Signed Key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	registerSignerKeyReqBody := *openapiclient.NewRegisterSignerKeyReqBody("SignerUuid_example", "Signature_example", int32(123), int32(123)) // RegisterSignerKeyReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.RegisterSignedKey(context.Background()).RegisterSignerKeyReqBody(registerSignerKeyReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.RegisterSignedKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterSignedKey`: Signer
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.RegisterSignedKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSignedKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerSignerKeyReqBody** | [**RegisterSignerKeyReqBody**](RegisterSignerKeyReqBody.md) |  | 

### Return type

[**Signer**](Signer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterSignedKeyForDeveloperManagedSigner

> DeveloperManagedSigner RegisterSignedKeyForDeveloperManagedSigner(ctx).RegisterDeveloperManagedSignedKeyReqBody(registerDeveloperManagedSignedKeyReqBody).Execute()

Register Signed Key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
	registerDeveloperManagedSignedKeyReqBody := *openapiclient.NewRegisterDeveloperManagedSignedKeyReqBody("PublicKey_example", "Signature_example", int32(123), int32(123)) // RegisterDeveloperManagedSignedKeyReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.RegisterSignedKeyForDeveloperManagedSigner(context.Background()).RegisterDeveloperManagedSignedKeyReqBody(registerDeveloperManagedSignedKeyReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.RegisterSignedKeyForDeveloperManagedSigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterSignedKeyForDeveloperManagedSigner`: DeveloperManagedSigner
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.RegisterSignedKeyForDeveloperManagedSigner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSignedKeyForDeveloperManagedSignerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerDeveloperManagedSignedKeyReqBody** | [**RegisterDeveloperManagedSignedKeyReqBody**](RegisterDeveloperManagedSignedKeyReqBody.md) |  | 

### Return type

[**DeveloperManagedSigner**](DeveloperManagedSigner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

