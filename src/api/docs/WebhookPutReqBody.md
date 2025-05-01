# WebhookPutReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Url** | **string** |  | 
**Subscription** | Pointer to [**WebhookSubscriptionFilters**](WebhookSubscriptionFilters.md) |  | [optional] 
**WebhookId** | **string** |  | 

## Methods

### NewWebhookPutReqBody

`func NewWebhookPutReqBody(name string, url string, webhookId string, ) *WebhookPutReqBody`

NewWebhookPutReqBody instantiates a new WebhookPutReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPutReqBodyWithDefaults

`func NewWebhookPutReqBodyWithDefaults() *WebhookPutReqBody`

NewWebhookPutReqBodyWithDefaults instantiates a new WebhookPutReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookPutReqBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookPutReqBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookPutReqBody) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *WebhookPutReqBody) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookPutReqBody) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookPutReqBody) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSubscription

`func (o *WebhookPutReqBody) GetSubscription() WebhookSubscriptionFilters`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *WebhookPutReqBody) GetSubscriptionOk() (*WebhookSubscriptionFilters, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *WebhookPutReqBody) SetSubscription(v WebhookSubscriptionFilters)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *WebhookPutReqBody) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetWebhookId

`func (o *WebhookPutReqBody) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookPutReqBody) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookPutReqBody) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


