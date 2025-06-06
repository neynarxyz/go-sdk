# WebhookSubscriptionFiltersCast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeAuthorFids** | Pointer to **[]int32** | Exclude casts that matches these authors. **Note:** This is applied as an AND operation against rest of the filters. Rest of the filters are bundled as an OR operation. | [optional] 
**AuthorFids** | Pointer to **[]int32** |  | [optional] 
**MentionedFids** | Pointer to **[]int32** |  | [optional] 
**ParentUrls** | Pointer to **[]string** |  | [optional] 
**RootParentUrls** | Pointer to **[]string** |  | [optional] 
**ParentHashes** | Pointer to **[]string** |  | [optional] 
**ParentAuthorFids** | Pointer to **[]int32** |  | [optional] 
**Text** | Pointer to **string** | Regex pattern to match the text key of the cast. **Note:**  1) Regex must be parsed by Go&#39;s RE2 engine (Test your expression here: https://www.lddgo.net/en/string/golangregex) 2) Use backslashes to escape special characters. For example: (?i)\\\\$degen should be written as (?i)\\\\\\\\$degen | [optional] 
**Embeds** | Pointer to **string** | Regex pattern to match the embeded_url (key embeds) of the cast. **Note:**  1) Regex must be parsed by Go&#39;s RE2 engine (Test your expression here: https://www.lddgo.net/en/string/golangregex) 2) Use backslashes to escape special characters. For example: \\\\b(farcaster|neynar)\\\\b should be written as \\\\\\\\b(farcaster|neynar)\\\\\\\\b | [optional] 

## Methods

### NewWebhookSubscriptionFiltersCast

`func NewWebhookSubscriptionFiltersCast() *WebhookSubscriptionFiltersCast`

NewWebhookSubscriptionFiltersCast instantiates a new WebhookSubscriptionFiltersCast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionFiltersCastWithDefaults

`func NewWebhookSubscriptionFiltersCastWithDefaults() *WebhookSubscriptionFiltersCast`

NewWebhookSubscriptionFiltersCastWithDefaults instantiates a new WebhookSubscriptionFiltersCast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeAuthorFids

`func (o *WebhookSubscriptionFiltersCast) GetExcludeAuthorFids() []int32`

GetExcludeAuthorFids returns the ExcludeAuthorFids field if non-nil, zero value otherwise.

### GetExcludeAuthorFidsOk

`func (o *WebhookSubscriptionFiltersCast) GetExcludeAuthorFidsOk() (*[]int32, bool)`

GetExcludeAuthorFidsOk returns a tuple with the ExcludeAuthorFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAuthorFids

`func (o *WebhookSubscriptionFiltersCast) SetExcludeAuthorFids(v []int32)`

SetExcludeAuthorFids sets ExcludeAuthorFids field to given value.

### HasExcludeAuthorFids

`func (o *WebhookSubscriptionFiltersCast) HasExcludeAuthorFids() bool`

HasExcludeAuthorFids returns a boolean if a field has been set.

### GetAuthorFids

`func (o *WebhookSubscriptionFiltersCast) GetAuthorFids() []int32`

GetAuthorFids returns the AuthorFids field if non-nil, zero value otherwise.

### GetAuthorFidsOk

`func (o *WebhookSubscriptionFiltersCast) GetAuthorFidsOk() (*[]int32, bool)`

GetAuthorFidsOk returns a tuple with the AuthorFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorFids

`func (o *WebhookSubscriptionFiltersCast) SetAuthorFids(v []int32)`

SetAuthorFids sets AuthorFids field to given value.

### HasAuthorFids

`func (o *WebhookSubscriptionFiltersCast) HasAuthorFids() bool`

HasAuthorFids returns a boolean if a field has been set.

### GetMentionedFids

`func (o *WebhookSubscriptionFiltersCast) GetMentionedFids() []int32`

GetMentionedFids returns the MentionedFids field if non-nil, zero value otherwise.

### GetMentionedFidsOk

`func (o *WebhookSubscriptionFiltersCast) GetMentionedFidsOk() (*[]int32, bool)`

GetMentionedFidsOk returns a tuple with the MentionedFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedFids

`func (o *WebhookSubscriptionFiltersCast) SetMentionedFids(v []int32)`

SetMentionedFids sets MentionedFids field to given value.

### HasMentionedFids

`func (o *WebhookSubscriptionFiltersCast) HasMentionedFids() bool`

HasMentionedFids returns a boolean if a field has been set.

### GetParentUrls

`func (o *WebhookSubscriptionFiltersCast) GetParentUrls() []string`

GetParentUrls returns the ParentUrls field if non-nil, zero value otherwise.

### GetParentUrlsOk

`func (o *WebhookSubscriptionFiltersCast) GetParentUrlsOk() (*[]string, bool)`

GetParentUrlsOk returns a tuple with the ParentUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrls

`func (o *WebhookSubscriptionFiltersCast) SetParentUrls(v []string)`

SetParentUrls sets ParentUrls field to given value.

### HasParentUrls

`func (o *WebhookSubscriptionFiltersCast) HasParentUrls() bool`

HasParentUrls returns a boolean if a field has been set.

### GetRootParentUrls

`func (o *WebhookSubscriptionFiltersCast) GetRootParentUrls() []string`

GetRootParentUrls returns the RootParentUrls field if non-nil, zero value otherwise.

### GetRootParentUrlsOk

`func (o *WebhookSubscriptionFiltersCast) GetRootParentUrlsOk() (*[]string, bool)`

GetRootParentUrlsOk returns a tuple with the RootParentUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootParentUrls

`func (o *WebhookSubscriptionFiltersCast) SetRootParentUrls(v []string)`

SetRootParentUrls sets RootParentUrls field to given value.

### HasRootParentUrls

`func (o *WebhookSubscriptionFiltersCast) HasRootParentUrls() bool`

HasRootParentUrls returns a boolean if a field has been set.

### GetParentHashes

`func (o *WebhookSubscriptionFiltersCast) GetParentHashes() []string`

GetParentHashes returns the ParentHashes field if non-nil, zero value otherwise.

### GetParentHashesOk

`func (o *WebhookSubscriptionFiltersCast) GetParentHashesOk() (*[]string, bool)`

GetParentHashesOk returns a tuple with the ParentHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHashes

`func (o *WebhookSubscriptionFiltersCast) SetParentHashes(v []string)`

SetParentHashes sets ParentHashes field to given value.

### HasParentHashes

`func (o *WebhookSubscriptionFiltersCast) HasParentHashes() bool`

HasParentHashes returns a boolean if a field has been set.

### GetParentAuthorFids

`func (o *WebhookSubscriptionFiltersCast) GetParentAuthorFids() []int32`

GetParentAuthorFids returns the ParentAuthorFids field if non-nil, zero value otherwise.

### GetParentAuthorFidsOk

`func (o *WebhookSubscriptionFiltersCast) GetParentAuthorFidsOk() (*[]int32, bool)`

GetParentAuthorFidsOk returns a tuple with the ParentAuthorFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAuthorFids

`func (o *WebhookSubscriptionFiltersCast) SetParentAuthorFids(v []int32)`

SetParentAuthorFids sets ParentAuthorFids field to given value.

### HasParentAuthorFids

`func (o *WebhookSubscriptionFiltersCast) HasParentAuthorFids() bool`

HasParentAuthorFids returns a boolean if a field has been set.

### GetText

`func (o *WebhookSubscriptionFiltersCast) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WebhookSubscriptionFiltersCast) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WebhookSubscriptionFiltersCast) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WebhookSubscriptionFiltersCast) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEmbeds

`func (o *WebhookSubscriptionFiltersCast) GetEmbeds() string`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *WebhookSubscriptionFiltersCast) GetEmbedsOk() (*string, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *WebhookSubscriptionFiltersCast) SetEmbeds(v string)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *WebhookSubscriptionFiltersCast) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


