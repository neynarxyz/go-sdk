# \UserDataAPI

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchUserData**](UserDataAPI.md#FetchUserData) | **Get** /v1/userDataByFid | Fetch UserData for a FID



## FetchUserData

> FetchUserData200Response FetchUserData(ctx).Fid(fid).UserDataType(userDataType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Fetch UserData for a FID



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
	fid := int32(3) // int32 | The FID that's being requested
	userDataType := openapiclient.UserDataType("USER_DATA_TYPE_PFP") // UserDataType | The type of user data, either as a numerical value or type string. If this is omitted, all user data for the FID is returned (optional) (default to "USER_DATA_TYPE_PFP")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataAPI.FetchUserData(context.Background()).Fid(fid).UserDataType(userDataType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataAPI.FetchUserData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchUserData`: FetchUserData200Response
	fmt.Fprintf(os.Stdout, "Response from `UserDataAPI.FetchUserData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchUserDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **int32** | The FID that&#39;s being requested | 
 **userDataType** | [**UserDataType**](UserDataType.md) | The type of user data, either as a numerical value or type string. If this is omitted, all user data for the FID is returned | [default to &quot;USER_DATA_TYPE_PFP&quot;]
 **pageSize** | **int32** | Maximum number of messages to return in a single response | 
 **reverse** | **bool** | Reverse the sort order, returning latest messages first | 
 **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FetchUserData200Response**](FetchUserData200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

