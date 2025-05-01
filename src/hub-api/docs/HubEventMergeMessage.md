# HubEventMergeMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **int32** |  | 
**MergeMessageBody** | [**MergeMessageBody**](MergeMessageBody.md) |  | 

## Methods

### NewHubEventMergeMessage

`func NewHubEventMergeMessage(type_ string, id int32, mergeMessageBody MergeMessageBody, ) *HubEventMergeMessage`

NewHubEventMergeMessage instantiates a new HubEventMergeMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubEventMergeMessageWithDefaults

`func NewHubEventMergeMessageWithDefaults() *HubEventMergeMessage`

NewHubEventMergeMessageWithDefaults instantiates a new HubEventMergeMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HubEventMergeMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HubEventMergeMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HubEventMergeMessage) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *HubEventMergeMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HubEventMergeMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HubEventMergeMessage) SetId(v int32)`

SetId sets Id field to given value.


### GetMergeMessageBody

`func (o *HubEventMergeMessage) GetMergeMessageBody() MergeMessageBody`

GetMergeMessageBody returns the MergeMessageBody field if non-nil, zero value otherwise.

### GetMergeMessageBodyOk

`func (o *HubEventMergeMessage) GetMergeMessageBodyOk() (*MergeMessageBody, bool)`

GetMergeMessageBodyOk returns a tuple with the MergeMessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeMessageBody

`func (o *HubEventMergeMessage) SetMergeMessageBody(v MergeMessageBody)`

SetMergeMessageBody sets MergeMessageBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


