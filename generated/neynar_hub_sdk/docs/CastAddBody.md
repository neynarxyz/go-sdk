# CastAddBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmbedsDeprecated** | **[]string** |  | 
**Mentions** | **[]int32** |  | 
**ParentCastId** | Pointer to [**CastId**](CastId.md) |  | [optional] 
**ParentUrl** | Pointer to **string** |  | [optional] 
**Text** | **string** |  | 
**MentionsPositions** | **[]int64** |  | 
**Embeds** | [**[]Embed**](Embed.md) |  | 

## Methods

### NewCastAddBody

`func NewCastAddBody(embedsDeprecated []string, mentions []int32, text string, mentionsPositions []int64, embeds []Embed, ) *CastAddBody`

NewCastAddBody instantiates a new CastAddBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastAddBodyWithDefaults

`func NewCastAddBodyWithDefaults() *CastAddBody`

NewCastAddBodyWithDefaults instantiates a new CastAddBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedsDeprecated

`func (o *CastAddBody) GetEmbedsDeprecated() []string`

GetEmbedsDeprecated returns the EmbedsDeprecated field if non-nil, zero value otherwise.

### GetEmbedsDeprecatedOk

`func (o *CastAddBody) GetEmbedsDeprecatedOk() (*[]string, bool)`

GetEmbedsDeprecatedOk returns a tuple with the EmbedsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedsDeprecated

`func (o *CastAddBody) SetEmbedsDeprecated(v []string)`

SetEmbedsDeprecated sets EmbedsDeprecated field to given value.


### GetMentions

`func (o *CastAddBody) GetMentions() []int32`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *CastAddBody) GetMentionsOk() (*[]int32, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *CastAddBody) SetMentions(v []int32)`

SetMentions sets Mentions field to given value.


### GetParentCastId

`func (o *CastAddBody) GetParentCastId() CastId`

GetParentCastId returns the ParentCastId field if non-nil, zero value otherwise.

### GetParentCastIdOk

`func (o *CastAddBody) GetParentCastIdOk() (*CastId, bool)`

GetParentCastIdOk returns a tuple with the ParentCastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCastId

`func (o *CastAddBody) SetParentCastId(v CastId)`

SetParentCastId sets ParentCastId field to given value.

### HasParentCastId

`func (o *CastAddBody) HasParentCastId() bool`

HasParentCastId returns a boolean if a field has been set.

### GetParentUrl

`func (o *CastAddBody) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *CastAddBody) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *CastAddBody) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.

### HasParentUrl

`func (o *CastAddBody) HasParentUrl() bool`

HasParentUrl returns a boolean if a field has been set.

### GetText

`func (o *CastAddBody) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CastAddBody) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CastAddBody) SetText(v string)`

SetText sets Text field to given value.


### GetMentionsPositions

`func (o *CastAddBody) GetMentionsPositions() []int64`

GetMentionsPositions returns the MentionsPositions field if non-nil, zero value otherwise.

### GetMentionsPositionsOk

`func (o *CastAddBody) GetMentionsPositionsOk() (*[]int64, bool)`

GetMentionsPositionsOk returns a tuple with the MentionsPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionsPositions

`func (o *CastAddBody) SetMentionsPositions(v []int64)`

SetMentionsPositions sets MentionsPositions field to given value.


### GetEmbeds

`func (o *CastAddBody) GetEmbeds() []Embed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *CastAddBody) GetEmbedsOk() (*[]Embed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *CastAddBody) SetEmbeds(v []Embed)`

SetEmbeds sets Embeds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


