# Conversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversation** | [**ConversationConversation**](ConversationConversation.md) |  | 
**Next** | Pointer to [**NextCursor**](NextCursor.md) |  | [optional] 

## Methods

### NewConversation

`func NewConversation(conversation ConversationConversation, ) *Conversation`

NewConversation instantiates a new Conversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationWithDefaults

`func NewConversationWithDefaults() *Conversation`

NewConversationWithDefaults instantiates a new Conversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversation

`func (o *Conversation) GetConversation() ConversationConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *Conversation) GetConversationOk() (*ConversationConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *Conversation) SetConversation(v ConversationConversation)`

SetConversation sets Conversation field to given value.


### GetNext

`func (o *Conversation) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Conversation) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Conversation) SetNext(v NextCursor)`

SetNext sets Next field to given value.

### HasNext

`func (o *Conversation) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


