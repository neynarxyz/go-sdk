# MergeMessageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | [**Message**](Message.md) |  | 
**DeletedMessages** | [**[]Message**](Message.md) |  | 

## Methods

### NewMergeMessageBody

`func NewMergeMessageBody(message Message, deletedMessages []Message, ) *MergeMessageBody`

NewMergeMessageBody instantiates a new MergeMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeMessageBodyWithDefaults

`func NewMergeMessageBodyWithDefaults() *MergeMessageBody`

NewMergeMessageBodyWithDefaults instantiates a new MergeMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MergeMessageBody) GetMessage() Message`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MergeMessageBody) GetMessageOk() (*Message, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MergeMessageBody) SetMessage(v Message)`

SetMessage sets Message field to given value.


### GetDeletedMessages

`func (o *MergeMessageBody) GetDeletedMessages() []Message`

GetDeletedMessages returns the DeletedMessages field if non-nil, zero value otherwise.

### GetDeletedMessagesOk

`func (o *MergeMessageBody) GetDeletedMessagesOk() (*[]Message, bool)`

GetDeletedMessagesOk returns a tuple with the DeletedMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedMessages

`func (o *MergeMessageBody) SetDeletedMessages(v []Message)`

SetDeletedMessages sets DeletedMessages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


