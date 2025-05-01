# WebhookPostReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Url** | **string** |  | 
**Subscription** | Pointer to [**WebhookSubscriptionFilters**](WebhookSubscriptionFilters.md) |  | [optional] 

## Methods

### NewWebhookPostReqBody

`func NewWebhookPostReqBody(name string, url string, ) *WebhookPostReqBody`

NewWebhookPostReqBody instantiates a new WebhookPostReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPostReqBodyWithDefaults

`func NewWebhookPostReqBodyWithDefaults() *WebhookPostReqBody`

NewWebhookPostReqBodyWithDefaults instantiates a new WebhookPostReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookPostReqBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookPostReqBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookPostReqBody) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *WebhookPostReqBody) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookPostReqBody) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookPostReqBody) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSubscription

`func (o *WebhookPostReqBody) GetSubscription() WebhookSubscriptionFilters`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *WebhookPostReqBody) GetSubscriptionOk() (*WebhookSubscriptionFilters, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *WebhookPostReqBody) SetSubscription(v WebhookSubscriptionFilters)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *WebhookPostReqBody) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


