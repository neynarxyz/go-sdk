# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**WebhookId** | **string** |  | 
**DeveloperUuid** | **string** |  | 
**TargetUrl** | **string** |  | 
**Title** | **string** |  | 
**Secrets** | [**[]WebhookSecret**](WebhookSecret.md) |  | 
**Description** | **string** |  | 
**HttpTimeout** | **string** |  | 
**RateLimit** | **int32** |  | 
**Active** | **bool** |  | 
**RateLimitDuration** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **time.Time** |  | 
**Subscription** | Pointer to [**WebhookSubscription**](WebhookSubscription.md) |  | [optional] 

## Methods

### NewWebhook

`func NewWebhook(object string, webhookId string, developerUuid string, targetUrl string, title string, secrets []WebhookSecret, description string, httpTimeout string, rateLimit int32, active bool, rateLimitDuration string, createdAt time.Time, updatedAt time.Time, deletedAt time.Time, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Webhook) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Webhook) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Webhook) SetObject(v string)`

SetObject sets Object field to given value.


### GetWebhookId

`func (o *Webhook) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *Webhook) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *Webhook) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetDeveloperUuid

`func (o *Webhook) GetDeveloperUuid() string`

GetDeveloperUuid returns the DeveloperUuid field if non-nil, zero value otherwise.

### GetDeveloperUuidOk

`func (o *Webhook) GetDeveloperUuidOk() (*string, bool)`

GetDeveloperUuidOk returns a tuple with the DeveloperUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperUuid

`func (o *Webhook) SetDeveloperUuid(v string)`

SetDeveloperUuid sets DeveloperUuid field to given value.


### GetTargetUrl

`func (o *Webhook) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *Webhook) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *Webhook) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetTitle

`func (o *Webhook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Webhook) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Webhook) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSecrets

`func (o *Webhook) GetSecrets() []WebhookSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *Webhook) GetSecretsOk() (*[]WebhookSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *Webhook) SetSecrets(v []WebhookSecret)`

SetSecrets sets Secrets field to given value.


### GetDescription

`func (o *Webhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Webhook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Webhook) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHttpTimeout

`func (o *Webhook) GetHttpTimeout() string`

GetHttpTimeout returns the HttpTimeout field if non-nil, zero value otherwise.

### GetHttpTimeoutOk

`func (o *Webhook) GetHttpTimeoutOk() (*string, bool)`

GetHttpTimeoutOk returns a tuple with the HttpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpTimeout

`func (o *Webhook) SetHttpTimeout(v string)`

SetHttpTimeout sets HttpTimeout field to given value.


### GetRateLimit

`func (o *Webhook) GetRateLimit() int32`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *Webhook) GetRateLimitOk() (*int32, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *Webhook) SetRateLimit(v int32)`

SetRateLimit sets RateLimit field to given value.


### GetActive

`func (o *Webhook) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Webhook) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Webhook) SetActive(v bool)`

SetActive sets Active field to given value.


### GetRateLimitDuration

`func (o *Webhook) GetRateLimitDuration() string`

GetRateLimitDuration returns the RateLimitDuration field if non-nil, zero value otherwise.

### GetRateLimitDurationOk

`func (o *Webhook) GetRateLimitDurationOk() (*string, bool)`

GetRateLimitDurationOk returns a tuple with the RateLimitDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitDuration

`func (o *Webhook) SetRateLimitDuration(v string)`

SetRateLimitDuration sets RateLimitDuration field to given value.


### GetCreatedAt

`func (o *Webhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Webhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Webhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Webhook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Webhook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Webhook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Webhook) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Webhook) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Webhook) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetSubscription

`func (o *Webhook) GetSubscription() WebhookSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Webhook) GetSubscriptionOk() (*WebhookSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Webhook) SetSubscription(v WebhookSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *Webhook) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


