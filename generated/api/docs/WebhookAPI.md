# \WebhookAPI

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhook**](WebhookAPI.md#DeleteWebhook) | **Delete** /farcaster/webhook | Delete a webhook
[**FetchWebhooks**](WebhookAPI.md#FetchWebhooks) | **Get** /farcaster/webhook/list | Associated webhooks of user
[**LookupWebhook**](WebhookAPI.md#LookupWebhook) | **Get** /farcaster/webhook | Fetch a webhook
[**PublishWebhook**](WebhookAPI.md#PublishWebhook) | **Post** /farcaster/webhook | Create a webhook
[**UpdateWebhook**](WebhookAPI.md#UpdateWebhook) | **Put** /farcaster/webhook | Update a webhook
[**UpdateWebhookActiveStatus**](WebhookAPI.md#UpdateWebhookActiveStatus) | **Patch** /farcaster/webhook | Update webhook status



## DeleteWebhook

> WebhookResponse DeleteWebhook(ctx).WebhookDeleteReqBody(webhookDeleteReqBody).Execute()

Delete a webhook



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
	webhookDeleteReqBody := *openapiclient.NewWebhookDeleteReqBody("WebhookId_example") // WebhookDeleteReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.DeleteWebhook(context.Background()).WebhookDeleteReqBody(webhookDeleteReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookDeleteReqBody** | [**WebhookDeleteReqBody**](WebhookDeleteReqBody.md) |  | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWebhooks

> WebhookListResponse FetchWebhooks(ctx).Execute()

Associated webhooks of user



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.FetchWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.FetchWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchWebhooks`: WebhookListResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.FetchWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchWebhooksRequest struct via the builder pattern


### Return type

[**WebhookListResponse**](WebhookListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupWebhook

> WebhookResponse LookupWebhook(ctx).WebhookId(webhookId).Execute()

Fetch a webhook



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
	webhookId := "webhookId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.LookupWebhook(context.Background()).WebhookId(webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.LookupWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.LookupWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **string** |  | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishWebhook

> WebhookResponse PublishWebhook(ctx).WebhookPostReqBody(webhookPostReqBody).Execute()

Create a webhook



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
	webhookPostReqBody := *openapiclient.NewWebhookPostReqBody("Name_example", "Url_example") // WebhookPostReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.PublishWebhook(context.Background()).WebhookPostReqBody(webhookPostReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.PublishWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.PublishWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPostReqBody** | [**WebhookPostReqBody**](WebhookPostReqBody.md) |  | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> WebhookResponse UpdateWebhook(ctx).WebhookPutReqBody(webhookPutReqBody).Execute()

Update a webhook



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
	webhookPutReqBody := *openapiclient.NewWebhookPutReqBody("Name_example", "Url_example", "WebhookId_example") // WebhookPutReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.UpdateWebhook(context.Background()).WebhookPutReqBody(webhookPutReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPutReqBody** | [**WebhookPutReqBody**](WebhookPutReqBody.md) |  | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookActiveStatus

> WebhookResponse UpdateWebhookActiveStatus(ctx).WebhookPatchReqBody(webhookPatchReqBody).Execute()

Update webhook status



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
	webhookPatchReqBody := *openapiclient.NewWebhookPatchReqBody("WebhookId_example", "Active_example") // WebhookPatchReqBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.UpdateWebhookActiveStatus(context.Background()).WebhookPatchReqBody(webhookPatchReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.UpdateWebhookActiveStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookActiveStatus`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.UpdateWebhookActiveStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookActiveStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPatchReqBody** | [**WebhookPatchReqBody**](WebhookPatchReqBody.md) |  | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

